// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErrorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v map[string]interface{}) *GetErrorResponseBody
	GetArgs() map[string]interface{}
	SetErrorCode(v int32) *GetErrorResponseBody
	GetErrorCode() *int32
	SetMessage(v string) *GetErrorResponseBody
	GetMessage() *string
	SetModel(v *GetErrorResponseBodyModel) *GetErrorResponseBody
	GetModel() *GetErrorResponseBodyModel
	SetRequestId(v string) *GetErrorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetErrorResponseBody
	GetSuccess() *bool
}

type GetErrorResponseBody struct {
	Args map[string]interface{} `json:"Args,omitempty" xml:"Args,omitempty"`
	// example:
	//
	// Success
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Successful
	Message *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *GetErrorResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// A8313212-EB4E-4E15-A7F9-D9C8F3FE8E94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetErrorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetErrorResponseBody) GoString() string {
	return s.String()
}

func (s *GetErrorResponseBody) GetArgs() map[string]interface{} {
	return s.Args
}

func (s *GetErrorResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetErrorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetErrorResponseBody) GetModel() *GetErrorResponseBodyModel {
	return s.Model
}

func (s *GetErrorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetErrorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetErrorResponseBody) SetArgs(v map[string]interface{}) *GetErrorResponseBody {
	s.Args = v
	return s
}

func (s *GetErrorResponseBody) SetErrorCode(v int32) *GetErrorResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetErrorResponseBody) SetMessage(v string) *GetErrorResponseBody {
	s.Message = &v
	return s
}

func (s *GetErrorResponseBody) SetModel(v *GetErrorResponseBodyModel) *GetErrorResponseBody {
	s.Model = v
	return s
}

func (s *GetErrorResponseBody) SetRequestId(v string) *GetErrorResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetErrorResponseBody) SetSuccess(v bool) *GetErrorResponseBody {
	s.Success = &v
	return s
}

func (s *GetErrorResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetErrorResponseBodyModel struct {
	// example:
	//
	// 4G
	Access *string `json:"Access,omitempty" xml:"Access,omitempty"`
	// example:
	//
	// null
	AccessSubType *string `json:"AccessSubType,omitempty" xml:"AccessSubType,omitempty"`
	// example:
	//
	// -
	AdditionalCallbackInfo *string `json:"AdditionalCallbackInfo,omitempty" xml:"AdditionalCallbackInfo,omitempty"`
	// example:
	//
	// -
	AdditionalCrossPlatformCrashInfo *string `json:"AdditionalCrossPlatformCrashInfo,omitempty" xml:"AdditionalCrossPlatformCrashInfo,omitempty"`
	// example:
	//
	// -
	AdditionalCustomInfo *string `json:"AdditionalCustomInfo,omitempty" xml:"AdditionalCustomInfo,omitempty"`
	// example:
	//
	// -
	AndroidVm *string `json:"AndroidVm,omitempty" xml:"AndroidVm,omitempty"`
	// example:
	//
	// ALIPUBDC3CA71131027
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// appKey
	//
	// example:
	//
	// 233588686
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// V20250224102631
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// -
	ArgHash *int64 `json:"ArgHash,omitempty" xml:"ArgHash,omitempty"`
	// example:
	//
	// h=60
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// example:
	//
	// java.lang.NullPointerException: Attempt to invoke virtual method \\"java.lang.Object
	Backtrace *string `json:"Backtrace,omitempty" xml:"Backtrace,omitempty"`
	// banner
	//
	// example:
	//
	// -
	Banner *string `json:"Banner,omitempty" xml:"Banner,omitempty"`
	// example:
	//
	// -
	BinaryUuids *string `json:"BinaryUuids,omitempty" xml:"BinaryUuids,omitempty"`
	// example:
	//
	// Hinova
	Brand *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	// example:
	//
	// Chrome
	BrowserName    *string `json:"BrowserName,omitempty" xml:"BrowserName,omitempty"`
	BrowserVersion *string `json:"BrowserVersion,omitempty" xml:"BrowserVersion,omitempty"`
	// example:
	//
	// 100
	Build *string `json:"Build,omitempty" xml:"Build,omitempty"`
	// example:
	//
	// china
	BusinessCountry *string `json:"BusinessCountry,omitempty" xml:"BusinessCountry,omitempty"`
	// example:
	//
	// -
	BusinessErrorExtData *string `json:"BusinessErrorExtData,omitempty" xml:"BusinessErrorExtData,omitempty"`
	// example:
	//
	// -
	BusinessLogType *string `json:"BusinessLogType,omitempty" xml:"BusinessLogType,omitempty"`
	// carrier
	//
	// example:
	//
	// CT
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// example:
	//
	// ant
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// example:
	//
	// 375
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// 100.104.16.217:51534
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// 1739808000000
	ClientTime *int64 `json:"ClientTime,omitempty" xml:"ClientTime,omitempty"`
	// example:
	//
	// 0
	ColNo *string `json:"ColNo,omitempty" xml:"ColNo,omitempty"`
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// example:
	//
	// arm
	CpuModel *string `json:"CpuModel,omitempty" xml:"CpuModel,omitempty"`
	// example:
	//
	// -
	CruxModule *string `json:"CruxModule,omitempty" xml:"CruxModule,omitempty"`
	// example:
	//
	// -
	CruxStack *string `json:"CruxStack,omitempty" xml:"CruxStack,omitempty"`
	// example:
	//
	// -
	CruxStackHash *int64 `json:"CruxStackHash,omitempty" xml:"CruxStackHash,omitempty"`
	// CruxStackTrace
	//
	// example:
	//
	// -
	CruxStackTrace *string `json:"CruxStackTrace,omitempty" xml:"CruxStackTrace,omitempty"`
	// example:
	//
	// -
	CruxStackTraceHash *int64 `json:"CruxStackTraceHash,omitempty" xml:"CruxStackTraceHash,omitempty"`
	// example:
	//
	// -
	CustomExceptionType *string `json:"CustomExceptionType,omitempty" xml:"CustomExceptionType,omitempty"`
	// example:
	//
	// -
	CustomInfo *string `json:"CustomInfo,omitempty" xml:"CustomInfo,omitempty"`
	// example:
	//
	// -
	DataDirectory *string `json:"DataDirectory,omitempty" xml:"DataDirectory,omitempty"`
	// example:
	//
	// ad-0001t9c1b6y48kqcd44s-105
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// ASUS_X00GD
	DeviceModel *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	// example:
	//
	// 22jtJ213saqfB1yBgBKhayYPwcZ2
	Did *string `json:"Did,omitempty" xml:"Did,omitempty"`
	// example:
	//
	// sha256:b3b79f1f6c2beca7ac93a9518a35d542e0bfeb7ae64c9e57e9edd4719489927c
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// example:
	//
	// -
	DigestHash *string `json:"DigestHash,omitempty" xml:"DigestHash,omitempty"`
	// example:
	//
	// 0.5
	DomScore *string `json:"DomScore,omitempty" xml:"DomScore,omitempty"`
	// example:
	//
	// -
	DownloadCacheDirectory *string `json:"DownloadCacheDirectory,omitempty" xml:"DownloadCacheDirectory,omitempty"`
	// example:
	//
	// OOM
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	// example:
	//
	// jsError
	ErrorType *string `json:"ErrorType,omitempty" xml:"ErrorType,omitempty"`
	// example:
	//
	// 0DB-22jtJZzBvqfB1yBgBKhayYPwcZ2
	EventId *int32 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// -
	EventLog *string `json:"EventLog,omitempty" xml:"EventLog,omitempty"`
	// example:
	//
	// -
	ExceptionArg1 *string `json:"ExceptionArg1,omitempty" xml:"ExceptionArg1,omitempty"`
	// example:
	//
	// -
	ExceptionArg2 *string `json:"ExceptionArg2,omitempty" xml:"ExceptionArg2,omitempty"`
	// example:
	//
	// -
	ExceptionArg3 *string `json:"ExceptionArg3,omitempty" xml:"ExceptionArg3,omitempty"`
	// example:
	//
	// -
	ExceptionCode *string `json:"ExceptionCode,omitempty" xml:"ExceptionCode,omitempty"`
	// example:
	//
	// -
	ExceptionCodes *string `json:"ExceptionCodes,omitempty" xml:"ExceptionCodes,omitempty"`
	// example:
	//
	// -
	ExceptionDetail *string `json:"ExceptionDetail,omitempty" xml:"ExceptionDetail,omitempty"`
	// example:
	//
	// -
	ExceptionId *string `json:"ExceptionId,omitempty" xml:"ExceptionId,omitempty"`
	// example:
	//
	// Table \\"123_temp.image_vw_6191426\\" doesn\\"t exist
	ExceptionMsg *string `json:"ExceptionMsg,omitempty" xml:"ExceptionMsg,omitempty"`
	// example:
	//
	// -
	ExceptionSubtype *string `json:"ExceptionSubtype,omitempty" xml:"ExceptionSubtype,omitempty"`
	// example:
	//
	// -
	ExceptionType *string `json:"ExceptionType,omitempty" xml:"ExceptionType,omitempty"`
	// example:
	//
	// -
	ExceptionVersion *string `json:"ExceptionVersion,omitempty" xml:"ExceptionVersion,omitempty"`
	// example:
	//
	// exporting
	ExportStatus *string `json:"ExportStatus,omitempty" xml:"ExportStatus,omitempty"`
	// example:
	//
	// -
	ExternalStorageDirectory *string `json:"ExternalStorageDirectory,omitempty" xml:"ExternalStorageDirectory,omitempty"`
	// example:
	//
	// -
	FeatureScene *string `json:"FeatureScene,omitempty" xml:"FeatureScene,omitempty"`
	// example:
	//
	// baselineCspm_20241005
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// dubhe-service/src/main/java/com/viatris/dubhe/monitortask/service/Impl/DataSyncMonitorRecordServiceImpl.java
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// example:
	//
	// chunk-56494f41.3df8bd75.css
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// example:
	//
	// -
	FlutterRoute *string `json:"FlutterRoute,omitempty" xml:"FlutterRoute,omitempty"`
	// example:
	//
	// 1
	ForceGround *int32 `json:"ForceGround,omitempty" xml:"ForceGround,omitempty"`
	// example:
	//
	// 1
	ForeGround *int32 `json:"ForeGround,omitempty" xml:"ForeGround,omitempty"`
	// example:
	//
	// -
	H5FullUrl *string `json:"H5FullUrl,omitempty" xml:"H5FullUrl,omitempty"`
	// example:
	//
	// -
	H5ShortUrl *string `json:"H5ShortUrl,omitempty" xml:"H5ShortUrl,omitempty"`
	// example:
	//
	// 0
	HasOpenMultiProcessMode *int32 `json:"HasOpenMultiProcessMode,omitempty" xml:"HasOpenMultiProcessMode,omitempty"`
	// example:
	//
	// -
	HasSdCard *int32 `json:"HasSdCard,omitempty" xml:"HasSdCard,omitempty"`
	// example:
	//
	// 0
	HasSgSecurityConfigKey *int32 `json:"HasSgSecurityConfigKey,omitempty" xml:"HasSgSecurityConfigKey,omitempty"`
	// IMEI
	//
	// example:
	//
	// 2704b9b0d94b2f3897f82119a1900cd9c7a8108d
	Imei *string `json:"Imei,omitempty" xml:"Imei,omitempty"`
	// IMSI
	//
	// example:
	//
	// 460049842500442
	Imsi *string `json:"Imsi,omitempty" xml:"Imsi,omitempty"`
	// example:
	//
	// 0
	InMainProcess *int32 `json:"InMainProcess,omitempty" xml:"InMainProcess,omitempty"`
	// example:
	//
	// -
	InstallSdCard *int32 `json:"InstallSdCard,omitempty" xml:"InstallSdCard,omitempty"`
	// example:
	//
	// -
	IsBackTrace *int32 `json:"IsBackTrace,omitempty" xml:"IsBackTrace,omitempty"`
	// example:
	//
	// 0
	IsSpeedVersion *int32 `json:"IsSpeedVersion,omitempty" xml:"IsSpeedVersion,omitempty"`
	// ISP
	//
	// example:
	//
	// cmcc
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// example:
	//
	// java.lang.NullPointerException: Attempt to invoke virtual method \\"java.lang.Object
	JsBacktrace *string `json:"JsBacktrace,omitempty" xml:"JsBacktrace,omitempty"`
	// example:
	//
	// cn
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 2025-02-25 18:02:50
	LaunchedTime *string `json:"LaunchedTime,omitempty" xml:"LaunchedTime,omitempty"`
	// example:
	//
	// 0
	LineNo *string `json:"LineNo,omitempty" xml:"LineNo,omitempty"`
	// example:
	//
	// -
	LogHash *int64 `json:"LogHash,omitempty" xml:"LogHash,omitempty"`
	// example:
	//
	// -
	MainLog *string `json:"MainLog,omitempty" xml:"MainLog,omitempty"`
	// example:
	//
	// -
	MemoryMap *string `json:"MemoryMap,omitempty" xml:"MemoryMap,omitempty"`
	// example:
	//
	// -
	MoreInfo1 *string `json:"MoreInfo1,omitempty" xml:"MoreInfo1,omitempty"`
	// example:
	//
	// -
	MoreInfo2 *string `json:"MoreInfo2,omitempty" xml:"MoreInfo2,omitempty"`
	// example:
	//
	// -
	MoreInfo3 *string `json:"MoreInfo3,omitempty" xml:"MoreInfo3,omitempty"`
	// example:
	//
	// -
	NativeAllThreadDump *string `json:"NativeAllThreadDump,omitempty" xml:"NativeAllThreadDump,omitempty"`
	// Native map
	//
	// example:
	//
	// -
	NativeMaps *string `json:"NativeMaps,omitempty" xml:"NativeMaps,omitempty"`
	// example:
	//
	// -
	NeedReCluster *int32 `json:"NeedReCluster,omitempty" xml:"NeedReCluster,omitempty"`
	// example:
	//
	// -
	OpenedFileCount *int32 `json:"OpenedFileCount,omitempty" xml:"OpenedFileCount,omitempty"`
	// example:
	//
	// [{\\"status\\":1,\\"target\\":\\"5461e0ab-340f-4a6b-9bbe-f75e44985420\\"}]
	Operations *string `json:"Operations,omitempty" xml:"Operations,omitempty"`
	// example:
	//
	// []
	OriginData *string `json:"OriginData,omitempty" xml:"OriginData,omitempty"`
	// example:
	//
	// -
	OriginUri *string `json:"OriginUri,omitempty" xml:"OriginUri,omitempty"`
	// example:
	//
	// android
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// example:
	//
	// 1
	OsVersion *string `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	// example:
	//
	// {\\"Content\\": [], \\"TotalPages\\": 2, \\"TotalItems\\": 31, \\"PageNumber\\": 4, \\"PageSize\\": 20}
	Page *string `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// com.aliyun.demo
	ParentProcessName *string `json:"ParentProcessName,omitempty" xml:"ParentProcessName,omitempty"`
	// example:
	//
	// ce0ak81asd@62706c656a7f1a1
	Pid *int32 `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// example:
	//
	// wechat
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// null
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// -
	ReportContent *string `json:"ReportContent,omitempty" xml:"ReportContent,omitempty"`
	// example:
	//
	// 2
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// example:
	//
	// 720*1280
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// example:
	//
	// -
	RootDirectory *string `json:"RootDirectory,omitempty" xml:"RootDirectory,omitempty"`
	// example:
	//
	// -
	RuntimeExtData *string `json:"RuntimeExtData,omitempty" xml:"RuntimeExtData,omitempty"`
	// example:
	//
	// -
	SceneValue *string `json:"SceneValue,omitempty" xml:"SceneValue,omitempty"`
	// example:
	//
	// software_nls_tts_offline_standard
	SdkType *string `json:"SdkType,omitempty" xml:"SdkType,omitempty"`
	// example:
	//
	// 5.2.3
	SdkVersion *string `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	// example:
	//
	// 3481243636390427020
	Seq *string `json:"Seq,omitempty" xml:"Seq,omitempty"`
	// example:
	//
	// 1739808000000
	ServerTime *int64 `json:"ServerTime,omitempty" xml:"ServerTime,omitempty"`
	// example:
	//
	// 9ec8e399-51cc-4283-8ba4-f6df78246249
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// -
	SimpleReportContent *string `json:"SimpleReportContent,omitempty" xml:"SimpleReportContent,omitempty"`
	// example:
	//
	// -
	SoLibBuild *string `json:"SoLibBuild,omitempty" xml:"SoLibBuild,omitempty"`
	// SpeedFlags
	//
	// example:
	//
	// -
	SpeedFlags *string `json:"SpeedFlags,omitempty" xml:"SpeedFlags,omitempty"`
	// example:
	//
	// java.lang.NullPointerException: Attempt to invoke virtual method \\"java.lang.Object java.lang.ref.WeakReference.get()\\" on a null object reference
	Stack *string `json:"Stack,omitempty" xml:"Stack,omitempty"`
	// example:
	//
	// success
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StructuredStack *string `json:"StructuredStack,omitempty" xml:"StructuredStack,omitempty"`
	// example:
	//
	// {\\"HandledCheckToday\\": 0, \\"RiskCheckCnt\\": 34, \\"RiskWarningCnt\\": 57, \\"RiskDays\\": 30, \\"HandledCheckTotal\\": 5, \\"HandledDays\\": 365}
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// -
	SymbolicFileType *string `json:"SymbolicFileType,omitempty" xml:"SymbolicFileType,omitempty"`
	// example:
	//
	// -
	SysLog *string `json:"SysLog,omitempty" xml:"SysLog,omitempty"`
	// example:
	//
	// null
	ThreadName *string `json:"ThreadName,omitempty" xml:"ThreadName,omitempty"`
	// example:
	//
	// Process Name: \\"com.muyuan.mytongdriver\\"
	//
	// Thread Name: \\"main\\"
	Threads *string `json:"Threads,omitempty" xml:"Threads,omitempty"`
	// example:
	//
	// 2359
	Tid *int32 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// example:
	//
	// -
	Track *string `json:"Track,omitempty" xml:"Track,omitempty"`
	// example:
	//
	// 2025-02-25 18:02:50
	TriggeredTime *string `json:"TriggeredTime,omitempty" xml:"TriggeredTime,omitempty"`
	// example:
	//
	// 1740412800000
	UploadTime *int64 `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
	// URI
	//
	// example:
	//
	// https://oss.aliyuncs.com/aliyun_id_photo_bucket/default_handsome.jpg
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// example:
	//
	// -
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// mark
	UserNick *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
	// utdid
	//
	// example:
	//
	// 2704b9b0d94b2f3897f82119a
	Utdid *string `json:"Utdid,omitempty" xml:"Utdid,omitempty"`
	// uuid
	//
	// example:
	//
	// 105a87c6-f299-42e5-8fc7-1066a47a96e5
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// view
	//
	// example:
	//
	// com.aliyun.view1
	View *string `json:"View,omitempty" xml:"View,omitempty"`
	// example:
	//
	// -
	WeexFullUrl *string `json:"WeexFullUrl,omitempty" xml:"WeexFullUrl,omitempty"`
	// example:
	//
	// -
	WeexShortUrl *string `json:"WeexShortUrl,omitempty" xml:"WeexShortUrl,omitempty"`
	// example:
	//
	// -
	WriteLimit *int32 `json:"WriteLimit,omitempty" xml:"WriteLimit,omitempty"`
}

func (s GetErrorResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s GetErrorResponseBodyModel) GoString() string {
	return s.String()
}

func (s *GetErrorResponseBodyModel) GetAccess() *string {
	return s.Access
}

func (s *GetErrorResponseBodyModel) GetAccessSubType() *string {
	return s.AccessSubType
}

func (s *GetErrorResponseBodyModel) GetAdditionalCallbackInfo() *string {
	return s.AdditionalCallbackInfo
}

func (s *GetErrorResponseBodyModel) GetAdditionalCrossPlatformCrashInfo() *string {
	return s.AdditionalCrossPlatformCrashInfo
}

func (s *GetErrorResponseBodyModel) GetAdditionalCustomInfo() *string {
	return s.AdditionalCustomInfo
}

func (s *GetErrorResponseBodyModel) GetAndroidVm() *string {
	return s.AndroidVm
}

func (s *GetErrorResponseBodyModel) GetAppId() *string {
	return s.AppId
}

func (s *GetErrorResponseBodyModel) GetAppKey() *int64 {
	return s.AppKey
}

func (s *GetErrorResponseBodyModel) GetAppVersion() *string {
	return s.AppVersion
}

func (s *GetErrorResponseBodyModel) GetArgHash() *int64 {
	return s.ArgHash
}

func (s *GetErrorResponseBodyModel) GetArgs() *string {
	return s.Args
}

func (s *GetErrorResponseBodyModel) GetBacktrace() *string {
	return s.Backtrace
}

func (s *GetErrorResponseBodyModel) GetBanner() *string {
	return s.Banner
}

func (s *GetErrorResponseBodyModel) GetBinaryUuids() *string {
	return s.BinaryUuids
}

func (s *GetErrorResponseBodyModel) GetBrand() *string {
	return s.Brand
}

func (s *GetErrorResponseBodyModel) GetBrowserName() *string {
	return s.BrowserName
}

func (s *GetErrorResponseBodyModel) GetBrowserVersion() *string {
	return s.BrowserVersion
}

func (s *GetErrorResponseBodyModel) GetBuild() *string {
	return s.Build
}

func (s *GetErrorResponseBodyModel) GetBusinessCountry() *string {
	return s.BusinessCountry
}

func (s *GetErrorResponseBodyModel) GetBusinessErrorExtData() *string {
	return s.BusinessErrorExtData
}

func (s *GetErrorResponseBodyModel) GetBusinessLogType() *string {
	return s.BusinessLogType
}

func (s *GetErrorResponseBodyModel) GetCarrier() *string {
	return s.Carrier
}

func (s *GetErrorResponseBodyModel) GetChannel() *string {
	return s.Channel
}

func (s *GetErrorResponseBodyModel) GetCity() *string {
	return s.City
}

func (s *GetErrorResponseBodyModel) GetClientIp() *string {
	return s.ClientIp
}

func (s *GetErrorResponseBodyModel) GetClientTime() *int64 {
	return s.ClientTime
}

func (s *GetErrorResponseBodyModel) GetColNo() *string {
	return s.ColNo
}

func (s *GetErrorResponseBodyModel) GetCountry() *string {
	return s.Country
}

func (s *GetErrorResponseBodyModel) GetCpuModel() *string {
	return s.CpuModel
}

func (s *GetErrorResponseBodyModel) GetCruxModule() *string {
	return s.CruxModule
}

func (s *GetErrorResponseBodyModel) GetCruxStack() *string {
	return s.CruxStack
}

func (s *GetErrorResponseBodyModel) GetCruxStackHash() *int64 {
	return s.CruxStackHash
}

func (s *GetErrorResponseBodyModel) GetCruxStackTrace() *string {
	return s.CruxStackTrace
}

func (s *GetErrorResponseBodyModel) GetCruxStackTraceHash() *int64 {
	return s.CruxStackTraceHash
}

func (s *GetErrorResponseBodyModel) GetCustomExceptionType() *string {
	return s.CustomExceptionType
}

func (s *GetErrorResponseBodyModel) GetCustomInfo() *string {
	return s.CustomInfo
}

func (s *GetErrorResponseBodyModel) GetDataDirectory() *string {
	return s.DataDirectory
}

func (s *GetErrorResponseBodyModel) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetErrorResponseBodyModel) GetDeviceModel() *string {
	return s.DeviceModel
}

func (s *GetErrorResponseBodyModel) GetDid() *string {
	return s.Did
}

func (s *GetErrorResponseBodyModel) GetDigest() *string {
	return s.Digest
}

func (s *GetErrorResponseBodyModel) GetDigestHash() *string {
	return s.DigestHash
}

func (s *GetErrorResponseBodyModel) GetDomScore() *string {
	return s.DomScore
}

func (s *GetErrorResponseBodyModel) GetDownloadCacheDirectory() *string {
	return s.DownloadCacheDirectory
}

func (s *GetErrorResponseBodyModel) GetErrorName() *string {
	return s.ErrorName
}

func (s *GetErrorResponseBodyModel) GetErrorType() *string {
	return s.ErrorType
}

func (s *GetErrorResponseBodyModel) GetEventId() *int32 {
	return s.EventId
}

func (s *GetErrorResponseBodyModel) GetEventLog() *string {
	return s.EventLog
}

func (s *GetErrorResponseBodyModel) GetExceptionArg1() *string {
	return s.ExceptionArg1
}

func (s *GetErrorResponseBodyModel) GetExceptionArg2() *string {
	return s.ExceptionArg2
}

func (s *GetErrorResponseBodyModel) GetExceptionArg3() *string {
	return s.ExceptionArg3
}

func (s *GetErrorResponseBodyModel) GetExceptionCode() *string {
	return s.ExceptionCode
}

func (s *GetErrorResponseBodyModel) GetExceptionCodes() *string {
	return s.ExceptionCodes
}

func (s *GetErrorResponseBodyModel) GetExceptionDetail() *string {
	return s.ExceptionDetail
}

func (s *GetErrorResponseBodyModel) GetExceptionId() *string {
	return s.ExceptionId
}

func (s *GetErrorResponseBodyModel) GetExceptionMsg() *string {
	return s.ExceptionMsg
}

func (s *GetErrorResponseBodyModel) GetExceptionSubtype() *string {
	return s.ExceptionSubtype
}

func (s *GetErrorResponseBodyModel) GetExceptionType() *string {
	return s.ExceptionType
}

func (s *GetErrorResponseBodyModel) GetExceptionVersion() *string {
	return s.ExceptionVersion
}

func (s *GetErrorResponseBodyModel) GetExportStatus() *string {
	return s.ExportStatus
}

func (s *GetErrorResponseBodyModel) GetExternalStorageDirectory() *string {
	return s.ExternalStorageDirectory
}

func (s *GetErrorResponseBodyModel) GetFeatureScene() *string {
	return s.FeatureScene
}

func (s *GetErrorResponseBodyModel) GetFileName() *string {
	return s.FileName
}

func (s *GetErrorResponseBodyModel) GetFilePath() *string {
	return s.FilePath
}

func (s *GetErrorResponseBodyModel) GetFilename() *string {
	return s.Filename
}

func (s *GetErrorResponseBodyModel) GetFlutterRoute() *string {
	return s.FlutterRoute
}

func (s *GetErrorResponseBodyModel) GetForceGround() *int32 {
	return s.ForceGround
}

func (s *GetErrorResponseBodyModel) GetForeGround() *int32 {
	return s.ForeGround
}

func (s *GetErrorResponseBodyModel) GetH5FullUrl() *string {
	return s.H5FullUrl
}

func (s *GetErrorResponseBodyModel) GetH5ShortUrl() *string {
	return s.H5ShortUrl
}

func (s *GetErrorResponseBodyModel) GetHasOpenMultiProcessMode() *int32 {
	return s.HasOpenMultiProcessMode
}

func (s *GetErrorResponseBodyModel) GetHasSdCard() *int32 {
	return s.HasSdCard
}

func (s *GetErrorResponseBodyModel) GetHasSgSecurityConfigKey() *int32 {
	return s.HasSgSecurityConfigKey
}

func (s *GetErrorResponseBodyModel) GetImei() *string {
	return s.Imei
}

func (s *GetErrorResponseBodyModel) GetImsi() *string {
	return s.Imsi
}

func (s *GetErrorResponseBodyModel) GetInMainProcess() *int32 {
	return s.InMainProcess
}

func (s *GetErrorResponseBodyModel) GetInstallSdCard() *int32 {
	return s.InstallSdCard
}

func (s *GetErrorResponseBodyModel) GetIsBackTrace() *int32 {
	return s.IsBackTrace
}

func (s *GetErrorResponseBodyModel) GetIsSpeedVersion() *int32 {
	return s.IsSpeedVersion
}

func (s *GetErrorResponseBodyModel) GetIsp() *string {
	return s.Isp
}

func (s *GetErrorResponseBodyModel) GetJsBacktrace() *string {
	return s.JsBacktrace
}

func (s *GetErrorResponseBodyModel) GetLanguage() *string {
	return s.Language
}

func (s *GetErrorResponseBodyModel) GetLaunchedTime() *string {
	return s.LaunchedTime
}

func (s *GetErrorResponseBodyModel) GetLineNo() *string {
	return s.LineNo
}

func (s *GetErrorResponseBodyModel) GetLogHash() *int64 {
	return s.LogHash
}

func (s *GetErrorResponseBodyModel) GetMainLog() *string {
	return s.MainLog
}

func (s *GetErrorResponseBodyModel) GetMemoryMap() *string {
	return s.MemoryMap
}

func (s *GetErrorResponseBodyModel) GetMoreInfo1() *string {
	return s.MoreInfo1
}

func (s *GetErrorResponseBodyModel) GetMoreInfo2() *string {
	return s.MoreInfo2
}

func (s *GetErrorResponseBodyModel) GetMoreInfo3() *string {
	return s.MoreInfo3
}

func (s *GetErrorResponseBodyModel) GetNativeAllThreadDump() *string {
	return s.NativeAllThreadDump
}

func (s *GetErrorResponseBodyModel) GetNativeMaps() *string {
	return s.NativeMaps
}

func (s *GetErrorResponseBodyModel) GetNeedReCluster() *int32 {
	return s.NeedReCluster
}

func (s *GetErrorResponseBodyModel) GetOpenedFileCount() *int32 {
	return s.OpenedFileCount
}

func (s *GetErrorResponseBodyModel) GetOperations() *string {
	return s.Operations
}

func (s *GetErrorResponseBodyModel) GetOriginData() *string {
	return s.OriginData
}

func (s *GetErrorResponseBodyModel) GetOriginUri() *string {
	return s.OriginUri
}

func (s *GetErrorResponseBodyModel) GetOs() *string {
	return s.Os
}

func (s *GetErrorResponseBodyModel) GetOsVersion() *string {
	return s.OsVersion
}

func (s *GetErrorResponseBodyModel) GetPage() *string {
	return s.Page
}

func (s *GetErrorResponseBodyModel) GetParentProcessName() *string {
	return s.ParentProcessName
}

func (s *GetErrorResponseBodyModel) GetPid() *int32 {
	return s.Pid
}

func (s *GetErrorResponseBodyModel) GetPlatform() *string {
	return s.Platform
}

func (s *GetErrorResponseBodyModel) GetProcessName() *string {
	return s.ProcessName
}

func (s *GetErrorResponseBodyModel) GetProvince() *string {
	return s.Province
}

func (s *GetErrorResponseBodyModel) GetReason() *string {
	return s.Reason
}

func (s *GetErrorResponseBodyModel) GetReportContent() *string {
	return s.ReportContent
}

func (s *GetErrorResponseBodyModel) GetReportType() *string {
	return s.ReportType
}

func (s *GetErrorResponseBodyModel) GetResolution() *string {
	return s.Resolution
}

func (s *GetErrorResponseBodyModel) GetRootDirectory() *string {
	return s.RootDirectory
}

func (s *GetErrorResponseBodyModel) GetRuntimeExtData() *string {
	return s.RuntimeExtData
}

func (s *GetErrorResponseBodyModel) GetSceneValue() *string {
	return s.SceneValue
}

func (s *GetErrorResponseBodyModel) GetSdkType() *string {
	return s.SdkType
}

func (s *GetErrorResponseBodyModel) GetSdkVersion() *string {
	return s.SdkVersion
}

func (s *GetErrorResponseBodyModel) GetSeq() *string {
	return s.Seq
}

func (s *GetErrorResponseBodyModel) GetServerTime() *int64 {
	return s.ServerTime
}

func (s *GetErrorResponseBodyModel) GetSessionId() *string {
	return s.SessionId
}

func (s *GetErrorResponseBodyModel) GetSimpleReportContent() *string {
	return s.SimpleReportContent
}

func (s *GetErrorResponseBodyModel) GetSoLibBuild() *string {
	return s.SoLibBuild
}

func (s *GetErrorResponseBodyModel) GetSpeedFlags() *string {
	return s.SpeedFlags
}

func (s *GetErrorResponseBodyModel) GetStack() *string {
	return s.Stack
}

func (s *GetErrorResponseBodyModel) GetStatus() *string {
	return s.Status
}

func (s *GetErrorResponseBodyModel) GetStructuredStack() *string {
	return s.StructuredStack
}

func (s *GetErrorResponseBodyModel) GetSummary() *string {
	return s.Summary
}

func (s *GetErrorResponseBodyModel) GetSymbolicFileType() *string {
	return s.SymbolicFileType
}

func (s *GetErrorResponseBodyModel) GetSysLog() *string {
	return s.SysLog
}

func (s *GetErrorResponseBodyModel) GetThreadName() *string {
	return s.ThreadName
}

func (s *GetErrorResponseBodyModel) GetThreads() *string {
	return s.Threads
}

func (s *GetErrorResponseBodyModel) GetTid() *int32 {
	return s.Tid
}

func (s *GetErrorResponseBodyModel) GetTrack() *string {
	return s.Track
}

func (s *GetErrorResponseBodyModel) GetTriggeredTime() *string {
	return s.TriggeredTime
}

func (s *GetErrorResponseBodyModel) GetUploadTime() *int64 {
	return s.UploadTime
}

func (s *GetErrorResponseBodyModel) GetUri() *string {
	return s.Uri
}

func (s *GetErrorResponseBodyModel) GetUser() *string {
	return s.User
}

func (s *GetErrorResponseBodyModel) GetUserId() *string {
	return s.UserId
}

func (s *GetErrorResponseBodyModel) GetUserNick() *string {
	return s.UserNick
}

func (s *GetErrorResponseBodyModel) GetUtdid() *string {
	return s.Utdid
}

func (s *GetErrorResponseBodyModel) GetUuid() *string {
	return s.Uuid
}

func (s *GetErrorResponseBodyModel) GetView() *string {
	return s.View
}

func (s *GetErrorResponseBodyModel) GetWeexFullUrl() *string {
	return s.WeexFullUrl
}

func (s *GetErrorResponseBodyModel) GetWeexShortUrl() *string {
	return s.WeexShortUrl
}

func (s *GetErrorResponseBodyModel) GetWriteLimit() *int32 {
	return s.WriteLimit
}

func (s *GetErrorResponseBodyModel) SetAccess(v string) *GetErrorResponseBodyModel {
	s.Access = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetAccessSubType(v string) *GetErrorResponseBodyModel {
	s.AccessSubType = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetAdditionalCallbackInfo(v string) *GetErrorResponseBodyModel {
	s.AdditionalCallbackInfo = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetAdditionalCrossPlatformCrashInfo(v string) *GetErrorResponseBodyModel {
	s.AdditionalCrossPlatformCrashInfo = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetAdditionalCustomInfo(v string) *GetErrorResponseBodyModel {
	s.AdditionalCustomInfo = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetAndroidVm(v string) *GetErrorResponseBodyModel {
	s.AndroidVm = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetAppId(v string) *GetErrorResponseBodyModel {
	s.AppId = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetAppKey(v int64) *GetErrorResponseBodyModel {
	s.AppKey = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetAppVersion(v string) *GetErrorResponseBodyModel {
	s.AppVersion = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetArgHash(v int64) *GetErrorResponseBodyModel {
	s.ArgHash = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetArgs(v string) *GetErrorResponseBodyModel {
	s.Args = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetBacktrace(v string) *GetErrorResponseBodyModel {
	s.Backtrace = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetBanner(v string) *GetErrorResponseBodyModel {
	s.Banner = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetBinaryUuids(v string) *GetErrorResponseBodyModel {
	s.BinaryUuids = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetBrand(v string) *GetErrorResponseBodyModel {
	s.Brand = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetBrowserName(v string) *GetErrorResponseBodyModel {
	s.BrowserName = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetBrowserVersion(v string) *GetErrorResponseBodyModel {
	s.BrowserVersion = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetBuild(v string) *GetErrorResponseBodyModel {
	s.Build = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetBusinessCountry(v string) *GetErrorResponseBodyModel {
	s.BusinessCountry = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetBusinessErrorExtData(v string) *GetErrorResponseBodyModel {
	s.BusinessErrorExtData = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetBusinessLogType(v string) *GetErrorResponseBodyModel {
	s.BusinessLogType = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetCarrier(v string) *GetErrorResponseBodyModel {
	s.Carrier = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetChannel(v string) *GetErrorResponseBodyModel {
	s.Channel = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetCity(v string) *GetErrorResponseBodyModel {
	s.City = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetClientIp(v string) *GetErrorResponseBodyModel {
	s.ClientIp = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetClientTime(v int64) *GetErrorResponseBodyModel {
	s.ClientTime = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetColNo(v string) *GetErrorResponseBodyModel {
	s.ColNo = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetCountry(v string) *GetErrorResponseBodyModel {
	s.Country = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetCpuModel(v string) *GetErrorResponseBodyModel {
	s.CpuModel = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetCruxModule(v string) *GetErrorResponseBodyModel {
	s.CruxModule = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetCruxStack(v string) *GetErrorResponseBodyModel {
	s.CruxStack = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetCruxStackHash(v int64) *GetErrorResponseBodyModel {
	s.CruxStackHash = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetCruxStackTrace(v string) *GetErrorResponseBodyModel {
	s.CruxStackTrace = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetCruxStackTraceHash(v int64) *GetErrorResponseBodyModel {
	s.CruxStackTraceHash = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetCustomExceptionType(v string) *GetErrorResponseBodyModel {
	s.CustomExceptionType = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetCustomInfo(v string) *GetErrorResponseBodyModel {
	s.CustomInfo = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetDataDirectory(v string) *GetErrorResponseBodyModel {
	s.DataDirectory = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetDeviceId(v string) *GetErrorResponseBodyModel {
	s.DeviceId = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetDeviceModel(v string) *GetErrorResponseBodyModel {
	s.DeviceModel = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetDid(v string) *GetErrorResponseBodyModel {
	s.Did = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetDigest(v string) *GetErrorResponseBodyModel {
	s.Digest = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetDigestHash(v string) *GetErrorResponseBodyModel {
	s.DigestHash = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetDomScore(v string) *GetErrorResponseBodyModel {
	s.DomScore = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetDownloadCacheDirectory(v string) *GetErrorResponseBodyModel {
	s.DownloadCacheDirectory = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetErrorName(v string) *GetErrorResponseBodyModel {
	s.ErrorName = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetErrorType(v string) *GetErrorResponseBodyModel {
	s.ErrorType = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetEventId(v int32) *GetErrorResponseBodyModel {
	s.EventId = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetEventLog(v string) *GetErrorResponseBodyModel {
	s.EventLog = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetExceptionArg1(v string) *GetErrorResponseBodyModel {
	s.ExceptionArg1 = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetExceptionArg2(v string) *GetErrorResponseBodyModel {
	s.ExceptionArg2 = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetExceptionArg3(v string) *GetErrorResponseBodyModel {
	s.ExceptionArg3 = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetExceptionCode(v string) *GetErrorResponseBodyModel {
	s.ExceptionCode = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetExceptionCodes(v string) *GetErrorResponseBodyModel {
	s.ExceptionCodes = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetExceptionDetail(v string) *GetErrorResponseBodyModel {
	s.ExceptionDetail = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetExceptionId(v string) *GetErrorResponseBodyModel {
	s.ExceptionId = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetExceptionMsg(v string) *GetErrorResponseBodyModel {
	s.ExceptionMsg = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetExceptionSubtype(v string) *GetErrorResponseBodyModel {
	s.ExceptionSubtype = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetExceptionType(v string) *GetErrorResponseBodyModel {
	s.ExceptionType = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetExceptionVersion(v string) *GetErrorResponseBodyModel {
	s.ExceptionVersion = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetExportStatus(v string) *GetErrorResponseBodyModel {
	s.ExportStatus = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetExternalStorageDirectory(v string) *GetErrorResponseBodyModel {
	s.ExternalStorageDirectory = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetFeatureScene(v string) *GetErrorResponseBodyModel {
	s.FeatureScene = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetFileName(v string) *GetErrorResponseBodyModel {
	s.FileName = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetFilePath(v string) *GetErrorResponseBodyModel {
	s.FilePath = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetFilename(v string) *GetErrorResponseBodyModel {
	s.Filename = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetFlutterRoute(v string) *GetErrorResponseBodyModel {
	s.FlutterRoute = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetForceGround(v int32) *GetErrorResponseBodyModel {
	s.ForceGround = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetForeGround(v int32) *GetErrorResponseBodyModel {
	s.ForeGround = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetH5FullUrl(v string) *GetErrorResponseBodyModel {
	s.H5FullUrl = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetH5ShortUrl(v string) *GetErrorResponseBodyModel {
	s.H5ShortUrl = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetHasOpenMultiProcessMode(v int32) *GetErrorResponseBodyModel {
	s.HasOpenMultiProcessMode = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetHasSdCard(v int32) *GetErrorResponseBodyModel {
	s.HasSdCard = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetHasSgSecurityConfigKey(v int32) *GetErrorResponseBodyModel {
	s.HasSgSecurityConfigKey = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetImei(v string) *GetErrorResponseBodyModel {
	s.Imei = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetImsi(v string) *GetErrorResponseBodyModel {
	s.Imsi = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetInMainProcess(v int32) *GetErrorResponseBodyModel {
	s.InMainProcess = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetInstallSdCard(v int32) *GetErrorResponseBodyModel {
	s.InstallSdCard = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetIsBackTrace(v int32) *GetErrorResponseBodyModel {
	s.IsBackTrace = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetIsSpeedVersion(v int32) *GetErrorResponseBodyModel {
	s.IsSpeedVersion = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetIsp(v string) *GetErrorResponseBodyModel {
	s.Isp = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetJsBacktrace(v string) *GetErrorResponseBodyModel {
	s.JsBacktrace = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetLanguage(v string) *GetErrorResponseBodyModel {
	s.Language = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetLaunchedTime(v string) *GetErrorResponseBodyModel {
	s.LaunchedTime = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetLineNo(v string) *GetErrorResponseBodyModel {
	s.LineNo = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetLogHash(v int64) *GetErrorResponseBodyModel {
	s.LogHash = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetMainLog(v string) *GetErrorResponseBodyModel {
	s.MainLog = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetMemoryMap(v string) *GetErrorResponseBodyModel {
	s.MemoryMap = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetMoreInfo1(v string) *GetErrorResponseBodyModel {
	s.MoreInfo1 = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetMoreInfo2(v string) *GetErrorResponseBodyModel {
	s.MoreInfo2 = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetMoreInfo3(v string) *GetErrorResponseBodyModel {
	s.MoreInfo3 = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetNativeAllThreadDump(v string) *GetErrorResponseBodyModel {
	s.NativeAllThreadDump = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetNativeMaps(v string) *GetErrorResponseBodyModel {
	s.NativeMaps = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetNeedReCluster(v int32) *GetErrorResponseBodyModel {
	s.NeedReCluster = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetOpenedFileCount(v int32) *GetErrorResponseBodyModel {
	s.OpenedFileCount = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetOperations(v string) *GetErrorResponseBodyModel {
	s.Operations = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetOriginData(v string) *GetErrorResponseBodyModel {
	s.OriginData = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetOriginUri(v string) *GetErrorResponseBodyModel {
	s.OriginUri = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetOs(v string) *GetErrorResponseBodyModel {
	s.Os = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetOsVersion(v string) *GetErrorResponseBodyModel {
	s.OsVersion = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetPage(v string) *GetErrorResponseBodyModel {
	s.Page = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetParentProcessName(v string) *GetErrorResponseBodyModel {
	s.ParentProcessName = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetPid(v int32) *GetErrorResponseBodyModel {
	s.Pid = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetPlatform(v string) *GetErrorResponseBodyModel {
	s.Platform = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetProcessName(v string) *GetErrorResponseBodyModel {
	s.ProcessName = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetProvince(v string) *GetErrorResponseBodyModel {
	s.Province = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetReason(v string) *GetErrorResponseBodyModel {
	s.Reason = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetReportContent(v string) *GetErrorResponseBodyModel {
	s.ReportContent = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetReportType(v string) *GetErrorResponseBodyModel {
	s.ReportType = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetResolution(v string) *GetErrorResponseBodyModel {
	s.Resolution = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetRootDirectory(v string) *GetErrorResponseBodyModel {
	s.RootDirectory = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetRuntimeExtData(v string) *GetErrorResponseBodyModel {
	s.RuntimeExtData = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetSceneValue(v string) *GetErrorResponseBodyModel {
	s.SceneValue = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetSdkType(v string) *GetErrorResponseBodyModel {
	s.SdkType = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetSdkVersion(v string) *GetErrorResponseBodyModel {
	s.SdkVersion = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetSeq(v string) *GetErrorResponseBodyModel {
	s.Seq = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetServerTime(v int64) *GetErrorResponseBodyModel {
	s.ServerTime = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetSessionId(v string) *GetErrorResponseBodyModel {
	s.SessionId = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetSimpleReportContent(v string) *GetErrorResponseBodyModel {
	s.SimpleReportContent = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetSoLibBuild(v string) *GetErrorResponseBodyModel {
	s.SoLibBuild = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetSpeedFlags(v string) *GetErrorResponseBodyModel {
	s.SpeedFlags = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetStack(v string) *GetErrorResponseBodyModel {
	s.Stack = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetStatus(v string) *GetErrorResponseBodyModel {
	s.Status = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetStructuredStack(v string) *GetErrorResponseBodyModel {
	s.StructuredStack = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetSummary(v string) *GetErrorResponseBodyModel {
	s.Summary = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetSymbolicFileType(v string) *GetErrorResponseBodyModel {
	s.SymbolicFileType = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetSysLog(v string) *GetErrorResponseBodyModel {
	s.SysLog = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetThreadName(v string) *GetErrorResponseBodyModel {
	s.ThreadName = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetThreads(v string) *GetErrorResponseBodyModel {
	s.Threads = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetTid(v int32) *GetErrorResponseBodyModel {
	s.Tid = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetTrack(v string) *GetErrorResponseBodyModel {
	s.Track = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetTriggeredTime(v string) *GetErrorResponseBodyModel {
	s.TriggeredTime = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetUploadTime(v int64) *GetErrorResponseBodyModel {
	s.UploadTime = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetUri(v string) *GetErrorResponseBodyModel {
	s.Uri = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetUser(v string) *GetErrorResponseBodyModel {
	s.User = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetUserId(v string) *GetErrorResponseBodyModel {
	s.UserId = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetUserNick(v string) *GetErrorResponseBodyModel {
	s.UserNick = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetUtdid(v string) *GetErrorResponseBodyModel {
	s.Utdid = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetUuid(v string) *GetErrorResponseBodyModel {
	s.Uuid = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetView(v string) *GetErrorResponseBodyModel {
	s.View = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetWeexFullUrl(v string) *GetErrorResponseBodyModel {
	s.WeexFullUrl = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetWeexShortUrl(v string) *GetErrorResponseBodyModel {
	s.WeexShortUrl = &v
	return s
}

func (s *GetErrorResponseBodyModel) SetWriteLimit(v int32) *GetErrorResponseBodyModel {
	s.WriteLimit = &v
	return s
}

func (s *GetErrorResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
