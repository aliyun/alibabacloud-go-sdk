// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConsoleMessagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessageList(v []*DescribeConsoleMessagesResponseBodyMessageList) *DescribeConsoleMessagesResponseBody
	GetMessageList() []*DescribeConsoleMessagesResponseBodyMessageList
	SetRequestId(v string) *DescribeConsoleMessagesResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeConsoleMessagesResponseBody
	GetTotalCount() *string
}

type DescribeConsoleMessagesResponseBody struct {
	// The list of console messages queried for the current page. If no messages match, an empty list is returned.
	MessageList []*DescribeConsoleMessagesResponseBodyMessageList `json:"MessageList,omitempty" xml:"MessageList,omitempty" type:"Repeated"`
	// The ID of the current request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9F1D7E42-06A8-4B50-9A21-6E0C8D3F5B72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of messages that match the query conditions, not the number of messages returned on the current page.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeConsoleMessagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsoleMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConsoleMessagesResponseBody) GetMessageList() []*DescribeConsoleMessagesResponseBodyMessageList {
	return s.MessageList
}

func (s *DescribeConsoleMessagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConsoleMessagesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeConsoleMessagesResponseBody) SetMessageList(v []*DescribeConsoleMessagesResponseBodyMessageList) *DescribeConsoleMessagesResponseBody {
	s.MessageList = v
	return s
}

func (s *DescribeConsoleMessagesResponseBody) SetRequestId(v string) *DescribeConsoleMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBody) SetTotalCount(v string) *DescribeConsoleMessagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBody) Validate() error {
	if s.MessageList != nil {
		for _, item := range s.MessageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeConsoleMessagesResponseBodyMessageList struct {
	// The list of client types to which the message applies. Standard names are returned: Windows, Mac, Web, and Hardware.
	ClientTypes []*string `json:"ClientTypes,omitempty" xml:"ClientTypes,omitempty" type:"Repeated"`
	// The list of message body configurations for each language.
	ContentConfigs []*DescribeConsoleMessagesResponseBodyMessageListContentConfigs `json:"ContentConfigs,omitempty" xml:"ContentConfigs,omitempty" type:"Repeated"`
	// The time when the operational message stops taking effect. The time is in ISO 8601 format. A time zone offset of +00:00 indicates UTC time.
	//
	// example:
	//
	// 2026-10-11T02:00:00.000+00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the message was created.
	//
	// example:
	//
	// 2026-10-10T00:00:00.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the message was last modified. This time is also updated when the message status changes.
	//
	// example:
	//
	// 2026-10-10T01:00:00.000+00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The list of hyperlink button configurations in the message.
	HyperLinkConfigs []*DescribeConsoleMessagesResponseBodyMessageListHyperLinkConfigs `json:"HyperLinkConfigs,omitempty" xml:"HyperLinkConfigs,omitempty" type:"Repeated"`
	// The storage path of the image file for the operational message.
	//
	// example:
	//
	// image_file/OPERATIONAL/00000000-0000-console-message.png
	ImageFilePath *string `json:"ImageFilePath,omitempty" xml:"ImageFilePath,omitempty"`
	// The list of image configurations, which is mainly used for custom wallpaper messages. Each item includes the image name, image type, storage path, and temporary download URL.
	ImageInfoList []*DescribeConsoleMessagesResponseBodyMessageListImageInfoList `json:"ImageInfoList,omitempty" xml:"ImageInfoList,omitempty" type:"Repeated"`
	// The temporary download URL of the operational message image. The link has a limited validity period. If the link expires, you can query again to obtain a new one.
	//
	// example:
	//
	// https://example.com/images/console-message.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The message ID, which can be used to send, modify, or delete a specified message.
	//
	// example:
	//
	// msg-123456
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The message level of the notification message. Valid values: STRONG (strong message) and WEAK (weak message).
	//
	// example:
	//
	// STRONG
	MessageLevel *string `json:"MessageLevel,omitempty" xml:"MessageLevel,omitempty"`
	// The message type. Valid values: OPERATIONAL (client operational message), WALLPAPER (desktop wallpaper text message), NOTIFICATION (message center notification), and CUSTOM_WALLPAPER (custom desktop wallpaper).
	//
	// example:
	//
	// NOTIFICATION
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The additional configuration of the message, including the associated workspace, AD domain name, and scheduled sending configuration. Whether specific fields are returned depends on the message type and its saved configuration.
	MetaData *DescribeConsoleMessagesResponseBodyMessageListMetaData `json:"MetaData,omitempty" xml:"MetaData,omitempty" type:"Struct"`
	// The display position of the message. For wallpaper text messages, valid values are CENTER (center), TOP_RIGHT (top right), BOTTOM_RIGHT (bottom right), and BOTTOM_LEFT (bottom left). For client operational messages, valid values are CLIENT_ENTERPRISE_RESOURCE_PAGE (Enterprise Edition client cloud resource page) and CLIENT_BUSINESS_RESOURCE_PAGE (Business Edition client cloud resource page).
	//
	// example:
	//
	// CENTER
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
	// The time when the operational message starts taking effect. The time is in ISO 8601 format. A time zone offset of +00:00 indicates UTC time.
	//
	// example:
	//
	// 2026-10-10T02:00:00.000+00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The current status of the message. Valid values: Init (initializing), Pushing (pushing), Published (published), Expired (expired), Deleting (deleting), and Failed (failed). Published indicates that the message has been published or the server-side push processing has finished, but it does not indicate that all target terminals have successfully received the message.
	//
	// example:
	//
	// Published
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The JSON string of the list of targets to which the message applies. For Organization and AdOrg, it is an array of objects containing orgId (organization ID) and orgName (organization name). The organization name may be empty, and records in the old format may not contain orgName. For Tag, it is an array of objects containing tagKey (tag key), tagValue (tag value), and regionId (region ID). For other target types, it is an array of strings. The all-users target is represented by ALL_USERS.
	//
	// example:
	//
	// ["user01"]
	TargetIds *string `json:"TargetIds,omitempty" xml:"TargetIds,omitempty"`
	// The type of the target to which the message applies. Valid values: AllUsers (all convenience users), EndUser (specified convenience user), Organization (specified convenience user organization), Desktop (cloud computer), ResourceGroup (resource group), Region (region), Tag (tag), AdAllUsers (all AD users), AdUser (specified AD user), and AdOrg (specified AD organization).
	//
	// example:
	//
	// EndUser
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The list of message title configurations for each language.
	TitleConfigs []*DescribeConsoleMessagesResponseBodyMessageListTitleConfigs `json:"TitleConfigs,omitempty" xml:"TitleConfigs,omitempty" type:"Repeated"`
}

func (s DescribeConsoleMessagesResponseBodyMessageList) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsoleMessagesResponseBodyMessageList) GoString() string {
	return s.String()
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) GetClientTypes() []*string {
	return s.ClientTypes
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) GetContentConfigs() []*DescribeConsoleMessagesResponseBodyMessageListContentConfigs {
	return s.ContentConfigs
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) GetHyperLinkConfigs() []*DescribeConsoleMessagesResponseBodyMessageListHyperLinkConfigs {
	return s.HyperLinkConfigs
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) GetImageFilePath() *string {
	return s.ImageFilePath
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) GetImageInfoList() []*DescribeConsoleMessagesResponseBodyMessageListImageInfoList {
	return s.ImageInfoList
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) GetMessageId() *string {
	return s.MessageId
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) GetMessageLevel() *string {
	return s.MessageLevel
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) GetMessageType() *string {
	return s.MessageType
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) GetMetaData() *DescribeConsoleMessagesResponseBodyMessageListMetaData {
	return s.MetaData
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) GetPosition() *string {
	return s.Position
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) GetStatus() *string {
	return s.Status
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) GetTargetIds() *string {
	return s.TargetIds
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) GetTitleConfigs() []*DescribeConsoleMessagesResponseBodyMessageListTitleConfigs {
	return s.TitleConfigs
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) SetClientTypes(v []*string) *DescribeConsoleMessagesResponseBodyMessageList {
	s.ClientTypes = v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) SetContentConfigs(v []*DescribeConsoleMessagesResponseBodyMessageListContentConfigs) *DescribeConsoleMessagesResponseBodyMessageList {
	s.ContentConfigs = v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) SetEndTime(v string) *DescribeConsoleMessagesResponseBodyMessageList {
	s.EndTime = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) SetGmtCreate(v string) *DescribeConsoleMessagesResponseBodyMessageList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) SetGmtModified(v string) *DescribeConsoleMessagesResponseBodyMessageList {
	s.GmtModified = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) SetHyperLinkConfigs(v []*DescribeConsoleMessagesResponseBodyMessageListHyperLinkConfigs) *DescribeConsoleMessagesResponseBodyMessageList {
	s.HyperLinkConfigs = v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) SetImageFilePath(v string) *DescribeConsoleMessagesResponseBodyMessageList {
	s.ImageFilePath = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) SetImageInfoList(v []*DescribeConsoleMessagesResponseBodyMessageListImageInfoList) *DescribeConsoleMessagesResponseBodyMessageList {
	s.ImageInfoList = v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) SetImageUrl(v string) *DescribeConsoleMessagesResponseBodyMessageList {
	s.ImageUrl = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) SetMessageId(v string) *DescribeConsoleMessagesResponseBodyMessageList {
	s.MessageId = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) SetMessageLevel(v string) *DescribeConsoleMessagesResponseBodyMessageList {
	s.MessageLevel = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) SetMessageType(v string) *DescribeConsoleMessagesResponseBodyMessageList {
	s.MessageType = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) SetMetaData(v *DescribeConsoleMessagesResponseBodyMessageListMetaData) *DescribeConsoleMessagesResponseBodyMessageList {
	s.MetaData = v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) SetPosition(v string) *DescribeConsoleMessagesResponseBodyMessageList {
	s.Position = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) SetStartTime(v string) *DescribeConsoleMessagesResponseBodyMessageList {
	s.StartTime = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) SetStatus(v string) *DescribeConsoleMessagesResponseBodyMessageList {
	s.Status = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) SetTargetIds(v string) *DescribeConsoleMessagesResponseBodyMessageList {
	s.TargetIds = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) SetTargetType(v string) *DescribeConsoleMessagesResponseBodyMessageList {
	s.TargetType = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) SetTitleConfigs(v []*DescribeConsoleMessagesResponseBodyMessageListTitleConfigs) *DescribeConsoleMessagesResponseBodyMessageList {
	s.TitleConfigs = v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageList) Validate() error {
	if s.ContentConfigs != nil {
		for _, item := range s.ContentConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HyperLinkConfigs != nil {
		for _, item := range s.HyperLinkConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ImageInfoList != nil {
		for _, item := range s.ImageInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MetaData != nil {
		if err := s.MetaData.Validate(); err != nil {
			return err
		}
	}
	if s.TitleConfigs != nil {
		for _, item := range s.TitleConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeConsoleMessagesResponseBodyMessageListContentConfigs struct {
	// The color of the body text. The color value saved when the message is created or modified is used.
	//
	// example:
	//
	// #666666
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
	// The message body text in the specified language.
	//
	// example:
	//
	// The cloud computer will undergo routine maintenance. Please save your work in advance
	ContentText *string `json:"ContentText,omitempty" xml:"ContentText,omitempty"`
	// The body language. Valid values: CN (Chinese), EN (English), and JP (Japanese). The query returns the language value saved when the message is created or modified, retaining the original case.
	//
	// example:
	//
	// CN
	LangType *string `json:"LangType,omitempty" xml:"LangType,omitempty"`
}

func (s DescribeConsoleMessagesResponseBodyMessageListContentConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsoleMessagesResponseBodyMessageListContentConfigs) GoString() string {
	return s.String()
}

func (s *DescribeConsoleMessagesResponseBodyMessageListContentConfigs) GetColor() *string {
	return s.Color
}

func (s *DescribeConsoleMessagesResponseBodyMessageListContentConfigs) GetContentText() *string {
	return s.ContentText
}

func (s *DescribeConsoleMessagesResponseBodyMessageListContentConfigs) GetLangType() *string {
	return s.LangType
}

func (s *DescribeConsoleMessagesResponseBodyMessageListContentConfigs) SetColor(v string) *DescribeConsoleMessagesResponseBodyMessageListContentConfigs {
	s.Color = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageListContentConfigs) SetContentText(v string) *DescribeConsoleMessagesResponseBodyMessageListContentConfigs {
	s.ContentText = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageListContentConfigs) SetLangType(v string) *DescribeConsoleMessagesResponseBodyMessageListContentConfigs {
	s.LangType = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageListContentConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeConsoleMessagesResponseBodyMessageListHyperLinkConfigs struct {
	// The text displayed on the hyperlink button.
	//
	// example:
	//
	// View Details
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The URL accessed after the hyperlink button is clicked.
	//
	// example:
	//
	// https://example.com/announcements/maintenance
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeConsoleMessagesResponseBodyMessageListHyperLinkConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsoleMessagesResponseBodyMessageListHyperLinkConfigs) GoString() string {
	return s.String()
}

func (s *DescribeConsoleMessagesResponseBodyMessageListHyperLinkConfigs) GetTitle() *string {
	return s.Title
}

func (s *DescribeConsoleMessagesResponseBodyMessageListHyperLinkConfigs) GetUrl() *string {
	return s.Url
}

func (s *DescribeConsoleMessagesResponseBodyMessageListHyperLinkConfigs) SetTitle(v string) *DescribeConsoleMessagesResponseBodyMessageListHyperLinkConfigs {
	s.Title = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageListHyperLinkConfigs) SetUrl(v string) *DescribeConsoleMessagesResponseBodyMessageListHyperLinkConfigs {
	s.Url = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageListHyperLinkConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeConsoleMessagesResponseBodyMessageListImageInfoList struct {
	// The name of the image file.
	//
	// example:
	//
	// img0.jpg
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The storage path of the image file.
	//
	// example:
	//
	// image_file/CUSTOM_WALLPAPER/00000000-0000-img0.jpg
	ImageFilePath *string `json:"ImageFilePath,omitempty" xml:"ImageFilePath,omitempty"`
	// The image type. Configuration types defined by the backend include Operational (operational message image), WallpaperBasic (basic wallpaper image), and WallpaperAppend (appended wallpaper image). The query returns the image type value saved when the message is created or modified, retaining the original case.
	//
	// example:
	//
	// WallpaperBasic
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The temporary download URL of the image file. The link has a limited validity period. If the link expires, you can query again to obtain a new one.
	//
	// example:
	//
	// https://example.com/images/img0.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s DescribeConsoleMessagesResponseBodyMessageListImageInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsoleMessagesResponseBodyMessageListImageInfoList) GoString() string {
	return s.String()
}

func (s *DescribeConsoleMessagesResponseBodyMessageListImageInfoList) GetFileName() *string {
	return s.FileName
}

func (s *DescribeConsoleMessagesResponseBodyMessageListImageInfoList) GetImageFilePath() *string {
	return s.ImageFilePath
}

func (s *DescribeConsoleMessagesResponseBodyMessageListImageInfoList) GetImageType() *string {
	return s.ImageType
}

func (s *DescribeConsoleMessagesResponseBodyMessageListImageInfoList) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DescribeConsoleMessagesResponseBodyMessageListImageInfoList) SetFileName(v string) *DescribeConsoleMessagesResponseBodyMessageListImageInfoList {
	s.FileName = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageListImageInfoList) SetImageFilePath(v string) *DescribeConsoleMessagesResponseBodyMessageListImageInfoList {
	s.ImageFilePath = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageListImageInfoList) SetImageType(v string) *DescribeConsoleMessagesResponseBodyMessageListImageInfoList {
	s.ImageType = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageListImageInfoList) SetImageUrl(v string) *DescribeConsoleMessagesResponseBodyMessageListImageInfoList {
	s.ImageUrl = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageListImageInfoList) Validate() error {
	return dara.Validate(s)
}

type DescribeConsoleMessagesResponseBodyMessageListMetaData struct {
	// The AD domain name associated with the message configuration.
	//
	// example:
	//
	// ad.example.com
	AdDomain *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	// The workspace ID associated with the message configuration.
	//
	// example:
	//
	// cn-shanghai+dir-example01
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The list of scheduled sending configurations. If TimerType is CRON, the elements are cron expressions. If TimerType is Timestamp, the elements are millisecond timestamps represented as strings.
	Timer []*string `json:"Timer,omitempty" xml:"Timer,omitempty" type:"Repeated"`
	// The scheduled sending type. Valid values: CRON (cron expression) and Timestamp (timestamp). The configuration values are case-sensitive.
	//
	// example:
	//
	// CRON
	TimerType *string `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
}

func (s DescribeConsoleMessagesResponseBodyMessageListMetaData) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsoleMessagesResponseBodyMessageListMetaData) GoString() string {
	return s.String()
}

func (s *DescribeConsoleMessagesResponseBodyMessageListMetaData) GetAdDomain() *string {
	return s.AdDomain
}

func (s *DescribeConsoleMessagesResponseBodyMessageListMetaData) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeConsoleMessagesResponseBodyMessageListMetaData) GetTimer() []*string {
	return s.Timer
}

func (s *DescribeConsoleMessagesResponseBodyMessageListMetaData) GetTimerType() *string {
	return s.TimerType
}

func (s *DescribeConsoleMessagesResponseBodyMessageListMetaData) SetAdDomain(v string) *DescribeConsoleMessagesResponseBodyMessageListMetaData {
	s.AdDomain = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageListMetaData) SetOfficeSiteId(v string) *DescribeConsoleMessagesResponseBodyMessageListMetaData {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageListMetaData) SetTimer(v []*string) *DescribeConsoleMessagesResponseBodyMessageListMetaData {
	s.Timer = v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageListMetaData) SetTimerType(v string) *DescribeConsoleMessagesResponseBodyMessageListMetaData {
	s.TimerType = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageListMetaData) Validate() error {
	return dara.Validate(s)
}

type DescribeConsoleMessagesResponseBodyMessageListTitleConfigs struct {
	// The color of the title text. The color value saved when the message is created or modified is used.
	//
	// example:
	//
	// #333333
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
	// The title language. Valid values: CN (Chinese), EN (English), and JP (Japanese). The query returns the language value saved when the message is created or modified, retaining the original case.
	//
	// example:
	//
	// CN
	LangType *string `json:"LangType,omitempty" xml:"LangType,omitempty"`
	// The message title text in the specified language.
	//
	// example:
	//
	// System Maintenance Notice
	TitleText *string `json:"TitleText,omitempty" xml:"TitleText,omitempty"`
}

func (s DescribeConsoleMessagesResponseBodyMessageListTitleConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsoleMessagesResponseBodyMessageListTitleConfigs) GoString() string {
	return s.String()
}

func (s *DescribeConsoleMessagesResponseBodyMessageListTitleConfigs) GetColor() *string {
	return s.Color
}

func (s *DescribeConsoleMessagesResponseBodyMessageListTitleConfigs) GetLangType() *string {
	return s.LangType
}

func (s *DescribeConsoleMessagesResponseBodyMessageListTitleConfigs) GetTitleText() *string {
	return s.TitleText
}

func (s *DescribeConsoleMessagesResponseBodyMessageListTitleConfigs) SetColor(v string) *DescribeConsoleMessagesResponseBodyMessageListTitleConfigs {
	s.Color = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageListTitleConfigs) SetLangType(v string) *DescribeConsoleMessagesResponseBodyMessageListTitleConfigs {
	s.LangType = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageListTitleConfigs) SetTitleText(v string) *DescribeConsoleMessagesResponseBodyMessageListTitleConfigs {
	s.TitleText = &v
	return s
}

func (s *DescribeConsoleMessagesResponseBodyMessageListTitleConfigs) Validate() error {
	return dara.Validate(s)
}
