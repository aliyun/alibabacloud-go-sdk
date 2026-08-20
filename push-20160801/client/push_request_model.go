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
	AndroidOppoPrivateContentParameters map[string]*string `json:"AndroidOppoPrivateContentParameters,omitempty" xml:"AndroidOppoPrivateContentParameters,omitempty"`
	// OPPO private message template ID
	//
	// example:
	//
	// 687557242b1634hzefs3d5013
	AndroidOppoPrivateMsgTemplateId *string `json:"AndroidOppoPrivateMsgTemplateId,omitempty" xml:"AndroidOppoPrivateMsgTemplateId,omitempty"`
	// OPPO private message template title parameters
	AndroidOppoPrivateTitleParameters map[string]*string `json:"AndroidOppoPrivateTitleParameters,omitempty" xml:"AndroidOppoPrivateTitleParameters,omitempty"`
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
