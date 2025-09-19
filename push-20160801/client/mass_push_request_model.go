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
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey          *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	IdempotentToken *string `json:"IdempotentToken,omitempty" xml:"IdempotentToken,omitempty"`
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
	return dara.Validate(s)
}

type MassPushRequestPushTask struct {
	// example:
	//
	// com.alibaba.cloudpushdemo.bizactivity
	AndroidActivity *string `json:"AndroidActivity,omitempty" xml:"AndroidActivity,omitempty"`
	// example:
	//
	// 99
	AndroidBadgeAddNum *int32  `json:"AndroidBadgeAddNum,omitempty" xml:"AndroidBadgeAddNum,omitempty"`
	AndroidBadgeClass  *string `json:"AndroidBadgeClass,omitempty" xml:"AndroidBadgeClass,omitempty"`
	// example:
	//
	// 99
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
	// 1
	AndroidHuaweiTargetUserType *int32 `json:"AndroidHuaweiTargetUserType,omitempty" xml:"AndroidHuaweiTargetUserType,omitempty"`
	// example:
	//
	// https://imag.example.com/image.png
	AndroidImageUrl  *string `json:"AndroidImageUrl,omitempty" xml:"AndroidImageUrl,omitempty"`
	AndroidInboxBody *string `json:"AndroidInboxBody,omitempty" xml:"AndroidInboxBody,omitempty"`
	// if can be null:
	// false
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
	// 0
	AndroidNotificationVivoChannel *string `json:"AndroidNotificationVivoChannel,omitempty" xml:"AndroidNotificationVivoChannel,omitempty"`
	// example:
	//
	// michannel
	AndroidNotificationXiaomiChannel *string `json:"AndroidNotificationXiaomiChannel,omitempty" xml:"AndroidNotificationXiaomiChannel,omitempty"`
	// example:
	//
	// VIBRATE
	AndroidNotifyType *string `json:"AndroidNotifyType,omitempty" xml:"AndroidNotifyType,omitempty"`
	// example:
	//
	// APPLICATION
	AndroidOpenType *string `json:"AndroidOpenType,omitempty" xml:"AndroidOpenType,omitempty"`
	// example:
	//
	// https://xxxx.xxx
	AndroidOpenUrl                      *string            `json:"AndroidOpenUrl,omitempty" xml:"AndroidOpenUrl,omitempty"`
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
	AndroidRenderStyle    *string `json:"AndroidRenderStyle,omitempty" xml:"AndroidRenderStyle,omitempty"`
	AndroidTargetUserType *int32  `json:"AndroidTargetUserType,omitempty" xml:"AndroidTargetUserType,omitempty"`
	// example:
	//
	// 1
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
	// true
	StoreOffline *bool `json:"StoreOffline,omitempty" xml:"StoreOffline,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DEVICE
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// deviceid1,deviceid2
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
	// ””
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

func (s *MassPushRequestPushTask) GetAndroidXiaomiImageUrl() *string {
	return s.AndroidXiaomiImageUrl
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

func (s *MassPushRequestPushTask) SetAndroidXiaomiImageUrl(v string) *MassPushRequestPushTask {
	s.AndroidXiaomiImageUrl = &v
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
