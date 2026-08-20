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
	// The push method. Optional parameter. Default value: `PUSH_IMMEDIATELY` (push immediately).
	//
	// example:
	//
	// PUSH_IMMEDIATELY
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The pass-through message data sent to the device. The total length cannot exceed 4,000 bytes.
	//
	// > Length calculation notes
	//
	// > - The length is calculated based on the byte length of the UTF-8 encoded string after the Message object is serialized to JSON.
	//
	// > - Chinese characters typically occupy 3 bytes in UTF-8 encoding.
	Message *PushTaskMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Struct"`
	// The vendor notification data sent to the device.
	//
	// 	Notice:
	//
	// When both `Message` and `Notification` are set, the device receives only one of them. The delivery rules are as follows:
	//
	// - When the device is online, the pass-through message data is delivered.
	//
	// - When the device is offline, the system notification is sent.
	Notification *PushTaskNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	// The push options.
	Options *PushTaskOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	// Specifies the target object for message push. This parameter is optional when the operation type `Action` is set to `CREATE_CONTINUOUS_PUSH` (create a continuous push task).
	Target *PushTaskTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
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
	// The body of the message to send.
	//
	// example:
	//
	// {"key": "value"}
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The title of the message to send.
	//
	// example:
	//
	// title
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
	// The Android notification configuration.
	Android *PushTaskNotificationAndroid `json:"Android,omitempty" xml:"Android,omitempty" type:"Struct"`
	// The body of the push notification.
	//
	// example:
	//
	// Dear customer, your reservation order has been successfully canceled
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The HarmonyOS notification configuration.
	Hmos *PushTaskNotificationHmos `json:"Hmos,omitempty" xml:"Hmos,omitempty" type:"Struct"`
	// The iOS notification configuration.
	Ios *PushTaskNotificationIos `json:"Ios,omitempty" xml:"Ios,omitempty" type:"Struct"`
	// The title of the push notification.
	//
	// > Length limits:
	//
	// > - iOS/Harmony: The **byte length*	- cannot exceed 200.
	//
	// > - Android: The **character length*	- cannot exceed 50.
	//
	// example:
	//
	// You have a new message
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
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
	// The full class name of the Activity for the badge setting application entry.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	BadgeActivity *string `json:"BadgeActivity,omitempty" xml:"BadgeActivity,omitempty"`
	// The incremental badge count value, which is added to the current badge count.
	//
	// > - Supported on `Huawei` and `Honor` channels.
	//
	// > - If both `BadgeAddNum` and `BadgeSetNum` are specified, `BadgeSetNum` takes precedence.
	//
	// example:
	//
	// 1
	BadgeAddNum *int32 `json:"BadgeAddNum,omitempty" xml:"BadgeAddNum,omitempty"`
	// The fixed badge number. Valid values: 1 to 99.
	//
	// example:
	//
	// 4
	BadgeSetNum *int32 `json:"BadgeSetNum,omitempty" xml:"BadgeSetNum,omitempty"`
	// The channelId of the Android app. This must match the channelId configured in the vendor app.
	//
	// example:
	//
	// 8.0up
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The custom extension attributes of the Android notification.
	//
	// example:
	//
	// {"key1":"value1"}
	ExtParameters *string `json:"ExtParameters,omitempty" xml:"ExtParameters,omitempty"`
	// The message group. Only the latest message and the total number of messages received in the group are displayed in the notification bar. All messages are not displayed and cannot be expanded. Currently supported channels:
	//
	// - Huawei channel
	//
	// - Honor channel
	//
	// - Chinese domestic channel with Android SDK 3.9.1 and earlier
	//
	// > The Chinese domestic channel no longer supports this parameter in Android SDK 3.9.2 and later.
	//
	// example:
	//
	// group-1
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The URL of the right-side icon. Currently supported:
	//
	// - `Huawei EMUI` (applicable only in long text mode and Inbox mode).
	//
	// - `Honor Magic UI` (applicable only in long text mode).
	//
	// - `Custom channel` (Android SDK 3.5.0 and later).
	//
	// example:
	//
	// https://imag.example.com/image.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The body content in Inbox mode. The value must be a valid JSON array with no more than 5 elements. Currently supported on:
	//
	// - Huawei: EMUI 9 and later
	//
	// - Honor: Magic UI 4.0 and later
	//
	// - Xiaomi: MIUI 10 and later
	//
	// - OPPO: ColorOS later than 5.0
	//
	// - Custom channel: Android SDK 3.6.0 and later
	InboxContent []*string `json:"InboxContent,omitempty" xml:"InboxContent,omitempty" type:"Repeated"`
	// The notification sound for the Huawei vendor channel. Specify the audio file name stored in the client project directory `app/src/main/res/raw/` without the file format extension. If not set, the default ringtone is used.
	//
	// example:
	//
	// alicloud_notification_sound
	Music *string `json:"Music,omitempty" xml:"Music,omitempty"`
	// The unique identifier of the Android notification bar message, used to control notification override and replacement behavior. A new notification with the same NotifyId automatically overrides the old notification.
	//
	// example:
	//
	// 233856727
	NotifyId *int32 `json:"NotifyId,omitempty" xml:"NotifyId,omitempty"`
	// The detailed channel configuration.
	Options *PushTaskNotificationAndroidOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	// The image URL in big picture mode. Currently supported: proprietary channel: Android SDK 3.6.0 and later.
	//
	// example:
	//
	// https://imag.example.com/image.png
	PictureUrl *string `json:"PictureUrl,omitempty" xml:"PictureUrl,omitempty"`
	// The notification style. Valid values:
	//
	// example:
	//
	// 0
	RenderStyle *string `json:"RenderStyle,omitempty" xml:"RenderStyle,omitempty"`
	// Specifies the notification type for the manufacturer channel. Valid values:
	//
	// - `false`: Production notification. This is the default value.
	//
	// - `true`: Test notification.
	//
	// > Currently supported: Huawei channel, Honor channel, vivo channel, and OPPO Fluid Cloud.
	//
	// example:
	//
	// false
	TestMessage *bool `json:"TestMessage,omitempty" xml:"TestMessage,omitempty"`
	// The Activity to open when the notification is tapped.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	VendorChannelActivity *string `json:"VendorChannelActivity,omitempty" xml:"VendorChannelActivity,omitempty"`
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
	// The Alibaba Cloud proprietary channel configuration.
	Accs *PushTaskNotificationAndroidOptionsAccs `json:"Accs,omitempty" xml:"Accs,omitempty" type:"Struct"`
	// The Honor channel configuration.
	Honor *PushTaskNotificationAndroidOptionsHonor `json:"Honor,omitempty" xml:"Honor,omitempty" type:"Struct"`
	// The Huawei channel configuration.
	Huawei *PushTaskNotificationAndroidOptionsHuawei `json:"Huawei,omitempty" xml:"Huawei,omitempty" type:"Struct"`
	// The Meizu channel configuration.
	Meizu *PushTaskNotificationAndroidOptionsMeizu `json:"Meizu,omitempty" xml:"Meizu,omitempty" type:"Struct"`
	// The OPPO channel configuration.
	Oppo *PushTaskNotificationAndroidOptionsOppo `json:"Oppo,omitempty" xml:"Oppo,omitempty" type:"Struct"`
	// The vivo channel configuration.
	Vivo *PushTaskNotificationAndroidOptionsVivo `json:"Vivo,omitempty" xml:"Vivo,omitempty" type:"Struct"`
	// The Xiaomi channel configuration.
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
	// The custom notification bar style for Android. Valid values: 1 to 100.
	//
	// > The style preset must be configured on the client. For more information, see [Custom notification style API](https://help.aliyun.com/document_detail/2834944.html).
	//
	// example:
	//
	// 1
	CustomStyle *int32 `json:"CustomStyle,omitempty" xml:"CustomStyle,omitempty"`
	// The notification alert type. Valid values:
	//
	// - `VIBRATE`: vibration (default)
	//
	// - `SOUND`: sound
	//
	// - `BOTH`: sound and vibration
	//
	// - `NONE`: silent
	//
	// example:
	//
	// NONE
	NotifyType *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// The activity to open when the notification is tapped. This parameter takes effect only when `OpenType` is set to `ACTIVITY`.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	OpenActivity *string `json:"OpenActivity,omitempty" xml:"OpenActivity,omitempty"`
	// The action after tapping the notification. Valid values:
	//
	// example:
	//
	// APPLICATION
	OpenType *string `json:"OpenType,omitempty" xml:"OpenType,omitempty"`
	// The URL to open when the notification is tapped on Android. This is valid when `OpenType` is set to `URL`.
	//
	// example:
	//
	// www.example.com
	OpenUrl *string `json:"OpenUrl,omitempty" xml:"OpenUrl,omitempty"`
	// The priority of the Android notification position in the notification bar. Valid values: -2, -1, 0, 1, 2.
	//
	// example:
	//
	// 0
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The message group. Messages in the same group are collapsed in the notification bar and can be expanded. Messages in different groups are displayed separately.
	//
	// > Android SDK 3.9.2 and later
	//
	// example:
	//
	// order_ORD20231201001
	ThreadId *string `json:"ThreadId,omitempty" xml:"ThreadId,omitempty"`
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
	// Specifies the importance parameter for Honor notification message classification, which determines the notification behavior on the user\\"s device. Valid values:
	//
	// - `0`: informational and marketing messages
	//
	// - `1`: service and communication messages
	//
	// You must apply for this on the Honor platform. [Application link](https://developer.honor.com/cn/docs/11002/guides/notification-class#%E8%87%AA%E5%88%86%E7%B1%BB%E6%9D%83%E7%9B%8A%E7%94%B3%E8%AF%B7).
	//
	// example:
	//
	// 0
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
	// The Huawei quick notification parameter.
	//
	// example:
	//
	// 1
	BusinessType *int32 `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// Purpose 1: After completing the [self-classification privilege](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835?#section3410731125514) application, this parameter identifies the message type, determines the [notification method](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#ZH-CN_TOPIC_0000001149358835__p3850133955718), and accelerates delivery for specific message types. For valid values, refer to the [message classification standard](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section1076611477914) in the official Huawei Push documentation. Use the value from the "Cloud notification category value" or "Local notification category value" column in the table.
	//
	// Purpose 2: After [applying for special permissions](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/faq-0000001050042183#section037425218509), this parameter identifies high-priority pass-through scenarios. Valid values:
	//
	// - `VOIP`: audio and video calls
	//
	// - `PLAY_VOICE`: voice broadcast
	//
	// > - For messages where the "Cloud notification category value" is "Not applicable", messages are sent through the Alibaba Cloud proprietary channel.
	//
	// > - For messages where the "Local notification category value" is "Not applicable", messages are sent through the Huawei channel.
	//
	// example:
	//
	// VOIP
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The importance parameter for Huawei notification message classification, which determines the notification behavior on the user device. Valid values:
	//
	// example:
	//
	// 0
	Importance *int32 `json:"Importance,omitempty" xml:"Importance,omitempty"`
	// The JSON string of the Huawei Android Live Notification data structure [LiveNotificationPayload](https://developer.huawei.com/consumer/cn/doc/HMSCore-References/rest-live-0000001562939968#ZH-CN_TOPIC_0000001700850537__p195121620102511). For development and integration, refer to [Huawei Live Notification Push Guide](https://help.aliyun.com/document_detail/2983768.html).
	//
	// example:
	//
	// {
	//
	//   "activityId": 1,
	//
	//   "operation": 1,
	//
	//   "event": "TAXI",
	//
	//   "activityData": {
	//
	//     "notificationData": {
	//
	//       "type": 3
	//
	//     }
	//
	//   }
	//
	// }
	LiveNotificationPayload *string `json:"LiveNotificationPayload,omitempty" xml:"LiveNotificationPayload,omitempty"`
	// The receipt ID of the Huawei channel. You can view this receipt ID in the receipt parameter configuration on the Huawei channel push operation platform.
	//
	// > If the default receipt configuration on the Huawei channel push operation platform is set to Alibaba Cloud receipt, you do not need to provide this parameter. If not, configure the default Huawei channel receipt ID in the Alibaba Cloud EMAS Mobile Push console first.
	//
	// example:
	//
	// RCP4C123456
	ReceiptId *string `json:"ReceiptId,omitempty" xml:"ReceiptId,omitempty"`
	// The delivery priority of the Huawei channel notification. Valid values:
	//
	// example:
	//
	// NORMAL
	Urgency *string `json:"Urgency,omitempty" xml:"Urgency,omitempty"`
}

func (s PushTaskNotificationAndroidOptionsHuawei) String() string {
	return dara.Prettify(s)
}

func (s PushTaskNotificationAndroidOptionsHuawei) GoString() string {
	return s.String()
}

func (s *PushTaskNotificationAndroidOptionsHuawei) GetBusinessType() *int32 {
	return s.BusinessType
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

func (s *PushTaskNotificationAndroidOptionsHuawei) SetBusinessType(v int32) *PushTaskNotificationAndroidOptionsHuawei {
	s.BusinessType = &v
	return s
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
	// The Meizu message type.
	//
	// example:
	//
	// 0
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
	// OPPO categorizes messages into two types for management: Communication & Service, and Content & Marketing.
	//
	// example:
	//
	// NEWS
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The JSON character string of the OPPO Fluid Cloud intent delete data structure [data](https://open.oppomobile.com/documentation/page/info?id=13578). This parameter is invalid when the AndroidOppoIntelligentIntent parameter is already specified. References: [OPPO Fluid Cloud Push Guide](https://help.aliyun.com/document_detail/2997310.html).
	//
	// example:
	//
	// {
	//
	//     "intentName": "Example.Progress",
	//
	//     "entityIds": [
	//
	//         "A580202509130712"
	//
	//     ],
	//
	//     "serviceId": {
	//
	//         "launcher": "999800001",
	//
	//         "fluidCloud": "999900001"
	//
	//     }
	//
	// }
	DeleteIntentData *string `json:"DeleteIntentData,omitempty" xml:"DeleteIntentData,omitempty"`
	// The JSON character string of the OPPO Fluid Cloud intent sharing data structure [IntelligentIntent](https://open.oppomobile.com/documentation/page/info?id=13565). References: [OPPO Fluid Cloud Push Guide](https://help.aliyun.com/document_detail/2997310.html).
	//
	// example:
	//
	// {
	//
	//     "intentName": "Example.Progress",
	//
	//     "identifier": "d71ebd3119877b12ecdb6c4fe96b068e",
	//
	//     "timestamp": 1729485000989,
	//
	//     "serviceId": {
	//
	//         "launcher": "999800001",
	//
	//         "fluidCloud": "999900001"
	//
	//     },
	//
	//     "intentAction": {
	//
	//         "actionStatus": 0
	//
	//     },
	//
	//     "intentEntity": {
	//
	//         "entityName": "TAXI"
	//
	//     }
	//
	// }
	IntelligentIntent *string `json:"IntelligentIntent,omitempty" xml:"IntelligentIntent,omitempty"`
	// The notification bar message alert level for the OPPO channel. Valid values:
	//
	// example:
	//
	// 1
	NotifyLevel *int64 `json:"NotifyLevel,omitempty" xml:"NotifyLevel,omitempty"`
	// The OPPO private message template content parameters.
	//
	// example:
	//
	// {
	//
	// "key1": "value1",
	//
	// "key2": "value2"
	//
	// }
	PrivateContentParameters *string `json:"PrivateContentParameters,omitempty" xml:"PrivateContentParameters,omitempty"`
	// The OPPO private message template ID.
	//
	// example:
	//
	// 687557242b1634hzefs3d5013
	PrivateMsgTemplateId *string `json:"PrivateMsgTemplateId,omitempty" xml:"PrivateMsgTemplateId,omitempty"`
	// The OPPO private message template title parameters.
	//
	// example:
	//
	// {"name": "John"}
	PrivateTitleParameters *string `json:"PrivateTitleParameters,omitempty" xml:"PrivateTitleParameters,omitempty"`
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
	AddBadge *bool `json:"AddBadge,omitempty" xml:"AddBadge,omitempty"`
	// vivo categorizes messages into two types: system messages and operational messages.
	//
	// **System messages:**
	//
	// - IM: instant messaging
	//
	// - ACCOUNT: accounts and assets
	//
	// - TODO: schedules and to-do items
	//
	// - DEVICE_REMINDER: device information
	//
	// - ORDER: orders and logistics
	//
	// - SUBSCRIPTION: subscription reminders
	//
	// **Operational messages:**
	//
	// - NEWS: news
	//
	// - CONTENT: content recommendation
	//
	// - MARKETING: operational activity
	//
	// - SOCIAL: social updates
	//
	// For more information, refer to [vivo category description](https://dev.vivo.com.cn/documentCenter/doc/359#s-ef3qugc3).
	//
	// example:
	//
	// MARKETING
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies the vivo notification message category. Valid values:
	//
	// - `0`: Operational message (default).
	//
	// - `1`: System message.
	//
	// > Use `Category` for notification classification. You need to apply on the vivo platform. For more information, see [Application link](https://dev.vivo.com.cn/documentCenter/doc/359).
	//
	// example:
	//
	// 0
	Importance *int32 `json:"Importance,omitempty" xml:"Importance,omitempty"`
	// The JSON character string of the vivo Atomic Island data structure [liveMessage](https://dev.vivo.com.cn/documentCenter/doc/896#s-fdagzbd4). References: [vivo Atomic Island Push Guide](https://www.alibabacloud.com/help/en/document_detail/3030718.html).
	//
	// example:
	//
	// {
	//
	//     "operation": 0,
	//
	//     "scene": "HEALTH_REGISTER",
	//
	//     "templateType": 1,
	//
	//     "showNotify": true,
	//
	//     "changeRecord": 999,
	//
	//     "capsuleData": {
	//
	//         "bgColor": "#32d4d4"
	//
	//     }
	//
	// }
	LiveMessage *string `json:"LiveMessage,omitempty" xml:"LiveMessage,omitempty"`
	// The message receipt identifier for the vivo vendor push channel, used to receive push result callback notifications.
	//
	// example:
	//
	// 1232221
	ReceiptId *string `json:"ReceiptId,omitempty" xml:"ReceiptId,omitempty"`
}

func (s PushTaskNotificationAndroidOptionsVivo) String() string {
	return dara.Prettify(s)
}

func (s PushTaskNotificationAndroidOptionsVivo) GoString() string {
	return s.String()
}

func (s *PushTaskNotificationAndroidOptionsVivo) GetAddBadge() *bool {
	return s.AddBadge
}

func (s *PushTaskNotificationAndroidOptionsVivo) GetCategory() *string {
	return s.Category
}

func (s *PushTaskNotificationAndroidOptionsVivo) GetImportance() *int32 {
	return s.Importance
}

func (s *PushTaskNotificationAndroidOptionsVivo) GetLiveMessage() *string {
	return s.LiveMessage
}

func (s *PushTaskNotificationAndroidOptionsVivo) GetReceiptId() *string {
	return s.ReceiptId
}

func (s *PushTaskNotificationAndroidOptionsVivo) SetAddBadge(v bool) *PushTaskNotificationAndroidOptionsVivo {
	s.AddBadge = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsVivo) SetCategory(v string) *PushTaskNotificationAndroidOptionsVivo {
	s.Category = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsVivo) SetImportance(v int32) *PushTaskNotificationAndroidOptionsVivo {
	s.Importance = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsVivo) SetLiveMessage(v string) *PushTaskNotificationAndroidOptionsVivo {
	s.LiveMessage = &v
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
	// The channel ID for Xiaomi notification types. You must apply for this on the Xiaomi platform. For more information, see [Application link](https://dev.mi.com/console/doc/detail?pId=2422#_4).
	//
	// > A single application can apply for a maximum of 8 channels on the Xiaomi channel. Plan ahead.
	//
	// example:
	//
	// michannel
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The JSON character string of the Xiaomi Super Island data structure [miui.focus.param](https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131). References: [Xiaomi Super Island Push Guide](https://www.alibabacloud.com/help/en/document_detail/3037956.html).
	//
	// example:
	//
	// {
	//
	//     "param_v2": {
	//
	//         "business": "taxi",
	//
	//         "updatable": true,
	//
	//         "orderId": "A580202509130712",
	//
	//         "param_island": {
	//
	//             "islandProperty": 1,
	//
	//             "bigIslandArea": {
	//
	//                 "imageTextInfoLeft": {
	//
	//                     "type": 1
	//
	//                 }
	//
	//             }
	//
	//         }
	//
	//     }
	//
	// }
	FocusParam *string `json:"FocusParam,omitempty" xml:"FocusParam,omitempty"`
	// The JSON character string of the Xiaomi Super Island image data [miui.focus.pic_xxx](https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131). References: [Xiaomi Super Island Push Guide](https://www.alibabacloud.com/help/en/document_detail/3037956.html).
	//
	// example:
	//
	// {
	//
	//     "miui.focus.pic_ticker": "https://example.com/ticker.jpg",
	//
	//     "miui.focus.pic_aod": "https://example.com/aod.jpg",
	//
	//     "miui.focus.pic_imageText": "https://example.com/imageText.jpg"
	//
	// }
	FocusPics *string `json:"FocusPics,omitempty" xml:"FocusPics,omitempty"`
	// The Xiaomi private message template ID.
	//
	// example:
	//
	// P10645
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The Xiaomi private message template parameters in JSON string format.
	//
	// example:
	//
	// {"keywords1":"Tom","keywords2":"phone"}
	TemplateParams *string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
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

func (s *PushTaskNotificationAndroidOptionsXiaomi) GetFocusParam() *string {
	return s.FocusParam
}

func (s *PushTaskNotificationAndroidOptionsXiaomi) GetFocusPics() *string {
	return s.FocusPics
}

func (s *PushTaskNotificationAndroidOptionsXiaomi) GetTemplateId() *string {
	return s.TemplateId
}

func (s *PushTaskNotificationAndroidOptionsXiaomi) GetTemplateParams() *string {
	return s.TemplateParams
}

func (s *PushTaskNotificationAndroidOptionsXiaomi) SetChannel(v string) *PushTaskNotificationAndroidOptionsXiaomi {
	s.Channel = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsXiaomi) SetFocusParam(v string) *PushTaskNotificationAndroidOptionsXiaomi {
	s.FocusParam = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsXiaomi) SetFocusPics(v string) *PushTaskNotificationAndroidOptionsXiaomi {
	s.FocusPics = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsXiaomi) SetTemplateId(v string) *PushTaskNotificationAndroidOptionsXiaomi {
	s.TemplateId = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsXiaomi) SetTemplateParams(v string) *PushTaskNotificationAndroidOptionsXiaomi {
	s.TemplateParams = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsXiaomi) Validate() error {
	return dara.Validate(s)
}

type PushTaskNotificationHmos struct {
	// The action that corresponds to the ability of the in-app page.
	//
	// > For more information, refer to [ClickAction.action](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section152462191216) on the HarmonyOS official website.
	//
	// example:
	//
	// com.example.action
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The incremental badge number for HarmonyOS applications.
	//
	// > - Supported since HarmonyOS SDK 1.2.0.
	//
	// > - Refer to the HarmonyOS badge [addNum field description](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section266310382145).
	//
	// example:
	//
	// 1
	BadgeAddNum *int32 `json:"BadgeAddNum,omitempty" xml:"BadgeAddNum,omitempty"`
	// The number to set for the HarmonyOS app badge.
	//
	// > - Refer to the HarmonyOS badge [setNum field](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section266310382145) description.
	//
	// > - Supported since HarmonyOS SDK version 1.2.0.
	//
	// example:
	//
	// 1
	BadgeSetNum *int32 `json:"BadgeSetNum,omitempty" xml:"BadgeSetNum,omitempty"`
	// The category of the notification message. This is an optional parameter. Default value: `MARKETING`.
	//
	// > After you complete the application for the notification message self-classification privilege, this parameter identifies the message type. Different notification message types affect how messages are displayed and how reminders are triggered. For more information, refer to [Notification.category](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117) on the HarmonyOS official website.
	//
	// example:
	//
	// IM
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The custom extension attributes of the notification message, used to pass additional business data.
	//
	// example:
	//
	// {"key": "value"}
	ExtParameters *string `json:"ExtParameters,omitempty" xml:"ExtParameters,omitempty"`
	// The extra data of the notification extension message.
	//
	// > - Valid when sending HarmonyOS notification extension messages.
	//
	// > - Conceptually equivalent to the extraData field of HarmonyOS notification extension messages. For the specific definition, refer to the HarmonyOS [ExtensionPayload](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section161192514234) documentation.
	//
	// > - Supported since HarmonyOS SDK 1.2.0.
	//
	// example:
	//
	// text
	ExtensionExtraData *string `json:"ExtensionExtraData,omitempty" xml:"ExtensionExtraData,omitempty"`
	// Enables HarmonyOS notification extension.
	//
	// > - To send notification extension messages, you must first apply for permissions on the HarmonyOS official website. For more information, refer to [HarmonyOS documentation](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V5/push-send-extend-noti-V5) on sending notification extension messages.
	//
	// > - Supported starting from HarmonyOS SDK 1.2.0.
	//
	// example:
	//
	// false
	ExtensionPush *bool `json:"ExtensionPush,omitempty" xml:"ExtensionPush,omitempty"`
	// The URL of the large icon displayed on the right side of the notification. The URL must use the HTTPS protocol.
	//
	// > - Supported image formats include png, jpg, jpeg, heif, gif, and bmp. The image length × width must be less than 25000 pixels.
	//
	// > - For more information, refer to the HarmonyOS official documentation [Notification.image](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117).
	//
	// example:
	//
	// https://example.com/xxx.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// When `RenderStyle` is set to `MULTI_LINE`, this field is required to define the content in multi-line text style. A maximum of 3 items are supported.
	InboxContent []*string `json:"InboxContent,omitempty" xml:"InboxContent,omitempty" type:"Repeated"`
	// The JSON string of the HarmonyOS Live View data structure [LiveViewPayload](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V13/push-scenariozed-api-request-param-V13#section66881469306). For development and integration, refer to [HarmonyOS Live View Push Guide](https://help.aliyun.com/document_detail/2982112.html).
	//
	// example:
	//
	// {
	//
	//   "activityId": 1,
	//
	//   "operation": 0,
	//
	//   "event": "TAXI",
	//
	//   "status": "DRIVER_ON_THE_WAY",
	//
	//   "activityData": {
	//
	//     "notificationData": {
	//
	//       "type": 3
	//
	//     }
	//
	//   }
	//
	// }
	LiveViewPayload *string `json:"LiveViewPayload,omitempty" xml:"LiveViewPayload,omitempty"`
	// The unique identifier (notifyId) for each message displayed in the notification bar. If not provided, the push service automatically generates a unique identifier. Different notification messages can use the same notifyId to enable new messages to overwrite old messages. For more information, see [Notification.notifyId](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117) on the HarmonyOS official website.
	//
	// example:
	//
	// 123456
	NotifyId *int32 `json:"NotifyId,omitempty" xml:"NotifyId,omitempty"`
	// The receipt ID of the HarmonyOS channel. You can view this receipt ID in the receipt parameter settings on the HarmonyOS channel push operation platform.
	//
	// example:
	//
	// RCPB***DFD5
	ReceiptId *string `json:"ReceiptId,omitempty" xml:"ReceiptId,omitempty"`
	// The notification message style. This is an optional parameter. Default value: normal notification.
	//
	// example:
	//
	// NORMAL
	RenderStyle *string `json:"RenderStyle,omitempty" xml:"RenderStyle,omitempty"`
	// Specifies the notification channel type to use.
	//
	// > - Valid only for the Alibaba Cloud proprietary channel.
	//
	// > - For more information, refer to the HarmonyOS official documentation [SlotType](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/js-apis-notificationmanager-V5#slottype).
	//
	// example:
	//
	// SOCIAL_COMMUNICATION
	SlotType *string `json:"SlotType,omitempty" xml:"SlotType,omitempty"`
	// The HarmonyOS custom ringtone file name.
	//
	// example:
	//
	// music.mp3
	Sound *string `json:"Sound,omitempty" xml:"Sound,omitempty"`
	// The custom notification ringtone duration in seconds. Valid values: 1 to 60. The ringtone loops if its duration is shorter than the specified value.
	//
	// example:
	//
	// 2
	SoundDuration *int32 `json:"SoundDuration,omitempty" xml:"SoundDuration,omitempty"`
	// Enables the test message.
	//
	// > - For more information, refer to the HarmonyOS push parameter [TestMessage](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section418321011212).
	//
	// example:
	//
	// true
	TestMessage *bool `json:"TestMessage,omitempty" xml:"TestMessage,omitempty"`
	// The URI that corresponds to the in-app page ability.
	//
	// > - When multiple Abilities exist, specify the action and URI for each Ability separately. The action is used first to find the corresponding in-app page.
	//
	// > - For more information, see [ClickAction.uri](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section152462191216) on the HarmonyOS official website.
	//
	// example:
	//
	// https://www.example.com:8080/push/example
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
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

func (s *PushTaskNotificationHmos) GetSound() *string {
	return s.Sound
}

func (s *PushTaskNotificationHmos) GetSoundDuration() *int32 {
	return s.SoundDuration
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

func (s *PushTaskNotificationHmos) SetSound(v string) *PushTaskNotificationHmos {
	s.Sound = &v
	return s
}

func (s *PushTaskNotificationHmos) SetSoundDuration(v int32) *PushTaskNotificationHmos {
	s.SoundDuration = &v
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
	// iOS notifications are sent through the APNs center. You need to specify the corresponding environment information. Optional parameter. Default value: production environment.
	//
	// example:
	//
	// DEV
	ApnsEnv *string `json:"ApnsEnv,omitempty" xml:"ApnsEnv,omitempty"`
	// The iOS application badge number.
	//
	// example:
	//
	// 1
	Badge *int32 `json:"Badge,omitempty" xml:"Badge,omitempty"`
	// Specifies whether to enable the badge auto-increment feature. Optional parameter. Default value: false.
	//
	// example:
	//
	// false
	BadgeAutoIncrement *bool `json:"BadgeAutoIncrement,omitempty" xml:"BadgeAutoIncrement,omitempty"`
	// The category identifier for the iOS notification, which defines the interaction behavior and display style of the notification.
	//
	// > - The category must be pre-registered in the app to take effect.
	//
	// > - Different categories can define different sets of actions.
	//
	// example:
	//
	// MESSAGE_REPLY
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The unique identifier for notification collapsing. Notifications with the same identifier are overwritten and displayed as one.
	//
	// example:
	//
	// order_status_update_12345
	CollapseId *string `json:"CollapseId,omitempty" xml:"CollapseId,omitempty"`
	// The custom extension attributes of the iOS notification.
	//
	// example:
	//
	// {"attachment": "https://xxxx.xxx/notification_pic.png"}
	ExtParameters *string `json:"ExtParameters,omitempty" xml:"ExtParameters,omitempty"`
	// The interruption level. Optional parameter. Valid values:
	//
	// example:
	//
	// active
	InterruptionLevel *string `json:"InterruptionLevel,omitempty" xml:"InterruptionLevel,omitempty"`
	// The Live Activity parameter object.
	//
	// 	Notice:
	//
	// - Live Activity push notifications can only be sent to a **single device*	- by specifying the `DEVICE` type.
	//
	// - When pushing Live Activity notifications, the title and body parameters are optional.
	LiveActivity *PushTaskNotificationIosLiveActivity `json:"LiveActivity,omitempty" xml:"LiveActivity,omitempty" type:"Struct"`
	// The notification sound for iOS. Specify the name of an audio file stored in the app bundle or the Library/Sounds directory of the sandbox. For more information, see [How to set notification sounds for iOS push](https://help.aliyun.com/document_detail/48906.html).
	//
	// > - If set to an empty string (""), the notification is silent.
	//
	// > - If not specified, the value defaults to "default", which plays the system alert sound.
	//
	// example:
	//
	// default
	Music *string `json:"Music,omitempty" xml:"Music,omitempty"`
	// Specifies whether to enable the notification extension, which controls whether iOS notifications support processing by Notification Service Extension.
	//
	// > - When sending silent notifications, this parameter must be set to true.
	//
	// > - The Extension processing time cannot exceed 30 seconds.
	//
	// > - A timeout causes the notification to display the original content.
	//
	// > - You must add a Notification Service Extension to your application.
	//
	// example:
	//
	// true
	Mutable *bool `json:"Mutable,omitempty" xml:"Mutable,omitempty"`
	// The relevance score of the notification message, used to control the priority and display strategy of the notification.
	//
	// example:
	//
	// 0.5
	RelevanceScore *float64 `json:"RelevanceScore,omitempty" xml:"RelevanceScore,omitempty"`
	// Specifies whether to enable silent push mode.
	//
	// example:
	//
	// false
	Silent *bool `json:"Silent,omitempty" xml:"Silent,omitempty"`
	// The subtitle content of the iOS notification.
	//
	// example:
	//
	// Please check your order
	Subtitle *string `json:"Subtitle,omitempty" xml:"Subtitle,omitempty"`
	// The thread identifier for iOS notification grouping, which is used to categorize and collapse related notifications.
	//
	// > - Notifications with the same thread-id are automatically grouped together.
	//
	// > - Multiple related notifications are collapsed into a single notification group.
	//
	// > - Users can expand the group to view all notifications within it.
	//
	// example:
	//
	// news_category_tech
	ThreadId *string `json:"ThreadId,omitempty" xml:"ThreadId,omitempty"`
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
	// The static pass-through parameter for iOS Live Activities push notifications, used to pass immutable business identifier information.
	//
	// > Required when `Event` is set to start.
	//
	// example:
	//
	// {
	//
	//   "orderId": "ORD20231201001",
	//
	//   "restaurantName": "Delicious Restaurant",
	//
	//   "customerAddress": "No. xx, xx Road, xx District",
	//
	//   "orderType": "delivery"
	//
	// }
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The type of the Live Activity to start.
	//
	// example:
	//
	// OrderActivityAttributes
	AttributesType *string `json:"AttributesType,omitempty" xml:"AttributesType,omitempty"`
	// The dynamic pass-through parameters of the Live Activity, containing real-time updatable status information and changing data.
	//
	// example:
	//
	// {
	//
	//     "status": "delivering",
	//
	//     "estimatedTime": "10 minutes",
	//
	//     "progress": 80,
	//
	//     "driverName": "Driver Li",
	//
	//     "currentStep": "The delivery driver is on the way"}
	//
	// }
	ContentState *string `json:"ContentState,omitempty" xml:"ContentState,omitempty"`
	// The retention time of an ended Live Activity on the lock screen, allowing users to view information after the activity ends. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1701439800
	DismissalDate *int64 `json:"DismissalDate,omitempty" xml:"DismissalDate,omitempty"`
	// Starts, updates, or ends a Live Activity.
	//
	// example:
	//
	// start
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// The unique identifier of the Live Activity, used to associate the device-side activity instance with the server-side push target.
	//
	// example:
	//
	// FOOD_DELIVERY_ORD20231201001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The expiration timestamp for the iOS Live Activity content, specified as a Unix timestamp in seconds.
	//
	// > - After the specified time is reached, the system automatically marks the activity as expired.
	//
	// > - Expired activities are removed from the Dynamic Island and Lock Screen.
	//
	// > - This prevents outdated information from occupying the user interface for an extended period.
	//
	// example:
	//
	// 1701425400
	StaleDate *int64 `json:"StaleDate,omitempty" xml:"StaleDate,omitempty"`
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
	// The expiration time of the message. The message will not be sent after it expires. Messages can be retained for up to 72 hours.
	//
	// > 	- The time follows the ISO 8601 standard in UTC. Format: YYYY-MM-DDThh:mm:ssZ.
	//
	// > 	- The expiration time must meet the following condition: ExpireTime > PushTime + 3 seconds (3 seconds is the redundancy for network and system latency).
	//
	// > 	- Recommendation: Set the expiration time to at least 1 minute for single push notifications and at least 10 minutes for full push or batch push notifications.
	//
	//
	// 	Notice: For pass-through messages, if no expiration time is set, the message is sent only to online devices. When the device is offline, the message is discarded.
	//
	// example:
	//
	// 2025-06-21T12:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The custom identifier for the push task. When JobKey is not empty, this field is included in the receipt log. For more information about receipt logs, see [Receipt logs](https://help.aliyun.com/document_detail/434651.html).
	//
	// example:
	//
	// jobkey1727749697913
	JobKey *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	// The unique ID used to identify the message. This parameter is valid only when the `Action` parameter is set to `CONTINUOUS_PUSH`.
	//
	// example:
	//
	// 1174754033128****
	MessageId *int64 `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The scheduled time to send the message. The value cannot be later than 7 days from the current time. This parameter takes effect only when `Action` is set to `SCHEDULED_PUSH`.
	//
	// > The time follows the ISO 8601 standard in UTC in the format of yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2025-06-19T12:00:00Z
	PushTime *string `json:"PushTime,omitempty" xml:"PushTime,omitempty"`
	// The supplementary SMS settings.
	Sms *PushTaskOptionsSms `json:"Sms,omitempty" xml:"Sms,omitempty" type:"Struct"`
	// Specifies whether to automatically truncate titles and content that exceed the length limit.
	//
	// >This parameter applies only to vendor channels that explicitly limit the title and content length. It does not apply to channels such as APNs, Huawei, and Honor that do not limit the title or content length but only limit the total request body size.
	//
	// example:
	//
	// false
	Trim *bool `json:"Trim,omitempty" xml:"Trim,omitempty"`
	// Specifies the delivery channels. Valid values:
	//
	// - `accs`: Alibaba Cloud proprietary channel
	//
	// - `huawei`: Huawei channel
	//
	// - `honor`: Honor channel
	//
	// - `xiaomi`: Xiaomi channel
	//
	// - `oppo`: OPPO channel
	//
	// - `vivo`: vivo channel
	//
	// - `meizu`: Meizu channel
	//
	// - `fcm`: Google Firebase channel (HTTP v1 API)
	//
	// - `apns`: APNs channel
	//
	// - `harmony`: HarmonyOS channel
	//
	// > - If this parameter is not specified, all channels are available.
	//
	// > - If this parameter is specified, only the specified channels are used.
	//
	// > - If the specified channels conflict with the delivery policy (for example, iOS notifications can only be delivered through the APNs channel, but apns is not included in this parameter), the message is not delivered.
	//
	// example:
	//
	// accs,apns
	UseChannels *string `json:"UseChannels,omitempty" xml:"UseChannels,omitempty"`
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
	// The delay before triggering the SMS message. Unit: seconds.
	//
	// This parameter is required when SMS linkage is used. We recommend that you set this parameter to at least 15 seconds and no more than 3 days to avoid duplicate notifications from both SMS and push.
	//
	// > When SMS linkage is used, the ExpireTime parameter does not take effect. The notification expiration time is calculated based on the DelaySecs parameter. The expiration time is the current time plus the DelaySecs value.
	//
	// example:
	//
	// 150
	DelaySecs *int64 `json:"DelaySecs,omitempty" xml:"DelaySecs,omitempty"`
	// The key-value pairs of variable names in the SMS template.
	//
	// example:
	//
	// key1=value1&key2=value2
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The SMS sending policy.
	//
	// example:
	//
	// PUSH_NOT_RECEIVED
	SendPolicy *string `json:"SendPolicy,omitempty" xml:"SendPolicy,omitempty"`
	// The SMS signature.
	//
	// example:
	//
	// SampleTech
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The SMS template name. You can obtain this name from the SMS template management page. This is the system-assigned name, not the name set by the developer.
	//
	// example:
	//
	// SMS_123456789
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
	// The platform type. Optional parameter.
	//
	// example:
	//
	// IOS
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The push target type.
	//
	// 	Notice:
	//
	// The batch push operation `MassPushV2` and continuous push `CONTINUOUS_PUSH` support only the following three target types:
	//
	// - `DEVICE`
	//
	// - `ACCOUNT`
	//
	// - `ALIAS`
	//
	// example:
	//
	// DEVICE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The push target based on `Target.Type`. Separate multiple targets with commas. The following describes the target types and target values:
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
