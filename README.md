# Twitch status badge

[![Maintainer streaming](https://twitch-status.soren.tools/sorencodes)](https://twitch.tv/sorencodes)

This little app returns an SVG indicating a streamer's status on Twitch.

It's hosted in Google Cloud Run at https://twitch-status.soren.tools/

Simply add your username to the end of the url and it'll return either ![offline](offline.svg) or ![online](online.svg) as appropriate.

To add at to your README here on Github, you can do something like (all on one line):

```
[![Maintainer online](https://twitch-status.soren.tools/sorencodes)](https://twitch.tv/sorencodes)
```

That's an [image with an alt text](https://docs.github.com/en/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax#images) that is used as the [link "text" for the link to Twitch](https://docs.github.com/en/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax#links).
