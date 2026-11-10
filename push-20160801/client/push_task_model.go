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
	// The push method. This parameter is optional. Default value: PUSH_IMMEDIATELY (push immediately).
	//
	// 	Notice:
	//
	// The MassPushV2 batch push API supports only the following push methods:
	//
	// - PUSH_IMMEDIATELY (push immediately)
	//
	// - SCHEDULED_PUSH (scheduled push)
	//
	// example:
	//
	// PUSH_IMMEDIATELY
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The pass-through message data sent to the device. The total length cannot exceed 4,000 bytes.
	//
	// > Length calculation rules
	//
	// > - The length is calculated based on the byte length of the UTF-8 encoded string after the Message object is serialized into JSON.
	//
	// > - A Chinese character typically occupies 3 bytes in UTF-8 encoding.
	Message *PushTaskMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Struct"`
	// The vendor notification data sent to the device.
	//
	// 	Notice:
	//
	// If both Message and Notification are configured, the device receives only one of them based on the following rules:
	//
	// - If the device is online, the pass-through message data is sent.
	//
	// - If the device is offline, the system notification is sent.
	Notification *PushTaskNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	// The push options.
	Options *PushTaskOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	// The target object for the message push. This parameter is optional when the Action parameter is set to CREATE_CONTINUOUS_PUSH (create a continuous push task).
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
	// The body of the message to be sent.
	//
	// example:
	//
	// {"key": "value"}
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The title of the message to be sent.
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
	// The body of the notification to be pushed.
	//
	// > Length limits:
	//
	// > - iOS/HarmonyOS/Android: the **character length*	- cannot exceed 200.
	//
	// example:
	//
	// Dear customer, your reservation has been successfully canceled
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// HarmonyOS notification configuration.
	Hmos *PushTaskNotificationHmos `json:"Hmos,omitempty" xml:"Hmos,omitempty" type:"Struct"`
	// iOS notification configuration.
	Ios *PushTaskNotificationIos `json:"Ios,omitempty" xml:"Ios,omitempty" type:"Struct"`
	// The title of the notification to be pushed.
	//
	// > Length limits:
	//
	// > - iOS/HarmonyOS: the **byte length*	- cannot exceed 200.
	//
	// > - Android: the **character length*	- cannot exceed 50.
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
	// The fully qualified class name of the application entry activity for badge settings.
	//
	// > Valid only for Huawei and Honor vendor channel pushes.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	BadgeActivity *string `json:"BadgeActivity,omitempty" xml:"BadgeActivity,omitempty"`
	// Sets the increment value for the badge, which is added to the original badge number.
	//
	// > - Supported by Huawei and Honor channels.
	//
	// > - If both BadgeAddNum and BadgeSetNum are specified, BadgeSetNum takes precedence.
	//
	// example:
	//
	// 1
	BadgeAddNum *int32 `json:"BadgeAddNum,omitempty" xml:"BadgeAddNum,omitempty"`
	// Sets a fixed value for the badge number. Valid values: 1 to 99.
	//
	// > - This parameter takes effect only for Huawei and Honor vendor channel pushes.
	//
	// > - For Alibaba Cloud proprietary channel pushes, this parameter takes effect only on Huawei, Honor, and vivo devices.
	//
	// example:
	//
	// 4
	BadgeSetNum *int32 `json:"BadgeSetNum,omitempty" xml:"BadgeSetNum,omitempty"`
	// Sets the channel ID of the Android application. The channel ID must match the channel ID in the vendor application.
	//
	// > - Because the channel_id of the OPPO private message channel is the same as the channel ID of the application, this value is used as the channel_id when pushing through the OPPO channel.
	//
	// > - This value is also used as the channel_id for Huawei, FCM, and Alibaba Cloud proprietary channel pushes.
	//
	// > - For specific use cases, see the FAQ: [Notifications are not received on devices running Android 8.0 or later](https://help.aliyun.com/document_detail/67398.html).
	//
	// example:
	//
	// 8.0up
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The custom extension attributes for Android notifications.
	//
	// > - The parameter must be passed in standard JSON map format. Incorrect formats will cause parsing failures.
	//
	// example:
	//
	// {"key1":"value1"}
	ExtParameters *string `json:"ExtParameters,omitempty" xml:"ExtParameters,omitempty"`
	// The message group. For messages in the same group, only the latest message and the total number of received messages in the group are displayed in the notification bar. All messages are not displayed and cannot be expanded. Currently supported by:
	//
	// - Huawei vendor channel.
	//
	// - Honor vendor channel.
	//
	// - Proprietary channels with Android SDK 3.9.1 or earlier.
	//
	// > Proprietary channels no longer support this parameter in Android SDK 3.9.2 or later.
	//
	// example:
	//
	// group-1
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The URL of the icon on the right side. Currently supported by:
	//
	// - Huawei EMUI (applicable only in long text mode and Inbox mode).
	//
	// - Honor Magic UI (applicable only in long text mode).
	//
	// - Proprietary channels (Android SDK 3.5.0 or later).
	//
	// example:
	//
	// https://imag.example.com/image.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The body text in Inbox mode. The content must be a valid JSON array with a maximum of five elements. Currently supported by:
	//
	// - Huawei: EMUI 9 or later.
	//
	// - Honor: Magic UI 4.0 or later.
	//
	// - Xiaomi: MIUI 10 or later.
	//
	// - OPPO: ColorOS 5.0 or later.
	//
	// - Proprietary channels: Android SDK 3.6.0 or later.
	InboxContent []*string `json:"InboxContent,omitempty" xml:"InboxContent,omitempty" type:"Repeated"`
	// The notification sound for the Huawei vendor channel. Specify the audio file name stored in the `app/src/main/res/raw/` directory of the client project. Do not include the file format extension. If not set, the default ringtone is used.
	//
	// example:
	//
	// alicloud_notification_sound
	Music *string `json:"Music,omitempty" xml:"Music,omitempty"`
	// The unique identifier of the Android notification bar message, used to control the override and replacement behavior of notifications. A new notification with the same NotifyId automatically overrides the old notification.
	//
	// example:
	//
	// 233856727
	NotifyId *int32 `json:"NotifyId,omitempty" xml:"NotifyId,omitempty"`
	// The detailed channel configuration.
	Options *PushTaskNotificationAndroidOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	// The image URL in large image mode. Currently supported by proprietary channels with Android SDK 3.6.0 or later.
	//
	// example:
	//
	// https://imag.example.com/image.png
	PictureUrl *string `json:"PictureUrl,omitempty" xml:"PictureUrl,omitempty"`
	// The notification style. Valid values:
	//
	// - 0: Standard mode (default).
	//
	// - 1: Long text mode (supported by Huawei, Honor, Xiaomi, OPPO, Meizu, and proprietary channels).
	//
	// - 2: Large image mode (supported by proprietary channels).
	//
	// - 3: List mode (supported by Huawei, Honor, Xiaomi, OPPO, and proprietary channels).
	//
	// example:
	//
	// 0
	RenderStyle *string `json:"RenderStyle,omitempty" xml:"RenderStyle,omitempty"`
	// Sets the notification type for vendor channels. Valid values:
	//
	// - false: Formal notification (default).
	//
	// - true: Test notification.
	//
	// > Currently supported by Huawei channel, Honor channel, vivo channel, and OPPO Fluid Cloud.
	//
	// example:
	//
	// false
	TestMessage *bool `json:"TestMessage,omitempty" xml:"TestMessage,omitempty"`
	// Specifies the activity to open when the notification is clicked.
	//
	// 	Warning: This parameter is required when you use Android vendor channels.</warning>
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
	// The Alibaba Cloud proprietary configuration.
	//
	// > Valid only when the Alibaba Cloud proprietary channel is used.
	Accs *PushTaskNotificationAndroidOptionsAccs `json:"Accs,omitempty" xml:"Accs,omitempty" type:"Struct"`
	// The Honor configuration.
	Honor *PushTaskNotificationAndroidOptionsHonor `json:"Honor,omitempty" xml:"Honor,omitempty" type:"Struct"`
	// The Huawei configuration.
	Huawei *PushTaskNotificationAndroidOptionsHuawei `json:"Huawei,omitempty" xml:"Huawei,omitempty" type:"Struct"`
	// The Meizu configuration.
	Meizu *PushTaskNotificationAndroidOptionsMeizu `json:"Meizu,omitempty" xml:"Meizu,omitempty" type:"Struct"`
	// The OPPO configuration.
	Oppo *PushTaskNotificationAndroidOptionsOppo `json:"Oppo,omitempty" xml:"Oppo,omitempty" type:"Struct"`
	// The vivo configuration.
	Vivo *PushTaskNotificationAndroidOptionsVivo `json:"Vivo,omitempty" xml:"Vivo,omitempty" type:"Struct"`
	// The Xiaomi configuration.
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
	// > The client must complete the style preset configuration. References: [Custom notification style API](https://help.aliyun.com/document_detail/2834944.html).
	//
	// example:
	//
	// 1
	CustomStyle *int32 `json:"CustomStyle,omitempty" xml:"CustomStyle,omitempty"`
	// The reminder method for the notification. Valid values:
	//
	// - VIBRATE: Vibration (default).
	//
	// - SOUND: Sound.
	//
	// - BOTH: Sound and vibration.
	//
	// - NONE: Silent.
	//
	// example:
	//
	// NONE
	NotifyType *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// Specifies the activity to open when the notification is clicked. Valid only when OpenType is set to ACTIVITY.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	OpenActivity *string `json:"OpenActivity,omitempty" xml:"OpenActivity,omitempty"`
	// The action to perform when the notification is clicked. Valid values:
	//
	// - APPLICATION: Open the application (default).
	//
	// - ACTIVITY: Open the specified activity OpenActivity.
	//
	// - URL: Open a URL.
	//
	// - NONE: No redirection.
	//
	// example:
	//
	// APPLICATION
	OpenType *string `json:"OpenType,omitempty" xml:"OpenType,omitempty"`
	// The URL to open when the notification is clicked after an Android device receives the push message. Valid only when OpenType is set to URL.
	//
	// example:
	//
	// www.example.com
	OpenUrl *string `json:"OpenUrl,omitempty" xml:"OpenUrl,omitempty"`
	// The priority of the display position for Android notifications in the notification bar. Valid values: -2, -1, 0, 1, and 2.
	//
	// example:
	//
	// 0
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The message group. Messages in the same group are collapsed and can be expanded in the notification bar. Notifications from different groups are displayed separately.
	//
	// > Supported by Android SDK 3.9.2 or later.
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
	// Sets the importance parameter for Honor notification message classification, which determines the notification behavior on user devices. Valid values:
	//
	// - 0: Information and marketing messages.
	//
	// - 1: Service and communication messages.
	//
	// You must apply on the Honor platform. See [Application link](https://developer.honor.com/cn/docs/11002/guides/notification-class#%E8%87%AA%E5%88%86%E7%B1%BB%E6%9D%83%E7%9B%8A%E7%94%B3%E8%AF%B7).
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
	// Sets the Huawei fast notification parameter. Valid values:
	//
	// - **0**: Send standard Huawei notifications (default).
	//
	// - **1**: Send Huawei fast notifications.
	//
	// example:
	//
	// 1
	BusinessType *int32 `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// Purpose 1: After applying for [self-classification benefits](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835?#section3410731125514), this parameter is used to identify the message type, determine the [message reminder method](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#ZH-CN_TOPIC_0000001149358835__p3850133955718), and accelerate the delivery of specific types of messages. For valid values, refer to the [message classification standards](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section1076611477914) in the official Huawei push documentation, and fill in the "Cloud notification category value" or "Local notification category value" from the documentation table.
	//
	// Purpose 2: After [applying for special permissions](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/faq-0000001050042183#section037425218509), this parameter is used to identify high-priority pass-through scenarios. Valid values:
	//
	// - VOIP: Audio and video calls.
	//
	// - PLAY_VOICE: Voice broadcast.
	//
	// > - If the "Cloud notification category value" is "Not applicable", the message is sent through the Alibaba Cloud proprietary channel.
	//
	// > - If the "Local notification category value" is "Not applicable", the message is sent through the Huawei channel.
	//
	// example:
	//
	// VOIP
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Sets the importance parameter for Huawei notification message classification, which determines the notification behavior on user devices. Valid values:
	//
	// - 0: Information and marketing messages.
	//
	// - 1: Service and communication messages.
	//
	// > We recommend that you use Category for notification classification. You must apply on the Huawei platform. See [Application link](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section893184112272).
	//
	// example:
	//
	// 0
	Importance *int32 `json:"Importance,omitempty" xml:"Importance,omitempty"`
	// The JSON string of the Huawei Android Live Window data structure [LiveNotificationPayload](https://developer.huawei.com/consumer/cn/doc/HMSCore-References/rest-live-0000001562939968#ZH-CN_TOPIC_0000001700850537__p195121620102511). For development and integration, refer to the [Huawei Live Window push guide](https://help.aliyun.com/document_detail/2983768.html).
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
	// The receipt ID of the Huawei channel. You can view this receipt ID in the receipt parameter settings of the Huawei channel push operations platform.
	//
	// > If the default receipt configuration on the Huawei channel push operations platform is set to Alibaba Cloud receipt, you do not need to provide this parameter. Otherwise, we recommend that you configure the default receipt ID for the Huawei channel in the Alibaba Cloud EMAS shift push console first.
	//
	// example:
	//
	// RCP4C123456
	ReceiptId *string `json:"ReceiptId,omitempty" xml:"ReceiptId,omitempty"`
	// The delivery priority for Huawei channel notifications. Valid values:
	//
	// - HIGH
	//
	// - NORMAL
	//
	// You must apply for permissions. For more information, see [Application link](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/faq-0000001050042183#section037425218509).
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
	// The Meizu message type. Valid values:
	//
	// - 0: Public message (default).
	//
	// - 1: Private message.
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
	// OPPO message categorization includes two categories for management: communication and service, and content and marketing.
	//
	// **Communication and service (requires permission application):**
	//
	// - IM: Instant messages.
	//
	// - ACCOUNT: Accounts and assets.
	//
	// - TODO: Schedules and to-do lists.
	//
	// - DEVICE_REMINDER: Device information.
	//
	// - ORDER: Orders and logistics.
	//
	// - SUBSCRIPTION: subscribe reminders.
	//
	// **Content and marketing:**
	//
	// - NEWS: News.
	//
	// - CONTENT: Content recommendations.
	//
	// - MARKETING: operational activity.
	//
	// - SOCIAL: Social updates.
	//
	// For more information, see the [OPPO categorization description](https://open.oppomobile.com/new/developmentDoc/info?id=13189).
	//
	// example:
	//
	// NEWS
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The JSON character string of the OPPO Fluid Cloud intent delete data structure [data](https://open.oppomobile.com/documentation/page/info?id=13578). This parameter is invalid if the AndroidOppoIntelligentIntent parameter is already specified. For development and integration, References: [OPPO Fluid Cloud push guide](https://help.aliyun.com/document_detail/2997310.html).
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
	// The JSON character string of the OPPO Fluid Cloud intent sharing data structure [IntelligentIntent](https://open.oppomobile.com/documentation/page/info?id=13565). For development and integration, References: [OPPO Fluid Cloud push guide](https://help.aliyun.com/document_detail/2997310.html).
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
	// The reminder level for OPPO channel notification bar messages. Valid values:
	//
	// - 1: Notification bar.
	//
	// - 2: Notification bar, lock screen, ringtone, and vibration (default notification level for communication and service messages).
	//
	// - 16: Notification bar, lock screen, ringtone, vibration, and banner (requires permission application).
	//
	// > When you use the NotifyLevel parameter, you must also pass the Category parameter.
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
	// {"name": "Zhang San"}
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
	// Sets the increment value for the vivo channel badge.
	//
	// example:
	//
	// true
	AddBadge *bool `json:"AddBadge,omitempty" xml:"AddBadge,omitempty"`
	// vivo message categorization includes two categories for management: system messages and operation messages.
	//
	// **System messages:**
	//
	// - IM: Instant messages.
	//
	// - ACCOUNT: Accounts and assets.
	//
	// - TODO: Schedules and to-do lists.
	//
	// - DEVICE_REMINDER: Device information.
	//
	// - ORDER: Orders and logistics.
	//
	// - SUBSCRIPTION: subscribe reminders.
	//
	// **Operation messages:**
	//
	// - NEWS: News.
	//
	// - CONTENT: Content recommendations.
	//
	// - MARKETING: operational activity.
	//
	// - SOCIAL: Social updates.
	//
	// For more information, see [vivo categorization description](https://dev.vivo.com.cn/documentCenter/doc/359#s-ef3qugc3).
	//
	// example:
	//
	// MARKETING
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Sets the vivo notification message classification. Valid values:
	//
	// - 0: Operation messages (default).
	//
	// - 1: System messages.
	//
	// > We recommend that you use Category for notification classification. You must apply on the vivo platform. For more information, see [Application link](https://dev.vivo.com.cn/documentCenter/doc/359).
	//
	// example:
	//
	// 0
	Importance *int32 `json:"Importance,omitempty" xml:"Importance,omitempty"`
	// The JSON character string of the vivo Atomic Island data structure [liveMessage](https://dev.vivo.com.cn/documentCenter/doc/896#s-fdagzbd4). For development and integration, References: [vivo Atomic Island push guide](https://www.alibabacloud.com/help/en/document_detail/3030718.html).
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
	// The message receipt identifier for the vivo vendor push channel, used to receive callback notifications for push results.
	//
	// > - Location: vivo Open Platform > Push Service > Application Information > Receipt Configuration.
	//
	// > - Configuration recommendation: Configure the default receipt ID in the Alibaba Cloud EMAS console first.
	//
	// > - Usage condition: Required only when the default receipt on the vivo platform is not the Alibaba Cloud receipt.
	//
	// example:
	//
	// 1232221
	ReceiptId *string `json:"ReceiptId,omitempty" xml:"ReceiptId,omitempty"`
	// example:
	//
	// t1-ge3-lq
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// {"keywords1":"S123"}
	TemplateParams *string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
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

func (s *PushTaskNotificationAndroidOptionsVivo) GetTemplateId() *string {
	return s.TemplateId
}

func (s *PushTaskNotificationAndroidOptionsVivo) GetTemplateParams() *string {
	return s.TemplateParams
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

func (s *PushTaskNotificationAndroidOptionsVivo) SetTemplateId(v string) *PushTaskNotificationAndroidOptionsVivo {
	s.TemplateId = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsVivo) SetTemplateParams(v string) *PushTaskNotificationAndroidOptionsVivo {
	s.TemplateParams = &v
	return s
}

func (s *PushTaskNotificationAndroidOptionsVivo) Validate() error {
	return dara.Validate(s)
}

type PushTaskNotificationAndroidOptionsXiaomi struct {
	// Sets the channel ID for the Xiaomi notification type. You must apply on the Xiaomi platform. For more information, see [Application link](https://dev.mi.com/console/doc/detail?pId=2422#_4).
	//
	// > A single application can apply for a maximum of eight channels on the Xiaomi channel. Plan accordingly in advance.
	//
	// example:
	//
	// michannel
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The JSON character string of the Xiaomi Super Island data structure [miui.focus.param](https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131). For development and integration, References: [Xiaomi Super Island push guide](https://www.alibabacloud.com/help/en/document_detail/3037956.html).
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
	// The JSON character string of the Xiaomi Super Island data images [miui.focus.pic_xxx](https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131). For development and integration, References: [Xiaomi Super Island push guide](https://www.alibabacloud.com/help/en/document_detail/3037956.html).
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
	// The Xiaomi private message template parameters, as a JSON string.
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
	// The action corresponding to the ability of the in-app page.
	//
	// > For more information, see [ClickAction.action](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section152462191216) on the HarmonyOS official website.
	//
	// example:
	//
	// com.example.action
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The number to add to the HarmonyOS application badge.
	//
	// > - Supported by HarmonyOS SDK 1.2.0 and later.
	//
	// > - For more information, see the [addNum field description](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section266310382145) for HarmonyOS badges.
	//
	// example:
	//
	// 1
	BadgeAddNum *int32 `json:"BadgeAddNum,omitempty" xml:"BadgeAddNum,omitempty"`
	// The number to set for the HarmonyOS application badge.
	//
	// > - For more information, see the [setNum field](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section266310382145) description for HarmonyOS badges.
	//
	// > - Supported by HarmonyOS SDK 1.2.0 and later.
	//
	// example:
	//
	// 1
	BadgeSetNum *int32 `json:"BadgeSetNum,omitempty" xml:"BadgeSetNum,omitempty"`
	// The category of the notification message. This parameter is optional. Default value: MARKETING.
	//
	// > After you apply for the notification message auto-classification permission, this parameter is used to identify the message type. Different notification message types affect the message display and reminder methods. For more information, see [Notification.category](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117) on the HarmonyOS official website.
	//
	// example:
	//
	// IM
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Sets the custom extension attributes of the notification message to pass additional business data.
	//
	// > The parameter must be passed in standard JSON map format. Incorrect formats will cause parsing failures.
	//
	// example:
	//
	// {"key": "value"}
	ExtParameters *string `json:"ExtParameters,omitempty" xml:"ExtParameters,omitempty"`
	// The additional data for notification extension messages.
	//
	// > - Valid only when sending HarmonyOS notification extension messages.
	//
	// > - Conceptually equivalent to the extraData field of HarmonyOS notification extension messages. For specific definitions, see the [ExtensionPayload](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section161192514234) description in HarmonyOS.
	//
	// > - Supported by HarmonyOS SDK 1.2.0 and later.
	//
	// example:
	//
	// text
	ExtensionExtraData *string `json:"ExtensionExtraData,omitempty" xml:"ExtensionExtraData,omitempty"`
	// Enables HarmonyOS notification extensions.
	//
	// > - To send notification extension messages, you must first apply for permissions on the HarmonyOS official website. For more information, see [HarmonyOS documentation](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V5/push-send-extend-noti-V5).
	//
	// > - Supported by HarmonyOS SDK 1.2.0 and later.
	//
	// example:
	//
	// false
	ExtensionPush *bool `json:"ExtensionPush,omitempty" xml:"ExtensionPush,omitempty"`
	// The URL of the large icon on the right side of the notification. The URL must use the HTTPS protocol.
	//
	// > - Supported image formats: PNG, JPG, JPEG, HEIF, GIF, and BMP. The total number of pixels (width × height) must be less than 25,000.
	//
	// > - For more information, see [Notification.image](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117) on the HarmonyOS official website.
	//
	// example:
	//
	// https://example.com/xxx.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// When RenderStyle is set to MULTI_LINE, this field is required to define the content for the multi-line text style. A maximum of three content items are supported.
	InboxContent []*string `json:"InboxContent,omitempty" xml:"InboxContent,omitempty" type:"Repeated"`
	// The JSON string of the HarmonyOS Live View data structure [LiveViewPayload](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V13/push-scenariozed-api-request-param-V13#section66881469306). For development and integration, refer to the [HarmonyOS Live View push guide](https://help.aliyun.com/document_detail/2982112.html).
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
	// The unique identifier (notifyId) for each message displayed in the notification bar. If this parameter is not provided, the push service automatically generates a unique identifier. Different notification messages can use the same notifyId to allow new messages to overwrite old ones. For more information, see [Notification.notifyId](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117) on the HarmonyOS official website.
	//
	// example:
	//
	// 123456
	NotifyId *int32 `json:"NotifyId,omitempty" xml:"NotifyId,omitempty"`
	// The receipt ID of the HarmonyOS channel. You can view this receipt ID in the receipt parameter settings of the HarmonyOS channel push operations platform.
	//
	// > - If the default receipt configuration on the HarmonyOS channel push operations platform is set to Alibaba Cloud receipt, you do not need to provide this parameter. Otherwise, we recommend that you configure the default receipt ID for the HarmonyOS channel in the Alibaba Cloud EMAS shift push console first.
	//
	// > - For details, see [pushOptions.receiptId](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section418321011212) on the HarmonyOS official website.
	//
	// example:
	//
	// RCPB***DFD5
	ReceiptId *string `json:"ReceiptId,omitempty" xml:"ReceiptId,omitempty"`
	// The notification message style. This parameter is optional. Default value: NORMAL.
	//
	// example:
	//
	// NORMAL
	RenderStyle *string `json:"RenderStyle,omitempty" xml:"RenderStyle,omitempty"`
	// Uses a notification channel of the specified type.
	//
	// > - This parameter is valid only for Alibaba Cloud proprietary channels.
	//
	// > - For more information, see [SlotType](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/js-apis-notificationmanager-V5#slottype) on the HarmonyOS official website.
	//
	// example:
	//
	// SOCIAL_COMMUNICATION
	SlotType *string `json:"SlotType,omitempty" xml:"SlotType,omitempty"`
	// The custom ringtone file name for HarmonyOS.
	//
	// example:
	//
	// music.mp3
	Sound *string `json:"Sound,omitempty" xml:"Sound,omitempty"`
	// The duration of the custom message notification ringtone, in seconds. Valid values: 1 to 60. If the ringtone duration is shorter than the notification display time, the ringtone loops.
	//
	// example:
	//
	// 2
	SoundDuration *int32 `json:"SoundDuration,omitempty" xml:"SoundDuration,omitempty"`
	// Enables test messages.
	//
	// > - For more information, see [TestMessage](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section418321011212) in the HarmonyOS push parameters.
	//
	// example:
	//
	// true
	TestMessage *bool `json:"TestMessage,omitempty" xml:"TestMessage,omitempty"`
	// The URI corresponding to the ability of the built-in app page.
	//
	// > - If multiple abilities exist, specify the action and URI for each ability. The action is prioritized when searching for the corresponding built-in app page.
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
	// To send iOS notifications through the APNs center, you must specify the corresponding environment context. This parameter is optional. Default value: production environment.
	//
	// - DEV: the development environment, which applies to applications directly installed and debugged by using Xcode.
	//
	// - PRODUCT: the production environment, which applies to applications distributed through the App Store, TestFlight, Ad Hoc, and enterprise distribution.
	//
	// example:
	//
	// DEV
	ApnsEnv *string `json:"ApnsEnv,omitempty" xml:"ApnsEnv,omitempty"`
	// The badge of the iOS application.
	//
	// example:
	//
	// 1
	Badge *int32 `json:"Badge,omitempty" xml:"Badge,omitempty"`
	// Specifies whether to enable the automatic badge increment feature. This parameter is optional. Default value: false.
	//
	// > - This parameter cannot be used together with the badge setting parameter.
	//
	// > - The automatic badge increment feature maintains the badge count for each device on the Alibaba Cloud push server. You must use SDK version 1.9.5 or later to actively synchronize the badge count to the server.
	//
	// example:
	//
	// false
	BadgeAutoIncrement *bool `json:"BadgeAutoIncrement,omitempty" xml:"BadgeAutoIncrement,omitempty"`
	// The category identifier of the iOS notification, which is used to define the interactive behavior and display style of the notification.
	//
	// > - The category must be pre-registered in the application to take effect.
	//
	// > - Different categories can define different sets of actions.
	//
	// example:
	//
	// MESSAGE_REPLY
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The unique identifier used to control notification grouping. Notifications with the same identifier are overwritten.
	//
	// example:
	//
	// order_status_update_12345
	CollapseId *string `json:"CollapseId,omitempty" xml:"CollapseId,omitempty"`
	// The custom extension attributes of the iOS notification.
	//
	// > - The parameter must be passed in standard JSON map format. Incorrect formats cause parsing failures.
	//
	// example:
	//
	// {"attachment": "https://www.example.com/notification_pic.png"}
	ExtParameters *string `json:"ExtParameters,omitempty" xml:"ExtParameters,omitempty"`
	// The interruption level. This parameter is optional. Valid values:
	//
	// - passive: The system adds the notification to the notification list without lighting up the screen or playing a sound.
	//
	// - active: The system displays the notification immediately, lights up the screen, and may play a sound.
	//
	// - time-sensitive: The system presents the notification immediately, lights up the screen, and may play a sound, but does not bypass the system notification controls.
	//
	// - critical: The system displays the notification immediately, lights up the screen, and plays a sound bypassing the silent switch.
	//
	// example:
	//
	// active
	InterruptionLevel *string `json:"InterruptionLevel,omitempty" xml:"InterruptionLevel,omitempty"`
	// The Live Activity parameter object.
	//
	// 	Notice:
	//
	// - Live Activity push supports pushing to a **single device*	- only by using the DEVICE type.
	//
	// - When pushing to a Live Activity, you do not need to specify the title and body parameters.
	LiveActivity *PushTaskNotificationIosLiveActivity `json:"LiveActivity,omitempty" xml:"LiveActivity,omitempty" type:"Struct"`
	// The sound of the iOS notification. Specify the name of the audio file stored in the app bundle or the Library/Sounds directory of the sandbox. For more information, see [How to set the notification sound for iOS push](https://help.aliyun.com/document_detail/48906.html).
	//
	// > - If an empty string ("") is specified, the notification is silent.
	//
	// > - If this parameter is not set, the default value default is used, which indicates the system prompt sound.
	//
	// example:
	//
	// default
	Music *string `json:"Music,omitempty" xml:"Music,omitempty"`
	// Specifies whether to enable extended notifications and allow iOS notifications to be processed by the Notification Service Extension.
	//
	// > - This parameter must be set to true when sending silent notifications.
	//
	// > - The processing time of the extension cannot exceed 30 seconds.
	//
	// > - A timeout causes the notification to display the original content.
	//
	// > - You must add the Notification Service Extension to your application.
	//
	// example:
	//
	// true
	Mutable *bool `json:"Mutable,omitempty" xml:"Mutable,omitempty"`
	// The relevance score of the notification message, which is used to control the priority and display strategy of the notification.
	//
	// example:
	//
	// 0.5
	RelevanceScore *float64 `json:"RelevanceScore,omitempty" xml:"RelevanceScore,omitempty"`
	// Specifies whether to enable the silent push mode.
	//
	// > - When sending a silent notification, you do not need to specify the title and body parameters.
	//
	// example:
	//
	// false
	Silent *bool `json:"Silent,omitempty" xml:"Silent,omitempty"`
	// The subtitle of the iOS notification.
	//
	// example:
	//
	// Please check your order
	Subtitle *string `json:"Subtitle,omitempty" xml:"Subtitle,omitempty"`
	// The thread identifier for iOS notification grouping, which is used to group and collapse related notifications.
	//
	// > - Notifications with the same thread ID are automatically grouped.
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
	// The static pass-through parameters for iOS Live Activity push, which are used to pass unchanged business identification information.
	//
	// > This parameter is required when Event is set to start.
	//
	// example:
	//
	// {
	//
	//   "orderId": "ORD20231201001",
	//
	//   "restaurantName": "Tasty Restaurant",
	//
	//   "customerAddress": "No. xx, xx Road, xx District",
	//
	//   "orderType": "delivery"
	//
	// }
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The type of the Live Activity to be started.
	//
	// > This parameter is required when Event is set to start.
	//
	// example:
	//
	// OrderActivityAttributes
	AttributesType *string `json:"AttributesType,omitempty" xml:"AttributesType,omitempty"`
	// The dynamic pass-through parameters for the Live Activity, which contain status information and changing data that can be updated in real time.
	//
	// > - Avoid excessively frequent updates. An interval of 5 seconds or more is recommended.
	//
	// > - Update multiple fields in batches to reduce the number of pushes.
	//
	// > - Consider the user experience and avoid interface flickering.
	//
	// > - Must be a valid JSON string.
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
	//     "currentStep": "The courier is on the way"
	//
	// }
	ContentState *string `json:"ContentState,omitempty" xml:"ContentState,omitempty"`
	// The retention time of a completed Live Activity on the Lock Screen, which allows users to view the information after the activity ends. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1701439800
	DismissalDate *int64 `json:"DismissalDate,omitempty" xml:"DismissalDate,omitempty"`
	// Starts, updates, or ends a live activity.
	//
	// example:
	//
	// start
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// The unique identifier of the Live Activity, which is used to associate the activity instance on the device with the push target on the server.
	//
	// 	Notice:
	//
	// - The ID must be the same as the forActivityId parameter in the registerLiveActivityPushToken method of the client SDK.
	//
	// - The server uses this ID to locate the specific activity instance during a push.
	//
	// example:
	//
	// FOOD_DELIVERY_ORD20231201001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The expiration timestamp for the content of the iOS Live Activity. The value is a UNIX timestamp in seconds.
	//
	// > - After the specified time is reached, the system automatically marks the activity as stale.
	//
	// > - Stale activities are removed from the Dynamic Island and the Lock Screen.
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
	// The expiration time of the message. The message is not sent after it expires. The maximum retention period is 72 hours.
	//
	// > 	- The time must follow the ISO 8601 standard and use UTC time in the YYYY-MM-DDThh:mm:ssZ format.
	//
	// > 	- The expiration time must meet the following condition: ExpireTime > PushTime + 3 seconds. The 3 seconds are reserved for network and system latency.
	//
	// > 	- Set the expiration time to at least 1 minute for single-target push tasks and at least 10 minutes for full or batch push tasks.
	//
	//
	// 	Notice: For pass-through messages, if the expiration time is not set, the message is sent only to online devices. If the device is offline, the message is discarded.</notice>
	//
	// example:
	//
	// 2025-06-21T12:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The custom identifier of the push task. If JobKey is not empty, this field is included in the receipt log. For more information, see [Receipt logs](https://help.aliyun.com/document_detail/434651.html).
	//
	// example:
	//
	// jobkey1727749697913
	JobKey *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	// The unique identity of the message. This parameter is valid only when the Action parameter is set to CONTINUOUS_PUSH.
	//
	// example:
	//
	// 1174754033128****
	MessageId *int64 `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// Specifies the sending time of the message, which can be up to seven days later. Valid only when the Action parameter is set to SCHEDULED_PUSH.
	//
	// > The time must follow the ISO 8601 standard and use UTC time in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2025-06-19T12:00:00Z
	PushTime *string `json:"PushTime,omitempty" xml:"PushTime,omitempty"`
	// The fallback SMS message configuration.
	//
	// > Currently, only Android and HarmonyOS devices are supported.
	Sms *PushTaskOptionsSms `json:"Sms,omitempty" xml:"Sms,omitempty" type:"Struct"`
	// Specifies whether to automatically truncate excessively long titles and content.
	//
	// >This parameter applies only to vendor channels that explicitly limit the title and content length. It does not apply to channels such as APNs, Huawei, and Honor, which do not limit the title and content length but only limit the total request body size.
	//
	// example:
	//
	// false
	Trim *bool `json:"Trim,omitempty" xml:"Trim,omitempty"`
	// Specifies the sending channels. Valid values:
	//
	// - accs: Alibaba Cloud proprietary channel.
	//
	// - huawei: Huawei channel.
	//
	// - honor: Honor channel.
	//
	// - xiaomi: Xiaomi channel.
	//
	// - oppo: OPPO channel.
	//
	// - vivo: vivo channel.
	//
	// - meizu: Meizu channel.
	//
	// - fcm: Google Firebase channel (HTTP v1 API).
	//
	// - apns: APNs channel.
	//
	// - harmony: HarmonyOS channel.
	//
	// > - If this parameter is not configured, all channels are available.
	//
	// > - If this parameter is configured, only the specified channels are used.
	//
	// > - If the configured channels conflict with the sending policy (for example, iOS notifications only use the APNs channel, but apns is not included in this parameter), the message is not sent.
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
	// The delay time before the SMS message is triggered. Unit: seconds.
	//
	// If you use SMS linkage, this parameter is required. Set this parameter to at least 15 seconds and no more than 3 days to avoid duplicate SMS and push messages.
	//
	// > When SMS linkage is used, the ExpireTime parameter becomes invalid. The expiration time of the notification is calculated based on the DelaySecs parameter. The expiration time is the current time plus the DelaySecs value.
	//
	// example:
	//
	// 150
	DelaySecs *int64 `json:"DelaySecs,omitempty" xml:"DelaySecs,omitempty"`
	// The key-value pairs of the variables in the SMS template.
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
	// The name of the SMS template. You can obtain the name from the SMS template management page. The name is assigned by the system, not specified by the developer.
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
	// The platform type. This parameter is optional.
	//
	// example:
	//
	// IOS
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The type of the push target.
	//
	// 	Notice:
	//
	// The MassPushV2 batch push API and the CONTINUOUS_PUSH continuous push support only the following three target types:
	//
	// - DEVICE (device)
	//
	// - ACCOUNT (account)
	//
	// - ALIAS (alias)
	//
	// example:
	//
	// DEVICE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The push targets based on Target.Type. Separate multiple targets with commas. The following section describes the target types and target values in detail:
	//
	// >	- DEVICE: device IDs, such as deviceid1,deviceid2. Maximum of 1,000.
	//
	// >	- ACCOUNT: account IDs, such as account1,account2. Maximum of 1,000.
	//
	// >	- ALIAS: aliases, such as alias1,alias2. Maximum of 1,000.
	//
	// >	- TAG: one or more tags. For the format, see [Tag format specifications](https://help.aliyun.com/document_detail/434847.html).
	//
	// >	- ALL: all targets. No value is required. Pushing to all targets may increase costs. Use this option with caution.
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
