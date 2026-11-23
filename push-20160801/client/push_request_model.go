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
	SetAndroidVivoTemplateId(v string) *PushRequest
	GetAndroidVivoTemplateId() *string
	SetAndroidVivoTemplateParams(v string) *PushRequest
	GetAndroidVivoTemplateParams() *string
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
	// The activity to be opened by the notification.
	//
	// This parameter is passed only when AndroidOpenType is set to ACTIVITY. Example: com.alibaba.cloudpushdemo.bizactivity.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidActivity *string `json:"AndroidActivity,omitempty" xml:"AndroidActivity,omitempty"`
	// Specifies the increment value for the badge. This value is added to the existing badge number. Valid values: 1 to 99.
	//
	// > This parameter is valid only for pushes through the Huawei or Honor vendor channel. If both AndroidBadgeAddNum and AndroidBadgeSetNum are specified, AndroidBadgeSetNum takes precedence.
	//
	// example:
	//
	// 1
	AndroidBadgeAddNum *int32 `json:"AndroidBadgeAddNum,omitempty" xml:"AndroidBadgeAddNum,omitempty"`
	// The fully qualified class name of the application entry activity for badge settings.
	//
	// > This parameter is valid only for pushes through the Huawei or Honor vendor channel.
	//
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidBadgeClass *string `json:"AndroidBadgeClass,omitempty" xml:"AndroidBadgeClass,omitempty"`
	// Specifies the fixed number for the badge. Valid values: 0 to 99.
	//
	// > This parameter takes effect only for the Huawei and Honor channels when pushing through vendor channels. When pushing through the Alibaba Cloud proprietary channel, this parameter takes effect only on Huawei, Honor, and vivo devices.
	//
	// example:
	//
	// 5
	AndroidBadgeSetNum *int32 `json:"AndroidBadgeSetNum,omitempty" xml:"AndroidBadgeSetNum,omitempty"`
	// The body in long text mode. Maximum length: 1,000 bytes (1 Chinese character counts as 3 bytes). The actual limit is subject to the specific vendor channel.
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
	// >If this parameter is not provided in long text mode, the first non-empty value from Body and AndroidPopupBody is used.
	//
	// example:
	//
	// Sample long text
	AndroidBigBody *string `json:"AndroidBigBody,omitempty" xml:"AndroidBigBody,omitempty"`
	// The image URL in big picture mode. Currently supported: Proprietary channel (Android SDK 3.6.0 and later).
	//
	// example:
	//
	// https://imag.example.com/image.png
	AndroidBigPictureUrl *string `json:"AndroidBigPictureUrl,omitempty" xml:"AndroidBigPictureUrl,omitempty"`
	// The title in long text mode. Maximum length: 200 bytes (1 Chinese character counts as 3 bytes).
	//
	// - Currently, only the Honor channel and the Huawei channel (EMUI 11 and later) are supported.
	//
	// - If this parameter is not provided in long text mode, the first non-empty value from Title and AndroidPopupTitle is used.
	//
	// example:
	//
	// Sample long title
	AndroidBigTitle *string `json:"AndroidBigTitle,omitempty" xml:"AndroidBigTitle,omitempty"`
	// The extended properties of the notification. This property does not take effect when the push type PushType is set to MESSAGE.
	//
	// This parameter must be passed in JSON map format. Otherwise, a parsing error occurs.
	//
	// example:
	//
	// {"key1":"value1","api_name":"PushNoticeToAndroidRequest"}
	AndroidExtParameters *string `json:"AndroidExtParameters,omitempty" xml:"AndroidExtParameters,omitempty"`
	// The notification type for the Honor channel. Valid values:
	//
	// - **0**: formal notification (default)
	//
	// - **1**: test notification
	//
	// > Each app can send 1,000 test notifications per day, which is not subject to the daily push limit per device.
	//
	// example:
	//
	// 0
	AndroidHonorTargetUserType *int32 `json:"AndroidHonorTargetUserType,omitempty" xml:"AndroidHonorTargetUserType,omitempty"`
	// The Huawei quick notification parameter. Valid values:
	//
	// - **0**: sends a standard Huawei notification (default)
	//
	// - **1**: sends a Huawei quick notification
	//
	// example:
	//
	// 1
	AndroidHuaweiBusinessType *int32 `json:"AndroidHuaweiBusinessType,omitempty" xml:"AndroidHuaweiBusinessType,omitempty"`
	// The JSON string of the Huawei Android Live Notification data structure [LiveNotificationPayload](https://developer.huawei.com/consumer/cn/doc/HMSCore-References/rest-live-0000001562939968#ZH-CN_TOPIC_0000001700850537__p195121620102511). For development integration, refer to the [Huawei Live Notification push guide](https://help.aliyun.com/document_detail/2983768.html).
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
	// > If the default receipt configuration on the Huawei push operations platform is set to the Alibaba Cloud receipt, you do not need to provide this parameter. Otherwise, configure the default receipt ID for the Huawei channel in the Alibaba Cloud EMAS shift Push console first.
	//
	// example:
	//
	// RCP4C123456
	AndroidHuaweiReceiptId *string `json:"AndroidHuaweiReceiptId,omitempty" xml:"AndroidHuaweiReceiptId,omitempty"`
	// The notification type for the Huawei channel. Valid values:
	//
	// - **0**: formal notification (default)
	//
	// - **1**: test notification
	//
	// > Each app can send 500 test notifications per day, which is not subject to the daily push limit per device.
	//
	// example:
	//
	// 0
	AndroidHuaweiTargetUserType *int32 `json:"AndroidHuaweiTargetUserType,omitempty" xml:"AndroidHuaweiTargetUserType,omitempty"`
	// The URL of the right-side icon.
	//
	// Currently supported:
	//
	// - Huawei EMUI (applicable only in long text mode and inbox mode).
	//
	// - Honor Magic UI (applicable only in long text mode).
	//
	// - Proprietary channel: Android SDK 3.5.0 and later.
	//
	// example:
	//
	// https://imag.example.com/image.png
	AndroidImageUrl *string `json:"AndroidImageUrl,omitempty" xml:"AndroidImageUrl,omitempty"`
	// The body in inbox mode. The content must be a valid JSON array with no more than 5 elements. Currently supported:
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
	// ["First line","Second line"]
	AndroidInboxBody *string `json:"AndroidInboxBody,omitempty" xml:"AndroidInboxBody,omitempty"`
	// The Meizu message type. Valid values:
	//
	// - 0: public message. This is the default value.
	//
	// - 1: private message.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 0
	AndroidMeizuNoticeMsgType *int32 `json:"AndroidMeizuNoticeMsgType,omitempty" xml:"AndroidMeizuNoticeMsgType,omitempty"`
	// Purpose 1: After you complete the [self-classification permission application](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835?#section3410731125514), this parameter is used to identify the message type and determine the [message reminder method](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#ZH-CN_TOPIC_0000001149358835__p3850133955718). It accelerates the delivery of specific types of messages. For valid values, refer to the [message classification standards](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section1076611477914) in the official Huawei push documentation, and fill in the "Cloud notification category value" or "Local notification category value" from the documentation table.
	//
	// Purpose 2: After you [apply for special permissions](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/faq-0000001050042183#section037425218509), this parameter is used to identify high-priority pass-through scenarios. Valid values:
	//
	// - VOIP: audio and video calls
	//
	// - PLAY_VOICE: voice broadcast
	//
	// > For messages whose "Cloud notification category value" is "Not applicable", the Alibaba Cloud proprietary channel is used. For messages whose "Local notification category value" is "Not applicable", the Huawei channel is used.
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
	// You must apply for the required permissions. For more information, see [Application link](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/faq-0000001050042183#section037425218509).
	//
	// example:
	//
	// HIGH
	AndroidMessageHuaweiUrgency *string `json:"AndroidMessageHuaweiUrgency,omitempty" xml:"AndroidMessageHuaweiUrgency,omitempty"`
	// OPPO message categorization: Communication and Service, and Content and Marketing.
	//
	// Communication and Service (requires permission application):
	//
	// - IM: instant chat, audio calls, and video calls
	//
	// - ACCOUNT: individual account and asset changes
	//
	// - DEVICE_REMINDER: personal device reminders
	//
	// - ORDER: personal order and logistics status changes
	//
	// - TODO: personal schedules and to-dos
	//
	// - SUBSCRIPTION: personal subscribe
	//
	// Content and Marketing:
	//
	// - NEWS: news and information
	//
	// - CONTENT: content recommendations
	//
	// - MARKETING: platform activities
	//
	// - SOCIAL: social updates
	//
	// For more information, see [OPUSH message classification rules](https://open.oppomobile.com/new/developmentDoc/info?id=13189).
	//
	// example:
	//
	// MARKETING
	AndroidMessageOppoCategory *string `json:"AndroidMessageOppoCategory,omitempty" xml:"AndroidMessageOppoCategory,omitempty"`
	// The reminder level of OPPO channel notification bar messages. Valid values:
	//
	// - 1: notification bar
	//
	// - 2: notification bar, lock screen, ringtone, and vibrate (default notification level for Communication and Service messages)
	//
	// - 16: notification bar, lock screen, ringtone, vibrate, and banner (requires permission application)
	//
	// > When you use the AndroidMessageOppoNotifyLevel parameter, you must also pass the AndroidMessageOppoCategory parameter.
	//
	// example:
	//
	// 1
	AndroidMessageOppoNotifyLevel *int32 `json:"AndroidMessageOppoNotifyLevel,omitempty" xml:"AndroidMessageOppoNotifyLevel,omitempty"`
	// vivo message categorization: System messages and Operation messages.
	//
	// System messages:
	//
	// - IM: instant messages
	//
	// - ACCOUNT: accounts and assets
	//
	// - TODO: schedules and to-dos
	//
	// - DEVICE_REMINDER: device information
	//
	// - ORDER: orders and logistics
	//
	// - SUBSCRIPTION: subscribe reminders
	//
	// Operation messages:
	//
	// - NEWS: news
	//
	// - CONTENT: content recommendations
	//
	// - MARKETING: operational activity
	//
	// - SOCIAL: social updates
	//
	// For more information, see [Categorization description](https://dev.vivo.com.cn/documentCenter/doc/359#s-ef3qugc3).
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
	// The priority of the display position of the Android notification in the notification bar. Valid values: -2, -1, 0, 1, and 2.
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
	// The channel ID of the Android app, which must correspond to the channel ID in the app.
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
	// The message group. For messages in the same group, only the latest message and the total number of messages received in the group are displayed in the notification bar. All messages are not displayed and cannot be expanded. Currently supported:
	//
	// - Huawei vendor channel
	//
	// - Honor vendor channel
	//
	// - Proprietary channel: Android SDK 3.9.1 and earlier
	//
	// > The proprietary channel no longer supports this parameter in Android SDK 3.9.2 and later.
	//
	// example:
	//
	// group-1
	AndroidNotificationGroup *string `json:"AndroidNotificationGroup,omitempty" xml:"AndroidNotificationGroup,omitempty"`
	// The importance parameter for Honor notification message classification, which determines the notification behavior on user devices. Valid values:
	//
	// - **LOW**: information and marketing messages
	//
	// - **NORMAL**: service and communication messages
	//
	// You must apply for this permission on the Honor platform. [Application link](https://developer.honor.com/cn/docs/11002/guides/notification-class#%E8%87%AA%E5%88%86%E7%B1%BB%E6%9D%83%E7%9B%8A%E7%94%B3%E8%AF%B7).
	//
	// example:
	//
	// LOW
	AndroidNotificationHonorChannel *string `json:"AndroidNotificationHonorChannel,omitempty" xml:"AndroidNotificationHonorChannel,omitempty"`
	// The importance parameter for Huawei notification message classification, which determines the notification behavior on user devices. Valid values:
	//
	// - LOW: information and marketing messages
	//
	// - NORMAL: service and communication messages
	//
	// >- Currently, we recommend that you use AndroidMessageHuaweiCategory for notification classification in the Huawei channel instead of AndroidNotificationHuaweiChannel.
	//
	// >- You must apply for this permission on the Huawei platform. [Application link](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section893184112272).
	//
	// example:
	//
	// LOW
	AndroidNotificationHuaweiChannel *string `json:"AndroidNotificationHuaweiChannel,omitempty" xml:"AndroidNotificationHuaweiChannel,omitempty"`
	// The unique identifier of each message when displayed in the notification. Different notification bar messages can have the same NotifyId, which allows new notification bar messages to overwrite old ones.
	//
	// example:
	//
	// 100001
	AndroidNotificationNotifyId *int32 `json:"AndroidNotificationNotifyId,omitempty" xml:"AndroidNotificationNotifyId,omitempty"`
	// The message thread. Messages in the same thread are collapsed and displayed in the notification bar, and can be expanded. Notifications in different threads are displayed separately. Currently supported:
	//
	// - Proprietary channel: Android SDK 3.9.2 and later
	//
	// example:
	//
	// thread-1
	AndroidNotificationThreadId *string `json:"AndroidNotificationThreadId,omitempty" xml:"AndroidNotificationThreadId,omitempty"`
	// The vivo notification message classification. Valid values:
	//
	// - 0: operation messages (default)
	//
	// - 1: system messages
	//
	// >- Currently, we recommend that you use AndroidMessageVivoCategory for notification classification in the vivo channel instead of AndroidNotificationVivoChannel.
	//
	// >- You must apply for this permission on the vivo platform. For more information, see [Application link](https://dev.vivo.com.cn/documentCenter/doc/359).
	//
	// example:
	//
	// classification
	AndroidNotificationVivoChannel *string `json:"AndroidNotificationVivoChannel,omitempty" xml:"AndroidNotificationVivoChannel,omitempty"`
	// The channel ID for Xiaomi notification types. You must apply for this permission on the Xiaomi platform. For more information, see [Application link](https://dev.mi.com/console/doc/detail?pId=2422#_4).
	//
	// >- You can apply for a maximum of 8 channels for a single app in the Xiaomi channel. Plan ahead.
	//
	// example:
	//
	// michannel
	AndroidNotificationXiaomiChannel *string `json:"AndroidNotificationXiaomiChannel,omitempty" xml:"AndroidNotificationXiaomiChannel,omitempty"`
	// The reminder method of the notification. Valid values:
	//
	// - **VIBRATE**: vibrate (default)
	//
	// - **SOUND**: sound
	//
	// - **BOTH**: sound and vibrate
	//
	// - **NONE**: silent
	//
	// example:
	//
	// BOTH
	AndroidNotifyType *string `json:"AndroidNotifyType,omitempty" xml:"AndroidNotifyType,omitempty"`
	// The action performed after the notification is clicked. Valid values:
	//
	// - **APPLICATION**: opens the app (default)
	//
	// - **ACTIVITY**: opens the specified Android activity
	//
	// - **URL**: opens the specified URL
	//
	// - **NONE**: no redirection
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
	// The JSON character string of the OPPO Fluid Cloud intent delete data structure [data](https://open.oppomobile.com/documentation/page/info?id=13578). This parameter is invalid if the AndroidOppoIntelligentIntent parameter is already specified. For development integration, see the References [OPPO Fluid Cloud push guide](https://help.aliyun.com/document_detail/2997310.html).
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
	// The JSON character string of the OPPO Fluid Cloud intent sharing data structure [IntelligentIntent](https://open.oppomobile.com/documentation/page/info?id=13565). For development integration, see the References [OPPO Fluid Cloud push guide](https://help.aliyun.com/document_detail/2997310.html).
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
	// - **0**: formal environment (default)
	//
	// - **1**: test environment
	//
	// > For the OPPO Fluid Cloud test environment, refer to [Environment setup](https://open.oppomobile.com/documentation/page/info?id=13590) to set up the client.
	//
	// example:
	//
	// 1
	AndroidOppoIntentEnv *int32 `json:"AndroidOppoIntentEnv,omitempty" xml:"AndroidOppoIntentEnv,omitempty"`
	// The content parameters for the OPPO private message template.
	AndroidOppoPrivateContentParameters map[string]*string `json:"AndroidOppoPrivateContentParameters,omitempty" xml:"AndroidOppoPrivateContentParameters,omitempty"`
	// The OPPO private message template ID.
	//
	// example:
	//
	// 687557242b1634hzefs3d5013
	AndroidOppoPrivateMsgTemplateId *string `json:"AndroidOppoPrivateMsgTemplateId,omitempty" xml:"AndroidOppoPrivateMsgTemplateId,omitempty"`
	// The title parameters for the OPPO private message template.
	AndroidOppoPrivateTitleParameters map[string]*string `json:"AndroidOppoPrivateTitleParameters,omitempty" xml:"AndroidOppoPrivateTitleParameters,omitempty"`
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
	// If the message is successfully converted into a notification and pushed, the data displayed in the received notification is the values of the AndroidPopupTitle and AndroidPopupBody parameters set on the server. When the notification is clicked, the data obtained in the onSysNoticeOpened method of the auxiliary pop-up window is the values of the Title and Body parameters set on the server.
	//
	// example:
	//
	// true
	AndroidRemind *bool `json:"AndroidRemind,omitempty" xml:"AndroidRemind,omitempty"`
	// The notification style. Valid values:
	//
	// - **0**: standard mode (default)
	//
	// - **1**: long text mode (supported by Huawei, Honor, Xiaomi, OPPO, Meizu, and the proprietary channel)
	//
	// - **2**: big picture mode (supported by the proprietary channel, but not supported by Xiaomi devices)
	//
	// - **3**: inbox mode (supported by Huawei, Honor, Xiaomi, OPPO, and the proprietary channel)
	//
	// > If a non-standard mode is used, this parameter must be provided.
	//
	// example:
	//
	// 1
	AndroidRenderStyle *int32 `json:"AndroidRenderStyle,omitempty" xml:"AndroidRenderStyle,omitempty"`
	// Settings for the vendor channel notification type. Valid values:
	//
	// - **0**: formal notification (default)
	//
	// - **1**: test notification
	//
	// >- These parameter settings are equivalent to configuring the AndroidHuaweiTargetUserType, AndroidHonorTargetUserType, AndroidVivoPushMode, and AndroidOppoIntentEnv parameters at the same time. The corresponding parameters for specific vendor channels can overwrite this parameter.
	//
	// >- Currently supported: Huawei channel, Honor channel, vivo channel, and OPPO Fluid Cloud.
	//
	// example:
	//
	// 0
	AndroidTargetUserType *int32 `json:"AndroidTargetUserType,omitempty" xml:"AndroidTargetUserType,omitempty"`
	// The JSON character string of the vivo Atomic Island data structure [liveMessage](https://dev.vivo.com.cn/documentCenter/doc/896#s-fdagzbd4). For development integration, see the References [vivo Atomic Island push guide](https://www.alibabacloud.com/help/en/document_detail/3030718.html).
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
	// - **0**: formal push (default)
	//
	// - **1**: test push
	//
	// > For test pushes, configure test devices in the vivo console in advance. You can search for "onReceiveRegId regId" in the device startup log to obtain the RegId of the test device.
	//
	// example:
	//
	// 0
	AndroidVivoPushMode *int32 `json:"AndroidVivoPushMode,omitempty" xml:"AndroidVivoPushMode,omitempty"`
	// The receipt ID for the vivo channel. You can view this receipt ID in the application information of the push service on the vivo Open Platform.
	//
	// > If the default receipt configuration on the vivo Open Platform is set to the Alibaba Cloud receipt, you do not need to provide this parameter. Otherwise, configure the default receipt ID for the vivo channel in the Alibaba Cloud EMAS Mobile Push console first.
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
	// **[Deprecated]*	- This parameter is deprecated. Starting from August 2023, Xiaomi no longer officially supports dynamically setting small icons, right-side icons, and big pictures during pushes on new devices or systems.
	//
	// example:
	//
	// https://f6.market.xiaomi.com/download/MiPass/aaa/bbb.png
	AndroidXiaomiBigPictureUrl *string `json:"AndroidXiaomiBigPictureUrl,omitempty" xml:"AndroidXiaomiBigPictureUrl,omitempty"`
	// The JSON character string of the Xiaomi Hyper Island data structure [miui.focus.param](https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131). For development integration, see the References [Xiaomi Hyper Island push guide](https://www.alibabacloud.com/help/en/document_detail/3037956.html).
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
	// The JSON character string of the Xiaomi Hyper Island data images [miui.focus.pic_xxx](https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131). For development integration, see the References [Xiaomi Hyper Island push guide](https://www.alibabacloud.com/help/en/document_detail/3037956.html).
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
	// **[Deprecated]*	- This parameter is deprecated. Starting from August 2023, Xiaomi no longer officially supports dynamically setting small icons, right-side icons, and big pictures during pushes on new devices or systems.
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
	// The content of the notification or message for Android and HarmonyOS pushes, or the content of the iOS message or notification. The content size is limited. For more information, see [Product limits](https://help.aliyun.com/document_detail/434629.html).
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
	// The expiration time of offline messages or notifications. This parameter is used together with StoreOffline. Expired messages or notifications are not delivered. The maximum retention period is 72 hours. Default value: 72 hours.
	//
	// The time must be in the ISO 8601 standard format and use UTC time. Format: YYYY-MM-DDThh:mm:ssZ. The expiration time must be later than the current time or the scheduled delivery time plus 3 seconds (ExpireTime > PushTime + 3 seconds). The 3-second buffer is used to compensate for errors caused by network and system latency. We recommend that you set the expiration time to at least 1 minute for single pushes and at least 10 minutes for full pushes and batch pushes.
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
	// The action to perform after the notification is clicked. Valid values:
	//
	// - APP_HOME_PAGE: opens the application home page.
	//
	// - APP_CUSTOM_PAGE: opens a custom page in the application.
	//
	// example:
	//
	// APP_HOME_PAGE
	HarmonyActionType *string `json:"HarmonyActionType,omitempty" xml:"HarmonyActionType,omitempty"`
	// The increment number for the HarmonyOS application badge. For more information, see [HarmonyOS badge addNum field description](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section266310382145).</br>This parameter is supported in HarmonyOS SDK 1.2.0 and later.
	//
	// example:
	//
	// 1
	HarmonyBadgeAddNum *int32 `json:"HarmonyBadgeAddNum,omitempty" xml:"HarmonyBadgeAddNum,omitempty"`
	// The fixed number for the HarmonyOS application badge. For more information, see [HarmonyOS badge setNum field description](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section266310382145). This parameter is supported in HarmonyOS SDK 1.2.0 and later.
	//
	// example:
	//
	// 1
	HarmonyBadgeSetNum *int32 `json:"HarmonyBadgeSetNum,omitempty" xml:"HarmonyBadgeSetNum,omitempty"`
	// The category of the notification message. After you apply for the notification message auto-categorization permission, this parameter is used as the identity for the message type. Different notification message types affect the display and reminder methods of the message. Valid values:
	//
	// - IM: instant chat.
	//
	// - VOIP: audio and video calls.
	//
	// - SUBSCRIPTION: subscribe.
	//
	// - TRAVEL: travel.
	//
	// - HEALTH: health.
	//
	// - WORK: work reminders.
	//
	// - ACCOUNT: account updates.
	//
	// - EXPRESS: orders and logistics.
	//
	// - FINANCE: finance.
	//
	// - DEVICE_REMINDER: device reminders.
	//
	// - MAIL: emails.
	//
	// - CUSTOMER_SERVICE: customer service messages.
	//
	// - MARKETING: news, content recommendations, social updates, product promotions, financial updates, lifestyle information, surveys, feature recommendations, and operational activity. This category only serves as the identity for the content and does not accelerate message send. These are collectively referred to as information and marketing messages.
	//
	// - PLAY_VOICE: voice broadcast. This value can only be used for HarmonyOS extension notification messages.
	//
	// For details, see [Notification.category](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117) on the HarmonyOS official website.
	//
	// example:
	//
	// IM
	HarmonyCategory *string `json:"HarmonyCategory,omitempty" xml:"HarmonyCategory,omitempty"`
	// The extended properties of the notification. This property does not take effect when the push type PushType is set to MESSAGE.
	//
	// This parameter must be passed in JSON map format. Otherwise, a parsing error occurs.
	//
	// example:
	//
	// {"key1":"value1","api_name":"PushNoticeToAndroidRequest"}
	HarmonyExtParameters *string `json:"HarmonyExtParameters,omitempty" xml:"HarmonyExtParameters,omitempty"`
	// The extra data for the extended notification message.</br>This parameter is valid when sending HarmonyOS extended notification messages.</br>This parameter is conceptually equivalent to the extraData field of HarmonyOS extended notification messages. For specific definitions, see [HarmonyOS ExtensionPayload description](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section161192514234).</br>This parameter is supported in HarmonyOS SDK 1.2.0 and later.
	//
	// example:
	//
	// Sample extra data
	HarmonyExtensionExtraData *string `json:"HarmonyExtensionExtraData,omitempty" xml:"HarmonyExtensionExtraData,omitempty"`
	// Specifies whether the message is a HarmonyOS extended notification message when PushType is set to NOTICE. Valid values:
	//
	// - true: sends an extended notification message.
	//
	// - false: sends a normal notification. This is the default value.
	//
	// You must apply for permissions on the HarmonyOS side before you can send extended notification messages. For more information, see [Send extended notification messages](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V5/push-send-extend-noti-V5) in the HarmonyOS documentation.</br>This parameter is supported in HarmonyOS SDK 1.2.0 and later.
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
	// The JSON string of the HarmonyOS Live View data structure [LiveViewPayload](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V13/push-scenariozed-api-request-param-V13#section66881469306). For development integration, refer to the [HarmonyOS Live View push guide](https://help.aliyun.com/document_detail/2982112.html).
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
	// Specifies the type of the notification channel to use. This parameter is valid only when the Alibaba Cloud proprietary channel is online. Valid values:
	//
	// - SOCIAL_COMMUNICATION: social communication.
	//
	// - SERVICE_INFORMATION: service reminders.
	//
	// - CONTENT_INFORMATION: content information.
	//
	// - CUSTOMER_SERVICE: customer service messages. This type is used for customer service messages between users and merchants, and must be actively initiated by the user.
	//
	// - OTHER_TYPES: other types.
	//
	// For more information, see [SlotType](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/js-apis-notificationmanager-V5#slottype) on the HarmonyOS official website.
	//
	// example:
	//
	// SOCIAL_COMMUNICATION
	HarmonyNotificationSlotType *string `json:"HarmonyNotificationSlotType,omitempty" xml:"HarmonyNotificationSlotType,omitempty"`
	// The unique identifier of each message when the notification is displayed. If this parameter is not specified, the push service automatically generates a unique identifier for each message. Different notification messages can share the same notifyId to allow new messages to overwrite old ones.
	//
	// For more information, see [Notification.notifyId](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section17371529101117) on the HarmonyOS official website.
	//
	// example:
	//
	// 0
	HarmonyNotifyId *int32 `json:"HarmonyNotifyId,omitempty" xml:"HarmonyNotifyId,omitempty"`
	// The receipt ID for the HarmonyOS channel. You can view this receipt ID in the receipt parameter settings of the HarmonyOS push operations platform.
	//
	// > If the default receipt configuration on the HarmonyOS push operations platform is set to the Alibaba Cloud receipt, you do not need to provide this parameter. Otherwise, configure the default receipt ID for the HarmonyOS channel in the Alibaba Cloud EMAS shift Push console first.
	//
	// For details, see [pushOptions.receiptId](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section418321011212) on the HarmonyOS official website.
	//
	// example:
	//
	// RCPB***DFD5
	HarmonyReceiptId *string `json:"HarmonyReceiptId,omitempty" xml:"HarmonyReceiptId,omitempty"`
	// Specifies whether to use the auxiliary pop-up window feature when the device is offline during a message push. Default value: false. This parameter takes effect only when PushType is set to MESSAGE.
	//
	// If the message-to-notification push is successful, the notification displays the data specified by the HarmonyRemindTitle and HarmonyRemindBody parameters on the server.
	//
	// example:
	//
	// false
	HarmonyRemind *bool `json:"HarmonyRemind,omitempty" xml:"HarmonyRemind,omitempty"`
	// The HarmonyOS notification content used when a HarmonyOS message is converted to a notification. This parameter is valid only when HarmonyRemind is set to true.
	//
	// example:
	//
	// You have a new message. Please check it
	HarmonyRemindBody *string `json:"HarmonyRemindBody,omitempty" xml:"HarmonyRemindBody,omitempty"`
	// The HarmonyOS notification title used when a HarmonyOS message is converted to a notification. This parameter is valid only when HarmonyRemind is set to true.
	//
	// example:
	//
	// New message
	HarmonyRemindTitle *string `json:"HarmonyRemindTitle,omitempty" xml:"HarmonyRemindTitle,omitempty"`
	// The style of the notification message. Valid values:
	//
	// - NORMAL: normal notification. This is the default value.
	//
	// - MULTI_LINE: multi-line text style.
	//
	// example:
	//
	// NORMAL
	HarmonyRenderStyle *string `json:"HarmonyRenderStyle,omitempty" xml:"HarmonyRenderStyle,omitempty"`
	// Specifies whether the message is a test message. Valid values:
	//
	// - false: normal message. This is the default value.
	//
	// - true: test message.
	//
	// For more information, see [pushOptions.testMessage](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section418321011212) on the HarmonyOS official website.
	//
	// example:
	//
	// true
	HarmonyTestMessage *bool `json:"HarmonyTestMessage,omitempty" xml:"HarmonyTestMessage,omitempty"`
	// The URI corresponding to the ability of the built-in page in the application.
	//
	// 	Notice: When HarmonyActionType is set to APP_CUSTOM_PAGE, you must specify at least one of the HarmonyUri and HarmonyAction fields. If multiple abilities exist, specify the action and URI for each ability. The action is prioritized when searching for the corresponding built-in page in the application.
	//
	// For more information, see [ClickAction.uri](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-param-V5#section152462191216) on the HarmonyOS official website.
	//
	// example:
	//
	// https://www.example.com:8080/push/example
	HarmonyUri *string `json:"HarmonyUri,omitempty" xml:"HarmonyUri,omitempty"`
	// An idempotent parameter used to prevent duplicate pushes on the server caused by retries from the API client. If the same IdempotentToken is used for calls within 15 minutes, only one push is performed, and subsequent calls return the result of the first successful push.
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
	// The custom identifier of the push task. If JobKey is not empty, this field is included in the receipt log. For more information about receipt logs, see [Receipt logs](https://help.aliyun.com/document_detail/434651.html).
	//
	// >Format requirements: Only letters, digits, underscores (_), and hyphens (-) are allowed. The length cannot exceed 32 characters.
	//
	// example:
	//
	// 123
	JobKey *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	// The scheduled delivery time. If this parameter is not set, the push is delivered immediately.
	//
	// The scheduled delivery time can be at most 7 days later than the current time.
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
	// - **NOTICE**: notification. Notifications are delivered to device endpoints through vendor channels such as APNs, Huawei, Xiaomi, and HarmonyOS, and are directly displayed in the notification bar of the device. When an Android device is online (the app process is alive), the notification is preferentially delivered through the Alibaba Cloud proprietary channel. The push SDK constructs and displays the notification to achieve better push performance and save vendor push message quotas in certain scenarios.
	//
	// - **MESSAGE**: message. Messages are delivered through the Alibaba Cloud proprietary online channel. They are not displayed in the notification bar by default. Instead, the app must receive and process them when the process is active, and the business logic determines whether to trigger specific business behaviors. When the device is offline (the app process is inactive), messages cannot be received in a timely manner. In this case, you can use the iOSRemind or AndroidRemind parameter described below to convert messages into notifications for delivery when the device is offline. Alternatively, you can set the StoreOffline parameter described below to allow the push system to save the message when the device is offline and automatically deliver it when the device comes online.
	//
	// This parameter is required.
	//
	// example:
	//
	// MESSAGE
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	// The specified delivery channels. Valid values:
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
	// >- If this parameter is not configured, all channels can be used.
	//
	// >- If this parameter is configured, only the specified channels are strictly used.
	//
	// >- If the configured channels conflict with the delivery policy (for example, iOS notifications are only delivered through the APNs channel, but apns is not included in this parameter), the push is not delivered.
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
	// This parameter is required if you use the SMS linkage feature. Set this parameter to 15 seconds or more, up to a maximum of 3 days, to prevent duplicate SMS messages and push notifications.
	//
	// > When the SMS linkage feature is used, the ExpireTime parameter becomes invalid. The expiration time of the notification is calculated based on the SmsDelaySecs parameter, which is the current time plus the value of SmsDelaySecs.
	//
	// example:
	//
	// 15
	SmsDelaySecs *int32 `json:"SmsDelaySecs,omitempty" xml:"SmsDelaySecs,omitempty"`
	// The variable name-value pairs for the SMS template. Format: `key1=value1&key2=value2`.
	//
	// example:
	//
	// key1=value1
	SmsParams *string `json:"SmsParams,omitempty" xml:"SmsParams,omitempty"`
	// The condition that triggers the SMS message. Valid values:
	//
	// - **0**: triggers the SMS message when the push notification is not received.
	//
	// - **1**: triggers the SMS message when the user does not open the notification.
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
	// The name of the SMS template used for fallback SMS messages. You can obtain this name from the SMS template management page. This is the name assigned by the system, not the name set by the developer.
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
	// - **DEVICE**: pushes by device.
	//
	// - **ACCOUNT**: pushes by account.
	//
	// - **ALIAS**: pushes by alias.
	//
	// - **TAG**: pushes by tag.
	//
	// - **ALL**: pushes to all devices. The interval between two full pushes of the same DeviceType must be at least 1 second.
	//
	//  > A full push to iOS devices is sent to devices that have been active within the last 24 months and have not been uninstalled. Once Apple Push Notification service (APNs) receives the push request and does not return an error message, the push is considered delivered. This may cause a surge in the number of active devices and incur high costs. Use this value with caution.
	//
	// - **TBD**: initializes a continuous push. The push target is specified by the subsequent [ContinuouslyPush](https://help.aliyun.com/document_detail/2249917.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The target value, which is set based on the Target type. Separate multiple values with commas (,). If the number of values exceeds the limit, you must send the push in multiple batches.
	//
	// - If Target is set to DEVICE, the value is in the format of deviceid1,deviceid2. A maximum of 1,000 devices are supported.
	//
	// - If Target is set to ACCOUNT, the value is in the format of account1,account2. A maximum of 1,000 accounts are supported.
	//
	// - If Target is set to ALIAS, the value is in the format of alias1,alias2. A maximum of 1,000 aliases are supported.
	//
	// - If Target is set to TAG, single-tag and multi-tag pushes are supported. For more information about the format, see [Tag format](https://help.aliyun.com/document_detail/434847.html).
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
	// - For iOS 8.2 to iOS 10 (excluding iOS 10), the title replaces the app name in the notification.
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
	// - **DEV**: the development environment, which applies to apps directly installed and debugged by using Xcode.
	//
	// - **PRODUCT**: the production environment, which applies to apps distributed through the App Store, TestFlight, Ad Hoc, and enterprise distribution.
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
	// For iOS 10 and later, you can specify the resource URL for a rich media push notification in this parameter: `{"attachment": "https://www.example.com/notification_pic.png"}`. This parameter must be passed in JSON map format. Otherwise, a parse error occurs.
	//
	// example:
	//
	// {"attachment": "https://www.example.com/notification_pic.png"}
	IOSExtParameters *string `json:"iOSExtParameters,omitempty" xml:"iOSExtParameters,omitempty"`
	// The interruption level. Valid values:
	//
	// - **passive**: The system adds the notification to the notification list without lighting up the screen or playing a sound.
	//
	// - **active**: The system immediately displays the notification, lights up the screen, and may play a sound.
	//
	// - **time-sensitive**: The system immediately presents the notification, lights up the screen, and may play a sound, but does not break through the system notification controls.
	//
	// - **critical**: The system immediately displays the notification, lights up the screen, and plays a sound bypassing the silent switch.
	//
	// example:
	//
	// active
	IOSInterruptionLevel *string `json:"iOSInterruptionLevel,omitempty" xml:"iOSInterruptionLevel,omitempty"`
	// The JSON character string that contains the static parameters for Dynamic Island pushes. It includes static user-defined information, such as product IDs and order information.
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
	// The dynamic parameters for Dynamic Island pushes. It includes real-time update information, such as price and inventory changes.
	//
	// example:
	//
	// {"status": "delivered", "estimatedArrival": "2023-12-31T12:00:00Z"}
	IOSLiveActivityContentState *string `json:"iOSLiveActivityContentState,omitempty" xml:"iOSLiveActivityContentState,omitempty"`
	// The UNIX timestamp in seconds. The ended Live Activity remains on the Lock Screen until this specified time. The maximum retention period is 4 hours.
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
	// The sound of the iOS notification. Specify the name of the audio file stored in the app bundle or the Library/Sounds directory of the sandbox. For more information, see [How to set the notification sound for iOS pushes](https://help.aliyun.com/document_detail/48906.html).
	//
	// If this parameter is set to an empty string (""), the notification is silent. If this parameter is not set, the default value "default" is used, which indicates the system prompt sound.
	//
	// example:
	//
	// ""
	IOSMusic *string `json:"iOSMusic,omitempty" xml:"iOSMusic,omitempty"`
	// The iOS notification processing extension flag (iOS 10 and later). If this parameter is set to true, the notification pushed by APNs can be processed by the extension before it is displayed. For silent notifications, this parameter must be set to true.
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
	// The collapse ID. If a device receives messages with the same CollapseId, the messages are merged into one. If the device is offline and consecutive messages with the same CollapseId are sent, only one notification is displayed in the notification bar. This parameter is supported in iOS 10 and later.
	//
	// example:
	//
	// ZD2011
	IOSNotificationCollapseId *string `json:"iOSNotificationCollapseId,omitempty" xml:"iOSNotificationCollapseId,omitempty"`
	// The thread ID used to group iOS remote notifications and mark the identifier of the collapsed group.
	//
	// This parameter is supported only in iOS 12.0 and later.
	//
	// example:
	//
	// abc
	IOSNotificationThreadId *string `json:"iOSNotificationThreadId,omitempty" xml:"iOSNotificationThreadId,omitempty"`
	// The relevance score for summary highlighting. Valid values: floating-point numbers in the range of [0,1\\].
	//
	// example:
	//
	// 0.01
	IOSRelevanceScore *float64 `json:"iOSRelevanceScore,omitempty" xml:"iOSRelevanceScore,omitempty"`
	// Specifies whether to convert a message into a notification and deliver it through the APNs channel when the device is offline (the persistent connection channel to the Mobile Push server is disconnected) during a message push.
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

func (s *PushRequest) GetAndroidVivoTemplateId() *string {
	return s.AndroidVivoTemplateId
}

func (s *PushRequest) GetAndroidVivoTemplateParams() *string {
	return s.AndroidVivoTemplateParams
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

func (s *PushRequest) SetAndroidVivoTemplateId(v string) *PushRequest {
	s.AndroidVivoTemplateId = &v
	return s
}

func (s *PushRequest) SetAndroidVivoTemplateParams(v string) *PushRequest {
	s.AndroidVivoTemplateParams = &v
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
