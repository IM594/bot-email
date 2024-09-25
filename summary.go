package templates

import (
	"text/template"
)

type Summary struct {
	DeviceName    string
	ResourcesName string
	MeetingTitle  string
	MeetingDay    string
	MeetingTime   string
	MeetingDate   string
	Runtime       string
	Speakers      int
	WordsSpoken   int
	LogSummary    string
	SummaryLink   string
}

// func (i *Summary) GetRecipients() []string {
// 	// 这里需要实现获取收件人的逻辑
// 	return []string{}
// }

// func (i *Summary) GetTemplate() *template.Template {
// 	return summaryTemplate
// }

// func (i *Summary) SetViaEmail(viaEmail string) {
// 	// 这里需要实现设置发件人邮箱的逻辑
// }

var summaryTemplate = template.Must(template.New("Summary").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Vibe Meeting Summary</title>
    <!--[if mso]>
    <noscript>
    <xml>
        <o:OfficeDocumentSettings>
            <o:PixelsPerInch>96</o:PixelsPerInch>
        </o:OfficeDocumentSettings>
    </xml>
    </noscript>
    <![endif]-->
    <style type="text/css">
        @media only screen and (max-width: 600px) {
            .main-table {
                width: 100% !important;
            }
            .mobile-padding {
                padding-left: 10px !important;
                padding-right: 10px !important;
            }
        }
    </style>
</head>
<body style="font-family: ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, 'Noto Sans', sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol', 'Noto Color Emoji'; background-color: #fff; margin: 0; padding: 0; width: 100%;">
    <center style="width: 100%; background-color: #ffffff;">
        <!--[if mso | IE]>
        <table role="presentation" border="0" cellpadding="0" cellspacing="0" width="600" align="center" style="width:600px;">
        <tr>
        <td style="line-height:0px;font-size:0px;mso-line-height-rule:exactly;">
        <![endif]-->
        <div style="max-width: 600px; margin: 0 auto;">
            <table role="presentation" cellpadding="0" cellspacing="0" style="width: 100%; border-collapse: collapse;" class="main-table">
                <tr>
                    <td style="background-color: #000; text-align: center; height: 48px; font-size: 0; line-height: 0;">
                        <img src="https://via.placeholder.com/68x22" alt="Vibe icon" style="height: 22px; width: 68px; vertical-align: middle; display: inline-block;">
                    </td>
                </tr>
                <tr>
                    <td style="background-color: #fff; padding: 24px 24px 16px 24px; font-size: 14px; line-height: 21px; text-align: left;" class="mobile-padding">
                        A recording log has been sent to you via {{.DeviceName}} in {{.ResourcesName}}
                    </td>
                </tr>
                <tr>
                    <td style="padding: 0 24px;" class="mobile-padding">
                        <table style="width: 100%; background-color: #001BAA; color: #fff; padding: 24px; border-radius: 8px; box-shadow: 0px 4px 16px 0px #00000040;">
                            <tr>
                                <td style="width: 70%;">
                                    <div>
                                        <h2 style="margin: 0; font-size: 20px; font-weight: 400; line-height: 1.2; color: #FFFFFFE5;">{{.MeetingTitle}}</h2>
                                        <p style="margin: 4px 0 0; font-size: 16px; font-weight: 400; line-height: 1.2; color: #FFFFFFE5;">{{.MeetingDay}} {{.MeetingTime}} • {{.MeetingDate}}</p>
                                    </div>
                                    <div style="margin-top: 20px;">
                                        <a href="{{.SummaryLink}}" style="background-color: transparent; border: 1px solid #FFFFFF; color: #fff; border-radius: 50px; cursor: pointer; font-size: 14px; padding: 10px 20px; display: inline-flex; align-items: center; justify-content: center; text-decoration: none; line-height: 1;">
                                            <img src="https://via.placeholder.com/15x15" alt="Play icon" style="width: 15px; height: 15px; margin-right: 8px;">
                                            <span style="display: inline-block; vertical-align: middle;">Play</span>
                                        </a>
                                    </div>
                                </td>
                                <td style="width: 30%; vertical-align: top; text-align: right;">
                                    <img src="https://via.placeholder.com/80x80" alt="Audio icon" style="width: 80px; height: 80px; mix-blend-mode: soft-light;">
                                </td>
                            </tr>
                        </table>
                    </td>
                </tr>
                <tr>
                    <td style="padding: 32px 24px 0;" class="mobile-padding">
                        <table style="width: 100%;">
                            <tr>
                                <td style="display: flex; align-items: center;">
                                    <img src="https://via.placeholder.com/24x24" alt="Talktime Breakdown icon" style="width: 24px; height: 24px; margin-right: 4px;">
                                    <span style="font-size: 14px; line-height: 24px; color: #222222; font-weight: 500;">Talktime Breakdown</span>
                                </td>
                            </tr>
                        </table>
                    </td>
                </tr>
                <tr>
                    <td style="padding: 16px 24px 0;" class="mobile-padding">
                        <table style="width: 100%; border-spacing: 0;">
                            <tr>
                                <td style="background-color: #F8F8F8; padding: 8px 16px; border-radius: 8px; width: 33.33%;">
                                    <table style="width: 100%;">
                                        <tr>
                                            <td style="font-size: 10px; color: #8B8B8B; text-align: left;">Runtime</td>
                                            <td style="font-size: 16px; font-weight: 500; color: #222222; text-align: right;">{{.Runtime}}</td>
                                        </tr>
                                    </table>
                                </td>
                                <td style="padding: 4px;"></td>
                                <td style="background-color: #F8F8F8; padding: 8px 16px; border-radius: 8px; width: 33.33%;">
                                    <table style="width: 100%;">
                                        <tr>
                                            <td style="font-size: 10px; color: #8B8B8B; text-align: left;">Speakers</td>
                                            <td style="font-size: 16px; font-weight: 500; color: #222222; text-align: right;">{{.Speakers}}</td>
                                        </tr>
                                    </table>
                                </td>
                                <td style="padding: 4px;"></td>
                                <td style="background-color: #F8F8F8; padding: 8px 16px; border-radius: 8px; width: 33.33%;">
                                    <table style="width: 100%;">
                                        <tr>
                                            <td style="font-size: 10px; color: #8B8B8B; text-align: left;">Words spoken</td>
                                            <td style="font-size: 16px; font-weight: 500; color: #222222; text-align: right;">{{.WordsSpoken}}</td>
                                        </tr>
                                    </table>
                                </td>
                            </tr>
                        </table>
                    </td>
                </tr>
                <tr>
                    <td style="padding: 32px 24px 0;" class="mobile-padding">
                        <table style="width: 100%;">
                            <tr>
                                <td style="display: flex; align-items: center;">
                                    <img src="https://via.placeholder.com/24x24" alt="Log Summary icon" style="width: 24px; height: 24px; margin-right: 4px;">
                                    <span style="font-size: 14px; line-height: 24px; color: #222222; font-weight: 500;">Log Summary</span>
                                </td>
                            </tr>
                        </table>
                    </td>
                </tr>
                <tr>
                    <td style="padding: 16px 24px 0;" class="mobile-padding">
                        <table style="width: 100%; background-color: #f8f8f8; padding: 24px; border-radius: 8px;">
                            <tr>
                                <td style="margin: 0; font-size: 14px; font-weight: 400; line-height: 21px; text-align: left;">
                                    {{.LogSummary}}
                                </td>
                            </tr>
                            <tr>
                                <td style="padding-top: 24px;">
                                    <!--[if mso]>
                                    <v:roundrect xmlns:v="urn:schemas-microsoft-com:vml" xmlns:w="urn:schemas-microsoft-com:office:word" href="{{.SummaryLink}}" style="height:40px;v-text-anchor:middle;width:150px;" arcsize="50%" strokecolor="#405EFF" fillcolor="#405EFF">
                                    <w:anchorlock/>
                                    <center style="color:#ffffff;font-family:sans-serif;font-size:14px;font-weight:bold;">See Full Summary</center>
                                    </v:roundrect>
                                    <![endif]-->
                                    <!--[if !mso]><!-->
                                    <a href="{{.SummaryLink}}" style="background-color: #405EFF; border: 1px solid #405EFF; border-radius: 20px; color: #ffffff; display: inline-block; font-size: 14px; font-weight: bold; line-height: 40px; text-align: center; text-decoration: none; width: 150px; -webkit-text-size-adjust: none;">See Full Summary</a>
                                    <!--<![endif]-->
                                </td>
                            </tr>
                        </table>
                    </td>
                </tr>
                <tr>
                    <td style="height: 32px; font-size: 1px; line-height: 1px;">&nbsp;</td>
                </tr>
                <tr>
                    <td style="padding: 32px 40px; text-align: center; font-size: 12px; color: #FFFFFF; background-color: #111111;" class="mobile-padding">
                        <table cellpadding="0" cellspacing="0" border="0" style="width: 100%; margin-bottom: 15px;">
                            <tr>
                                <td align="center">
                                    <table cellpadding="0" cellspacing="0" border="0">
                                        <tr>
                                            <td style="padding: 0 16px;"><img src="https://via.placeholder.com/24x24" alt="Facebook icon" style="width: 24px; height: 24px; display: block;"></td>
                                            <td style="padding: 0 16px;"><img src="https://via.placeholder.com/29x24" alt="Instagram icon" style="width: 29px; height: 24px; display: block;"></td>
                                            <td style="padding: 0 16px;"><img src="https://via.placeholder.com/24x24" alt="LinkedIn icon" style="width: 24px; height: 24px; display: block;"></td>
                                            <td style="padding: 0 16px;"><img src="https://via.placeholder.com/24x24" alt="X icon" style="width: 24px; height: 24px; display: block;"></td>
                                        </tr>
                                    </table>
                                </td>
                            </tr>
                        </table>
                        <div style="margin-bottom: 16px;">
                            <p style="margin: 0;">Vibe Inc, 10400 NE 4th ST, STE 500, Bellevue, Washington 98004, USA</p>
                        </div>
                        <div style="margin-bottom: 32px;">
                            <p style="margin: 0; color: #FFFFFF;">
                                <a href="#" style="color: #FFFFFF; text-decoration: underline;">Unsubscribe</a> • 
                                <a href="#" style="color: #FFFFFF; text-decoration: underline;">Manage preferences</a>
                            </p>
                        </div>
                        <div>
                            <p style="margin: 0;">© Vibe vibe.us</p>
                        </div>
                    </td>
                </tr>
            </table>
        </div>
        <!--[if mso | IE]>
        </td>
        </tr>
        </table>
        <![endif]-->
    </center>
</body>
</html>
`))