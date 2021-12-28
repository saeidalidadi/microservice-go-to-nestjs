import { Controller, Get } from '@nestjs/common';
import { Ctx, EventPattern, MessagePattern, Payload, RmqContext } from '@nestjs/microservices';
import { AppService } from './app.service';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get()
  getHello(): string {
    return this.appService.getHello();
  }

  // @MessagePattern({cmd: ""})
  // getMessage(name: string) {
  //   console.log("name is ->", name)
  // }

  @EventPattern("test")
  getEventMessage(@Payload() data: any, @Ctx() context: RmqContext) {
    console.log("name is ->", data)
    console.log("ctx is ->", context.getMessage())
    console.log("ctx is ->", context.getPattern())
    console.log("ctx is ->", JSON.parse(context.getMessage().content.toString()))
  }
}
