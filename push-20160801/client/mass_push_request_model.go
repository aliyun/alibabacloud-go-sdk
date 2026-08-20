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
	// AppKey information.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// An idempotency parameter used to prevent duplicate pushes caused by API caller retries. When calls are made with the same IdempotentToken within 15 minutes, only one push is performed, and subsequent calls return the result of the first successful push.
	//
	// >
	//
	// > - The parameter format is a standard 36-character UUID (8-4-4-4-12). Each valid character is a hexadecimal digit in the range 0-9 or a-f, case-insensitive.
	//
	// > - This parameter is only used to prevent duplicate pushes caused by retries and cannot prevent duplicate pushes caused by concurrent calls.
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
	// Specify the Activity to open from the notification.
	//
	// Only applicable when PushTask.N.AndroidOpenType="Activity", e.g., `com.alibaba.cloudpushdemo.bizactivity`.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidActivity *string `json:"AndroidActivity,omitempty" xml:"AndroidActivity,omitempty"`
	// Set the badge increment value. The value is added to the existing badge count. Value range: [1-99].
	//
	// > Only valid for Huawei/Honor vendor channel pushes. When both AndroidBadgeAddNum and AndroidBadgeSetNum are present, AndroidBadgeSetNum takes precedence.
	//
	// example:
	//
	// 1
	AndroidBadgeAddNum *int32 `json:"AndroidBadgeAddNum,omitempty" xml:"AndroidBadgeAddNum,omitempty"`
	// Full class name of the app entry Activity for badge settings.
	//
	// > Only valid for Huawei/Honor vendor channel pushes.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidBadgeClass *string `json:"AndroidBadgeClass,omitempty" xml:"AndroidBadgeClass,omitempty"`
	// Set the badge to a fixed number. Value range: [0-99].
	//
	// > For vendor channel pushes, this only takes effect on Huawei and Honor channels. For Alibaba Cloud proprietary channel pushes, this only takes effect on Huawei, Honor, and vivo devices.
	//
	// example:
	//
	// 5
	AndroidBadgeSetNum *int32 `json:"AndroidBadgeSetNum,omitempty" xml:"AndroidBadgeSetNum,omitempty"`
	// Body in long text mode. Length limit: 1,000 bytes (1 Chinese character counts as 3 bytes). Subject to specific vendor channel restrictions when sending.
	//
	// Currently supported:
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
	// - Proprietary channel: Android SDK 3.6.0 and later
	//
	// If this parameter is not provided in long text mode, the first non-empty value from Body and AndroidPopupBody is used.
	//
	// example:
	//
	// 示例长文本
	AndroidBigBody *string `json:"AndroidBigBody,omitempty" xml:"AndroidBigBody,omitempty"`
	// Image URL in big picture mode. Currently supported: Proprietary channel: Android SDK 3.6.0 and later.
	//
	// example:
	//
	// https://imag.example.com/image.png
	AndroidBigPictureUrl *string `json:"AndroidBigPictureUrl,omitempty" xml:"AndroidBigPictureUrl,omitempty"`
	// Title in long text mode. Length limit: 200 bytes (1 Chinese character counts as 3 bytes).
	//
	// - Currently only supported by the Honor channel and Huawei channel EMUI 11 and later.
	//
	// - If this parameter is not provided in long text mode, the first non-empty value from Title and AndroidPopupTitle is used.
	//
	// example:
	//
	// 示例长标题
	AndroidBigTitle *string `json:"AndroidBigTitle,omitempty" xml:"AndroidBigTitle,omitempty"`
	// Set notification extended properties. This property does not take effect when the push type PushType is set to MESSAGE.
	//
	// This parameter must be passed in JSON map format; otherwise, parsing errors will occur.
	//
	// example:
	//
	// {"key1":"value1","api_name":"PushNoticeToAndroidRequest"}
	AndroidExtParameters *string `json:"AndroidExtParameters,omitempty" xml:"AndroidExtParameters,omitempty"`
	// Set the Honor channel notification type:
	//
	// - **0**: Production notification (default).
	//
	// - **1**: Test notification.
	//
	// > Each application can send up to 1,000 test notifications per day, and this is not subject to the daily per-device push limit.
	//
	// example:
	//
	// 1
	AndroidHonorTargetUserType *int32 `json:"AndroidHonorTargetUserType,omitempty" xml:"AndroidHonorTargetUserType,omitempty"`
	// Set the Huawei instant notification parameter:
	//
	// - **0**: Send a regular Huawei notification (default).
	//
	// - **1**: Send a Huawei instant notification.
	//
	// example:
	//
	// 1
	AndroidHuaweiBusinessType *int32 `json:"AndroidHuaweiBusinessType,omitempty" xml:"AndroidHuaweiBusinessType,omitempty"`
	// JSON string of the Huawei Android Live Notification data structure [LiveNotificationPayload](https://developer.huawei.com/consumer/cn/doc/HMSCore-References/rest-live-0000001562939968#ZH-CN_TOPIC_0000001700850537__p195121620102511). For development integration, see [Huawei Live Notification Push Guide](https://help.aliyun.com/document_detail/2983768.html)
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
	// Huawei channel receipt ID. This receipt ID can be found in the receipt parameter configuration on the Huawei channel push operations platform.
	//
	// > If the default receipt configuration on the Huawei channel push operations platform is set to Alibaba Cloud receipt, this is not required. If not, we recommend configuring the default Huawei channel receipt ID in the Alibaba Cloud EMAS Mobile Push console first.
	//
	// example:
	//
	// RCP4C123456
	AndroidHuaweiReceiptId *string `json:"AndroidHuaweiReceiptId,omitempty" xml:"AndroidHuaweiReceiptId,omitempty"`
	// Set the Huawei channel notification type:
	//
	// - **0**: Production notification (default).
	//
	// - **1**: Test notification.
	//
	// > Each application can send up to 500 test notifications per day, and this is not subject to the daily per-device push limit.
	//
	// example:
	//
	// 1
	AndroidHuaweiTargetUserType *int32 `json:"AndroidHuaweiTargetUserType,omitempty" xml:"AndroidHuaweiTargetUserType,omitempty"`
	// Right-side icon URL. Currently supported:
	//
	// - Huawei EMUI (only applicable in long text mode and Inbox mode)
	//
	// - Honor Magic UI (only applicable in long text mode)
	//
	// - Proprietary channel: Android SDK 3.5.0 and later
	//
	// example:
	//
	// https://imag.example.com/image.png
	AndroidImageUrl *string `json:"AndroidImageUrl,omitempty" xml:"AndroidImageUrl,omitempty"`
	// Body content in Inbox mode. The content must be a valid JSON Array with no more than 5 elements. Currently supported:
	//
	// - Huawei: EMUI 9 and later
	//
	// - Honor: Magic UI 4.0 and later
	//
	// - Xiaomi: MIUI 10 and later
	//
	// - OPPO: ColorOS 5.0 and later
	//
	// - Proprietary channel: Android SDK 3.6.0 and later
	//
	// example:
	//
	// ["第一行","第二行"]
	AndroidInboxBody *string `json:"AndroidInboxBody,omitempty" xml:"AndroidInboxBody,omitempty"`
	// Meizu message type
	//
	// - 0: Public message (default)
	//
	// - 1: Private message
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 0
	AndroidMeizuNoticeMsgType *int32 `json:"AndroidMeizuNoticeMsgType,omitempty" xml:"AndroidMeizuNoticeMsgType,omitempty"`
	// Purpose 1: After completing the [self-classification privilege application](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835?#section3410731125514), this is used to identify the message type, determine the [notification alert method](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#ZH-CN_TOPIC_0000001149358835__p3850133955718), and speed up delivery for specific message types. For valid values, refer to the Huawei Push official documentation\\"s [Message Classification Standard](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section1076611477914). Use the "Cloud notification category value" or "Local notification category value" from the documentation table.
	//
	// Purpose 2: After [applying for special permissions](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/faq-0000001050042183#section037425218509), this is used to identify high-priority pass-through scenarios. Valid values:
	//
	// - VOIP: Audio/video calls
	//
	// - PLAY_VOICE: Voice playback
	//
	// > For "Cloud notification category values" marked as "Not applicable", the Alibaba Cloud proprietary channel is used. For "Local notification category values" marked as "Not applicable", the Huawei channel is used.
	//
	// example:
	//
	// SUBSCRIPTION
	AndroidMessageHuaweiCategory *string `json:"AndroidMessageHuaweiCategory,omitempty" xml:"AndroidMessageHuaweiCategory,omitempty"`
	// Huawei channel notification delivery priority. Valid values:
	//
	// - HIGH
	//
	// - NORMAL
	//
	// Permission application is required. See: [Application link](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/faq-0000001050042183#section037425218509).
	//
	// example:
	//
	// HIGH
	AndroidMessageHuaweiUrgency *string `json:"AndroidMessageHuaweiUrgency,omitempty" xml:"AndroidMessageHuaweiUrgency,omitempty"`
	// OPPO classifies messages into two categories: Communication & Service, and Content & Marketing.
	//
	// Communication & Service (permission application required):
	//
	// - IM: Instant messaging, audio, and video calls
	//
	// - ACCOUNT: Personal account and asset changes
	//
	// - DEVICE_REMINDER: Personal device reminders
	//
	// - ORDER: Personal order/logistics status changes
	//
	// - TODO: Personal schedule/to-do items
	//
	// - SUBSCRIPTION: Personal subscriptions
	//
	// Content & Marketing:
	//
	// - NEWS: News and information
	//
	// - CONTENT: Content recommendations
	//
	// - MARKETING: Platform promotions
	//
	// - SOCIAL: Social updates
	//
	// For details, see [OPUSH Message Classification Rules](https://open.oppomobile.com/new/developmentDoc/info?id=13189)
	//
	// example:
	//
	// MARKETING
	AndroidMessageOppoCategory *string `json:"AndroidMessageOppoCategory,omitempty" xml:"AndroidMessageOppoCategory,omitempty"`
	// OPPO channel notification bar message alert level. Valid values:
	//
	// - 1: Notification bar
	//
	// - 2: Notification bar, lock screen, ringtone, and vibration (default notification level for Communication & Service messages)
	//
	// - 16: Notification bar, lock screen, ringtone, vibration, and banner (permission application required)
	//
	// > When using the AndroidMessageOppoNotifyLevel parameter, the AndroidMessageOppoCategory parameter must also be provided.
	//
	// example:
	//
	// 1
	AndroidMessageOppoNotifyLevel *int32 `json:"AndroidMessageOppoNotifyLevel,omitempty" xml:"AndroidMessageOppoNotifyLevel,omitempty"`
	// vivo classifies messages into two categories: System messages and Operational messages.
	//
	// System messages:
	//
	// - IM: Instant messages
	//
	// - ACCOUNT: Account and assets
	//
	// - TODO: Schedule and to-do items
	//
	// - DEVICE_REMINDER: Device information
	//
	// - ORDER: Orders and logistics
	//
	// - SUBSCRIPTION: Subscription reminders
	//
	// Operational messages:
	//
	// - NEWS: News
	//
	// - CONTENT: Content recommendations
	//
	// - MARKETING: Operational campaigns
	//
	// - SOCIAL: Social updates
	//
	// > For details, see [Classification Description](https://dev.vivo.com.cn/documentCenter/doc/359#s-ef3qugc3)
	//
	// example:
	//
	// TODO
	AndroidMessageVivoCategory *string `json:"AndroidMessageVivoCategory,omitempty" xml:"AndroidMessageVivoCategory,omitempty"`
	// Huawei vendor channel notification sound. Specify the name of an audio file stored in the client project\\"s app/src/main/res/raw/ directory. The file extension is not required.
	//
	// If not set, the default ringtone is used.
	//
	// example:
	//
	// alicloud_notification_sound
	AndroidMusic *string `json:"AndroidMusic,omitempty" xml:"AndroidMusic,omitempty"`
	// Priority of the notification position in the Android notification bar. Valid values: -2, -1, 0, 1, 2.
	//
	// example:
	//
	// 0
	AndroidNotificationBarPriority *int32 `json:"AndroidNotificationBarPriority,omitempty" xml:"AndroidNotificationBarPriority,omitempty"`
	// Android custom notification bar style. Value range: 1-100.
	//
	// example:
	//
	// 2
	AndroidNotificationBarType *int32 `json:"AndroidNotificationBarType,omitempty" xml:"AndroidNotificationBarType,omitempty"`
	// The channelId of the Android app. Must match the channelId configured in the app.
	//
	// - Set the NotificationChannel parameter. For specific usage, see [FAQ: Notifications not received on Android 8.0+ devices](https://help.aliyun.com/document_detail/67398.html).
	//
	// - Since the OPPO notification private channel\\"s channel_id is the same as the app\\"s channelId, the channel_id takes this value when pushing through the OPPO channel.
	//
	// - For Huawei, FCM, and Alibaba Cloud proprietary channel pushes, the channel_id takes this value.
	//
	// example:
	//
	// 1
	AndroidNotificationChannel *string `json:"AndroidNotificationChannel,omitempty" xml:"AndroidNotificationChannel,omitempty"`
	// Message grouping. Messages in the same group display only the latest one and the total count of messages received in that group in the notification bar. All messages are not displayed and cannot be expanded. Currently supported:
	//
	// - Huawei vendor channel
	//
	// - Honor vendor channel
	//
	// - Proprietary channel: Android SDK 3.9.1 and earlier
	//
	// > The proprietary channel no longer supports this parameter on Android SDK 3.9.2 and later.
	//
	// example:
	//
	// group-1
	AndroidNotificationGroup *string `json:"AndroidNotificationGroup,omitempty" xml:"AndroidNotificationGroup,omitempty"`
	// Set the Honor notification message classification importance parameter, which determines the notification behavior on user devices. Valid values:
	//
	// - LOW: Information and marketing messages
	//
	// - NORMAL: Service and communication messages
	//
	// Application is required on the Honor platform. [Application link](https://developer.honor.com/cn/docs/11002/guides/notification-class#%E8%87%AA%E5%88%86%E7%B1%BB%E6%9D%83%E7%9B%8A%E7%94%B3%E8%AF%B7).
	//
	// example:
	//
	// LOW
	AndroidNotificationHonorChannel *string `json:"AndroidNotificationHonorChannel,omitempty" xml:"AndroidNotificationHonorChannel,omitempty"`
	// Set the Huawei notification message classification importance parameter, which determines the notification behavior on user devices. Valid values:
	//
	// - LOW: Information and marketing messages
	//
	// - NORMAL: Service and communication messages
	//
	// >- The Huawei channel now recommends using AndroidMessageHuaweiCategory for notification classification. AndroidNotificationHuaweiChannel is no longer required.
	//
	// >- Application is required on the Huawei platform. [Application link](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section893184112272).
	//
	// example:
	//
	// LOW
	AndroidNotificationHuaweiChannel *string `json:"AndroidNotificationHuaweiChannel,omitempty" xml:"AndroidNotificationHuaweiChannel,omitempty"`
	// Unique identifier for each message displayed in the notification bar. Different notification bar messages can share the same NotifyId, allowing new notifications to replace old ones.
	//
	// example:
	//
	// 100001
	AndroidNotificationNotifyId *int32 `json:"AndroidNotificationNotifyId,omitempty" xml:"AndroidNotificationNotifyId,omitempty"`
	// Message grouping. Messages in the same group are displayed collapsed in the notification bar and can be expanded. Notifications from different groups are displayed separately. Currently supported:
	//
	// - Proprietary channel: Android SDK 3.9.2 and later
	//
	// example:
	//
	// thread-1
	AndroidNotificationThreadId *string `json:"AndroidNotificationThreadId,omitempty" xml:"AndroidNotificationThreadId,omitempty"`
	// Set the vivo notification message classification. Valid values:
	//
	// - 0: Operational messages (default)
	//
	// - 1: System messages
	//
	// >- The vivo channel now recommends using AndroidMessageVivoCategory for notification classification. AndroidNotificationVivoChannel is no longer required.
	//
	// >- Application is required on the vivo platform. See: [Application link](https://dev.vivo.com.cn/documentCenter/doc/359).
	//
	// example:
	//
	// 0
	AndroidNotificationVivoChannel *string `json:"AndroidNotificationVivoChannel,omitempty" xml:"AndroidNotificationVivoChannel,omitempty"`
	// Set the channelId for Xiaomi notification types. Application is required on the Xiaomi platform. See: [Application link](https://dev.mi.com/console/doc/detail?pId=2422#_4).
	//
	// >- A single application on the Xiaomi channel can apply for up to 8 channels. Please plan ahead.
	//
	// example:
	//
	// michannel
	AndroidNotificationXiaomiChannel *string `json:"AndroidNotificationXiaomiChannel,omitempty" xml:"AndroidNotificationXiaomiChannel,omitempty"`
	// Notification alert type. Valid values:
	//
	// - **VIBRATE**: Vibration (default)
	//
	// - **SOUND**: Sound
	//
	// - **BOTH**: Sound and vibration
	//
	// - **NONE**: Silent
	//
	// example:
	//
	// VIBRATE
	AndroidNotifyType *string `json:"AndroidNotifyType,omitempty" xml:"AndroidNotifyType,omitempty"`
	// Action after clicking the notification. Valid values:
	//
	// - APPLICATION: Open the app (default)
	//
	// - ACTIVITY: Open an Android Activity
	//
	// - URL: Open a URL
	//
	// - NONE: No navigation
	//
	// example:
	//
	// APPLICATION
	AndroidOpenType *string `json:"AndroidOpenType,omitempty" xml:"AndroidOpenType,omitempty"`
	// The URL to open after Android receives the push. Only applicable when PushTask.N.AndroidOpenType="URL".
	//
	// example:
	//
	// https://xxxx.xxx
	AndroidOpenUrl *string `json:"AndroidOpenUrl,omitempty" xml:"AndroidOpenUrl,omitempty"`
	// JSON string of the OPPO Fluid Cloud intent deletion data structure [data](https://open.oppomobile.com/documentation/page/info?id=13578). This parameter is ignored when the AndroidOppoIntelligentIntent parameter is already provided. For development integration, see [OPPO Fluid Cloud Push Guide](https://help.aliyun.com/document_detail/2997310.html)
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
	// JSON string of the OPPO Fluid Cloud intent sharing data structure [IntelligentIntent](https://open.oppomobile.com/documentation/page/info?id=13565). For development integration, see [OPPO Fluid Cloud Push Guide](https://help.aliyun.com/document_detail/2997310.html)
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
	// Set the OPPO Fluid Cloud push environment:
	//
	// - **0**: Production environment (default).
	//
	// - **1**: Test environment.
	//
	// > The OPPO Fluid Cloud test environment requires client-side setup. See [Environment Setup](https://open.oppomobile.com/documentation/page/info?id=13590).
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
	// 	Warning: The OPPO private message template feature is no longer supported by MaasPush. To use this feature, please use the Push, PushV2, or MassPushV2 API instead.
	//
	// example:
	//
	// 687557242b1634hzef3zd5013
	AndroidOppoPrivateMsgTemplateId *string `json:"AndroidOppoPrivateMsgTemplateId,omitempty" xml:"AndroidOppoPrivateMsgTemplateId,omitempty"`
	// Deprecated
	//
	// OPPO private message template title parameters
	AndroidOppoPrivateTitleParameters map[string]*string `json:"AndroidOppoPrivateTitleParameters,omitempty" xml:"AndroidOppoPrivateTitleParameters,omitempty"`
	// Specify the Activity to navigate to when the notification is clicked.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidPopupActivity *string `json:"AndroidPopupActivity,omitempty" xml:"AndroidPopupActivity,omitempty"`
	// Body content in auxiliary popup mode. Required when the AndroidPopupActivity parameter is not empty.
	//
	// Length limit: 200 characters (both Chinese and English characters count as one character each).
	//
	// When using vendor channels, you must also comply with the vendor channel restrictions. For details, see [Android Auxiliary Channel Push Limits](https://help.aliyun.com/document_detail/165253.html).
	//
	// example:
	//
	// hello
	AndroidPopupBody *string `json:"AndroidPopupBody,omitempty" xml:"AndroidPopupBody,omitempty"`
	// Title content in auxiliary popup mode. Required when the AndroidPopupActivity parameter is not empty.
	//
	// Length limit: 50 characters (both Chinese and English characters count as one character each).
	//
	// When using vendor channels, you must also comply with the vendor channel restrictions. For details, see [Android Auxiliary Channel Push Limits](https://help.aliyun.com/document_detail/165253.html).
	//
	// example:
	//
	// hello
	AndroidPopupTitle *string `json:"AndroidPopupTitle,omitempty" xml:"AndroidPopupTitle,omitempty"`
	// When the push type is message and the device is offline, this push will use the auxiliary popup feature. Defaults to false. Only takes effect when PushType=MESSAGE.
	//
	// If the message-to-notification conversion push is successful, the notification displays the data set by the server\\"s AndroidPopupTitle and AndroidPopupBody parameter values. The data obtained when clicking the notification in the auxiliary popup\\"s onSysNoticeOpened method is the server-set Title and Body parameter values.
	//
	// example:
	//
	// true
	AndroidRemind *bool `json:"AndroidRemind,omitempty" xml:"AndroidRemind,omitempty"`
	// Notification style. Valid values:
	//
	// - **0**: Standard mode (default)
	//
	// - **1**: Long text mode (supported by Huawei, Honor, Xiaomi, OPPO, Meizu, and proprietary channels)
	//
	// - **2**: Big picture mode (supported by proprietary channel, not supported on Xiaomi devices)
	//
	// - **3**: List mode (supported by Huawei, Honor, Xiaomi, OPPO, and proprietary channels)
	//
	// > This parameter is required when using non-standard modes.
	//
	// example:
	//
	// 1
	AndroidRenderStyle *string `json:"AndroidRenderStyle,omitempty" xml:"AndroidRenderStyle,omitempty"`
	// Set the vendor channel notification type:
	//
	// - **0**: Production notification (default).
	//
	// - **1**: Test notification.
	//
	// >- Configuring this parameter is equivalent to simultaneously configuring the AndroidHuaweiTargetUserType, AndroidHonorTargetUserType, AndroidVivoPushMode, and AndroidOppoIntentEnv parameters. The corresponding parameter for a specific vendor channel can override this parameter.
	//
	// >- Currently supported: Huawei channel, Honor channel, vivo channel, and OPPO Fluid Cloud.
	//
	// example:
	//
	// 1
	AndroidTargetUserType *int32 `json:"AndroidTargetUserType,omitempty" xml:"AndroidTargetUserType,omitempty"`
	// JSON string of the vivo Atomic Island data structure [liveMessage](https://dev.vivo.com.cn/documentCenter/doc/896#s-fdagzbd4). For development integration, see [vivo Atomic Island Push Guide](https://help.aliyun.com/zh/document_detail/3030718.html)
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
	// Set the vivo channel notification type:
	//
	// - **0**: Production push (default).
	//
	// - **1**: Test push.
	//
	// > For test pushes, configure test devices in the vivo console in advance. The test device RegId can be obtained by searching for "onReceiveRegId regId" in the device startup logs.
	//
	// example:
	//
	// 1
	AndroidVivoPushMode *int32 `json:"AndroidVivoPushMode,omitempty" xml:"AndroidVivoPushMode,omitempty"`
	// vivo channel receipt ID. This receipt ID can be found in the app information section of the push service on the vivo open platform.
	//
	// > If the default receipt configuration on the vivo open platform is set to Alibaba Cloud receipt, this is not required. If not, we recommend configuring the default vivo channel receipt ID in the Alibaba Cloud EMAS Mobile Push console first.
	//
	// example:
	//
	// 123
	AndroidVivoReceiptId *string `json:"AndroidVivoReceiptId,omitempty" xml:"AndroidVivoReceiptId,omitempty"`
	// Deprecated
	//
	// This parameter has been deprecated. All third-party auxiliary popups are now supported by the new parameter **AndroidPopupActivity**.
	//
	// example:
	//
	// 无
	AndroidXiaoMiActivity *string `json:"AndroidXiaoMiActivity,omitempty" xml:"AndroidXiaoMiActivity,omitempty"`
	// Deprecated
	//
	// This parameter has been deprecated. All third-party auxiliary popups are now supported by the new parameter **AndroidPopupBody**.
	//
	// example:
	//
	// 无
	AndroidXiaoMiNotifyBody *string `json:"AndroidXiaoMiNotifyBody,omitempty" xml:"AndroidXiaoMiNotifyBody,omitempty"`
	// Deprecated
	//
	// This parameter has been deprecated. All third-party auxiliary popups are now supported by the new parameter **AndroidPopupTitle**.
	//
	// example:
	//
	// 无
	AndroidXiaoMiNotifyTitle *string `json:"AndroidXiaoMiNotifyTitle,omitempty" xml:"AndroidXiaoMiNotifyTitle,omitempty"`
	// Deprecated
	//
	// This parameter has been deprecated. Since August 2023, Xiaomi has officially discontinued support for dynamically setting small icons, right-side icons, and big pictures during push on new devices/systems.
	//
	// example:
	//
	// https://f6.market.xiaomi.com/download/MiPass/aaa/bbb.png
	AndroidXiaomiBigPictureUrl *string `json:"AndroidXiaomiBigPictureUrl,omitempty" xml:"AndroidXiaomiBigPictureUrl,omitempty"`
	// JSON string of the Xiaomi Super Island data structure [miui.focus.param](https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131). For development integration, see [Xiaomi Super Island Push Guide](https://help.aliyun.com/zh/document_detail/3037956.html)
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
	// JSON string of the Xiaomi Super Island images [miui.focus.pic_xxx](https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131). For development integration, see [Xiaomi Super Island Push Guide](https://help.aliyun.com/zh/document_detail/3037956.html)
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
	// This parameter has been deprecated. Since August 2023, Xiaomi has officially discontinued support for dynamically setting small icons, right-side icons, and big pictures during push on new devices/systems.
	//
	// example:
	//
	// https://imag.example.com/image.png
	AndroidXiaomiImageUrl *string `json:"AndroidXiaomiImageUrl,omitempty" xml:"AndroidXiaomiImageUrl,omitempty"`
	// Xiaomi private message template ID
	//
	// example:
	//
	// P10645
	AndroidXiaomiTemplateId *string `json:"AndroidXiaomiTemplateId,omitempty" xml:"AndroidXiaomiTemplateId,omitempty"`
	// Xiaomi private message template parameters, JSON string
	//
	// example:
	//
	// {"keywords1":"Tom","keywords2":"phone"}
	AndroidXiaomiTemplateParams *string `json:"AndroidXiaomiTemplateParams,omitempty" xml:"AndroidXiaomiTemplateParams,omitempty"`
	// Content of the notification/message for Android and HarmonyOS pushes; iOS message/notification content. The push content size is limited. See [Product Limits](https://help.aliyun.com/document_detail/92832.html).
	//
	// example:
	//
	// hello
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// Device type. Valid values:
	//
	// - HARMONY: HarmonyOS device
	//
	// - iOS: iOS device
	//
	// - ANDROID: Android device
	//
	// - ALL: When the AppKey is for a legacy dual-platform app, this pushes to both Android and iOS devices simultaneously. When the AppKey is for a new single-platform app, the effect is the same as specifying the device type corresponding to the app type.
	//
	// This parameter is required.
	//
	// example:
	//
	// HARMONY
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// Expiration time of offline messages/notifications. Used together with StoreOffline. Expired messages will no longer be sent. The maximum retention period is 72 hours. The default is 72 hours.
	//
	// The time format follows the ISO 8601 standard and must use UTC time in the format YYYY-MM-DDThh:mm:ssZ. The expiration time cannot be earlier than the current time or the scheduled push time plus 3 seconds (`ExpireTime > PushTime + 3 seconds`). The 3-second buffer accounts for network and system latency. We recommend at least 1 minute for unicast pushes and at least 10 minutes for broadcast and batch pushes.
	//
	// example:
	//
	// 2019-02-20T00:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The action corresponding to the built-in page ability of the app.
	//
	// 	Notice: When HarmonyActionType is APP_CUSTOM_PAGE, at least one of HarmonyUri and HarmonyAction must be provided.
	//
	// For details, see HarmonyOS official documentation [ClickAction.action](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section152462191216)
	//
	// example:
	//
	// com.example.action
	HarmonyAction *string `json:"HarmonyAction,omitempty" xml:"HarmonyAction,omitempty"`
	// Action after clicking the notification. Valid values:
	//
	// - APP_HOME_PAGE: Open the app home page
	//
	// - APP_CUSTOM_PAGE: Open a custom app page
	//
	// example:
	//
	// APP_HOME_PAGE
	HarmonyActionType *string `json:"HarmonyActionType,omitempty" xml:"HarmonyActionType,omitempty"`
	// HarmonyOS app badge increment number. See [HarmonyOS badge addNum field description](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section266310382145).</br>
	//
	// Supported from HarmonyOS SDK 1.2.0.
	//
	// example:
	//
	// 1
	HarmonyBadgeAddNum *int32 `json:"HarmonyBadgeAddNum,omitempty" xml:"HarmonyBadgeAddNum,omitempty"`
	// HarmonyOS app badge set number. See [HarmonyOS badge setNum field description](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section266310382145).</br>
	//
	// Supported from HarmonyOS SDK 1.2.0.
	//
	// example:
	//
	// 1
	HarmonyBadgeSetNum *int32 `json:"HarmonyBadgeSetNum,omitempty" xml:"HarmonyBadgeSetNum,omitempty"`
	// Notification message category. After completing the notification message self-classification privilege application, this is used to identify the message type. Different notification message types affect the display and alert methods. Valid values:
	//
	// - IM: Instant messaging
	//
	// - VOIP: Audio/video calls
	//
	// - SUBSCRIPTION: Subscriptions
	//
	// - TRAVEL: Travel
	//
	// - HEALTH: Health
	//
	// - WORK: Work task reminders
	//
	// - ACCOUNT: Account updates
	//
	// - EXPRESS: Orders & logistics
	//
	// - FINANCE: Finance
	//
	// - DEVICE_REMINDER: Device reminders
	//
	// - MAIL: Email
	//
	// - CUSTOMER_SERVICE: Customer service messages
	//
	// - MARKETING: News, content recommendations, social updates, product promotions, financial updates, lifestyle information, surveys, feature recommendations, and operational campaigns (only marks the content, does not speed up message delivery), collectively referred to as information and marketing messages
	//
	// For details, see HarmonyOS official documentation [Notification.category](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117)
	//
	// example:
	//
	// IM
	HarmonyCategory *string `json:"HarmonyCategory,omitempty" xml:"HarmonyCategory,omitempty"`
	// Set notification extended properties. This property does not take effect when the push type PushType is set to MESSAGE.
	//
	// This parameter must be passed in JSON map format; otherwise, parsing errors will occur.
	//
	// example:
	//
	// {"key1":"value1","api_name":"PushNoticeToAndroidRequest"}
	HarmonyExtParameters *string `json:"HarmonyExtParameters,omitempty" xml:"HarmonyExtParameters,omitempty"`
	// Extra data for notification extension messages.</br>
	//
	// Valid when sending HarmonyOS notification extension messages.</br>
	//
	// Conceptually equivalent to the extraData field of HarmonyOS notification extension messages. For the specific definition, see [HarmonyOS ExtensionPayload Description](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section161192514234).</br>
	//
	// Supported from HarmonyOS SDK 1.2.0.
	//
	// example:
	//
	// 示例额外数据
	HarmonyExtensionExtraData *string `json:"HarmonyExtensionExtraData,omitempty" xml:"HarmonyExtensionExtraData,omitempty"`
	// When PushType is NOTICE, whether this is a HarmonyOS notification extension message.
	//
	// - true: Send a notification extension message
	//
	// - false: Send a regular notification (default)
	//
	// Notification extension messages require permission to be applied for on the HarmonyOS side before sending. For details, see the HarmonyOS documentation [Send Notification Extension Messages](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V5/push-send-extend-noti-V5).</br>
	//
	// Supported from HarmonyOS SDK 1.2.0.
	//
	// example:
	//
	// true
	HarmonyExtensionPush *bool `json:"HarmonyExtensionPush,omitempty" xml:"HarmonyExtensionPush,omitempty"`
	// URL for the large icon on the right side of the notification. The URL must use the HTTPS protocol.
	//
	// > Supported image formats: png, jpg, jpeg, heif, gif, bmp. Image dimensions must satisfy height × width < 25,000 pixels.
	//
	// For details, see HarmonyOS official documentation [Notification.image](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117)
	//
	// example:
	//
	// https://example.com/xxx.png
	HarmonyImageUrl *string `json:"HarmonyImageUrl,omitempty" xml:"HarmonyImageUrl,omitempty"`
	// Content for multi-line text style. Required when HarmonyRenderStyle is MULTI_LINE. Up to 3 content items are supported.
	//
	// example:
	//
	// ["1.content1","2.content2","3.content3"]
	HarmonyInboxContent *string `json:"HarmonyInboxContent,omitempty" xml:"HarmonyInboxContent,omitempty"`
	// JSON string of the HarmonyOS Live View data structure [LiveViewPayload](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V13/push-scenariozed-api-request-param-V13#section66881469306). For development integration, see [HarmonyOS Live View Push Guide](https://help.aliyun.com/document_detail/2982112.html)
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
	// Use the specified type of notification slot. Only valid when the Alibaba Cloud proprietary channel is online.
	//
	// - SOCIAL_COMMUNICATION: Social communication.
	//
	// - SERVICE_INFORMATION: Service reminders.
	//
	// - CONTENT_INFORMATION: Content information.
	//
	// - CUSTOMER_SERVICE: Customer service messages. This type is used for customer service messages between users and merchants, and must be initiated by the user.
	//
	// - OTHER_TYPES: Others.
	//
	// For details, see HarmonyOS official documentation [SlotType](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/js-apis-notificationmanager-V5#slottype)
	//
	// example:
	//
	// SOCIAL_COMMUNICATION
	HarmonyNotificationSlotType *string `json:"HarmonyNotificationSlotType,omitempty" xml:"HarmonyNotificationSlotType,omitempty"`
	// Unique identifier for each message displayed in the notification. When not provided, the push service automatically generates a unique identifier for each message. Different notification messages can share the same notifyId, enabling new messages to replace old ones.
	//
	// For details, see HarmonyOS official documentation [Notification.notifyId](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117)
	//
	// example:
	//
	// 0
	HarmonyNotifyId *int32 `json:"HarmonyNotifyId,omitempty" xml:"HarmonyNotifyId,omitempty"`
	// HarmonyOS channel receipt ID. This receipt ID can be found in the receipt parameter configuration on the HarmonyOS channel push operations platform.
	//
	// > If the default receipt configuration on the HarmonyOS channel push operations platform is set to Alibaba Cloud receipt, this is not required. If not, we recommend configuring the default HarmonyOS channel receipt ID in the Alibaba Cloud EMAS Mobile Push console first.
	//
	// For details, see HarmonyOS official documentation [pushOptions.receiptId](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section418321011212)
	//
	// example:
	//
	// RCPB***DFD5
	HarmonyReceiptId *string `json:"HarmonyReceiptId,omitempty" xml:"HarmonyReceiptId,omitempty"`
	// When the push type is message and the device is offline, this push will use the auxiliary popup feature. Defaults to false. Only takes effect when PushType=MESSAGE.
	//
	// If the message-to-notification conversion push is successful, the notification displays the data set by the server\\"s HarmonyRemindTitle and HarmonyRemindBody parameter values.
	//
	// example:
	//
	// false
	HarmonyRemind *bool `json:"HarmonyRemind,omitempty" xml:"HarmonyRemind,omitempty"`
	// HarmonyOS notification content used when converting HarmonyOS messages to notifications. Only valid when HarmonyRemind is true.
	//
	// example:
	//
	// 您有一条新消息，请查收
	HarmonyRemindBody *string `json:"HarmonyRemindBody,omitempty" xml:"HarmonyRemindBody,omitempty"`
	// HarmonyOS notification title used when converting HarmonyOS messages to notifications. Only valid when HarmonyRemind is true.
	//
	// example:
	//
	// 新消息
	HarmonyRemindTitle *string `json:"HarmonyRemindTitle,omitempty" xml:"HarmonyRemindTitle,omitempty"`
	// Notification message style:
	//
	// - NORMAL: Standard notification (default)
	//
	// - MULTI_LINE: Multi-line text style
	//
	// example:
	//
	// NORMAL
	HarmonyRenderStyle *string `json:"HarmonyRenderStyle,omitempty" xml:"HarmonyRenderStyle,omitempty"`
	// Test message flag:
	//
	// - false: Normal message (default)
	//
	// - true: Test message
	//
	// For details, see HarmonyOS official documentation [pushOptions.testMessage](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section418321011212)
	//
	// example:
	//
	// true
	HarmonyTestMessage *bool `json:"HarmonyTestMessage,omitempty" xml:"HarmonyTestMessage,omitempty"`
	// The URI corresponding to the built-in page ability of the app.
	//
	// 	Notice: When HarmonyActionType is APP_CUSTOM_PAGE, at least one of HarmonyUri and HarmonyAction must be provided. When multiple Abilities exist, provide different action and URI values for each Ability. The action is prioritized when looking up the corresponding built-in app page.
	//
	// For details, see HarmonyOS official documentation [ClickAction.uri](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section152462191216)
	//
	// example:
	//
	// https://www.example.com:8080/push/example
	HarmonyUri *string `json:"HarmonyUri,omitempty" xml:"HarmonyUri,omitempty"`
	// Custom identifier for the push task. When JobKey is not empty, this field will be included in the receipt logs. For receipt log details, see [Receipt Logs](https://help.aliyun.com/document_detail/434651.html).
	//
	// example:
	//
	// 123
	JobKey *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	// Used for scheduled sending. If not set, the default is immediate sending.
	//
	// The time format follows the ISO 8601 standard and must use UTC time in the format YYYY-MM-DDThh:mm:ssZ.
	//
	// example:
	//
	// 2019-02-20T00:00:00Z
	PushTime *string `json:"PushTime,omitempty" xml:"PushTime,omitempty"`
	// Push type. Valid values:
	//
	// - MESSAGE: indicates a message.
	//
	// - NOTICE: indicates a notification.
	//
	// This parameter is required.
	//
	// example:
	//
	// MESSAGE
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	// Specify sending channels. Valid values:
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
	// >- If this parameter is not configured, all channels are available.
	//
	// >- If this parameter is configured, only the specified channels are used.
	//
	// >- If the configured channels conflict with the sending strategy (e.g., iOS notifications only go through the APNs channel, but this parameter does not include apns), the push will not be sent.
	//
	// >- If gcm is configured, both Google GCM and FCM channels can be used. If fcm is configured, only the Google FCM channel can be used.
	//
	// example:
	//
	// accs,huawei,xiaomi
	SendChannels *string `json:"SendChannels,omitempty" xml:"SendChannels,omitempty"`
	// Deprecated
	//
	// This parameter has been deprecated.
	//
	// example:
	//
	// 0
	SendSpeed *int32 `json:"SendSpeed,omitempty" xml:"SendSpeed,omitempty"`
	// Whether to store offline messages/notifications. StoreOffline defaults to false.
	//
	// If stored, when the user is offline during push, the message will be resent when the user comes online within the expiration time (ExpireTime). ExpireTime defaults to 72 hours. iOS notifications are delivered through the APNs channel and are not affected by StoreOffline.
	//
	// example:
	//
	// true
	StoreOffline *bool `json:"StoreOffline,omitempty" xml:"StoreOffline,omitempty"`
	// Push target. Valid values:
	//
	// - DEVICE: push by device.
	//
	// - ACCOUNT: push by account.
	//
	// - ALIAS: push by alias.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEVICE
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// Set based on Target. Multiple values are separated by commas. If the limit is exceeded, you need to split the push into multiple calls.
	//
	// - Target=DEVICE: values such as `deviceid1,deviceid2` (up to 1,000 supported).
	//
	// - Target=ACCOUNT: values such as `account1,account2` (up to 1,000 supported).
	//
	// - Target=ALIAS: values such as `alias1,alias2` (up to 1,000 supported).
	//
	// This parameter is required.
	//
	// example:
	//
	// deviceid1,deviceid2
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
	// Title of the notification/message during push. Length limit: 200 bytes.
	//
	// Required for Android and HarmonyOS pushes. Optional for iOS push notifications. If provided:
	//
	// 	- iOS 10+: the notification displays the title.
	//
	// 	- iOS 8.2 <= iOS version < iOS 10: replaces the notification app name.
	//
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Whether to automatically truncate titles and content that are too long.
	//
	// Note: This only applies to vendor channels that explicitly limit title and content length. It does not apply to channels like APNs, Huawei, and Honor that do not limit title and content individually but only limit the total request body size.
	//
	// example:
	//
	// false
	Trim *bool `json:"Trim,omitempty" xml:"Trim,omitempty"`
	// iOS notifications are sent through the APNs center. You need to specify the corresponding environment information.
	//
	// - DEV: Development environment, applicable to apps installed and debugged directly via Xcode.
	//
	// - PRODUCT: Production environment, applicable to apps distributed through App Store, TestFlight, Ad Hoc, and enterprise distribution.
	//
	// example:
	//
	// DEV
	IOSApnsEnv *string `json:"iOSApnsEnv,omitempty" xml:"iOSApnsEnv,omitempty"`
	// iOS app icon badge number in the upper-right corner.
	//
	// > If iOSBadgeAutoIncrement is set to True, this field must be empty.
	//
	// example:
	//
	// 0
	IOSBadge *int32 `json:"iOSBadge,omitempty" xml:"iOSBadge,omitempty"`
	// Whether to enable badge auto-increment. Defaults to false.
	//
	// > When this is set to true, iOSBadge must be empty.
	//
	// The badge auto-increment feature is maintained by the push server for each device\\"s badge count. Users must use SDK version V1.9.5 or later and actively sync the badge number to the server.
	//
	// example:
	//
	// true
	IOSBadgeAutoIncrement *bool `json:"iOSBadgeAutoIncrement,omitempty" xml:"iOSBadgeAutoIncrement,omitempty"`
	// Extended properties of iOS notifications.
	//
	// On iOS 10+, you can specify the resource URL for rich media push notifications here: `{"attachment": "https://xxxx.xxx/notification_pic.png"} `. This parameter must be passed in JSON map format; otherwise, parsing errors will occur.
	//
	// example:
	//
	// {"attachment": "https://xxxx.xxx/notification_pic.png"}
	IOSExtParameters *string `json:"iOSExtParameters,omitempty" xml:"iOSExtParameters,omitempty"`
	// Interruption level. Valid values:
	//
	// - passive: The system adds the notification to the notification list without lighting up the screen or playing sound.
	//
	// - active: The system immediately displays the notification, lights up the screen, and can play sound.
	//
	// - time-sensitive: The system immediately presents the notification, lights up the screen, and can play sound, but does not break through system notification controls.
	//
	// - critical: The system immediately displays the notification, lights up the screen, and plays sound bypassing the mute switch.
	//
	// example:
	//
	// active
	IOSInterruptionLevel *string `json:"iOSInterruptionLevel,omitempty" xml:"iOSInterruptionLevel,omitempty"`
	// JSON string. Static pass-through parameters for Dynamic Island push. Contains static user-defined information such as product ID and order information.
	//
	// > Required when iOSLiveActivityEvent is start.
	//
	// example:
	//
	// {"orderId": "12345", "product": "Shoes"}
	IOSLiveActivityAttributes *string `json:"iOSLiveActivityAttributes,omitempty" xml:"iOSLiveActivityAttributes,omitempty"`
	// The type of Live Activity to start.
	//
	// > Required when iOSLiveActivityEvent is start.
	//
	// example:
	//
	// OrderActivityAttributes
	IOSLiveActivityAttributesType *string `json:"iOSLiveActivityAttributesType,omitempty" xml:"iOSLiveActivityAttributesType,omitempty"`
	// Dynamic pass-through parameters for Dynamic Island push. Contains real-time update information such as price and inventory changes.
	//
	// example:
	//
	// {"status": "delivered", "estimatedArrival": "2023-12-31T12:00:00Z"}
	IOSLiveActivityContentState *string `json:"iOSLiveActivityContentState,omitempty" xml:"iOSLiveActivityContentState,omitempty"`
	// The ended Live Activity will remain on the lock screen until the specified time, up to a maximum of 4 hours.
	//
	// example:
	//
	// 1743131967
	IOSLiveActivityDismissalDate *int64 `json:"iOSLiveActivityDismissalDate,omitempty" xml:"iOSLiveActivityDismissalDate,omitempty"`
	// Start, update, or end a Live Activity.
	//
	// - Enum: start | update | end
	//
	// example:
	//
	// start
	IOSLiveActivityEvent *string `json:"iOSLiveActivityEvent,omitempty" xml:"iOSLiveActivityEvent,omitempty"`
	// The Live Activity ID reported by the device to the user\\"s server. A unique identifier for the Live Activity.
	//
	// example:
	//
	// 66B94673-B32E-4CA7-863C-3E523054FD46
	IOSLiveActivityId *string `json:"iOSLiveActivityId,omitempty" xml:"iOSLiveActivityId,omitempty"`
	// Timestamp in seconds, marking the expiration time of the activity content.
	//
	// example:
	//
	// 1743131967
	IOSLiveActivityStaleDate *int64 `json:"iOSLiveActivityStaleDate,omitempty" xml:"iOSLiveActivityStaleDate,omitempty"`
	// iOS notification sound. Specify the name of an audio file stored in the app bundle or the sandbox Library/Sounds directory. See: How to set notification sound for iOS push.
	//
	// If set to an empty string (""), the notification is silent. If not set, the default system alert sound is used.
	//
	// example:
	//
	// ””
	IOSMusic *string `json:"iOSMusic,omitempty" xml:"iOSMusic,omitempty"`
	// iOS notification processing extension flag (iOS 10+). If set to true, APNs push notifications can reach the Extension for processing before being displayed. Must be set to true for silent notifications.
	//
	// example:
	//
	// true
	IOSMutableContent *bool `json:"iOSMutableContent,omitempty" xml:"iOSMutableContent,omitempty"`
	// Specify the iOS notification Category (iOS 10+).
	//
	// example:
	//
	// ios
	IOSNotificationCategory *string `json:"iOSNotificationCategory,omitempty" xml:"iOSNotificationCategory,omitempty"`
	// When a device receives messages with the same CollapseId, they are merged into one. When the device is offline, consecutive messages with the same CollapseId will show only one notification in the notification bar. Supported on iOS 10+.
	//
	// example:
	//
	// ZD2011
	IOSNotificationCollapseId *string `json:"iOSNotificationCollapseId,omitempty" xml:"iOSNotificationCollapseId,omitempty"`
	// Groups iOS remote notifications using this property, marking the group identifier for collapsed notifications. Only supported on iOS 12.0+.
	//
	// example:
	//
	// abc
	IOSNotificationThreadId *string `json:"iOSNotificationThreadId,omitempty" xml:"iOSNotificationThreadId,omitempty"`
	// Summary highlight score. Value range: a floating-point number in [0,1\\].
	//
	// example:
	//
	// 0.01
	IOSRelevanceScore *float64 `json:"iOSRelevanceScore,omitempty" xml:"iOSRelevanceScore,omitempty"`
	// When the device is offline during message push (i.e., the persistent connection channel to the Mobile Push server is disconnected), this push will be delivered as a notification through Apple\\"s APNs channel once.
	//
	// > Offline message-to-notification conversion only applies to the production environment.
	//
	// example:
	//
	// true
	IOSRemind *bool `json:"iOSRemind,omitempty" xml:"iOSRemind,omitempty"`
	// iOS notification content used when converting iOS messages to notifications. Only valid when iOSApnsEnv=PRODUCT and iOSRemind is true.
	//
	// example:
	//
	// ios通知body
	IOSRemindBody *string `json:"iOSRemindBody,omitempty" xml:"iOSRemindBody,omitempty"`
	// Whether to enable iOS silent notification.
	//
	// example:
	//
	// true
	IOSSilentNotification *bool `json:"iOSSilentNotification,omitempty" xml:"iOSSilentNotification,omitempty"`
	// iOS notification subtitle content (iOS 10+).
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
