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
	// The AppKey of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// An idempotency parameter that prevents duplicate pushes caused by API client retries. If you make a call with the same IdempotentToken within 15 minutes, only one push is performed, and subsequent calls return the result of the first successful push.
	//
	// > - The parameter format is a standard 36-character UUID (8-4-4-4-12). Each valid character is a hexadecimal digit from 0-9 or a-f, case-insensitive.
	//
	// >
	//
	// > - This parameter only prevents duplicate pushes caused by retries. It cannot prevent duplicate pushes caused by concurrent calls.
	//
	// example:
	//
	// c8016d13-6e76-410c-9bda-769383d11787
	IdempotentToken *string `json:"IdempotentToken,omitempty" xml:"IdempotentToken,omitempty"`
	// An array of independent push tasks.
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
	// Specifies the activity to open when the notification is tapped.
	//
	// This is required only when PushTask.N.AndroidOpenType is set to "Activity". For example: `com.alibaba.cloudpushdemo.bizactivity`.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidActivity *string `json:"AndroidActivity,omitempty" xml:"AndroidActivity,omitempty"`
	// Sets the value to add to the badge number. The value is added to the original badge number. The value range is [1, 99].
	//
	// > This is effective only for pushes through Huawei or Honor vendor channels. If both AndroidBadgeAddNum and AndroidBadgeSetNum are present, AndroidBadgeSetNum takes precedence.
	//
	// example:
	//
	// 1
	AndroidBadgeAddNum *int32 `json:"AndroidBadgeAddNum,omitempty" xml:"AndroidBadgeAddNum,omitempty"`
	// The full class name of the entry Activity of the application for badge settings.
	//
	// > This is effective only for pushes through Huawei or Honor vendor channels.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidBadgeClass *string `json:"AndroidBadgeClass,omitempty" xml:"AndroidBadgeClass,omitempty"`
	// Sets a fixed number for the badge. The value range is [0, 99].
	//
	// > For vendor channel pushes, this is effective only for Huawei and Honor channels. For pushes through Alibaba Cloud\\"s proprietary channel, this is effective only on Huawei, Honor, and vivo models.
	//
	// example:
	//
	// 5
	AndroidBadgeSetNum *int32 `json:"AndroidBadgeSetNum,omitempty" xml:"AndroidBadgeSetNum,omitempty"`
	// The body in long text mode. Length limit: 1,000 bytes (one Chinese character is counted as 3 bytes). The actual limit depends on the specific vendor channel.
	//
	// Currently supported on:
	//
	// - Huawei: EMUI 10 and later
	//
	// - Honor: Magic UI 4.0 and later
	//
	// - Xiaomi: MIUI 10 and later
	//
	// - OPPO: ColorOS 5.0 and later
	//
	// - Meizu: Flyme
	//
	// - Alibaba Cloud\\"s proprietary channel: Android SDK 3.6.0 and later
	//
	// If this parameter is not provided in long text mode, the first non-empty value from Body or AndroidPopupBody is used.
	//
	// example:
	//
	// 示例长文本
	AndroidBigBody *string `json:"AndroidBigBody,omitempty" xml:"AndroidBigBody,omitempty"`
	// The image URL in big picture mode. Currently supported on: Alibaba Cloud\\"s proprietary channel with Android SDK 3.6.0 or later.
	//
	// example:
	//
	// https://imag.example.com/image.png
	AndroidBigPictureUrl *string `json:"AndroidBigPictureUrl,omitempty" xml:"AndroidBigPictureUrl,omitempty"`
	// The title in long text mode. Length limit: 200 bytes (one Chinese character is counted as 3 bytes).
	//
	// - Currently, this is only supported by Honor channels and Huawei channels on EMUI 11 and later.
	//
	// - If this parameter is not provided in long text mode, the first non-empty value from Title or AndroidPopupTitle is used.
	//
	// example:
	//
	// 示例长标题
	AndroidBigTitle *string `json:"AndroidBigTitle,omitempty" xml:"AndroidBigTitle,omitempty"`
	// Sets the extended properties of the notification. This parameter does not take effect when the push type PushType is set to MESSAGE.
	//
	// This parameter must be passed in JSON map format, or it will fail to parse.
	//
	// example:
	//
	// {"key1":"value1","api_name":"PushNoticeToAndroidRequest"}
	AndroidExtParameters *string `json:"AndroidExtParameters,omitempty" xml:"AndroidExtParameters,omitempty"`
	// Sets the Honor channel notification type:
	//
	// - **0**: Formal notification (default).
	//
	// - **1**: Test notification.
	//
	// > Each application can send 1,000 test notifications per day, and these are not subject to the daily push limit per device.
	//
	// example:
	//
	// 1
	AndroidHonorTargetUserType *int32 `json:"AndroidHonorTargetUserType,omitempty" xml:"AndroidHonorTargetUserType,omitempty"`
	// Sets the parameters for Huawei quick notifications
	//
	// - **0**: Send a normal Huawei notification (default).
	//
	// - **1**: Send a Huawei quick notification.
	//
	// example:
	//
	// 1
	AndroidHuaweiBusinessType *int32 `json:"AndroidHuaweiBusinessType,omitempty" xml:"AndroidHuaweiBusinessType,omitempty"`
	// A JSON string of the Huawei Android Live Notification data structure [LiveNotificationPayload](https://developer.huawei.com/consumer/cn/doc/HMSCore-References/rest-live-0000001562939968#ZH-CN_TOPIC_0000001700850537__p195121620102511). For development and integration, see the [Huawei Live Notification Push Guide](https://help.aliyun.com/document_detail/2983768.html).
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
	// The receipt ID for the Huawei channel. View this receipt ID in the receipt parameter configuration on the Huawei Push operations platform.
	//
	// > If the default receipt configuration on the Huawei Push operations platform is the Alibaba Cloud receipt, you do not need to provide this. If not, we recommend that you first configure the default receipt ID for the Huawei channel in the Alibaba Cloud EMAS Mobile Push console.
	//
	// example:
	//
	// RCP4C123456
	AndroidHuaweiReceiptId *string `json:"AndroidHuaweiReceiptId,omitempty" xml:"AndroidHuaweiReceiptId,omitempty"`
	// Sets the Huawei channel notification type:
	//
	// - **0**: Formal notification (default).
	//
	// - **1**: Test notification.
	//
	// > Each application can send 500 test notifications per day, and these are not subject to the daily push limit per device.
	//
	// example:
	//
	// 1
	AndroidHuaweiTargetUserType *int32 `json:"AndroidHuaweiTargetUserType,omitempty" xml:"AndroidHuaweiTargetUserType,omitempty"`
	// The URL for the right-side icon. Currently supported on:
	//
	// - Huawei EMUI (applicable only in long text mode and inbox mode)
	//
	// - Honor Magic UI (applicable only in long text mode)
	//
	// - Alibaba Cloud\\"s proprietary channel: Android SDK 3.5.0 and later
	//
	// example:
	//
	// https://imag.example.com/image.png
	AndroidImageUrl *string `json:"AndroidImageUrl,omitempty" xml:"AndroidImageUrl,omitempty"`
	// The body in inbox mode. The content must be a valid JSON array with no more than 5 elements. Currently supported on:
	//
	// - Huawei: EMUI 9 and later
	//
	// - Honor: Magic UI 4.0 and later
	//
	// - Xiaomi: MIUI 10 and later
	//
	// - OPPO: ColorOS 5.0 and later
	//
	// - Alibaba Cloud\\"s proprietary channel: Android SDK 3.6.0 and later
	//
	// example:
	//
	// ["第一行","第二行"]
	AndroidInboxBody *string `json:"AndroidInboxBody,omitempty" xml:"AndroidInboxBody,omitempty"`
	// Meizu message type
	//
	// - 0 Public message (default)
	//
	// - 1 Private message
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 0
	AndroidMeizuNoticeMsgType *int32 `json:"AndroidMeizuNoticeMsgType,omitempty" xml:"AndroidMeizuNoticeMsgType,omitempty"`
	// Function 1: After applying for [self-classification permissions](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835?#section3410731125514), use this parameter to identify the message type, determine the [message reminder method](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#ZH-CN_TOPIC_0000001149358835__p3850133955718), and expedite the sending of specific message types. For valid values, see the [message classification standards](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section1076611477914) in the official Huawei Push documentation. Fill in the "Cloud-side notification category value" or "Local notification category value" from the documentation table.
	//
	// Function 2: After applying for [special permissions](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/faq-0000001050042183#section037425218509), use this parameter to identify high-priority pass-through scenarios. Valid values:
	//
	// - VOIP: Video calls
	//
	// - PLAY_VOICE: Voice playback
	//
	// > For "Cloud-side notification category values" that are "Not applicable", all pushes go through Alibaba Cloud\\"s proprietary channel. For "Local notification category values" that are "Not applicable", all pushes go through the Huawei channel.
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
	// You must apply for permissions. For more information, see [Application Link](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/faq-0000001050042183#section037425218509).
	//
	// example:
	//
	// HIGH
	AndroidMessageHuaweiUrgency *string `json:"AndroidMessageHuaweiUrgency,omitempty" xml:"AndroidMessageHuaweiUrgency,omitempty"`
	// OPPO manages messages in two categories: Communication and Services, and Content and Marketing.
	//
	// Communication and Services (requires permission application):
	//
	// - IM: Instant messaging, audio, and video calls
	//
	// - ACCOUNT: Changes in personal accounts and assets
	//
	// - DEVICE_REMINDER: Personal device reminders
	//
	// - ORDER: Changes in personal order or logistics status
	//
	// - TODO: Personal schedules or to-do items
	//
	// - SUBSCRIPTION: Personal subscriptions
	//
	// Content and Marketing:
	//
	// - NEWS: News and information
	//
	// - CONTENT: Content recommendations
	//
	// - MARKETING: Platform activities
	//
	// - SOCIAL: Social updates
	//
	// For more information, see [OPUSH Message Classification Rules](https://open.oppomobile.com/new/developmentDoc/info?id=13189)
	//
	// example:
	//
	// MARKETING
	AndroidMessageOppoCategory *string `json:"AndroidMessageOppoCategory,omitempty" xml:"AndroidMessageOppoCategory,omitempty"`
	// The reminder level for OPPO channel notification bar messages. Valid values:
	//
	// - 1: Notification bar
	//
	// - 2: Notification bar, lock screen, ringtone, vibration (default notification level for Communication and Services messages)
	//
	// - 16: Notification bar, lock screen, ringtone, vibration, banner (requires permission application)
	//
	// > When using the AndroidMessageOppoNotifyLevel parameter, you must also pass the AndroidMessageOppoCategory parameter.
	//
	// example:
	//
	// 1
	AndroidMessageOppoNotifyLevel *int32 `json:"AndroidMessageOppoNotifyLevel,omitempty" xml:"AndroidMessageOppoNotifyLevel,omitempty"`
	// vivo manages messages in two categories: system messages and operational messages.
	//
	// System messages:
	//
	// - IM: Instant messages
	//
	// - ACCOUNT: Account and asset
	//
	// - TODO: Schedule and to-do
	//
	// - DEVICE_REMINDER: Device information
	//
	// - ORDER: Order and logistics
	//
	// - SUBSCRIPTION: Subscription reminder
	//
	// Operational messages:
	//
	// - NEWS: News
	//
	// - CONTENT: Content recommendation
	//
	// - MARKETING: Operational activity
	//
	// - SOCIAL: Social updates
	//
	// > For more information, see [Classification Description](https://dev.vivo.com.cn/documentCenter/doc/359#s-ef3qugc3)
	//
	// example:
	//
	// TODO
	AndroidMessageVivoCategory *string `json:"AndroidMessageVivoCategory,omitempty" xml:"AndroidMessageVivoCategory,omitempty"`
	// The notification sound for the Huawei vendor channel. Specify the name of the audio file stored in the app/src/main/res/raw/ directory of the client project. Do not include the file format suffix.
	//
	// If you do not set this parameter, the default ringtone is used.
	//
	// example:
	//
	// alicloud_notification_sound
	AndroidMusic *string `json:"AndroidMusic,omitempty" xml:"AndroidMusic,omitempty"`
	// The priority that determines the position of the Android notification in the notification bar. Valid values: -2, -1, 0, 1, and 2.
	//
	// example:
	//
	// 0
	AndroidNotificationBarPriority *int32 `json:"AndroidNotificationBarPriority,omitempty" xml:"AndroidNotificationBarPriority,omitempty"`
	// The custom Android notification bar style. Valid values: 1 to 100.
	//
	// example:
	//
	// 2
	AndroidNotificationBarType *int32 `json:"AndroidNotificationBarType,omitempty" xml:"AndroidNotificationBarType,omitempty"`
	// The channel ID for the Android app. It must correspond to a channel ID in the app.
	//
	// - Set the NotificationChannel parameter. For more information about its use, see [FAQ: Why are notifications not received on devices with Android 8.0 or later?](https://help.aliyun.com/document_detail/67398.html).
	//
	// - Because the channel_id for the OPPO private message channel is the same as the app\\"s channelId, this value is used for the channel_id when pushing through the OPPO channel.
	//
	// - For pushes through Huawei, FCM, and Alibaba Cloud\\"s proprietary channels, this value is used for the channel_id.
	//
	// example:
	//
	// 1
	AndroidNotificationChannel *string `json:"AndroidNotificationChannel,omitempty" xml:"AndroidNotificationChannel,omitempty"`
	// Message grouping. For messages in the same group, the notification bar displays only the latest message and the total number of messages received for that group. It does not display all messages and cannot be expanded. Currently supported on:
	//
	// - Huawei vendor channel
	//
	// - Honor vendor channel
	//
	// - Alibaba Cloud\\"s proprietary channel with Android SDK 3.9.1 and earlier
	//
	// > This parameter is not supported by Alibaba Cloud\\"s proprietary channel on Android SDK 3.9.2 and later.
	//
	// example:
	//
	// group-1
	AndroidNotificationGroup *string `json:"AndroidNotificationGroup,omitempty" xml:"AndroidNotificationGroup,omitempty"`
	// Sets the importance parameter for Honor notification message classification, which determines the notification behavior on the user\\"s device. Valid values:
	//
	// - LOW: Marketing messages
	//
	// - NORMAL: Service and communication messages
	//
	// Apply for this on the Honor platform. [Application Link](https://developer.honor.com/cn/docs/11002/guides/notification-class#%E8%87%AA%E5%88%86%E7%B1%BB%E6%9D%83%E7%9B%8A%E7%94%B3%E8%AF%B7).
	//
	// example:
	//
	// LOW
	AndroidNotificationHonorChannel *string `json:"AndroidNotificationHonorChannel,omitempty" xml:"AndroidNotificationHonorChannel,omitempty"`
	// Sets the importance parameter for Huawei notification message classification, which determines the notification behavior on the user\\"s device. Valid values:
	//
	// - LOW: Marketing messages
	//
	// - NORMAL: Service and communication messages
	//
	// > 	- For the Huawei channel, use AndroidMessageHuaweiCategory for notification classification. AndroidNotificationHuaweiChannel is no longer required.
	//
	// >
	//
	// > 	- You must apply for this on the Huawei platform. [Application Link](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section893184112272).
	//
	// example:
	//
	// LOW
	AndroidNotificationHuaweiChannel *string `json:"AndroidNotificationHuaweiChannel,omitempty" xml:"AndroidNotificationHuaweiChannel,omitempty"`
	// The unique ID for each message when it is displayed as a notification. Different notification messages can have the same NotifyId to allow new notifications to overwrite old ones.
	//
	// example:
	//
	// 100001
	AndroidNotificationNotifyId *int32 `json:"AndroidNotificationNotifyId,omitempty" xml:"AndroidNotificationNotifyId,omitempty"`
	// Message grouping. Messages in the same group are displayed collapsed in the notification bar and can be expanded. Notifications from different groups are displayed separately. Currently supported on:
	//
	// - Alibaba Cloud\\"s proprietary channel with Android SDK 3.9.2 and later
	//
	// example:
	//
	// thread-1
	AndroidNotificationThreadId *string `json:"AndroidNotificationThreadId,omitempty" xml:"AndroidNotificationThreadId,omitempty"`
	// Sets the vivo notification message classification. Valid values:
	//
	// - 0: Operational messages (default)
	//
	// - 1: System messages
	//
	// > 	- For the vivo channel, use AndroidMessageVivoCategory for notification classification. AndroidNotificationVivoChannel is no longer required.
	//
	// >
	//
	// > 	- Apply for this on the vivo platform. For more information, see [Application Link](https://dev.vivo.com.cn/documentCenter/doc/359).
	//
	// example:
	//
	// 0
	AndroidNotificationVivoChannel *string `json:"AndroidNotificationVivoChannel,omitempty" xml:"AndroidNotificationVivoChannel,omitempty"`
	// Sets the channel ID for the Xiaomi notification type. Apply for it on the Xiaomi platform. For more information, see [Application Link](https://dev.mi.com/console/doc/detail?pId=2422#_4).
	//
	// > - A single application can apply for a maximum of 8 channels on the Xiaomi platform. Plan accordingly.
	//
	// example:
	//
	// michannel
	AndroidNotificationXiaomiChannel *string `json:"AndroidNotificationXiaomiChannel,omitempty" xml:"AndroidNotificationXiaomiChannel,omitempty"`
	// The notification reminder method. Valid values:
	//
	// - **VIBRATE**: Vibrate (default)
	//
	// - **SOUND**: Sound
	//
	// - **BOTH**: Sound and vibrate
	//
	// - **NONE**: Silent
	//
	// example:
	//
	// VIBRATE
	AndroidNotifyType *string `json:"AndroidNotifyType,omitempty" xml:"AndroidNotifyType,omitempty"`
	// The action to take after a notification is tapped. Valid values:
	//
	// - APPLICATION: Open the application (default)
	//
	// - ACTIVITY: Open the application\\"s AndroidActivity
	//
	// - URL: Open a URL
	//
	// - NONE: No action
	//
	// example:
	//
	// APPLICATION
	AndroidOpenType *string `json:"AndroidOpenType,omitempty" xml:"AndroidOpenType,omitempty"`
	// The URL to open after the Android device receives the push. This is required only when PushTask.N.AndroidOpenType is set to "URL".
	//
	// example:
	//
	// https://xxxx.xxx
	AndroidOpenUrl *string `json:"AndroidOpenUrl,omitempty" xml:"AndroidOpenUrl,omitempty"`
	// A JSON string of the OPPO Fluid Cloud intent deletion data structure [data](https://open.oppomobile.com/documentation/page/info?id=13578). This parameter is invalid if the AndroidOppoIntelligentIntent parameter is already filled. For development and integration, see the [OPPO Fluid Cloud Push Guide](https://help.aliyun.com/document_detail/2997310.html).
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
	// A JSON string of the OPPO Fluid Cloud intent sharing data structure [IntelligentIntent](https://open.oppomobile.com/documentation/page/info?id=13565). For development and integration, see the [OPPO Fluid Cloud Push Guide](https://help.aliyun.com/document_detail/2997310.html).
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
	// Sets the OPPO Fluid Cloud push environment
	//
	// - **0**: Production environment (default).
	//
	// - **1**: Staging environment.
	//
	// > The OPPO Fluid Cloud staging environment needs to be set up on the client side. For more information, see [Environment Setup](https://open.oppomobile.com/documentation/page/info?id=13590).
	//
	// example:
	//
	// 1
	AndroidOppoIntentEnv *int32 `json:"AndroidOppoIntentEnv,omitempty" xml:"AndroidOppoIntentEnv,omitempty"`
	// Deprecated
	//
	// OPPO private message template content parameters
	AndroidOppoPrivateContentParameters map[string]*string `json:"AndroidOppoPrivateContentParameters,omitempty" xml:"AndroidOppoPrivateContentParameters,omitempty"`
	// Deprecated
	//
	// OPPO private message template ID
	//
	// 	Warning:
	//
	// The OPPO private message template feature is no longer supported by MaasPush. To use this feature, use the Push, PushV2, or MassPushV2 API instead.
	//
	// example:
	//
	// 687557242b1634hzef3zd5013
	AndroidOppoPrivateMsgTemplateId *string `json:"AndroidOppoPrivateMsgTemplateId,omitempty" xml:"AndroidOppoPrivateMsgTemplateId,omitempty"`
	// Deprecated
	//
	// OPPO private message template title parameters
	AndroidOppoPrivateTitleParameters map[string]*string `json:"AndroidOppoPrivateTitleParameters,omitempty" xml:"AndroidOppoPrivateTitleParameters,omitempty"`
	// Specifies the Activity to which the user is redirected after tapping the notification.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidPopupActivity *string `json:"AndroidPopupActivity,omitempty" xml:"AndroidPopupActivity,omitempty"`
	// The body content in auxiliary pop-up mode. This parameter is required if the AndroidPopupActivity parameter is not empty.
	//
	// Length limit: 200 characters. Both Chinese and English characters count as one.
	//
	// If you use a vendor channel, comply with the vendor channel\\"s restrictions. For more information, see [Limits on auxiliary channel pushes for Android](https://help.aliyun.com/document_detail/165253.html).
	//
	// example:
	//
	// hello
	AndroidPopupBody *string `json:"AndroidPopupBody,omitempty" xml:"AndroidPopupBody,omitempty"`
	// The title content in auxiliary pop-up mode. This parameter is required if the AndroidPopupActivity parameter is not empty.
	//
	// Length limit: 50 characters. Both Chinese and English characters count as one.
	//
	// If you use a vendor channel, comply with the vendor channel\\"s restrictions. For more information, see [Limits on auxiliary channel pushes for Android](https://help.aliyun.com/document_detail/165253.html).
	//
	// example:
	//
	// hello
	AndroidPopupTitle *string `json:"AndroidPopupTitle,omitempty" xml:"AndroidPopupTitle,omitempty"`
	// If the push type is MESSAGE and the device is offline, this push uses the auxiliary pop-up feature. The default value is false. This parameter takes effect only when PushType is MESSAGE.
	//
	// If a message is successfully converted to a notification, the displayed notification uses the values of the AndroidPopupTitle and AndroidPopupBody parameters. When the user taps the notification, the data retrieved by the onSysNoticeOpened method of the auxiliary pop-up uses the values of the Title and Body parameters.
	//
	// example:
	//
	// true
	AndroidRemind *bool `json:"AndroidRemind,omitempty" xml:"AndroidRemind,omitempty"`
	// The notification style. Valid values:
	//
	// - **0**: Standard mode (default)
	//
	// - **1**: Long text mode (supported by Huawei, Honor, Xiaomi, OPPO, Meizu, and Alibaba Cloud\\"s proprietary channel)
	//
	// - **2**: Big picture mode (supported by Alibaba Cloud\\"s proprietary channel, not supported on Xiaomi models)
	//
	// - **3**: List mode (supported by Huawei, Honor, Xiaomi, OPPO, and Alibaba Cloud\\"s proprietary channel)
	//
	// > This parameter is required for non-standard modes.
	//
	// example:
	//
	// 1
	AndroidRenderStyle *string `json:"AndroidRenderStyle,omitempty" xml:"AndroidRenderStyle,omitempty"`
	// Sets the vendor channel notification type:
	//
	// - **0**: Formal notification (default).
	//
	// - **1**: Test notification.
	//
	// > 	- Configuring this parameter is equivalent to configuring the AndroidHuaweiTargetUserType, AndroidHonorTargetUserType, AndroidVivoPushMode, and AndroidOppoIntentEnv parameters at the same time. A specific vendor channel parameter can override this parameter.
	//
	// >
	//
	// > 	- Currently supported: Huawei channel, Honor channel, vivo channel, OPPO Fluid Cloud.
	//
	// example:
	//
	// 1
	AndroidTargetUserType *int32 `json:"AndroidTargetUserType,omitempty" xml:"AndroidTargetUserType,omitempty"`
	// A JSON string of the vivo Atomic Island data structure [liveMessage](https://dev.vivo.com.cn/documentCenter/doc/896#s-fdagzbd4). For development and integration, see the [vivo Atomic Island Push Guide](https://help.aliyun.com/zh/document_detail/3030718.html).
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
	// Sets the vivo channel notification type:
	//
	// - **0**: Formal push (default).
	//
	// - **1**: Test push.
	//
	// > For test pushes, configure the test devices in the vivo console beforehand. You can obtain the test device\\"s RegId by searching for "onReceiveRegId regId" in the device startup logs.
	//
	// example:
	//
	// 1
	AndroidVivoPushMode *int32 `json:"AndroidVivoPushMode,omitempty" xml:"AndroidVivoPushMode,omitempty"`
	// The receipt ID for the vivo channel. View this receipt ID in the application information of the push service on the vivo open platform.
	//
	// > If the default receipt configuration on the vivo open platform is the Alibaba Cloud receipt, you do not need to provide this. If not, we recommend that you first configure the default receipt ID for the vivo channel in the Alibaba Cloud EMAS Mobile Push console.
	//
	// example:
	//
	// 123
	AndroidVivoReceiptId *string `json:"AndroidVivoReceiptId,omitempty" xml:"AndroidVivoReceiptId,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. All third-party auxiliary pop-ups are now supported by the new **AndroidPopupActivity*	- parameter.
	//
	// example:
	//
	// 无
	AndroidXiaoMiActivity *string `json:"AndroidXiaoMiActivity,omitempty" xml:"AndroidXiaoMiActivity,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. All third-party auxiliary pop-ups are now supported by the new **AndroidPopupBody*	- parameter.
	//
	// example:
	//
	// 无
	AndroidXiaoMiNotifyBody *string `json:"AndroidXiaoMiNotifyBody,omitempty" xml:"AndroidXiaoMiNotifyBody,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. All third-party auxiliary pop-ups are now supported by the new **AndroidPopupTitle*	- parameter.
	//
	// example:
	//
	// 无
	AndroidXiaoMiNotifyTitle *string `json:"AndroidXiaoMiNotifyTitle,omitempty" xml:"AndroidXiaoMiNotifyTitle,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Since August 2023, Xiaomi no longer supports dynamically setting small icons, right-side icons, or large pictures during pushes on new devices or systems.
	//
	// example:
	//
	// https://f6.market.xiaomi.com/download/MiPass/aaa/bbb.png
	AndroidXiaomiBigPictureUrl *string `json:"AndroidXiaomiBigPictureUrl,omitempty" xml:"AndroidXiaomiBigPictureUrl,omitempty"`
	// A JSON string of the Xiaomi HyperOS Island data structure [miui.focus.param](https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131). For development and integration, see the [Xiaomi HyperOS Island Push Guide](https://help.aliyun.com/zh/document_detail/3037956.html).
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
	// A JSON string of the Xiaomi HyperOS Island data image [miui.focus.pic_xxx](https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131). For development and integration, see the [Xiaomi HyperOS Island Push Guide](https://help.aliyun.com/zh/document_detail/3037956.html).
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
	// This parameter is deprecated. Since August 2023, Xiaomi no longer supports dynamically setting small icons, right-side icons, or large pictures during pushes on new devices or systems.
	//
	// example:
	//
	// https://imag.example.com/image.png
	AndroidXiaomiImageUrl       *string `json:"AndroidXiaomiImageUrl,omitempty" xml:"AndroidXiaomiImageUrl,omitempty"`
	AndroidXiaomiTemplateId     *string `json:"AndroidXiaomiTemplateId,omitempty" xml:"AndroidXiaomiTemplateId,omitempty"`
	AndroidXiaomiTemplateParams *string `json:"AndroidXiaomiTemplateParams,omitempty" xml:"AndroidXiaomiTemplateParams,omitempty"`
	// The content of the notification or message for Android and HarmonyOS pushes. The content of the message or notification for iOS. The content size is limited. For more information, see [Product limits](https://help.aliyun.com/document_detail/92832.html).
	//
	// example:
	//
	// hello
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The device type. Valid values:
	//
	// - HARMONY: HarmonyOS devices
	//
	// - iOS: iOS devices
	//
	// - ANDROID: Android devices
	//
	// - ALL: If the AppKey is for an old version of a dual-platform application, this value indicates that pushes are sent to both Android and iOS devices. If the AppKey is for a new version of a single-platform application, the effect is the same as specifying the device type corresponding to that application type.
	//
	// This parameter is required.
	//
	// example:
	//
	// HARMONY
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The time-to-live (TTL) for offline messages or notifications. Use this with StoreOffline. After the TTL expires, the message or notification is no longer sent. The maximum TTL is 72 hours. The default is 72 hours.
	//
	// The time must be in ISO 8601 format and in UTC: YYYY-MM-DDThh:mm:ssZ. The expiration time must be at least 3 seconds later than the current time or the scheduled push time (`ExpireTime > PushTime + 3 seconds`). The 3-second buffer accounts for potential network and system latency. Set the TTL to at least 1 minute for individual pushes and at least 10 minutes for full or batch pushes.
	//
	// example:
	//
	// 2019-02-20T00:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The action corresponding to the in-app page ability.
	//
	// 	Notice:
	//
	// When HarmonyActionType is APP_CUSTOM_PAGE, at least one of HarmonyUri and HarmonyAction must be filled in.
	//
	//
	//
	// For more information, see [ClickAction.action](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section152462191216) on the HarmonyOS website.
	//
	// example:
	//
	// com.example.action
	HarmonyAction *string `json:"HarmonyAction,omitempty" xml:"HarmonyAction,omitempty"`
	// The action to take after a notification is tapped. Valid values:
	//
	// - APP_HOME_PAGE: Open the application home page
	//
	// - APP_CUSTOM_PAGE: Open a custom application page
	//
	// example:
	//
	// APP_HOME_PAGE
	HarmonyActionType *string `json:"HarmonyActionType,omitempty" xml:"HarmonyActionType,omitempty"`
	// The number to add to the HarmonyOS application badge. See the [HarmonyOS badge addNum field description](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section266310382145).<br>
	//
	// Supported starting from HarmonyOS SDK version 1.2.0.<br>
	//
	// example:
	//
	// 1
	HarmonyBadgeAddNum *int32 `json:"HarmonyBadgeAddNum,omitempty" xml:"HarmonyBadgeAddNum,omitempty"`
	// The number to set for the HarmonyOS application badge. See the [HarmonyOS badge setNum field description](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section266310382145).<br>
	//
	// Supported starting from HarmonyOS SDK version 1.2.0.<br>
	//
	// example:
	//
	// 1
	HarmonyBadgeSetNum *int32 `json:"HarmonyBadgeSetNum,omitempty" xml:"HarmonyBadgeSetNum,omitempty"`
	// The notification message category. After applying for notification message self-classification permissions, use this to identify the message type. Different notification message types affect how messages are displayed and reminded. Valid values:
	//
	// - IM: Instant messaging
	//
	// - VOIP: Video call
	//
	// - SUBSCRIPTION: Subscription
	//
	// - TRAVEL: Travel
	//
	// - HEALTH: Health
	//
	// - WORK: Work item reminder
	//
	// - ACCOUNT: Account updates
	//
	// - EXPRESS: Order & logistics
	//
	// - FINANCE: Finance
	//
	// - DEVICE_REMINDER: Device reminder
	//
	// - MAIL: Email
	//
	// - CUSTOMER_SERVICE: Customer service message
	//
	// - MARKETING: News, content recommendations, social updates, product promotions, financial updates, lifestyle information, surveys, feature recommendations, and operational activities (only identifies content, does not expedite message sending), collectively referred to as marketing messages.
	//
	// For more information, see [Notification.category](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117) on the HarmonyOS website.
	//
	// example:
	//
	// IM
	HarmonyCategory *string `json:"HarmonyCategory,omitempty" xml:"HarmonyCategory,omitempty"`
	// Sets the extended properties of the notification. This parameter does not take effect when the push type PushType is set to MESSAGE.
	//
	// This parameter must be passed in JSON map format, or it will fail to parse.
	//
	// example:
	//
	// {"key1":"value1","api_name":"PushNoticeToAndroidRequest"}
	HarmonyExtParameters *string `json:"HarmonyExtParameters,omitempty" xml:"HarmonyExtParameters,omitempty"`
	// Extra data for the extended notification message.<br>
	//
	// Effective when sending HarmonyOS extended notification messages.<br>
	//
	// Conceptually equivalent to the extraData field of a HarmonyOS extended notification message. For a detailed definition, see [HarmonyOS ExtensionPayload Description](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section161192514234).<br>
	//
	// Supported starting from HarmonyOS SDK version 1.2.0.<br><br><br>
	//
	// example:
	//
	// 示例额外数据
	HarmonyExtensionExtraData *string `json:"HarmonyExtensionExtraData,omitempty" xml:"HarmonyExtensionExtraData,omitempty"`
	// When PushType is NOTICE, specifies whether it is a HarmonyOS extended notification message.
	//
	// - true: Send an extended notification message
	//
	// - false: Send a normal notification (default)
	//
	// You must apply for permission on the HarmonyOS side before sending extended notification messages. For more information, see [Send Extended Notification Messages](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V5/push-send-extend-noti-V5) in the HarmonyOS documentation.<br>
	//
	// Supported starting from HarmonyOS SDK version 1.2.0.<br>
	//
	// example:
	//
	// true
	HarmonyExtensionPush *bool `json:"HarmonyExtensionPush,omitempty" xml:"HarmonyExtensionPush,omitempty"`
	// The URL for the large icon on the right of the notification. The URL must use the HTTPS protocol.
	//
	// > Supported image formats are png, jpg, jpeg, heif, gif, and bmp. The image length × width must be less than 25,000 pixels.
	//
	// For more information, see [Notification.image](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117) on the HarmonyOS website.
	//
	// example:
	//
	// https://example.com/xxx.png
	HarmonyImageUrl *string `json:"HarmonyImageUrl,omitempty" xml:"HarmonyImageUrl,omitempty"`
	// The content for the multi-line text style. This field is required when HarmonyRenderStyle is MULTI_LINE. A maximum of 3 content entries are supported.
	//
	// example:
	//
	// ["1.content1","2.content2","3.content3"]
	HarmonyInboxContent *string `json:"HarmonyInboxContent,omitempty" xml:"HarmonyInboxContent,omitempty"`
	// A JSON string of the HarmonyOS Live Window data structure [LiveViewPayload](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V13/push-scenariozed-api-request-param-V13#section66881469306). For development and integration, see the [HarmonyOS Live Window Push Guide](https://help.aliyun.com/document_detail/2982112.html).
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
	// Uses the specified type of notification channel. This is effective only when the Alibaba Cloud proprietary channel is online.
	//
	// - SOCIAL_COMMUNICATION: Social communication.
	//
	// - SERVICE_INFORMATION: Service reminder.
	//
	// - CONTENT_INFORMATION: Content information.
	//
	// - CUSTOMER_SERVICE: Customer service message. This type is used for customer service messages between users and businesses and must be initiated by the user.
	//
	// - OTHER_TYPES: Other.
	//
	// For more information, see [SlotType](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/js-apis-notificationmanager-V5#slottype) on the HarmonyOS website.
	//
	// example:
	//
	// SOCIAL_COMMUNICATION
	HarmonyNotificationSlotType *string `json:"HarmonyNotificationSlotType,omitempty" xml:"HarmonyNotificationSlotType,omitempty"`
	// The unique ID for each message when it is displayed as a notification. If not included, the push service automatically generates a unique ID for each message. Different notification messages can have the same notifyId to allow new messages to overwrite old ones.
	//
	// For more information, see [Notification.notifyId](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117) on the HarmonyOS website.
	//
	// example:
	//
	// 0
	HarmonyNotifyId *int32 `json:"HarmonyNotifyId,omitempty" xml:"HarmonyNotifyId,omitempty"`
	// The receipt ID for the HarmonyOS channel. View this receipt ID in the receipt parameter configuration on the HarmonyOS Push operations platform.
	//
	// > If the default receipt configuration on the HarmonyOS Push operations platform is the Alibaba Cloud receipt, you do not need to provide this. If not, we recommend that you first configure the default receipt ID for the HarmonyOS channel in the Alibaba Cloud EMAS Mobile Push console.
	//
	// For more information, see [pushOptions.receiptId](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section418321011212) on the HarmonyOS website.
	//
	// example:
	//
	// RCPB***DFD5
	HarmonyReceiptId *string `json:"HarmonyReceiptId,omitempty" xml:"HarmonyReceiptId,omitempty"`
	// If the push type is MESSAGE and the device is offline, this push uses the auxiliary pop-up feature. The default value is false. This parameter is effective only when PushType is set to MESSAGE.
	//
	// If a message is successfully converted to a notification, the data displayed in the notification is the value of the server-side HarmonyRemindTitle and HarmonyRemindBody parameters.
	//
	// example:
	//
	// false
	HarmonyRemind *bool `json:"HarmonyRemind,omitempty" xml:"HarmonyRemind,omitempty"`
	// The HarmonyOS notification content used when a message is converted to a notification. This is effective only when HarmonyRemind is set to true.
	//
	// example:
	//
	// 您有一条新消息，请查收
	HarmonyRemindBody *string `json:"HarmonyRemindBody,omitempty" xml:"HarmonyRemindBody,omitempty"`
	// The HarmonyOS notification title used when a message is converted to a notification. This is effective only when HarmonyRemind is set to true.
	//
	// example:
	//
	// 新消息
	HarmonyRemindTitle *string `json:"HarmonyRemindTitle,omitempty" xml:"HarmonyRemindTitle,omitempty"`
	// The notification message style:
	//
	// - NORMAL: Normal notification (default)
	//
	// - MULTI_LINE: Multi-line text style
	//
	// example:
	//
	// NORMAL
	HarmonyRenderStyle *string `json:"HarmonyRenderStyle,omitempty" xml:"HarmonyRenderStyle,omitempty"`
	// The test message flag:
	//
	// - false: Normal message (default)
	//
	// - true: Test message
	//
	// For more information, see [pushOptions.testMessage](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section418321011212) on the HarmonyOS website.
	//
	// example:
	//
	// true
	HarmonyTestMessage *bool `json:"HarmonyTestMessage,omitempty" xml:"HarmonyTestMessage,omitempty"`
	// The URI corresponding to the in-app page ability.
	//
	// 	Notice: When HarmonyActionType is APP_CUSTOM_PAGE, at least one of HarmonyUri and HarmonyAction must be filled in. When there are multiple Abilities, fill in the action and URI for each Ability separately. The action is used with priority to find the corresponding in-app page.
	//
	// For more information, see [ClickAction.uri](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section152462191216) on the HarmonyOS website.
	//
	// example:
	//
	// https://www.example.com:8080/push/example
	HarmonyUri *string `json:"HarmonyUri,omitempty" xml:"HarmonyUri,omitempty"`
	// A custom ID for the push task. If JobKey is not empty, this field is included in the receipt logs. For more information about how to view receipt logs, see [Receipt logs](https://help.aliyun.com/document_detail/434651.html).
	//
	// example:
	//
	// 123
	JobKey *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	// Specifies the time for a scheduled push. If you do not set this parameter, the push is sent immediately.
	//
	// The time must be in ISO 8601 format and in UTC: YYYY-MM-DDThh:mm:ssZ.
	//
	// example:
	//
	// 2019-02-20T00:00:00Z
	PushTime *string `json:"PushTime,omitempty" xml:"PushTime,omitempty"`
	// The push type. Valid values:
	//
	// - MESSAGE: a message.
	//
	// - NOTICE: a notification.
	//
	// This parameter is required.
	//
	// example:
	//
	// MESSAGE
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	// Specifies the sending channels. Valid values:
	//
	// - accs: Alibaba Cloud\\"s proprietary channel
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
	// > 	- If you do not configure this parameter, all channels can be used.
	//
	// >
	//
	// > 	- If you configure this parameter, only the specified channels are used.
	//
	// >
	//
	// > 	- If the configured channels conflict with the sending policy (for example, iOS notifications are sent only through the APNs channel, but this parameter does not include apns), the push is not sent.
	//
	// >
	//
	// > 	- If you configure gcm, both Google GCM and FCM channels can be used. If you configure fcm, only the Google FCM channel can be used.
	//
	// example:
	//
	// accs,huawei,xiaomi
	SendChannels *string `json:"SendChannels,omitempty" xml:"SendChannels,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated.
	//
	// example:
	//
	// 0
	SendSpeed *int32 `json:"SendSpeed,omitempty" xml:"SendSpeed,omitempty"`
	// Specifies whether to save offline messages or notifications. The default value is false.
	//
	// If you save them, and a user is offline, the message or notification is resent when the user comes online before the time-to-live (TTL) specified by ExpireTime expires. The default TTL is 72 hours. iOS notifications are sent through the APNs channel and are not affected by this parameter.
	//
	// example:
	//
	// true
	StoreOffline *bool `json:"StoreOffline,omitempty" xml:"StoreOffline,omitempty"`
	// The push target. Valid values:
	//
	// - DEVICE: Push by device.
	//
	// - ACCOUNT: Push by account.
	//
	// - ALIAS: Push by alias.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEVICE
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// Set this parameter based on the value of Target. To specify multiple values, separate them with commas. If you exceed the limit, send multiple pushes.
	//
	// - If you set Target to DEVICE, specify device IDs, such as `deviceid1,deviceid2`. You can specify up to 1,000 device IDs.
	//
	// - If you set Target to ACCOUNT, specify accounts, such as `account1,account2`. You can specify up to 1,000 accounts.
	//
	// - If you set Target to ALIAS, specify aliases, such as `alias1,alias2`. You can specify up to 1,000 aliases.
	//
	// This parameter is required.
	//
	// example:
	//
	// deviceid1,deviceid2
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
	// The title of the notification or message. The length is limited to 200 bytes.
	//
	// This parameter is required for Android and HarmonyOS pushes. It is optional for iOS notification pushes. If you specify it for iOS:
	//
	// - For iOS 10 and later, the notification title is displayed.
	//
	// - For iOS versions from 8.2 to 10, it replaces the application name in the notification.
	//
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Specifies whether to automatically truncate titles and content that are too long.
	//
	// Note: This applies only to vendor channels that have explicit limits on title and content length. It does not apply to channels such as APNs, Huawei, and Honor, which limit the total request body size instead of the title and content length.
	//
	// example:
	//
	// false
	Trim *bool `json:"Trim,omitempty" xml:"Trim,omitempty"`
	// iOS notifications are sent through APNs. Specify the environment information.
	//
	// - DEV: The development environment. This applies to applications installed and debugged directly from Xcode.
	//
	// - PRODUCT: The production environment. This applies to applications distributed through the App Store, TestFlight, Ad Hoc, or enterprise distribution.
	//
	// example:
	//
	// DEV
	IOSApnsEnv *string `json:"iOSApnsEnv,omitempty" xml:"iOSApnsEnv,omitempty"`
	// The badge number on the top-right corner of the iOS application icon.
	//
	// > If iOSBadgeAutoIncrement is set to true, this parameter must be empty.
	//
	// example:
	//
	// 0
	IOSBadge *int32 `json:"iOSBadge,omitempty" xml:"iOSBadge,omitempty"`
	// Specifies whether to enable the auto-increment badge feature. The default value is false.
	//
	// > When this is set to true, iOSBadge must be empty.
	//
	// The auto-increment badge feature is maintained by the push server, which keeps a badge count for each device. To use this feature, use SDK version 1.9.5 or later and actively sync the badge number to the server.
	//
	// example:
	//
	// true
	IOSBadgeAutoIncrement *bool `json:"iOSBadgeAutoIncrement,omitempty" xml:"iOSBadgeAutoIncrement,omitempty"`
	// The extended properties for iOS notifications.
	//
	// For iOS 10 and later, specify the resource URL for a rich push notification, such as `{"attachment": "https://xxxx.xxx/notification_pic.png"}`. This parameter must be passed in JSON map format, or it will fail to parse.
	//
	// example:
	//
	// {"attachment": "https://xxxx.xxx/notification_pic.png"}
	IOSExtParameters *string `json:"iOSExtParameters,omitempty" xml:"iOSExtParameters,omitempty"`
	// The interruption level. Valid values:
	//
	// - passive: The system adds the notification to the notification list without lighting up the screen or playing a sound.
	//
	// - active: The system immediately displays the notification, lights up the screen, and can play a sound.
	//
	// - time-sensitive: The system immediately presents the notification, lights up the screen, and can play a sound, but does not break through system notification controls.
	//
	// - critical: The system immediately displays the notification, lights up the screen, and plays a sound, bypassing the mute switch.
	//
	// example:
	//
	// active
	IOSInterruptionLevel *string `json:"iOSInterruptionLevel,omitempty" xml:"iOSInterruptionLevel,omitempty"`
	// A JSON string for the static pass-through parameters of a Dynamic Island push. It contains static, user-defined information, such as product numbers and order information.
	//
	// > Required when iOSLiveActivityEvent is set to start.
	//
	// example:
	//
	// {"orderId": "12345", "product": "Shoes"}
	IOSLiveActivityAttributes *string `json:"iOSLiveActivityAttributes,omitempty" xml:"iOSLiveActivityAttributes,omitempty"`
	// The type of Live Activity to start.
	//
	// > Required when iOSLiveActivityEvent is set to start.
	//
	// example:
	//
	// OrderActivityAttributes
	IOSLiveActivityAttributesType *string `json:"iOSLiveActivityAttributesType,omitempty" xml:"iOSLiveActivityAttributesType,omitempty"`
	// The dynamic pass-through parameters for a Dynamic Island push. It contains real-time updated information, such as price or inventory changes.
	//
	// example:
	//
	// {"status": "delivered", "estimatedArrival": "2023-12-31T12:00:00Z"}
	IOSLiveActivityContentState *string `json:"iOSLiveActivityContentState,omitempty" xml:"iOSLiveActivityContentState,omitempty"`
	// The time until which an ended Live Activity remains on the lock screen. The maximum duration is 4 hours.
	//
	// example:
	//
	// 1743131967
	IOSLiveActivityDismissalDate *int64 `json:"iOSLiveActivityDismissalDate,omitempty" xml:"iOSLiveActivityDismissalDate,omitempty"`
	// Starts, updates, or ends a Live Activity.
	//
	// - Enumeration: start | update | end
	//
	// example:
	//
	// start
	IOSLiveActivityEvent *string `json:"iOSLiveActivityEvent,omitempty" xml:"iOSLiveActivityEvent,omitempty"`
	// The Live Activity ID reported by the device to your server. This is the unique identifier for a Live Activity.
	//
	// example:
	//
	// 66B94673-B32E-4CA7-863C-3E523054FD46
	IOSLiveActivityId *string `json:"iOSLiveActivityId,omitempty" xml:"iOSLiveActivityId,omitempty"`
	// A UNIX timestamp in seconds that marks the content of the activity as outdated.
	//
	// example:
	//
	// 1743131967
	IOSLiveActivityStaleDate *int64 `json:"iOSLiveActivityStaleDate,omitempty" xml:"iOSLiveActivityStaleDate,omitempty"`
	// The sound for the iOS notification. Specify the name of the audio file stored in the app bundle or the Library/Sounds directory of the sandbox. For more information, see How to set notification sounds for iOS pushes.
	//
	// If you specify an empty string (""), the notification is silent. If you do not set this parameter, the default value is \\"default\\", which is the system alert sound.
	//
	// example:
	//
	// ””
	IOSMusic *string `json:"iOSMusic,omitempty" xml:"iOSMusic,omitempty"`
	// The mutable content flag for iOS notifications (for iOS 10 and later). If set to true, notifications pushed through APNs can be processed by an extension before being displayed. For silent notifications, this must be set to true.
	//
	// example:
	//
	// true
	IOSMutableContent *bool `json:"iOSMutableContent,omitempty" xml:"iOSMutableContent,omitempty"`
	// Specifies the iOS notification category (for iOS 10 and later).
	//
	// example:
	//
	// ios
	IOSNotificationCategory *string `json:"iOSNotificationCategory,omitempty" xml:"iOSNotificationCategory,omitempty"`
	// If a device receives messages with the same CollapseId, they are merged into one. If the device is offline and receives multiple messages with the same CollapseId, only one is displayed in the notification bar. This parameter is supported on iOS 10 and later.
	//
	// example:
	//
	// ZD2011
	IOSNotificationCollapseId *string `json:"iOSNotificationCollapseId,omitempty" xml:"iOSNotificationCollapseId,omitempty"`
	// Groups iOS remote notifications using this property. It marks the identifier for a collapsed group. This is supported only on iOS 12.0 and later.
	//
	// example:
	//
	// abc
	IOSNotificationThreadId *string `json:"iOSNotificationThreadId,omitempty" xml:"iOSNotificationThreadId,omitempty"`
	// The score for highlighting the summary. The value must be a floating-point number from 0 to 1.
	//
	// example:
	//
	// 0.01
	IOSRelevanceScore *float64 `json:"iOSRelevanceScore,omitempty" xml:"iOSRelevanceScore,omitempty"`
	// If a device is offline when a message is pushed (meaning the persistent connection to the Mobile Push server is unavailable), this push is sent once as a notification through the Apple APNs channel.
	//
	// > Converting offline messages to notifications is only applicable to the production environment.
	//
	// example:
	//
	// true
	IOSRemind *bool `json:"iOSRemind,omitempty" xml:"iOSRemind,omitempty"`
	// The content of the iOS notification when an iOS message is converted to a notification. This parameter is valid only when iOSApnsEnv is set to PRODUCT and iOSRemind is set to true.
	//
	// example:
	//
	// ios通知body
	IOSRemindBody *string `json:"iOSRemindBody,omitempty" xml:"iOSRemindBody,omitempty"`
	// Specifies whether to enable iOS silent notifications.
	//
	// example:
	//
	// true
	IOSSilentNotification *bool `json:"iOSSilentNotification,omitempty" xml:"iOSSilentNotification,omitempty"`
	// The subtitle of the iOS notification (for iOS 10 and later).
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
