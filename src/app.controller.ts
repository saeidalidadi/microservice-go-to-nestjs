import { Controller } from '@nestjs/common';
import { Ctx, EventPattern, Payload, RmqContext } from '@nestjs/microservices';

@Controller()
export class AppController {
  constructor() { }

  @EventPattern("test")
  getEventMessage(@Payload() data: string, @Payload() age: string, @Ctx() context: RmqContext) {
    console.log("data is -> ", data)
    console.log("age is -> ", age)
    console.log(
      "content of message is -> ",
      JSON.parse(
        context.getMessage().content.toString()
      )
    )
  }
}
