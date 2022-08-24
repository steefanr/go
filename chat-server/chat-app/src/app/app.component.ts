import { Component, OnDestroy, OnInit } from '@angular/core';
import { of } from 'rxjs';
import { SocketService } from './socket.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.sass']
})
export class AppComponent implements OnInit, OnDestroy{
  public messages: Array<any>;
  public chatBox: string;

  public constructor(private socket: SocketService) {
    this.messages = [];
    this.chatBox = "";
  }

  public ngOnInit(): void {
    this.socket.getEventListener().subscribe(event => {
      if(event.type == "message") {
        let data = event.data.Content;
        if(event.data.sender) {
          data = event.data.sender + ": " + data;
        }
        this.messages.push(data);
      }
      if(event.type == "close") {
        this.messages.push("/The socket connection has been closed");
      }
      if(event.type == "open") {
        this.messages.push("/The socket connection has been establised");
      }
    });
  }

  public ngOnDestroy(): void {
    this.socket.close();
  }

  public send() {
    if(this.chatBox) {
      this.socket.send(this.chatBox);
      this.chatBox = "";
    }
  }

  public isSystemMessage(message: string) {
    return !!message && message.startsWith("/") ? "<strong>" + message.substring(1) + "</strong>" : message;
  }
}
