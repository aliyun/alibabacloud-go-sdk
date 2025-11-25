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
	SetAndroidXiaomiImageUrl(v string) *PushRequest
	GetAndroidXiaomiImageUrl() *string
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
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidActivity    *string `json:"AndroidActivity,omitempty" xml:"AndroidActivity,omitempty"`
	AndroidBadgeAddNum *int32  `json:"AndroidBadgeAddNum,omitempty" xml:"AndroidBadgeAddNum,omitempty"`
	AndroidBadgeClass  *string `json:"AndroidBadgeClass,omitempty" xml:"AndroidBadgeClass,omitempty"`
	AndroidBadgeSetNum *int32  `json:"AndroidBadgeSetNum,omitempty" xml:"AndroidBadgeSetNum,omitempty"`
	AndroidBigBody     *string `json:"AndroidBigBody,omitempty" xml:"AndroidBigBody,omitempty"`
	// example:
	//
	// https://imag.example.com/image.png
	AndroidBigPictureUrl *string `json:"AndroidBigPictureUrl,omitempty" xml:"AndroidBigPictureUrl,omitempty"`
	AndroidBigTitle      *string `json:"AndroidBigTitle,omitempty" xml:"AndroidBigTitle,omitempty"`
	// example:
	//
	// {"key1":"value1","api_name":"PushNoticeToAndroidRequest"}
	AndroidExtParameters                 *string `json:"AndroidExtParameters,omitempty" xml:"AndroidExtParameters,omitempty"`
	AndroidHonorTargetUserType           *int32  `json:"AndroidHonorTargetUserType,omitempty" xml:"AndroidHonorTargetUserType,omitempty"`
	AndroidHuaweiLiveNotificationPayload *string `json:"AndroidHuaweiLiveNotificationPayload,omitempty" xml:"AndroidHuaweiLiveNotificationPayload,omitempty"`
	// example:
	//
	// RCP4C123456
	AndroidHuaweiReceiptId *string `json:"AndroidHuaweiReceiptId,omitempty" xml:"AndroidHuaweiReceiptId,omitempty"`
	// example:
	//
	// 0
	AndroidHuaweiTargetUserType *int32 `json:"AndroidHuaweiTargetUserType,omitempty" xml:"AndroidHuaweiTargetUserType,omitempty"`
	// example:
	//
	// https://imag.example.com/image.png
	AndroidImageUrl  *string `json:"AndroidImageUrl,omitempty" xml:"AndroidImageUrl,omitempty"`
	AndroidInboxBody *string `json:"AndroidInboxBody,omitempty" xml:"AndroidInboxBody,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 0
	AndroidMeizuNoticeMsgType *int32 `json:"AndroidMeizuNoticeMsgType,omitempty" xml:"AndroidMeizuNoticeMsgType,omitempty"`
	// example:
	//
	// VOIP
	AndroidMessageHuaweiCategory *string `json:"AndroidMessageHuaweiCategory,omitempty" xml:"AndroidMessageHuaweiCategory,omitempty"`
	// example:
	//
	// HIGH
	AndroidMessageHuaweiUrgency   *string `json:"AndroidMessageHuaweiUrgency,omitempty" xml:"AndroidMessageHuaweiUrgency,omitempty"`
	AndroidMessageOppoCategory    *string `json:"AndroidMessageOppoCategory,omitempty" xml:"AndroidMessageOppoCategory,omitempty"`
	AndroidMessageOppoNotifyLevel *int32  `json:"AndroidMessageOppoNotifyLevel,omitempty" xml:"AndroidMessageOppoNotifyLevel,omitempty"`
	// example:
	//
	// TODO
	AndroidMessageVivoCategory *string `json:"AndroidMessageVivoCategory,omitempty" xml:"AndroidMessageVivoCategory,omitempty"`
	AndroidMusic               *string `json:"AndroidMusic,omitempty" xml:"AndroidMusic,omitempty"`
	// example:
	//
	// 0
	AndroidNotificationBarPriority *int32 `json:"AndroidNotificationBarPriority,omitempty" xml:"AndroidNotificationBarPriority,omitempty"`
	// example:
	//
	// 2
	AndroidNotificationBarType *int32 `json:"AndroidNotificationBarType,omitempty" xml:"AndroidNotificationBarType,omitempty"`
	// example:
	//
	// 1
	AndroidNotificationChannel *string `json:"AndroidNotificationChannel,omitempty" xml:"AndroidNotificationChannel,omitempty"`
	// example:
	//
	// group-1
	AndroidNotificationGroup *string `json:"AndroidNotificationGroup,omitempty" xml:"AndroidNotificationGroup,omitempty"`
	// example:
	//
	// LOW
	AndroidNotificationHonorChannel *string `json:"AndroidNotificationHonorChannel,omitempty" xml:"AndroidNotificationHonorChannel,omitempty"`
	// example:
	//
	// LOW
	AndroidNotificationHuaweiChannel *string `json:"AndroidNotificationHuaweiChannel,omitempty" xml:"AndroidNotificationHuaweiChannel,omitempty"`
	// example:
	//
	// 100001
	AndroidNotificationNotifyId *int32  `json:"AndroidNotificationNotifyId,omitempty" xml:"AndroidNotificationNotifyId,omitempty"`
	AndroidNotificationThreadId *string `json:"AndroidNotificationThreadId,omitempty" xml:"AndroidNotificationThreadId,omitempty"`
	// example:
	//
	// classification
	AndroidNotificationVivoChannel *string `json:"AndroidNotificationVivoChannel,omitempty" xml:"AndroidNotificationVivoChannel,omitempty"`
	// example:
	//
	// michannel
	AndroidNotificationXiaomiChannel *string `json:"AndroidNotificationXiaomiChannel,omitempty" xml:"AndroidNotificationXiaomiChannel,omitempty"`
	// example:
	//
	// BOTH
	AndroidNotifyType *string `json:"AndroidNotifyType,omitempty" xml:"AndroidNotifyType,omitempty"`
	// example:
	//
	// APPLICATION
	AndroidOpenType *string `json:"AndroidOpenType,omitempty" xml:"AndroidOpenType,omitempty"`
	// example:
	//
	// https://xxxx.xxx
	AndroidOpenUrl                      *string            `json:"AndroidOpenUrl,omitempty" xml:"AndroidOpenUrl,omitempty"`
	AndroidOppoDeleteIntentData         *string            `json:"AndroidOppoDeleteIntentData,omitempty" xml:"AndroidOppoDeleteIntentData,omitempty"`
	AndroidOppoIntelligentIntent        *string            `json:"AndroidOppoIntelligentIntent,omitempty" xml:"AndroidOppoIntelligentIntent,omitempty"`
	AndroidOppoIntentEnv                *int32             `json:"AndroidOppoIntentEnv,omitempty" xml:"AndroidOppoIntentEnv,omitempty"`
	AndroidOppoPrivateContentParameters map[string]*string `json:"AndroidOppoPrivateContentParameters,omitempty" xml:"AndroidOppoPrivateContentParameters,omitempty"`
	AndroidOppoPrivateMsgTemplateId     *string            `json:"AndroidOppoPrivateMsgTemplateId,omitempty" xml:"AndroidOppoPrivateMsgTemplateId,omitempty"`
	AndroidOppoPrivateTitleParameters   map[string]*string `json:"AndroidOppoPrivateTitleParameters,omitempty" xml:"AndroidOppoPrivateTitleParameters,omitempty"`
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidPopupActivity *string `json:"AndroidPopupActivity,omitempty" xml:"AndroidPopupActivity,omitempty"`
	// example:
	//
	// hello
	AndroidPopupBody *string `json:"AndroidPopupBody,omitempty" xml:"AndroidPopupBody,omitempty"`
	// example:
	//
	// hello
	AndroidPopupTitle *string `json:"AndroidPopupTitle,omitempty" xml:"AndroidPopupTitle,omitempty"`
	// example:
	//
	// true
	AndroidRemind *bool `json:"AndroidRemind,omitempty" xml:"AndroidRemind,omitempty"`
	// example:
	//
	// 1
	AndroidRenderStyle    *int32 `json:"AndroidRenderStyle,omitempty" xml:"AndroidRenderStyle,omitempty"`
	AndroidTargetUserType *int32 `json:"AndroidTargetUserType,omitempty" xml:"AndroidTargetUserType,omitempty"`
	// example:
	//
	// 0
	AndroidVivoPushMode  *int32  `json:"AndroidVivoPushMode,omitempty" xml:"AndroidVivoPushMode,omitempty"`
	AndroidVivoReceiptId *string `json:"AndroidVivoReceiptId,omitempty" xml:"AndroidVivoReceiptId,omitempty"`
	// Deprecated
	AndroidXiaoMiActivity *string `json:"AndroidXiaoMiActivity,omitempty" xml:"AndroidXiaoMiActivity,omitempty"`
	// Deprecated
	AndroidXiaoMiNotifyBody *string `json:"AndroidXiaoMiNotifyBody,omitempty" xml:"AndroidXiaoMiNotifyBody,omitempty"`
	// Deprecated
	AndroidXiaoMiNotifyTitle *string `json:"AndroidXiaoMiNotifyTitle,omitempty" xml:"AndroidXiaoMiNotifyTitle,omitempty"`
	// Deprecated
	//
	// example:
	//
	// https://f6.market.xiaomi.com/download/MiPass/aaa/bbb.png
	AndroidXiaomiBigPictureUrl *string `json:"AndroidXiaomiBigPictureUrl,omitempty" xml:"AndroidXiaomiBigPictureUrl,omitempty"`
	// Deprecated
	//
	// example:
	//
	// https://imag.example.com/image.png
	AndroidXiaomiImageUrl *string `json:"AndroidXiaomiImageUrl,omitempty" xml:"AndroidXiaomiImageUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// hello
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALL
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// 2019-02-20T00:00:00Z
	ExpireTime                  *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	HarmonyAction               *string `json:"HarmonyAction,omitempty" xml:"HarmonyAction,omitempty"`
	HarmonyActionType           *string `json:"HarmonyActionType,omitempty" xml:"HarmonyActionType,omitempty"`
	HarmonyBadgeAddNum          *int32  `json:"HarmonyBadgeAddNum,omitempty" xml:"HarmonyBadgeAddNum,omitempty"`
	HarmonyBadgeSetNum          *int32  `json:"HarmonyBadgeSetNum,omitempty" xml:"HarmonyBadgeSetNum,omitempty"`
	HarmonyCategory             *string `json:"HarmonyCategory,omitempty" xml:"HarmonyCategory,omitempty"`
	HarmonyExtParameters        *string `json:"HarmonyExtParameters,omitempty" xml:"HarmonyExtParameters,omitempty"`
	HarmonyExtensionExtraData   *string `json:"HarmonyExtensionExtraData,omitempty" xml:"HarmonyExtensionExtraData,omitempty"`
	HarmonyExtensionPush        *bool   `json:"HarmonyExtensionPush,omitempty" xml:"HarmonyExtensionPush,omitempty"`
	HarmonyImageUrl             *string `json:"HarmonyImageUrl,omitempty" xml:"HarmonyImageUrl,omitempty"`
	HarmonyInboxContent         *string `json:"HarmonyInboxContent,omitempty" xml:"HarmonyInboxContent,omitempty"`
	HarmonyLiveViewPayload      *string `json:"HarmonyLiveViewPayload,omitempty" xml:"HarmonyLiveViewPayload,omitempty"`
	HarmonyNotificationSlotType *string `json:"HarmonyNotificationSlotType,omitempty" xml:"HarmonyNotificationSlotType,omitempty"`
	HarmonyNotifyId             *int32  `json:"HarmonyNotifyId,omitempty" xml:"HarmonyNotifyId,omitempty"`
	HarmonyReceiptId            *string `json:"HarmonyReceiptId,omitempty" xml:"HarmonyReceiptId,omitempty"`
	HarmonyRemind               *bool   `json:"HarmonyRemind,omitempty" xml:"HarmonyRemind,omitempty"`
	HarmonyRemindBody           *string `json:"HarmonyRemindBody,omitempty" xml:"HarmonyRemindBody,omitempty"`
	HarmonyRemindTitle          *string `json:"HarmonyRemindTitle,omitempty" xml:"HarmonyRemindTitle,omitempty"`
	HarmonyRenderStyle          *string `json:"HarmonyRenderStyle,omitempty" xml:"HarmonyRenderStyle,omitempty"`
	HarmonyTestMessage          *bool   `json:"HarmonyTestMessage,omitempty" xml:"HarmonyTestMessage,omitempty"`
	HarmonyUri                  *string `json:"HarmonyUri,omitempty" xml:"HarmonyUri,omitempty"`
	IdempotentToken             *string `json:"IdempotentToken,omitempty" xml:"IdempotentToken,omitempty"`
	// example:
	//
	// 123
	JobKey *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	// example:
	//
	// 2019-02-20T00:00:00Z
	PushTime *string `json:"PushTime,omitempty" xml:"PushTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MESSAGE
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	// example:
	//
	// accs,huawei,xiaomi
	SendChannels *string `json:"SendChannels,omitempty" xml:"SendChannels,omitempty"`
	// Deprecated
	//
	// example:
	//
	// 0
	SendSpeed *int32 `json:"SendSpeed,omitempty" xml:"SendSpeed,omitempty"`
	// example:
	//
	// 15
	SmsDelaySecs *int32 `json:"SmsDelaySecs,omitempty" xml:"SmsDelaySecs,omitempty"`
	// example:
	//
	// key1=value1
	SmsParams *string `json:"SmsParams,omitempty" xml:"SmsParams,omitempty"`
	// example:
	//
	// 0
	SmsSendPolicy   *int32  `json:"SmsSendPolicy,omitempty" xml:"SmsSendPolicy,omitempty"`
	SmsSignName     *string `json:"SmsSignName,omitempty" xml:"SmsSignName,omitempty"`
	SmsTemplateName *string `json:"SmsTemplateName,omitempty" xml:"SmsTemplateName,omitempty"`
	// example:
	//
	// false
	StoreOffline *bool `json:"StoreOffline,omitempty" xml:"StoreOffline,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALL
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALL
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// false
	Trim *bool `json:"Trim,omitempty" xml:"Trim,omitempty"`
	// example:
	//
	// DEV
	IOSApnsEnv *string `json:"iOSApnsEnv,omitempty" xml:"iOSApnsEnv,omitempty"`
	// example:
	//
	// 0
	IOSBadge *int32 `json:"iOSBadge,omitempty" xml:"iOSBadge,omitempty"`
	// example:
	//
	// true
	IOSBadgeAutoIncrement *bool `json:"iOSBadgeAutoIncrement,omitempty" xml:"iOSBadgeAutoIncrement,omitempty"`
	// example:
	//
	// {“attachment”: “https://xxxx.xxx/notification_pic.png"}
	IOSExtParameters *string `json:"iOSExtParameters,omitempty" xml:"iOSExtParameters,omitempty"`
	// example:
	//
	// active
	IOSInterruptionLevel          *string `json:"iOSInterruptionLevel,omitempty" xml:"iOSInterruptionLevel,omitempty"`
	IOSLiveActivityAttributes     *string `json:"iOSLiveActivityAttributes,omitempty" xml:"iOSLiveActivityAttributes,omitempty"`
	IOSLiveActivityAttributesType *string `json:"iOSLiveActivityAttributesType,omitempty" xml:"iOSLiveActivityAttributesType,omitempty"`
	IOSLiveActivityContentState   *string `json:"iOSLiveActivityContentState,omitempty" xml:"iOSLiveActivityContentState,omitempty"`
	IOSLiveActivityDismissalDate  *int64  `json:"iOSLiveActivityDismissalDate,omitempty" xml:"iOSLiveActivityDismissalDate,omitempty"`
	IOSLiveActivityEvent          *string `json:"iOSLiveActivityEvent,omitempty" xml:"iOSLiveActivityEvent,omitempty"`
	IOSLiveActivityId             *string `json:"iOSLiveActivityId,omitempty" xml:"iOSLiveActivityId,omitempty"`
	IOSLiveActivityStaleDate      *int64  `json:"iOSLiveActivityStaleDate,omitempty" xml:"iOSLiveActivityStaleDate,omitempty"`
	// example:
	//
	// ""
	IOSMusic *string `json:"iOSMusic,omitempty" xml:"iOSMusic,omitempty"`
	// example:
	//
	// true
	IOSMutableContent *bool `json:"iOSMutableContent,omitempty" xml:"iOSMutableContent,omitempty"`
	// example:
	//
	// ios
	IOSNotificationCategory *string `json:"iOSNotificationCategory,omitempty" xml:"iOSNotificationCategory,omitempty"`
	// example:
	//
	// ZD2011
	IOSNotificationCollapseId *string `json:"iOSNotificationCollapseId,omitempty" xml:"iOSNotificationCollapseId,omitempty"`
	// example:
	//
	// abc
	IOSNotificationThreadId *string `json:"iOSNotificationThreadId,omitempty" xml:"iOSNotificationThreadId,omitempty"`
	// example:
	//
	// 0.01
	IOSRelevanceScore *float64 `json:"iOSRelevanceScore,omitempty" xml:"iOSRelevanceScore,omitempty"`
	// example:
	//
	// true
	IOSRemind     *bool   `json:"iOSRemind,omitempty" xml:"iOSRemind,omitempty"`
	IOSRemindBody *string `json:"iOSRemindBody,omitempty" xml:"iOSRemindBody,omitempty"`
	// example:
	//
	// true
	IOSSilentNotification *bool `json:"iOSSilentNotification,omitempty" xml:"iOSSilentNotification,omitempty"`
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

func (s *PushRequest) GetAndroidXiaomiImageUrl() *string {
	return s.AndroidXiaomiImageUrl
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

func (s *PushRequest) SetAndroidXiaomiImageUrl(v string) *PushRequest {
	s.AndroidXiaomiImageUrl = &v
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
