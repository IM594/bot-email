package templates

import (
	"text/template"
	"time"
)

type Summary struct {
	Email    string
	ViaEmail string

	StartTime         time.Time
	Runtime           int
	ParticipantsCount int
	Words             int
	Overview          []string
	Link              string
}

func (i *Summary) GetRecipients() []string {
	return []string{i.Email}
}

func (i *Summary) GetTemplate() *template.Template {
	return summaryTemplate
}

func (i *Summary) SetViaEmail(viaEmail string) {
	i.ViaEmail = viaEmail
}

var summaryTemplate = template.Must(template.New("ResetPassword").Funcs(funcMap).Parse(
	`From: "Vibe" {{if .ViaEmail}}<{{.ViaEmail}}>{{else}}<no-reply@vibe.us>{{end}}
To: {{.Email}}
Subject: Summary
MIME-Version: 1.0
Content-Type: multipart/alternative; boundary="VN9IOIHhMoPcH0icu"

--VN9IOIHhMoPcH0icu
Content-Type: text/plain; charset="utf-8"
Content-Transfer-Encoding: quoted-printable
Content-Disposition: inline

Summary...

Didn't request this email?
No worries! Your address may have been entered by mistake. You can safely ignore or delete this message.

Vibe Inc.
https://vibe.us

--VN9IOIHhMoPcH0icu
Content-Type: text/html; charset="utf-8"
Content-Disposition: inline

<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="viewport" content="width=device-width,initial-scale=1"/>
<title>Reset Password</title>
<style type="text/css">.background_main,body,div,p,table,td{font-family:'Helvetica Neue',Helvetica,Arial,sans-serif}img{border:none;-ms-interpolation-mode:bicubic;max-width:100%}p{padding-bottom:2px}body{background:#fff;font-size:17px;line-height:24px;margin:0;padding:0;-ms-text-size-adjust:100%;-webkit-text-size-adjust:100%}table{border-collapse:collapse;mso-table-lspace:0;mso-table-rspace:0;width:100%}td{font-size:17px;line-height:24px;vertical-align:top}.email_footer a,.email_footer p,.email_footer span,.email_footer td{font-size:15px;text-align:center}.email_footer td{padding-top:20px}h1,h2,h3,h4{color:#434245;font-weight:400;line-height:1.4;margin:0;margin-bottom:12px}h1{font-size:30px;line-height:36px;font-weight:900;letter-spacing:-.75px;text-align:left}ol,p,ul{font-size:17px;line-height:24px;font-weight:400;margin:0;margin-bottom:15px}ol li,p li,ul li{list-style-position:inside;margin-left:5px}a{color:#3498db;text-decoration:none}a:hover{text-decoration:underline}.button_link::after{position:absolute;content:'';top:0;right:0;bottom:0;left:0;border-radius:4px}.button_link:hover::after{box-shadow:inset 0 -2px #237c4a}.preview_text{color:transparent;display:none;height:0;max-height:0;max-width:0;opacity:0;overflow:hidden;mso-hide:all;visibility:hidden;width:0;font-size:1px;line-height:1px}.preview_text a{color:#3aa3e3!important;font-weight:700}@media only screen and (max-width:600px){table[class=background_main] .sm_full_width{width:100%!important}table[class=background_main] .sm_align_center{text-align:center!important}table[class=background_main] .sm_auto_width{width:auto!important}table[class=background_main] .sm_auto_height{height:auto!important}table[class=background_main] .sm_border_box{box-sizing:border-box!important}table[class=background_main] .sm_block{display:block!important}table[class=background_main] .sm_inline_block{display:inline-block!important}table[class=background_main] .sm_table{display:table!important}table[class=background_main] .sm_no_side_padding{padding-right:0!important;padding-left:0!important}table[class=background_main] .sm_no_border_radius{border-radius:0!important}table[class=background_main] .sm_no_padding{padding-right:0!important;padding-left:0!important}table[class=background_main] .sm_os_icons_height{height:44px}.social_img_bottom_margin{margin-bottom:20px!important}.social_p_bottom_margin{margin-bottom:40px!important}}@media all{.ExternalClass{width:100%}.ExternalClass,.ExternalClass div,.ExternalClass font,.ExternalClass p,.ExternalClass span,.ExternalClass td{line-height:100%}.email_footer a{color:#434245!important;font-family:inherit!important;font-size:inherit!important;font-weight:inherit!important;line-height:inherit!important;text-decoration:none!important}}a:hover{text-decoration:underline!important}</style>
</head>
<body>
<!--[if mso]>
    <style type="text/css">

    .background_main, table, table td, p, div, h1, h2, h3, h4, h5, h6 {
    font-family: Arial, sans-serif !important;
    }

    </style>
  <![endif]-->
<table width="100%" cellpadding="0" cellspacing="0" border="0" class="background_main" style="background-color:#fff;padding-top:20px;color:#434245;width:100%;-webkit-font-smoothing:antialiased;-moz-osx-font-smoothing:grayscale">
<tr>
<td valign="top" class="sm_full_width" style="margin:0 auto;width:100%;max-width:600px;display:block">
<div class="sm_no_padding" style="margin:0 auto;padding:30px 0 40px;display:block;box-sizing:border-box">
<table style="width:100%;color:#434245" border="0" cellpadding="0" cellspacing="0">
<tr>
<td style="box-sizing:border-box">
<table border="0" cellpadding="0" cellspacing="0">
<tr>
<td>
<!--[if mso]>
              <table cellpadding="0" cellspacing="0" border="0" style="padding: 0; margin: 0; width: 100%;">
              <tr>
              <td colspan="3" style="padding: 0; margin: 0; font-size: 20px; height: 20px;" height="20">&nbsp;</td>
              </tr>
              <tr>
              <td style="padding: 0; margin: 0;">&nbsp;</td>
              <td style="padding: 0; margin: 0;" width="540">
            <![endif]-->
<img style="width:38px;height:38px;margin:0 0 15px 0;padding-right:30px;padding-left:30px" alt="" width="38" height="38" src="https://app.vibe.us/static/media/vibe-icon-128.png"/>
<br>
<br>
<h1 style="font-size:30px;padding-right:30px;padding-left:30px">Summary</h1>
<div style="padding-right:30px;padding-left:30px;margin-top:40px;margin-bottom:30px">
<table style="width:100%" class="sm_table">
<tr style="width:100%">
<p style="font-size:15px;padding-right:30px;padding-left:0px;padding-top:8px;">start time: {{.StartTime}}</p>
</tr>
<tr style="width:100%">
<p style="font-size:15px;padding-right:30px;padding-left:0px;padding-top:8px;">runtime: {{.Runtime}} seconds</p>
</tr>
<tr style="width:100%">
<p style="font-size:15px;padding-right:30px;padding-left:0px;padding-top:8px;">participants: {{.ParticipantsCount}}</p>
</tr>
<tr style="width:100%">
<p style="font-size:15px;padding-right:30px;padding-left:0px;padding-top:8px;">words: {{.Words}}</p>
</tr>
<tr style="width:100%">
<p style="font-size:15px;padding-right:30px;padding-left:0px;padding-top:8px;">overview: </p>
{{range $index,$value := .Overview}}
<p style="font-size:15px;padding-right:30px;padding-left:0px;padding-top:8px;">{{$value}}</p>
{{end}}
</tr>
<tr style="width:100%">
<td style="width:100%">
<span style="display:inline-block;position:relative;border-radius:4px;background-color:#463ff4" class="sm_full_width">
<a class="button_link sm_full_width sm_border_box" href="{{.Link}}" style="padding:16px 40px;border-radius:4px;background-color:#463ff4;color:#fff;font-size:20px;line-height:24px;word-break:break-word;display:inline-block;text-align:center;font-weight:700;text-decoration:none!important">
Summary
</a>
</span>
</td>
</tr>
</table>
</div>
<p style="font-size:17px;padding-right:30px;padding-left:30px;margin-top:40px;margin-bottom:30px">
<strong>Didn't request this email?</strong>
<br/>No worries! Your address may have been entered by mistake. You can safely ignore or delete this message.</p>
<!--[if mso]>
              </td>
              <td style="padding: 0; margin: 0;">&nbsp;</td>
              </tr>
              <tr>
              <td colspan="3" style="padding: 0; margin: 0; font-size: 20px; height: 20px;" height="20">&nbsp;</td>
              </tr>
              </table>
              <![endif]-->
</td>
</tr>
</table>
</td>
</tr>
</table>
</div>
</td>
</tr>
<tr>
<td class="email_footer" style="padding:0 30px 40px;border-top:1px solid #e1e1e4;line-height:24px;font-size:15px;color:#717274;text-align:center;width:100%">
<table border="0" cellpadding="0" cellspacing="0" width="100%" style="margin-top:20px;background-color:#fff">
<tr>
<td>&nbsp;</td>
<td>
<span style="display:block"><a href="https://vibe.us" style="text-decoration:none;color:#434245">Vibe Inc.</a></span>
</td>
<td>&nbsp;</td>
</tr>
</table>
</td>
</tr>
</table>
</body>
</html>

--VN9IOIHhMoPcH0icu--
`))