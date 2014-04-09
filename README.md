Slack Fogbugz Integration proxy
===============================

This is a quick and dirty proxy that we use to relay information from our FogBugz environment
to Slack and help coordinate our support activities.

Enjoy, pull requests are very welcome :)

You probably want to change some of the settings in the proxy to suit your needs.

Configure the channel you want to send the notification on:
	m := Message{text, "#support", "fogbugz", "http://www.fogcreek.com/images/fogbugz/pricing/kiwi.png"}
posts to #support channel as user "fogbugz" with the FogBugz kiwi bird as avatar.

Configure the port that the proxy listens on:
	port := ":10333"

Remember to use your own webhook token when starting the proxy as documented in usage help.


On the FogBugz side, use the URL Trigger plugin to post to the proxy which
will in turn relay the information to Slack.

In the URL Trigger Plugin configuration add a new trigger and configure the
URL like you normally would but point it to the proxy.

E.g. if your plugin is running at http://some.host:9090 and you want to notify
Slack users when a case is opened/resolved/closed/reactivated you could enter
http://some.host:9090//`{CaseNumber}` {EventType}: {StatusName} - <http:/office.lekane.net:9099/default.asp?{CaseNumber}|{Title}> <{PersonEditingName}>
as the URL.
You can use GET verb for the trigger.

FogBugz is a bug tracking system from FogCreek Software:
https://www.fogcreek.com/fogbugz/

Slack is a team communication and sharing environment:
https://slack.com/
