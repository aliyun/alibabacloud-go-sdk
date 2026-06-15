// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidActivity(v string) *PushRequest
	GetAndroidActivity() *string
	SetAndroidBadgeAddNum(v int32) *PushRequest
	GetAndroidBadgeAddNum() *int32
	SetAndroidBadgeClass(v string) *PushRequest
	GetAndroidBadgeClass() *string
	SetAndroidBadgeSetNum(v int32) *PushRequest
	GetAndroidBadgeSetNum() *int32
	SetAndroidBigBody(v string) *PushRequest
	GetAndroidBigBody() *string
	SetAndroidBigPictureUrl(v string) *PushRequest
	GetAndroidBigPictureUrl() *string
	SetAndroidBigTitle(v string) *PushRequest
	GetAndroidBigTitle() *string
	SetAndroidExtParameters(v string) *PushRequest
	GetAndroidExtParameters() *string
	SetAndroidHonorTargetUserType(v int32) *PushRequest
	GetAndroidHonorTargetUserType() *int32
	SetAndroidHuaweiBusinessType(v int32) *PushRequest
	GetAndroidHuaweiBusinessType() *int32
	SetAndroidHuaweiLiveNotificationPayload(v string) *PushRequest
	GetAndroidHuaweiLiveNotificationPayload() *string
	SetAndroidHuaweiReceiptId(v string) *PushRequest
	GetAndroidHuaweiReceiptId() *string
	SetAndroidHuaweiTargetUserType(v int32) *PushRequest
	GetAndroidHuaweiTargetUserType() *int32
	SetAndroidImageUrl(v string) *PushRequest
	GetAndroidImageUrl() *string
	SetAndroidInboxBody(v string) *PushRequest
	GetAndroidInboxBody() *string
	SetAndroidMeizuNoticeMsgType(v int32) *PushRequest
	GetAndroidMeizuNoticeMsgType() *int32
	SetAndroidMessageHuaweiCategory(v string) *PushRequest
	GetAndroidMessageHuaweiCategory() *string
	SetAndroidMessageHuaweiUrgency(v string) *PushRequest
	GetAndroidMessageHuaweiUrgency() *string
	SetAndroidMessageOppoCategory(v string) *PushRequest
	GetAndroidMessageOppoCategory() *string
	SetAndroidMessageOppoNotifyLevel(v int32) *PushRequest
	GetAndroidMessageOppoNotifyLevel() *int32
	SetAndroidMessageVivoCategory(v string) *PushRequest
	GetAndroidMessageVivoCategory() *string
	SetAndroidMusic(v string) *PushRequest
	GetAndroidMusic() *string
	SetAndroidNotificationBarPriority(v int32) *PushRequest
	GetAndroidNotificationBarPriority() *int32
	SetAndroidNotificationBarType(v int32) *PushRequest
	GetAndroidNotificationBarType() *int32
	SetAndroidNotificationChannel(v string) *PushRequest
	GetAndroidNotificationChannel() *string
	SetAndroidNotificationGroup(v string) *PushRequest
	GetAndroidNotificationGroup() *string
	SetAndroidNotificationHonorChannel(v string) *PushRequest
	GetAndroidNotificationHonorChannel() *string
	SetAndroidNotificationHuaweiChannel(v string) *PushRequest
	GetAndroidNotificationHuaweiChannel() *string
	SetAndroidNotificationNotifyId(v int32) *PushRequest
	GetAndroidNotificationNotifyId() *int32
	SetAndroidNotificationThreadId(v string) *PushRequest
	GetAndroidNotificationThreadId() *string
	SetAndroidNotificationVivoChannel(v string) *PushRequest
	GetAndroidNotificationVivoChannel() *string
	SetAndroidNotificationXiaomiChannel(v string) *PushRequest
	GetAndroidNotificationXiaomiChannel() *string
	SetAndroidNotifyType(v string) *PushRequest
	GetAndroidNotifyType() *string
	SetAndroidOpenType(v string) *PushRequest
	GetAndroidOpenType() *string
	SetAndroidOpenUrl(v string) *PushRequest
	GetAndroidOpenUrl() *string
	SetAndroidOppoDeleteIntentData(v string) *PushRequest
	GetAndroidOppoDeleteIntentData() *string
	SetAndroidOppoIntelligentIntent(v string) *PushRequest
	GetAndroidOppoIntelligentIntent() *string
	SetAndroidOppoIntentEnv(v int32) *PushRequest
	GetAndroidOppoIntentEnv() *int32
	SetAndroidOppoPrivateContentParameters(v map[string]*string) *PushRequest
	GetAndroidOppoPrivateContentParameters() map[string]*string
	SetAndroidOppoPrivateMsgTemplateId(v string) *PushRequest
	GetAndroidOppoPrivateMsgTemplateId() *string
	SetAndroidOppoPrivateTitleParameters(v map[string]*string) *PushRequest
	GetAndroidOppoPrivateTitleParameters() map[string]*string
	SetAndroidPopupActivity(v string) *PushRequest
	GetAndroidPopupActivity() *string
	SetAndroidPopupBody(v string) *PushRequest
	GetAndroidPopupBody() *string
	SetAndroidPopupTitle(v string) *PushRequest
	GetAndroidPopupTitle() *string
	SetAndroidRemind(v bool) *PushRequest
	GetAndroidRemind() *bool
	SetAndroidRenderStyle(v int32) *PushRequest
	GetAndroidRenderStyle() *int32
	SetAndroidTargetUserType(v int32) *PushRequest
	GetAndroidTargetUserType() *int32
	SetAndroidVivoLiveMessage(v string) *PushRequest
	GetAndroidVivoLiveMessage() *string
	SetAndroidVivoPushMode(v int32) *PushRequest
	GetAndroidVivoPushMode() *int32
	SetAndroidVivoReceiptId(v string) *PushRequest
	GetAndroidVivoReceiptId() *string
	SetAndroidXiaoMiActivity(v string) *PushRequest
	GetAndroidXiaoMiActivity() *string
	SetAndroidXiaoMiNotifyBody(v string) *PushRequest
	GetAndroidXiaoMiNotifyBody() *string
	SetAndroidXiaoMiNotifyTitle(v string) *PushRequest
	GetAndroidXiaoMiNotifyTitle() *string
	SetAndroidXiaomiBigPictureUrl(v string) *PushRequest
	GetAndroidXiaomiBigPictureUrl() *string
	SetAndroidXiaomiFocusParam(v string) *PushRequest
	GetAndroidXiaomiFocusParam() *string
	SetAndroidXiaomiFocusPics(v string) *PushRequest
	GetAndroidXiaomiFocusPics() *string
	SetAndroidXiaomiImageUrl(v string) *PushRequest
	GetAndroidXiaomiImageUrl() *string
	SetAndroidXiaomiTemplateId(v string) *PushRequest
	GetAndroidXiaomiTemplateId() *string
	SetAndroidXiaomiTemplateParams(v string) *PushRequest
	GetAndroidXiaomiTemplateParams() *string
	SetAppKey(v int64) *PushRequest
	GetAppKey() *int64
	SetBody(v string) *PushRequest
	GetBody() *string
	SetDeviceType(v string) *PushRequest
	GetDeviceType() *string
	SetExpireTime(v string) *PushRequest
	GetExpireTime() *string
	SetHarmonyAction(v string) *PushRequest
	GetHarmonyAction() *string
	SetHarmonyActionType(v string) *PushRequest
	GetHarmonyActionType() *string
	SetHarmonyBadgeAddNum(v int32) *PushRequest
	GetHarmonyBadgeAddNum() *int32
	SetHarmonyBadgeSetNum(v int32) *PushRequest
	GetHarmonyBadgeSetNum() *int32
	SetHarmonyCategory(v string) *PushRequest
	GetHarmonyCategory() *string
	SetHarmonyExtParameters(v string) *PushRequest
	GetHarmonyExtParameters() *string
	SetHarmonyExtensionExtraData(v string) *PushRequest
	GetHarmonyExtensionExtraData() *string
	SetHarmonyExtensionPush(v bool) *PushRequest
	GetHarmonyExtensionPush() *bool
	SetHarmonyImageUrl(v string) *PushRequest
	GetHarmonyImageUrl() *string
	SetHarmonyInboxContent(v string) *PushRequest
	GetHarmonyInboxContent() *string
	SetHarmonyLiveViewPayload(v string) *PushRequest
	GetHarmonyLiveViewPayload() *string
	SetHarmonyNotificationSlotType(v string) *PushRequest
	GetHarmonyNotificationSlotType() *string
	SetHarmonyNotifyId(v int32) *PushRequest
	GetHarmonyNotifyId() *int32
	SetHarmonyReceiptId(v string) *PushRequest
	GetHarmonyReceiptId() *string
	SetHarmonyRemind(v bool) *PushRequest
	GetHarmonyRemind() *bool
	SetHarmonyRemindBody(v string) *PushRequest
	GetHarmonyRemindBody() *string
	SetHarmonyRemindTitle(v string) *PushRequest
	GetHarmonyRemindTitle() *string
	SetHarmonyRenderStyle(v string) *PushRequest
	GetHarmonyRenderStyle() *string
	SetHarmonyTestMessage(v bool) *PushRequest
	GetHarmonyTestMessage() *bool
	SetHarmonyUri(v string) *PushRequest
	GetHarmonyUri() *string
	SetIdempotentToken(v string) *PushRequest
	GetIdempotentToken() *string
	SetJobKey(v string) *PushRequest
	GetJobKey() *string
	SetPushTime(v string) *PushRequest
	GetPushTime() *string
	SetPushType(v string) *PushRequest
	GetPushType() *string
	SetSendChannels(v string) *PushRequest
	GetSendChannels() *string
	SetSendSpeed(v int32) *PushRequest
	GetSendSpeed() *int32
	SetSmsDelaySecs(v int32) *PushRequest
	GetSmsDelaySecs() *int32
	SetSmsParams(v string) *PushRequest
	GetSmsParams() *string
	SetSmsSendPolicy(v int32) *PushRequest
	GetSmsSendPolicy() *int32
	SetSmsSignName(v string) *PushRequest
	GetSmsSignName() *string
	SetSmsTemplateName(v string) *PushRequest
	GetSmsTemplateName() *string
	SetStoreOffline(v bool) *PushRequest
	GetStoreOffline() *bool
	SetTarget(v string) *PushRequest
	GetTarget() *string
	SetTargetValue(v string) *PushRequest
	GetTargetValue() *string
	SetTitle(v string) *PushRequest
	GetTitle() *string
	SetTrim(v bool) *PushRequest
	GetTrim() *bool
	SetIOSApnsEnv(v string) *PushRequest
	GetIOSApnsEnv() *string
	SetIOSBadge(v int32) *PushRequest
	GetIOSBadge() *int32
	SetIOSBadgeAutoIncrement(v bool) *PushRequest
	GetIOSBadgeAutoIncrement() *bool
	SetIOSExtParameters(v string) *PushRequest
	GetIOSExtParameters() *string
	SetIOSInterruptionLevel(v string) *PushRequest
	GetIOSInterruptionLevel() *string
	SetIOSLiveActivityAttributes(v string) *PushRequest
	GetIOSLiveActivityAttributes() *string
	SetIOSLiveActivityAttributesType(v string) *PushRequest
	GetIOSLiveActivityAttributesType() *string
	SetIOSLiveActivityContentState(v string) *PushRequest
	GetIOSLiveActivityContentState() *string
	SetIOSLiveActivityDismissalDate(v int64) *PushRequest
	GetIOSLiveActivityDismissalDate() *int64
	SetIOSLiveActivityEvent(v string) *PushRequest
	GetIOSLiveActivityEvent() *string
	SetIOSLiveActivityId(v string) *PushRequest
	GetIOSLiveActivityId() *string
	SetIOSLiveActivityStaleDate(v int64) *PushRequest
	GetIOSLiveActivityStaleDate() *int64
	SetIOSMusic(v string) *PushRequest
	GetIOSMusic() *string
	SetIOSMutableContent(v bool) *PushRequest
	GetIOSMutableContent() *bool
	SetIOSNotificationCategory(v string) *PushRequest
	GetIOSNotificationCategory() *string
	SetIOSNotificationCollapseId(v string) *PushRequest
	GetIOSNotificationCollapseId() *string
	SetIOSNotificationThreadId(v string) *PushRequest
	GetIOSNotificationThreadId() *string
	SetIOSRelevanceScore(v float64) *PushRequest
	GetIOSRelevanceScore() *float64
	SetIOSRemind(v bool) *PushRequest
	GetIOSRemind() *bool
	SetIOSRemindBody(v string) *PushRequest
	GetIOSRemindBody() *string
	SetIOSSilentNotification(v bool) *PushRequest
	GetIOSSilentNotification() *bool
	SetIOSSubtitle(v string) *PushRequest
	GetIOSSubtitle() *string
}

type PushRequest struct {
	// Specifies the activity to open when the notification is tapped.
	//
	// This is required only when \\`AndroidOpenType\\` is \\`Activity\\`. For example: \\`com.alibaba.cloudpushdemo.bizactivity\\`.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidActivity *string `json:"AndroidActivity,omitempty" xml:"AndroidActivity,omitempty"`
	// Sets the value to add to the badge number. This value is added to the original badge number. The value must be between 1 and 99.
	//
	// > This is effective only for pushes through Huawei/Honor vendor channels. If both \\`AndroidBadgeAddNum\\` and \\`AndroidBadgeSetNum\\` are present, \\`AndroidBadgeSetNum\\` takes precedence.
	//
	// example:
	//
	// 1
	AndroidBadgeAddNum *int32 `json:"AndroidBadgeAddNum,omitempty" xml:"AndroidBadgeAddNum,omitempty"`
	// The fully qualified class name of the app\\"s entry Activity for badge setting.
	//
	// > This is effective only for pushes through Huawei/Honor vendor channels.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidBadgeClass *string `json:"AndroidBadgeClass,omitempty" xml:"AndroidBadgeClass,omitempty"`
	// Sets a fixed number for the badge. The value must be between 0 and 99.
	//
	// > For vendor channel pushes, this is effective only for Huawei and Honor channels. For pushes through Alibaba Cloud\\"s proprietary channel, this is effective only on Huawei, Honor, and vivo models.
	//
	// example:
	//
	// 5
	AndroidBadgeSetNum *int32 `json:"AndroidBadgeSetNum,omitempty" xml:"AndroidBadgeSetNum,omitempty"`
	// The body in long text mode. Length limit: 1,000 bytes (1 Chinese character is counted as 3 bytes). The actual limit depends on the specific vendor channel.
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
	// - Proprietary channel: Android SDK 3.6.0 and later
	//
	// > If this parameter is not provided in long text mode, the system uses the first non-empty value from \\`Body\\` or \\`AndroidPopupBody\\`.
	//
	// example:
	//
	// 示例长文本
	AndroidBigBody *string `json:"AndroidBigBody,omitempty" xml:"AndroidBigBody,omitempty"`
	// The image URL for big picture mode. Currently supported by the proprietary channel on Android SDK 3.6.0 and later.
	//
	// example:
	//
	// https://imag.example.com/image.png
	AndroidBigPictureUrl *string `json:"AndroidBigPictureUrl,omitempty" xml:"AndroidBigPictureUrl,omitempty"`
	// The title in long text mode. Length limit: 200 bytes (1 Chinese character is counted as 3 bytes).
	//
	// - Currently, this is only supported by Honor channels and Huawei channels on EMUI 11 and later.
	//
	// - If this parameter is not provided in long text mode, the system uses the first non-empty value from \\`Title\\` or \\`AndroidPopupTitle\\`.
	//
	// example:
	//
	// 示例长标题
	AndroidBigTitle *string `json:"AndroidBigTitle,omitempty" xml:"AndroidBigTitle,omitempty"`
	// Sets the extended properties of the notification. This property is not effective when \\`PushType\\` is \\`MESSAGE\\`.
	//
	// This parameter must be in JSON map format to avoid parsing errors.
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
	// > Each app can send 1,000 test notifications per day. These are not subject to the daily push limit per device.
	//
	// example:
	//
	// 0
	AndroidHonorTargetUserType *int32 `json:"AndroidHonorTargetUserType,omitempty" xml:"AndroidHonorTargetUserType,omitempty"`
	// Sets the Huawei quick notification parameter.
	//
	// - **0**: Send a standard Huawei notification (default).
	//
	// - **1**: Send a Huawei quick notification.
	//
	// example:
	//
	// 1
	AndroidHuaweiBusinessType *int32 `json:"AndroidHuaweiBusinessType,omitempty" xml:"AndroidHuaweiBusinessType,omitempty"`
	// A JSON string of the Huawei Android Live Notification data structure [LiveNotificationPayload](https://developer.huawei.com/consumer/cn/doc/HMSCore-References/rest-live-0000001562939968#ZH-CN_TOPIC_0000001700850537__p195121620102511). For development and integration, see [Huawei Live Notification Push Guide](https://help.aliyun.com/document_detail/2983768.html).
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
	// The receipt ID for the Huawei channel. You can find this ID in the receipt parameter configuration on the Huawei Push service platform.
	//
	// > If the default receipt configuration on the Huawei Push service platform is the Alibaba Cloud receipt, do not provide this. If not, first configure the default Huawei channel receipt ID in the Alibaba Cloud EMAS Mobile Push console.
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
	// > Each app can send 500 test notifications per day. These are not subject to the daily push limit per device.
	//
	// example:
	//
	// 0
	AndroidHuaweiTargetUserType *int32 `json:"AndroidHuaweiTargetUserType,omitempty" xml:"AndroidHuaweiTargetUserType,omitempty"`
	// The URL for the right-side icon.
	//
	// Currently supported on:
	//
	// - Huawei EMUI (only in long text and inbox modes).
	//
	// - Honor Magic UI (only in long text mode).
	//
	// - Proprietary channel: Android SDK 3.5.0 and later.
	//
	// example:
	//
	// https://imag.example.com/image.png
	AndroidImageUrl *string `json:"AndroidImageUrl,omitempty" xml:"AndroidImageUrl,omitempty"`
	// The body content for inbox mode. The content must be a valid JSON array with no more than 5 elements. Currently supported on:
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
	// - 0 Public message (default)
	//
	// - 1 Private message
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 0
	AndroidMeizuNoticeMsgType *int32 `json:"AndroidMeizuNoticeMsgType,omitempty" xml:"AndroidMeizuNoticeMsgType,omitempty"`
	// Function 1: After applying for [self-classification rights](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835?#section3410731125514), this is used to identify the message type and determine the [message alert method](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#ZH-CN_TOPIC_0000001149358835__p3850133955718). It accelerates the sending of specific message types. For valid values, refer to the [message classification standards](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section1076611477914) in the official Huawei Push documentation. Fill in the \\"Cloud notification category value\\" or \\"Local notification category value\\" from the document\\"s table.
	//
	// Function 2: After applying for [special permissions](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/faq-0000001050042183#section037425218509), this is used to identify high-priority pass-through scenarios. Valid values:
	//
	// - VOIP: Voice and video calls
	//
	// - PLAY_VOICE: Voice playback
	//
	// > If the \\"Cloud notification category value\\" is \\"Not applicable\\", the push is sent through Alibaba Cloud\\"s proprietary channel. If the \\"Local notification category value\\" is \\"Not applicable\\", the push is sent through the Huawei channel.
	//
	// example:
	//
	// VOIP
	AndroidMessageHuaweiCategory *string `json:"AndroidMessageHuaweiCategory,omitempty" xml:"AndroidMessageHuaweiCategory,omitempty"`
	// The delivery priority for notifications on the Huawei channel. Valid values:
	//
	// - **HIGH**
	//
	// - **NORMAL**
	//
	// Apply for permission. For more information, see [Application link](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/faq-0000001050042183#section037425218509).
	//
	// example:
	//
	// HIGH
	AndroidMessageHuaweiUrgency *string `json:"AndroidMessageHuaweiUrgency,omitempty" xml:"AndroidMessageHuaweiUrgency,omitempty"`
	// OPPO classifies and manages messages in two categories: Communication & Service, and Content & Marketing.
	//
	// Communication & Service (requires permission):
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
	// - MARKETING: Platform activities
	//
	// - SOCIAL: Social updates
	//
	// For more information, see [OPUSH Message Classification Rules](https://open.oppomobile.com/new/developmentDoc/info?id=13189).
	//
	// example:
	//
	// MARKETING
	AndroidMessageOppoCategory *string `json:"AndroidMessageOppoCategory,omitempty" xml:"AndroidMessageOppoCategory,omitempty"`
	// The alert level for notification bar messages on the OPPO channel. Valid values:
	//
	// - 1: Notification bar
	//
	// - 2: Notification bar, lock screen, ringtone, vibration (default level for Communication & Service messages)
	//
	// - 16: Notification bar, lock screen, ringtone, vibration, banner (requires permission)
	//
	// > When using the \\`AndroidMessageOppoNotifyLevel\\` parameter, you must also pass the \\`AndroidMessageOppoCategory\\` parameter.
	//
	// example:
	//
	// 1
	AndroidMessageOppoNotifyLevel *int32 `json:"AndroidMessageOppoNotifyLevel,omitempty" xml:"AndroidMessageOppoNotifyLevel,omitempty"`
	// vivo classifies and manages messages in two categories: System messages and Operational messages.
	//
	// System messages:
	//
	// - IM: Instant messages
	//
	// - ACCOUNT: Account and assets
	//
	// - TODO: Schedule and to-do
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
	// - MARKETING: Operational activities
	//
	// - SOCIAL: Social updates
	//
	// For more information, see [Classification description](https://dev.vivo.com.cn/documentCenter/doc/359#s-ef3qugc3).
	//
	// example:
	//
	// TODO
	AndroidMessageVivoCategory *string `json:"AndroidMessageVivoCategory,omitempty" xml:"AndroidMessageVivoCategory,omitempty"`
	// The notification sound for the Huawei vendor channel. Specify the name of the audio file located in the \\`app/src/main/res/raw/\\` directory of the client project. Do not include the file format suffix.
	//
	// If this is not set, the default ringtone is used.
	//
	// example:
	//
	// alicloud_notification_sound
	AndroidMusic *string `json:"AndroidMusic,omitempty" xml:"AndroidMusic,omitempty"`
	// The priority for arranging the Android notification in the notification bar. Valid values: -2, -1, 0, 1, 2.
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
	// The \\`channelId\\` for the Android app. This must correspond to a \\`channelId\\` in the app.
	//
	// - Set the \\`NotificationChannel\\` parameter. For more information about its usage, see [FAQ: Why are notifications not received on devices running Android 8.0 or later?](https://help.aliyun.com/document_detail/67398.html).
	//
	// - Because the \\`channel_id\\` for the OPPO private message channel is the same as the app\\"s \\`channelId\\`, this value is used for pushes through the OPPO channel.
	//
	// - This value is used for pushes through Huawei, FCM, and Alibaba Cloud\\"s proprietary channels.
	//
	// example:
	//
	// 1
	AndroidNotificationChannel *string `json:"AndroidNotificationChannel,omitempty" xml:"AndroidNotificationChannel,omitempty"`
	// Message grouping. For messages in the same group, the notification bar shows only the latest message and the total number of messages received for that group. It does not display all messages and cannot be expanded. Currently supported on:
	//
	// - Huawei vendor channel
	//
	// - Honor vendor channel
	//
	// - Proprietary channel for Android SDK 3.9.1 and earlier
	//
	// > This parameter is no longer supported by the proprietary channel for Android SDK 3.9.2 and later.
	//
	// example:
	//
	// group-1
	AndroidNotificationGroup *string `json:"AndroidNotificationGroup,omitempty" xml:"AndroidNotificationGroup,omitempty"`
	// Sets the \\`importance\\` parameter for Honor notification message classification. This determines the notification behavior on the user\\"s device. Valid values:
	//
	// - **LOW**: For informational and marketing messages.
	//
	// - **NORMAL**: For service and communication messages.
	//
	// Apply for this on the Honor platform. [Application link](https://developer.honor.com/cn/docs/11002/guides/notification-class#%E8%87%AA%E5%88%86%E7%B1%BB%E6%9D%83%E7%9B%8A%E7%94%B3%E8%AF%B7).
	//
	// example:
	//
	// LOW
	AndroidNotificationHonorChannel *string `json:"AndroidNotificationHonorChannel,omitempty" xml:"AndroidNotificationHonorChannel,omitempty"`
	// Sets the \\`importance\\` parameter for Huawei notification message classification. This determines the notification behavior on the user\\"s device. Valid values:
	//
	// - LOW: For informational and marketing messages.
	//
	// - NORMAL: For service and communication messages.
	//
	// > 	- For the Huawei channel, use \\`AndroidMessageHuaweiCategory\\` for notification classification. You may no longer need to use \\`AndroidNotificationHuaweiChannel\\`.
	//
	// >
	//
	// > 	- Apply for this on the Huawei platform. [Application link](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section893184112272).
	//
	// example:
	//
	// LOW
	AndroidNotificationHuaweiChannel *string `json:"AndroidNotificationHuaweiChannel,omitempty" xml:"AndroidNotificationHuaweiChannel,omitempty"`
	// A unique identifier for each message when it is displayed as a notification. Different notifications can have the same \\`NotifyId\\`, which allows a new notification to overwrite an old one.
	//
	// example:
	//
	// 100001
	AndroidNotificationNotifyId *int32 `json:"AndroidNotificationNotifyId,omitempty" xml:"AndroidNotificationNotifyId,omitempty"`
	// Message grouping. Messages in the same group are displayed in a collapsed state in the notification bar and can be expanded. Notifications from different groups are displayed separately. Currently supported on:
	//
	// - Proprietary channel for Android SDK 3.9.2 and later
	//
	// example:
	//
	// thread-1
	AndroidNotificationThreadId *string `json:"AndroidNotificationThreadId,omitempty" xml:"AndroidNotificationThreadId,omitempty"`
	// Sets the classification for vivo notification messages. Valid values:
	//
	// - 0: Operational messages (default)
	//
	// - 1: System messages
	//
	// > 	- For the vivo channel, use \\`AndroidMessageVivoCategory\\` for notification classification. You may no longer need to use \\`AndroidNotificationVivoChannel\\`.
	//
	// >
	//
	// > 	- Apply for this on the vivo platform. For more information, see [Application link](https://dev.vivo.com.cn/documentCenter/doc/359).
	//
	// example:
	//
	// classification
	AndroidNotificationVivoChannel *string `json:"AndroidNotificationVivoChannel,omitempty" xml:"AndroidNotificationVivoChannel,omitempty"`
	// Sets the \\`channelId\\` for the Xiaomi notification type. Apply for this on the Xiaomi platform. For more information, see [Application link](https://dev.mi.com/console/doc/detail?pId=2422#_4).
	//
	// > - A single app can apply for a maximum of 8 channels through the Xiaomi channel. Plan accordingly.
	//
	// example:
	//
	// michannel
	AndroidNotificationXiaomiChannel *string `json:"AndroidNotificationXiaomiChannel,omitempty" xml:"AndroidNotificationXiaomiChannel,omitempty"`
	// The notification alert type. Valid values:
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
	// BOTH
	AndroidNotifyType *string `json:"AndroidNotifyType,omitempty" xml:"AndroidNotifyType,omitempty"`
	// The action to take after a notification is tapped. Valid values:
	//
	// - **APPLICATION**: Open the application (default).
	//
	// - **ACTIVITY**: Open a specific Android Activity.
	//
	// - **URL**: Open a URL.
	//
	// - **NONE**: No action.
	//
	// example:
	//
	// APPLICATION
	AndroidOpenType *string `json:"AndroidOpenType,omitempty" xml:"AndroidOpenType,omitempty"`
	// The URL to open after the Android device receives the push.
	//
	// This is required only when \\`AndroidOpenType\\` is \\`URL\\`.
	//
	// example:
	//
	// https://xxxx.xxx
	AndroidOpenUrl *string `json:"AndroidOpenUrl,omitempty" xml:"AndroidOpenUrl,omitempty"`
	// A JSON string of the OPPO Fluid Cloud intent deletion data structure [data](https://open.oppomobile.com/documentation/page/info?id=13578). This parameter is invalid if the \\`AndroidOppoIntelligentIntent\\` parameter is filled. For development and integration, see [OPPO Fluid Cloud Push Guide](https://help.aliyun.com/document_detail/2997310.html).
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
	// A JSON string of the OPPO Fluid Cloud intent sharing data structure [IntelligentIntent](https://open.oppomobile.com/documentation/page/info?id=13565). For development and integration, see [OPPO Fluid Cloud Push Guide](https://help.aliyun.com/document_detail/2997310.html).
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
	// Sets the OPPO Fluid Cloud push environment.
	//
	// - **0**: Production environment (default).
	//
	// - **1**: Staging environment.
	//
	// > The OPPO Fluid Cloud staging environment must be set up on the client side. For more information, see [Environment setup](https://open.oppomobile.com/documentation/page/info?id=13590).
	//
	// example:
	//
	// 1
	AndroidOppoIntentEnv *int32 `json:"AndroidOppoIntentEnv,omitempty" xml:"AndroidOppoIntentEnv,omitempty"`
	// OPPO private message template content parameters
	AndroidOppoPrivateContentParameters map[string]*string `json:"AndroidOppoPrivateContentParameters,omitempty" xml:"AndroidOppoPrivateContentParameters,omitempty"`
	// OPPO private message template ID
	//
	// example:
	//
	// 687557242b1634hzefs3d5013
	AndroidOppoPrivateMsgTemplateId *string `json:"AndroidOppoPrivateMsgTemplateId,omitempty" xml:"AndroidOppoPrivateMsgTemplateId,omitempty"`
	// OPPO private message template title parameters
	AndroidOppoPrivateTitleParameters map[string]*string `json:"AndroidOppoPrivateTitleParameters,omitempty" xml:"AndroidOppoPrivateTitleParameters,omitempty"`
	// Specifies the Activity to launch after the notification is tapped.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidPopupActivity *string `json:"AndroidPopupActivity,omitempty" xml:"AndroidPopupActivity,omitempty"`
	// The body content in auxiliary pop-up mode. This parameter is required if \\`AndroidPopupActivity\\` is not empty.
	//
	// Length limit: 200 characters. Both Chinese and English characters count as one.
	//
	// If you use a vendor channel, comply with its restrictions. For more information, see [Limits on pushes through auxiliary channels on Android](https://help.aliyun.com/document_detail/165253.html).
	//
	// example:
	//
	// hello
	AndroidPopupBody *string `json:"AndroidPopupBody,omitempty" xml:"AndroidPopupBody,omitempty"`
	// The title content in auxiliary pop-up mode. This parameter is required if \\`AndroidPopupActivity\\` is not empty.
	//
	// Length limit: 50 characters. Both Chinese and English characters count as one.
	//
	// If you use a vendor channel, comply with its restrictions. For more information, see [Limits on pushes through auxiliary channels on Android](https://help.aliyun.com/document_detail/165253.html).
	//
	// example:
	//
	// hello
	AndroidPopupTitle *string `json:"AndroidPopupTitle,omitempty" xml:"AndroidPopupTitle,omitempty"`
	// If the device is offline when a message is pushed, this push uses the auxiliary pop-up feature. The default value is \\`false\\`. This is effective only when \\`PushType\\` is \\`MESSAGE\\`.
	//
	// If the message is successfully converted to a notification, the data displayed in the notification is the value of the \\`AndroidPopupTitle\\` and \\`AndroidPopupBody\\` parameters set on the server. When the notification is tapped, the data obtained in the \\`onSysNoticeOpened\\` method of the auxiliary pop-up is the value of the \\`Title\\` and \\`Body\\` parameters set on the server.
	//
	// example:
	//
	// true
	AndroidRemind *bool `json:"AndroidRemind,omitempty" xml:"AndroidRemind,omitempty"`
	// The notification style. Valid values:
	//
	// - **0**: Standard mode (default)
	//
	// - **1**: Long text mode (supported by Huawei, Honor, Xiaomi, OPPO, Meizu, and proprietary channels)
	//
	// - **2**: Big picture mode (supported by proprietary channels, but not by Xiaomi models)
	//
	// - **3**: List mode (supported by Huawei, Honor, Xiaomi, OPPO, and proprietary channels)
	//
	// > This parameter is required if you use a non-standard mode.
	//
	// example:
	//
	// 1
	AndroidRenderStyle *int32 `json:"AndroidRenderStyle,omitempty" xml:"AndroidRenderStyle,omitempty"`
	// Sets the vendor channel notification type:
	//
	// - **0**: Formal notification (default).
	//
	// - **1**: Test notification.
	//
	// > 	- Configuring this parameter is equivalent to configuring \\`AndroidHuaweiTargetUserType\\`, \\`AndroidHonorTargetUserType\\`, \\`AndroidVivoPushMode\\`, and \\`AndroidOppoIntentEnv\\` simultaneously. Specific vendor channel parameters can override this setting.
	//
	// >
	//
	// > 	- Currently supported by: Huawei channel, Honor channel, vivo channel, and OPPO Fluid Cloud.
	//
	// example:
	//
	// 0
	AndroidTargetUserType *int32 `json:"AndroidTargetUserType,omitempty" xml:"AndroidTargetUserType,omitempty"`
	// A JSON string of the vivo Atomic Island data structure [liveMessage](https://dev.vivo.com.cn/documentCenter/doc/896#s-fdagzbd4). For development and integration, see [vivo Atomic Island Push Guide](https://help.aliyun.com/zh/document_detail/3030718.html).
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
	// > For test pushes, configure test devices in the vivo console beforehand. Find the test device\\"s \\`RegId\\` by searching for "onReceiveRegId regId" in the device startup logs.
	//
	// example:
	//
	// 0
	AndroidVivoPushMode *int32 `json:"AndroidVivoPushMode,omitempty" xml:"AndroidVivoPushMode,omitempty"`
	// The receipt ID for the vivo channel. You can find this ID in the application information section of the vivo open platform\\"s push service.
	//
	// > If the default receipt configuration on the vivo open platform is the Alibaba Cloud receipt, do not provide this. If not, first configure the default vivo channel receipt ID in the Alibaba Cloud EMAS Mobile Push console.
	//
	// example:
	//
	// 123
	AndroidVivoReceiptId *string `json:"AndroidVivoReceiptId,omitempty" xml:"AndroidVivoReceiptId,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. All third-party auxiliary pop-ups are now supported by the new parameter **AndroidPopupActivity**.
	//
	// example:
	//
	// 无
	AndroidXiaoMiActivity *string `json:"AndroidXiaoMiActivity,omitempty" xml:"AndroidXiaoMiActivity,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. All third-party auxiliary pop-ups are now supported by the new parameter **AndroidPopupBody**.
	//
	// example:
	//
	// 无
	AndroidXiaoMiNotifyBody *string `json:"AndroidXiaoMiNotifyBody,omitempty" xml:"AndroidXiaoMiNotifyBody,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. All third-party auxiliary pop-ups are now supported by the new parameter **AndroidPopupTitle**.
	//
	// example:
	//
	// 无
	AndroidXiaoMiNotifyTitle *string `json:"AndroidXiaoMiNotifyTitle,omitempty" xml:"AndroidXiaoMiNotifyTitle,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Starting from August 2023, Xiaomi no longer supports dynamically setting small icons, right-side icons, or large pictures during pushes on new devices/systems.
	//
	// example:
	//
	// https://f6.market.xiaomi.com/download/MiPass/aaa/bbb.png
	AndroidXiaomiBigPictureUrl *string `json:"AndroidXiaomiBigPictureUrl,omitempty" xml:"AndroidXiaomiBigPictureUrl,omitempty"`
	// A JSON string of the Xiaomi Super Island data structure [miui.focus.param](https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131). For development and integration, see [Xiaomi Super Island Push Guide](https://help.aliyun.com/zh/document_detail/3037956.html).
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
	// A JSON string of the Xiaomi Super Island data images [miui.focus.pic_xxx](https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131). For development and integration, see [Xiaomi Super Island Push Guide](https://help.aliyun.com/zh/document_detail/3037956.html).
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
	// This parameter is deprecated. Starting from August 2023, Xiaomi no longer supports dynamically setting small icons, right-side icons, or large pictures during pushes on new devices/systems.
	//
	// example:
	//
	// https://imag.example.com/image.png
	AndroidXiaomiImageUrl       *string `json:"AndroidXiaomiImageUrl,omitempty" xml:"AndroidXiaomiImageUrl,omitempty"`
	AndroidXiaomiTemplateId     *string `json:"AndroidXiaomiTemplateId,omitempty" xml:"AndroidXiaomiTemplateId,omitempty"`
	AndroidXiaomiTemplateParams *string `json:"AndroidXiaomiTemplateParams,omitempty" xml:"AndroidXiaomiTemplateParams,omitempty"`
	// The AppKey.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The content of the notification or message for Android and HarmonyOS pushes. The content of the message or notification for iOS. The size of the push content is limited. For more information, see [Product limits](https://help.aliyun.com/document_detail/434629.html).
	//
	// example:
	//
	// hello
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The device type. Valid values:
	//
	// - **HARMONY**: A HarmonyOS device.
	//
	// - **iOS**: An iOS device.
	//
	// - **ANDROID**: An Android device.
	//
	// - **ALL**: For older dual-platform apps, this sends pushes to both Android and iOS devices. For newer single-platform apps, this has the same effect as specifying the device type for that app.
	//
	// This parameter is required.
	//
	// example:
	//
	// HARMONY
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The expiration time for offline messages or notifications. Use this with \\`StoreOffline\\`. The message is not sent after this time. The maximum retention period is 72 hours, which is also the default.
	//
	// The time must be in ISO 8601 format and in UTC: \\`YYYY-MM-DDThh:mm:ssZ\\`. The expiration time must be at least 3 seconds after the current time or the scheduled push time (\\`ExpireTime\\` > \\`PushTime\\` + 3 seconds). The 3-second buffer accounts for network and system delays. For single pushes, use a value of at least 1 minute. For batch pushes or pushes to all devices, use a value of at least 10 minutes.
	//
	// example:
	//
	// 2019-02-20T00:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The action corresponding to the in-app page ability.
	//
	// 	Notice:
	//
	// When \\`HarmonyActionType\\` is \\`APP_CUSTOM_PAGE\\`, fill in at least one of \\`HarmonyUri\\` or \\`HarmonyAction\\`.
	//
	//
	//
	// For more information, see [ClickAction.action](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section152462191216) in the HarmonyOS documentation.
	//
	// example:
	//
	// com.example.action
	HarmonyAction *string `json:"HarmonyAction,omitempty" xml:"HarmonyAction,omitempty"`
	// The action to take after a notification is tapped. Valid values:
	//
	// - APP_HOME_PAGE: Open the app\\"s home page.
	//
	// - APP_CUSTOM_PAGE: Open a custom page in the app.
	//
	// example:
	//
	// APP_HOME_PAGE
	HarmonyActionType *string `json:"HarmonyActionType,omitempty" xml:"HarmonyActionType,omitempty"`
	// The number to add to the HarmonyOS app badge. See the description of the [HarmonyOS badge addNum field](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section266310382145).<br>
	//
	// Supported starting from HarmonyOS SDK version 1.2.0.<br>
	//
	// example:
	//
	// 1
	HarmonyBadgeAddNum *int32 `json:"HarmonyBadgeAddNum,omitempty" xml:"HarmonyBadgeAddNum,omitempty"`
	// The number to set for the HarmonyOS app badge. See the description of the [HarmonyOS badge setNum field](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section266310382145).
	//
	// Supported starting from HarmonyOS SDK version 1.2.0.
	//
	// example:
	//
	// 1
	HarmonyBadgeSetNum *int32 `json:"HarmonyBadgeSetNum,omitempty" xml:"HarmonyBadgeSetNum,omitempty"`
	// The notification message category. After you apply for notification message self-classification rights, this is used to identify the message type. Different notification message types affect how messages are displayed and alerted. Valid values:
	//
	// - IM: Instant messaging
	//
	// - VOIP: Voice and video calls
	//
	// - SUBSCRIPTION: Subscriptions
	//
	// - TRAVEL: Travel
	//
	// - HEALTH: Health
	//
	// - WORK: Work reminders
	//
	// - ACCOUNT: Account updates
	//
	// - EXPRESS: Orders & logistics
	//
	// - FINANCE: Finance
	//
	// - DEVICE_REMINDER: Device reminders
	//
	// - MAIL: Mail
	//
	// - CUSTOMER_SERVICE: Customer service messages
	//
	// - MARKETING: News, content recommendations, social updates, product promotions, financial updates, lifestyle information, surveys, feature recommendations, and operational activities. This only identifies the content and does not speed up message delivery. These are collectively known as informational and marketing messages.
	//
	// For more information, see [Notification.category](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117) in the HarmonyOS documentation.
	//
	// example:
	//
	// IM
	HarmonyCategory *string `json:"HarmonyCategory,omitempty" xml:"HarmonyCategory,omitempty"`
	// Sets the extended properties of the notification. This property is not effective when \\`PushType\\` is \\`MESSAGE\\`.
	//
	// This parameter must be in JSON map format to avoid parsing errors.
	//
	// example:
	//
	// {"key1":"value1","api_name":"PushNoticeToAndroidRequest"}
	HarmonyExtParameters *string `json:"HarmonyExtParameters,omitempty" xml:"HarmonyExtParameters,omitempty"`
	// The extra data for the extended notification message.<br>
	//
	// This is effective when sending a HarmonyOS extended notification message.<br>
	//
	// Conceptually, this is equivalent to the \\`extraData\\` field of a HarmonyOS extended notification message. For the specific definition, see [HarmonyOS ExtensionPayload Description](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section161192514234).<br>
	//
	// Supported starting from HarmonyOS SDK version 1.2.0.<br><br><br>
	//
	// example:
	//
	// 示例额外数据
	HarmonyExtensionExtraData *string `json:"HarmonyExtensionExtraData,omitempty" xml:"HarmonyExtensionExtraData,omitempty"`
	// When \\`PushType\\` is \\`NOTICE\\`, specifies whether this is a HarmonyOS extended notification message.
	//
	// - true: Send an extended notification message.
	//
	// - false: Send a normal notification (default).
	//
	// Apply for permission on the HarmonyOS side before you can send extended notification messages. For more information, see [Send Extended Notification Messages](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V5/push-send-extend-noti-V5) in the HarmonyOS documentation.<br>
	//
	// Supported starting from HarmonyOS SDK version 1.2.0.<br>
	//
	// example:
	//
	// true
	HarmonyExtensionPush *bool `json:"HarmonyExtensionPush,omitempty" xml:"HarmonyExtensionPush,omitempty"`
	// The URL for the large icon on the right of the notification. The URL must use the HTTPS protocol.
	//
	// > Supported image formats are PNG, JPG, JPEG, HEIF, GIF, and BMP. The image dimensions (height × width) must be less than 25,000 pixels.
	//
	// For more information, see [Notification.image](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117) in the HarmonyOS documentation.
	//
	// example:
	//
	// https://example.com/xxx.png
	HarmonyImageUrl *string `json:"HarmonyImageUrl,omitempty" xml:"HarmonyImageUrl,omitempty"`
	// The content for the multi-line text style. This field is required when \\`HarmonyRenderStyle\\` is \\`MULTI_LINE\\`. It supports up to 3 lines of content.
	//
	// example:
	//
	// ["1.content1","2.content2","3.content3"]
	HarmonyInboxContent *string `json:"HarmonyInboxContent,omitempty" xml:"HarmonyInboxContent,omitempty"`
	// A JSON string of the HarmonyOS Live Window data structure [LiveViewPayload](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V13/push-scenariozed-api-request-param-V13#section66881469306). For development and integration, see [HarmonyOS Live Window Push Guide](https://help.aliyun.com/document_detail/2982112.html).
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
	// - SERVICE_INFORMATION: Service reminders.
	//
	// - CONTENT_INFORMATION: Content information.
	//
	// - CUSTOMER_SERVICE: Customer service messages. This type is for messages between users and businesses and must be initiated by the user.
	//
	// - OTHER_TYPES: Others.
	//
	// For more information, see [SlotType](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/js-apis-notificationmanager-V5#slottype) in the HarmonyOS documentation.
	//
	// example:
	//
	// SOCIAL_COMMUNICATION
	HarmonyNotificationSlotType *string `json:"HarmonyNotificationSlotType,omitempty" xml:"HarmonyNotificationSlotType,omitempty"`
	// A unique identifier for each message when it is displayed as a notification. If not provided, the push service automatically generates a unique ID for each message. Different notifications can have the same \\`notifyId\\`, which allows a new message to overwrite an old one.
	//
	// For more information, see [Notification.notifyId](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117) in the HarmonyOS documentation.
	//
	// example:
	//
	// 0
	HarmonyNotifyId *int32 `json:"HarmonyNotifyId,omitempty" xml:"HarmonyNotifyId,omitempty"`
	// The receipt ID for the HarmonyOS channel. You can find this ID in the receipt parameter configuration on the HarmonyOS Push service platform.
	//
	// > If the default receipt configuration on the HarmonyOS Push service platform is the Alibaba Cloud receipt, do not provide this. If not, first configure the default HarmonyOS channel receipt ID in the Alibaba Cloud EMAS Mobile Push console.
	//
	// For more information, see [pushOptions.receiptId](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section418321011212) in the HarmonyOS documentation.
	//
	// example:
	//
	// RCPB***DFD5
	HarmonyReceiptId *string `json:"HarmonyReceiptId,omitempty" xml:"HarmonyReceiptId,omitempty"`
	// If the device is offline when a message is pushed, this push uses the auxiliary pop-up feature. The default value is \\`false\\`. This is effective only when \\`PushType\\` is \\`MESSAGE\\`.
	//
	// If the message is successfully converted to a notification, the data displayed in the notification is the value of the \\`HarmonyRemindTitle\\` and \\`HarmonyRemindBody\\` parameters set on the server.
	//
	// example:
	//
	// false
	HarmonyRemind *bool `json:"HarmonyRemind,omitempty" xml:"HarmonyRemind,omitempty"`
	// The HarmonyOS notification content used when a message is converted to a notification. This is effective only when \\`HarmonyRemind\\` is \\`true\\`.
	//
	// example:
	//
	// 您有一条新消息，请查收
	HarmonyRemindBody *string `json:"HarmonyRemindBody,omitempty" xml:"HarmonyRemindBody,omitempty"`
	// The HarmonyOS notification title used when a message is converted to a notification. This is effective only when \\`HarmonyRemind\\` is \\`true\\`.
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
	// Test message flag:
	//
	// - false: Normal message (default)
	//
	// - true: Test message
	//
	// For more information, see [pushOptions.testMessage](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section418321011212) in the HarmonyOS documentation.
	//
	// example:
	//
	// true
	HarmonyTestMessage *bool `json:"HarmonyTestMessage,omitempty" xml:"HarmonyTestMessage,omitempty"`
	// The URI corresponding to the in-app page ability.
	//
	// 	Notice: When \\`HarmonyActionType\\` is \\`APP_CUSTOM_PAGE\\`, fill in at least one of \\`HarmonyUri\\` or \\`HarmonyAction\\`. If there are multiple abilities, fill in the action and URI for each. The action is used with priority to find the corresponding in-app page.
	//
	// For more information, see [ClickAction.uri](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section152462191216) in the HarmonyOS documentation.
	//
	// example:
	//
	// https://www.example.com:8080/push/example
	HarmonyUri *string `json:"HarmonyUri,omitempty" xml:"HarmonyUri,omitempty"`
	// An idempotent parameter to prevent duplicate pushes caused by API call retries. If you make a call with the same \\`IdempotentToken\\` within 15 minutes, only one push is sent. Subsequent calls return the result of the first successful push.
	//
	// > - The parameter must be a standard 36-character UUID (8-4-4-4-12). Each valid character must be a hexadecimal digit from 0-9 or a-f, case-insensitive.
	//
	// >
	//
	// > - This parameter only prevents duplicate pushes from retries. It cannot prevent duplicate pushes from concurrent calls.
	//
	// example:
	//
	// c8016d13-6e76-410c-9bda-769383d11787
	IdempotentToken *string `json:"IdempotentToken,omitempty" xml:"IdempotentToken,omitempty"`
	// A custom ID for the push task. If \\`JobKey\\` is not empty, this field is included in the receipt logs. For more information about receipt logs, see [Receipt logs](https://help.aliyun.com/document_detail/434651.html).
	//
	// > The format must consist of letters, numbers, underscores (_), or hyphens (-). The length cannot exceed 32 characters.
	//
	// example:
	//
	// 123
	JobKey *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	// Used for scheduled sending. If you do not set this parameter, the push is sent immediately.
	//
	// The scheduled time can be no more than 7 days in the future.
	//
	// The time must be in ISO 8601 format and in UTC: \\`YYYY-MM-DDThh:mm:ssZ\\`.
	//
	// > Scheduled sending is not supported when \\`Target\\` is \\`TBD\\` (continuous push).
	//
	// example:
	//
	// 2019-02-20T00:00:00Z
	PushTime *string `json:"PushTime,omitempty" xml:"PushTime,omitempty"`
	// The push type. Valid values:
	//
	// - **NOTICE**: A notification. Notifications are sent to devices through vendor channels, such as APNs, Huawei, Xiaomi, and HarmonyOS, and appear directly in the device\\"s notification bar. When an Android device is online (the app process is active), the notification is preferentially sent through Alibaba Cloud\\"s proprietary channel. The Push software development kit (SDK) then constructs and displays the notification. This improves push performance and can save on vendor channel message quotas in some scenarios.
	//
	// - **MESSAGE**: A message. Messages are sent through Alibaba Cloud\\"s proprietary online channel. They do not appear in the notification bar by default. Instead, the app must be active to receive and process them. Your business logic determines whether to trigger any actions. If a device is offline (the app process is inactive), it cannot receive messages immediately. In this case, use the \\`iOSRemind\\` or \\`AndroidRemind\\` parameter to convert the message into a notification. Alternatively, set the \\`StoreOffline\\` parameter to have the push system save the message. The system then delivers the message automatically when the device comes back online.
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
	// > 	- If you do not set this parameter, all channels can be used.
	//
	// >
	//
	// > 	- If you set this parameter, only the specified channels are used.
	//
	// >
	//
	// > 	- If the specified channels conflict with the sending policy, the push is not sent. For example, if an iOS notification can only be sent through the APNs channel, but \\`apns\\` is not included in this parameter, the push will fail.
	//
	// >
	//
	// > 	- If you specify \\`gcm\\`, pushes can be sent through both Google GCM and FCM channels. If you specify \\`fcm\\`, pushes can only be sent through the Google FCM channel.
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
	// The delay time in seconds before triggering the text message.
	//
	// This must be set if using SMS filter interaction. Set it to 15 seconds or more, with a maximum of 3 days, to avoid duplicate pushes and text messages.
	//
	// > When using SMS filter interaction, the \\`ExpireTime\\` parameter is invalid. The notification expiration time is calculated based on the \\`SmsDelaySecs\\` parameter. The expiration time is the current time plus the \\`SmsDelaySecs\\` time.
	//
	// example:
	//
	// 15
	SmsDelaySecs *int32 `json:"SmsDelaySecs,omitempty" xml:"SmsDelaySecs,omitempty"`
	// The key-value pairs for the variables in the SMS template. Format: `key1=value1&key2=value2`.
	//
	// example:
	//
	// key1=value1
	SmsParams *string `json:"SmsParams,omitempty" xml:"SmsParams,omitempty"`
	// The condition for triggering the text message. Valid values:
	//
	// - **0**: Triggered when the push is not received.
	//
	// - **1**: Triggered when the user does not open the push.
	//
	// example:
	//
	// 0
	SmsSendPolicy *int32 `json:"SmsSendPolicy,omitempty" xml:"SmsSendPolicy,omitempty"`
	// The signature for the supplementary text message.
	//
	// example:
	//
	// 短信签名
	SmsSignName *string `json:"SmsSignName,omitempty" xml:"SmsSignName,omitempty"`
	// The name of the SMS template for supplementary sending. Get this from the SMS template management interface. This is the system-assigned name, not the name set by the developer.
	//
	// example:
	//
	// 短信模板名称
	SmsTemplateName *string `json:"SmsTemplateName,omitempty" xml:"SmsTemplateName,omitempty"`
	// Specifies whether to save offline messages and notifications. The default value is **false**.
	//
	// If set to true, and a user is offline, the message is sent again when the user comes online before the \\`ExpireTime\\`. The default \\`ExpireTime\\` is 72 hours. iOS notifications are sent through APNs and are not affected by this parameter.
	//
	// example:
	//
	// false
	StoreOffline *bool `json:"StoreOffline,omitempty" xml:"StoreOffline,omitempty"`
	// The push target. Valid values:
	//
	// - **DEVICE**: Push to devices.
	//
	// - **ACCOUNT**: Push to accounts.
	//
	// - **ALIAS**: Push to aliases.
	//
	// - **TAG**: Push to tags.
	//
	// - **ALL**: Push to all devices. The interval between two consecutive pushes to all devices of the same \\`DeviceType\\` must be at least 1 second.
	//
	// > When pushing to all iOS devices, the push is sent to devices that have been active in the last 24 months and have not uninstalled the app. A push is considered delivered once the Apple Push Notification service (APNs) receives the request and does not return an error. This can cause a sharp increase in the number of active devices and lead to significant costs. Use this feature with caution.
	//
	// - **TBD**: Initializes a continuous push. The target is specified by a subsequent call to the [ContinuouslyPush](https://help.aliyun.com/document_detail/2249917.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// Set this based on the \\`Target\\` type. Use commas to separate multiple values. If you exceed the limit, send multiple pushes.
	//
	// - If \\`Target\\` is \\`DEVICE\\`, provide device IDs, such as \\`deviceid1,deviceid2\\`. You can specify up to 1,000 device IDs.
	//
	// - If \\`Target\\` is \\`ACCOUNT\\`, provide account IDs, such as \\`account1,account2\\`. You can specify up to 1,000 account IDs.
	//
	// - If \\`Target\\` is \\`ALIAS\\`, provide aliases, such as \\`alias1,alias2\\`. You can specify up to 1,000 aliases.
	//
	// - If \\`Target\\` is \\`TAG\\`, you can use single or multiple tags. For more information about the format, see [Tag format](https://help.aliyun.com/document_detail/434847.html).
	//
	// - If \\`Target\\` is \\`ALL\\`, set the value to **ALL**. This is a fixed parameter combination for pushing to all devices.
	//
	// - If \\`Target\\` is \\`TBD\\`, set the value to **TBD**. This is a fixed parameter combination for continuous pushes.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
	// The title of the notification or message. The maximum length is 200 bytes.
	//
	// This is required for pushes to Android and HarmonyOS. It is optional for iOS notifications. If you provide a title for an iOS notification:
	//
	// - For iOS 10 and later, the notification displays the title.
	//
	// - For iOS 8.2 to iOS 9.x, the title replaces the app name in the notification.
	//
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Specifies whether to automatically truncate titles and content that are too long.
	//
	// > This only applies to vendor channels that have explicit limits on title and content length. It does not apply to channels like APNs, Huawei, and Honor, which only limit the total request body size.
	//
	// example:
	//
	// false
	Trim *bool `json:"Trim,omitempty" xml:"Trim,omitempty"`
	// iOS notifications are sent through APNs. Specify the environment.
	//
	// - **DEV**: The development environment. Use this for apps installed and debugged directly from Xcode.
	//
	// - **PRODUCT**: The production environment. Use this for apps distributed through the App Store, TestFlight, Ad Hoc, or enterprise distribution.
	//
	// example:
	//
	// DEV
	IOSApnsEnv *string `json:"iOSApnsEnv,omitempty" xml:"iOSApnsEnv,omitempty"`
	// The badge number on the top-right corner of the app icon on iOS.
	//
	// > If \\`iOSBadgeAutoIncrement\\` is set to \\`true\\`, this parameter must be empty.
	//
	// example:
	//
	// 0
	IOSBadge *int32 `json:"iOSBadge,omitempty" xml:"iOSBadge,omitempty"`
	// Specifies whether to enable the auto-increment feature for the badge number. The default value is \\`false\\`.
	//
	// > When this is \\`true\\`, \\`iOSBadge\\` must be empty.
	//
	// The auto-increment feature is managed by the push server, which maintains a badge count for each device. This requires SDK version 1.9.5 or later. The user must also actively sync the badge number to the server.
	//
	// example:
	//
	// true
	IOSBadgeAutoIncrement *bool `json:"iOSBadgeAutoIncrement,omitempty" xml:"iOSBadgeAutoIncrement,omitempty"`
	// The extended properties of the iOS notification.
	//
	// For iOS 10 and later, specify the resource URL for a rich push notification, such as \\`{"attachment": "https\\://xxxx.xxx/notification_pic.png"}\\`. This parameter must be in JSON map format to avoid parsing errors.
	//
	// example:
	//
	// {"attachment": "https://xxxx.xxx/notification_pic.png"}
	IOSExtParameters *string `json:"iOSExtParameters,omitempty" xml:"iOSExtParameters,omitempty"`
	// The interruption level. Valid values:
	//
	// - **passive**: The system adds the notification to the notification list without lighting up the screen or playing a sound.
	//
	// - **active**: The system displays the notification immediately, lights up the screen, and can play a sound.
	//
	// - **time-sensitive**: The system presents the notification immediately, lights up the screen, and can play a sound, but it does not break through system notification controls.
	//
	// - **critical**: The system displays the notification immediately, lights up the screen, and plays a sound, bypassing the mute switch.
	//
	// example:
	//
	// active
	IOSInterruptionLevel *string `json:"iOSInterruptionLevel,omitempty" xml:"iOSInterruptionLevel,omitempty"`
	// A JSON string containing static pass-through parameters for Dynamic Island pushes. It includes static, custom user information, such as product numbers and order details.
	//
	// > This is required when \\`iOSLiveActivityEvent\\` is \\`start\\`.
	//
	// example:
	//
	// {"orderId": "12345", "product": "Shoes"}
	IOSLiveActivityAttributes *string `json:"iOSLiveActivityAttributes,omitempty" xml:"iOSLiveActivityAttributes,omitempty"`
	// The type of Live Activity to start.
	//
	// > This is required when \\`iOSLiveActivityEvent\\` is \\`start\\`.
	//
	// example:
	//
	// OrderActivityAttributes
	IOSLiveActivityAttributesType *string `json:"iOSLiveActivityAttributesType,omitempty" xml:"iOSLiveActivityAttributesType,omitempty"`
	// Dynamic pass-through parameters for Dynamic Island pushes. It includes real-time updates, such as price or inventory changes.
	//
	// example:
	//
	// {"status": "delivered", "estimatedArrival": "2023-12-31T12:00:00Z"}
	IOSLiveActivityContentState *string `json:"iOSLiveActivityContentState,omitempty" xml:"iOSLiveActivityContentState,omitempty"`
	// A UNIX timestamp in seconds. The ended Live Activity remains on the lock screen until this specified time. The maximum duration is 4 hours.
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
	// The Live Activity ID reported by the device to your server. This is the unique identifier for the Live Activity.
	//
	// example:
	//
	// 66B94673-B32E-4CA7-863C-3E523054FD46
	IOSLiveActivityId *string `json:"iOSLiveActivityId,omitempty" xml:"iOSLiveActivityId,omitempty"`
	// A UNIX timestamp in seconds. Marks the time when the activity\\"s content becomes outdated.
	//
	// example:
	//
	// 1743131967
	IOSLiveActivityStaleDate *int64 `json:"iOSLiveActivityStaleDate,omitempty" xml:"iOSLiveActivityStaleDate,omitempty"`
	// The sound for an iOS notification. Specify the name of an audio file located in the app bundle or the \\`Library/Sounds\\` directory of the sandbox. For more information, see [How to set notification sounds for iOS pushes](https://help.aliyun.com/document_detail/48906.html).
	//
	// If you specify an empty string (""), the notification is silent. If you do not set this parameter, the default system sound is used.
	//
	// example:
	//
	// ""
	IOSMusic *string `json:"iOSMusic,omitempty" xml:"iOSMusic,omitempty"`
	// The flag for the iOS notification content extension (iOS 10+). If set to \\`true\\`, an APNs notification can be processed by the extension before it is displayed. This must be set to \\`true\\` for silent notifications.
	//
	// example:
	//
	// true
	IOSMutableContent *bool `json:"iOSMutableContent,omitempty" xml:"iOSMutableContent,omitempty"`
	// Specifies the iOS notification category (iOS 10+).
	//
	// example:
	//
	// ios
	IOSNotificationCategory *string `json:"iOSNotificationCategory,omitempty" xml:"iOSNotificationCategory,omitempty"`
	// If a device receives multiple notifications with the same \\`CollapseId\\`, they are merged into a single notification. If the device is offline and receives consecutive notifications with the same \\`CollapseId\\`, only one is shown in the notification bar. This parameter is supported on iOS 10 and later.
	//
	// example:
	//
	// ZD2011
	IOSNotificationCollapseId *string `json:"iOSNotificationCollapseId,omitempty" xml:"iOSNotificationCollapseId,omitempty"`
	// Groups iOS remote notifications using this property. It marks the identifier for the collapsed group.
	//
	// This is supported only on iOS 12.0 and later.
	//
	// example:
	//
	// abc
	IOSNotificationThreadId *string `json:"iOSNotificationThreadId,omitempty" xml:"iOSNotificationThreadId,omitempty"`
	// The score for highlighting the summary. The value must be a floating-point number between 0 and 1.
	//
	// example:
	//
	// 0.01
	IOSRelevanceScore *float64 `json:"iOSRelevanceScore,omitempty" xml:"iOSRelevanceScore,omitempty"`
	// If a device is offline when a message is pushed (meaning the persistent connection to the Mobile Push server is down), the push is sent once as a notification through Apple\\"s APNs channel.
	//
	// > Converting offline messages to notifications is only supported in the production environment.
	//
	// example:
	//
	// true
	IOSRemind *bool `json:"iOSRemind,omitempty" xml:"iOSRemind,omitempty"`
	// The content of the iOS notification used when a message is converted to a notification. This is valid only when \\`iOSApnsEnv\\` is \\`PRODUCT\\` and \\`iOSRemind\\` is \\`true\\`.
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
	// The subtitle of the iOS notification (iOS 10+).
	//
	// example:
	//
	// su\\"b
	IOSSubtitle *string `json:"iOSSubtitle,omitempty" xml:"iOSSubtitle,omitempty"`
}

func (s PushRequest) String() string {
	return dara.Prettify(s)
}

func (s PushRequest) GoString() string {
	return s.String()
}

func (s *PushRequest) GetAndroidActivity() *string {
	return s.AndroidActivity
}

func (s *PushRequest) GetAndroidBadgeAddNum() *int32 {
	return s.AndroidBadgeAddNum
}

func (s *PushRequest) GetAndroidBadgeClass() *string {
	return s.AndroidBadgeClass
}

func (s *PushRequest) GetAndroidBadgeSetNum() *int32 {
	return s.AndroidBadgeSetNum
}

func (s *PushRequest) GetAndroidBigBody() *string {
	return s.AndroidBigBody
}

func (s *PushRequest) GetAndroidBigPictureUrl() *string {
	return s.AndroidBigPictureUrl
}

func (s *PushRequest) GetAndroidBigTitle() *string {
	return s.AndroidBigTitle
}

func (s *PushRequest) GetAndroidExtParameters() *string {
	return s.AndroidExtParameters
}

func (s *PushRequest) GetAndroidHonorTargetUserType() *int32 {
	return s.AndroidHonorTargetUserType
}

func (s *PushRequest) GetAndroidHuaweiBusinessType() *int32 {
	return s.AndroidHuaweiBusinessType
}

func (s *PushRequest) GetAndroidHuaweiLiveNotificationPayload() *string {
	return s.AndroidHuaweiLiveNotificationPayload
}

func (s *PushRequest) GetAndroidHuaweiReceiptId() *string {
	return s.AndroidHuaweiReceiptId
}

func (s *PushRequest) GetAndroidHuaweiTargetUserType() *int32 {
	return s.AndroidHuaweiTargetUserType
}

func (s *PushRequest) GetAndroidImageUrl() *string {
	return s.AndroidImageUrl
}

func (s *PushRequest) GetAndroidInboxBody() *string {
	return s.AndroidInboxBody
}

func (s *PushRequest) GetAndroidMeizuNoticeMsgType() *int32 {
	return s.AndroidMeizuNoticeMsgType
}

func (s *PushRequest) GetAndroidMessageHuaweiCategory() *string {
	return s.AndroidMessageHuaweiCategory
}

func (s *PushRequest) GetAndroidMessageHuaweiUrgency() *string {
	return s.AndroidMessageHuaweiUrgency
}

func (s *PushRequest) GetAndroidMessageOppoCategory() *string {
	return s.AndroidMessageOppoCategory
}

func (s *PushRequest) GetAndroidMessageOppoNotifyLevel() *int32 {
	return s.AndroidMessageOppoNotifyLevel
}

func (s *PushRequest) GetAndroidMessageVivoCategory() *string {
	return s.AndroidMessageVivoCategory
}

func (s *PushRequest) GetAndroidMusic() *string {
	return s.AndroidMusic
}

func (s *PushRequest) GetAndroidNotificationBarPriority() *int32 {
	return s.AndroidNotificationBarPriority
}

func (s *PushRequest) GetAndroidNotificationBarType() *int32 {
	return s.AndroidNotificationBarType
}

func (s *PushRequest) GetAndroidNotificationChannel() *string {
	return s.AndroidNotificationChannel
}

func (s *PushRequest) GetAndroidNotificationGroup() *string {
	return s.AndroidNotificationGroup
}

func (s *PushRequest) GetAndroidNotificationHonorChannel() *string {
	return s.AndroidNotificationHonorChannel
}

func (s *PushRequest) GetAndroidNotificationHuaweiChannel() *string {
	return s.AndroidNotificationHuaweiChannel
}

func (s *PushRequest) GetAndroidNotificationNotifyId() *int32 {
	return s.AndroidNotificationNotifyId
}

func (s *PushRequest) GetAndroidNotificationThreadId() *string {
	return s.AndroidNotificationThreadId
}

func (s *PushRequest) GetAndroidNotificationVivoChannel() *string {
	return s.AndroidNotificationVivoChannel
}

func (s *PushRequest) GetAndroidNotificationXiaomiChannel() *string {
	return s.AndroidNotificationXiaomiChannel
}

func (s *PushRequest) GetAndroidNotifyType() *string {
	return s.AndroidNotifyType
}

func (s *PushRequest) GetAndroidOpenType() *string {
	return s.AndroidOpenType
}

func (s *PushRequest) GetAndroidOpenUrl() *string {
	return s.AndroidOpenUrl
}

func (s *PushRequest) GetAndroidOppoDeleteIntentData() *string {
	return s.AndroidOppoDeleteIntentData
}

func (s *PushRequest) GetAndroidOppoIntelligentIntent() *string {
	return s.AndroidOppoIntelligentIntent
}

func (s *PushRequest) GetAndroidOppoIntentEnv() *int32 {
	return s.AndroidOppoIntentEnv
}

func (s *PushRequest) GetAndroidOppoPrivateContentParameters() map[string]*string {
	return s.AndroidOppoPrivateContentParameters
}

func (s *PushRequest) GetAndroidOppoPrivateMsgTemplateId() *string {
	return s.AndroidOppoPrivateMsgTemplateId
}

func (s *PushRequest) GetAndroidOppoPrivateTitleParameters() map[string]*string {
	return s.AndroidOppoPrivateTitleParameters
}

func (s *PushRequest) GetAndroidPopupActivity() *string {
	return s.AndroidPopupActivity
}

func (s *PushRequest) GetAndroidPopupBody() *string {
	return s.AndroidPopupBody
}

func (s *PushRequest) GetAndroidPopupTitle() *string {
	return s.AndroidPopupTitle
}

func (s *PushRequest) GetAndroidRemind() *bool {
	return s.AndroidRemind
}

func (s *PushRequest) GetAndroidRenderStyle() *int32 {
	return s.AndroidRenderStyle
}

func (s *PushRequest) GetAndroidTargetUserType() *int32 {
	return s.AndroidTargetUserType
}

func (s *PushRequest) GetAndroidVivoLiveMessage() *string {
	return s.AndroidVivoLiveMessage
}

func (s *PushRequest) GetAndroidVivoPushMode() *int32 {
	return s.AndroidVivoPushMode
}

func (s *PushRequest) GetAndroidVivoReceiptId() *string {
	return s.AndroidVivoReceiptId
}

func (s *PushRequest) GetAndroidXiaoMiActivity() *string {
	return s.AndroidXiaoMiActivity
}

func (s *PushRequest) GetAndroidXiaoMiNotifyBody() *string {
	return s.AndroidXiaoMiNotifyBody
}

func (s *PushRequest) GetAndroidXiaoMiNotifyTitle() *string {
	return s.AndroidXiaoMiNotifyTitle
}

func (s *PushRequest) GetAndroidXiaomiBigPictureUrl() *string {
	return s.AndroidXiaomiBigPictureUrl
}

func (s *PushRequest) GetAndroidXiaomiFocusParam() *string {
	return s.AndroidXiaomiFocusParam
}

func (s *PushRequest) GetAndroidXiaomiFocusPics() *string {
	return s.AndroidXiaomiFocusPics
}

func (s *PushRequest) GetAndroidXiaomiImageUrl() *string {
	return s.AndroidXiaomiImageUrl
}

func (s *PushRequest) GetAndroidXiaomiTemplateId() *string {
	return s.AndroidXiaomiTemplateId
}

func (s *PushRequest) GetAndroidXiaomiTemplateParams() *string {
	return s.AndroidXiaomiTemplateParams
}

func (s *PushRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *PushRequest) GetBody() *string {
	return s.Body
}

func (s *PushRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *PushRequest) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *PushRequest) GetHarmonyAction() *string {
	return s.HarmonyAction
}

func (s *PushRequest) GetHarmonyActionType() *string {
	return s.HarmonyActionType
}

func (s *PushRequest) GetHarmonyBadgeAddNum() *int32 {
	return s.HarmonyBadgeAddNum
}

func (s *PushRequest) GetHarmonyBadgeSetNum() *int32 {
	return s.HarmonyBadgeSetNum
}

func (s *PushRequest) GetHarmonyCategory() *string {
	return s.HarmonyCategory
}

func (s *PushRequest) GetHarmonyExtParameters() *string {
	return s.HarmonyExtParameters
}

func (s *PushRequest) GetHarmonyExtensionExtraData() *string {
	return s.HarmonyExtensionExtraData
}

func (s *PushRequest) GetHarmonyExtensionPush() *bool {
	return s.HarmonyExtensionPush
}

func (s *PushRequest) GetHarmonyImageUrl() *string {
	return s.HarmonyImageUrl
}

func (s *PushRequest) GetHarmonyInboxContent() *string {
	return s.HarmonyInboxContent
}

func (s *PushRequest) GetHarmonyLiveViewPayload() *string {
	return s.HarmonyLiveViewPayload
}

func (s *PushRequest) GetHarmonyNotificationSlotType() *string {
	return s.HarmonyNotificationSlotType
}

func (s *PushRequest) GetHarmonyNotifyId() *int32 {
	return s.HarmonyNotifyId
}

func (s *PushRequest) GetHarmonyReceiptId() *string {
	return s.HarmonyReceiptId
}

func (s *PushRequest) GetHarmonyRemind() *bool {
	return s.HarmonyRemind
}

func (s *PushRequest) GetHarmonyRemindBody() *string {
	return s.HarmonyRemindBody
}

func (s *PushRequest) GetHarmonyRemindTitle() *string {
	return s.HarmonyRemindTitle
}

func (s *PushRequest) GetHarmonyRenderStyle() *string {
	return s.HarmonyRenderStyle
}

func (s *PushRequest) GetHarmonyTestMessage() *bool {
	return s.HarmonyTestMessage
}

func (s *PushRequest) GetHarmonyUri() *string {
	return s.HarmonyUri
}

func (s *PushRequest) GetIdempotentToken() *string {
	return s.IdempotentToken
}

func (s *PushRequest) GetJobKey() *string {
	return s.JobKey
}

func (s *PushRequest) GetPushTime() *string {
	return s.PushTime
}

func (s *PushRequest) GetPushType() *string {
	return s.PushType
}

func (s *PushRequest) GetSendChannels() *string {
	return s.SendChannels
}

func (s *PushRequest) GetSendSpeed() *int32 {
	return s.SendSpeed
}

func (s *PushRequest) GetSmsDelaySecs() *int32 {
	return s.SmsDelaySecs
}

func (s *PushRequest) GetSmsParams() *string {
	return s.SmsParams
}

func (s *PushRequest) GetSmsSendPolicy() *int32 {
	return s.SmsSendPolicy
}

func (s *PushRequest) GetSmsSignName() *string {
	return s.SmsSignName
}

func (s *PushRequest) GetSmsTemplateName() *string {
	return s.SmsTemplateName
}

func (s *PushRequest) GetStoreOffline() *bool {
	return s.StoreOffline
}

func (s *PushRequest) GetTarget() *string {
	return s.Target
}

func (s *PushRequest) GetTargetValue() *string {
	return s.TargetValue
}

func (s *PushRequest) GetTitle() *string {
	return s.Title
}

func (s *PushRequest) GetTrim() *bool {
	return s.Trim
}

func (s *PushRequest) GetIOSApnsEnv() *string {
	return s.IOSApnsEnv
}

func (s *PushRequest) GetIOSBadge() *int32 {
	return s.IOSBadge
}

func (s *PushRequest) GetIOSBadgeAutoIncrement() *bool {
	return s.IOSBadgeAutoIncrement
}

func (s *PushRequest) GetIOSExtParameters() *string {
	return s.IOSExtParameters
}

func (s *PushRequest) GetIOSInterruptionLevel() *string {
	return s.IOSInterruptionLevel
}

func (s *PushRequest) GetIOSLiveActivityAttributes() *string {
	return s.IOSLiveActivityAttributes
}

func (s *PushRequest) GetIOSLiveActivityAttributesType() *string {
	return s.IOSLiveActivityAttributesType
}

func (s *PushRequest) GetIOSLiveActivityContentState() *string {
	return s.IOSLiveActivityContentState
}

func (s *PushRequest) GetIOSLiveActivityDismissalDate() *int64 {
	return s.IOSLiveActivityDismissalDate
}

func (s *PushRequest) GetIOSLiveActivityEvent() *string {
	return s.IOSLiveActivityEvent
}

func (s *PushRequest) GetIOSLiveActivityId() *string {
	return s.IOSLiveActivityId
}

func (s *PushRequest) GetIOSLiveActivityStaleDate() *int64 {
	return s.IOSLiveActivityStaleDate
}

func (s *PushRequest) GetIOSMusic() *string {
	return s.IOSMusic
}

func (s *PushRequest) GetIOSMutableContent() *bool {
	return s.IOSMutableContent
}

func (s *PushRequest) GetIOSNotificationCategory() *string {
	return s.IOSNotificationCategory
}

func (s *PushRequest) GetIOSNotificationCollapseId() *string {
	return s.IOSNotificationCollapseId
}

func (s *PushRequest) GetIOSNotificationThreadId() *string {
	return s.IOSNotificationThreadId
}

func (s *PushRequest) GetIOSRelevanceScore() *float64 {
	return s.IOSRelevanceScore
}

func (s *PushRequest) GetIOSRemind() *bool {
	return s.IOSRemind
}

func (s *PushRequest) GetIOSRemindBody() *string {
	return s.IOSRemindBody
}

func (s *PushRequest) GetIOSSilentNotification() *bool {
	return s.IOSSilentNotification
}

func (s *PushRequest) GetIOSSubtitle() *string {
	return s.IOSSubtitle
}

func (s *PushRequest) SetAndroidActivity(v string) *PushRequest {
	s.AndroidActivity = &v
	return s
}

func (s *PushRequest) SetAndroidBadgeAddNum(v int32) *PushRequest {
	s.AndroidBadgeAddNum = &v
	return s
}

func (s *PushRequest) SetAndroidBadgeClass(v string) *PushRequest {
	s.AndroidBadgeClass = &v
	return s
}

func (s *PushRequest) SetAndroidBadgeSetNum(v int32) *PushRequest {
	s.AndroidBadgeSetNum = &v
	return s
}

func (s *PushRequest) SetAndroidBigBody(v string) *PushRequest {
	s.AndroidBigBody = &v
	return s
}

func (s *PushRequest) SetAndroidBigPictureUrl(v string) *PushRequest {
	s.AndroidBigPictureUrl = &v
	return s
}

func (s *PushRequest) SetAndroidBigTitle(v string) *PushRequest {
	s.AndroidBigTitle = &v
	return s
}

func (s *PushRequest) SetAndroidExtParameters(v string) *PushRequest {
	s.AndroidExtParameters = &v
	return s
}

func (s *PushRequest) SetAndroidHonorTargetUserType(v int32) *PushRequest {
	s.AndroidHonorTargetUserType = &v
	return s
}

func (s *PushRequest) SetAndroidHuaweiBusinessType(v int32) *PushRequest {
	s.AndroidHuaweiBusinessType = &v
	return s
}

func (s *PushRequest) SetAndroidHuaweiLiveNotificationPayload(v string) *PushRequest {
	s.AndroidHuaweiLiveNotificationPayload = &v
	return s
}

func (s *PushRequest) SetAndroidHuaweiReceiptId(v string) *PushRequest {
	s.AndroidHuaweiReceiptId = &v
	return s
}

func (s *PushRequest) SetAndroidHuaweiTargetUserType(v int32) *PushRequest {
	s.AndroidHuaweiTargetUserType = &v
	return s
}

func (s *PushRequest) SetAndroidImageUrl(v string) *PushRequest {
	s.AndroidImageUrl = &v
	return s
}

func (s *PushRequest) SetAndroidInboxBody(v string) *PushRequest {
	s.AndroidInboxBody = &v
	return s
}

func (s *PushRequest) SetAndroidMeizuNoticeMsgType(v int32) *PushRequest {
	s.AndroidMeizuNoticeMsgType = &v
	return s
}

func (s *PushRequest) SetAndroidMessageHuaweiCategory(v string) *PushRequest {
	s.AndroidMessageHuaweiCategory = &v
	return s
}

func (s *PushRequest) SetAndroidMessageHuaweiUrgency(v string) *PushRequest {
	s.AndroidMessageHuaweiUrgency = &v
	return s
}

func (s *PushRequest) SetAndroidMessageOppoCategory(v string) *PushRequest {
	s.AndroidMessageOppoCategory = &v
	return s
}

func (s *PushRequest) SetAndroidMessageOppoNotifyLevel(v int32) *PushRequest {
	s.AndroidMessageOppoNotifyLevel = &v
	return s
}

func (s *PushRequest) SetAndroidMessageVivoCategory(v string) *PushRequest {
	s.AndroidMessageVivoCategory = &v
	return s
}

func (s *PushRequest) SetAndroidMusic(v string) *PushRequest {
	s.AndroidMusic = &v
	return s
}

func (s *PushRequest) SetAndroidNotificationBarPriority(v int32) *PushRequest {
	s.AndroidNotificationBarPriority = &v
	return s
}

func (s *PushRequest) SetAndroidNotificationBarType(v int32) *PushRequest {
	s.AndroidNotificationBarType = &v
	return s
}

func (s *PushRequest) SetAndroidNotificationChannel(v string) *PushRequest {
	s.AndroidNotificationChannel = &v
	return s
}

func (s *PushRequest) SetAndroidNotificationGroup(v string) *PushRequest {
	s.AndroidNotificationGroup = &v
	return s
}

func (s *PushRequest) SetAndroidNotificationHonorChannel(v string) *PushRequest {
	s.AndroidNotificationHonorChannel = &v
	return s
}

func (s *PushRequest) SetAndroidNotificationHuaweiChannel(v string) *PushRequest {
	s.AndroidNotificationHuaweiChannel = &v
	return s
}

func (s *PushRequest) SetAndroidNotificationNotifyId(v int32) *PushRequest {
	s.AndroidNotificationNotifyId = &v
	return s
}

func (s *PushRequest) SetAndroidNotificationThreadId(v string) *PushRequest {
	s.AndroidNotificationThreadId = &v
	return s
}

func (s *PushRequest) SetAndroidNotificationVivoChannel(v string) *PushRequest {
	s.AndroidNotificationVivoChannel = &v
	return s
}

func (s *PushRequest) SetAndroidNotificationXiaomiChannel(v string) *PushRequest {
	s.AndroidNotificationXiaomiChannel = &v
	return s
}

func (s *PushRequest) SetAndroidNotifyType(v string) *PushRequest {
	s.AndroidNotifyType = &v
	return s
}

func (s *PushRequest) SetAndroidOpenType(v string) *PushRequest {
	s.AndroidOpenType = &v
	return s
}

func (s *PushRequest) SetAndroidOpenUrl(v string) *PushRequest {
	s.AndroidOpenUrl = &v
	return s
}

func (s *PushRequest) SetAndroidOppoDeleteIntentData(v string) *PushRequest {
	s.AndroidOppoDeleteIntentData = &v
	return s
}

func (s *PushRequest) SetAndroidOppoIntelligentIntent(v string) *PushRequest {
	s.AndroidOppoIntelligentIntent = &v
	return s
}

func (s *PushRequest) SetAndroidOppoIntentEnv(v int32) *PushRequest {
	s.AndroidOppoIntentEnv = &v
	return s
}

func (s *PushRequest) SetAndroidOppoPrivateContentParameters(v map[string]*string) *PushRequest {
	s.AndroidOppoPrivateContentParameters = v
	return s
}

func (s *PushRequest) SetAndroidOppoPrivateMsgTemplateId(v string) *PushRequest {
	s.AndroidOppoPrivateMsgTemplateId = &v
	return s
}

func (s *PushRequest) SetAndroidOppoPrivateTitleParameters(v map[string]*string) *PushRequest {
	s.AndroidOppoPrivateTitleParameters = v
	return s
}

func (s *PushRequest) SetAndroidPopupActivity(v string) *PushRequest {
	s.AndroidPopupActivity = &v
	return s
}

func (s *PushRequest) SetAndroidPopupBody(v string) *PushRequest {
	s.AndroidPopupBody = &v
	return s
}

func (s *PushRequest) SetAndroidPopupTitle(v string) *PushRequest {
	s.AndroidPopupTitle = &v
	return s
}

func (s *PushRequest) SetAndroidRemind(v bool) *PushRequest {
	s.AndroidRemind = &v
	return s
}

func (s *PushRequest) SetAndroidRenderStyle(v int32) *PushRequest {
	s.AndroidRenderStyle = &v
	return s
}

func (s *PushRequest) SetAndroidTargetUserType(v int32) *PushRequest {
	s.AndroidTargetUserType = &v
	return s
}

func (s *PushRequest) SetAndroidVivoLiveMessage(v string) *PushRequest {
	s.AndroidVivoLiveMessage = &v
	return s
}

func (s *PushRequest) SetAndroidVivoPushMode(v int32) *PushRequest {
	s.AndroidVivoPushMode = &v
	return s
}

func (s *PushRequest) SetAndroidVivoReceiptId(v string) *PushRequest {
	s.AndroidVivoReceiptId = &v
	return s
}

func (s *PushRequest) SetAndroidXiaoMiActivity(v string) *PushRequest {
	s.AndroidXiaoMiActivity = &v
	return s
}

func (s *PushRequest) SetAndroidXiaoMiNotifyBody(v string) *PushRequest {
	s.AndroidXiaoMiNotifyBody = &v
	return s
}

func (s *PushRequest) SetAndroidXiaoMiNotifyTitle(v string) *PushRequest {
	s.AndroidXiaoMiNotifyTitle = &v
	return s
}

func (s *PushRequest) SetAndroidXiaomiBigPictureUrl(v string) *PushRequest {
	s.AndroidXiaomiBigPictureUrl = &v
	return s
}

func (s *PushRequest) SetAndroidXiaomiFocusParam(v string) *PushRequest {
	s.AndroidXiaomiFocusParam = &v
	return s
}

func (s *PushRequest) SetAndroidXiaomiFocusPics(v string) *PushRequest {
	s.AndroidXiaomiFocusPics = &v
	return s
}

func (s *PushRequest) SetAndroidXiaomiImageUrl(v string) *PushRequest {
	s.AndroidXiaomiImageUrl = &v
	return s
}

func (s *PushRequest) SetAndroidXiaomiTemplateId(v string) *PushRequest {
	s.AndroidXiaomiTemplateId = &v
	return s
}

func (s *PushRequest) SetAndroidXiaomiTemplateParams(v string) *PushRequest {
	s.AndroidXiaomiTemplateParams = &v
	return s
}

func (s *PushRequest) SetAppKey(v int64) *PushRequest {
	s.AppKey = &v
	return s
}

func (s *PushRequest) SetBody(v string) *PushRequest {
	s.Body = &v
	return s
}

func (s *PushRequest) SetDeviceType(v string) *PushRequest {
	s.DeviceType = &v
	return s
}

func (s *PushRequest) SetExpireTime(v string) *PushRequest {
	s.ExpireTime = &v
	return s
}

func (s *PushRequest) SetHarmonyAction(v string) *PushRequest {
	s.HarmonyAction = &v
	return s
}

func (s *PushRequest) SetHarmonyActionType(v string) *PushRequest {
	s.HarmonyActionType = &v
	return s
}

func (s *PushRequest) SetHarmonyBadgeAddNum(v int32) *PushRequest {
	s.HarmonyBadgeAddNum = &v
	return s
}

func (s *PushRequest) SetHarmonyBadgeSetNum(v int32) *PushRequest {
	s.HarmonyBadgeSetNum = &v
	return s
}

func (s *PushRequest) SetHarmonyCategory(v string) *PushRequest {
	s.HarmonyCategory = &v
	return s
}

func (s *PushRequest) SetHarmonyExtParameters(v string) *PushRequest {
	s.HarmonyExtParameters = &v
	return s
}

func (s *PushRequest) SetHarmonyExtensionExtraData(v string) *PushRequest {
	s.HarmonyExtensionExtraData = &v
	return s
}

func (s *PushRequest) SetHarmonyExtensionPush(v bool) *PushRequest {
	s.HarmonyExtensionPush = &v
	return s
}

func (s *PushRequest) SetHarmonyImageUrl(v string) *PushRequest {
	s.HarmonyImageUrl = &v
	return s
}

func (s *PushRequest) SetHarmonyInboxContent(v string) *PushRequest {
	s.HarmonyInboxContent = &v
	return s
}

func (s *PushRequest) SetHarmonyLiveViewPayload(v string) *PushRequest {
	s.HarmonyLiveViewPayload = &v
	return s
}

func (s *PushRequest) SetHarmonyNotificationSlotType(v string) *PushRequest {
	s.HarmonyNotificationSlotType = &v
	return s
}

func (s *PushRequest) SetHarmonyNotifyId(v int32) *PushRequest {
	s.HarmonyNotifyId = &v
	return s
}

func (s *PushRequest) SetHarmonyReceiptId(v string) *PushRequest {
	s.HarmonyReceiptId = &v
	return s
}

func (s *PushRequest) SetHarmonyRemind(v bool) *PushRequest {
	s.HarmonyRemind = &v
	return s
}

func (s *PushRequest) SetHarmonyRemindBody(v string) *PushRequest {
	s.HarmonyRemindBody = &v
	return s
}

func (s *PushRequest) SetHarmonyRemindTitle(v string) *PushRequest {
	s.HarmonyRemindTitle = &v
	return s
}

func (s *PushRequest) SetHarmonyRenderStyle(v string) *PushRequest {
	s.HarmonyRenderStyle = &v
	return s
}

func (s *PushRequest) SetHarmonyTestMessage(v bool) *PushRequest {
	s.HarmonyTestMessage = &v
	return s
}

func (s *PushRequest) SetHarmonyUri(v string) *PushRequest {
	s.HarmonyUri = &v
	return s
}

func (s *PushRequest) SetIdempotentToken(v string) *PushRequest {
	s.IdempotentToken = &v
	return s
}

func (s *PushRequest) SetJobKey(v string) *PushRequest {
	s.JobKey = &v
	return s
}

func (s *PushRequest) SetPushTime(v string) *PushRequest {
	s.PushTime = &v
	return s
}

func (s *PushRequest) SetPushType(v string) *PushRequest {
	s.PushType = &v
	return s
}

func (s *PushRequest) SetSendChannels(v string) *PushRequest {
	s.SendChannels = &v
	return s
}

func (s *PushRequest) SetSendSpeed(v int32) *PushRequest {
	s.SendSpeed = &v
	return s
}

func (s *PushRequest) SetSmsDelaySecs(v int32) *PushRequest {
	s.SmsDelaySecs = &v
	return s
}

func (s *PushRequest) SetSmsParams(v string) *PushRequest {
	s.SmsParams = &v
	return s
}

func (s *PushRequest) SetSmsSendPolicy(v int32) *PushRequest {
	s.SmsSendPolicy = &v
	return s
}

func (s *PushRequest) SetSmsSignName(v string) *PushRequest {
	s.SmsSignName = &v
	return s
}

func (s *PushRequest) SetSmsTemplateName(v string) *PushRequest {
	s.SmsTemplateName = &v
	return s
}

func (s *PushRequest) SetStoreOffline(v bool) *PushRequest {
	s.StoreOffline = &v
	return s
}

func (s *PushRequest) SetTarget(v string) *PushRequest {
	s.Target = &v
	return s
}

func (s *PushRequest) SetTargetValue(v string) *PushRequest {
	s.TargetValue = &v
	return s
}

func (s *PushRequest) SetTitle(v string) *PushRequest {
	s.Title = &v
	return s
}

func (s *PushRequest) SetTrim(v bool) *PushRequest {
	s.Trim = &v
	return s
}

func (s *PushRequest) SetIOSApnsEnv(v string) *PushRequest {
	s.IOSApnsEnv = &v
	return s
}

func (s *PushRequest) SetIOSBadge(v int32) *PushRequest {
	s.IOSBadge = &v
	return s
}

func (s *PushRequest) SetIOSBadgeAutoIncrement(v bool) *PushRequest {
	s.IOSBadgeAutoIncrement = &v
	return s
}

func (s *PushRequest) SetIOSExtParameters(v string) *PushRequest {
	s.IOSExtParameters = &v
	return s
}

func (s *PushRequest) SetIOSInterruptionLevel(v string) *PushRequest {
	s.IOSInterruptionLevel = &v
	return s
}

func (s *PushRequest) SetIOSLiveActivityAttributes(v string) *PushRequest {
	s.IOSLiveActivityAttributes = &v
	return s
}

func (s *PushRequest) SetIOSLiveActivityAttributesType(v string) *PushRequest {
	s.IOSLiveActivityAttributesType = &v
	return s
}

func (s *PushRequest) SetIOSLiveActivityContentState(v string) *PushRequest {
	s.IOSLiveActivityContentState = &v
	return s
}

func (s *PushRequest) SetIOSLiveActivityDismissalDate(v int64) *PushRequest {
	s.IOSLiveActivityDismissalDate = &v
	return s
}

func (s *PushRequest) SetIOSLiveActivityEvent(v string) *PushRequest {
	s.IOSLiveActivityEvent = &v
	return s
}

func (s *PushRequest) SetIOSLiveActivityId(v string) *PushRequest {
	s.IOSLiveActivityId = &v
	return s
}

func (s *PushRequest) SetIOSLiveActivityStaleDate(v int64) *PushRequest {
	s.IOSLiveActivityStaleDate = &v
	return s
}

func (s *PushRequest) SetIOSMusic(v string) *PushRequest {
	s.IOSMusic = &v
	return s
}

func (s *PushRequest) SetIOSMutableContent(v bool) *PushRequest {
	s.IOSMutableContent = &v
	return s
}

func (s *PushRequest) SetIOSNotificationCategory(v string) *PushRequest {
	s.IOSNotificationCategory = &v
	return s
}

func (s *PushRequest) SetIOSNotificationCollapseId(v string) *PushRequest {
	s.IOSNotificationCollapseId = &v
	return s
}

func (s *PushRequest) SetIOSNotificationThreadId(v string) *PushRequest {
	s.IOSNotificationThreadId = &v
	return s
}

func (s *PushRequest) SetIOSRelevanceScore(v float64) *PushRequest {
	s.IOSRelevanceScore = &v
	return s
}

func (s *PushRequest) SetIOSRemind(v bool) *PushRequest {
	s.IOSRemind = &v
	return s
}

func (s *PushRequest) SetIOSRemindBody(v string) *PushRequest {
	s.IOSRemindBody = &v
	return s
}

func (s *PushRequest) SetIOSSilentNotification(v bool) *PushRequest {
	s.IOSSilentNotification = &v
	return s
}

func (s *PushRequest) SetIOSSubtitle(v string) *PushRequest {
	s.IOSSubtitle = &v
	return s
}

func (s *PushRequest) Validate() error {
	return dara.Validate(s)
}
