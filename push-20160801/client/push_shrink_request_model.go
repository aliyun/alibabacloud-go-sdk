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
	SetAndroidVivoTemplateId(v string) *PushShrinkRequest
	GetAndroidVivoTemplateId() *string
	SetAndroidVivoTemplateParams(v string) *PushShrinkRequest
	GetAndroidVivoTemplateParams() *string
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
	// The activity to be opened by the notification.
	//
	// This parameter is passed only when AndroidOpenType is set to ACTIVITY. Example: `com.alibaba.cloudpushdemo.bizactivity`.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidActivity *string `json:"AndroidActivity,omitempty" xml:"AndroidActivity,omitempty"`
	// The increment value for the badge. The value is added to the current badge number. Valid values: 1 to 99.
	//
	// > This parameter is valid only for push notifications sent through the Huawei or Honor vendor channel. If both AndroidBadgeAddNum and AndroidBadgeSetNum are specified, AndroidBadgeSetNum takes precedence.
	//
	// example:
	//
	// 1
	AndroidBadgeAddNum *int32 `json:"AndroidBadgeAddNum,omitempty" xml:"AndroidBadgeAddNum,omitempty"`
	// The fully qualified class name of the entry Activity for badge settings.
	//
	// > This parameter is valid only for push notifications sent through the Huawei or Honor vendor channel.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidBadgeClass *string `json:"AndroidBadgeClass,omitempty" xml:"AndroidBadgeClass,omitempty"`
	// The fixed number for the badge. Valid values: 0 to 99.
	//
	// > This parameter takes effect only on the Huawei and Honor channels when push notifications are sent through vendor channels. When push notifications are sent through the Alibaba Cloud proprietary channel, this parameter takes effect only on Huawei, Honor, and vivo devices.
	//
	// example:
	//
	// 5
	AndroidBadgeSetNum *int32 `json:"AndroidBadgeSetNum,omitempty" xml:"AndroidBadgeSetNum,omitempty"`
	// The body in long text mode. Maximum length: 1,000 bytes. One Chinese character counts as 3 bytes. The actual limit is subject to the specific vendor channel during delivery.
	//
	// Currently supported by:
	//
	// - Huawei: EMUI 10 and later.
	//
	// - Honor: Magic UI 4.0 and later.
	//
	// - Xiaomi: MIUI 10 and later.
	//
	// - OPPO: ColorOS 5.0 and later.
	//
	// - Meizu: Flyme.
	//
	// - Proprietary channel: Android SDK 3.6.0 and later.
	//
	// >If this parameter is not provided in long text mode, the first non-empty value from Body and AndroidPopupBody is used.
	//
	// example:
	//
	// Sample long text
	AndroidBigBody *string `json:"AndroidBigBody,omitempty" xml:"AndroidBigBody,omitempty"`
	// The image URL in big image mode. Currently supported by the proprietary channel for Android SDK 3.6.0 and later.
	//
	// example:
	//
	// https://imag.example.com/image.png
	AndroidBigPictureUrl *string `json:"AndroidBigPictureUrl,omitempty" xml:"AndroidBigPictureUrl,omitempty"`
	// The title in long text mode. Maximum length: 200 bytes. One Chinese character counts as 3 bytes.
	//
	// - This parameter is currently supported only by the Honor channel and the Huawei channel for EMUI 11 and later.
	//
	// - If this parameter is not provided in long text mode, the first non-empty value from Title and AndroidPopupTitle is used.
	//
	// example:
	//
	// Sample long title
	AndroidBigTitle *string `json:"AndroidBigTitle,omitempty" xml:"AndroidBigTitle,omitempty"`
	// The extended attributes of the notification. This attribute does not take effect when the push type PushType is set to MESSAGE.
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
	// > Each app can send 1,000 test notifications per day, which is not subject to the daily push limit per device.
	//
	// example:
	//
	// 0
	AndroidHonorTargetUserType *int32 `json:"AndroidHonorTargetUserType,omitempty" xml:"AndroidHonorTargetUserType,omitempty"`
	// The Huawei fast notification parameter. Valid values:
	//
	// - **0**: Send Huawei standard notifications. This is the default value.
	//
	// - **1**: Send Huawei fast notifications.
	//
	// example:
	//
	// 1
	AndroidHuaweiBusinessType *int32 `json:"AndroidHuaweiBusinessType,omitempty" xml:"AndroidHuaweiBusinessType,omitempty"`
	// The JSON string of the Huawei Android Live Notification data structure [LiveNotificationPayload](https://developer.huawei.com/consumer/cn/doc/HMSCore-References/rest-live-0000001562939968#ZH-CN_TOPIC_0000001700850537__p195121620102511). For development and integration, refer to the [Huawei Live Notification push guide](https://help.aliyun.com/document_detail/2983768.html).
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
	// The receipt ID for the Huawei channel. You can view this receipt ID in the receipt parameter settings of the Huawei push operations platform.
	//
	// > If the default receipt configuration on the Huawei push operations platform is set to the Alibaba Cloud receipt, this parameter is not required. Otherwise, we recommend that you configure the default receipt ID for the Huawei channel in the Alibaba Cloud EMAS Mobile Push console first.
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
	// > Each app can send 500 test notifications per day, which is not subject to the daily push limit per device.
	//
	// example:
	//
	// 0
	AndroidHuaweiTargetUserType *int32 `json:"AndroidHuaweiTargetUserType,omitempty" xml:"AndroidHuaweiTargetUserType,omitempty"`
	// The URL of the right-side icon. Currently supported by:
	//
	// - Huawei EMUI: applicable only in long text mode and inbox mode.
	//
	// - Honor Magic UI: applicable only in long text mode.
	//
	// - Proprietary channel: Android SDK 3.5.0 and later.
	//
	// example:
	//
	// https://imag.example.com/image.png
	AndroidImageUrl *string `json:"AndroidImageUrl,omitempty" xml:"AndroidImageUrl,omitempty"`
	// The body in inbox mode. The content must be a valid JSON array with no more than 5 elements. Currently supported by:
	//
	// - Huawei: EMUI 9 and later.
	//
	// - Honor: Magic UI 4.0 and later.
	//
	// - Xiaomi: MIUI 10 and later.
	//
	// - OPPO: ColorOS 5.0 and later.
	//
	// - Proprietary channel: Android SDK 3.6.0 and later.
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
	// true
	//
	// example:
	//
	// 0
	AndroidMeizuNoticeMsgType *int32 `json:"AndroidMeizuNoticeMsgType,omitempty" xml:"AndroidMeizuNoticeMsgType,omitempty"`
	// Purpose 1: After you complete the [self-classification permission application](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835?#section3410731125514), this parameter is used to identify the message type and determine the [message reminder method](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#ZH-CN_TOPIC_0000001149358835__p3850133955718). It accelerates the delivery of specific types of messages. For valid values, refer to the [message classification standards](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section1076611477914) in the official Huawei push documentation, and fill in the "Cloud notification category value" or "Local notification category value" in the documentation table.
	//
	// Purpose 2: After you [apply for special permissions](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/faq-0000001050042183#section037425218509), this parameter is used to identify high-priority pass-through scenarios. Valid values:
	//
	// - VOIP: Audio and video calls.
	//
	// - PLAY_VOICE: Voice broadcast.
	//
	// > For messages whose "Cloud notification category value" is "Not involved", the Alibaba Cloud proprietary channel is used. For messages whose "Local notification category value" is "Not involved", the Huawei channel is used.
	//
	// example:
	//
	// VOIP
	AndroidMessageHuaweiCategory *string `json:"AndroidMessageHuaweiCategory,omitempty" xml:"AndroidMessageHuaweiCategory,omitempty"`
	// The delivery priority of Huawei channel notifications. Valid values:
	//
	// - **HIGH**
	//
	// - **NORMAL**
	//
	// You must apply for the permission. For more information, see [Application link](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/faq-0000001050042183#section037425218509).
	//
	// example:
	//
	// HIGH
	AndroidMessageHuaweiUrgency *string `json:"AndroidMessageHuaweiUrgency,omitempty" xml:"AndroidMessageHuaweiUrgency,omitempty"`
	// OPPO message categorization includes two categories for management: Communication and Service, and Content and Marketing.
	//
	// Communication and Service (permission application required):
	//
	// - IM: Instant chat, audio, and video calls.
	//
	// - ACCOUNT: Individual account and asset changes.
	//
	// - DEVICE_REMINDER: Personal device reminders.
	//
	// - ORDER: Personal order and logistics status changes.
	//
	// - TODO: Personal schedules and to-dos.
	//
	// - SUBSCRIPTION: Individual subscribe.
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
	// The reminder level of OPPO channel notification bar messages. Valid values:
	//
	// - 1: Notification bar.
	//
	// - 2: Notification bar, lock screen, ringtone, and vibrate. This is the default notification level for Communication and Service messages.
	//
	// - 16: Notification bar, lock screen, ringtone, vibrate, and banner. Permission application is required.
	//
	// > When you use the AndroidMessageOppoNotifyLevel parameter, you must also pass the AndroidMessageOppoCategory parameter.
	//
	// example:
	//
	// 1
	AndroidMessageOppoNotifyLevel *int32 `json:"AndroidMessageOppoNotifyLevel,omitempty" xml:"AndroidMessageOppoNotifyLevel,omitempty"`
	// vivo message categorization includes two categories for management: System Messages and Operation Messages.
	//
	// System Messages:
	//
	// - IM: Instant messages.
	//
	// - ACCOUNT: Accounts and assets.
	//
	// - TODO: Schedules and to-dos.
	//
	// - DEVICE_REMINDER: Device information.
	//
	// - ORDER: Orders and logistics.
	//
	// - SUBSCRIPTION: subscribe reminders.
	//
	// Operation Messages:
	//
	// - NEWS: News.
	//
	// - CONTENT: Content recommendations.
	//
	// - MARKETING: Operational activity.
	//
	// - SOCIAL: Social updates.
	//
	// For more information, see [Classification description](https://dev.vivo.com.cn/documentCenter/doc/359#s-ef3qugc3).
	//
	// example:
	//
	// TODO
	AndroidMessageVivoCategory *string `json:"AndroidMessageVivoCategory,omitempty" xml:"AndroidMessageVivoCategory,omitempty"`
	// The notification sound for the Huawei vendor channel. Specify the name of the audio file stored in the `app/src/main/res/raw/` directory of the client project. Do not include the file extension.
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
	// The channel ID of the Android app. It must correspond to the channel ID in the app.
	//
	// - For more information about how to set the NotificationChannel parameter, see [FAQ: Notifications are not received on Android 8.0 and later devices](https://help.aliyun.com/document_detail/67398.html).
	//
	// - Because the channel_id of the OPPO private message channel is the same as the channel ID of the app, this value is used as the channel_id when pushing through the OPPO channel.
	//
	// - This value is used as the channel_id when pushing through the Huawei, FCM, and Alibaba Cloud proprietary channels.
	//
	// example:
	//
	// 1
	AndroidNotificationChannel *string `json:"AndroidNotificationChannel,omitempty" xml:"AndroidNotificationChannel,omitempty"`
	// The message group. For messages in the same group, only the latest message and the total number of messages received by the group are displayed in the notification bar. All messages are not displayed and cannot be expanded. Currently supported by:
	//
	// - Huawei vendor channel.
	//
	// - Honor vendor channel.
	//
	// - Proprietary channel: Android SDK 3.9.1 and earlier.
	//
	// > The proprietary channel no longer supports this parameter in Android SDK 3.9.2 and later.
	//
	// example:
	//
	// group-1
	AndroidNotificationGroup *string `json:"AndroidNotificationGroup,omitempty" xml:"AndroidNotificationGroup,omitempty"`
	// The importance parameter for Honor notification message classification, which determines the notification behavior on user devices. Valid values:
	//
	// - **LOW**: Information and marketing messages.
	//
	// - **NORMAL**: Service and communication messages.
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
	// >- Use AndroidMessageHuaweiCategory for notification classification in the Huawei channel. AndroidNotificationHuaweiChannel is no longer required.
	//
	// >- You must apply for this permission on the Huawei platform. For more information, see [Application link](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section893184112272).
	//
	// example:
	//
	// LOW
	AndroidNotificationHuaweiChannel *string `json:"AndroidNotificationHuaweiChannel,omitempty" xml:"AndroidNotificationHuaweiChannel,omitempty"`
	// The unique identifier of each message when displayed in the notification bar. Different notification bar messages can use the same NotifyId to allow new notification bar messages to overwrite old ones.
	//
	// example:
	//
	// 100001
	AndroidNotificationNotifyId *int32 `json:"AndroidNotificationNotifyId,omitempty" xml:"AndroidNotificationNotifyId,omitempty"`
	// The message thread ID. Messages in the same group are collapsed and can be expanded in the notification bar. Notifications in different groups are displayed separately. Currently supported by:
	//
	// - Proprietary channel: Android SDK 3.9.2 and later.
	//
	// example:
	//
	// thread-1
	AndroidNotificationThreadId *string `json:"AndroidNotificationThreadId,omitempty" xml:"AndroidNotificationThreadId,omitempty"`
	// The vivo notification message classification. Valid values:
	//
	// - 0: Operation messages. This is the default value.
	//
	// - 1: System messages.
	//
	// >- Use AndroidMessageVivoCategory for notification classification in the vivo channel. AndroidNotificationVivoChannel is no longer required.
	//
	// >- You must apply for this permission on the vivo platform. For more information, see [Application link](https://dev.vivo.com.cn/documentCenter/doc/359).
	//
	// example:
	//
	// classification
	AndroidNotificationVivoChannel *string `json:"AndroidNotificationVivoChannel,omitempty" xml:"AndroidNotificationVivoChannel,omitempty"`
	// The channel ID for Xiaomi notification types. You must apply for it on the Xiaomi platform. For more information, see [Application link](https://dev.mi.com/console/doc/detail?pId=2422#_4).
	//
	// >- A single app can apply for a maximum of 8 channels on the Xiaomi channel. Plan ahead.
	//
	// example:
	//
	// michannel
	AndroidNotificationXiaomiChannel *string `json:"AndroidNotificationXiaomiChannel,omitempty" xml:"AndroidNotificationXiaomiChannel,omitempty"`
	// The reminder method of the notification. Valid values:
	//
	// - **VIBRATE**: Vibrate. This is the default value.
	//
	// - **SOUND**: Sound.
	//
	// - **BOTH**: Sound and vibrate.
	//
	// - **NONE**: Silent.
	//
	// example:
	//
	// BOTH
	AndroidNotifyType *string `json:"AndroidNotifyType,omitempty" xml:"AndroidNotifyType,omitempty"`
	// The action performed after the notification is clicked. Valid values:
	//
	// - **APPLICATION**: Open the app. This is the default value.
	//
	// - **ACTIVITY**: Open the specified Android activity.
	//
	// - **URL**: Open the specified URL.
	//
	// - **NONE**: No redirection.
	//
	// example:
	//
	// APPLICATION
	AndroidOpenType *string `json:"AndroidOpenType,omitempty" xml:"AndroidOpenType,omitempty"`
	// The URL to be opened after the Android device receives the push.
	//
	// This parameter is passed only when AndroidOpenType is set to URL.
	//
	// example:
	//
	// https://www.example.com
	AndroidOpenUrl *string `json:"AndroidOpenUrl,omitempty" xml:"AndroidOpenUrl,omitempty"`
	// The JSON character string of the OPPO Fluid Cloud intent delete data structure [data](https://open.oppomobile.com/documentation/page/info?id=13578). This parameter is invalid when the AndroidOppoIntelligentIntent parameter is specified. For development and integration, see References: [OPPO Fluid Cloud push guide](https://help.aliyun.com/document_detail/2997310.html).
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
	// The JSON character string of the OPPO Fluid Cloud intent sharing data structure [IntelligentIntent](https://open.oppomobile.com/documentation/page/info?id=13565). For development and integration, see References: [OPPO Fluid Cloud push guide](https://help.aliyun.com/document_detail/2997310.html).
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
	// > For the OPPO Fluid Cloud test environment, refer to [Environment setup](https://open.oppomobile.com/documentation/page/info?id=13590) to set up the client.
	//
	// example:
	//
	// 1
	AndroidOppoIntentEnv *int32 `json:"AndroidOppoIntentEnv,omitempty" xml:"AndroidOppoIntentEnv,omitempty"`
	// The content parameters for the OPPO private message template.
	AndroidOppoPrivateContentParametersShrink *string `json:"AndroidOppoPrivateContentParameters,omitempty" xml:"AndroidOppoPrivateContentParameters,omitempty"`
	// The OPPO private message template ID.
	//
	// example:
	//
	// 687557242b1634hzefs3d5013
	AndroidOppoPrivateMsgTemplateId *string `json:"AndroidOppoPrivateMsgTemplateId,omitempty" xml:"AndroidOppoPrivateMsgTemplateId,omitempty"`
	// The title parameters for the OPPO private message template.
	AndroidOppoPrivateTitleParametersShrink *string `json:"AndroidOppoPrivateTitleParameters,omitempty" xml:"AndroidOppoPrivateTitleParameters,omitempty"`
	// The activity to which the user is redirected after the notification is clicked.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidPopupActivity *string `json:"AndroidPopupActivity,omitempty" xml:"AndroidPopupActivity,omitempty"`
	// The body content in auxiliary pop-up window mode. This parameter is required when the **AndroidPopupActivity*	- parameter is not empty.
	//
	// Maximum length: 200 characters. Both Chinese and English characters count as one character.
	//
	// If a vendor channel is used, the limits of the vendor channel must also be met. For more information, see [Limits on auxiliary channel pushes for Android](https://help.aliyun.com/document_detail/165253.html).
	//
	// example:
	//
	// hello
	AndroidPopupBody *string `json:"AndroidPopupBody,omitempty" xml:"AndroidPopupBody,omitempty"`
	// The title content in auxiliary pop-up window mode. This parameter is required when the **AndroidPopupActivity*	- parameter is not empty.
	//
	// Maximum length: 50 characters. Both Chinese and English characters count as one character.
	//
	// If a vendor channel is used, the limits of the vendor channel must also be met. For more information, see [Limits on auxiliary channel pushes for Android](https://help.aliyun.com/document_detail/165253.html).
	//
	// example:
	//
	// hello
	AndroidPopupTitle *string `json:"AndroidPopupTitle,omitempty" xml:"AndroidPopupTitle,omitempty"`
	// Specifies whether to use the auxiliary pop-up window feature when the device is offline during a message push. Default value: false. This parameter takes effect only when PushType is set to MESSAGE.
	//
	// If the message is successfully converted into a notification and pushed, the data displayed in the received notification is the values of the AndroidPopupTitle and AndroidPopupBody parameters set on the server. The data obtained in the onSysNoticeOpened method of the auxiliary pop-up window when the notification is clicked is the values of the Title and Body parameters set on the server.
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
	// > If a non-standard mode is used, this parameter must be provided.
	//
	// example:
	//
	// 1
	AndroidRenderStyle *int32 `json:"AndroidRenderStyle,omitempty" xml:"AndroidRenderStyle,omitempty"`
	// Settings for the vendor channel notification type. Valid values:
	//
	// - **0**: Formal notification. This is the default value.
	//
	// - **1**: Test notification.
	//
	// >- These parameter settings are equivalent to configuring the AndroidHuaweiTargetUserType, AndroidHonorTargetUserType, AndroidVivoPushMode, and AndroidOppoIntentEnv parameters at the same time. The corresponding parameters for specific vendor channels can overwrite this parameter.
	//
	// >- Currently supported by: Huawei channel, Honor channel, vivo channel, and OPPO Fluid Cloud.
	//
	// example:
	//
	// 0
	AndroidTargetUserType *int32 `json:"AndroidTargetUserType,omitempty" xml:"AndroidTargetUserType,omitempty"`
	// The JSON character string of the vivo Atomic Island data structure [liveMessage](https://dev.vivo.com.cn/documentCenter/doc/896#s-fdagzbd4). For development and integration, see References: [vivo Atomic Island push guide](https://www.alibabacloud.com/help/en/document_detail/3030718.html).
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
	// > For test pushes, configure test devices in the vivo console in advance. You can search for "onReceiveRegId regId" in the device startup logs to obtain the RegId of the test device.
	//
	// example:
	//
	// 0
	AndroidVivoPushMode *int32 `json:"AndroidVivoPushMode,omitempty" xml:"AndroidVivoPushMode,omitempty"`
	// The receipt ID for the vivo channel. You can view this receipt ID in the application information of the push service on the vivo Open Platform.
	//
	// > If the default receipt configuration on the vivo Open Platform is set to the Alibaba Cloud receipt, this parameter is not required. Otherwise, we recommend that you configure the default receipt ID for the vivo channel in the Alibaba Cloud EMAS Mobile Push console first.
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
	// The parameters for the vivo template. The value is a JSON object string.
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
	// **[Deprecated]*	- Starting from August 2023, Xiaomi officially no longer supports dynamically setting small icons, right-side icons, and big images during pushes on new devices and systems.
	//
	// example:
	//
	// https://f6.market.xiaomi.com/download/MiPass/aaa/bbb.png
	AndroidXiaomiBigPictureUrl *string `json:"AndroidXiaomiBigPictureUrl,omitempty" xml:"AndroidXiaomiBigPictureUrl,omitempty"`
	// The JSON character string of the Xiaomi Hyper Island data structure [miui.focus.param](https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131). For development and integration, see References: [Xiaomi Hyper Island push guide](https://www.alibabacloud.com/help/en/document_detail/3037956.html).
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
	// The JSON character string of the Xiaomi Hyper Island image data [miui.focus.pic_xxx](https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131). For development and integration, see References: [Xiaomi Hyper Island push guide](https://www.alibabacloud.com/help/en/document_detail/3037956.html).
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
	// **[Deprecated]*	- Starting from August 2023, Xiaomi officially no longer supports dynamically setting small icons, right-side icons, and big images during pushes on new devices and systems.
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
	// The parameters for the Xiaomi private message template. The value is a JSON string.
	//
	// example:
	//
	// {"keywords1":"Tom","keywords2":"phone"}
	AndroidXiaomiTemplateParams *string `json:"AndroidXiaomiTemplateParams,omitempty" xml:"AndroidXiaomiTemplateParams,omitempty"`
	// The AppKey information.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The content of the notification or message for Android and HarmonyOS pushes, or the content of the iOS message or notification. The content size is limited. For more information, see [Limits](https://help.aliyun.com/document_detail/434629.html).
	//
	// example:
	//
	// hello
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The device type. Valid values:
	//
	// - **HARMONY**: HarmonyOS devices.
	//
	// - **iOS**: iOS devices.
	//
	// - **ANDROID**: Android devices.
	//
	// - **ALL**: If the AppKey belongs to a legacy dual-platform app, this value indicates that the push is sent to both Android and iOS devices. If the AppKey belongs to a new single-platform app, this value has the same effect as specifying the device type corresponding to the app type.
	//
	// This parameter is required.
	//
	// example:
	//
	// HARMONY
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The expiration time of offline messages and notifications. This parameter is used together with StoreOffline. Expired messages and notifications are not delivered. The maximum retention period is 72 hours. Default value: 72 hours.
	//
	// The time must be in the ISO 8601 standard format and use UTC time. Format: YYYY-MM-DDThh:mm:ssZ. The expiration time cannot be earlier than the current time or the scheduled delivery time plus 3 seconds (ExpireTime > PushTime + 3 seconds). The 3 seconds are reserved to compensate for network and system latency. Set the expiration time to at least 1 minute for single pushes and at least 10 minutes for full pushes and batch pushes.
	//
	// example:
	//
	// 2019-02-20T00:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The action corresponding to the ability of the built-in page in the application.
	//
	// 	Notice: When HarmonyActionType is set to APP_CUSTOM_PAGE, you must specify at least one of the HarmonyUri and HarmonyAction fields.
	//
	// For more information, see [ClickAction.action](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section152462191216) on the HarmonyOS official website.
	//
	// example:
	//
	// com.example.action
	HarmonyAction *string `json:"HarmonyAction,omitempty" xml:"HarmonyAction,omitempty"`
	// The action to perform after the notification is tapped. Valid values:
	//
	// - APP_HOME_PAGE: Opens the application home page.
	//
	// - APP_CUSTOM_PAGE: Opens a custom page in the application.
	//
	// example:
	//
	// APP_HOME_PAGE
	HarmonyActionType *string `json:"HarmonyActionType,omitempty" xml:"HarmonyActionType,omitempty"`
	// The increment value for the badge of the HarmonyOS application. For more information, see [Description of the addNum field for HarmonyOS badges](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section266310382145).</br>
	//
	// This parameter is supported in HarmonyOS SDK 1.2.0 and later.
	//
	// example:
	//
	// 1
	HarmonyBadgeAddNum *int32 `json:"HarmonyBadgeAddNum,omitempty" xml:"HarmonyBadgeAddNum,omitempty"`
	// The fixed number for the badge of the HarmonyOS application. For more information, see [Description of the setNum field for HarmonyOS badges](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section266310382145).
	//
	// This parameter is supported in HarmonyOS SDK 1.2.0 and later.
	//
	// example:
	//
	// 1
	HarmonyBadgeSetNum *int32 `json:"HarmonyBadgeSetNum,omitempty" xml:"HarmonyBadgeSetNum,omitempty"`
	// The category of the notification message. After you apply for the notification message auto-categorization permission, this parameter is used as the identity of the message type. Different notification message types affect the display and reminder methods of the message. Valid values:
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
	// - MARKETING: News, content recommendations, social updates, product promotions, financial updates, lifestyle information, surveys, feature recommendations, and operational activity. This category only serves as the identity for the content and does not send messages faster. These are collectively referred to as information and marketing messages.
	//
	// For details, see [Notification.category](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117) on the HarmonyOS official website.
	//
	// example:
	//
	// IM
	HarmonyCategory *string `json:"HarmonyCategory,omitempty" xml:"HarmonyCategory,omitempty"`
	// The extended attributes of the notification. This attribute does not take effect when the push type PushType is set to MESSAGE.
	//
	// This parameter must be passed in JSON map format. Otherwise, a parsing error occurs.
	//
	// example:
	//
	// {"key1":"value1","api_name":"PushNoticeToAndroidRequest"}
	HarmonyExtParameters *string `json:"HarmonyExtParameters,omitempty" xml:"HarmonyExtParameters,omitempty"`
	// The extra data for the extended notification message.</br>
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
	// - false: Sends a normal notification message. This is the default value.
	//
	// You must apply for the required permissions on the HarmonyOS side before you can send extended notification messages. For more information, see [Sending extended notification messages](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V5/push-send-extend-noti-V5) in the HarmonyOS documentation.</br>
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
	// The content for the multi-line text style. This field is required when HarmonyRenderStyle is set to MULTI_LINE. A maximum of three content items are supported.
	//
	// example:
	//
	// ["1.content1","2.content2","3.content3"]
	HarmonyInboxContent *string `json:"HarmonyInboxContent,omitempty" xml:"HarmonyInboxContent,omitempty"`
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
	HarmonyLiveViewPayload *string `json:"HarmonyLiveViewPayload,omitempty" xml:"HarmonyLiveViewPayload,omitempty"`
	// The type of the notification channel to use. This parameter is valid only when the Alibaba Cloud proprietary channel is online. Valid values:
	//
	// - SOCIAL_COMMUNICATION: Social communication.
	//
	// - SERVICE_INFORMATION: Service reminders.
	//
	// - CONTENT_INFORMATION: Content information.
	//
	// - CUSTOMER_SERVICE: Customer service messages. This type is used for customer service messages between users and merchants, and must be initiated by the user.
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
	// The receipt ID for the HarmonyOS channel. You can view this receipt ID in the receipt parameter settings of the HarmonyOS push operations platform.
	//
	// > If the default receipt configuration on the HarmonyOS push operations platform is set to the Alibaba Cloud receipt, this parameter is not required. Otherwise, we recommend that you configure the default receipt ID for the HarmonyOS channel in the Alibaba Cloud EMAS Mobile Push console first.
	//
	// For more information, see [pushOptions.receiptId](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section418321011212) on the HarmonyOS official website.
	//
	// example:
	//
	// RCPB***DFD5
	HarmonyReceiptId *string `json:"HarmonyReceiptId,omitempty" xml:"HarmonyReceiptId,omitempty"`
	// Specifies whether to use the auxiliary pop-up window feature when the device is offline and the push type is message. Default value: false. This parameter takes effect only when PushType is set to MESSAGE.
	//
	// If the message-to-notification push is successful, the notification displays the data specified by the HarmonyRemindTitle and HarmonyRemindBody parameters on the server.
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
	// The test message identifier. Valid values:
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
	// The URI corresponding to the ability of the built-in page in the application.
	//
	// 	Notice: When HarmonyActionType is set to APP_CUSTOM_PAGE, you must specify at least one of the HarmonyUri and HarmonyAction fields. If multiple abilities exist, specify the action and URI for each ability. The action is used first to find the corresponding built-in page in the application.
	//
	// For more information, see [ClickAction.uri](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section152462191216) on the HarmonyOS official website.
	//
	// example:
	//
	// https://www.example.com:8080/push/example
	HarmonyUri *string `json:"HarmonyUri,omitempty" xml:"HarmonyUri,omitempty"`
	// An idempotent parameter used to prevent duplicate push notifications on the server caused by retries from the API client. If the same IdempotentToken is used for calls within 15 minutes, only one push notification is sent, and subsequent calls return the result of the first successful push.
	//
	// >
	//
	// > - The parameter format is a standard 36-character UUID (8-4-4-4-12). Each valid character is a hexadecimal digit in the range of 0-9 or a-f, and it is case-insensitive.
	//
	// > - This parameter is used only to prevent duplicate push notifications caused by retries. It cannot prevent duplicate push notifications caused by concurrent calls.
	//
	// example:
	//
	// c8016d13-6e76-410c-9bda-769383d11787
	IdempotentToken *string `json:"IdempotentToken,omitempty" xml:"IdempotentToken,omitempty"`
	// The custom identifier of the push task. If JobKey is not empty, this field is included in the receipt log. For more information, see [Receipt logs](https://help.aliyun.com/document_detail/434651.html).
	//
	// >Format requirements: The value can only contain letters, digits, underscores (_), and hyphens (-), and cannot exceed 32 characters in length.
	//
	// example:
	//
	// 123
	JobKey *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	// The scheduled delivery time. If this parameter is not set, the push is delivered immediately.
	//
	// The scheduled delivery time cannot be later than 7 days from the current time.
	//
	// The time must be in the ISO 8601 standard format and use UTC time. Format: YYYY-MM-DDThh:mm:ssZ.
	//
	// >This parameter is not supported when Target is set to TBD (continuous push).
	//
	// example:
	//
	// 2019-02-20T00:00:00Z
	PushTime *string `json:"PushTime,omitempty" xml:"PushTime,omitempty"`
	// The push type. Valid values:
	//
	// - **NOTICE**: Notification. Notifications are delivered to device endpoints through vendor channels such as APNs, Huawei, Xiaomi, and HarmonyOS, and are displayed directly in the notification bar of the device. When an Android device is online (the app process is alive), the notification is preferentially delivered through the Alibaba Cloud proprietary channel. The push SDK constructs and displays the notification to achieve better push performance and save vendor push message quotas in certain scenarios.
	//
	// - **MESSAGE**: Message. Messages are delivered through the Alibaba Cloud proprietary online channel. They are not displayed in the notification bar by default. Instead, the app must receive and process them when the process is active, and the business logic determines whether to trigger specific business behaviors. When the device is offline (the app process is inactive), messages cannot be received in a timely manner. In this case, you can use the iOSRemind or AndroidRemind parameter described below to convert messages into notifications for delivery when the device is offline. Alternatively, you can set the StoreOffline parameter described below to allow the push system to save the message when the device is offline and automatically deliver it when the device comes online.
	//
	// This parameter is required.
	//
	// example:
	//
	// MESSAGE
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	// The specified delivery channels. Valid values:
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
	// - gcm: Google GCM channel (legacy HTTP).
	//
	// - fcm: Google Firebase channel (HTTP v1 API).
	//
	// - apns: APNs channel.
	//
	// - harmony: HarmonyOS channel.
	//
	// >- If this parameter is not configured, all channels can be used.
	//
	// >- If this parameter is configured, only the specified channels are used.
	//
	// >- If the configured channels conflict with the delivery policy (for example, iOS notifications can only be delivered through the APNs channel, but apns is not included in this parameter), the push is not delivered.
	//
	// >- If gcm is configured, both Google GCM and FCM channels can be used. If fcm is configured, only the Google FCM channel can be used.
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
	// The delay time before the SMS message is triggered. Unit: seconds.
	//
	// This parameter is required if you use the SMS linkage feature. We recommend that you set this parameter to 15 seconds or more, and the maximum value cannot exceed 3 days. This prevents duplicate SMS and push messages.
	//
	// > When the SMS linkage feature is used, the ExpireTime parameter becomes invalid. The expiration time of the notification is calculated based on the SmsDelaySecs parameter. The expiration time is the current time plus the value of SmsDelaySecs.
	//
	// example:
	//
	// 15
	SmsDelaySecs *int32 `json:"SmsDelaySecs,omitempty" xml:"SmsDelaySecs,omitempty"`
	// The key-value pairs of the variables in the SMS template. Format: `key1=value1&key2=value2`.
	//
	// example:
	//
	// key1=value1
	SmsParams *string `json:"SmsParams,omitempty" xml:"SmsParams,omitempty"`
	// The condition that triggers the SMS message. Valid values:
	//
	// - **0**: Triggered when the push message is not received.
	//
	// - **1**: Triggered when the user does not open the push message.
	//
	// example:
	//
	// 0
	SmsSendPolicy *int32 `json:"SmsSendPolicy,omitempty" xml:"SmsSendPolicy,omitempty"`
	// The signature for the fallback SMS message.
	//
	// example:
	//
	// SMS signature
	SmsSignName *string `json:"SmsSignName,omitempty" xml:"SmsSignName,omitempty"`
	// The name of the SMS template used for fallback SMS messages. You can obtain this name on the SMS template management page. This is the name assigned by the system, not the name set by the developer.
	//
	// example:
	//
	// SMS template name
	SmsTemplateName *string `json:"SmsTemplateName,omitempty" xml:"SmsTemplateName,omitempty"`
	// Specifies whether to save offline messages and notifications. Default value: **false**.
	//
	// If this parameter is set to true and the user is offline during the push, the message or notification is sent again when the user comes online within the expiration time specified by ExpireTime. The default value of ExpireTime is 72 hours. iOS notifications are delivered through the APNs channel and are not affected by StoreOffline.
	//
	// example:
	//
	// false
	StoreOffline *bool `json:"StoreOffline,omitempty" xml:"StoreOffline,omitempty"`
	// The push target. Valid values:
	//
	// - **DEVICE**: Push by device.
	//
	// - **ACCOUNT**: Push by account.
	//
	// - **ALIAS**: Push by alias.
	//
	// - **TAG**: Push by tag.
	//
	// - **ALL**: Push to all devices. The interval between two full pushes of the same DeviceType must be at least 1 second.
	//
	//  > For a full push to iOS devices, the push is sent to devices that have been active within the last 24 months and have not been uninstalled. Once Apple Push Notification service (APNs) receives the push request and does not return an error message, the push is considered delivered. This may cause a surge in the number of active devices and incur high costs. Use this value with caution.
	//
	// - **TBD**: Initialize a continuous push. The push target is specified by the subsequent [ContinuouslyPush](https://help.aliyun.com/document_detail/2249917.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The target value, which is set based on the Target type. Separate multiple values with commas. If the number of values exceeds the limit, you must send the push in multiple batches.
	//
	// - If Target is set to DEVICE, the value is in the format of `deviceid1,deviceid2`. A maximum of 1,000 devices are supported.
	//
	// - If Target is set to ACCOUNT, the value is in the format of `account1,account2`. A maximum of 1,000 accounts are supported.
	//
	// - If Target is set to ALIAS, the value is in the format of `alias1,alias2`. A maximum of 1,000 aliases are supported.
	//
	// - If Target is set to TAG, single or multiple tags are supported. For more information about the format, see [Tag format](https://help.aliyun.com/document_detail/434847.html).
	//
	// - If Target is set to ALL, the value is **ALL**, which is a fixed parameter combination for full pushes.
	//
	// - If Target is set to TBD, the value is **TBD**, which is a fixed parameter combination for continuous pushes.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
	// The title of the notification or message during the push. Maximum length: 200 bytes.
	//
	// This parameter is required for Android and HarmonyOS pushes. This parameter is optional for iOS notifications. If this parameter is specified:
	//
	// - For iOS 10 and later, the title is displayed in the notification.
	//
	// - For iOS 8.2 to iOS 9.x, the title replaces the app name in the notification.
	//
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Specifies whether to automatically truncate excessively long titles and content.
	//
	// >This parameter applies only to vendor channels that explicitly limit the length of titles and content. It does not apply to channels such as APNs, Huawei, and Honor, which do not limit the length of titles and content but only limit the total request body size.
	//
	// example:
	//
	// false
	Trim *bool `json:"Trim,omitempty" xml:"Trim,omitempty"`
	// To send iOS notifications through the APNs center, you must specify the corresponding environment context. Valid values:
	//
	// - **DEV**: Development environment. This value applies to apps that are directly installed and debugged by using Xcode.
	//
	// - **PRODUCT**: Production environment. This value applies to apps distributed through the App Store, TestFlight, Ad Hoc, and enterprise distribution.
	//
	// example:
	//
	// DEV
	IOSApnsEnv *string `json:"iOSApnsEnv,omitempty" xml:"iOSApnsEnv,omitempty"`
	// The badge number displayed in the upper-right corner of the iOS app icon.
	//
	// > If iOSBadgeAutoIncrement is set to true, this parameter must be left empty.
	//
	// example:
	//
	// 0
	IOSBadge *int32 `json:"iOSBadge,omitempty" xml:"iOSBadge,omitempty"`
	// Specifies whether to enable the automatic badge increment feature. Default value: false.
	//
	// >When this parameter is set to true, iOSBadge must be left empty.
	//
	// The automatic badge increment feature maintains the badge count for each device on the push server. To use this feature, you must use SDK version 1.9.5 or later and actively synchronize the badge number to the server.
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
	// - **passive**: The system adds the notification to the notification list without lighting up the screen or playing a sound.
	//
	// - **active**: The system immediately displays the notification, lights up the screen, and can play a sound.
	//
	// - **time-sensitive**: The system immediately presents the notification, lights up the screen, and can play a sound, but does not break through the system notification controls.
	//
	// - **critical**: The system immediately displays the notification, lights up the screen, and plays a sound bypassing the silent switch.
	//
	// example:
	//
	// active
	IOSInterruptionLevel *string `json:"iOSInterruptionLevel,omitempty" xml:"iOSInterruptionLevel,omitempty"`
	// The JSON character string that contains the pass-through static parameter for Dynamic Island pushes. It includes static user-defined information, such as product IDs and order information.
	//
	// > This parameter is required when iOSLiveActivityEvent is set to start.
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
	// The pass-through dynamic parameter for Dynamic Island pushes. It includes real-time update information, such as price and inventory changes.
	//
	// example:
	//
	// {"status": "delivered", "estimatedArrival": "2023-12-31T12:00:00Z"}
	IOSLiveActivityContentState *string `json:"iOSLiveActivityContentState,omitempty" xml:"iOSLiveActivityContentState,omitempty"`
	// The UNIX timestamp in seconds. The ended Live Activity is retained on the Lock Screen until the specified time. The maximum retention period is 4 hours.
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
	// The Live Activity ID reported by the device to the user server. It is the unique identifier of the Live Activity.
	//
	// example:
	//
	// 66B94673-B32E-4CA7-863C-3E523054FD46
	IOSLiveActivityId *string `json:"iOSLiveActivityId,omitempty" xml:"iOSLiveActivityId,omitempty"`
	// The UNIX timestamp in seconds that marks the expiration time of the activity content.
	//
	// example:
	//
	// 1743131967
	IOSLiveActivityStaleDate *int64 `json:"iOSLiveActivityStaleDate,omitempty" xml:"iOSLiveActivityStaleDate,omitempty"`
	// The sound of the iOS notification. Specify the name of the audio file stored in the app bundle or the sandbox Library/Sounds directory. For more information, see [How to set the notification sound for iOS pushes](https://help.aliyun.com/document_detail/48906.html).
	//
	// If this parameter is set to an empty string (""), the notification is silent. If this parameter is not set, the default value is default, which indicates the system prompt sound.
	//
	// example:
	//
	// ""
	IOSMusic *string `json:"iOSMusic,omitempty" xml:"iOSMusic,omitempty"`
	// The iOS notification processing extension flag (iOS 10 and later). If this parameter is set to true, notifications pushed by APNs can be processed by the extension before they are displayed. This parameter must be set to true for silent notifications.
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
	// The collapse ID of the iOS notification. If a device receives messages with the same CollapseId, the messages are merged into one. If the device is offline and continuously receives messages with the same CollapseId, only one notification is displayed in the notification bar. This parameter is supported in iOS 10 and later.
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
	// The relevance score for summary highlighting. Valid values: a floating-point number in the range of [0,1\\].
	//
	// example:
	//
	// 0.01
	IOSRelevanceScore *float64 `json:"iOSRelevanceScore,omitempty" xml:"iOSRelevanceScore,omitempty"`
	// Specifies whether to convert a message into a notification and deliver it once through the Apple APNs channel when the device is offline (that is, the persistent connection channel to the Mobile Push server is disconnected) during a message push.
	//
	// > Converting offline messages into notifications applies only to the production environment.
	//
	// example:
	//
	// true
	IOSRemind *bool `json:"iOSRemind,omitempty" xml:"iOSRemind,omitempty"`
	// The iOS notification content used when an iOS message is converted into a notification. This parameter is valid only when iOSApnsEnv is set to PRODUCT and iOSRemind is set to true.
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

func (s *PushShrinkRequest) GetAndroidVivoTemplateId() *string {
	return s.AndroidVivoTemplateId
}

func (s *PushShrinkRequest) GetAndroidVivoTemplateParams() *string {
	return s.AndroidVivoTemplateParams
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

func (s *PushShrinkRequest) SetAndroidVivoTemplateId(v string) *PushShrinkRequest {
	s.AndroidVivoTemplateId = &v
	return s
}

func (s *PushShrinkRequest) SetAndroidVivoTemplateParams(v string) *PushShrinkRequest {
	s.AndroidVivoTemplateParams = &v
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
