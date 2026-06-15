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
	// The push method. This is an optional parameter. The default value is `PUSH_IMMEDIATELY` (immediate push).
	//
	// 	Notice:
	//
	// The `MassPushV2` batch push API supports only the following push methods:
	//
	// - `PUSH_IMMEDIATELY` (immediate push)
	//
	// - `SCHEDULED_PUSH` (scheduled push)
	//
	// example:
	//
	// PUSH_IMMEDIATELY
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The pass-through message data sent to the device. The total length cannot exceed 4,000 bytes.
	//
	// > Length calculation
	//
	// >
	//
	// > - The length is calculated based on the byte length of the UTF-8 encoded string after the Message object is serialized into JSON.
	//
	// >
	//
	// > - A Chinese character typically occupies 3 bytes in UTF-8 encoding.
	Message *PushTaskMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Struct"`
	// The vendor notification data sent to the device.
	//
	// 	Notice:
	//
	// If you set both `Message` and `Notification`, the device receives only one. The sending rules are as follows:
	//
	// - If the device is online, pass-through message data is sent.
	//
	// - If the device is offline, a system notification is sent.
	Notification *PushTaskNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	// Push options
	Options *PushTaskOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	// The target object for the message push. This parameter is optional when the `Action` operation type is `CREATE_CONTINUOUS_PUSH` (create a continuous push task).
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
	// The content of the message to send.
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
	// Android notification configuration
	Android *PushTaskNotificationAndroid `json:"Android,omitempty" xml:"Android,omitempty" type:"Struct"`
	// The content of the push notification.
	//
	// > The length limits are as follows:
	//
	// >
	//
	// > - For iOS, HarmonyOS, and Android, the character length cannot exceed 200.
	//
	// example:
	//
	// 尊敬的客户，您好！您的预约订单已取消成功。
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// HarmonyOS notification configuration.
	Hmos *PushTaskNotificationHmos `json:"Hmos,omitempty" xml:"Hmos,omitempty" type:"Struct"`
	// iOS notification configuration
	Ios *PushTaskNotificationIos `json:"Ios,omitempty" xml:"Ios,omitempty" type:"Struct"`
	// The title of the push notification.
	//
	// > The length limits are as follows:
	//
	// >
	//
	// > - For iOS/HarmonyOS, the byte length cannot exceed 200.
	//
	// >
	//
	// > - For Android, the character length cannot exceed 50.
	//
	// example:
	//
	// 您有一条新消息
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
	// The full class name of the application entry Activity for badge settings.
	//
	// > This is only valid when pushing through the Huawei or Honor vendor channel.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	BadgeActivity *string `json:"BadgeActivity,omitempty" xml:"BadgeActivity,omitempty"`
	// Sets a cumulative value for the badge, which is added to the original badge number.
	//
	// > - This is supported by `Huawei` and `Honor` channels.
	//
	// >
	//
	// > - If both `BadgeAddNum` and `BadgeSetNum` are present, the latter takes precedence.
	//
	// example:
	//
	// 1
	BadgeAddNum *int32 `json:"BadgeAddNum,omitempty" xml:"BadgeAddNum,omitempty"`
	// Sets a fixed value for the badge number. The value range is [1, 99].
	//
	// > - For vendor channel pushes, this is only effective for Huawei and Honor channels.
	//
	// >
	//
	// > - When pushing through Alibaba Cloud\\"s proprietary channel, this is only effective on Huawei, Honor, and vivo models.
	//
	// example:
	//
	// 4
	BadgeSetNum *int32 `json:"BadgeSetNum,omitempty" xml:"BadgeSetNum,omitempty"`
	// Sets the channelId for the Android app. It must correspond to the channelId in the vendor\\"s app.
	//
	// > - Because the channel_id for OPPO\\"s private message notification channel is the same as the app\\"s channelId, the channel_id takes this value when pushing through the OPPO channel.
	//
	// >
	//
	// > - For pushes through Huawei, FCM, and Alibaba Cloud\\"s proprietary channels, the channel_id takes this value.
	//
	// >
	//
	// > - For specific uses, see the FAQ: [Notifications not received on Android 8.0 and later devices](https://help.aliyun.com/document_detail/67398.html).
	//
	// example:
	//
	// 8.0up
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// Custom extension properties for Android notifications.
	//
	// > - The parameter must be passed in a standard JSON Map format. An incorrect format causes parsing to fail.
	//
	// example:
	//
	// {"key1":"value1"}
	ExtParameters *string `json:"ExtParameters,omitempty" xml:"ExtParameters,omitempty"`
	// Message grouping. For messages in the same group, only the latest one and the total number of messages received in that group are displayed in the notification bar. Not all messages are displayed, and they cannot be expanded. Currently supported by:
	//
	// - Huawei vendor channel
	//
	// - Honor vendor channel
	//
	// - Proprietary channels with Android SDK 3.9.1 and earlier
	//
	// > This parameter is no longer supported by proprietary channels in Android SDK 3.9.2 and later versions.
	//
	// example:
	//
	// group-1
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The URL for the icon on the right. Currently supported by:
	//
	// - `Huawei EMUI` (only applicable in long text mode and Inbox mode).
	//
	// - `Honor Magic UI` (only applicable in long text mode).
	//
	// - `Proprietary channels` (Android SDK 3.5.0 and later).
	//
	// example:
	//
	// https://imag.example.com/image.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The body text in Inbox mode. The content is a valid JSON Array with no more than 5 elements. Currently supported by:
	//
	// - Huawei: EMUI 9 and later
	//
	// - Honor: Magic UI 4.0 and later
	//
	// - Xiaomi: MIUI 10 and later
	//
	// - OPPO: ColorOS 5.0 and later
	//
	// - Proprietary channels: Android SDK 3.6.0 and later
	InboxContent []*string `json:"InboxContent,omitempty" xml:"InboxContent,omitempty" type:"Repeated"`
	// The Huawei vendor channel notification sound. Specify the name of the audio file stored in the `app/src/main/res/raw/` directory of the client project, without the file format suffix. If not set, the default ringtone is used.
	//
	// example:
	//
	// alicloud_notification_sound
	Music *string `json:"Music,omitempty" xml:"Music,omitempty"`
	// The unique identifier for an Android notification bar message. It controls the overwriting and replacement behavior of notifications. A new notification with the same NotifyId automatically overwrites the old one.
	//
	// example:
	//
	// 233856727
	NotifyId *int32 `json:"NotifyId,omitempty" xml:"NotifyId,omitempty"`
	// Detailed channel configuration.
	Options *PushTaskNotificationAndroidOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	// The image URL in large image mode. Currently supported by: proprietary channels with Android SDK 3.6.0 and later.
	//
	// example:
	//
	// https://imag.example.com/image.png
	PictureUrl *string `json:"PictureUrl,omitempty" xml:"PictureUrl,omitempty"`
	// The notification style. Valid values are:
	//
	// - `0`: Standard mode (default)
	//
	// - `1`: Long text mode (supported by Huawei, Honor, Xiaomi, OPPO, Meizu, and proprietary channels)
	//
	// - `2`: Large image mode (supported by proprietary channels)
	//
	// - `3`: List mode (supported by Huawei, Honor, Xiaomi, OPPO, and proprietary channels)
	//
	// example:
	//
	// 0
	RenderStyle *string `json:"RenderStyle,omitempty" xml:"RenderStyle,omitempty"`
	// Sets the vendor channel notification type:
	//
	// - `false`: Formal notification (default).
	//
	// - `true`: Test notification.
	//
	// > Currently supported by: Huawei channel, Honor channel, vivo channel, and OPPO Fluid Cloud.
	//
	// example:
	//
	// false
	TestMessage *bool `json:"TestMessage,omitempty" xml:"TestMessage,omitempty"`
	// Specifies the Activity to open after the notification is clicked.
	//
	// 	Warning:
	//
	// You must fill in this option when you use an Android vendor channel.
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
	// Alibaba Cloud proprietary configuration
	//
	// > This is only valid when using Alibaba Cloud\\"s proprietary channel.
	Accs *PushTaskNotificationAndroidOptionsAccs `json:"Accs,omitempty" xml:"Accs,omitempty" type:"Struct"`
	// Honor configuration
	Honor *PushTaskNotificationAndroidOptionsHonor `json:"Honor,omitempty" xml:"Honor,omitempty" type:"Struct"`
	// Huawei configuration
	Huawei *PushTaskNotificationAndroidOptionsHuawei `json:"Huawei,omitempty" xml:"Huawei,omitempty" type:"Struct"`
	// Meizu configuration
	Meizu *PushTaskNotificationAndroidOptionsMeizu `json:"Meizu,omitempty" xml:"Meizu,omitempty" type:"Struct"`
	// OPPO configuration
	Oppo *PushTaskNotificationAndroidOptionsOppo `json:"Oppo,omitempty" xml:"Oppo,omitempty" type:"Struct"`
	// vivo configuration
	Vivo *PushTaskNotificationAndroidOptionsVivo `json:"Vivo,omitempty" xml:"Vivo,omitempty" type:"Struct"`
	// Xiaomi configuration
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
	// The custom Android notification bar style. The value can be from 1 to 100.
	//
	// > The client must complete the style preset configuration. For more information, see the [Custom Notification Style API](https://help.aliyun.com/document_detail/2834944.html) document.
	//
	// example:
	//
	// 1
	CustomStyle *int32 `json:"CustomStyle,omitempty" xml:"CustomStyle,omitempty"`
	// The notification reminder method. Valid values:
	//
	// - `VIBRATE`: Vibrate (default)
	//
	// - `SOUND`: Sound
	//
	// - `BOTH`: Sound and vibration
	//
	// - `NONE`: Silent
	//
	// example:
	//
	// NONE
	NotifyType *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// Sets the activity to open when the notification is clicked. This is valid when `OpenType` is `ACTIVITY`.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	OpenActivity *string `json:"OpenActivity,omitempty" xml:"OpenActivity,omitempty"`
	// The action to take after the notification is clicked. Valid values:
	//
	// - `APPLICATION`: Open the application (default).
	//
	// - `ACTIVITY`: Open the specified page `OpenActivity`.
	//
	// - `URL`: Open a URL.
	//
	// - `NONE`: No action.
	//
	// example:
	//
	// APPLICATION
	OpenType *string `json:"OpenType,omitempty" xml:"OpenType,omitempty"`
	// After an Android device receives a push, clicking the notification opens the corresponding URL. This is valid when `OpenType` is `URL`.
	//
	// example:
	//
	// www.example.com
	OpenUrl *string `json:"OpenUrl,omitempty" xml:"OpenUrl,omitempty"`
	// The priority of the Android notification\\"s position in the notification bar. Valid values: -2, -1, 0, 1, 2.
	//
	// example:
	//
	// 0
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Message grouping. Messages in the same group are displayed collapsed in the notification bar and can be expanded. Different groups of notifications are displayed separately.
	//
	// > This is for Android SDK 3.9.2 and later.
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
	// Sets the importance parameter for Honor notification message classification, which determines the notification behavior on the user\\"s device. Valid values are:
	//
	// - `0`: Marketing message
	//
	// - `1`: Service and communication message
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
	// Sets the Huawei quick notification parameters.
	//
	// - **0**: Send a normal Huawei notification (default).
	//
	// - **1**: Send a Huawei quick notification.
	//
	// example:
	//
	// 1
	BusinessType *int32 `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// Function 1: After you apply for [self-classification rights](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835?#section3410731125514), this is used to identify the message type and determine the [message reminder method](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#ZH-CN_TOPIC_0000001149358835__p3850133955718). It speeds up the sending of specific types of messages. For valid values, see the [message classification standards](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section1076611477914) in the official Huawei Push documentation. Fill in the "Cloud notification category value" or "Local notification category value" from the document\\"s table.
	//
	// Function 2: After [applying for special permissions](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/faq-0000001050042183#section037425218509), this is used to identify high-priority pass-through scenarios. Valid values are:
	//
	// - `VOIP`: Video call
	//
	// - `PLAY_VOICE`: Voice playback
	//
	// > 	- For "Cloud notification category value" that is "Not applicable," all messages go through Alibaba Cloud\\"s proprietary channel.
	//
	// >
	//
	// > 	- For "Local notification category value" that is "Not applicable," all messages go through the Huawei channel.
	//
	// example:
	//
	// VOIP
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Sets the importance parameter for Huawei notification message classification, which determines the notification behavior on the user\\"s device. Valid values are:
	//
	// - `0`: Marketing message
	//
	// - `1`: Service and communication message
	//
	// > We recommend using `Category` for notification classification. You must apply for this on the Huawei platform. [Application link](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section893184112272).
	//
	// example:
	//
	// 0
	Importance *int32 `json:"Importance,omitempty" xml:"Importance,omitempty"`
	// The JSON string of the Huawei Android Live Window data structure [LiveNotificationPayload](https://developer.huawei.com/consumer/cn/doc/HMSCore-References/rest-live-0000001562939968#ZH-CN_TOPIC_0000001700850537__p195121620102511). For developer integration, see the document [Huawei Live Window Push Guide](https://help.aliyun.com/document_detail/2983768.html).
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
	// The receipt ID for the Huawei channel. This ID can be found in the receipt parameter settings on the Huawei channel push operations platform.
	//
	// > If the default receipt configuration on the Huawei channel push operations platform is the Alibaba Cloud receipt, you do not need to provide this. If not, we recommend that you first configure the default Huawei channel receipt ID in the Alibaba Cloud EMAS Mobile Push console.
	//
	// example:
	//
	// RCP4C123456
	ReceiptId *string `json:"ReceiptId,omitempty" xml:"ReceiptId,omitempty"`
	// The Huawei channel notification delivery priority. Valid values are:
	//
	// - `HIGH`
	//
	// - `NORMAL`
	//
	// You must apply for permission. For more information, see: [Application link](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/faq-0000001050042183#section037425218509).
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
	// - 0 Public message (default)
	//
	// - 1 Private message
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
	// OPPO classifies messages into two categories for management: communication and services, and content and marketing.
	//
	// **Communication and services (requires permission application):**
	//
	// - IM: Instant messages
	//
	// - ACCOUNT: Account and asset
	//
	// - TODO: To-do list
	//
	// - DEVICE_REMINDER: Device information
	//
	// - ORDER: Order and logistics
	//
	// - SUBSCRIPTION: Subscription reminder
	//
	// **Content and marketing:**
	//
	// - NEWS: News
	//
	// - CONTENT: Content recommendation
	//
	// - MARKETING: Operational activity
	//
	// - SOCIAL: Social dynamics
	//
	// For more information, see [vivo classification description](https://open.oppomobile.com/new/developmentDoc/info?id=13189).
	//
	// example:
	//
	// NEWS
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The JSON string of the OPPO Fluid Cloud\\"s intent deletion data structure [data](https://open.oppomobile.com/documentation/page/info?id=13578). This parameter is invalid if the AndroidOppoIntelligentIntent parameter is already filled. For developer integration, see the document [OPPO Fluid Cloud Push Guide](https://help.aliyun.com/document_detail/2997310.html).
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
	// The JSON string of the OPPO Fluid Cloud\\"s intent sharing data structure [IntelligentIntent](https://open.oppomobile.com/documentation/page/info?id=13565). For developer integration, see the document [OPPO Fluid Cloud Push Guide](https://help.aliyun.com/document_detail/2997310.html).
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
	// The OPPO channel notification bar message reminder level. Valid values are:
	//
	// - `1`: Notification bar
	//
	// - `2`: Notification bar, lock screen, ringtone, vibration (default notification level for communication and service messages)
	//
	// - `16`: Notification bar, lock screen, ringtone, vibration, banner (requires permission application)
	//
	// > When you use the `NotifyLevel` parameter, you must also pass the `Category` parameter.
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
	// {"name": "张三"}
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
	// vivo classifies messages into two categories for management: system messages and operational messages.
	//
	// **System messages:**
	//
	// - IM: Instant messages
	//
	// - ACCOUNT: Account and asset
	//
	// - TODO: To-do list
	//
	// - DEVICE_REMINDER: Device information
	//
	// - ORDER: Order and logistics
	//
	// - SUBSCRIPTION: Subscription reminder
	//
	// **Operational messages:**
	//
	// - NEWS: News
	//
	// - CONTENT: Content recommendation
	//
	// - MARKETING: Operational activity
	//
	// - SOCIAL: Social dynamics
	//
	// For more information, see [vivo classification description](https://dev.vivo.com.cn/documentCenter/doc/359#s-ef3qugc3).
	//
	// example:
	//
	// MARKETING
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Sets the vivo notification message classification. Valid values are:
	//
	// - `0`: Operational message (default)
	//
	// - `1`: System message
	//
	// > We recommend using `Category` for notification classification. You must apply for this on the vivo platform. For more information, see: [Application link](https://dev.vivo.com.cn/documentCenter/doc/359).
	//
	// example:
	//
	// 0
	Importance *int32 `json:"Importance,omitempty" xml:"Importance,omitempty"`
	// The JSON string of the vivo Atomic Island data structure [liveMessage](https://dev.vivo.com.cn/documentCenter/doc/896#s-fdagzbd4). For developer integration, see the document [vivo Atomic Island Push Guide](https://help.aliyun.com/zh/document_detail/3030718.html).
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
	// The message receipt identifier for the vivo vendor push channel. It is used to receive push result callback notifications.
	//
	// > - Location: vivo Open Platform → Push Service → Application Information → Receipt Configuration
	//
	// >
	//
	// > - Recommendation: First, configure the default receipt ID in the Alibaba Cloud EMAS console.
	//
	// >
	//
	// > - Condition: This must be configured only if the default receipt on the vivo platform is not the Alibaba Cloud receipt.
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
	// Sets the channelId for the Xiaomi notification type. You must apply for this on the Xiaomi platform. For more information, see: [Application link](https://dev.mi.com/console/doc/detail?pId=2422#_4).
	//
	// > A single application can apply for a maximum of 8 channels on the Xiaomi channel. Plan accordingly.
	//
	// example:
	//
	// michannel
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The JSON string of the Xiaomi Super Island data structure [miui.focus.param](https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131). For developer integration, see the document [Xiaomi Super Island Push Guide](https://help.aliyun.com/zh/document_detail/3037956.html).
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
	// The JSON string of the Xiaomi Super Island data image [miui.focus.pic_xxx](https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131). For developer integration, see the document [Xiaomi Super Island Push Guide](https://help.aliyun.com/zh/document_detail/3037956.html).
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
	FocusPics      *string `json:"FocusPics,omitempty" xml:"FocusPics,omitempty"`
	TemplateId     *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
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
	// Specifies the action corresponding to the ability of an in-app page.
	//
	// > For more information, see [ClickAction.action](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section152462191216) on the official HarmonyOS website.
	//
	// example:
	//
	// com.example.action
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The HarmonyOS application badge cumulative number.
	//
	// > - This is supported starting from HarmonyOS SDK 1.2.0.
	//
	// >
	//
	// > - See the description of the [addNum field](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section266310382145) for HarmonyOS badges.
	//
	// example:
	//
	// 1
	BadgeAddNum *int32 `json:"BadgeAddNum,omitempty" xml:"BadgeAddNum,omitempty"`
	// The HarmonyOS application badge number setting.
	//
	// > - See the description of the [setNum field](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section266310382145) for HarmonyOS badges.
	//
	// >
	//
	// > - This is supported starting from HarmonyOS SDK 1.2.0.
	//
	// example:
	//
	// 1
	BadgeSetNum *int32 `json:"BadgeSetNum,omitempty" xml:"BadgeSetNum,omitempty"`
	// The notification message category. This is an optional parameter. The default category is `MARKETING`.
	//
	// > After you apply for the right to self-classify notification messages, this parameter is used to identify the message type. Different notification message types affect how messages are displayed and how users are reminded. For more information, see [Notification.category](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117) on the official HarmonyOS website.
	//
	// example:
	//
	// IM
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Sets custom extension properties for the notification message. This is used to pass additional business data.
	//
	// > The parameter must be passed in a standard JSON Map format. An incorrect format causes parsing to fail.
	//
	// example:
	//
	// {"key": "value"}
	ExtParameters *string `json:"ExtParameters,omitempty" xml:"ExtParameters,omitempty"`
	// Extra data for the notification extension message.
	//
	// > - This is valid when sending a HarmonyOS notification extension message.
	//
	// >
	//
	// > - It is conceptually equivalent to the extraData field of a HarmonyOS notification extension message. For a specific definition, see the HarmonyOS [ExtensionPayload](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section161192514234) description.
	//
	// >
	//
	// > - This is supported starting from HarmonyOS SDK 1.2.0.
	//
	// example:
	//
	// text
	ExtensionExtraData *string `json:"ExtensionExtraData,omitempty" xml:"ExtensionExtraData,omitempty"`
	// Enables the HarmonyOS notification extension.
	//
	// > - You must first apply for permission on the official HarmonyOS website to send notification extension messages. For related content, see the [HarmonyOS documentation](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V5/push-send-extend-noti-V5) on sending notification extension messages.
	//
	// >
	//
	// > - This is supported starting from HarmonyOS SDK 1.2.0.
	//
	// example:
	//
	// false
	ExtensionPush *bool `json:"ExtensionPush,omitempty" xml:"ExtensionPush,omitempty"`
	// The URL for the large icon on the right side of the notification. The URL must use the HTTPS protocol.
	//
	// > - Supported image formats are png, jpg, jpeg, heif, gif, and bmp. The image dimensions (length × width) must be less than 25,000 pixels.
	//
	// >
	//
	// > - For more information, see [Notification.image](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117) on the official HarmonyOS website.
	//
	// example:
	//
	// https://example.com/xxx.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// When `RenderStyle` is `MULTI_LINE`, you must fill in this field to define the content for the multi-line text style. It supports up to 3 lines of content.
	InboxContent []*string `json:"InboxContent,omitempty" xml:"InboxContent,omitempty" type:"Repeated"`
	// The JSON string of the HarmonyOS Live Window data structure [LiveViewPayload](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V13/push-scenariozed-api-request-param-V13#section66881469306). For developer integration, see the document [HarmonyOS Live Window Push Guide](https://help.aliyun.com/document_detail/2982112.html).
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
	// Specifies the unique identifier (notifyId) for each message when it is displayed in the notification bar. If not provided, the push service automatically generates a unique identifier. Different notification messages can use the same notifyId to allow new messages to overwrite old ones. For more information, see [Notification.notifyId](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117) on the official HarmonyOS website.
	//
	// example:
	//
	// 123456
	NotifyId *int32 `json:"NotifyId,omitempty" xml:"NotifyId,omitempty"`
	// The receipt ID for the HarmonyOS channel. This ID can be found in the receipt parameter settings on the HarmonyOS channel push operations platform.
	//
	// > - If the default receipt configuration on the HarmonyOS channel push operations platform is the Alibaba Cloud receipt, you do not need to provide this. If not, we recommend that you first configure the default HarmonyOS channel receipt ID in the Alibaba Cloud EMAS Mobile Push console.
	//
	// >
	//
	// > - For more information, see [pushOptions.receiptId](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section418321011212) on the official HarmonyOS website.
	//
	// example:
	//
	// RCPB***DFD5
	ReceiptId *string `json:"ReceiptId,omitempty" xml:"ReceiptId,omitempty"`
	// The notification message style. This is an optional parameter. The default is a normal notification.
	//
	// example:
	//
	// NORMAL
	RenderStyle *string `json:"RenderStyle,omitempty" xml:"RenderStyle,omitempty"`
	// Uses the specified type of notification channel.
	//
	// > - This is only valid for Alibaba Cloud\\"s proprietary channels.
	//
	// >
	//
	// > - For more information, see [SlotType](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/js-apis-notificationmanager-V5#slottype) on the official HarmonyOS website.
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
	// The duration of the custom message notification ringtone in seconds. The range is [1, 60]. If the ringtone duration is too short, it will loop.
	//
	// example:
	//
	// 2
	SoundDuration *int32 `json:"SoundDuration,omitempty" xml:"SoundDuration,omitempty"`
	// Enables test messages.
	//
	// > - For more information, see the HarmonyOS push parameter [TestMessage](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section418321011212).
	//
	// example:
	//
	// true
	TestMessage *bool `json:"TestMessage,omitempty" xml:"TestMessage,omitempty"`
	// The URI corresponding to the ability of an in-app page.
	//
	// > - If there are multiple abilities, specify the action and URI for each ability separately. The system prioritizes using the action to find the corresponding in-app page.
	//
	// >
	//
	// > - For more information, see [ClickAction.uri](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section152462191216) on the official HarmonyOS website.
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
	// iOS notifications are sent through the Apple Push Notification service (APNs) center. You must specify the environment information. This is an optional parameter. The default is the production environment.
	//
	// - DEV: Development environment, for applications installed and tested directly from Xcode.
	//
	// - PRODUCT: Production environment, for applications distributed through the App Store, TestFlight, Ad Hoc, and enterprise channels.
	//
	// example:
	//
	// DEV
	ApnsEnv *string `json:"ApnsEnv,omitempty" xml:"ApnsEnv,omitempty"`
	// The iOS application badge.
	//
	// example:
	//
	// 1
	Badge *int32 `json:"Badge,omitempty" xml:"Badge,omitempty"`
	// Specifies whether to enable the badge auto-increment feature. This is an optional parameter. The default value is false.
	//
	// > - This parameter cannot be used with the badge setting parameter.
	//
	// >
	//
	// > - The badge auto-increment feature is maintained by the Alibaba Cloud push server, which counts the badges for each device. You must use SDK version 1.9.5 or later and actively sync the badge number to the server through the SDK.
	//
	// example:
	//
	// false
	BadgeAutoIncrement *bool `json:"BadgeAutoIncrement,omitempty" xml:"BadgeAutoIncrement,omitempty"`
	// Specifies the category identifier for an iOS notification. This defines the notification\\"s interactive behavior and display style.
	//
	// > - The category must be pre-registered in the app to take effect.
	//
	// >
	//
	// > - Different categories can define different sets of actions.
	//
	// example:
	//
	// MESSAGE_REPLY
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// A unique identifier that controls notification merging. Notifications with the same identifier are overwritten.
	//
	// example:
	//
	// order_status_update_12345
	CollapseId *string `json:"CollapseId,omitempty" xml:"CollapseId,omitempty"`
	// Custom extension properties for iOS notifications.
	//
	// > - The parameter must be passed in a standard JSON Map format. An incorrect format causes parsing to fail.
	//
	// example:
	//
	// {"attachment": "https://xxxx.xxx/notification_pic.png"}
	ExtParameters *string `json:"ExtParameters,omitempty" xml:"ExtParameters,omitempty"`
	// The interruption level. This is an optional parameter. Valid values are:
	//
	// - `passive`: The system adds the notification to the notification list without lighting up the screen or playing a sound.
	//
	// - `active`: The system displays the notification immediately, lights up the screen, and can play a sound.
	//
	// - `time-sensitive`: The system presents the notification immediately, lights up the screen, and can play a sound, but does not override system notification controls.
	//
	// - `critical`: The system displays the notification immediately, lights up the screen, and plays a sound, bypassing the mute switch.
	//
	// example:
	//
	// active
	InterruptionLevel *string `json:"InterruptionLevel,omitempty" xml:"InterruptionLevel,omitempty"`
	// Live Activities parameter object.
	//
	// 	Notice:
	//
	// - Live Activities push only supports pushing to a single device of the `DEVICE` type.
	//
	// - When you push to Live Activities, you can leave the title and body parameters empty.
	LiveActivity *PushTaskNotificationIosLiveActivity `json:"LiveActivity,omitempty" xml:"LiveActivity,omitempty" type:"Struct"`
	// The iOS notification sound. Specify the name of the audio file stored in the app bundle or the sandbox Library/Sounds directory. For more information, see [How to set the notification sound for iOS push](https://help.aliyun.com/document_detail/48906.html).
	//
	// > - If you specify an empty string (""), the notification is silent.
	//
	// >
	//
	// > - If this parameter is not set, the default value is \\`default\\`, which is the system prompt sound.
	//
	// example:
	//
	// default
	Music *string `json:"Music,omitempty" xml:"Music,omitempty"`
	// Enables extended notifications and controls whether iOS notifications support processing by the Notification Service Extension.
	//
	// > - This must be set to true when you send a silent notification.
	//
	// >
	//
	// > - The extension processing time cannot exceed 30 seconds.
	//
	// >
	//
	// > - A timeout causes the notification to display the original content.
	//
	// >
	//
	// > - You must add a Notification Service Extension to your application.
	//
	// example:
	//
	// true
	Mutable *bool `json:"Mutable,omitempty" xml:"Mutable,omitempty"`
	// The relevance score of the notification message. It is used to control the priority and display policy of the notification.
	//
	// example:
	//
	// 0.5
	RelevanceScore *float64 `json:"RelevanceScore,omitempty" xml:"RelevanceScore,omitempty"`
	// Controls whether to enable silent push mode.
	//
	// > - When you send a silent notification, you can leave the `title` and `body` parameters empty.
	//
	// example:
	//
	// false
	Silent *bool `json:"Silent,omitempty" xml:"Silent,omitempty"`
	// The subtitle of the iOS notification.
	//
	// example:
	//
	// 请查收订单。
	Subtitle *string `json:"Subtitle,omitempty" xml:"Subtitle,omitempty"`
	// The thread identifier for iOS notification grouping. It is used to classify and collapse related notifications.
	//
	// > - Notifications with the same thread-id are automatically grouped.
	//
	// >
	//
	// > - Multiple related notifications are collapsed into one notification group.
	//
	// >
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
	// Static pass-through parameters for iOS Live Activities push. They are used to transmit immutable business identification information.
	//
	// > This is required when `Event` is \\`start\\`.
	//
	// example:
	//
	// {
	//
	//   "orderId": "ORD20231201001",
	//
	//   "restaurantName": "美味餐厅",
	//
	//   "customerAddress": "xx区xx路xx号",
	//
	//   "orderType": "delivery"
	//
	// }
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The type of Live Activity to start.
	//
	// > This is required when `Event` is \\`start\\`.
	//
	// example:
	//
	// OrderActivityAttributes
	AttributesType *string `json:"AttributesType,omitempty" xml:"AttributesType,omitempty"`
	// Dynamic pass-through parameters for a Live Activity. They contain real-time updatable status information and changing data.
	//
	// > - Avoid overly frequent updates. An interval of 5 seconds or more is recommended.
	//
	// >
	//
	// > - Update multiple fields in a batch to reduce the number of pushes.
	//
	// >
	//
	// > - Consider the user experience and avoid screen flickering.
	//
	// >
	//
	// > - Must be a valid JSON string.
	//
	// example:
	//
	// {
	//
	//     "status": "delivering",
	//
	//     "estimatedTime": "10分钟",
	//
	//     "progress": 80,
	//
	//     "driverName": "李师傅",
	//
	//     "currentStep": "配送员正在路上"}
	//
	// }
	ContentState *string `json:"ContentState,omitempty" xml:"ContentState,omitempty"`
	// Sets the retention period for a finished Live Activity on the lock screen. This lets users view information after the activity has ended. It is a Unix timestamp in seconds.
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
	// The unique identifier for a Live Activity. It associates the activity instance on the device with the push target on the server.
	//
	// 	Notice:
	//
	// - This `ID` must be the same as the `forActivityId` parameter of the `registerLiveActivityPushToken` method in the client SDK.
	//
	// - The server uses this `ID` to locate the specific activity instance during a push.
	//
	// example:
	//
	// FOOD_DELIVERY_ORD20231201001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Sets the expiration timestamp for the content of an iOS Live Activity. It is a Unix timestamp in seconds.
	//
	// > - After the specified time is reached, the system automatically marks the activity as expired.
	//
	// >
	//
	// > - Expired activities are removed from the Live Activity and the lock screen.
	//
	// >
	//
	// > - This prevents outdated information from occupying the user interface for a long time.
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
	// Sets the expiration time of the message. After this time, the message will no longer be sent. The maximum retention period is 72 hours.
	//
	// > - This uses the ISO 8601 standard and UTC time. The format is YYYY-MM-DDThh:mm:ssZ.
	//
	// >
	//
	// > - The expiration time must satisfy: ExpireTime > PushTime + 3 seconds (3 seconds is a buffer for network and system delays).
	//
	// >
	//
	// > - Recommendation: The expiration time for a single push should be at least 1 minute. For a push to all or a batch push, it should be at least 10 minutes.
	//
	// 	Notice:
	//
	// For pass-through messages, if you do not set an expiration time, the message is only sent to online devices. If the device is offline, the message is discarded.
	//
	// example:
	//
	// 2025-06-21T12:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// A custom identifier for the push task. If JobKey is not empty, this field will be included in the receipt log. To view receipt logs, see [Receipt logs](https://help.aliyun.com/document_detail/434651.html).
	//
	// example:
	//
	// jobkey1727749697913
	JobKey *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	// A unique ID used to identify the message. This is only valid when the `Action` parameter is `CONTINUOUS_PUSH`.
	//
	// example:
	//
	// 1174754033128****
	MessageId *int64 `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// Specifies the sending time of the message, up to 7 days in the future. This is only valid when the `Action` parameter is `SCHEDULED_PUSH`.
	//
	// > This uses the ISO 8601 standard and UTC time. The format is yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2025-06-19T12:00:00Z
	PushTime *string `json:"PushTime,omitempty" xml:"PushTime,omitempty"`
	// Resends the message as a text message.
	//
	// > Currently, this is only supported for `Android` and `HarmonyOS` devices.
	Sms *PushTaskOptionsSms `json:"Sms,omitempty" xml:"Sms,omitempty" type:"Struct"`
	// Specifies whether to automatically truncate oversized titles and content.
	//
	// > This is only supported for vendor channels that have explicit limits on title and content length. It does not apply to channels like APNs, Huawei, and Honor, which do not limit title and content length but only the total request body size.
	//
	// example:
	//
	// false
	Trim *bool `json:"Trim,omitempty" xml:"Trim,omitempty"`
	// Specifies the sending channel. Valid values are:
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
	// > 	- If this parameter is not configured, all channels can be used.
	//
	// >
	//
	// > 	- If this parameter is configured, only the channels specified in the parameter are used.
	//
	// >
	//
	// > 	- If the configured channel conflicts with the sending policy (for example, iOS notifications only go through the APNs channel, but this parameter does not include \\`apns\\`), the message is not sent.
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
	// The delay time to trigger the text message, in seconds.
	//
	// This must be set if you use SMS filter interaction. We recommend setting it to 15 seconds or more, with a maximum of 3 days, to avoid duplicate text messages and pushes.
	//
	// > When you use SMS filter interaction, the ExpireTime parameter is invalid. The notification expiration time is calculated based on the DelaySecs parameter. The expiration time is the current time plus the DelaySecs time.
	//
	// example:
	//
	// 150
	DelaySecs *int64 `json:"DelaySecs,omitempty" xml:"DelaySecs,omitempty"`
	// Key-value pairs for the variables in the SMS template.
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
	// 某某科技
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The SMS template name. You can get this from the SMS template management interface. It is the name assigned by the system, not the name set by the developer.
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
	// The platform type. This is an optional parameter.
	//
	// example:
	//
	// IOS
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The push target type.
	//
	// 	Notice:
	//
	// The `MassPushV2` batch push API and `CONTINUOUS_PUSH` continuous push support only the following three target types:
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
	// Set the push target based on `Target.Type`. Separate multiple targets with commas. The target types and their values are described as follows:
	//
	// > - `DEVICE`: Device ID, such as deviceid1,deviceid2. You can specify up to 1,000 device IDs.
	//
	// >
	//
	// > - `ACCOUNT`: Account ID, such as account1,account2. You can specify up to 1,000 account IDs.
	//
	// >
	//
	// > - `ALIAS`: Alias, such as alias1,alias2. You can specify up to 1,000 aliases.
	//
	// >
	//
	// > - `TAG`: Supports one or more tags. For more information about the format, see [Tag format specifications](https://help.aliyun.com/document_detail/434847.html).
	//
	// >
	//
	// > - `ALL`: Push to all devices. You do not need to set a value. Pushing to all devices may increase costs. Use this feature with caution.
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
