// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMassPushRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *MassPushRequest
	GetAppKey() *int64
	SetIdempotentToken(v string) *MassPushRequest
	GetIdempotentToken() *string
	SetPushTask(v []*MassPushRequestPushTask) *MassPushRequest
	GetPushTask() []*MassPushRequestPushTask
}

type MassPushRequest struct {
	// The AppKey.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The idempotent parameter used to prevent duplicate pushes caused by API client retries. If the same IdempotentToken is used within 15 minutes, only one push is performed, and subsequent calls return the result of the first successful push.
	//
	// >
	//
	// > - The parameter format is a standard 36-character UUID (8-4-4-4-12). Each valid character is a hexadecimal digit in the range of 0-9 or a-f, and it is case-insensitive.
	//
	// > - This parameter is only used to prevent duplicate pushes caused by retries. It cannot prevent duplicate pushes caused by concurrent calls.
	//
	// example:
	//
	// c8016d13-6e76-410c-9bda-769383d11787
	IdempotentToken *string `json:"IdempotentToken,omitempty" xml:"IdempotentToken,omitempty"`
	// The array of independent push tasks.
	//
	// This parameter is required.
	PushTask []*MassPushRequestPushTask `json:"PushTask,omitempty" xml:"PushTask,omitempty" type:"Repeated"`
}

func (s MassPushRequest) String() string {
	return dara.Prettify(s)
}

func (s MassPushRequest) GoString() string {
	return s.String()
}

func (s *MassPushRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *MassPushRequest) GetIdempotentToken() *string {
	return s.IdempotentToken
}

func (s *MassPushRequest) GetPushTask() []*MassPushRequestPushTask {
	return s.PushTask
}

func (s *MassPushRequest) SetAppKey(v int64) *MassPushRequest {
	s.AppKey = &v
	return s
}

func (s *MassPushRequest) SetIdempotentToken(v string) *MassPushRequest {
	s.IdempotentToken = &v
	return s
}

func (s *MassPushRequest) SetPushTask(v []*MassPushRequestPushTask) *MassPushRequest {
	s.PushTask = v
	return s
}

func (s *MassPushRequest) Validate() error {
	if s.PushTask != nil {
		for _, item := range s.PushTask {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MassPushRequestPushTask struct {
	// The activity to be opened by the notification.
	//
	// This parameter is valid only when PushTask.N.AndroidOpenType is set to ACTIVITY. Example: `com.alibaba.cloudpushdemo.bizactivity`.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidActivity *string `json:"AndroidActivity,omitempty" xml:"AndroidActivity,omitempty"`
	// Specifies the increment value for the badge. The value is added to the existing badge number. Valid values: 1 to 99.
	//
	// > This parameter is valid only for push notifications sent through Huawei or Honor vendor channels. If both AndroidBadgeAddNum and AndroidBadgeSetNum are specified, AndroidBadgeSetNum takes precedence.
	//
	// example:
	//
	// 1
	AndroidBadgeAddNum *int32 `json:"AndroidBadgeAddNum,omitempty" xml:"AndroidBadgeAddNum,omitempty"`
	// The fully qualified class name of the entry Activity for badge settings.
	//
	// > This parameter is valid only for push notifications sent through Huawei or Honor vendor channels.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidBadgeClass *string `json:"AndroidBadgeClass,omitempty" xml:"AndroidBadgeClass,omitempty"`
	// Specifies the fixed number for the badge. Valid values: 0 to 99.
	//
	// > This parameter takes effect only on Huawei and Honor channels when push notifications are sent through vendor channels. When push notifications are sent through the Alibaba Cloud proprietary channel, this parameter takes effect only on Huawei, Honor, and vivo devices.
	//
	// example:
	//
	// 5
	AndroidBadgeSetNum *int32 `json:"AndroidBadgeSetNum,omitempty" xml:"AndroidBadgeSetNum,omitempty"`
	// The body in long text mode. Maximum length: 1,000 bytes. One Chinese character counts as 3 bytes. The actual limit is subject to the specific vendor channel during sending.
	//
	// Currently supported by:
	//
	// - Huawei: EMUI 10 or later
	//
	// - Honor: Magic UI 4.0 or later
	//
	// - Xiaomi: MIUI 10 or later
	//
	// - OPPO: ColorOS 5.0 or later
	//
	// - Meizu: Flyme
	//
	// - Proprietary channel: Android SDK 3.6.0 or later
	//
	// If this parameter is not provided in long text mode, the first non-empty value from Body and AndroidPopupBody is used.
	//
	// example:
	//
	// Sample long text
	AndroidBigBody *string `json:"AndroidBigBody,omitempty" xml:"AndroidBigBody,omitempty"`
	// The image URL in big image mode. Currently supported by the proprietary channel with Android SDK 3.6.0 or later.
	//
	// example:
	//
	// https://imag.example.com/image.png
	AndroidBigPictureUrl *string `json:"AndroidBigPictureUrl,omitempty" xml:"AndroidBigPictureUrl,omitempty"`
	// The title in long text mode. Maximum length: 200 bytes. One Chinese character counts as 3 bytes.
	//
	// - This parameter is currently supported only by the Honor channel and the Huawei channel with EMUI 11 or later.
	//
	// - If this parameter is not provided in long text mode, the first non-empty value from Title and AndroidPopupTitle is used.
	//
	// example:
	//
	// Sample long title
	AndroidBigTitle *string `json:"AndroidBigTitle,omitempty" xml:"AndroidBigTitle,omitempty"`
	// The extended properties of the notification. This parameter does not take effect when the push type PushType is set to MESSAGE.
	//
	// This parameter must be passed in JSON map format. Otherwise, a parsing error occurs.
	//
	// example:
	//
	// {"key1":"value1","api_name":"PushNoticeToAndroidRequest"}
	AndroidExtParameters *string `json:"AndroidExtParameters,omitempty" xml:"AndroidExtParameters,omitempty"`
	// The notification type for the Honor channel. Valid values:
	//
	// - **0**: Formal notification. This is the default value.
	//
	// - **1**: Test notification.
	//
	// > Each application can send 1,000 test notifications per day, which is not subject to the daily push limit per device.
	//
	// example:
	//
	// 1
	AndroidHonorTargetUserType *int32 `json:"AndroidHonorTargetUserType,omitempty" xml:"AndroidHonorTargetUserType,omitempty"`
	// The Huawei quick notification parameter. Valid values:
	//
	// - **0**: Send a standard Huawei notification. This is the default value.
	//
	// - **1**: Send a Huawei quick notification.
	//
	// example:
	//
	// 1
	AndroidHuaweiBusinessType *int32 `json:"AndroidHuaweiBusinessType,omitempty" xml:"AndroidHuaweiBusinessType,omitempty"`
	// The JSON string of the Huawei Android Live Notification data structure [LiveNotificationPayload](https://developer.huawei.com/consumer/cn/doc/HMSCore-References/rest-live-0000001562939968#ZH-CN_TOPIC_0000001700850537__p195121620102511). For development integration, refer to [Huawei Live Notification push guide](https://help.aliyun.com/document_detail/2983768.html).
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
	AndroidHuaweiLiveNotificationPayload *string `json:"AndroidHuaweiLiveNotificationPayload,omitempty" xml:"AndroidHuaweiLiveNotificationPayload,omitempty"`
	// The receipt ID for the Huawei channel. You can view this receipt ID in the receipt parameter settings of the Huawei channel push operations platform.
	//
	// > If the default receipt configuration on the Huawei channel push operations platform is set to Alibaba Cloud receipt, you do not need to provide this parameter. Otherwise, configure the default receipt ID for the Huawei channel in the Alibaba Cloud EMAS Mobile Push console first.
	//
	// example:
	//
	// RCP4C123456
	AndroidHuaweiReceiptId *string `json:"AndroidHuaweiReceiptId,omitempty" xml:"AndroidHuaweiReceiptId,omitempty"`
	// The notification type for the Huawei channel. Valid values:
	//
	// - **0**: Formal notification. This is the default value.
	//
	// - **1**: Test notification.
	//
	// > Each application can send 500 test notifications per day, which is not subject to the daily push limit per device.
	//
	// example:
	//
	// 1
	AndroidHuaweiTargetUserType *int32 `json:"AndroidHuaweiTargetUserType,omitempty" xml:"AndroidHuaweiTargetUserType,omitempty"`
	// The URL of the right-side icon. Currently supported by:
	//
	// - Huawei EMUI: Applicable only in long text mode and inbox mode.
	//
	// - Honor Magic UI: Applicable only in long text mode.
	//
	// - Proprietary channel: Android SDK 3.5.0 or later.
	//
	// example:
	//
	// https://imag.example.com/image.png
	AndroidImageUrl *string `json:"AndroidImageUrl,omitempty" xml:"AndroidImageUrl,omitempty"`
	// The body in inbox mode. The content must be a valid JSON array with no more than five elements. Currently supported by:
	//
	// - Huawei: EMUI 9 or later
	//
	// - Honor: Magic UI 4.0 or later
	//
	// - Xiaomi: MIUI 10 or later
	//
	// - OPPO: ColorOS 5.0 or later
	//
	// - Proprietary channel: Android SDK 3.6.0 or later
	//
	// example:
	//
	// ["First line","Second line"]
	AndroidInboxBody *string `json:"AndroidInboxBody,omitempty" xml:"AndroidInboxBody,omitempty"`
	// The Meizu message type. Valid values:
	//
	// - 0: Public message. This is the default value.
	//
	// - 1: Private message.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 0
	AndroidMeizuNoticeMsgType *int32 `json:"AndroidMeizuNoticeMsgType,omitempty" xml:"AndroidMeizuNoticeMsgType,omitempty"`
	// Purpose 1: After you complete the [self-classification permission application](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835?#section3410731125514), this parameter is used to identify the message type, determine the [message reminder method](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#ZH-CN_TOPIC_0000001149358835__p3850133955718), and accelerate the sending of specific types of messages. For valid values, refer to the [Message classification standards](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section1076611477914) in the official Huawei push documentation, and fill in the "Cloud notification category value" or "Local notification category value" in the document table.
	//
	// Purpose 2: After you [apply for special permissions](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/faq-0000001050042183#section037425218509), this parameter is used to identify high-priority pass-through scenarios. Valid values:
	//
	// - VOIP: Audio and video calls.
	//
	// - PLAY_VOICE: Voice broadcast.
	//
	// > If the "Cloud notification category value" is "Not applicable", the message is sent through the Alibaba Cloud proprietary channel. If the "Local notification category value" is "Not applicable", the message is sent through the Huawei channel.
	//
	// example:
	//
	// SUBSCRIPTION
	AndroidMessageHuaweiCategory *string `json:"AndroidMessageHuaweiCategory,omitempty" xml:"AndroidMessageHuaweiCategory,omitempty"`
	// The delivery priority for Huawei channel notifications. Valid values:
	//
	// - HIGH
	//
	// - NORMAL
	//
	// You must apply for this permission. For more information, see [Application link](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/faq-0000001050042183#section037425218509).
	//
	// example:
	//
	// HIGH
	AndroidMessageHuaweiUrgency *string `json:"AndroidMessageHuaweiUrgency,omitempty" xml:"AndroidMessageHuaweiUrgency,omitempty"`
	// OPPO message categorization includes two categories for management: Communication and Service, and Content and Marketing.
	//
	// Communication and Service (requires permission application):
	//
	// - IM: Instant chat, audio calls, and video calls.
	//
	// - ACCOUNT: Individual account and asset changes.
	//
	// - DEVICE_REMINDER: Personal device reminders.
	//
	// - ORDER: Personal order and logistics status changes.
	//
	// - TODO: Personal schedules and to-do lists.
	//
	// - SUBSCRIPTION: Personal subscriptions.
	//
	// Content and Marketing:
	//
	// - NEWS: News and information.
	//
	// - CONTENT: Content recommendations.
	//
	// - MARKETING: Platform activities.
	//
	// - SOCIAL: Social updates.
	//
	// For more information, see [OPUSH message classification rules](https://open.oppomobile.com/new/developmentDoc/info?id=13189).
	//
	// example:
	//
	// MARKETING
	AndroidMessageOppoCategory *string `json:"AndroidMessageOppoCategory,omitempty" xml:"AndroidMessageOppoCategory,omitempty"`
	// The notification bar message reminder level for the OPPO channel. Valid values:
	//
	// - 1: Notification bar.
	//
	// - 2: Notification bar, lock screen, ringtone, and vibration. This is the default notification level for communication and service messages.
	//
	// - 16: Notification bar, lock screen, ringtone, vibration, and banner. Requires permission application.
	//
	// > When you use the AndroidMessageOppoNotifyLevel parameter, you must also pass the AndroidMessageOppoCategory parameter.
	//
	// example:
	//
	// 1
	AndroidMessageOppoNotifyLevel *int32 `json:"AndroidMessageOppoNotifyLevel,omitempty" xml:"AndroidMessageOppoNotifyLevel,omitempty"`
	// vivo message categorization includes two categories for management: System messages and Operational messages.
	//
	// System messages:
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
	// - SUBSCRIPTION: Subscription reminders.
	//
	// Operational messages:
	//
	// - NEWS: News.
	//
	// - CONTENT: Content recommendations.
	//
	// - MARKETING: Operational activities.
	//
	// - SOCIAL: Social updates.
	//
	// > For more information, see [Classification description](https://dev.vivo.com.cn/documentCenter/doc/359#s-ef3qugc3).
	//
	// example:
	//
	// TODO
	AndroidMessageVivoCategory *string `json:"AndroidMessageVivoCategory,omitempty" xml:"AndroidMessageVivoCategory,omitempty"`
	// The notification sound for the Huawei vendor channel. Specify the name of the audio file stored in the app/src/main/res/raw/ directory of the client project. Do not include the file extension.
	//
	// If this parameter is not specified, the default ringtone is used.
	//
	// example:
	//
	// alicloud_notification_sound
	AndroidMusic *string `json:"AndroidMusic,omitempty" xml:"AndroidMusic,omitempty"`
	// The priority of the Android notification display position in the notification bar. Valid values: -2, -1, 0, 1, and 2.
	//
	// example:
	//
	// 0
	AndroidNotificationBarPriority *int32 `json:"AndroidNotificationBarPriority,omitempty" xml:"AndroidNotificationBarPriority,omitempty"`
	// The custom notification bar style for Android. Valid values: 1 to 100.
	//
	// example:
	//
	// 2
	AndroidNotificationBarType *int32 `json:"AndroidNotificationBarType,omitempty" xml:"AndroidNotificationBarType,omitempty"`
	// The channel ID of the Android application. This value must match the channel ID in the application.
	//
	// - For more information about the NotificationChannel parameter, see [FAQ: Notifications are not received on devices running Android 8.0 or later](https://help.aliyun.com/document_detail/67398.html).
	//
	// - Because the channel_id of the OPPO private message channel is the same as the channel ID of the application, this value is used as the channel_id for OPPO channel pushes.
	//
	// - This value is used as the channel_id for Huawei, FCM, and Alibaba Cloud proprietary channel pushes.
	//
	// example:
	//
	// 1
	AndroidNotificationChannel *string `json:"AndroidNotificationChannel,omitempty" xml:"AndroidNotificationChannel,omitempty"`
	// The message group. For messages in the same group, only the latest message and the total number of received messages in the group are displayed in the notification bar. All messages are not displayed and cannot be expanded. Currently supported by:
	//
	// - Huawei vendor channel
	//
	// - Honor vendor channel
	//
	// - Proprietary channel with Android SDK 3.9.1 or earlier
	//
	// > The proprietary channel no longer supports this parameter in Android SDK 3.9.2 or later.
	//
	// example:
	//
	// group-1
	AndroidNotificationGroup *string `json:"AndroidNotificationGroup,omitempty" xml:"AndroidNotificationGroup,omitempty"`
	// The importance parameter for Honor notification message classification, which determines the notification behavior on user devices. Valid values:
	//
	// - LOW: Information and marketing messages.
	//
	// - NORMAL: Service and communication messages.
	//
	// You must apply for this permission on the Honor platform. For more information, see [Application link](https://developer.honor.com/cn/docs/11002/guides/notification-class#%E8%87%AA%E5%88%86%E7%B1%BB%E6%9D%83%E7%9B%8A%E7%94%B3%E8%AF%B7).
	//
	// example:
	//
	// LOW
	AndroidNotificationHonorChannel *string `json:"AndroidNotificationHonorChannel,omitempty" xml:"AndroidNotificationHonorChannel,omitempty"`
	// The importance parameter for Huawei notification message classification, which determines the notification behavior on user devices. Valid values:
	//
	// - LOW: Information and marketing messages.
	//
	// - NORMAL: Service and communication messages.
	//
	// > - For the Huawei channel, use AndroidMessageHuaweiCategory for notification classification instead of AndroidNotificationHuaweiChannel.
	//
	// > - You must apply for this permission on the Huawei platform. For more information, see [Application link](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section893184112272).
	//
	// example:
	//
	// LOW
	AndroidNotificationHuaweiChannel *string `json:"AndroidNotificationHuaweiChannel,omitempty" xml:"AndroidNotificationHuaweiChannel,omitempty"`
	// The unique identifier for each message when displayed in the notification bar. Different notification bar messages can share the same NotifyId, allowing new notification bar messages to overwrite old ones.
	//
	// example:
	//
	// 100001
	AndroidNotificationNotifyId *int32 `json:"AndroidNotificationNotifyId,omitempty" xml:"AndroidNotificationNotifyId,omitempty"`
	// The message thread ID. Messages in the same group are collapsed and can be expanded in the notification bar. Notifications in different groups are displayed separately. Currently supported by:
	//
	// - Proprietary channel with Android SDK 3.9.2 or later
	//
	// example:
	//
	// thread-1
	AndroidNotificationThreadId *string `json:"AndroidNotificationThreadId,omitempty" xml:"AndroidNotificationThreadId,omitempty"`
	// The vivo notification message classification. Valid values:
	//
	// - 0: Operational messages. This is the default value.
	//
	// - 1: System messages.
	//
	// > - For the vivo channel, use AndroidMessageVivoCategory for notification classification instead of AndroidNotificationVivoChannel.
	//
	// > - You must apply for this permission on the vivo platform. For more information, see [Application link](https://dev.vivo.com.cn/documentCenter/doc/359).
	//
	// example:
	//
	// 0
	AndroidNotificationVivoChannel *string `json:"AndroidNotificationVivoChannel,omitempty" xml:"AndroidNotificationVivoChannel,omitempty"`
	// The channel ID for Xiaomi notification types. You must apply for this permission on the Xiaomi platform. For more information, see [Application link](https://dev.mi.com/console/doc/detail?pId=2422#_4).
	//
	// > - A single application can apply for a maximum of eight channels on the Xiaomi channel. Plan ahead.
	//
	// example:
	//
	// michannel
	AndroidNotificationXiaomiChannel *string `json:"AndroidNotificationXiaomiChannel,omitempty" xml:"AndroidNotificationXiaomiChannel,omitempty"`
	// The reminder method for the notification. Valid values:
	//
	// - **VIBRATE**: Vibration. This is the default value.
	//
	// - **SOUND**: Sound.
	//
	// - **BOTH**: Sound and vibration.
	//
	// - **NONE**: Silent.
	//
	// example:
	//
	// VIBRATE
	AndroidNotifyType *string `json:"AndroidNotifyType,omitempty" xml:"AndroidNotifyType,omitempty"`
	// The action to perform after the notification is clicked. Valid values:
	//
	// - APPLICATION: Open the application. This is the default value.
	//
	// - ACTIVITY: Open the specified Android activity.
	//
	// - URL: Open the specified URL.
	//
	// - NONE: Perform no action.
	//
	// example:
	//
	// APPLICATION
	AndroidOpenType *string `json:"AndroidOpenType,omitempty" xml:"AndroidOpenType,omitempty"`
	// The URL to be opened after the Android device receives the push. This parameter is valid only when PushTask.N.AndroidOpenType is set to URL.
	//
	// example:
	//
	// https://www.example.com
	AndroidOpenUrl *string `json:"AndroidOpenUrl,omitempty" xml:"AndroidOpenUrl,omitempty"`
	// The JSON character string of the OPPO Fluid Cloud intent delete data structure [data](https://open.oppomobile.com/documentation/page/info?id=13578). This parameter is invalid if the AndroidOppoIntelligentIntent parameter is already specified. For development integration, see References: [OPPO Fluid Cloud push guide](https://help.aliyun.com/document_detail/2997310.html).
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
	AndroidOppoDeleteIntentData *string `json:"AndroidOppoDeleteIntentData,omitempty" xml:"AndroidOppoDeleteIntentData,omitempty"`
	// The JSON character string of the OPPO Fluid Cloud intent sharing data structure [IntelligentIntent](https://open.oppomobile.com/documentation/page/info?id=13565). For development integration, see References: [OPPO Fluid Cloud push guide](https://help.aliyun.com/document_detail/2997310.html).
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
	AndroidOppoIntelligentIntent *string `json:"AndroidOppoIntelligentIntent,omitempty" xml:"AndroidOppoIntelligentIntent,omitempty"`
	// The push environment for OPPO Fluid Cloud. Valid values:
	//
	// - **0**: Formal environment. This is the default value.
	//
	// - **1**: Test environment.
	//
	// > To use the OPPO Fluid Cloud test environment, you must build it on the client side by referring to [Environment setup](https://open.oppomobile.com/documentation/page/info?id=13590).
	//
	// example:
	//
	// 1
	AndroidOppoIntentEnv *int32 `json:"AndroidOppoIntentEnv,omitempty" xml:"AndroidOppoIntentEnv,omitempty"`
	// Deprecated
	//
	// The content parameter of the OPPO private message template.
	AndroidOppoPrivateContentParameters map[string]*string `json:"AndroidOppoPrivateContentParameters,omitempty" xml:"AndroidOppoPrivateContentParameters,omitempty"`
	// Deprecated
	//
	// The OPPO private message template ID.
	//
	// 	Warning: The OPPO private message template feature is no longer supported by MaasPush. To use this feature, call the Push, PushV2, or MassPushV2 operation.
	//
	// example:
	//
	// 687557242b1634hzef3zd5013
	AndroidOppoPrivateMsgTemplateId *string `json:"AndroidOppoPrivateMsgTemplateId,omitempty" xml:"AndroidOppoPrivateMsgTemplateId,omitempty"`
	// Deprecated
	//
	// The title parameter of the OPPO private message template.
	AndroidOppoPrivateTitleParameters map[string]*string `json:"AndroidOppoPrivateTitleParameters,omitempty" xml:"AndroidOppoPrivateTitleParameters,omitempty"`
	// The activity to navigate to after the notification is clicked.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidPopupActivity *string `json:"AndroidPopupActivity,omitempty" xml:"AndroidPopupActivity,omitempty"`
	// The body content in auxiliary pop-up window mode. This parameter is required when the AndroidPopupActivity parameter is not empty.
	//
	// Maximum length: 200 characters. Both Chinese and English characters count as one character.
	//
	// If a vendor channel is used, the body must also meet the limits of the vendor channel. For more information, see [Limits on auxiliary channel pushes for Android](https://help.aliyun.com/document_detail/165253.html).
	//
	// example:
	//
	// hello
	AndroidPopupBody *string `json:"AndroidPopupBody,omitempty" xml:"AndroidPopupBody,omitempty"`
	// The title content in auxiliary pop-up window mode. This parameter is required when the AndroidPopupActivity parameter is not empty.
	//
	// Maximum length: 50 characters. Both Chinese and English characters count as one character.
	//
	// If a vendor channel is used, the title must also meet the limits of the vendor channel. For more information, see [Limits on auxiliary channel pushes for Android](https://help.aliyun.com/document_detail/165253.html).
	//
	// example:
	//
	// hello
	AndroidPopupTitle *string `json:"AndroidPopupTitle,omitempty" xml:"AndroidPopupTitle,omitempty"`
	// Specifies whether to use the auxiliary pop-up window feature when the device is offline during a message push. Default value: false. This parameter takes effect only when PushType is set to MESSAGE.
	//
	// If the message-to-notification push is successful, the data displayed in the received notification is the value of the AndroidPopupTitle and AndroidPopupBody parameters set on the server. When you click the notification, the data obtained in the onSysNoticeOpened method of the auxiliary pop-up window is the value of the Title and Body parameters set on the server.
	//
	// example:
	//
	// true
	AndroidRemind *bool `json:"AndroidRemind,omitempty" xml:"AndroidRemind,omitempty"`
	// The notification style. Valid values:
	//
	// - **0**: Standard mode. This is the default value.
	//
	// - **1**: Long text mode. Supported by Huawei, Honor, Xiaomi, OPPO, Meizu, and the proprietary channel.
	//
	// - **2**: Big image mode. Supported by the proprietary channel, but not supported by Xiaomi devices.
	//
	// - **3**: Inbox mode. Supported by Huawei, Honor, Xiaomi, OPPO, and the proprietary channel.
	//
	// > This parameter is required if a non-standard mode is used.
	//
	// example:
	//
	// 1
	AndroidRenderStyle *string `json:"AndroidRenderStyle,omitempty" xml:"AndroidRenderStyle,omitempty"`
	// Settings for the vendor channel notification type. Valid values:
	//
	// - **0**: Formal notification. This is the default value.
	//
	// - **1**: Test notification.
	//
	// > - The parameter settings are equivalent to configuring the AndroidHuaweiTargetUserType, AndroidHonorTargetUserType, AndroidVivoPushMode, and AndroidOppoIntentEnv parameters at the same time. The corresponding parameters for specific vendor channels can overwrite this parameter.
	//
	// > - Currently supported by: Huawei channel, Honor channel, vivo channel, and OPPO Fluid Cloud.
	//
	// example:
	//
	// 1
	AndroidTargetUserType *int32 `json:"AndroidTargetUserType,omitempty" xml:"AndroidTargetUserType,omitempty"`
	// The JSON character string of the vivo Atomic Island data structure [liveMessage](https://dev.vivo.com.cn/documentCenter/doc/896#s-fdagzbd4). For development integration, see References: [vivo Atomic Island push guide](https://www.alibabacloud.com/help/en/document_detail/3030718.html).
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
	AndroidVivoLiveMessage *string `json:"AndroidVivoLiveMessage,omitempty" xml:"AndroidVivoLiveMessage,omitempty"`
	// The notification type for the vivo channel. Valid values:
	//
	// - **0**: Formal push. This is the default value.
	//
	// - **1**: Test push.
	//
	// > For test pushes, configure the test devices in the vivo console in advance. You can obtain the RegId of the test device by searching for "onReceiveRegId regId" in the device startup log.
	//
	// example:
	//
	// 1
	AndroidVivoPushMode *int32 `json:"AndroidVivoPushMode,omitempty" xml:"AndroidVivoPushMode,omitempty"`
	// The receipt ID for the vivo channel. You can view this receipt ID in the application information of the push service on the vivo open platform.
	//
	// > If the default receipt configuration on the vivo open platform is set to Alibaba Cloud receipt, you do not need to provide this parameter. Otherwise, configure the default receipt ID for the vivo channel in the Alibaba Cloud EMAS Mobile Push console first.
	//
	// example:
	//
	// 123
	AndroidVivoReceiptId *string `json:"AndroidVivoReceiptId,omitempty" xml:"AndroidVivoReceiptId,omitempty"`
	// The vivo template ID.
	//
	// example:
	//
	// t1-ge3-lq
	AndroidVivoTemplateId *string `json:"AndroidVivoTemplateId,omitempty" xml:"AndroidVivoTemplateId,omitempty"`
	// The parameters of the vivo template. The value is a JSON object string.
	//
	// example:
	//
	// {"keywords1":"S123"}
	AndroidVivoTemplateParams *string `json:"AndroidVivoTemplateParams,omitempty" xml:"AndroidVivoTemplateParams,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- This parameter is deprecated. All third-party auxiliary pop-up windows are uniformly supported by the new **AndroidPopupActivity*	- parameter.
	//
	// example:
	//
	// None
	AndroidXiaoMiActivity *string `json:"AndroidXiaoMiActivity,omitempty" xml:"AndroidXiaoMiActivity,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- This parameter is deprecated. All third-party auxiliary pop-up windows are uniformly supported by the new **AndroidPopupBody*	- parameter.
	//
	// example:
	//
	// None
	AndroidXiaoMiNotifyBody *string `json:"AndroidXiaoMiNotifyBody,omitempty" xml:"AndroidXiaoMiNotifyBody,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- This parameter is deprecated. All third-party auxiliary pop-up windows are uniformly supported by the new **AndroidPopupTitle*	- parameter.
	//
	// example:
	//
	// None
	AndroidXiaoMiNotifyTitle *string `json:"AndroidXiaoMiNotifyTitle,omitempty" xml:"AndroidXiaoMiNotifyTitle,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- This parameter is deprecated. Since August 2023, Xiaomi no longer supports dynamically setting small icons, right-side icons, and big images during pushes on new devices or systems.
	//
	// example:
	//
	// https://f6.market.xiaomi.com/download/MiPass/aaa/bbb.png
	AndroidXiaomiBigPictureUrl *string `json:"AndroidXiaomiBigPictureUrl,omitempty" xml:"AndroidXiaomiBigPictureUrl,omitempty"`
	// The JSON character string of the Xiaomi Super Island data structure [miui.focus.param](https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131). For development integration, see References: [Xiaomi Super Island push guide](https://www.alibabacloud.com/help/en/document_detail/3037956.html).
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
	AndroidXiaomiFocusParam *string `json:"AndroidXiaomiFocusParam,omitempty" xml:"AndroidXiaomiFocusParam,omitempty"`
	// The JSON character string of the Xiaomi Super Island data images [miui.focus.pic_xxx](https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131). For development integration, see References: [Xiaomi Super Island push guide](https://www.alibabacloud.com/help/en/document_detail/3037956.html).
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
	AndroidXiaomiFocusPics *string `json:"AndroidXiaomiFocusPics,omitempty" xml:"AndroidXiaomiFocusPics,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- This parameter is deprecated. Since August 2023, Xiaomi no longer supports dynamically setting small icons, right-side icons, and big images during pushes on new devices or systems.
	//
	// example:
	//
	// https://imag.example.com/image.png
	AndroidXiaomiImageUrl *string `json:"AndroidXiaomiImageUrl,omitempty" xml:"AndroidXiaomiImageUrl,omitempty"`
	// The Xiaomi private message template ID.
	//
	// example:
	//
	// P10645
	AndroidXiaomiTemplateId *string `json:"AndroidXiaomiTemplateId,omitempty" xml:"AndroidXiaomiTemplateId,omitempty"`
	// The parameters of the Xiaomi private message template. The value is a JSON string.
	//
	// example:
	//
	// {"keywords1":"Tom","keywords2":"phone"}
	AndroidXiaomiTemplateParams *string `json:"AndroidXiaomiTemplateParams,omitempty" xml:"AndroidXiaomiTemplateParams,omitempty"`
	// The content of the notification or message for Android and HarmonyOS pushes, or the content of the message or notification for iOS pushes. The content size is limited. For more information, see [Limits](https://help.aliyun.com/document_detail/92832.html).
	//
	// example:
	//
	// hello
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The device type. Valid values:
	//
	// - HARMONY: HarmonyOS device
	//
	// - iOS: iOS device
	//
	// - ANDROID: Android device
	//
	// - ALL: If the AppKey belongs to a legacy dual-platform application, this value indicates that the message is pushed to both Android and iOS devices. If the AppKey belongs to a new single-platform application, this value has the same effect as specifying the device type corresponding to the application type.
	//
	// This parameter is required.
	//
	// example:
	//
	// HARMONY
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The expiration time of the offline message or notification. This parameter is used together with StoreOffline. An expired message or notification is not sent. The maximum retention period is 72 hours. Default value: 72 hours.
	//
	// Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC. The expiration time cannot be earlier than the current time or 3 seconds after the scheduled sending time (ExpireTime > PushTime + 3 seconds). The 3-second buffer is used to compensate for network and system latency. Set the expiration time to at least 1 minute for single pushes and at least 10 minutes for full pushes and batch pushes.
	//
	// example:
	//
	// 2019-02-20T00:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The action corresponding to the ability of the built-in page in the app.
	//
	// 	Notice: When HarmonyActionType is set to APP_CUSTOM_PAGE, you must specify at least one of the HarmonyUri and HarmonyAction fields.
	//
	// For more information, see [ClickAction.action](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section152462191216) on the HarmonyOS official website.
	//
	// example:
	//
	// com.example.action
	HarmonyAction *string `json:"HarmonyAction,omitempty" xml:"HarmonyAction,omitempty"`
	// The action to perform after the notification is clicked. Valid values:
	//
	// - APP_HOME_PAGE: Opens the app home page.
	//
	// - APP_CUSTOM_PAGE: Opens a custom page in the app.
	//
	// example:
	//
	// APP_HOME_PAGE
	HarmonyActionType *string `json:"HarmonyActionType,omitempty" xml:"HarmonyActionType,omitempty"`
	// The increment value for the badge of the HarmonyOS app. For more information, see [Description of the addNum field for HarmonyOS badges](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section266310382145).</br>
	//
	// This parameter is supported in HarmonyOS SDK 1.2.0 and later.
	//
	// example:
	//
	// 1
	HarmonyBadgeAddNum *int32 `json:"HarmonyBadgeAddNum,omitempty" xml:"HarmonyBadgeAddNum,omitempty"`
	// The fixed number for the badge of the HarmonyOS app. For more information, see [Description of the setNum field for HarmonyOS badges](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section266310382145).</br>
	//
	// This parameter is supported in HarmonyOS SDK 1.2.0 and later.
	//
	// example:
	//
	// 1
	HarmonyBadgeSetNum *int32 `json:"HarmonyBadgeSetNum,omitempty" xml:"HarmonyBadgeSetNum,omitempty"`
	// The category of the notification message. After you apply for the notification message self-categorization permission, this parameter is used as the identity of the message type. Different notification message types affect the display and reminder methods of messages. Valid values:
	//
	// - IM: Instant chat.
	//
	// - VOIP: Audio and video calls.
	//
	// - SUBSCRIPTION: subscribe.
	//
	// - TRAVEL: Travel.
	//
	// - HEALTH: Health.
	//
	// - WORK: Work reminders.
	//
	// - ACCOUNT: Account updates.
	//
	// - EXPRESS: Orders and logistics.
	//
	// - FINANCE: Finance.
	//
	// - DEVICE_REMINDER: Device reminders.
	//
	// - MAIL: Emails.
	//
	// - CUSTOMER_SERVICE: Customer service messages.
	//
	// - MARKETING: News, content recommendations, social updates, product promotions, financial updates, lifestyle information, surveys, feature recommendations, and operational activity. This category only identifies the content and does not accelerate message sending. These are collectively referred to as information and marketing messages.
	//
	// For details, see [Notification.category](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117) on the HarmonyOS official website.
	//
	// example:
	//
	// IM
	HarmonyCategory *string `json:"HarmonyCategory,omitempty" xml:"HarmonyCategory,omitempty"`
	// The extended properties of the notification. This parameter does not take effect when the push type PushType is set to MESSAGE.
	//
	// This parameter must be passed in JSON map format. Otherwise, a parsing error occurs.
	//
	// example:
	//
	// {"key1":"value1","api_name":"PushNoticeToAndroidRequest"}
	HarmonyExtParameters *string `json:"HarmonyExtParameters,omitempty" xml:"HarmonyExtParameters,omitempty"`
	// The extra data of the extended notification message.</br>
	//
	// This parameter is valid when you send HarmonyOS extended notification messages.</br>
	//
	// This parameter is conceptually equivalent to the extraData field of a HarmonyOS extended notification message. For the specific definition, see [HarmonyOS ExtensionPayload description](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section161192514234).</br>
	//
	// This parameter is supported in HarmonyOS SDK 1.2.0 and later.
	//
	// example:
	//
	// Sample extra data
	HarmonyExtensionExtraData *string `json:"HarmonyExtensionExtraData,omitempty" xml:"HarmonyExtensionExtraData,omitempty"`
	// Specifies whether to send a HarmonyOS extended notification message when PushType is set to NOTICE. Valid values:
	//
	// - true: Sends an extended notification message.
	//
	// - false: Sends a normal notification. This is the default value.
	//
	// You must apply for permissions on the HarmonyOS side before you can send extended notification messages. For more information, see [Send extended notification messages](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V5/push-send-extend-noti-V5) in the HarmonyOS documentation.</br>
	//
	// This parameter is supported in HarmonyOS SDK 1.2.0 and later.
	//
	// example:
	//
	// true
	HarmonyExtensionPush *bool `json:"HarmonyExtensionPush,omitempty" xml:"HarmonyExtensionPush,omitempty"`
	// The URL of the large icon on the right side of the notification. The URL must use the HTTPS protocol.
	//
	// > Supported image formats include PNG, JPG, JPEG, HEIF, GIF, and BMP. The product of the image width and height must be less than 25,000 pixels.
	//
	// For more information, see [Notification.image](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117) on the HarmonyOS official website.
	//
	// example:
	//
	// https://example.com/xxx.png
	HarmonyImageUrl *string `json:"HarmonyImageUrl,omitempty" xml:"HarmonyImageUrl,omitempty"`
	// The content in the multi-line text style. This field is required when HarmonyRenderStyle is set to MULTI_LINE. A maximum of three content items are supported.
	//
	// example:
	//
	// ["1.content1","2.content2","3.content3"]
	HarmonyInboxContent *string `json:"HarmonyInboxContent,omitempty" xml:"HarmonyInboxContent,omitempty"`
	// The JSON string of the HarmonyOS Live View data structure [LiveViewPayload](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V13/push-scenariozed-api-request-param-V13#section66881469306). For development integration, refer to [HarmonyOS Live View push guide](https://help.aliyun.com/document_detail/2982112.html).
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
	HarmonyLiveViewPayload *string `json:"HarmonyLiveViewPayload,omitempty" xml:"HarmonyLiveViewPayload,omitempty"`
	// The type of the notification channel to use. This parameter is valid only when the Alibaba Cloud proprietary channel is online. Valid values:
	//
	// - SOCIAL_COMMUNICATION: Social communication.
	//
	// - SERVICE_INFORMATION: Service reminders.
	//
	// - CONTENT_INFORMATION: Content information.
	//
	// - CUSTOMER_SERVICE: Customer service messages. This type is used for customer service messages between users and merchants and must be initiated by the user.
	//
	// - OTHER_TYPES: Other types.
	//
	// For more information, see [SlotType](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/js-apis-notificationmanager-V5#slottype) on the HarmonyOS official website.
	//
	// example:
	//
	// SOCIAL_COMMUNICATION
	HarmonyNotificationSlotType *string `json:"HarmonyNotificationSlotType,omitempty" xml:"HarmonyNotificationSlotType,omitempty"`
	// The unique identifier of each message when the notification is displayed. If this parameter is not specified, the push service automatically generates a unique identifier for each message. Different notification messages can have the same notifyId to allow new messages to overwrite old ones.
	//
	// For more information, see [Notification.notifyId](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117) on the HarmonyOS official website.
	//
	// example:
	//
	// 0
	HarmonyNotifyId *int32 `json:"HarmonyNotifyId,omitempty" xml:"HarmonyNotifyId,omitempty"`
	// The receipt ID for the HarmonyOS channel. You can view this receipt ID in the receipt parameter settings of the HarmonyOS channel push operations platform.
	//
	// > If the default receipt configured on the HarmonyOS channel push operations platform is the Alibaba Cloud receipt, you do not need to provide this parameter. Otherwise, we recommend that you configure the default receipt ID for the HarmonyOS channel in the Alibaba Cloud EMAS shift Push console first.
	//
	// For details, see [pushOptions.receiptId](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section418321011212) on the HarmonyOS official website.
	//
	// example:
	//
	// RCPB***DFD5
	HarmonyReceiptId *string `json:"HarmonyReceiptId,omitempty" xml:"HarmonyReceiptId,omitempty"`
	// Specifies whether to use the auxiliary pop-up window feature when a push notification of the message type is sent to an offline device. Default value: false. This parameter takes effect only when PushType is set to MESSAGE.
	//
	// If the message is successfully converted to a notification and pushed, the received notification displays the data specified by the HarmonyRemindTitle and HarmonyRemindBody parameters on the server.
	//
	// example:
	//
	// false
	HarmonyRemind *bool `json:"HarmonyRemind,omitempty" xml:"HarmonyRemind,omitempty"`
	// The HarmonyOS notification content used when a message is converted to a notification. This parameter is valid only when HarmonyRemind is set to true.
	//
	// example:
	//
	// You have a new message. Please check it
	HarmonyRemindBody *string `json:"HarmonyRemindBody,omitempty" xml:"HarmonyRemindBody,omitempty"`
	// The HarmonyOS notification title used when a message is converted to a notification. This parameter is valid only when HarmonyRemind is set to true.
	//
	// example:
	//
	// New message
	HarmonyRemindTitle *string `json:"HarmonyRemindTitle,omitempty" xml:"HarmonyRemindTitle,omitempty"`
	// The style of the notification message. Valid values:
	//
	// - NORMAL: Normal notification. This is the default value.
	//
	// - MULTI_LINE: Multi-line text style.
	//
	// example:
	//
	// NORMAL
	HarmonyRenderStyle *string `json:"HarmonyRenderStyle,omitempty" xml:"HarmonyRenderStyle,omitempty"`
	// Specifies whether the message is a test message. Valid values:
	//
	// - false: Normal message. This is the default value.
	//
	// - true: Test message.
	//
	// For more information, see [pushOptions.testMessage](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section418321011212) on the HarmonyOS official website.
	//
	// example:
	//
	// true
	HarmonyTestMessage *bool `json:"HarmonyTestMessage,omitempty" xml:"HarmonyTestMessage,omitempty"`
	// The URI corresponding to the ability of the built-in page in the app.
	//
	// 	Notice: When HarmonyActionType is set to APP_CUSTOM_PAGE, you must specify at least one of the HarmonyUri and HarmonyAction fields. If multiple abilities exist, specify the action and URI for each ability. The action is preferentially used to find the corresponding built-in page in the app.
	//
	// For more information, see [ClickAction.uri](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section152462191216) on the HarmonyOS official website.
	//
	// example:
	//
	// https://www.example.com:8080/push/example
	HarmonyUri *string `json:"HarmonyUri,omitempty" xml:"HarmonyUri,omitempty"`
	// The custom identifier of the push task. If JobKey is not empty, this field is included in the receipt log. For more information, see [Receipt logs](https://help.aliyun.com/document_detail/434651.html).
	//
	// example:
	//
	// 123
	JobKey *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	// The scheduled sending time. If this parameter is not specified, the message is sent immediately.
	//
	// Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-02-20T00:00:00Z
	PushTime *string `json:"PushTime,omitempty" xml:"PushTime,omitempty"`
	// The push type. Valid values:
	//
	// - MESSAGE: message
	//
	// - NOTICE: notification
	//
	// This parameter is required.
	//
	// example:
	//
	// MESSAGE
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	// The sending channels. Valid values:
	//
	// - accs: Alibaba Cloud proprietary channel
	//
	// - huawei: Huawei channel
	//
	// - honor: Honor channel
	//
	// - xiaomi: Xiaomi channel
	//
	// - oppo: OPPO channel
	//
	// - vivo: vivo channel
	//
	// - meizu: Meizu channel
	//
	// - gcm: Google GCM channel (legacy HTTP)
	//
	// - fcm: Google Firebase channel (HTTP v1 API)
	//
	// - apns: APNs channel
	//
	// - harmony: HarmonyOS channel
	//
	// > - If this parameter is not configured, all channels are available.
	//
	// > - If this parameter is configured, only the specified channels are used.
	//
	// > - If the configured channels conflict with the sending policy (for example, iOS notifications can only be sent through the APNs channel, but apns is not included in this parameter), the message is not sent.
	//
	// > - If gcm is configured, both Google GCM and FCM channels can be used. If fcm is configured, only the Google FCM channel can be used.
	//
	// example:
	//
	// accs,huawei,xiaomi
	SendChannels *string `json:"SendChannels,omitempty" xml:"SendChannels,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- This parameter is deprecated.
	//
	// example:
	//
	// 0
	SendSpeed *int32 `json:"SendSpeed,omitempty" xml:"SendSpeed,omitempty"`
	// Specifies whether to save offline messages or notifications. Default value: false.
	//
	// If this parameter is set to true and the user is offline during the push, the message or notification is sent again when the user comes online within the expiration time specified by ExpireTime. The default value of ExpireTime is 72 hours. iOS notifications are sent through the APNs channel and are not affected by the StoreOffline parameter.
	//
	// example:
	//
	// true
	StoreOffline *bool `json:"StoreOffline,omitempty" xml:"StoreOffline,omitempty"`
	// The push target. Valid values:
	//
	// - DEVICE: device
	//
	// - ACCOUNT: account
	//
	// - ALIAS: alias
	//
	// This parameter is required.
	//
	// example:
	//
	// DEVICE
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The target value, which is specified based on the Target parameter. Separate multiple values with commas. If the number of values exceeds the limit, you must push the message in multiple batches.
	//
	// - If Target is set to DEVICE, the value is in the format of `deviceid1,deviceid2`. A maximum of 1,000 devices are supported.
	//
	// - If Target is set to ACCOUNT, the value is in the format of `account1,account2`. A maximum of 1,000 accounts are supported.
	//
	// - If Target is set to ALIAS, the value is in the format of `alias1,alias2`. A maximum of 1,000 aliases are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// deviceid1,deviceid2
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
	// The title of the notification or message. Maximum length: 200 bytes.
	//
	// This parameter is required for Android and HarmonyOS pushes. This parameter is optional for iOS notification pushes. If you specify this parameter for iOS:
	//
	// 	- For iOS 10 and later, the title is displayed in the notification.
	//
	// 	- For iOS 8.2 to iOS 9.x, the title replaces the application name in the notification.
	//
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Specifies whether to automatically truncate excessively long titles and content.
	//
	// Note: This parameter applies only to vendor channels that explicitly limit the length of titles and content. It does not apply to channels such as APNs, Huawei, and Honor, which do not limit the length of titles and content but only limit the total request body size.
	//
	// example:
	//
	// false
	Trim *bool `json:"Trim,omitempty" xml:"Trim,omitempty"`
	// To send iOS notifications through the APNs center, you must specify the corresponding environment context. Valid values:
	//
	// - DEV: The development environment. This value applies to applications that are directly installed and debugged by using Xcode.
	//
	// - PRODUCT: The production environment. This value applies to applications distributed through the App Store, TestFlight, Ad Hoc, and enterprise distribution.
	//
	// example:
	//
	// DEV
	IOSApnsEnv *string `json:"iOSApnsEnv,omitempty" xml:"iOSApnsEnv,omitempty"`
	// The badge number displayed in the upper-right corner of the iOS application icon.
	//
	// > If iOSBadgeAutoIncrement is set to true, this parameter must be left empty.
	//
	// example:
	//
	// 0
	IOSBadge *int32 `json:"iOSBadge,omitempty" xml:"iOSBadge,omitempty"`
	// Specifies whether to enable the automatic badge increment feature. Default value: false.
	//
	// > When this parameter is set to true, iOSBadge must be left empty.
	//
	// The automatic badge increment feature maintains the badge count for each device on the push server. This feature requires SDK V1.9.5 or later, and users must actively synchronize the badge number to the server.
	//
	// example:
	//
	// true
	IOSBadgeAutoIncrement *bool `json:"iOSBadgeAutoIncrement,omitempty" xml:"iOSBadgeAutoIncrement,omitempty"`
	// The extension property of the iOS notification.
	//
	// For iOS 10 and later, you can specify the resource URL for rich media push notifications in this parameter: `{"attachment": "https://www.example.com/notification_pic.png"}`. This parameter must be passed in JSON map format. Otherwise, a parse error occurs.
	//
	// example:
	//
	// {"attachment": "https://www.example.com/notification_pic.png"}
	IOSExtParameters *string `json:"iOSExtParameters,omitempty" xml:"iOSExtParameters,omitempty"`
	// The interruption level. Valid values:
	//
	// - passive: The system adds the notification to the notification list without turning on the screen or playing a sound.
	//
	// - active: The system displays the notification immediately, turns on the screen, and can play a sound.
	//
	// - time-sensitive: The system displays the notification immediately, turns on the screen, and can play a sound, but does not break through the system notification controls.
	//
	// - critical: The system displays the notification immediately, turns on the screen, and plays a sound bypassing the silent switch.
	//
	// example:
	//
	// active
	IOSInterruptionLevel *string `json:"iOSInterruptionLevel,omitempty" xml:"iOSInterruptionLevel,omitempty"`
	// The JSON string that contains the static pass-through parameters for Dynamic Island pushes. Static parameter includes static user-defined information, such as product IDs and order information.
	//
	// > Static parameter is required when iOSLiveActivityEvent is set to start.
	//
	// example:
	//
	// {"orderId": "12345", "product": "Shoes"}
	IOSLiveActivityAttributes *string `json:"iOSLiveActivityAttributes,omitempty" xml:"iOSLiveActivityAttributes,omitempty"`
	// The type of the Live Activity to be started.
	//
	// > This parameter is required when iOSLiveActivityEvent is set to start.
	//
	// example:
	//
	// OrderActivityAttributes
	IOSLiveActivityAttributesType *string `json:"iOSLiveActivityAttributesType,omitempty" xml:"iOSLiveActivityAttributesType,omitempty"`
	// The pass-through dynamic parameter for Dynamic Island pushes. This parameter includes real-time update information, such as price and inventory changes.
	//
	// example:
	//
	// {"status": "delivered", "estimatedArrival": "2023-12-31T12:00:00Z"}
	IOSLiveActivityContentState *string `json:"iOSLiveActivityContentState,omitempty" xml:"iOSLiveActivityContentState,omitempty"`
	// The time until which the ended Live Activity remains on the Lock Screen. The maximum retention period is 4 hours.
	//
	// example:
	//
	// 1743131967
	IOSLiveActivityDismissalDate *int64 `json:"iOSLiveActivityDismissalDate,omitempty" xml:"iOSLiveActivityDismissalDate,omitempty"`
	// The event to start, update, or end a Live Activity. Valid values: start, update, and end.
	//
	// example:
	//
	// start
	IOSLiveActivityEvent *string `json:"iOSLiveActivityEvent,omitempty" xml:"iOSLiveActivityEvent,omitempty"`
	// The Live Activity ID reported by the device to the user server. This is the unique identifier of the Live Activity.
	//
	// example:
	//
	// 66B94673-B32E-4CA7-863C-3E523054FD46
	IOSLiveActivityId *string `json:"iOSLiveActivityId,omitempty" xml:"iOSLiveActivityId,omitempty"`
	// The second-level timestamp that marks the expiration time of the activity content.
	//
	// example:
	//
	// 1743131967
	IOSLiveActivityStaleDate *int64 `json:"iOSLiveActivityStaleDate,omitempty" xml:"iOSLiveActivityStaleDate,omitempty"`
	// The sound of the iOS notification. Specify the name of the audio file stored in the app bundle or the sandbox Library/Sounds directory. For more information, see How to set the notification sound for iOS pushes.
	//
	// If this parameter is set to an empty string (""), the notification is silent. If this parameter is not specified, the default value "default" is used, which indicates the system prompt sound.
	//
	// example:
	//
	// ””
	IOSMusic *string `json:"iOSMusic,omitempty" xml:"iOSMusic,omitempty"`
	// The iOS notification processing extension flag (iOS 10 and later). If this parameter is set to true, the notification pushed by APNs can be processed by the extension before it is displayed. This parameter must be set to true for silent notifications.
	//
	// example:
	//
	// true
	IOSMutableContent *bool `json:"iOSMutableContent,omitempty" xml:"iOSMutableContent,omitempty"`
	// The category of the iOS notification (iOS 10 and later).
	//
	// example:
	//
	// ios
	IOSNotificationCategory *string `json:"iOSNotificationCategory,omitempty" xml:"iOSNotificationCategory,omitempty"`
	// The collapse ID. If a device receives messages with the same CollapseId, the messages are merged into one. If the device is offline and continuously receives messages with the same CollapseId, only one notification is displayed in the notification bar. This parameter is supported in iOS 10 and later.
	//
	// example:
	//
	// ZD2011
	IOSNotificationCollapseId *string `json:"iOSNotificationCollapseId,omitempty" xml:"iOSNotificationCollapseId,omitempty"`
	// The thread ID used to group iOS remote notifications and mark the identifier of the collapsed group. This parameter is supported only in iOS 12.0 and later.
	//
	// example:
	//
	// abc
	IOSNotificationThreadId *string `json:"iOSNotificationThreadId,omitempty" xml:"iOSNotificationThreadId,omitempty"`
	// The relevance score for summary highlighting. Valid values: a floating-point number in the range of [0,1].
	//
	// example:
	//
	// 0.01
	IOSRelevanceScore *float64 `json:"iOSRelevanceScore,omitempty" xml:"iOSRelevanceScore,omitempty"`
	// If the device is offline during a message push (that is, the persistent connection to the Mobile Push server is disconnected), the message is delivered once as a notification through the Apple APNs channel.
	//
	// > Offline message-to-notification conversion applies only to the production environment.
	//
	// example:
	//
	// true
	IOSRemind *bool `json:"iOSRemind,omitempty" xml:"iOSRemind,omitempty"`
	// The iOS notification content used when an iOS message is converted to a notification. This parameter is valid only when iOSApnsEnv is set to PRODUCT and iOSRemind is set to true.
	//
	// example:
	//
	// iOS notification body
	IOSRemindBody *string `json:"iOSRemindBody,omitempty" xml:"iOSRemindBody,omitempty"`
	// Specifies whether to enable iOS silent notifications.
	//
	// example:
	//
	// true
	IOSSilentNotification *bool `json:"iOSSilentNotification,omitempty" xml:"iOSSilentNotification,omitempty"`
	// The subtitle of the iOS notification (iOS 10 and later).
	//
	// example:
	//
	// subtitle
	IOSSubtitle *string `json:"iOSSubtitle,omitempty" xml:"iOSSubtitle,omitempty"`
}

func (s MassPushRequestPushTask) String() string {
	return dara.Prettify(s)
}

func (s MassPushRequestPushTask) GoString() string {
	return s.String()
}

func (s *MassPushRequestPushTask) GetAndroidActivity() *string {
	return s.AndroidActivity
}

func (s *MassPushRequestPushTask) GetAndroidBadgeAddNum() *int32 {
	return s.AndroidBadgeAddNum
}

func (s *MassPushRequestPushTask) GetAndroidBadgeClass() *string {
	return s.AndroidBadgeClass
}

func (s *MassPushRequestPushTask) GetAndroidBadgeSetNum() *int32 {
	return s.AndroidBadgeSetNum
}

func (s *MassPushRequestPushTask) GetAndroidBigBody() *string {
	return s.AndroidBigBody
}

func (s *MassPushRequestPushTask) GetAndroidBigPictureUrl() *string {
	return s.AndroidBigPictureUrl
}

func (s *MassPushRequestPushTask) GetAndroidBigTitle() *string {
	return s.AndroidBigTitle
}

func (s *MassPushRequestPushTask) GetAndroidExtParameters() *string {
	return s.AndroidExtParameters
}

func (s *MassPushRequestPushTask) GetAndroidHonorTargetUserType() *int32 {
	return s.AndroidHonorTargetUserType
}

func (s *MassPushRequestPushTask) GetAndroidHuaweiBusinessType() *int32 {
	return s.AndroidHuaweiBusinessType
}

func (s *MassPushRequestPushTask) GetAndroidHuaweiLiveNotificationPayload() *string {
	return s.AndroidHuaweiLiveNotificationPayload
}

func (s *MassPushRequestPushTask) GetAndroidHuaweiReceiptId() *string {
	return s.AndroidHuaweiReceiptId
}

func (s *MassPushRequestPushTask) GetAndroidHuaweiTargetUserType() *int32 {
	return s.AndroidHuaweiTargetUserType
}

func (s *MassPushRequestPushTask) GetAndroidImageUrl() *string {
	return s.AndroidImageUrl
}

func (s *MassPushRequestPushTask) GetAndroidInboxBody() *string {
	return s.AndroidInboxBody
}

func (s *MassPushRequestPushTask) GetAndroidMeizuNoticeMsgType() *int32 {
	return s.AndroidMeizuNoticeMsgType
}

func (s *MassPushRequestPushTask) GetAndroidMessageHuaweiCategory() *string {
	return s.AndroidMessageHuaweiCategory
}

func (s *MassPushRequestPushTask) GetAndroidMessageHuaweiUrgency() *string {
	return s.AndroidMessageHuaweiUrgency
}

func (s *MassPushRequestPushTask) GetAndroidMessageOppoCategory() *string {
	return s.AndroidMessageOppoCategory
}

func (s *MassPushRequestPushTask) GetAndroidMessageOppoNotifyLevel() *int32 {
	return s.AndroidMessageOppoNotifyLevel
}

func (s *MassPushRequestPushTask) GetAndroidMessageVivoCategory() *string {
	return s.AndroidMessageVivoCategory
}

func (s *MassPushRequestPushTask) GetAndroidMusic() *string {
	return s.AndroidMusic
}

func (s *MassPushRequestPushTask) GetAndroidNotificationBarPriority() *int32 {
	return s.AndroidNotificationBarPriority
}

func (s *MassPushRequestPushTask) GetAndroidNotificationBarType() *int32 {
	return s.AndroidNotificationBarType
}

func (s *MassPushRequestPushTask) GetAndroidNotificationChannel() *string {
	return s.AndroidNotificationChannel
}

func (s *MassPushRequestPushTask) GetAndroidNotificationGroup() *string {
	return s.AndroidNotificationGroup
}

func (s *MassPushRequestPushTask) GetAndroidNotificationHonorChannel() *string {
	return s.AndroidNotificationHonorChannel
}

func (s *MassPushRequestPushTask) GetAndroidNotificationHuaweiChannel() *string {
	return s.AndroidNotificationHuaweiChannel
}

func (s *MassPushRequestPushTask) GetAndroidNotificationNotifyId() *int32 {
	return s.AndroidNotificationNotifyId
}

func (s *MassPushRequestPushTask) GetAndroidNotificationThreadId() *string {
	return s.AndroidNotificationThreadId
}

func (s *MassPushRequestPushTask) GetAndroidNotificationVivoChannel() *string {
	return s.AndroidNotificationVivoChannel
}

func (s *MassPushRequestPushTask) GetAndroidNotificationXiaomiChannel() *string {
	return s.AndroidNotificationXiaomiChannel
}

func (s *MassPushRequestPushTask) GetAndroidNotifyType() *string {
	return s.AndroidNotifyType
}

func (s *MassPushRequestPushTask) GetAndroidOpenType() *string {
	return s.AndroidOpenType
}

func (s *MassPushRequestPushTask) GetAndroidOpenUrl() *string {
	return s.AndroidOpenUrl
}

func (s *MassPushRequestPushTask) GetAndroidOppoDeleteIntentData() *string {
	return s.AndroidOppoDeleteIntentData
}

func (s *MassPushRequestPushTask) GetAndroidOppoIntelligentIntent() *string {
	return s.AndroidOppoIntelligentIntent
}

func (s *MassPushRequestPushTask) GetAndroidOppoIntentEnv() *int32 {
	return s.AndroidOppoIntentEnv
}

func (s *MassPushRequestPushTask) GetAndroidOppoPrivateContentParameters() map[string]*string {
	return s.AndroidOppoPrivateContentParameters
}

func (s *MassPushRequestPushTask) GetAndroidOppoPrivateMsgTemplateId() *string {
	return s.AndroidOppoPrivateMsgTemplateId
}

func (s *MassPushRequestPushTask) GetAndroidOppoPrivateTitleParameters() map[string]*string {
	return s.AndroidOppoPrivateTitleParameters
}

func (s *MassPushRequestPushTask) GetAndroidPopupActivity() *string {
	return s.AndroidPopupActivity
}

func (s *MassPushRequestPushTask) GetAndroidPopupBody() *string {
	return s.AndroidPopupBody
}

func (s *MassPushRequestPushTask) GetAndroidPopupTitle() *string {
	return s.AndroidPopupTitle
}

func (s *MassPushRequestPushTask) GetAndroidRemind() *bool {
	return s.AndroidRemind
}

func (s *MassPushRequestPushTask) GetAndroidRenderStyle() *string {
	return s.AndroidRenderStyle
}

func (s *MassPushRequestPushTask) GetAndroidTargetUserType() *int32 {
	return s.AndroidTargetUserType
}

func (s *MassPushRequestPushTask) GetAndroidVivoLiveMessage() *string {
	return s.AndroidVivoLiveMessage
}

func (s *MassPushRequestPushTask) GetAndroidVivoPushMode() *int32 {
	return s.AndroidVivoPushMode
}

func (s *MassPushRequestPushTask) GetAndroidVivoReceiptId() *string {
	return s.AndroidVivoReceiptId
}

func (s *MassPushRequestPushTask) GetAndroidVivoTemplateId() *string {
	return s.AndroidVivoTemplateId
}

func (s *MassPushRequestPushTask) GetAndroidVivoTemplateParams() *string {
	return s.AndroidVivoTemplateParams
}

func (s *MassPushRequestPushTask) GetAndroidXiaoMiActivity() *string {
	return s.AndroidXiaoMiActivity
}

func (s *MassPushRequestPushTask) GetAndroidXiaoMiNotifyBody() *string {
	return s.AndroidXiaoMiNotifyBody
}

func (s *MassPushRequestPushTask) GetAndroidXiaoMiNotifyTitle() *string {
	return s.AndroidXiaoMiNotifyTitle
}

func (s *MassPushRequestPushTask) GetAndroidXiaomiBigPictureUrl() *string {
	return s.AndroidXiaomiBigPictureUrl
}

func (s *MassPushRequestPushTask) GetAndroidXiaomiFocusParam() *string {
	return s.AndroidXiaomiFocusParam
}

func (s *MassPushRequestPushTask) GetAndroidXiaomiFocusPics() *string {
	return s.AndroidXiaomiFocusPics
}

func (s *MassPushRequestPushTask) GetAndroidXiaomiImageUrl() *string {
	return s.AndroidXiaomiImageUrl
}

func (s *MassPushRequestPushTask) GetAndroidXiaomiTemplateId() *string {
	return s.AndroidXiaomiTemplateId
}

func (s *MassPushRequestPushTask) GetAndroidXiaomiTemplateParams() *string {
	return s.AndroidXiaomiTemplateParams
}

func (s *MassPushRequestPushTask) GetBody() *string {
	return s.Body
}

func (s *MassPushRequestPushTask) GetDeviceType() *string {
	return s.DeviceType
}

func (s *MassPushRequestPushTask) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *MassPushRequestPushTask) GetHarmonyAction() *string {
	return s.HarmonyAction
}

func (s *MassPushRequestPushTask) GetHarmonyActionType() *string {
	return s.HarmonyActionType
}

func (s *MassPushRequestPushTask) GetHarmonyBadgeAddNum() *int32 {
	return s.HarmonyBadgeAddNum
}

func (s *MassPushRequestPushTask) GetHarmonyBadgeSetNum() *int32 {
	return s.HarmonyBadgeSetNum
}

func (s *MassPushRequestPushTask) GetHarmonyCategory() *string {
	return s.HarmonyCategory
}

func (s *MassPushRequestPushTask) GetHarmonyExtParameters() *string {
	return s.HarmonyExtParameters
}

func (s *MassPushRequestPushTask) GetHarmonyExtensionExtraData() *string {
	return s.HarmonyExtensionExtraData
}

func (s *MassPushRequestPushTask) GetHarmonyExtensionPush() *bool {
	return s.HarmonyExtensionPush
}

func (s *MassPushRequestPushTask) GetHarmonyImageUrl() *string {
	return s.HarmonyImageUrl
}

func (s *MassPushRequestPushTask) GetHarmonyInboxContent() *string {
	return s.HarmonyInboxContent
}

func (s *MassPushRequestPushTask) GetHarmonyLiveViewPayload() *string {
	return s.HarmonyLiveViewPayload
}

func (s *MassPushRequestPushTask) GetHarmonyNotificationSlotType() *string {
	return s.HarmonyNotificationSlotType
}

func (s *MassPushRequestPushTask) GetHarmonyNotifyId() *int32 {
	return s.HarmonyNotifyId
}

func (s *MassPushRequestPushTask) GetHarmonyReceiptId() *string {
	return s.HarmonyReceiptId
}

func (s *MassPushRequestPushTask) GetHarmonyRemind() *bool {
	return s.HarmonyRemind
}

func (s *MassPushRequestPushTask) GetHarmonyRemindBody() *string {
	return s.HarmonyRemindBody
}

func (s *MassPushRequestPushTask) GetHarmonyRemindTitle() *string {
	return s.HarmonyRemindTitle
}

func (s *MassPushRequestPushTask) GetHarmonyRenderStyle() *string {
	return s.HarmonyRenderStyle
}

func (s *MassPushRequestPushTask) GetHarmonyTestMessage() *bool {
	return s.HarmonyTestMessage
}

func (s *MassPushRequestPushTask) GetHarmonyUri() *string {
	return s.HarmonyUri
}

func (s *MassPushRequestPushTask) GetJobKey() *string {
	return s.JobKey
}

func (s *MassPushRequestPushTask) GetPushTime() *string {
	return s.PushTime
}

func (s *MassPushRequestPushTask) GetPushType() *string {
	return s.PushType
}

func (s *MassPushRequestPushTask) GetSendChannels() *string {
	return s.SendChannels
}

func (s *MassPushRequestPushTask) GetSendSpeed() *int32 {
	return s.SendSpeed
}

func (s *MassPushRequestPushTask) GetStoreOffline() *bool {
	return s.StoreOffline
}

func (s *MassPushRequestPushTask) GetTarget() *string {
	return s.Target
}

func (s *MassPushRequestPushTask) GetTargetValue() *string {
	return s.TargetValue
}

func (s *MassPushRequestPushTask) GetTitle() *string {
	return s.Title
}

func (s *MassPushRequestPushTask) GetTrim() *bool {
	return s.Trim
}

func (s *MassPushRequestPushTask) GetIOSApnsEnv() *string {
	return s.IOSApnsEnv
}

func (s *MassPushRequestPushTask) GetIOSBadge() *int32 {
	return s.IOSBadge
}

func (s *MassPushRequestPushTask) GetIOSBadgeAutoIncrement() *bool {
	return s.IOSBadgeAutoIncrement
}

func (s *MassPushRequestPushTask) GetIOSExtParameters() *string {
	return s.IOSExtParameters
}

func (s *MassPushRequestPushTask) GetIOSInterruptionLevel() *string {
	return s.IOSInterruptionLevel
}

func (s *MassPushRequestPushTask) GetIOSLiveActivityAttributes() *string {
	return s.IOSLiveActivityAttributes
}

func (s *MassPushRequestPushTask) GetIOSLiveActivityAttributesType() *string {
	return s.IOSLiveActivityAttributesType
}

func (s *MassPushRequestPushTask) GetIOSLiveActivityContentState() *string {
	return s.IOSLiveActivityContentState
}

func (s *MassPushRequestPushTask) GetIOSLiveActivityDismissalDate() *int64 {
	return s.IOSLiveActivityDismissalDate
}

func (s *MassPushRequestPushTask) GetIOSLiveActivityEvent() *string {
	return s.IOSLiveActivityEvent
}

func (s *MassPushRequestPushTask) GetIOSLiveActivityId() *string {
	return s.IOSLiveActivityId
}

func (s *MassPushRequestPushTask) GetIOSLiveActivityStaleDate() *int64 {
	return s.IOSLiveActivityStaleDate
}

func (s *MassPushRequestPushTask) GetIOSMusic() *string {
	return s.IOSMusic
}

func (s *MassPushRequestPushTask) GetIOSMutableContent() *bool {
	return s.IOSMutableContent
}

func (s *MassPushRequestPushTask) GetIOSNotificationCategory() *string {
	return s.IOSNotificationCategory
}

func (s *MassPushRequestPushTask) GetIOSNotificationCollapseId() *string {
	return s.IOSNotificationCollapseId
}

func (s *MassPushRequestPushTask) GetIOSNotificationThreadId() *string {
	return s.IOSNotificationThreadId
}

func (s *MassPushRequestPushTask) GetIOSRelevanceScore() *float64 {
	return s.IOSRelevanceScore
}

func (s *MassPushRequestPushTask) GetIOSRemind() *bool {
	return s.IOSRemind
}

func (s *MassPushRequestPushTask) GetIOSRemindBody() *string {
	return s.IOSRemindBody
}

func (s *MassPushRequestPushTask) GetIOSSilentNotification() *bool {
	return s.IOSSilentNotification
}

func (s *MassPushRequestPushTask) GetIOSSubtitle() *string {
	return s.IOSSubtitle
}

func (s *MassPushRequestPushTask) SetAndroidActivity(v string) *MassPushRequestPushTask {
	s.AndroidActivity = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidBadgeAddNum(v int32) *MassPushRequestPushTask {
	s.AndroidBadgeAddNum = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidBadgeClass(v string) *MassPushRequestPushTask {
	s.AndroidBadgeClass = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidBadgeSetNum(v int32) *MassPushRequestPushTask {
	s.AndroidBadgeSetNum = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidBigBody(v string) *MassPushRequestPushTask {
	s.AndroidBigBody = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidBigPictureUrl(v string) *MassPushRequestPushTask {
	s.AndroidBigPictureUrl = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidBigTitle(v string) *MassPushRequestPushTask {
	s.AndroidBigTitle = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidExtParameters(v string) *MassPushRequestPushTask {
	s.AndroidExtParameters = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidHonorTargetUserType(v int32) *MassPushRequestPushTask {
	s.AndroidHonorTargetUserType = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidHuaweiBusinessType(v int32) *MassPushRequestPushTask {
	s.AndroidHuaweiBusinessType = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidHuaweiLiveNotificationPayload(v string) *MassPushRequestPushTask {
	s.AndroidHuaweiLiveNotificationPayload = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidHuaweiReceiptId(v string) *MassPushRequestPushTask {
	s.AndroidHuaweiReceiptId = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidHuaweiTargetUserType(v int32) *MassPushRequestPushTask {
	s.AndroidHuaweiTargetUserType = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidImageUrl(v string) *MassPushRequestPushTask {
	s.AndroidImageUrl = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidInboxBody(v string) *MassPushRequestPushTask {
	s.AndroidInboxBody = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidMeizuNoticeMsgType(v int32) *MassPushRequestPushTask {
	s.AndroidMeizuNoticeMsgType = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidMessageHuaweiCategory(v string) *MassPushRequestPushTask {
	s.AndroidMessageHuaweiCategory = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidMessageHuaweiUrgency(v string) *MassPushRequestPushTask {
	s.AndroidMessageHuaweiUrgency = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidMessageOppoCategory(v string) *MassPushRequestPushTask {
	s.AndroidMessageOppoCategory = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidMessageOppoNotifyLevel(v int32) *MassPushRequestPushTask {
	s.AndroidMessageOppoNotifyLevel = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidMessageVivoCategory(v string) *MassPushRequestPushTask {
	s.AndroidMessageVivoCategory = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidMusic(v string) *MassPushRequestPushTask {
	s.AndroidMusic = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidNotificationBarPriority(v int32) *MassPushRequestPushTask {
	s.AndroidNotificationBarPriority = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidNotificationBarType(v int32) *MassPushRequestPushTask {
	s.AndroidNotificationBarType = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidNotificationChannel(v string) *MassPushRequestPushTask {
	s.AndroidNotificationChannel = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidNotificationGroup(v string) *MassPushRequestPushTask {
	s.AndroidNotificationGroup = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidNotificationHonorChannel(v string) *MassPushRequestPushTask {
	s.AndroidNotificationHonorChannel = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidNotificationHuaweiChannel(v string) *MassPushRequestPushTask {
	s.AndroidNotificationHuaweiChannel = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidNotificationNotifyId(v int32) *MassPushRequestPushTask {
	s.AndroidNotificationNotifyId = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidNotificationThreadId(v string) *MassPushRequestPushTask {
	s.AndroidNotificationThreadId = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidNotificationVivoChannel(v string) *MassPushRequestPushTask {
	s.AndroidNotificationVivoChannel = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidNotificationXiaomiChannel(v string) *MassPushRequestPushTask {
	s.AndroidNotificationXiaomiChannel = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidNotifyType(v string) *MassPushRequestPushTask {
	s.AndroidNotifyType = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidOpenType(v string) *MassPushRequestPushTask {
	s.AndroidOpenType = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidOpenUrl(v string) *MassPushRequestPushTask {
	s.AndroidOpenUrl = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidOppoDeleteIntentData(v string) *MassPushRequestPushTask {
	s.AndroidOppoDeleteIntentData = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidOppoIntelligentIntent(v string) *MassPushRequestPushTask {
	s.AndroidOppoIntelligentIntent = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidOppoIntentEnv(v int32) *MassPushRequestPushTask {
	s.AndroidOppoIntentEnv = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidOppoPrivateContentParameters(v map[string]*string) *MassPushRequestPushTask {
	s.AndroidOppoPrivateContentParameters = v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidOppoPrivateMsgTemplateId(v string) *MassPushRequestPushTask {
	s.AndroidOppoPrivateMsgTemplateId = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidOppoPrivateTitleParameters(v map[string]*string) *MassPushRequestPushTask {
	s.AndroidOppoPrivateTitleParameters = v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidPopupActivity(v string) *MassPushRequestPushTask {
	s.AndroidPopupActivity = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidPopupBody(v string) *MassPushRequestPushTask {
	s.AndroidPopupBody = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidPopupTitle(v string) *MassPushRequestPushTask {
	s.AndroidPopupTitle = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidRemind(v bool) *MassPushRequestPushTask {
	s.AndroidRemind = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidRenderStyle(v string) *MassPushRequestPushTask {
	s.AndroidRenderStyle = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidTargetUserType(v int32) *MassPushRequestPushTask {
	s.AndroidTargetUserType = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidVivoLiveMessage(v string) *MassPushRequestPushTask {
	s.AndroidVivoLiveMessage = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidVivoPushMode(v int32) *MassPushRequestPushTask {
	s.AndroidVivoPushMode = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidVivoReceiptId(v string) *MassPushRequestPushTask {
	s.AndroidVivoReceiptId = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidVivoTemplateId(v string) *MassPushRequestPushTask {
	s.AndroidVivoTemplateId = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidVivoTemplateParams(v string) *MassPushRequestPushTask {
	s.AndroidVivoTemplateParams = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidXiaoMiActivity(v string) *MassPushRequestPushTask {
	s.AndroidXiaoMiActivity = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidXiaoMiNotifyBody(v string) *MassPushRequestPushTask {
	s.AndroidXiaoMiNotifyBody = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidXiaoMiNotifyTitle(v string) *MassPushRequestPushTask {
	s.AndroidXiaoMiNotifyTitle = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidXiaomiBigPictureUrl(v string) *MassPushRequestPushTask {
	s.AndroidXiaomiBigPictureUrl = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidXiaomiFocusParam(v string) *MassPushRequestPushTask {
	s.AndroidXiaomiFocusParam = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidXiaomiFocusPics(v string) *MassPushRequestPushTask {
	s.AndroidXiaomiFocusPics = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidXiaomiImageUrl(v string) *MassPushRequestPushTask {
	s.AndroidXiaomiImageUrl = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidXiaomiTemplateId(v string) *MassPushRequestPushTask {
	s.AndroidXiaomiTemplateId = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidXiaomiTemplateParams(v string) *MassPushRequestPushTask {
	s.AndroidXiaomiTemplateParams = &v
	return s
}

func (s *MassPushRequestPushTask) SetBody(v string) *MassPushRequestPushTask {
	s.Body = &v
	return s
}

func (s *MassPushRequestPushTask) SetDeviceType(v string) *MassPushRequestPushTask {
	s.DeviceType = &v
	return s
}

func (s *MassPushRequestPushTask) SetExpireTime(v string) *MassPushRequestPushTask {
	s.ExpireTime = &v
	return s
}

func (s *MassPushRequestPushTask) SetHarmonyAction(v string) *MassPushRequestPushTask {
	s.HarmonyAction = &v
	return s
}

func (s *MassPushRequestPushTask) SetHarmonyActionType(v string) *MassPushRequestPushTask {
	s.HarmonyActionType = &v
	return s
}

func (s *MassPushRequestPushTask) SetHarmonyBadgeAddNum(v int32) *MassPushRequestPushTask {
	s.HarmonyBadgeAddNum = &v
	return s
}

func (s *MassPushRequestPushTask) SetHarmonyBadgeSetNum(v int32) *MassPushRequestPushTask {
	s.HarmonyBadgeSetNum = &v
	return s
}

func (s *MassPushRequestPushTask) SetHarmonyCategory(v string) *MassPushRequestPushTask {
	s.HarmonyCategory = &v
	return s
}

func (s *MassPushRequestPushTask) SetHarmonyExtParameters(v string) *MassPushRequestPushTask {
	s.HarmonyExtParameters = &v
	return s
}

func (s *MassPushRequestPushTask) SetHarmonyExtensionExtraData(v string) *MassPushRequestPushTask {
	s.HarmonyExtensionExtraData = &v
	return s
}

func (s *MassPushRequestPushTask) SetHarmonyExtensionPush(v bool) *MassPushRequestPushTask {
	s.HarmonyExtensionPush = &v
	return s
}

func (s *MassPushRequestPushTask) SetHarmonyImageUrl(v string) *MassPushRequestPushTask {
	s.HarmonyImageUrl = &v
	return s
}

func (s *MassPushRequestPushTask) SetHarmonyInboxContent(v string) *MassPushRequestPushTask {
	s.HarmonyInboxContent = &v
	return s
}

func (s *MassPushRequestPushTask) SetHarmonyLiveViewPayload(v string) *MassPushRequestPushTask {
	s.HarmonyLiveViewPayload = &v
	return s
}

func (s *MassPushRequestPushTask) SetHarmonyNotificationSlotType(v string) *MassPushRequestPushTask {
	s.HarmonyNotificationSlotType = &v
	return s
}

func (s *MassPushRequestPushTask) SetHarmonyNotifyId(v int32) *MassPushRequestPushTask {
	s.HarmonyNotifyId = &v
	return s
}

func (s *MassPushRequestPushTask) SetHarmonyReceiptId(v string) *MassPushRequestPushTask {
	s.HarmonyReceiptId = &v
	return s
}

func (s *MassPushRequestPushTask) SetHarmonyRemind(v bool) *MassPushRequestPushTask {
	s.HarmonyRemind = &v
	return s
}

func (s *MassPushRequestPushTask) SetHarmonyRemindBody(v string) *MassPushRequestPushTask {
	s.HarmonyRemindBody = &v
	return s
}

func (s *MassPushRequestPushTask) SetHarmonyRemindTitle(v string) *MassPushRequestPushTask {
	s.HarmonyRemindTitle = &v
	return s
}

func (s *MassPushRequestPushTask) SetHarmonyRenderStyle(v string) *MassPushRequestPushTask {
	s.HarmonyRenderStyle = &v
	return s
}

func (s *MassPushRequestPushTask) SetHarmonyTestMessage(v bool) *MassPushRequestPushTask {
	s.HarmonyTestMessage = &v
	return s
}

func (s *MassPushRequestPushTask) SetHarmonyUri(v string) *MassPushRequestPushTask {
	s.HarmonyUri = &v
	return s
}

func (s *MassPushRequestPushTask) SetJobKey(v string) *MassPushRequestPushTask {
	s.JobKey = &v
	return s
}

func (s *MassPushRequestPushTask) SetPushTime(v string) *MassPushRequestPushTask {
	s.PushTime = &v
	return s
}

func (s *MassPushRequestPushTask) SetPushType(v string) *MassPushRequestPushTask {
	s.PushType = &v
	return s
}

func (s *MassPushRequestPushTask) SetSendChannels(v string) *MassPushRequestPushTask {
	s.SendChannels = &v
	return s
}

func (s *MassPushRequestPushTask) SetSendSpeed(v int32) *MassPushRequestPushTask {
	s.SendSpeed = &v
	return s
}

func (s *MassPushRequestPushTask) SetStoreOffline(v bool) *MassPushRequestPushTask {
	s.StoreOffline = &v
	return s
}

func (s *MassPushRequestPushTask) SetTarget(v string) *MassPushRequestPushTask {
	s.Target = &v
	return s
}

func (s *MassPushRequestPushTask) SetTargetValue(v string) *MassPushRequestPushTask {
	s.TargetValue = &v
	return s
}

func (s *MassPushRequestPushTask) SetTitle(v string) *MassPushRequestPushTask {
	s.Title = &v
	return s
}

func (s *MassPushRequestPushTask) SetTrim(v bool) *MassPushRequestPushTask {
	s.Trim = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSApnsEnv(v string) *MassPushRequestPushTask {
	s.IOSApnsEnv = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSBadge(v int32) *MassPushRequestPushTask {
	s.IOSBadge = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSBadgeAutoIncrement(v bool) *MassPushRequestPushTask {
	s.IOSBadgeAutoIncrement = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSExtParameters(v string) *MassPushRequestPushTask {
	s.IOSExtParameters = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSInterruptionLevel(v string) *MassPushRequestPushTask {
	s.IOSInterruptionLevel = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSLiveActivityAttributes(v string) *MassPushRequestPushTask {
	s.IOSLiveActivityAttributes = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSLiveActivityAttributesType(v string) *MassPushRequestPushTask {
	s.IOSLiveActivityAttributesType = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSLiveActivityContentState(v string) *MassPushRequestPushTask {
	s.IOSLiveActivityContentState = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSLiveActivityDismissalDate(v int64) *MassPushRequestPushTask {
	s.IOSLiveActivityDismissalDate = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSLiveActivityEvent(v string) *MassPushRequestPushTask {
	s.IOSLiveActivityEvent = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSLiveActivityId(v string) *MassPushRequestPushTask {
	s.IOSLiveActivityId = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSLiveActivityStaleDate(v int64) *MassPushRequestPushTask {
	s.IOSLiveActivityStaleDate = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSMusic(v string) *MassPushRequestPushTask {
	s.IOSMusic = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSMutableContent(v bool) *MassPushRequestPushTask {
	s.IOSMutableContent = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSNotificationCategory(v string) *MassPushRequestPushTask {
	s.IOSNotificationCategory = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSNotificationCollapseId(v string) *MassPushRequestPushTask {
	s.IOSNotificationCollapseId = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSNotificationThreadId(v string) *MassPushRequestPushTask {
	s.IOSNotificationThreadId = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSRelevanceScore(v float64) *MassPushRequestPushTask {
	s.IOSRelevanceScore = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSRemind(v bool) *MassPushRequestPushTask {
	s.IOSRemind = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSRemindBody(v string) *MassPushRequestPushTask {
	s.IOSRemindBody = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSSilentNotification(v bool) *MassPushRequestPushTask {
	s.IOSSilentNotification = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSSubtitle(v string) *MassPushRequestPushTask {
	s.IOSSubtitle = &v
	return s
}

func (s *MassPushRequestPushTask) Validate() error {
	return dara.Validate(s)
}
