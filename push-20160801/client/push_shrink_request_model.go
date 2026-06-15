// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidActivity(v string) *PushShrinkRequest
	GetAndroidActivity() *string
	SetAndroidBadgeAddNum(v int32) *PushShrinkRequest
	GetAndroidBadgeAddNum() *int32
	SetAndroidBadgeClass(v string) *PushShrinkRequest
	GetAndroidBadgeClass() *string
	SetAndroidBadgeSetNum(v int32) *PushShrinkRequest
	GetAndroidBadgeSetNum() *int32
	SetAndroidBigBody(v string) *PushShrinkRequest
	GetAndroidBigBody() *string
	SetAndroidBigPictureUrl(v string) *PushShrinkRequest
	GetAndroidBigPictureUrl() *string
	SetAndroidBigTitle(v string) *PushShrinkRequest
	GetAndroidBigTitle() *string
	SetAndroidExtParameters(v string) *PushShrinkRequest
	GetAndroidExtParameters() *string
	SetAndroidHonorTargetUserType(v int32) *PushShrinkRequest
	GetAndroidHonorTargetUserType() *int32
	SetAndroidHuaweiBusinessType(v int32) *PushShrinkRequest
	GetAndroidHuaweiBusinessType() *int32
	SetAndroidHuaweiLiveNotificationPayload(v string) *PushShrinkRequest
	GetAndroidHuaweiLiveNotificationPayload() *string
	SetAndroidHuaweiReceiptId(v string) *PushShrinkRequest
	GetAndroidHuaweiReceiptId() *string
	SetAndroidHuaweiTargetUserType(v int32) *PushShrinkRequest
	GetAndroidHuaweiTargetUserType() *int32
	SetAndroidImageUrl(v string) *PushShrinkRequest
	GetAndroidImageUrl() *string
	SetAndroidInboxBody(v string) *PushShrinkRequest
	GetAndroidInboxBody() *string
	SetAndroidMeizuNoticeMsgType(v int32) *PushShrinkRequest
	GetAndroidMeizuNoticeMsgType() *int32
	SetAndroidMessageHuaweiCategory(v string) *PushShrinkRequest
	GetAndroidMessageHuaweiCategory() *string
	SetAndroidMessageHuaweiUrgency(v string) *PushShrinkRequest
	GetAndroidMessageHuaweiUrgency() *string
	SetAndroidMessageOppoCategory(v string) *PushShrinkRequest
	GetAndroidMessageOppoCategory() *string
	SetAndroidMessageOppoNotifyLevel(v int32) *PushShrinkRequest
	GetAndroidMessageOppoNotifyLevel() *int32
	SetAndroidMessageVivoCategory(v string) *PushShrinkRequest
	GetAndroidMessageVivoCategory() *string
	SetAndroidMusic(v string) *PushShrinkRequest
	GetAndroidMusic() *string
	SetAndroidNotificationBarPriority(v int32) *PushShrinkRequest
	GetAndroidNotificationBarPriority() *int32
	SetAndroidNotificationBarType(v int32) *PushShrinkRequest
	GetAndroidNotificationBarType() *int32
	SetAndroidNotificationChannel(v string) *PushShrinkRequest
	GetAndroidNotificationChannel() *string
	SetAndroidNotificationGroup(v string) *PushShrinkRequest
	GetAndroidNotificationGroup() *string
	SetAndroidNotificationHonorChannel(v string) *PushShrinkRequest
	GetAndroidNotificationHonorChannel() *string
	SetAndroidNotificationHuaweiChannel(v string) *PushShrinkRequest
	GetAndroidNotificationHuaweiChannel() *string
	SetAndroidNotificationNotifyId(v int32) *PushShrinkRequest
	GetAndroidNotificationNotifyId() *int32
	SetAndroidNotificationThreadId(v string) *PushShrinkRequest
	GetAndroidNotificationThreadId() *string
	SetAndroidNotificationVivoChannel(v string) *PushShrinkRequest
	GetAndroidNotificationVivoChannel() *string
	SetAndroidNotificationXiaomiChannel(v string) *PushShrinkRequest
	GetAndroidNotificationXiaomiChannel() *string
	SetAndroidNotifyType(v string) *PushShrinkRequest
	GetAndroidNotifyType() *string
	SetAndroidOpenType(v string) *PushShrinkRequest
	GetAndroidOpenType() *string
	SetAndroidOpenUrl(v string) *PushShrinkRequest
	GetAndroidOpenUrl() *string
	SetAndroidOppoDeleteIntentData(v string) *PushShrinkRequest
	GetAndroidOppoDeleteIntentData() *string
	SetAndroidOppoIntelligentIntent(v string) *PushShrinkRequest
	GetAndroidOppoIntelligentIntent() *string
	SetAndroidOppoIntentEnv(v int32) *PushShrinkRequest
	GetAndroidOppoIntentEnv() *int32
	SetAndroidOppoPrivateContentParametersShrink(v string) *PushShrinkRequest
	GetAndroidOppoPrivateContentParametersShrink() *string
	SetAndroidOppoPrivateMsgTemplateId(v string) *PushShrinkRequest
	GetAndroidOppoPrivateMsgTemplateId() *string
	SetAndroidOppoPrivateTitleParametersShrink(v string) *PushShrinkRequest
	GetAndroidOppoPrivateTitleParametersShrink() *string
	SetAndroidPopupActivity(v string) *PushShrinkRequest
	GetAndroidPopupActivity() *string
	SetAndroidPopupBody(v string) *PushShrinkRequest
	GetAndroidPopupBody() *string
	SetAndroidPopupTitle(v string) *PushShrinkRequest
	GetAndroidPopupTitle() *string
	SetAndroidRemind(v bool) *PushShrinkRequest
	GetAndroidRemind() *bool
	SetAndroidRenderStyle(v int32) *PushShrinkRequest
	GetAndroidRenderStyle() *int32
	SetAndroidTargetUserType(v int32) *PushShrinkRequest
	GetAndroidTargetUserType() *int32
	SetAndroidVivoLiveMessage(v string) *PushShrinkRequest
	GetAndroidVivoLiveMessage() *string
	SetAndroidVivoPushMode(v int32) *PushShrinkRequest
	GetAndroidVivoPushMode() *int32
	SetAndroidVivoReceiptId(v string) *PushShrinkRequest
	GetAndroidVivoReceiptId() *string
	SetAndroidXiaoMiActivity(v string) *PushShrinkRequest
	GetAndroidXiaoMiActivity() *string
	SetAndroidXiaoMiNotifyBody(v string) *PushShrinkRequest
	GetAndroidXiaoMiNotifyBody() *string
	SetAndroidXiaoMiNotifyTitle(v string) *PushShrinkRequest
	GetAndroidXiaoMiNotifyTitle() *string
	SetAndroidXiaomiBigPictureUrl(v string) *PushShrinkRequest
	GetAndroidXiaomiBigPictureUrl() *string
	SetAndroidXiaomiFocusParam(v string) *PushShrinkRequest
	GetAndroidXiaomiFocusParam() *string
	SetAndroidXiaomiFocusPics(v string) *PushShrinkRequest
	GetAndroidXiaomiFocusPics() *string
	SetAndroidXiaomiImageUrl(v string) *PushShrinkRequest
	GetAndroidXiaomiImageUrl() *string
	SetAndroidXiaomiTemplateId(v string) *PushShrinkRequest
	GetAndroidXiaomiTemplateId() *string
	SetAndroidXiaomiTemplateParams(v string) *PushShrinkRequest
	GetAndroidXiaomiTemplateParams() *string
	SetAppKey(v int64) *PushShrinkRequest
	GetAppKey() *int64
	SetBody(v string) *PushShrinkRequest
	GetBody() *string
	SetDeviceType(v string) *PushShrinkRequest
	GetDeviceType() *string
	SetExpireTime(v string) *PushShrinkRequest
	GetExpireTime() *string
	SetHarmonyAction(v string) *PushShrinkRequest
	GetHarmonyAction() *string
	SetHarmonyActionType(v string) *PushShrinkRequest
	GetHarmonyActionType() *string
	SetHarmonyBadgeAddNum(v int32) *PushShrinkRequest
	GetHarmonyBadgeAddNum() *int32
	SetHarmonyBadgeSetNum(v int32) *PushShrinkRequest
	GetHarmonyBadgeSetNum() *int32
	SetHarmonyCategory(v string) *PushShrinkRequest
	GetHarmonyCategory() *string
	SetHarmonyExtParameters(v string) *PushShrinkRequest
	GetHarmonyExtParameters() *string
	SetHarmonyExtensionExtraData(v string) *PushShrinkRequest
	GetHarmonyExtensionExtraData() *string
	SetHarmonyExtensionPush(v bool) *PushShrinkRequest
	GetHarmonyExtensionPush() *bool
	SetHarmonyImageUrl(v string) *PushShrinkRequest
	GetHarmonyImageUrl() *string
	SetHarmonyInboxContent(v string) *PushShrinkRequest
	GetHarmonyInboxContent() *string
	SetHarmonyLiveViewPayload(v string) *PushShrinkRequest
	GetHarmonyLiveViewPayload() *string
	SetHarmonyNotificationSlotType(v string) *PushShrinkRequest
	GetHarmonyNotificationSlotType() *string
	SetHarmonyNotifyId(v int32) *PushShrinkRequest
	GetHarmonyNotifyId() *int32
	SetHarmonyReceiptId(v string) *PushShrinkRequest
	GetHarmonyReceiptId() *string
	SetHarmonyRemind(v bool) *PushShrinkRequest
	GetHarmonyRemind() *bool
	SetHarmonyRemindBody(v string) *PushShrinkRequest
	GetHarmonyRemindBody() *string
	SetHarmonyRemindTitle(v string) *PushShrinkRequest
	GetHarmonyRemindTitle() *string
	SetHarmonyRenderStyle(v string) *PushShrinkRequest
	GetHarmonyRenderStyle() *string
	SetHarmonyTestMessage(v bool) *PushShrinkRequest
	GetHarmonyTestMessage() *bool
	SetHarmonyUri(v string) *PushShrinkRequest
	GetHarmonyUri() *string
	SetIdempotentToken(v string) *PushShrinkRequest
	GetIdempotentToken() *string
	SetJobKey(v string) *PushShrinkRequest
	GetJobKey() *string
	SetPushTime(v string) *PushShrinkRequest
	GetPushTime() *string
	SetPushType(v string) *PushShrinkRequest
	GetPushType() *string
	SetSendChannels(v string) *PushShrinkRequest
	GetSendChannels() *string
	SetSendSpeed(v int32) *PushShrinkRequest
	GetSendSpeed() *int32
	SetSmsDelaySecs(v int32) *PushShrinkRequest
	GetSmsDelaySecs() *int32
	SetSmsParams(v string) *PushShrinkRequest
	GetSmsParams() *string
	SetSmsSendPolicy(v int32) *PushShrinkRequest
	GetSmsSendPolicy() *int32
	SetSmsSignName(v string) *PushShrinkRequest
	GetSmsSignName() *string
	SetSmsTemplateName(v string) *PushShrinkRequest
	GetSmsTemplateName() *string
	SetStoreOffline(v bool) *PushShrinkRequest
	GetStoreOffline() *bool
	SetTarget(v string) *PushShrinkRequest
	GetTarget() *string
	SetTargetValue(v string) *PushShrinkRequest
	GetTargetValue() *string
	SetTitle(v string) *PushShrinkRequest
	GetTitle() *string
	SetTrim(v bool) *PushShrinkRequest
	GetTrim() *bool
	SetIOSApnsEnv(v string) *PushShrinkRequest
	GetIOSApnsEnv() *string
	SetIOSBadge(v int32) *PushShrinkRequest
	GetIOSBadge() *int32
	SetIOSBadgeAutoIncrement(v bool) *PushShrinkRequest
	GetIOSBadgeAutoIncrement() *bool
	SetIOSExtParameters(v string) *PushShrinkRequest
	GetIOSExtParameters() *string
	SetIOSInterruptionLevel(v string) *PushShrinkRequest
	GetIOSInterruptionLevel() *string
	SetIOSLiveActivityAttributes(v string) *PushShrinkRequest
	GetIOSLiveActivityAttributes() *string
	SetIOSLiveActivityAttributesType(v string) *PushShrinkRequest
	GetIOSLiveActivityAttributesType() *string
	SetIOSLiveActivityContentState(v string) *PushShrinkRequest
	GetIOSLiveActivityContentState() *string
	SetIOSLiveActivityDismissalDate(v int64) *PushShrinkRequest
	GetIOSLiveActivityDismissalDate() *int64
	SetIOSLiveActivityEvent(v string) *PushShrinkRequest
	GetIOSLiveActivityEvent() *string
	SetIOSLiveActivityId(v string) *PushShrinkRequest
	GetIOSLiveActivityId() *string
	SetIOSLiveActivityStaleDate(v int64) *PushShrinkRequest
	GetIOSLiveActivityStaleDate() *int64
	SetIOSMusic(v string) *PushShrinkRequest
	GetIOSMusic() *string
	SetIOSMutableContent(v bool) *PushShrinkRequest
	GetIOSMutableContent() *bool
	SetIOSNotificationCategory(v string) *PushShrinkRequest
	GetIOSNotificationCategory() *string
	SetIOSNotificationCollapseId(v string) *PushShrinkRequest
	GetIOSNotificationCollapseId() *string
	SetIOSNotificationThreadId(v string) *PushShrinkRequest
	GetIOSNotificationThreadId() *string
	SetIOSRelevanceScore(v float64) *PushShrinkRequest
	GetIOSRelevanceScore() *float64
	SetIOSRemind(v bool) *PushShrinkRequest
	GetIOSRemind() *bool
	SetIOSRemindBody(v string) *PushShrinkRequest
	GetIOSRemindBody() *string
	SetIOSSilentNotification(v bool) *PushShrinkRequest
	GetIOSSilentNotification() *bool
	SetIOSSubtitle(v string) *PushShrinkRequest
	GetIOSSubtitle() *string
}

type PushShrinkRequest struct {
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
	AndroidOppoPrivateContentParametersShrink *string `json:"AndroidOppoPrivateContentParameters,omitempty" xml:"AndroidOppoPrivateContentParameters,omitempty"`
	// OPPO private message template ID
	//
	// example:
	//
	// 687557242b1634hzefs3d5013
	AndroidOppoPrivateMsgTemplateId *string `json:"AndroidOppoPrivateMsgTemplateId,omitempty" xml:"AndroidOppoPrivateMsgTemplateId,omitempty"`
	// OPPO private message template title parameters
	AndroidOppoPrivateTitleParametersShrink *string `json:"AndroidOppoPrivateTitleParameters,omitempty" xml:"AndroidOppoPrivateTitleParameters,omitempty"`
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

func (s PushShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PushShrinkRequest) GoString() string {
	return s.String()
}

func (s *PushShrinkRequest) GetAndroidActivity() *string {
	return s.AndroidActivity
}

func (s *PushShrinkRequest) GetAndroidBadgeAddNum() *int32 {
	return s.AndroidBadgeAddNum
}

func (s *PushShrinkRequest) GetAndroidBadgeClass() *string {
	return s.AndroidBadgeClass
}

func (s *PushShrinkRequest) GetAndroidBadgeSetNum() *int32 {
	return s.AndroidBadgeSetNum
}

func (s *PushShrinkRequest) GetAndroidBigBody() *string {
	return s.AndroidBigBody
}

func (s *PushShrinkRequest) GetAndroidBigPictureUrl() *string {
	return s.AndroidBigPictureUrl
}

func (s *PushShrinkRequest) GetAndroidBigTitle() *string {
	return s.AndroidBigTitle
}

func (s *PushShrinkRequest) GetAndroidExtParameters() *string {
	return s.AndroidExtParameters
}

func (s *PushShrinkRequest) GetAndroidHonorTargetUserType() *int32 {
	return s.AndroidHonorTargetUserType
}

func (s *PushShrinkRequest) GetAndroidHuaweiBusinessType() *int32 {
	return s.AndroidHuaweiBusinessType
}

func (s *PushShrinkRequest) GetAndroidHuaweiLiveNotificationPayload() *string {
	return s.AndroidHuaweiLiveNotificationPayload
}

func (s *PushShrinkRequest) GetAndroidHuaweiReceiptId() *string {
	return s.AndroidHuaweiReceiptId
}

func (s *PushShrinkRequest) GetAndroidHuaweiTargetUserType() *int32 {
	return s.AndroidHuaweiTargetUserType
}

func (s *PushShrinkRequest) GetAndroidImageUrl() *string {
	return s.AndroidImageUrl
}

func (s *PushShrinkRequest) GetAndroidInboxBody() *string {
	return s.AndroidInboxBody
}

func (s *PushShrinkRequest) GetAndroidMeizuNoticeMsgType() *int32 {
	return s.AndroidMeizuNoticeMsgType
}

func (s *PushShrinkRequest) GetAndroidMessageHuaweiCategory() *string {
	return s.AndroidMessageHuaweiCategory
}

func (s *PushShrinkRequest) GetAndroidMessageHuaweiUrgency() *string {
	return s.AndroidMessageHuaweiUrgency
}

func (s *PushShrinkRequest) GetAndroidMessageOppoCategory() *string {
	return s.AndroidMessageOppoCategory
}

func (s *PushShrinkRequest) GetAndroidMessageOppoNotifyLevel() *int32 {
	return s.AndroidMessageOppoNotifyLevel
}

func (s *PushShrinkRequest) GetAndroidMessageVivoCategory() *string {
	return s.AndroidMessageVivoCategory
}

func (s *PushShrinkRequest) GetAndroidMusic() *string {
	return s.AndroidMusic
}

func (s *PushShrinkRequest) GetAndroidNotificationBarPriority() *int32 {
	return s.AndroidNotificationBarPriority
}

func (s *PushShrinkRequest) GetAndroidNotificationBarType() *int32 {
	return s.AndroidNotificationBarType
}

func (s *PushShrinkRequest) GetAndroidNotificationChannel() *string {
	return s.AndroidNotificationChannel
}

func (s *PushShrinkRequest) GetAndroidNotificationGroup() *string {
	return s.AndroidNotificationGroup
}

func (s *PushShrinkRequest) GetAndroidNotificationHonorChannel() *string {
	return s.AndroidNotificationHonorChannel
}

func (s *PushShrinkRequest) GetAndroidNotificationHuaweiChannel() *string {
	return s.AndroidNotificationHuaweiChannel
}

func (s *PushShrinkRequest) GetAndroidNotificationNotifyId() *int32 {
	return s.AndroidNotificationNotifyId
}

func (s *PushShrinkRequest) GetAndroidNotificationThreadId() *string {
	return s.AndroidNotificationThreadId
}

func (s *PushShrinkRequest) GetAndroidNotificationVivoChannel() *string {
	return s.AndroidNotificationVivoChannel
}

func (s *PushShrinkRequest) GetAndroidNotificationXiaomiChannel() *string {
	return s.AndroidNotificationXiaomiChannel
}

func (s *PushShrinkRequest) GetAndroidNotifyType() *string {
	return s.AndroidNotifyType
}

func (s *PushShrinkRequest) GetAndroidOpenType() *string {
	return s.AndroidOpenType
}

func (s *PushShrinkRequest) GetAndroidOpenUrl() *string {
	return s.AndroidOpenUrl
}

func (s *PushShrinkRequest) GetAndroidOppoDeleteIntentData() *string {
	return s.AndroidOppoDeleteIntentData
}

func (s *PushShrinkRequest) GetAndroidOppoIntelligentIntent() *string {
	return s.AndroidOppoIntelligentIntent
}

func (s *PushShrinkRequest) GetAndroidOppoIntentEnv() *int32 {
	return s.AndroidOppoIntentEnv
}

func (s *PushShrinkRequest) GetAndroidOppoPrivateContentParametersShrink() *string {
	return s.AndroidOppoPrivateContentParametersShrink
}

func (s *PushShrinkRequest) GetAndroidOppoPrivateMsgTemplateId() *string {
	return s.AndroidOppoPrivateMsgTemplateId
}

func (s *PushShrinkRequest) GetAndroidOppoPrivateTitleParametersShrink() *string {
	return s.AndroidOppoPrivateTitleParametersShrink
}

func (s *PushShrinkRequest) GetAndroidPopupActivity() *string {
	return s.AndroidPopupActivity
}

func (s *PushShrinkRequest) GetAndroidPopupBody() *string {
	return s.AndroidPopupBody
}

func (s *PushShrinkRequest) GetAndroidPopupTitle() *string {
	return s.AndroidPopupTitle
}

func (s *PushShrinkRequest) GetAndroidRemind() *bool {
	return s.AndroidRemind
}

func (s *PushShrinkRequest) GetAndroidRenderStyle() *int32 {
	return s.AndroidRenderStyle
}

func (s *PushShrinkRequest) GetAndroidTargetUserType() *int32 {
	return s.AndroidTargetUserType
}

func (s *PushShrinkRequest) GetAndroidVivoLiveMessage() *string {
	return s.AndroidVivoLiveMessage
}

func (s *PushShrinkRequest) GetAndroidVivoPushMode() *int32 {
	return s.AndroidVivoPushMode
}

func (s *PushShrinkRequest) GetAndroidVivoReceiptId() *string {
	return s.AndroidVivoReceiptId
}

func (s *PushShrinkRequest) GetAndroidXiaoMiActivity() *string {
	return s.AndroidXiaoMiActivity
}

func (s *PushShrinkRequest) GetAndroidXiaoMiNotifyBody() *string {
	return s.AndroidXiaoMiNotifyBody
}

func (s *PushShrinkRequest) GetAndroidXiaoMiNotifyTitle() *string {
	return s.AndroidXiaoMiNotifyTitle
}

func (s *PushShrinkRequest) GetAndroidXiaomiBigPictureUrl() *string {
	return s.AndroidXiaomiBigPictureUrl
}

func (s *PushShrinkRequest) GetAndroidXiaomiFocusParam() *string {
	return s.AndroidXiaomiFocusParam
}

func (s *PushShrinkRequest) GetAndroidXiaomiFocusPics() *string {
	return s.AndroidXiaomiFocusPics
}

func (s *PushShrinkRequest) GetAndroidXiaomiImageUrl() *string {
	return s.AndroidXiaomiImageUrl
}

func (s *PushShrinkRequest) GetAndroidXiaomiTemplateId() *string {
	return s.AndroidXiaomiTemplateId
}

func (s *PushShrinkRequest) GetAndroidXiaomiTemplateParams() *string {
	return s.AndroidXiaomiTemplateParams
}

func (s *PushShrinkRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *PushShrinkRequest) GetBody() *string {
	return s.Body
}

func (s *PushShrinkRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *PushShrinkRequest) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *PushShrinkRequest) GetHarmonyAction() *string {
	return s.HarmonyAction
}

func (s *PushShrinkRequest) GetHarmonyActionType() *string {
	return s.HarmonyActionType
}

func (s *PushShrinkRequest) GetHarmonyBadgeAddNum() *int32 {
	return s.HarmonyBadgeAddNum
}

func (s *PushShrinkRequest) GetHarmonyBadgeSetNum() *int32 {
	return s.HarmonyBadgeSetNum
}

func (s *PushShrinkRequest) GetHarmonyCategory() *string {
	return s.HarmonyCategory
}

func (s *PushShrinkRequest) GetHarmonyExtParameters() *string {
	return s.HarmonyExtParameters
}

func (s *PushShrinkRequest) GetHarmonyExtensionExtraData() *string {
	return s.HarmonyExtensionExtraData
}

func (s *PushShrinkRequest) GetHarmonyExtensionPush() *bool {
	return s.HarmonyExtensionPush
}

func (s *PushShrinkRequest) GetHarmonyImageUrl() *string {
	return s.HarmonyImageUrl
}

func (s *PushShrinkRequest) GetHarmonyInboxContent() *string {
	return s.HarmonyInboxContent
}

func (s *PushShrinkRequest) GetHarmonyLiveViewPayload() *string {
	return s.HarmonyLiveViewPayload
}

func (s *PushShrinkRequest) GetHarmonyNotificationSlotType() *string {
	return s.HarmonyNotificationSlotType
}

func (s *PushShrinkRequest) GetHarmonyNotifyId() *int32 {
	return s.HarmonyNotifyId
}

func (s *PushShrinkRequest) GetHarmonyReceiptId() *string {
	return s.HarmonyReceiptId
}

func (s *PushShrinkRequest) GetHarmonyRemind() *bool {
	return s.HarmonyRemind
}

func (s *PushShrinkRequest) GetHarmonyRemindBody() *string {
	return s.HarmonyRemindBody
}

func (s *PushShrinkRequest) GetHarmonyRemindTitle() *string {
	return s.HarmonyRemindTitle
}

func (s *PushShrinkRequest) GetHarmonyRenderStyle() *string {
	return s.HarmonyRenderStyle
}

func (s *PushShrinkRequest) GetHarmonyTestMessage() *bool {
	return s.HarmonyTestMessage
}

func (s *PushShrinkRequest) GetHarmonyUri() *string {
	return s.HarmonyUri
}

func (s *PushShrinkRequest) GetIdempotentToken() *string {
	return s.IdempotentToken
}

func (s *PushShrinkRequest) GetJobKey() *string {
	return s.JobKey
}

func (s *PushShrinkRequest) GetPushTime() *string {
	return s.PushTime
}

func (s *PushShrinkRequest) GetPushType() *string {
	return s.PushType
}

func (s *PushShrinkRequest) GetSendChannels() *string {
	return s.SendChannels
}

func (s *PushShrinkRequest) GetSendSpeed() *int32 {
	return s.SendSpeed
}

func (s *PushShrinkRequest) GetSmsDelaySecs() *int32 {
	return s.SmsDelaySecs
}

func (s *PushShrinkRequest) GetSmsParams() *string {
	return s.SmsParams
}

func (s *PushShrinkRequest) GetSmsSendPolicy() *int32 {
	return s.SmsSendPolicy
}

func (s *PushShrinkRequest) GetSmsSignName() *string {
	return s.SmsSignName
}

func (s *PushShrinkRequest) GetSmsTemplateName() *string {
	return s.SmsTemplateName
}

func (s *PushShrinkRequest) GetStoreOffline() *bool {
	return s.StoreOffline
}

func (s *PushShrinkRequest) GetTarget() *string {
	return s.Target
}

func (s *PushShrinkRequest) GetTargetValue() *string {
	return s.TargetValue
}

func (s *PushShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *PushShrinkRequest) GetTrim() *bool {
	return s.Trim
}

func (s *PushShrinkRequest) GetIOSApnsEnv() *string {
	return s.IOSApnsEnv
}

func (s *PushShrinkRequest) GetIOSBadge() *int32 {
	return s.IOSBadge
}

func (s *PushShrinkRequest) GetIOSBadgeAutoIncrement() *bool {
	return s.IOSBadgeAutoIncrement
}

func (s *PushShrinkRequest) GetIOSExtParameters() *string {
	return s.IOSExtParameters
}

func (s *PushShrinkRequest) GetIOSInterruptionLevel() *string {
	return s.IOSInterruptionLevel
}

func (s *PushShrinkRequest) GetIOSLiveActivityAttributes() *string {
	return s.IOSLiveActivityAttributes
}

func (s *PushShrinkRequest) GetIOSLiveActivityAttributesType() *string {
	return s.IOSLiveActivityAttributesType
}

func (s *PushShrinkRequest) GetIOSLiveActivityContentState() *string {
	return s.IOSLiveActivityContentState
}

func (s *PushShrinkRequest) GetIOSLiveActivityDismissalDate() *int64 {
	return s.IOSLiveActivityDismissalDate
}

func (s *PushShrinkRequest) GetIOSLiveActivityEvent() *string {
	return s.IOSLiveActivityEvent
}

func (s *PushShrinkRequest) GetIOSLiveActivityId() *string {
	return s.IOSLiveActivityId
}

func (s *PushShrinkRequest) GetIOSLiveActivityStaleDate() *int64 {
	return s.IOSLiveActivityStaleDate
}

func (s *PushShrinkRequest) GetIOSMusic() *string {
	return s.IOSMusic
}

func (s *PushShrinkRequest) GetIOSMutableContent() *bool {
	return s.IOSMutableContent
}

func (s *PushShrinkRequest) GetIOSNotificationCategory() *string {
	return s.IOSNotificationCategory
}

func (s *PushShrinkRequest) GetIOSNotificationCollapseId() *string {
	return s.IOSNotificationCollapseId
}

func (s *PushShrinkRequest) GetIOSNotificationThreadId() *string {
	return s.IOSNotificationThreadId
}

func (s *PushShrinkRequest) GetIOSRelevanceScore() *float64 {
	return s.IOSRelevanceScore
}

func (s *PushShrinkRequest) GetIOSRemind() *bool {
	return s.IOSRemind
}

func (s *PushShrinkRequest) GetIOSRemindBody() *string {
	return s.IOSRemindBody
}

func (s *PushShrinkRequest) GetIOSSilentNotification() *bool {
	return s.IOSSilentNotification
}

func (s *PushShrinkRequest) GetIOSSubtitle() *string {
	return s.IOSSubtitle
}

func (s *PushShrinkRequest) SetAndroidActivity(v string) *PushShrinkRequest {
	s.AndroidActivity = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidBadgeAddNum(v int32) *PushShrinkRequest {
	s.AndroidBadgeAddNum = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidBadgeClass(v string) *PushShrinkRequest {
	s.AndroidBadgeClass = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidBadgeSetNum(v int32) *PushShrinkRequest {
	s.AndroidBadgeSetNum = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidBigBody(v string) *PushShrinkRequest {
	s.AndroidBigBody = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidBigPictureUrl(v string) *PushShrinkRequest {
	s.AndroidBigPictureUrl = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidBigTitle(v string) *PushShrinkRequest {
	s.AndroidBigTitle = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidExtParameters(v string) *PushShrinkRequest {
	s.AndroidExtParameters = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidHonorTargetUserType(v int32) *PushShrinkRequest {
	s.AndroidHonorTargetUserType = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidHuaweiBusinessType(v int32) *PushShrinkRequest {
	s.AndroidHuaweiBusinessType = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidHuaweiLiveNotificationPayload(v string) *PushShrinkRequest {
	s.AndroidHuaweiLiveNotificationPayload = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidHuaweiReceiptId(v string) *PushShrinkRequest {
	s.AndroidHuaweiReceiptId = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidHuaweiTargetUserType(v int32) *PushShrinkRequest {
	s.AndroidHuaweiTargetUserType = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidImageUrl(v string) *PushShrinkRequest {
	s.AndroidImageUrl = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidInboxBody(v string) *PushShrinkRequest {
	s.AndroidInboxBody = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidMeizuNoticeMsgType(v int32) *PushShrinkRequest {
	s.AndroidMeizuNoticeMsgType = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidMessageHuaweiCategory(v string) *PushShrinkRequest {
	s.AndroidMessageHuaweiCategory = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidMessageHuaweiUrgency(v string) *PushShrinkRequest {
	s.AndroidMessageHuaweiUrgency = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidMessageOppoCategory(v string) *PushShrinkRequest {
	s.AndroidMessageOppoCategory = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidMessageOppoNotifyLevel(v int32) *PushShrinkRequest {
	s.AndroidMessageOppoNotifyLevel = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidMessageVivoCategory(v string) *PushShrinkRequest {
	s.AndroidMessageVivoCategory = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidMusic(v string) *PushShrinkRequest {
	s.AndroidMusic = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidNotificationBarPriority(v int32) *PushShrinkRequest {
	s.AndroidNotificationBarPriority = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidNotificationBarType(v int32) *PushShrinkRequest {
	s.AndroidNotificationBarType = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidNotificationChannel(v string) *PushShrinkRequest {
	s.AndroidNotificationChannel = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidNotificationGroup(v string) *PushShrinkRequest {
	s.AndroidNotificationGroup = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidNotificationHonorChannel(v string) *PushShrinkRequest {
	s.AndroidNotificationHonorChannel = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidNotificationHuaweiChannel(v string) *PushShrinkRequest {
	s.AndroidNotificationHuaweiChannel = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidNotificationNotifyId(v int32) *PushShrinkRequest {
	s.AndroidNotificationNotifyId = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidNotificationThreadId(v string) *PushShrinkRequest {
	s.AndroidNotificationThreadId = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidNotificationVivoChannel(v string) *PushShrinkRequest {
	s.AndroidNotificationVivoChannel = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidNotificationXiaomiChannel(v string) *PushShrinkRequest {
	s.AndroidNotificationXiaomiChannel = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidNotifyType(v string) *PushShrinkRequest {
	s.AndroidNotifyType = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidOpenType(v string) *PushShrinkRequest {
	s.AndroidOpenType = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidOpenUrl(v string) *PushShrinkRequest {
	s.AndroidOpenUrl = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidOppoDeleteIntentData(v string) *PushShrinkRequest {
	s.AndroidOppoDeleteIntentData = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidOppoIntelligentIntent(v string) *PushShrinkRequest {
	s.AndroidOppoIntelligentIntent = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidOppoIntentEnv(v int32) *PushShrinkRequest {
	s.AndroidOppoIntentEnv = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidOppoPrivateContentParametersShrink(v string) *PushShrinkRequest {
	s.AndroidOppoPrivateContentParametersShrink = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidOppoPrivateMsgTemplateId(v string) *PushShrinkRequest {
	s.AndroidOppoPrivateMsgTemplateId = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidOppoPrivateTitleParametersShrink(v string) *PushShrinkRequest {
	s.AndroidOppoPrivateTitleParametersShrink = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidPopupActivity(v string) *PushShrinkRequest {
	s.AndroidPopupActivity = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidPopupBody(v string) *PushShrinkRequest {
	s.AndroidPopupBody = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidPopupTitle(v string) *PushShrinkRequest {
	s.AndroidPopupTitle = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidRemind(v bool) *PushShrinkRequest {
	s.AndroidRemind = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidRenderStyle(v int32) *PushShrinkRequest {
	s.AndroidRenderStyle = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidTargetUserType(v int32) *PushShrinkRequest {
	s.AndroidTargetUserType = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidVivoLiveMessage(v string) *PushShrinkRequest {
	s.AndroidVivoLiveMessage = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidVivoPushMode(v int32) *PushShrinkRequest {
	s.AndroidVivoPushMode = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidVivoReceiptId(v string) *PushShrinkRequest {
	s.AndroidVivoReceiptId = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidXiaoMiActivity(v string) *PushShrinkRequest {
	s.AndroidXiaoMiActivity = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidXiaoMiNotifyBody(v string) *PushShrinkRequest {
	s.AndroidXiaoMiNotifyBody = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidXiaoMiNotifyTitle(v string) *PushShrinkRequest {
	s.AndroidXiaoMiNotifyTitle = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidXiaomiBigPictureUrl(v string) *PushShrinkRequest {
	s.AndroidXiaomiBigPictureUrl = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidXiaomiFocusParam(v string) *PushShrinkRequest {
	s.AndroidXiaomiFocusParam = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidXiaomiFocusPics(v string) *PushShrinkRequest {
	s.AndroidXiaomiFocusPics = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidXiaomiImageUrl(v string) *PushShrinkRequest {
	s.AndroidXiaomiImageUrl = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidXiaomiTemplateId(v string) *PushShrinkRequest {
	s.AndroidXiaomiTemplateId = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidXiaomiTemplateParams(v string) *PushShrinkRequest {
	s.AndroidXiaomiTemplateParams = &v
	return s
}

func (s *PushShrinkRequest) SetAppKey(v int64) *PushShrinkRequest {
	s.AppKey = &v
	return s
}

func (s *PushShrinkRequest) SetBody(v string) *PushShrinkRequest {
	s.Body = &v
	return s
}

func (s *PushShrinkRequest) SetDeviceType(v string) *PushShrinkRequest {
	s.DeviceType = &v
	return s
}

func (s *PushShrinkRequest) SetExpireTime(v string) *PushShrinkRequest {
	s.ExpireTime = &v
	return s
}

func (s *PushShrinkRequest) SetHarmonyAction(v string) *PushShrinkRequest {
	s.HarmonyAction = &v
	return s
}

func (s *PushShrinkRequest) SetHarmonyActionType(v string) *PushShrinkRequest {
	s.HarmonyActionType = &v
	return s
}

func (s *PushShrinkRequest) SetHarmonyBadgeAddNum(v int32) *PushShrinkRequest {
	s.HarmonyBadgeAddNum = &v
	return s
}

func (s *PushShrinkRequest) SetHarmonyBadgeSetNum(v int32) *PushShrinkRequest {
	s.HarmonyBadgeSetNum = &v
	return s
}

func (s *PushShrinkRequest) SetHarmonyCategory(v string) *PushShrinkRequest {
	s.HarmonyCategory = &v
	return s
}

func (s *PushShrinkRequest) SetHarmonyExtParameters(v string) *PushShrinkRequest {
	s.HarmonyExtParameters = &v
	return s
}

func (s *PushShrinkRequest) SetHarmonyExtensionExtraData(v string) *PushShrinkRequest {
	s.HarmonyExtensionExtraData = &v
	return s
}

func (s *PushShrinkRequest) SetHarmonyExtensionPush(v bool) *PushShrinkRequest {
	s.HarmonyExtensionPush = &v
	return s
}

func (s *PushShrinkRequest) SetHarmonyImageUrl(v string) *PushShrinkRequest {
	s.HarmonyImageUrl = &v
	return s
}

func (s *PushShrinkRequest) SetHarmonyInboxContent(v string) *PushShrinkRequest {
	s.HarmonyInboxContent = &v
	return s
}

func (s *PushShrinkRequest) SetHarmonyLiveViewPayload(v string) *PushShrinkRequest {
	s.HarmonyLiveViewPayload = &v
	return s
}

func (s *PushShrinkRequest) SetHarmonyNotificationSlotType(v string) *PushShrinkRequest {
	s.HarmonyNotificationSlotType = &v
	return s
}

func (s *PushShrinkRequest) SetHarmonyNotifyId(v int32) *PushShrinkRequest {
	s.HarmonyNotifyId = &v
	return s
}

func (s *PushShrinkRequest) SetHarmonyReceiptId(v string) *PushShrinkRequest {
	s.HarmonyReceiptId = &v
	return s
}

func (s *PushShrinkRequest) SetHarmonyRemind(v bool) *PushShrinkRequest {
	s.HarmonyRemind = &v
	return s
}

func (s *PushShrinkRequest) SetHarmonyRemindBody(v string) *PushShrinkRequest {
	s.HarmonyRemindBody = &v
	return s
}

func (s *PushShrinkRequest) SetHarmonyRemindTitle(v string) *PushShrinkRequest {
	s.HarmonyRemindTitle = &v
	return s
}

func (s *PushShrinkRequest) SetHarmonyRenderStyle(v string) *PushShrinkRequest {
	s.HarmonyRenderStyle = &v
	return s
}

func (s *PushShrinkRequest) SetHarmonyTestMessage(v bool) *PushShrinkRequest {
	s.HarmonyTestMessage = &v
	return s
}

func (s *PushShrinkRequest) SetHarmonyUri(v string) *PushShrinkRequest {
	s.HarmonyUri = &v
	return s
}

func (s *PushShrinkRequest) SetIdempotentToken(v string) *PushShrinkRequest {
	s.IdempotentToken = &v
	return s
}

func (s *PushShrinkRequest) SetJobKey(v string) *PushShrinkRequest {
	s.JobKey = &v
	return s
}

func (s *PushShrinkRequest) SetPushTime(v string) *PushShrinkRequest {
	s.PushTime = &v
	return s
}

func (s *PushShrinkRequest) SetPushType(v string) *PushShrinkRequest {
	s.PushType = &v
	return s
}

func (s *PushShrinkRequest) SetSendChannels(v string) *PushShrinkRequest {
	s.SendChannels = &v
	return s
}

func (s *PushShrinkRequest) SetSendSpeed(v int32) *PushShrinkRequest {
	s.SendSpeed = &v
	return s
}

func (s *PushShrinkRequest) SetSmsDelaySecs(v int32) *PushShrinkRequest {
	s.SmsDelaySecs = &v
	return s
}

func (s *PushShrinkRequest) SetSmsParams(v string) *PushShrinkRequest {
	s.SmsParams = &v
	return s
}

func (s *PushShrinkRequest) SetSmsSendPolicy(v int32) *PushShrinkRequest {
	s.SmsSendPolicy = &v
	return s
}

func (s *PushShrinkRequest) SetSmsSignName(v string) *PushShrinkRequest {
	s.SmsSignName = &v
	return s
}

func (s *PushShrinkRequest) SetSmsTemplateName(v string) *PushShrinkRequest {
	s.SmsTemplateName = &v
	return s
}

func (s *PushShrinkRequest) SetStoreOffline(v bool) *PushShrinkRequest {
	s.StoreOffline = &v
	return s
}

func (s *PushShrinkRequest) SetTarget(v string) *PushShrinkRequest {
	s.Target = &v
	return s
}

func (s *PushShrinkRequest) SetTargetValue(v string) *PushShrinkRequest {
	s.TargetValue = &v
	return s
}

func (s *PushShrinkRequest) SetTitle(v string) *PushShrinkRequest {
	s.Title = &v
	return s
}

func (s *PushShrinkRequest) SetTrim(v bool) *PushShrinkRequest {
	s.Trim = &v
	return s
}

func (s *PushShrinkRequest) SetIOSApnsEnv(v string) *PushShrinkRequest {
	s.IOSApnsEnv = &v
	return s
}

func (s *PushShrinkRequest) SetIOSBadge(v int32) *PushShrinkRequest {
	s.IOSBadge = &v
	return s
}

func (s *PushShrinkRequest) SetIOSBadgeAutoIncrement(v bool) *PushShrinkRequest {
	s.IOSBadgeAutoIncrement = &v
	return s
}

func (s *PushShrinkRequest) SetIOSExtParameters(v string) *PushShrinkRequest {
	s.IOSExtParameters = &v
	return s
}

func (s *PushShrinkRequest) SetIOSInterruptionLevel(v string) *PushShrinkRequest {
	s.IOSInterruptionLevel = &v
	return s
}

func (s *PushShrinkRequest) SetIOSLiveActivityAttributes(v string) *PushShrinkRequest {
	s.IOSLiveActivityAttributes = &v
	return s
}

func (s *PushShrinkRequest) SetIOSLiveActivityAttributesType(v string) *PushShrinkRequest {
	s.IOSLiveActivityAttributesType = &v
	return s
}

func (s *PushShrinkRequest) SetIOSLiveActivityContentState(v string) *PushShrinkRequest {
	s.IOSLiveActivityContentState = &v
	return s
}

func (s *PushShrinkRequest) SetIOSLiveActivityDismissalDate(v int64) *PushShrinkRequest {
	s.IOSLiveActivityDismissalDate = &v
	return s
}

func (s *PushShrinkRequest) SetIOSLiveActivityEvent(v string) *PushShrinkRequest {
	s.IOSLiveActivityEvent = &v
	return s
}

func (s *PushShrinkRequest) SetIOSLiveActivityId(v string) *PushShrinkRequest {
	s.IOSLiveActivityId = &v
	return s
}

func (s *PushShrinkRequest) SetIOSLiveActivityStaleDate(v int64) *PushShrinkRequest {
	s.IOSLiveActivityStaleDate = &v
	return s
}

func (s *PushShrinkRequest) SetIOSMusic(v string) *PushShrinkRequest {
	s.IOSMusic = &v
	return s
}

func (s *PushShrinkRequest) SetIOSMutableContent(v bool) *PushShrinkRequest {
	s.IOSMutableContent = &v
	return s
}

func (s *PushShrinkRequest) SetIOSNotificationCategory(v string) *PushShrinkRequest {
	s.IOSNotificationCategory = &v
	return s
}

func (s *PushShrinkRequest) SetIOSNotificationCollapseId(v string) *PushShrinkRequest {
	s.IOSNotificationCollapseId = &v
	return s
}

func (s *PushShrinkRequest) SetIOSNotificationThreadId(v string) *PushShrinkRequest {
	s.IOSNotificationThreadId = &v
	return s
}

func (s *PushShrinkRequest) SetIOSRelevanceScore(v float64) *PushShrinkRequest {
	s.IOSRelevanceScore = &v
	return s
}

func (s *PushShrinkRequest) SetIOSRemind(v bool) *PushShrinkRequest {
	s.IOSRemind = &v
	return s
}

func (s *PushShrinkRequest) SetIOSRemindBody(v string) *PushShrinkRequest {
	s.IOSRemindBody = &v
	return s
}

func (s *PushShrinkRequest) SetIOSSilentNotification(v bool) *PushShrinkRequest {
	s.IOSSilentNotification = &v
	return s
}

func (s *PushShrinkRequest) SetIOSSubtitle(v string) *PushShrinkRequest {
	s.IOSSubtitle = &v
	return s
}

func (s *PushShrinkRequest) Validate() error {
	return dara.Validate(s)
}
