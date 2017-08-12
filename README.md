# slack-facts-bot

When one of keywords appeared in Slack message - send random fact from fact's listback to the channel

### Build and Run

Projects uses [gb](https://getgb.io/) for build application. So, before build the application run command to load dependencies

    gb vendor restore

For build application use command

    gb build

For test application use command

    gb test -v

For run application build application with `gb build`. Then fill with your data `config.xml`. Prepare data files with facts. After thant, run your bot from directory with `config.xml`.

    bin/factsbot

### Default config is for PHP facts:

`<config>
    <facts_dir>/path/to/dir/with/facts</facts_dir>
    <slack_token>[Slack Bot Token]</slack_token>
    <slack_bot_user_id>[Slack Bot User Id]</slack_bot_user_id>
</config>`

### Dependencies

- [gb](https://getgb.io/) - A project based build tool for the Go programming language.
- [nlopes/slack](https://github.com/nlopes/slack) - supports most if not all of the api.slack.com REST calls
- [randfacts-lib](https://github.com/sheremetat/randfacts-lib) - library to processing random facts
- [logrus](https://github.com/Sirupsen/logrus) - Structured, pluggable logging for Go
- [testify](https://github.com/stretchr/testify) - testing library

### MIT License

Copyright (c) 2017 Taras Sheremeta

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.