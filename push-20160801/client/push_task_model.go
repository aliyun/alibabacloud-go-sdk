// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushTask interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *PushTask
	GetAction() *string
	SetMessage(v *PushTaskMessage) *PushTask
	GetMessage() *PushTaskMessage
	SetNotification(v *PushTaskNotification) *PushTask
	GetNotification() *PushTaskNotification
	SetOptions(v *PushTaskOptions) *PushTask
	GetOptions() *PushTaskOptions
	SetTarget(v *PushTaskTarget) *PushTask
	GetTarget() *PushTaskTarget
}

type PushTask struct {
	Action       *string               `json:"Action,omitempty" xml:"Action,omitempty"`
	Message      *PushTaskMessage      `json:"Message,omitempty" xml:"Message,omitempty" type:"Struct"`
	Notification *PushTaskNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	Options      *PushTaskOptions      `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	Target       *PushTaskTarget       `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
}

func (s PushTask) String() string {
	return dara.Prettify(s)
}

func (s PushTask) GoString() string {
	return s.String()
}

func (s *PushTask) GetAction() *string {
	return s.Action
}

func (s *PushTask) GetMessage() *PushTaskMessage {
	return s.Message
}

func (s *PushTask) GetNotification() *PushTaskNotification {
	return s.Notification
}

func (s *PushTask) GetOptions() *PushTaskOptions {
	return s.Options
}

func (s *PushTask) GetTarget() *PushTaskTarget {
	return s.Target
}

func (s *PushTask) SetAction(v string) *PushTask {
	s.Action = &v
	return s
}

func (s *PushTask) SetMessage(v *PushTaskMessage) *PushTask {
	s.Message = v
	return s
}

func (s *PushTask) SetNotification(v *PushTaskNotification) *PushTask {
	s.Notification = v
	return s
}

func (s *PushTask) SetOptions(v *PushTaskOptions) *PushTask {
	s.Options = v
	return s
}

func (s *PushTask) SetTarget(v *PushTaskTarget) *PushTask {
	s.Target = v
	return s
}

func (s *PushTask) Validate() error {
	if s.Message != nil {
		if err := s.Message.Validate(); err != nil {
			return err
		}
	}
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			return err
		}
	}
	if s.Options != nil {
		if err := s.Options.Validate(); err != nil {
			return err
		}
	}
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PushTaskMessage struct {
	Body  *string `json:"Body,omitempty" xml:"Body,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PushTaskMessage) String() string {
	return dara.Prettify(s)
}

func (s PushTaskMessage) GoString() string {
	return s.String()
}

func (s *PushTaskMessage) GetBody() *string {
	return s.Body
}

func (s *PushTaskMessage) GetTitle() *string {
	return s.Title
}

func (s *PushTaskMessage) SetBody(v string) *PushTaskMessage {
	s.Body = &v
	return s
}

func (s *PushTaskMessage) SetTitle(v string) *PushTaskMessage {
	s.Title = &v
	return s
}

func (s *PushTaskMessage) Validate() error {
	return dara.Validate(s)
}

type PushTaskNotification struct {
	Android *PushTaskNotificationAndroid `json:"Android,omitempty" xml:"Android,omitempty" type:"Struct"`
	Body    *string                      `json:"Body,omitempty" xml:"Body,omitempty"`
	Hmos    *PushTaskNotificationHmos    `json:"Hmos,omitempty" xml:"Hmos,omitempty" type:"Struct"`
	Ios     *PushTaskNotificationIos     `json:"Ios,omitempty" xml:"Ios,omitempty" type:"Struct"`
	Title   *string                      `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PushTaskNotification) String() string {
	return dara.Prettify(s)
}

func (s PushTaskNotification) GoString() string {
	return s.String()
}

func (s *PushTaskNotification) GetAndroid() *PushTaskNotificationAndroid {
	return s.Android
}

func (s *PushTaskNotification) GetBody() *string {
	return s.Body
}

func (s *PushTaskNotification) GetHmos() *PushTaskNotificationHmos {
	return s.Hmos
}

func (s *PushTaskNotification) GetIos() *PushTaskNotificationIos {
	return s.Ios
}

func (s *PushTaskNotification) GetTitle() *string {
	return s.Title
}

func (s *PushTaskNotification) SetAndroid(v *PushTaskNotificationAndroid) *PushTaskNotification {
	s.Android = v
	return s
}

func (s *PushTaskNotification) SetBody(v string) *PushTaskNotification {
	s.Body = &v
	return s
}

func (s *PushTaskNotification) SetHmos(v *PushTaskNotificationHmos) *PushTaskNotification {
	s.Hmos = v
	return s
}

func (s *PushTaskNotification) SetIos(v *PushTaskNotificationIos) *PushTaskNotification {
	s.Ios = v
	return s
}

func (s *PushTaskNotification) SetTitle(v string) *PushTaskNotification {
	s.Title = &v
	return s
}

func (s *PushTaskNotification) Validate() error {
	if s.Android != nil {
		if err := s.Android.Validate(); err != nil {
			return err
		}
	}
	if s.Hmos != nil {
		if err := s.Hmos.Validate(); err != nil {
			return err
		}
	}
	if s.Ios != nil {
		if err := s.Ios.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PushTaskNotificationAndroid struct {
	BadgeActivity         *string                             `json:"BadgeActivity,omitempty" xml:"BadgeActivity,omitempty"`
	BadgeAddNum           *int32                              `json:"BadgeAddNum,omitempty" xml:"BadgeAddNum,omitempty"`
	BadgeSetNum           *int32                              `json:"BadgeSetNum,omitempty" xml:"BadgeSetNum,omitempty"`
	ChannelId             *string                             `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ExtParameters         *string                             `json:"ExtParameters,omitempty" xml:"ExtParameters,omitempty"`
	GroupId               *string                             `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ImageUrl              *string                             `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	InboxContent          []*string                           `json:"InboxContent,omitempty" xml:"InboxContent,omitempty" type:"Repeated"`
	Music                 *string                             `json:"Music,omitempty" xml:"Music,omitempty"`
	NotifyId              *int32                              `json:"NotifyId,omitempty" xml:"NotifyId,omitempty"`
	Options               *PushTaskNotificationAndroidOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	PictureUrl            *string                             `json:"PictureUrl,omitempty" xml:"PictureUrl,omitempty"`
	RenderStyle           *string                             `json:"RenderStyle,omitempty" xml:"RenderStyle,omitempty"`
	TestMessage           *bool                               `json:"TestMessage,omitempty" xml:"TestMessage,omitempty"`
	VendorChannelActivity *string                             `json:"VendorChannelActivity,omitempty" xml:"VendorChannelActivity,omitempty"`
}

func (s PushTaskNotificationAndroid) String() string {
	return dara.Prettify(s)
}

func (s PushTaskNotificationAndroid) GoString() string {
	return s.String()
}

func (s *PushTaskNotificationAndroid) GetBadgeActivity() *string {
	return s.BadgeActivity
}

func (s *PushTaskNotificationAndroid) GetBadgeAddNum() *int32 {
	return s.BadgeAddNum
}

func (s *PushTaskNotificationAndroid) GetBadgeSetNum() *int32 {
	return s.BadgeSetNum
}

func (s *PushTaskNotificationAndroid) GetChannelId() *string {
	return s.ChannelId
}

func (s *PushTaskNotificationAndroid) GetExtParameters() *string {
	return s.ExtParameters
}

func (s *PushTaskNotificationAndroid) GetGroupId() *string {
	return s.GroupId
}

func (s *PushTaskNotificationAndroid) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *PushTaskNotificationAndroid) GetInboxContent() []*string {
	return s.InboxContent
}

func (s *PushTaskNotificationAndroid) GetMusic() *string {
	return s.Music
}

func (s *PushTaskNotificationAndroid) GetNotifyId() *int32 {
	return s.NotifyId
}

func (s *PushTaskNotificationAndroid) GetOptions() *PushTaskNotificationAndroidOptions {
	return s.Options
}

func (s *PushTaskNotificationAndroid) GetPictureUrl() *string {
	return s.PictureUrl
}

func (s *PushTaskNotificationAndroid) GetRenderStyle() *string {
	return s.RenderStyle
}

func (s *PushTaskNotificationAndroid) GetTestMessage() *bool {
	return s.TestMessage
}

func (s *PushTaskNotificationAndroid) GetVendorChannelActivity() *string {
	return s.VendorChannelActivity
}

func (s *PushTaskNotificationAndroid) SetBadgeActivity(v string) *PushTaskNotificationAndroid {
	s.BadgeActivity = &v
	return s
}

func (s *PushTaskNotificationAndroid) SetBadgeAddNum(v int32) *PushTaskNotificationAndroid {
	s.BadgeAddNum = &v
	return s
}

func (s *PushTaskNotificationAndroid) SetBadgeSetNum(v int32) *PushTaskNotificationAndroid {
	s.BadgeSetNum = &v
	return s
}

func (s *PushTaskNotificationAndroid) SetChannelId(v string) *PushTaskNotificationAndroid {
	s.ChannelId = &v
	return s
}

func (s *PushTaskNotificationAndroid) SetExtParameters(v string) *PushTaskNotificationAndroid {
	s.ExtParameters = &v
	return s
}

func (s *PushTaskNotificationAndroid) SetGroupId(v string) *PushTaskNotificationAndroid {
	s.GroupId = &v
	return s
}

func (s *PushTaskNotificationAndroid) SetImageUrl(v string) *PushTaskNotificationAndroid {
	s.ImageUrl = &v
	return s
}

func (s *PushTaskNotificationAndroid) SetInboxContent(v []*string) *PushTaskNotificationAndroid {
	s.InboxContent = v
	return s
}

func (s *PushTaskNotificationAndroid) SetMusic(v string) *PushTaskNotificationAndroid {
	s.Music = &v
	return s
}

func (s *PushTaskNotificationAndroid) SetNotifyId(v int32) *PushTaskNotificationAndroid {
	s.NotifyId = &v
	return s
}

func (s *PushTaskNotificationAndroid) SetOptions(v *PushTaskNotificationAndroidOptions) *PushTaskNotificationAndroid {
	s.Options = v
	return s
}

func (s *PushTaskNotificationAndroid) SetPictureUrl(v string) *PushTaskNotificationAndroid {
	s.PictureUrl = &v
	return s
}

func (s *PushTaskNotificationAndroid) SetRenderStyle(v string) *PushTaskNotificationAndroid {
	s.RenderStyle = &v
	return s
}

func (s *PushTaskNotificationAndroid) SetTestMessage(v bool) *PushTaskNotificationAndroid {
	s.TestMessage = &v
	return s
}

func (s *PushTaskNotificationAndroid) SetVendorChannelActivity(v string) *PushTaskNotificationAndroid {
	s.VendorChannelActivity = &v
	return s
}

func (s *PushTaskNotificationAndroid) Validate() error {
	if s.Options != nil {
		if err := s.Options.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PushTaskNotificationAndroidOptions struct {
	Accs   *PushTaskNotificationAndroidOptionsAccs   `json:"Accs,omitempty" xml:"Accs,omitempty" type:"Struct"`
	Honor  *PushTaskNotificationAndroidOptionsHonor  `json:"Honor,omitempty" xml:"Honor,omitempty" type:"Struct"`
	Huawei *PushTaskNotificationAndroidOptionsHuawei `json:"Huawei,omitempty" xml:"Huawei,omitempty" type:"Struct"`
	Meizu  *PushTaskNotificationAndroidOptionsMeizu  `json:"Meizu,omitempty" xml:"Meizu,omitempty" type:"Struct"`
	Oppo   *PushTaskNotificationAndroidOptionsOppo   `json:"Oppo,omitempty" xml:"Oppo,omitempty" type:"Struct"`
	Vivo   *PushTaskNotificationAndroidOptionsVivo   `json:"Vivo,omitempty" xml:"Vivo,omitempty" type:"Struct"`
	Xiaomi *PushTaskNotificationAndroidOptionsXiaomi `json:"Xiaomi,omitempty" xml:"Xiaomi,omitempty" type:"Struct"`
}

func (s PushTaskNotificationAndroidOptions) String() string {
	return dara.Prettify(s)
}

func (s PushTaskNotificationAndroidOptions) GoString() string {
	return s.String()
}

func (s *PushTaskNotificationAndroidOptions) GetAccs() *PushTaskNotificationAndroidOptionsAccs {
	return s.Accs
}

func (s *PushTaskNotificationAndroidOptions) GetHonor() *PushTaskNotificationAndroidOptionsHonor {
	return s.Honor
}

func (s *PushTaskNotificationAndroidOptions) GetHuawei() *PushTaskNotificationAndroidOptionsHuawei {
	return s.Huawei
}

func (s *PushTaskNotificationAndroidOptions) GetMeizu() *PushTaskNotificationAndroidOptionsMeizu {
	return s.Meizu
}

func (s *PushTaskNotificationAndroidOptions) GetOppo() *PushTaskNotificationAndroidOptionsOppo {
	return s.Oppo
}

func (s *PushTaskNotificationAndroidOptions) GetVivo() *PushTaskNotificationAndroidOptionsVivo {
	return s.Vivo
}

func (s *PushTaskNotificationAndroidOptions) GetXiaomi() *PushTaskNotificationAndroidOptionsXiaomi {
	return s.Xiaomi
}

func (s *PushTaskNotificationAndroidOptions) SetAccs(v *PushTaskNotificationAndroidOptionsAccs) *PushTaskNotificationAndroidOptions {
	s.Accs = v
	return s
}

func (s *PushTaskNotificationAndroidOptions) SetHonor(v *PushTaskNotificationAndroidOptionsHonor) *PushTaskNotificationAndroidOptions {
	s.Honor = v
	return s
}

func (s *PushTaskNotificationAndroidOptions) SetHuawei(v *PushTaskNotificationAndroidOptionsHuawei) *PushTaskNotificationAndroidOptions {
	s.Huawei = v
	return s
}

func (s *PushTaskNotificationAndroidOptions) SetMeizu(v *PushTaskNotificationAndroidOptionsMeizu) *PushTaskNotificationAndroidOptions {
	s.Meizu = v
	return s
}

func (s *PushTaskNotificationAndroidOptions) SetOppo(v *PushTaskNotificationAndroidOptionsOppo) *PushTaskNotificationAndroidOptions {
	s.Oppo = v
	return s
}

func (s *PushTaskNotificationAndroidOptions) SetVivo(v *PushTaskNotificationAndroidOptionsVivo) *PushTaskNotificationAndroidOptions {
	s.Vivo = v
	return s
}

func (s *PushTaskNotificationAndroidOptions) SetXiaomi(v *PushTaskNotificationAndroidOptionsXiaomi) *PushTaskNotificationAndroidOptions {
	s.Xiaomi = v
	return s
}

func (s *PushTaskNotificationAndroidOptions) Validate() error {
	if s.Accs != nil {
		if err := s.Accs.Validate(); err != nil {
			return err
		}
	}
	if s.Honor != nil {
		if err := s.Honor.Validate(); err != nil {
			return err
		}
	}
	if s.Huawei != nil {
		if err := s.Huawei.Validate(); err != nil {
			return err
		}
	}
	if s.Meizu != nil {
		if err := s.Meizu.Validate(); err != nil {
			return err
		}
	}
	if s.Oppo != nil {
		if err := s.Oppo.Validate(); err != nil {
			return err
		}
	}
	if s.Vivo != nil {
		if err := s.Vivo.Validate(); err != nil {
			return err
		}
	}
	if s.Xiaomi != nil {
		if err := s.Xiaomi.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PushTaskNotificationAndroidOptionsAccs struct {
	CustomStyle  *int32  `json:"CustomStyle,omitempty" xml:"CustomStyle,omitempty"`
	NotifyType   *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	OpenActivity *string `json:"OpenActivity,omitempty" xml:"OpenActivity,omitempty"`
	OpenType     *string `json:"OpenType,omitempty" xml:"OpenType,omitempty"`
	OpenUrl      *string `json:"OpenUrl,omitempty" xml:"OpenUrl,omitempty"`
	Priority     *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ThreadId     *string `json:"ThreadId,omitempty" xml:"ThreadId,omitempty"`
}

func (s PushTaskNotificationAndroidOptionsAccs) String() string {
	return dara.Prettify(s)
}

func (s PushTaskNotificationAndroidOptionsAccs) GoString() string {
	return s.String()
}

func (s *PushTaskNotificationAndroidOptionsAccs) GetCustomStyle() *int32 {
	return s.CustomStyle
}

func (s *PushTaskNotificationAndroidOptionsAccs) GetNotifyType() *string {
	return s.NotifyType
}

func (s *PushTaskNotificationAndroidOptionsAccs) GetOpenActivity() *string {
	return s.OpenActivity
}

func (s *PushTaskNotificationAndroidOptionsAccs) GetOpenType() *string {
	return s.OpenType
}

func (s *PushTaskNotificationAndroidOptionsAccs) GetOpenUrl() *string {
	return s.OpenUrl
}

func (s *PushTaskNotificationAndroidOptionsAccs) GetPriority() *int32 {
	return s.Priority
}

func (s *PushTaskNotificationAndroidOptionsAccs) GetThreadId() *string {
	return s.ThreadId
}

func (s *PushTaskNotificationAndroidOptionsAccs) SetCustomStyle(v int32) *PushTaskNotificationAndroidOptionsAccs {
	s.CustomStyle = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsAccs) SetNotifyType(v string) *PushTaskNotificationAndroidOptionsAccs {
	s.NotifyType = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsAccs) SetOpenActivity(v string) *PushTaskNotificationAndroidOptionsAccs {
	s.OpenActivity = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsAccs) SetOpenType(v string) *PushTaskNotificationAndroidOptionsAccs {
	s.OpenType = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsAccs) SetOpenUrl(v string) *PushTaskNotificationAndroidOptionsAccs {
	s.OpenUrl = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsAccs) SetPriority(v int32) *PushTaskNotificationAndroidOptionsAccs {
	s.Priority = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsAccs) SetThreadId(v string) *PushTaskNotificationAndroidOptionsAccs {
	s.ThreadId = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsAccs) Validate() error {
	return dara.Validate(s)
}

type PushTaskNotificationAndroidOptionsHonor struct {
	Importance *int32 `json:"Importance,omitempty" xml:"Importance,omitempty"`
}

func (s PushTaskNotificationAndroidOptionsHonor) String() string {
	return dara.Prettify(s)
}

func (s PushTaskNotificationAndroidOptionsHonor) GoString() string {
	return s.String()
}

func (s *PushTaskNotificationAndroidOptionsHonor) GetImportance() *int32 {
	return s.Importance
}

func (s *PushTaskNotificationAndroidOptionsHonor) SetImportance(v int32) *PushTaskNotificationAndroidOptionsHonor {
	s.Importance = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsHonor) Validate() error {
	return dara.Validate(s)
}

type PushTaskNotificationAndroidOptionsHuawei struct {
	Category                *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Importance              *int32  `json:"Importance,omitempty" xml:"Importance,omitempty"`
	LiveNotificationPayload *string `json:"LiveNotificationPayload,omitempty" xml:"LiveNotificationPayload,omitempty"`
	ReceiptId               *string `json:"ReceiptId,omitempty" xml:"ReceiptId,omitempty"`
	Urgency                 *string `json:"Urgency,omitempty" xml:"Urgency,omitempty"`
}

func (s PushTaskNotificationAndroidOptionsHuawei) String() string {
	return dara.Prettify(s)
}

func (s PushTaskNotificationAndroidOptionsHuawei) GoString() string {
	return s.String()
}

func (s *PushTaskNotificationAndroidOptionsHuawei) GetCategory() *string {
	return s.Category
}

func (s *PushTaskNotificationAndroidOptionsHuawei) GetImportance() *int32 {
	return s.Importance
}

func (s *PushTaskNotificationAndroidOptionsHuawei) GetLiveNotificationPayload() *string {
	return s.LiveNotificationPayload
}

func (s *PushTaskNotificationAndroidOptionsHuawei) GetReceiptId() *string {
	return s.ReceiptId
}

func (s *PushTaskNotificationAndroidOptionsHuawei) GetUrgency() *string {
	return s.Urgency
}

func (s *PushTaskNotificationAndroidOptionsHuawei) SetCategory(v string) *PushTaskNotificationAndroidOptionsHuawei {
	s.Category = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsHuawei) SetImportance(v int32) *PushTaskNotificationAndroidOptionsHuawei {
	s.Importance = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsHuawei) SetLiveNotificationPayload(v string) *PushTaskNotificationAndroidOptionsHuawei {
	s.LiveNotificationPayload = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsHuawei) SetReceiptId(v string) *PushTaskNotificationAndroidOptionsHuawei {
	s.ReceiptId = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsHuawei) SetUrgency(v string) *PushTaskNotificationAndroidOptionsHuawei {
	s.Urgency = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsHuawei) Validate() error {
	return dara.Validate(s)
}

type PushTaskNotificationAndroidOptionsMeizu struct {
	NoticeMsgType *int32 `json:"NoticeMsgType,omitempty" xml:"NoticeMsgType,omitempty"`
}

func (s PushTaskNotificationAndroidOptionsMeizu) String() string {
	return dara.Prettify(s)
}

func (s PushTaskNotificationAndroidOptionsMeizu) GoString() string {
	return s.String()
}

func (s *PushTaskNotificationAndroidOptionsMeizu) GetNoticeMsgType() *int32 {
	return s.NoticeMsgType
}

func (s *PushTaskNotificationAndroidOptionsMeizu) SetNoticeMsgType(v int32) *PushTaskNotificationAndroidOptionsMeizu {
	s.NoticeMsgType = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsMeizu) Validate() error {
	return dara.Validate(s)
}

type PushTaskNotificationAndroidOptionsOppo struct {
	Category                 *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteIntentData         *string `json:"DeleteIntentData,omitempty" xml:"DeleteIntentData,omitempty"`
	IntelligentIntent        *string `json:"IntelligentIntent,omitempty" xml:"IntelligentIntent,omitempty"`
	NotifyLevel              *int64  `json:"NotifyLevel,omitempty" xml:"NotifyLevel,omitempty"`
	PrivateContentParameters *string `json:"PrivateContentParameters,omitempty" xml:"PrivateContentParameters,omitempty"`
	PrivateMsgTemplateId     *string `json:"PrivateMsgTemplateId,omitempty" xml:"PrivateMsgTemplateId,omitempty"`
	PrivateTitleParameters   *string `json:"PrivateTitleParameters,omitempty" xml:"PrivateTitleParameters,omitempty"`
}

func (s PushTaskNotificationAndroidOptionsOppo) String() string {
	return dara.Prettify(s)
}

func (s PushTaskNotificationAndroidOptionsOppo) GoString() string {
	return s.String()
}

func (s *PushTaskNotificationAndroidOptionsOppo) GetCategory() *string {
	return s.Category
}

func (s *PushTaskNotificationAndroidOptionsOppo) GetDeleteIntentData() *string {
	return s.DeleteIntentData
}

func (s *PushTaskNotificationAndroidOptionsOppo) GetIntelligentIntent() *string {
	return s.IntelligentIntent
}

func (s *PushTaskNotificationAndroidOptionsOppo) GetNotifyLevel() *int64 {
	return s.NotifyLevel
}

func (s *PushTaskNotificationAndroidOptionsOppo) GetPrivateContentParameters() *string {
	return s.PrivateContentParameters
}

func (s *PushTaskNotificationAndroidOptionsOppo) GetPrivateMsgTemplateId() *string {
	return s.PrivateMsgTemplateId
}

func (s *PushTaskNotificationAndroidOptionsOppo) GetPrivateTitleParameters() *string {
	return s.PrivateTitleParameters
}

func (s *PushTaskNotificationAndroidOptionsOppo) SetCategory(v string) *PushTaskNotificationAndroidOptionsOppo {
	s.Category = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsOppo) SetDeleteIntentData(v string) *PushTaskNotificationAndroidOptionsOppo {
	s.DeleteIntentData = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsOppo) SetIntelligentIntent(v string) *PushTaskNotificationAndroidOptionsOppo {
	s.IntelligentIntent = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsOppo) SetNotifyLevel(v int64) *PushTaskNotificationAndroidOptionsOppo {
	s.NotifyLevel = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsOppo) SetPrivateContentParameters(v string) *PushTaskNotificationAndroidOptionsOppo {
	s.PrivateContentParameters = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsOppo) SetPrivateMsgTemplateId(v string) *PushTaskNotificationAndroidOptionsOppo {
	s.PrivateMsgTemplateId = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsOppo) SetPrivateTitleParameters(v string) *PushTaskNotificationAndroidOptionsOppo {
	s.PrivateTitleParameters = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsOppo) Validate() error {
	return dara.Validate(s)
}

type PushTaskNotificationAndroidOptionsVivo struct {
	Category   *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Importance *int32  `json:"Importance,omitempty" xml:"Importance,omitempty"`
	ReceiptId  *string `json:"ReceiptId,omitempty" xml:"ReceiptId,omitempty"`
}

func (s PushTaskNotificationAndroidOptionsVivo) String() string {
	return dara.Prettify(s)
}

func (s PushTaskNotificationAndroidOptionsVivo) GoString() string {
	return s.String()
}

func (s *PushTaskNotificationAndroidOptionsVivo) GetCategory() *string {
	return s.Category
}

func (s *PushTaskNotificationAndroidOptionsVivo) GetImportance() *int32 {
	return s.Importance
}

func (s *PushTaskNotificationAndroidOptionsVivo) GetReceiptId() *string {
	return s.ReceiptId
}

func (s *PushTaskNotificationAndroidOptionsVivo) SetCategory(v string) *PushTaskNotificationAndroidOptionsVivo {
	s.Category = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsVivo) SetImportance(v int32) *PushTaskNotificationAndroidOptionsVivo {
	s.Importance = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsVivo) SetReceiptId(v string) *PushTaskNotificationAndroidOptionsVivo {
	s.ReceiptId = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsVivo) Validate() error {
	return dara.Validate(s)
}

type PushTaskNotificationAndroidOptionsXiaomi struct {
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
}

func (s PushTaskNotificationAndroidOptionsXiaomi) String() string {
	return dara.Prettify(s)
}

func (s PushTaskNotificationAndroidOptionsXiaomi) GoString() string {
	return s.String()
}

func (s *PushTaskNotificationAndroidOptionsXiaomi) GetChannel() *string {
	return s.Channel
}

func (s *PushTaskNotificationAndroidOptionsXiaomi) SetChannel(v string) *PushTaskNotificationAndroidOptionsXiaomi {
	s.Channel = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsXiaomi) Validate() error {
	return dara.Validate(s)
}

type PushTaskNotificationHmos struct {
	Action             *string   `json:"Action,omitempty" xml:"Action,omitempty"`
	BadgeAddNum        *int32    `json:"BadgeAddNum,omitempty" xml:"BadgeAddNum,omitempty"`
	BadgeSetNum        *int32    `json:"BadgeSetNum,omitempty" xml:"BadgeSetNum,omitempty"`
	Category           *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	ExtParameters      *string   `json:"ExtParameters,omitempty" xml:"ExtParameters,omitempty"`
	ExtensionExtraData *string   `json:"ExtensionExtraData,omitempty" xml:"ExtensionExtraData,omitempty"`
	ExtensionPush      *bool     `json:"ExtensionPush,omitempty" xml:"ExtensionPush,omitempty"`
	ImageUrl           *string   `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	InboxContent       []*string `json:"InboxContent,omitempty" xml:"InboxContent,omitempty" type:"Repeated"`
	LiveViewPayload    *string   `json:"LiveViewPayload,omitempty" xml:"LiveViewPayload,omitempty"`
	NotifyId           *int32    `json:"NotifyId,omitempty" xml:"NotifyId,omitempty"`
	ReceiptId          *string   `json:"ReceiptId,omitempty" xml:"ReceiptId,omitempty"`
	RenderStyle        *string   `json:"RenderStyle,omitempty" xml:"RenderStyle,omitempty"`
	SlotType           *string   `json:"SlotType,omitempty" xml:"SlotType,omitempty"`
	TestMessage        *bool     `json:"TestMessage,omitempty" xml:"TestMessage,omitempty"`
	Uri                *string   `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s PushTaskNotificationHmos) String() string {
	return dara.Prettify(s)
}

func (s PushTaskNotificationHmos) GoString() string {
	return s.String()
}

func (s *PushTaskNotificationHmos) GetAction() *string {
	return s.Action
}

func (s *PushTaskNotificationHmos) GetBadgeAddNum() *int32 {
	return s.BadgeAddNum
}

func (s *PushTaskNotificationHmos) GetBadgeSetNum() *int32 {
	return s.BadgeSetNum
}

func (s *PushTaskNotificationHmos) GetCategory() *string {
	return s.Category
}

func (s *PushTaskNotificationHmos) GetExtParameters() *string {
	return s.ExtParameters
}

func (s *PushTaskNotificationHmos) GetExtensionExtraData() *string {
	return s.ExtensionExtraData
}

func (s *PushTaskNotificationHmos) GetExtensionPush() *bool {
	return s.ExtensionPush
}

func (s *PushTaskNotificationHmos) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *PushTaskNotificationHmos) GetInboxContent() []*string {
	return s.InboxContent
}

func (s *PushTaskNotificationHmos) GetLiveViewPayload() *string {
	return s.LiveViewPayload
}

func (s *PushTaskNotificationHmos) GetNotifyId() *int32 {
	return s.NotifyId
}

func (s *PushTaskNotificationHmos) GetReceiptId() *string {
	return s.ReceiptId
}

func (s *PushTaskNotificationHmos) GetRenderStyle() *string {
	return s.RenderStyle
}

func (s *PushTaskNotificationHmos) GetSlotType() *string {
	return s.SlotType
}

func (s *PushTaskNotificationHmos) GetTestMessage() *bool {
	return s.TestMessage
}

func (s *PushTaskNotificationHmos) GetUri() *string {
	return s.Uri
}

func (s *PushTaskNotificationHmos) SetAction(v string) *PushTaskNotificationHmos {
	s.Action = &v
	return s
}

func (s *PushTaskNotificationHmos) SetBadgeAddNum(v int32) *PushTaskNotificationHmos {
	s.BadgeAddNum = &v
	return s
}

func (s *PushTaskNotificationHmos) SetBadgeSetNum(v int32) *PushTaskNotificationHmos {
	s.BadgeSetNum = &v
	return s
}

func (s *PushTaskNotificationHmos) SetCategory(v string) *PushTaskNotificationHmos {
	s.Category = &v
	return s
}

func (s *PushTaskNotificationHmos) SetExtParameters(v string) *PushTaskNotificationHmos {
	s.ExtParameters = &v
	return s
}

func (s *PushTaskNotificationHmos) SetExtensionExtraData(v string) *PushTaskNotificationHmos {
	s.ExtensionExtraData = &v
	return s
}

func (s *PushTaskNotificationHmos) SetExtensionPush(v bool) *PushTaskNotificationHmos {
	s.ExtensionPush = &v
	return s
}

func (s *PushTaskNotificationHmos) SetImageUrl(v string) *PushTaskNotificationHmos {
	s.ImageUrl = &v
	return s
}

func (s *PushTaskNotificationHmos) SetInboxContent(v []*string) *PushTaskNotificationHmos {
	s.InboxContent = v
	return s
}

func (s *PushTaskNotificationHmos) SetLiveViewPayload(v string) *PushTaskNotificationHmos {
	s.LiveViewPayload = &v
	return s
}

func (s *PushTaskNotificationHmos) SetNotifyId(v int32) *PushTaskNotificationHmos {
	s.NotifyId = &v
	return s
}

func (s *PushTaskNotificationHmos) SetReceiptId(v string) *PushTaskNotificationHmos {
	s.ReceiptId = &v
	return s
}

func (s *PushTaskNotificationHmos) SetRenderStyle(v string) *PushTaskNotificationHmos {
	s.RenderStyle = &v
	return s
}

func (s *PushTaskNotificationHmos) SetSlotType(v string) *PushTaskNotificationHmos {
	s.SlotType = &v
	return s
}

func (s *PushTaskNotificationHmos) SetTestMessage(v bool) *PushTaskNotificationHmos {
	s.TestMessage = &v
	return s
}

func (s *PushTaskNotificationHmos) SetUri(v string) *PushTaskNotificationHmos {
	s.Uri = &v
	return s
}

func (s *PushTaskNotificationHmos) Validate() error {
	return dara.Validate(s)
}

type PushTaskNotificationIos struct {
	ApnsEnv            *string                              `json:"ApnsEnv,omitempty" xml:"ApnsEnv,omitempty"`
	Badge              *int32                               `json:"Badge,omitempty" xml:"Badge,omitempty"`
	BadgeAutoIncrement *bool                                `json:"BadgeAutoIncrement,omitempty" xml:"BadgeAutoIncrement,omitempty"`
	Category           *string                              `json:"Category,omitempty" xml:"Category,omitempty"`
	CollapseId         *string                              `json:"CollapseId,omitempty" xml:"CollapseId,omitempty"`
	ExtParameters      *string                              `json:"ExtParameters,omitempty" xml:"ExtParameters,omitempty"`
	InterruptionLevel  *string                              `json:"InterruptionLevel,omitempty" xml:"InterruptionLevel,omitempty"`
	LiveActivity       *PushTaskNotificationIosLiveActivity `json:"LiveActivity,omitempty" xml:"LiveActivity,omitempty" type:"Struct"`
	Music              *string                              `json:"Music,omitempty" xml:"Music,omitempty"`
	Mutable            *bool                                `json:"Mutable,omitempty" xml:"Mutable,omitempty"`
	RelevanceScore     *float64                             `json:"RelevanceScore,omitempty" xml:"RelevanceScore,omitempty"`
	Silent             *bool                                `json:"Silent,omitempty" xml:"Silent,omitempty"`
	Subtitle           *string                              `json:"Subtitle,omitempty" xml:"Subtitle,omitempty"`
	ThreadId           *string                              `json:"ThreadId,omitempty" xml:"ThreadId,omitempty"`
}

func (s PushTaskNotificationIos) String() string {
	return dara.Prettify(s)
}

func (s PushTaskNotificationIos) GoString() string {
	return s.String()
}

func (s *PushTaskNotificationIos) GetApnsEnv() *string {
	return s.ApnsEnv
}

func (s *PushTaskNotificationIos) GetBadge() *int32 {
	return s.Badge
}

func (s *PushTaskNotificationIos) GetBadgeAutoIncrement() *bool {
	return s.BadgeAutoIncrement
}

func (s *PushTaskNotificationIos) GetCategory() *string {
	return s.Category
}

func (s *PushTaskNotificationIos) GetCollapseId() *string {
	return s.CollapseId
}

func (s *PushTaskNotificationIos) GetExtParameters() *string {
	return s.ExtParameters
}

func (s *PushTaskNotificationIos) GetInterruptionLevel() *string {
	return s.InterruptionLevel
}

func (s *PushTaskNotificationIos) GetLiveActivity() *PushTaskNotificationIosLiveActivity {
	return s.LiveActivity
}

func (s *PushTaskNotificationIos) GetMusic() *string {
	return s.Music
}

func (s *PushTaskNotificationIos) GetMutable() *bool {
	return s.Mutable
}

func (s *PushTaskNotificationIos) GetRelevanceScore() *float64 {
	return s.RelevanceScore
}

func (s *PushTaskNotificationIos) GetSilent() *bool {
	return s.Silent
}

func (s *PushTaskNotificationIos) GetSubtitle() *string {
	return s.Subtitle
}

func (s *PushTaskNotificationIos) GetThreadId() *string {
	return s.ThreadId
}

func (s *PushTaskNotificationIos) SetApnsEnv(v string) *PushTaskNotificationIos {
	s.ApnsEnv = &v
	return s
}

func (s *PushTaskNotificationIos) SetBadge(v int32) *PushTaskNotificationIos {
	s.Badge = &v
	return s
}

func (s *PushTaskNotificationIos) SetBadgeAutoIncrement(v bool) *PushTaskNotificationIos {
	s.BadgeAutoIncrement = &v
	return s
}

func (s *PushTaskNotificationIos) SetCategory(v string) *PushTaskNotificationIos {
	s.Category = &v
	return s
}

func (s *PushTaskNotificationIos) SetCollapseId(v string) *PushTaskNotificationIos {
	s.CollapseId = &v
	return s
}

func (s *PushTaskNotificationIos) SetExtParameters(v string) *PushTaskNotificationIos {
	s.ExtParameters = &v
	return s
}

func (s *PushTaskNotificationIos) SetInterruptionLevel(v string) *PushTaskNotificationIos {
	s.InterruptionLevel = &v
	return s
}

func (s *PushTaskNotificationIos) SetLiveActivity(v *PushTaskNotificationIosLiveActivity) *PushTaskNotificationIos {
	s.LiveActivity = v
	return s
}

func (s *PushTaskNotificationIos) SetMusic(v string) *PushTaskNotificationIos {
	s.Music = &v
	return s
}

func (s *PushTaskNotificationIos) SetMutable(v bool) *PushTaskNotificationIos {
	s.Mutable = &v
	return s
}

func (s *PushTaskNotificationIos) SetRelevanceScore(v float64) *PushTaskNotificationIos {
	s.RelevanceScore = &v
	return s
}

func (s *PushTaskNotificationIos) SetSilent(v bool) *PushTaskNotificationIos {
	s.Silent = &v
	return s
}

func (s *PushTaskNotificationIos) SetSubtitle(v string) *PushTaskNotificationIos {
	s.Subtitle = &v
	return s
}

func (s *PushTaskNotificationIos) SetThreadId(v string) *PushTaskNotificationIos {
	s.ThreadId = &v
	return s
}

func (s *PushTaskNotificationIos) Validate() error {
	if s.LiveActivity != nil {
		if err := s.LiveActivity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PushTaskNotificationIosLiveActivity struct {
	Attributes     *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	AttributesType *string `json:"AttributesType,omitempty" xml:"AttributesType,omitempty"`
	ContentState   *string `json:"ContentState,omitempty" xml:"ContentState,omitempty"`
	DismissalDate  *int64  `json:"DismissalDate,omitempty" xml:"DismissalDate,omitempty"`
	Event          *string `json:"Event,omitempty" xml:"Event,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	StaleDate      *int64  `json:"StaleDate,omitempty" xml:"StaleDate,omitempty"`
}

func (s PushTaskNotificationIosLiveActivity) String() string {
	return dara.Prettify(s)
}

func (s PushTaskNotificationIosLiveActivity) GoString() string {
	return s.String()
}

func (s *PushTaskNotificationIosLiveActivity) GetAttributes() *string {
	return s.Attributes
}

func (s *PushTaskNotificationIosLiveActivity) GetAttributesType() *string {
	return s.AttributesType
}

func (s *PushTaskNotificationIosLiveActivity) GetContentState() *string {
	return s.ContentState
}

func (s *PushTaskNotificationIosLiveActivity) GetDismissalDate() *int64 {
	return s.DismissalDate
}

func (s *PushTaskNotificationIosLiveActivity) GetEvent() *string {
	return s.Event
}

func (s *PushTaskNotificationIosLiveActivity) GetId() *string {
	return s.Id
}

func (s *PushTaskNotificationIosLiveActivity) GetStaleDate() *int64 {
	return s.StaleDate
}

func (s *PushTaskNotificationIosLiveActivity) SetAttributes(v string) *PushTaskNotificationIosLiveActivity {
	s.Attributes = &v
	return s
}

func (s *PushTaskNotificationIosLiveActivity) SetAttributesType(v string) *PushTaskNotificationIosLiveActivity {
	s.AttributesType = &v
	return s
}

func (s *PushTaskNotificationIosLiveActivity) SetContentState(v string) *PushTaskNotificationIosLiveActivity {
	s.ContentState = &v
	return s
}

func (s *PushTaskNotificationIosLiveActivity) SetDismissalDate(v int64) *PushTaskNotificationIosLiveActivity {
	s.DismissalDate = &v
	return s
}

func (s *PushTaskNotificationIosLiveActivity) SetEvent(v string) *PushTaskNotificationIosLiveActivity {
	s.Event = &v
	return s
}

func (s *PushTaskNotificationIosLiveActivity) SetId(v string) *PushTaskNotificationIosLiveActivity {
	s.Id = &v
	return s
}

func (s *PushTaskNotificationIosLiveActivity) SetStaleDate(v int64) *PushTaskNotificationIosLiveActivity {
	s.StaleDate = &v
	return s
}

func (s *PushTaskNotificationIosLiveActivity) Validate() error {
	return dara.Validate(s)
}

type PushTaskOptions struct {
	ExpireTime  *string             `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	JobKey      *string             `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	MessageId   *int64              `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	PushTime    *string             `json:"PushTime,omitempty" xml:"PushTime,omitempty"`
	Sms         *PushTaskOptionsSms `json:"Sms,omitempty" xml:"Sms,omitempty" type:"Struct"`
	Trim        *bool               `json:"Trim,omitempty" xml:"Trim,omitempty"`
	UseChannels *string             `json:"UseChannels,omitempty" xml:"UseChannels,omitempty"`
}

func (s PushTaskOptions) String() string {
	return dara.Prettify(s)
}

func (s PushTaskOptions) GoString() string {
	return s.String()
}

func (s *PushTaskOptions) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *PushTaskOptions) GetJobKey() *string {
	return s.JobKey
}

func (s *PushTaskOptions) GetMessageId() *int64 {
	return s.MessageId
}

func (s *PushTaskOptions) GetPushTime() *string {
	return s.PushTime
}

func (s *PushTaskOptions) GetSms() *PushTaskOptionsSms {
	return s.Sms
}

func (s *PushTaskOptions) GetTrim() *bool {
	return s.Trim
}

func (s *PushTaskOptions) GetUseChannels() *string {
	return s.UseChannels
}

func (s *PushTaskOptions) SetExpireTime(v string) *PushTaskOptions {
	s.ExpireTime = &v
	return s
}

func (s *PushTaskOptions) SetJobKey(v string) *PushTaskOptions {
	s.JobKey = &v
	return s
}

func (s *PushTaskOptions) SetMessageId(v int64) *PushTaskOptions {
	s.MessageId = &v
	return s
}

func (s *PushTaskOptions) SetPushTime(v string) *PushTaskOptions {
	s.PushTime = &v
	return s
}

func (s *PushTaskOptions) SetSms(v *PushTaskOptionsSms) *PushTaskOptions {
	s.Sms = v
	return s
}

func (s *PushTaskOptions) SetTrim(v bool) *PushTaskOptions {
	s.Trim = &v
	return s
}

func (s *PushTaskOptions) SetUseChannels(v string) *PushTaskOptions {
	s.UseChannels = &v
	return s
}

func (s *PushTaskOptions) Validate() error {
	if s.Sms != nil {
		if err := s.Sms.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PushTaskOptionsSms struct {
	DelaySecs    *int64  `json:"DelaySecs,omitempty" xml:"DelaySecs,omitempty"`
	Params       *string `json:"Params,omitempty" xml:"Params,omitempty"`
	SendPolicy   *string `json:"SendPolicy,omitempty" xml:"SendPolicy,omitempty"`
	SignName     *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s PushTaskOptionsSms) String() string {
	return dara.Prettify(s)
}

func (s PushTaskOptionsSms) GoString() string {
	return s.String()
}

func (s *PushTaskOptionsSms) GetDelaySecs() *int64 {
	return s.DelaySecs
}

func (s *PushTaskOptionsSms) GetParams() *string {
	return s.Params
}

func (s *PushTaskOptionsSms) GetSendPolicy() *string {
	return s.SendPolicy
}

func (s *PushTaskOptionsSms) GetSignName() *string {
	return s.SignName
}

func (s *PushTaskOptionsSms) GetTemplateName() *string {
	return s.TemplateName
}

func (s *PushTaskOptionsSms) SetDelaySecs(v int64) *PushTaskOptionsSms {
	s.DelaySecs = &v
	return s
}

func (s *PushTaskOptionsSms) SetParams(v string) *PushTaskOptionsSms {
	s.Params = &v
	return s
}

func (s *PushTaskOptionsSms) SetSendPolicy(v string) *PushTaskOptionsSms {
	s.SendPolicy = &v
	return s
}

func (s *PushTaskOptionsSms) SetSignName(v string) *PushTaskOptionsSms {
	s.SignName = &v
	return s
}

func (s *PushTaskOptionsSms) SetTemplateName(v string) *PushTaskOptionsSms {
	s.TemplateName = &v
	return s
}

func (s *PushTaskOptionsSms) Validate() error {
	return dara.Validate(s)
}

type PushTaskTarget struct {
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PushTaskTarget) String() string {
	return dara.Prettify(s)
}

func (s PushTaskTarget) GoString() string {
	return s.String()
}

func (s *PushTaskTarget) GetPlatform() *string {
	return s.Platform
}

func (s *PushTaskTarget) GetType() *string {
	return s.Type
}

func (s *PushTaskTarget) GetValue() *string {
	return s.Value
}

func (s *PushTaskTarget) SetPlatform(v string) *PushTaskTarget {
	s.Platform = &v
	return s
}

func (s *PushTaskTarget) SetType(v string) *PushTaskTarget {
	s.Type = &v
	return s
}

func (s *PushTaskTarget) SetValue(v string) *PushTaskTarget {
	s.Value = &v
	return s
}

func (s *PushTaskTarget) Validate() error {
	return dara.Validate(s)
}
