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
	// Specify the activity to open from the notification.
	//
	// Only pass this when AndroidOpenType="Activity", e.g.: `com.alibaba.cloudpushdemo.bizactivity`.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidActivity *string `json:"AndroidActivity,omitempty" xml:"AndroidActivity,omitempty"`
	// Set the badge increment value, which is added to the current badge count. Value range: [1-99].
	//
	// > Only effective for Huawei/Honor vendor channel push. When both AndroidBadgeAddNum and AndroidBadgeSetNum are present, AndroidBadgeSetNum takes precedence.
	//
	// example:
	//
	// 1
	AndroidBadgeAddNum *int32 `json:"AndroidBadgeAddNum,omitempty" xml:"AndroidBadgeAddNum,omitempty"`
	// Full class name of the app entry Activity for badge settings.
	//
	// > Only effective for Huawei/Honor vendor channel push.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidBadgeClass *string `json:"AndroidBadgeClass,omitempty" xml:"AndroidBadgeClass,omitempty"`
	// Set a fixed badge number. Value range: [0-99].
	//
	// > For vendor channel push, only effective on Huawei and Honor channels. For Alibaba Cloud proprietary channel push, only effective on Huawei, Honor, and vivo devices.
	//
	// example:
	//
	// 5
	AndroidBadgeSetNum *int32 `json:"AndroidBadgeSetNum,omitempty" xml:"AndroidBadgeSetNum,omitempty"`
	// Body in long text mode. Length limit: 1000 bytes (1 Chinese character counts as 3 bytes). Subject to specific vendor channel limits when sending.
	//
	// Currently supported by:
	//
	// - Huawei: EMUI 10 and above
	//
	// - Honor: Magic UI 4.0 and above
	//
	// - Xiaomi: MIUI 10 and above
	//
	// - OPPO: ColorOS 5.0 and above
	//
	// - Meizu: Flyme
	//
	// - Proprietary channel: Android SDK 3.6.0 and above
	//
	// >If not provided in long text mode, the first non-empty value from Body or AndroidPopupBody is used.
	//
	// example:
	//
	// 示例长文本
	AndroidBigBody *string `json:"AndroidBigBody,omitempty" xml:"AndroidBigBody,omitempty"`
	// Image URL in big picture mode. Currently supported by: Proprietary channel: Android SDK 3.6.0 and above.
	//
	// example:
	//
	// https://imag.example.com/image.png
	AndroidBigPictureUrl *string `json:"AndroidBigPictureUrl,omitempty" xml:"AndroidBigPictureUrl,omitempty"`
	// Title in long text mode. Length limit: 200 bytes (1 Chinese character counts as 3 bytes).
	//
	// - Currently only supported by the Honor channel and Huawei channel EMUI 11 and above.
	//
	// - If not provided in long text mode, the first non-empty value from Title or AndroidPopupTitle is used.
	//
	// example:
	//
	// 示例长标题
	AndroidBigTitle *string `json:"AndroidBigTitle,omitempty" xml:"AndroidBigTitle,omitempty"`
	// Set the extension attributes of the notification. This attribute does not take effect when PushType is set to MESSAGE.
	//
	// This parameter must be passed in JSON map format, otherwise parsing will fail.
	//
	// example:
	//
	// {"key1":"value1","api_name":"PushNoticeToAndroidRequest"}
	AndroidExtParameters *string `json:"AndroidExtParameters,omitempty" xml:"AndroidExtParameters,omitempty"`
	// Set Honor channel notification type:
	//
	// - **0**: Official notification (default).
	//
	// - **1**: Test notification.
	//
	// > Each application can send up to 1000 test notifications per day, and these are not subject to the daily per-device push limit.
	//
	// example:
	//
	// 0
	AndroidHonorTargetUserType *int32 `json:"AndroidHonorTargetUserType,omitempty" xml:"AndroidHonorTargetUserType,omitempty"`
	// Set Huawei Quick Notification parameter:
	//
	// - **0**: Send Huawei standard notification (default).
	//
	// - **1**: Send Huawei Quick Notification.
	//
	// example:
	//
	// 1
	AndroidHuaweiBusinessType *int32 `json:"AndroidHuaweiBusinessType,omitempty" xml:"AndroidHuaweiBusinessType,omitempty"`
	// JSON string of the Huawei Android Live Notification data structure [LiveNotificationPayload](https://developer.huawei.com/consumer/cn/doc/HMSCore-References/rest-live-0000001562939968#ZH-CN_TOPIC_0000001700850537__p195121620102511). For development integration, refer to the documentation [Huawei Live Notification Push Guide](https://help.aliyun.com/document_detail/2983768.html).
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
	// Huawei channel receipt ID. This receipt ID can be found in the receipt parameter configuration on the Huawei channel push management platform.
	//
	// > If the default receipt configuration on the Huawei channel push management platform is set to the Alibaba Cloud receipt, this is not required. If not, it is recommended to configure the Huawei channel default receipt ID in the Alibaba Cloud EMAS Mobile Push console first.
	//
	// example:
	//
	// RCP4C123456
	AndroidHuaweiReceiptId *string `json:"AndroidHuaweiReceiptId,omitempty" xml:"AndroidHuaweiReceiptId,omitempty"`
	// Set Huawei channel notification type:
	//
	// - **0**: Official notification (default).
	//
	// - **1**: Test notification.
	//
	// > Each application can send up to 500 test notifications per day, and these are not subject to the daily per-device push limit.
	//
	// example:
	//
	// 0
	AndroidHuaweiTargetUserType *int32 `json:"AndroidHuaweiTargetUserType,omitempty" xml:"AndroidHuaweiTargetUserType,omitempty"`
	// Right-side icon URL.
	//
	// Currently supported by:
	//
	// - Huawei EMUI (only applicable in long text mode and Inbox mode).
	//
	// - Honor Magic UI (only applicable in long text mode).
	//
	// - Proprietary channel: Android SDK 3.5.0 and above.
	//
	// example:
	//
	// https://imag.example.com/image.png
	AndroidImageUrl *string `json:"AndroidImageUrl,omitempty" xml:"AndroidImageUrl,omitempty"`
	// Body content in Inbox mode. The content must be a valid JSON Array with no more than 5 elements. Currently supported by:
	//
	// - Huawei: EMUI 9 and above
	//
	// - Honor: Magic UI 4.0 and above
	//
	// - Xiaomi: MIUI 10 and above
	//
	// - OPPO: ColorOS 5.0 and above
	//
	// - Proprietary channel: Android SDK 3.6.0 and above
	//
	// example:
	//
	// ["第一行","第二行"]
	AndroidInboxBody *string `json:"AndroidInboxBody,omitempty" xml:"AndroidInboxBody,omitempty"`
	// Meizu message type:
	//
	// - 0: Public message (default)
	//
	// - 1: Private message
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 0
	AndroidMeizuNoticeMsgType *int32 `json:"AndroidMeizuNoticeMsgType,omitempty" xml:"AndroidMeizuNoticeMsgType,omitempty"`
	// Purpose 1: After completing the [self-classification rights application](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835?#section3410731125514), this is used to identify the message type, determine the [message notification method](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#ZH-CN_TOPIC_0000001149358835__p3850133955718), and accelerate delivery for specific message types. For valid values, refer to the [Message Classification Standard](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section1076611477914) in Huawei\\"s official push documentation, using the "Cloud notification category value" or "Local notification category value" from the table.
	//
	// Purpose 2: After [applying for special permissions](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/faq-0000001050042183#section037425218509), this is used to identify high-priority transparent transmission scenarios. Valid values:
	//
	// - VOIP: Audio/video calls
	//
	// - PLAY_VOICE: Voice playback
	//
	// > For items where "Cloud notification category value" is "Not applicable", they are delivered through the Alibaba Cloud proprietary channel. For items where "Local notification category value" is "Not applicable", they are delivered through the Huawei channel.
	//
	// example:
	//
	// VOIP
	AndroidMessageHuaweiCategory *string `json:"AndroidMessageHuaweiCategory,omitempty" xml:"AndroidMessageHuaweiCategory,omitempty"`
	// Huawei channel notification delivery priority. Valid values:
	//
	// - **HIGH**
	//
	// - **NORMAL**
	//
	// Requires permission application. For details, see: [Application Link](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/faq-0000001050042183#section037425218509).
	//
	// example:
	//
	// HIGH
	AndroidMessageHuaweiUrgency *string `json:"AndroidMessageHuaweiUrgency,omitempty" xml:"AndroidMessageHuaweiUrgency,omitempty"`
	// OPPO classifies messages into two categories for management: Communication & Service, and Content & Marketing.
	//
	// Communication & Service (requires permission application):
	//
	// - IM: Instant messaging, audio, video calls
	//
	// - ACCOUNT: Personal account and asset changes
	//
	// - DEVICE_REMINDER: Personal device reminders
	//
	// - ORDER: Personal order/logistics status changes
	//
	// - TODO: Personal schedules/to-dos
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
	// For details, refer to [OPUSH Message Classification Rules](https://open.oppomobile.com/new/developmentDoc/info?id=13189).
	//
	// example:
	//
	// MARKETING
	AndroidMessageOppoCategory *string `json:"AndroidMessageOppoCategory,omitempty" xml:"AndroidMessageOppoCategory,omitempty"`
	// OPPO channel notification bar message notification level. Valid values:
	//
	// - 1: Notification bar
	//
	// - 2: Notification bar, lock screen, ringtone, vibration (default notification level for Communication & Service messages)
	//
	// - 16: Notification bar, lock screen, ringtone, vibration, banner (requires permission application)
	//
	// > When using the AndroidMessageOppoNotifyLevel parameter, the AndroidMessageOppoCategory parameter must also be provided.
	//
	// example:
	//
	// 1
	AndroidMessageOppoNotifyLevel *int32 `json:"AndroidMessageOppoNotifyLevel,omitempty" xml:"AndroidMessageOppoNotifyLevel,omitempty"`
	// vivo classifies messages into two categories for management: System messages and Operational messages.
	//
	// System messages:
	//
	// - IM: Instant messages
	//
	// - ACCOUNT: Accounts and assets
	//
	// - TODO: Schedules and to-dos
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
	// - MARKETING: Operational promotions
	//
	// - SOCIAL: Social updates
	//
	// For details, refer to [Classification Description](https://dev.vivo.com.cn/documentCenter/doc/359#s-ef3qugc3).
	//
	// example:
	//
	// TODO
	AndroidMessageVivoCategory *string `json:"AndroidMessageVivoCategory,omitempty" xml:"AndroidMessageVivoCategory,omitempty"`
	// Huawei vendor channel notification sound. Specify the name of an audio file stored in the client project\\"s app/src/main/res/raw/ directory, without the file extension.
	//
	// If not set, the default ringtone is used.
	//
	// example:
	//
	// alicloud_notification_sound
	AndroidMusic *string `json:"AndroidMusic,omitempty" xml:"AndroidMusic,omitempty"`
	// Priority of the Android notification position in the notification bar. Valid values: -2, -1, 0, 1, 2.
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
	// The channelId of the Android app, which must correspond to the channelId in the app.
	//
	// - Set the NotificationChannel parameter. For specific usage, see [FAQ: Notifications Not Received on Android 8.0+ Devices](https://help.aliyun.com/document_detail/67398.html).
	//
	// - Since the OPPO private message channel\\"s channel_id is the same as the app\\"s channelId, the channel_id for OPPO channel push takes this value.
	//
	// - For Huawei, FCM, and Alibaba Cloud proprietary channel push, the channel_id takes this value.
	//
	// example:
	//
	// 1
	AndroidNotificationChannel *string `json:"AndroidNotificationChannel,omitempty" xml:"AndroidNotificationChannel,omitempty"`
	// Message grouping. Messages in the same group are displayed as only the latest one in the notification bar along with the total count of messages received for that group. All messages are not shown and cannot be expanded. Currently supported by:
	//
	// - Huawei vendor channel
	//
	// - Honor vendor channel
	//
	// - Proprietary channel: Android SDK 3.9.1 and below
	//
	// > The proprietary channel no longer supports this parameter on Android SDK 3.9.2 and above.
	//
	// example:
	//
	// group-1
	AndroidNotificationGroup *string `json:"AndroidNotificationGroup,omitempty" xml:"AndroidNotificationGroup,omitempty"`
	// Set the Honor notification message classification importance parameter, which determines notification behavior on user devices. Valid values:
	//
	// - **LOW**: Information and marketing messages
	//
	// - **NORMAL**: Service and communication messages
	//
	// Requires application on the Honor platform. [Application Link](https://developer.honor.com/cn/docs/11002/guides/notification-class#%E8%87%AA%E5%88%86%E7%B1%BB%E6%9D%83%E7%9B%8A%E7%94%B3%E8%AF%B7).
	//
	// example:
	//
	// LOW
	AndroidNotificationHonorChannel *string `json:"AndroidNotificationHonorChannel,omitempty" xml:"AndroidNotificationHonorChannel,omitempty"`
	// Set the Huawei notification message classification importance parameter, which determines notification behavior on user devices. Valid values:
	//
	// - LOW: Information and marketing messages
	//
	// - NORMAL: Service and communication messages
	//
	// >- Huawei channel currently recommends using AndroidMessageHuaweiCategory for notification classification. AndroidNotificationHuaweiChannel is no longer required.
	//
	// >- Requires application on the Huawei platform. [Application Link](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section893184112272).
	//
	// example:
	//
	// LOW
	AndroidNotificationHuaweiChannel *string `json:"AndroidNotificationHuaweiChannel,omitempty" xml:"AndroidNotificationHuaweiChannel,omitempty"`
	// Unique identifier for each message when displayed in the notification bar. Different notification bar messages can share the same NotifyId, allowing new notifications to replace old ones.
	//
	// example:
	//
	// 100001
	AndroidNotificationNotifyId *int32 `json:"AndroidNotificationNotifyId,omitempty" xml:"AndroidNotificationNotifyId,omitempty"`
	// Message grouping. Messages in the same group are collapsed in the notification bar and can be expanded. Notifications from different groups are displayed separately. Currently supported by:
	//
	// - Proprietary channel: Android SDK 3.9.2 and above
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
	// >- vivo channel currently recommends using AndroidMessageVivoCategory for notification classification. AndroidNotificationVivoChannel is no longer required.
	//
	// >- Requires application on the vivo platform. For details, see: [Application Link](https://dev.vivo.com.cn/documentCenter/doc/359).
	//
	// example:
	//
	// classification
	AndroidNotificationVivoChannel *string `json:"AndroidNotificationVivoChannel,omitempty" xml:"AndroidNotificationVivoChannel,omitempty"`
	// Set the Xiaomi notification type channelId. Requires application on the Xiaomi platform. For details, see: [Application Link](https://dev.mi.com/console/doc/detail?pId=2422#_4).
	//
	// >- A single application can apply for a maximum of 8 channels on the Xiaomi channel. Please plan ahead.
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
	// BOTH
	AndroidNotifyType *string `json:"AndroidNotifyType,omitempty" xml:"AndroidNotifyType,omitempty"`
	// Action after clicking the notification. Valid values:
	//
	// - **APPLICATION**: Open the application (default)
	//
	// - **ACTIVITY**: Open an Android Activity
	//
	// - **URL**: Open a URL
	//
	// - **NONE**: No redirect
	//
	// example:
	//
	// APPLICATION
	AndroidOpenType *string `json:"AndroidOpenType,omitempty" xml:"AndroidOpenType,omitempty"`
	// URL to open when Android receives the push.
	//
	// Only pass this when AndroidOpenType="URL".
	//
	// example:
	//
	// https://xxxx.xxx
	AndroidOpenUrl *string `json:"AndroidOpenUrl,omitempty" xml:"AndroidOpenUrl,omitempty"`
	// JSON string of the OPPO Fluid Cloud intent deletion data structure [data](https://open.oppomobile.com/documentation/page/info?id=13578). When the AndroidOppoIntelligentIntent parameter is already provided, this parameter is ignored. For development integration, refer to the documentation [OPPO Fluid Cloud Push Guide](https://help.aliyun.com/document_detail/2997310.html).
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
	// JSON string of the OPPO Fluid Cloud intent sharing data structure [IntelligentIntent](https://open.oppomobile.com/documentation/page/info?id=13565). For development integration, refer to the documentation [OPPO Fluid Cloud Push Guide](https://help.aliyun.com/document_detail/2997310.html).
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
	// Set OPPO Fluid Cloud push environment:
	//
	// - **0**: Production environment (default).
	//
	// - **1**: Test environment.
	//
	// > OPPO Fluid Cloud test environment requires setting up the client environment as described in [Environment Setup](https://open.oppomobile.com/documentation/page/info?id=13590).
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
	// Specify the Activity to navigate to after clicking the notification.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidPopupActivity *string `json:"AndroidPopupActivity,omitempty" xml:"AndroidPopupActivity,omitempty"`
	// Body content in supplementary popup mode. Required when the **AndroidPopupActivity*	- parameter is not empty.
	//
	// Length limit: 200 characters (both Chinese and English characters count as one character).
	//
	// If using vendor channels, it must also comply with vendor channel limits. For details, see: [Android Supplementary Channel Push Limits](https://help.aliyun.com/document_detail/165253.html).
	//
	// example:
	//
	// hello
	AndroidPopupBody *string `json:"AndroidPopupBody,omitempty" xml:"AndroidPopupBody,omitempty"`
	// Title content in supplementary popup mode. Required when the **AndroidPopupActivity*	- parameter is not empty.
	//
	// Length limit: 50 characters (both Chinese and English characters count as one character).
	//
	// If using vendor channels, it must also comply with vendor channel limits. For details, see: [Android Supplementary Channel Push Limits](https://help.aliyun.com/document_detail/165253.html).
	//
	// example:
	//
	// hello
	AndroidPopupTitle *string `json:"AndroidPopupTitle,omitempty" xml:"AndroidPopupTitle,omitempty"`
	// When the push type is message and the device is offline, this push will use the supplementary popup feature. Default is false. Only effective when PushType=MESSAGE.
	//
	// If the message-to-notification push is successful, the notification displays the AndroidPopupTitle and AndroidPopupBody parameter values set on the server. The data obtained in the onSysNoticeOpened method of the supplementary popup when clicking the notification is the Title and Body parameter values set on the server.
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
	// - **2**: Big picture mode (supported by the proprietary channel, not supported on Xiaomi devices)
	//
	// - **3**: List mode (supported by Huawei, Honor, Xiaomi, OPPO, and proprietary channels)
	//
	// > If using a non-standard mode, this parameter must be provided.
	//
	// example:
	//
	// 1
	AndroidRenderStyle *int32 `json:"AndroidRenderStyle,omitempty" xml:"AndroidRenderStyle,omitempty"`
	// Set vendor channel notification type:
	//
	// - **0**: Official notification (default).
	//
	// - **1**: Test notification.
	//
	// >- When this parameter is configured, it is equivalent to simultaneously configuring AndroidHuaweiTargetUserType, AndroidHonorTargetUserType, AndroidVivoPushMode, and AndroidOppoIntentEnv. The specific vendor channel parameters can override this parameter.
	//
	// >- Currently supported by: Huawei channel, Honor channel, vivo channel, and OPPO Fluid Cloud.
	//
	// example:
	//
	// 0
	AndroidTargetUserType *int32 `json:"AndroidTargetUserType,omitempty" xml:"AndroidTargetUserType,omitempty"`
	// JSON string of the vivo Atomic Island data structure [liveMessage](https://dev.vivo.com.cn/documentCenter/doc/896#s-fdagzbd4). For development integration, refer to the documentation [vivo Atomic Island Push Guide](https://help.aliyun.com/zh/document_detail/3030718.html).
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
	// Set vivo channel notification type:
	//
	// - **0**: Official push (default).
	//
	// - **1**: Test push.
	//
	// > For test push, please configure the test device on the vivo console in advance. The test device RegId can be obtained by searching for "onReceiveRegId regId" in the device startup logs.
	//
	// example:
	//
	// 0
	AndroidVivoPushMode *int32 `json:"AndroidVivoPushMode,omitempty" xml:"AndroidVivoPushMode,omitempty"`
	// vivo channel receipt ID. This receipt ID can be found in the application information of the push service on the vivo open platform.
	//
	// > If the default receipt configuration on the vivo open platform is set to the Alibaba Cloud receipt, this is not required. If not, it is recommended to configure the vivo channel default receipt ID in the Alibaba Cloud EMAS Mobile Push console first.
	//
	// example:
	//
	// 123
	AndroidVivoReceiptId *string `json:"AndroidVivoReceiptId,omitempty" xml:"AndroidVivoReceiptId,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. All third-party supplementary popups are now supported by the new parameter **AndroidPopupActivity**.
	//
	// example:
	//
	// 无
	AndroidXiaoMiActivity *string `json:"AndroidXiaoMiActivity,omitempty" xml:"AndroidXiaoMiActivity,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. All third-party supplementary popups are now supported by the new parameter **AndroidPopupBody**.
	//
	// example:
	//
	// 无
	AndroidXiaoMiNotifyBody *string `json:"AndroidXiaoMiNotifyBody,omitempty" xml:"AndroidXiaoMiNotifyBody,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. All third-party supplementary popups are now supported by the new parameter **AndroidPopupTitle**.
	//
	// example:
	//
	// 无
	AndroidXiaoMiNotifyTitle *string `json:"AndroidXiaoMiNotifyTitle,omitempty" xml:"AndroidXiaoMiNotifyTitle,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Starting from August 2023, Xiaomi officially no longer supports dynamically setting small icons, right-side icons, and big pictures during push on new devices/systems.
	//
	// example:
	//
	// https://f6.market.xiaomi.com/download/MiPass/aaa/bbb.png
	AndroidXiaomiBigPictureUrl *string `json:"AndroidXiaomiBigPictureUrl,omitempty" xml:"AndroidXiaomiBigPictureUrl,omitempty"`
	// JSON string of the Xiaomi Super Island data structure [miui.focus.param](https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131). For development integration, refer to the documentation [Xiaomi Super Island Push Guide](https://help.aliyun.com/zh/document_detail/3037956.html).
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
	// JSON string of the Xiaomi Super Island image data [miui.focus.pic_xxx](https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131). For development integration, refer to the documentation [Xiaomi Super Island Push Guide](https://help.aliyun.com/zh/document_detail/3037956.html).
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
	// This parameter is deprecated. Starting from August 2023, Xiaomi officially no longer supports dynamically setting small icons, right-side icons, and big pictures during push on new devices/systems.
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
	// AppKey information.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// Notification content/message content for Android and HarmonyOS push; iOS message/notification content. The push content size is limited. See [Product Limits](https://help.aliyun.com/document_detail/434629.html).
	//
	// example:
	//
	// hello
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// Device type. Valid values:
	//
	// - **HARMONY**: HarmonyOS device
	//
	// - **iOS**: iOS device
	//
	// - **ANDROID**: Android device
	//
	// - **ALL**: When the AppKey is for a legacy dual-platform application, this represents pushing to both Android and iOS devices simultaneously; when the AppKey is for a new single-platform application, the effect is the same as specifying the device type corresponding to the application type.
	//
	// This parameter is required.
	//
	// example:
	//
	// HARMONY
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// Expiration time for offline messages/notifications, used in conjunction with StoreOffline. Expired messages will no longer be sent. Maximum retention is 72 hours. Default is 72 hours.
	//
	// The time format follows the ISO8601 standard and must use UTC time, in the format YYYY-MM-DDThh:mm:ssZ. The expiration time must be greater than the current time or the scheduled send time plus 3 seconds (`ExpireTime > PushTime + 3 seconds`). The 3-second buffer accounts for network and system delay tolerance. It is recommended to set at least 1 minute for single push, and at least 10 minutes for full push or batch push.
	//
	// example:
	//
	// 2019-02-20T00:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The action corresponding to the in-app page ability.
	//
	// 	Notice: When HarmonyActionType is APP_CUSTOM_PAGE, at least one of HarmonyUri and HarmonyAction must be provided.
	//
	// For details, see the HarmonyOS official documentation [ClickAction.action](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section152462191216).
	//
	// example:
	//
	// com.example.action
	HarmonyAction *string `json:"HarmonyAction,omitempty" xml:"HarmonyAction,omitempty"`
	// Action after clicking the notification. Valid values:
	//
	// - APP_HOME_PAGE: Open app home page
	//
	// - APP_CUSTOM_PAGE: Open app custom page
	//
	// example:
	//
	// APP_HOME_PAGE
	HarmonyActionType *string `json:"HarmonyActionType,omitempty" xml:"HarmonyActionType,omitempty"`
	// HarmonyOS app badge increment number. Refer to [HarmonyOS badge addNum field description](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section266310382145).</br>
	//
	// Supported from HarmonyOS SDK 1.2.0.
	//
	// example:
	//
	// 1
	HarmonyBadgeAddNum *int32 `json:"HarmonyBadgeAddNum,omitempty" xml:"HarmonyBadgeAddNum,omitempty"`
	// HarmonyOS app badge set number. Refer to [HarmonyOS badge setNum field description](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section266310382145).
	//
	// Supported from HarmonyOS SDK 1.2.0.
	//
	// example:
	//
	// 1
	HarmonyBadgeSetNum *int32 `json:"HarmonyBadgeSetNum,omitempty" xml:"HarmonyBadgeSetNum,omitempty"`
	// Notification message category. After completing the notification message self-classification rights application, this is used to identify the message type. Different notification message types affect how messages are displayed and how alerts are triggered. Valid values:
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
	// - MARKETING: News, content recommendations, social updates, product promotions, financial updates, lifestyle information, surveys, feature recommendations, operational promotions (only identifies content, does not accelerate message delivery), collectively referred to as information and marketing messages
	//
	// For details, see the HarmonyOS official documentation [Notification.category](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117).
	//
	// example:
	//
	// IM
	HarmonyCategory *string `json:"HarmonyCategory,omitempty" xml:"HarmonyCategory,omitempty"`
	// Set the extension attributes of the notification. This attribute does not take effect when PushType is set to MESSAGE.
	//
	// This parameter must be passed in JSON map format, otherwise parsing will fail.
	//
	// example:
	//
	// {"key1":"value1","api_name":"PushNoticeToAndroidRequest"}
	HarmonyExtParameters *string `json:"HarmonyExtParameters,omitempty" xml:"HarmonyExtParameters,omitempty"`
	// Extra data for notification extension messages.</br>
	//
	// Effective when sending HarmonyOS notification extension messages.</br>
	//
	// Conceptually equivalent to the extraData field of HarmonyOS notification extension messages. For the specific definition, refer to [HarmonyOS ExtensionPayload Description](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section161192514234).</br>
	//
	// Supported from HarmonyOS SDK 1.2.0.
	//
	// example:
	//
	// 示例额外数据
	HarmonyExtensionExtraData *string `json:"HarmonyExtensionExtraData,omitempty" xml:"HarmonyExtensionExtraData,omitempty"`
	// When PushType is NOTICE, whether to send as a HarmonyOS notification extension message.
	//
	// - true: Send notification extension message
	//
	// - false: Send standard notification (default)
	//
	// Notification extension messages require permission application on the HarmonyOS side before sending. For details, refer to the HarmonyOS documentation [Send Notification Extension Messages](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V5/push-send-extend-noti-V5).</br>
	//
	// Supported from HarmonyOS SDK 1.2.0.
	//
	// example:
	//
	// true
	HarmonyExtensionPush *bool `json:"HarmonyExtensionPush,omitempty" xml:"HarmonyExtensionPush,omitempty"`
	// URL for the large icon on the right side of the notification. The URL must use the HTTPS protocol.
	//
	// > Supported image formats: png, jpg, jpeg, heif, gif, bmp. Image width 	- height must be less than 25000 pixels.
	//
	// For details, see the HarmonyOS official documentation [Notification.image](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117).
	//
	// example:
	//
	// https://example.com/xxx.png
	HarmonyImageUrl *string `json:"HarmonyImageUrl,omitempty" xml:"HarmonyImageUrl,omitempty"`
	// Content for multi-line text style. Required when HarmonyRenderStyle is MULTI_LINE. Supports up to 3 items.
	//
	// example:
	//
	// ["1.content1","2.content2","3.content3"]
	HarmonyInboxContent *string `json:"HarmonyInboxContent,omitempty" xml:"HarmonyInboxContent,omitempty"`
	// JSON string of the HarmonyOS Live View data structure [LiveViewPayload](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V13/push-scenariozed-api-request-param-V13#section66881469306). For development integration, refer to the documentation [HarmonyOS Live View Push Guide](https://help.aliyun.com/document_detail/2982112.html).
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
	// Use the specified notification channel type. Only effective when the Alibaba Cloud proprietary channel is online.
	//
	// - SOCIAL_COMMUNICATION: Social communication.
	//
	// - SERVICE_INFORMATION: Service reminders.
	//
	// - CONTENT_INFORMATION: Content information.
	//
	// - CUSTOMER_SERVICE: Customer service messages. This type is used for customer service messages between users and merchants, and must be initiated by the user.
	//
	// - OTHER_TYPES: Other.
	//
	// For details, see the HarmonyOS official documentation [SlotType](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/js-apis-notificationmanager-V5#slottype).
	//
	// example:
	//
	// SOCIAL_COMMUNICATION
	HarmonyNotificationSlotType *string `json:"HarmonyNotificationSlotType,omitempty" xml:"HarmonyNotificationSlotType,omitempty"`
	// Unique identifier for each message when displayed as a notification. If not provided, the push service automatically generates a unique identifier for each message. Different notification messages can share the same notifyId, enabling the new message to replace the old one.
	//
	// For details, see the HarmonyOS official documentation [Notification.notifyId](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117).
	//
	// example:
	//
	// 0
	HarmonyNotifyId *int32 `json:"HarmonyNotifyId,omitempty" xml:"HarmonyNotifyId,omitempty"`
	// HarmonyOS channel receipt ID. This receipt ID can be found in the receipt parameter configuration on the HarmonyOS channel push management platform.
	//
	// > If the default receipt configuration on the HarmonyOS channel push management platform is set to the Alibaba Cloud receipt, this is not required. If not, it is recommended to configure the HarmonyOS channel default receipt ID in the Alibaba Cloud EMAS Mobile Push console first.
	//
	// For details, see the HarmonyOS official documentation [pushOptions.receiptId](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section418321011212).
	//
	// example:
	//
	// RCPB***DFD5
	HarmonyReceiptId *string `json:"HarmonyReceiptId,omitempty" xml:"HarmonyReceiptId,omitempty"`
	// When the push type is message and the device is offline, this push will use the supplementary popup feature. Default is false. Only effective when PushType=MESSAGE.
	//
	// If the message-to-notification push is successful, the notification displays the HarmonyRemindTitle and HarmonyRemindBody parameter values set on the server.
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
	// - false: Official message (default)
	//
	// - true: Test message
	//
	// For details, see the HarmonyOS official documentation [pushOptions.testMessage](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section418321011212).
	//
	// example:
	//
	// true
	HarmonyTestMessage *bool `json:"HarmonyTestMessage,omitempty" xml:"HarmonyTestMessage,omitempty"`
	// The URI corresponding to the in-app page ability.
	//
	// 	Notice: When HarmonyActionType is APP_CUSTOM_PAGE, at least one of HarmonyUri and HarmonyAction must be provided. When multiple Abilities exist, fill in the action and uri of each Ability separately. The action is used first to find the corresponding in-app page.
	//
	// For details, see the HarmonyOS official documentation [ClickAction.uri](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section152462191216).
	//
	// example:
	//
	// https://www.example.com:8080/push/example
	HarmonyUri *string `json:"HarmonyUri,omitempty" xml:"HarmonyUri,omitempty"`
	// An idempotent parameter to prevent duplicate pushes caused by API client retries. When the same IdempotentToken is used for calls within 15 minutes, only one push will be made, and subsequent calls will return the result of the first successful push.
	//
	// >
	//
	// > - The parameter format is a standard 36-character UUID (8-4-4-4-12). Each valid character is a hexadecimal digit in the range 0-9 or a-f, case-insensitive.
	//
	// > - This parameter is only used to prevent duplicate pushes caused by retries. It cannot prevent duplicate pushes caused by concurrent calls.
	//
	// example:
	//
	// c8016d13-6e76-410c-9bda-769383d11787
	IdempotentToken *string `json:"IdempotentToken,omitempty" xml:"IdempotentToken,omitempty"`
	// Custom identifier for the push task. When JobKey is not empty, the receipt log will include this field. For viewing receipt logs, see [Receipt Logs](https://help.aliyun.com/document_detail/434651.html).
	//
	// >Format requirements: Only letters, digits, or the symbols \\"_\\" and \\"-\\" (any combination) are allowed, and the length must not exceed 32 characters.
	//
	// example:
	//
	// 123
	JobKey *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	// Used for scheduled sending. If not set, the default is immediate sending.
	//
	// Scheduled sending must be no later than 7 days from now.
	//
	// The time format follows the ISO8601 standard and must use UTC time, in the format YYYY-MM-DDThh:mm:ssZ.
	//
	// >When Target is TBD (continuous push), scheduled sending is not supported.
	//
	// example:
	//
	// 2019-02-20T00:00:00Z
	PushTime *string `json:"PushTime,omitempty" xml:"PushTime,omitempty"`
	// Push type. Valid values:
	//
	// - **NOTICE**: Notification. Notifications are delivered to devices through vendor channels such as APNs, Huawei, Xiaomi, and HarmonyOS, and are displayed directly in the device notification bar. When an Android device is online (app process is alive), the notification is preferentially delivered through the Alibaba Cloud proprietary channel, where the Push SDK constructs and displays the notification, providing better push performance and potentially saving vendor push message quotas in some scenarios.
	//
	// - **MESSAGE**: Message. Messages are delivered through the Alibaba Cloud proprietary online channel. They are not displayed in the notification bar by default, but need to be received and processed by the app when the process is active, allowing the business to decide whether to trigger certain business behaviors. When the device is offline (app process is inactive), messages cannot be received in a timely manner. In this case, you can use the `iOSRemind` or `AndroidRemind` parameters below to convert messages to notifications when the device is offline; or set the `StoreOffline` parameter below so the push system saves the message when the device is offline and automatically delivers it when the device comes online.
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
	// This parameter is deprecated.
	//
	// example:
	//
	// 0
	SendSpeed *int32 `json:"SendSpeed,omitempty" xml:"SendSpeed,omitempty"`
	// Delay time before triggering SMS, in seconds.
	//
	// Must be set when using SMS convergence. Recommended to be 15 seconds or more, with a maximum of 3 days, to avoid duplication between SMS and push notifications.
	//
	// > When SMS convergence is used, the ExpireTime parameter becomes ineffective. The notification expiration time is calculated based on the SmsDelaySecs parameter, with the expiration time being the current time plus SmsDelaySecs.
	//
	// example:
	//
	// 15
	SmsDelaySecs *int32 `json:"SmsDelaySecs,omitempty" xml:"SmsDelaySecs,omitempty"`
	// Variable name-value pairs for the SMS template, in the format: `key1=value1&key2=value2`.
	//
	// example:
	//
	// key1=value1
	SmsParams *string `json:"SmsParams,omitempty" xml:"SmsParams,omitempty"`
	// Condition for triggering SMS. Valid values:
	//
	// - **0**: Triggered when push is not received.
	//
	// - **1**: Triggered when user has not opened the notification.
	//
	// example:
	//
	// 0
	SmsSendPolicy *int32 `json:"SmsSendPolicy,omitempty" xml:"SmsSendPolicy,omitempty"`
	// The signature for supplementary SMS.
	//
	// example:
	//
	// 短信签名
	SmsSignName *string `json:"SmsSignName,omitempty" xml:"SmsSignName,omitempty"`
	// The template name for supplementary SMS. This can be obtained from the SMS template management page and is a system-assigned name, not a developer-defined name.
	//
	// example:
	//
	// 短信模板名称
	SmsTemplateName *string `json:"SmsTemplateName,omitempty" xml:"SmsTemplateName,omitempty"`
	// Whether to store offline messages/notifications. StoreOffline defaults to **false**.
	//
	// If enabled, when the user is offline during push, the message will be resent when the user comes online within the expiration time (ExpireTime). ExpireTime defaults to 72 hours. iOS notifications go through the APNs channel and are not affected by StoreOffline.
	//
	// example:
	//
	// false
	StoreOffline *bool `json:"StoreOffline,omitempty" xml:"StoreOffline,omitempty"`
	// Push target. Valid values:
	//
	// - **DEVICE**: Push by device.
	//
	// - **ACCOUNT**: Push by account.
	//
	// - **ALIAS**: Push by alias.
	//
	// - **TAG**: Push by tag.
	//
	// - **ALL**: Push to all devices (the interval between two full pushes of the same DeviceType must be at least 1 second).
	//
	//  > Pushing to all iOS devices will push to devices that have been active within the last 24 months but have not uninstalled the app. Once APNs (Apple Push Notification service) receives the push request without returning an error, it is considered delivered, which may cause a surge in active device counts and generate significant costs. Please use with discretion.
	//
	// - **TBD**: Initialize continuous push. The push target is specified by the subsequent [ContinuouslyPush](https://help.aliyun.com/document_detail/2249917.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// Set based on the Target type. Multiple values are separated by commas. If the limit is exceeded, split into multiple pushes.
	//
	// - Target=DEVICE: Values such as `deviceid1,deviceid2` (up to 1000).
	//
	// - Target=ACCOUNT: Values such as `account1,account2` (up to 1000).
	//
	// - Target=ALIAS: Values such as `alias1,alias2` (up to 1000).
	//
	// - Target=TAG: Supports single and multiple tags. For the format, see [Tag Format](https://help.aliyun.com/document_detail/434847.html).
	//
	// - Target=ALL: Value is **ALL*	- (fixed parameter for full push).
	//
	// - Target=TBD: Value is **TBD*	- (fixed parameter for continuous push).
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
	// Title of the notification/message during push. Length limit: 200 bytes.
	//
	// Required for Android and HarmonyOS push; optional for iOS notifications. If provided:
	//
	// - iOS 10+: Displayed as the notification title.
	//
	// - iOS 8.2 <= iOS version < iOS 10: Replaces the notification app name.
	//
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Whether to automatically truncate overly long titles and content.
	//
	// >Only applies to vendor channels that explicitly limit title and content length. Does not apply to APNs, Huawei, Honor, and other channels that do not limit title or content individually but only limit the total request body size.
	//
	// example:
	//
	// false
	Trim *bool `json:"Trim,omitempty" xml:"Trim,omitempty"`
	// iOS notifications are sent through the APNs center, and the corresponding environment information must be provided.
	//
	// - **DEV**: Development environment, applicable to apps installed and debugged directly via Xcode.
	//
	// - **PRODUCT**: Production environment, applicable to apps distributed via App Store, TestFlight, Ad Hoc, and enterprise distribution.
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
	// Whether to enable badge auto-increment. Default is false.
	//
	// >When this is set to true, iOSBadge must be empty.
	//
	// The badge auto-increment feature is maintained by the push server for each device\\"s badge count. Users must use SDK version 1.9.5 or above and actively sync the badge count to the server.
	//
	// example:
	//
	// true
	IOSBadgeAutoIncrement *bool `json:"iOSBadgeAutoIncrement,omitempty" xml:"iOSBadgeAutoIncrement,omitempty"`
	// Extension attributes for iOS notifications.
	//
	// For iOS 10+, you can specify the resource URL for rich media push notifications here: `{"attachment": "https://xxxx.xxx/notification_pic.png"}`. This parameter must be passed in JSON map format, otherwise parsing will fail.
	//
	// example:
	//
	// {"attachment": "https://xxxx.xxx/notification_pic.png"}
	IOSExtParameters *string `json:"iOSExtParameters,omitempty" xml:"iOSExtParameters,omitempty"`
	// Interruption level. Valid values:
	//
	// - **passive**: The system adds the notification to the notification list without lighting up the screen or playing a sound.
	//
	// - **active**: The system displays the notification immediately, lights up the screen, and can play a sound.
	//
	// - **time-sensitive**: The system displays the notification immediately, lights up the screen, and can play a sound, but does not break through system notification controls.
	//
	// - **critical**: The system displays the notification immediately, lights up the screen, and plays a sound bypassing the silent switch.
	//
	// example:
	//
	// active
	IOSInterruptionLevel *string `json:"iOSInterruptionLevel,omitempty" xml:"iOSInterruptionLevel,omitempty"`
	// JSON string, static parameters for Live Activity (Dynamic Island) push. Contains static user-defined information such as product IDs and order information.
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
	// Dynamic parameters for Live Activity (Dynamic Island) push, containing real-time update information such as price and inventory changes.
	//
	// example:
	//
	// {"status": "delivered", "estimatedArrival": "2023-12-31T12:00:00Z"}
	IOSLiveActivityContentState *string `json:"iOSLiveActivityContentState,omitempty" xml:"iOSLiveActivityContentState,omitempty"`
	// Timestamp in seconds. The ended Live Activity will remain on the lock screen until this specified time, with a maximum of 4 hours.
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
	// The Live Activity ID reported from the device to the user\\"s server. The unique identifier of the Live Activity.
	//
	// example:
	//
	// 66B94673-B32E-4CA7-863C-3E523054FD46
	IOSLiveActivityId *string `json:"iOSLiveActivityId,omitempty" xml:"iOSLiveActivityId,omitempty"`
	// Timestamp in seconds. Marks the expiration time of the activity content.
	//
	// example:
	//
	// 1743131967
	IOSLiveActivityStaleDate *int64 `json:"iOSLiveActivityStaleDate,omitempty" xml:"iOSLiveActivityStaleDate,omitempty"`
	// iOS notification sound. Specify the name of an audio file stored in the app bundle or the sandbox Library/Sounds directory. See: [How to Set iOS Push Notification Sound](https://help.aliyun.com/document_detail/48906.html).
	//
	// If set to an empty string (""), the notification will be silent; if not set, it defaults to the system alert sound.
	//
	// example:
	//
	// ""
	IOSMusic *string `json:"iOSMusic,omitempty" xml:"iOSMusic,omitempty"`
	// iOS notification processing extension flag (iOS 10+). If set to true, the APNs push notification can reach the Extension for processing before being displayed. For silent notifications, this must be set to true.
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
	// When a device receives messages with the same CollapseId, they will be merged into one. When the device is offline and consecutive messages with the same CollapseId are sent, only the latest one is displayed in the notification bar. iOS 10+ supports this parameter.
	//
	// example:
	//
	// ZD2011
	IOSNotificationCollapseId *string `json:"iOSNotificationCollapseId,omitempty" xml:"iOSNotificationCollapseId,omitempty"`
	// This attribute is used to group iOS remote notifications, identifying the group name for collapsed notifications.
	//
	// Only supported on iOS 12.0+.
	//
	// example:
	//
	// abc
	IOSNotificationThreadId *string `json:"iOSNotificationThreadId,omitempty" xml:"iOSNotificationThreadId,omitempty"`
	// Summary highlight score. Value range: floating-point number in [0,1\\].
	//
	// example:
	//
	// 0.01
	IOSRelevanceScore *float64 `json:"iOSRelevanceScore,omitempty" xml:"iOSRelevanceScore,omitempty"`
	// When the device is offline during message push (i.e., the persistent connection to the push server is disconnected), this push will be delivered as a notification through Apple\\"s APNs channel once.
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
