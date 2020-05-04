import React, { Component } from "react";
import "./Message.scss";

class Message extends Component {
    constructor(props) {
        super(props);
        console.log("this.prop.message");
        console.log(this.props.message);
        let temp = JSON.parse(this.props.message);
        console.log("temp=")
        console.log(temp)
        this.state = {
            message: temp
        }
    }

    render() {
        return <div className='Message'>{this.state.message.Body}</div>;
    };
}

export default Message;