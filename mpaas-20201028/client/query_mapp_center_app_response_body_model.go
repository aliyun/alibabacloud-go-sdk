// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMappCenterAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueryMappCenterAppResult(v *QueryMappCenterAppResponseBodyQueryMappCenterAppResult) *QueryMappCenterAppResponseBody
	GetQueryMappCenterAppResult() *QueryMappCenterAppResponseBodyQueryMappCenterAppResult
	SetRequestId(v string) *QueryMappCenterAppResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryMappCenterAppResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *QueryMappCenterAppResponseBody
	GetResultMessage() *string
}

type QueryMappCenterAppResponseBody struct {
	QueryMappCenterAppResult *QueryMappCenterAppResponseBodyQueryMappCenterAppResult `json:"QueryMappCenterAppResult,omitempty" xml:"QueryMappCenterAppResult,omitempty" type:"Struct"`
	RequestId                *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode               *string                                                 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage            *string                                                 `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMappCenterAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMappCenterAppResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMappCenterAppResponseBody) GetQueryMappCenterAppResult() *QueryMappCenterAppResponseBodyQueryMappCenterAppResult {
	return s.QueryMappCenterAppResult
}

func (s *QueryMappCenterAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMappCenterAppResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryMappCenterAppResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *QueryMappCenterAppResponseBody) SetQueryMappCenterAppResult(v *QueryMappCenterAppResponseBodyQueryMappCenterAppResult) *QueryMappCenterAppResponseBody {
	s.QueryMappCenterAppResult = v
	return s
}

func (s *QueryMappCenterAppResponseBody) SetRequestId(v string) *QueryMappCenterAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMappCenterAppResponseBody) SetResultCode(v string) *QueryMappCenterAppResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMappCenterAppResponseBody) SetResultMessage(v string) *QueryMappCenterAppResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryMappCenterAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryMappCenterAppResponseBodyQueryMappCenterAppResult struct {
	MappCenterApp *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp `json:"MappCenterApp,omitempty" xml:"MappCenterApp,omitempty" type:"Struct"`
	ResultMsg     *string                                                              `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success       *bool                                                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMappCenterAppResponseBodyQueryMappCenterAppResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMappCenterAppResponseBodyQueryMappCenterAppResult) GoString() string {
	return s.String()
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResult) GetMappCenterApp() *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp {
	return s.MappCenterApp
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResult) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResult) SetMappCenterApp(v *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) *QueryMappCenterAppResponseBodyQueryMappCenterAppResult {
	s.MappCenterApp = v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResult) SetResultMsg(v string) *QueryMappCenterAppResponseBodyQueryMappCenterAppResult {
	s.ResultMsg = &v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResult) SetSuccess(v bool) *QueryMappCenterAppResponseBodyQueryMappCenterAppResult {
	s.Success = &v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResult) Validate() error {
	return dara.Validate(s)
}

type QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp struct {
	AndroidConfig *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppAndroidConfig `json:"AndroidConfig,omitempty" xml:"AndroidConfig,omitempty" type:"Struct"`
	AppDesc       *string                                                                           `json:"AppDesc,omitempty" xml:"AppDesc,omitempty"`
	AppIcon       *string                                                                           `json:"AppIcon,omitempty" xml:"AppIcon,omitempty"`
	AppId         *string                                                                           `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName       *string                                                                           `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppSecret     *string                                                                           `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	Creator       *string                                                                           `json:"Creator,omitempty" xml:"Creator,omitempty"`
	GmtCreate     *string                                                                           `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified   *string                                                                           `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id            *int64                                                                            `json:"Id,omitempty" xml:"Id,omitempty"`
	IosConfig     *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppIosConfig     `json:"IosConfig,omitempty" xml:"IosConfig,omitempty" type:"Struct"`
	Modifier      *string                                                                           `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	MonitorJson   *string                                                                           `json:"MonitorJson,omitempty" xml:"MonitorJson,omitempty"`
	Status        *int64                                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId      *string                                                                           `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Type          *int64                                                                            `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) String() string {
	return dara.Prettify(s)
}

func (s QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) GoString() string {
	return s.String()
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) GetAndroidConfig() *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppAndroidConfig {
	return s.AndroidConfig
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) GetAppDesc() *string {
	return s.AppDesc
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) GetAppIcon() *string {
	return s.AppIcon
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) GetAppId() *string {
	return s.AppId
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) GetAppName() *string {
	return s.AppName
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) GetAppSecret() *string {
	return s.AppSecret
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) GetCreator() *string {
	return s.Creator
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) GetId() *int64 {
	return s.Id
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) GetIosConfig() *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppIosConfig {
	return s.IosConfig
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) GetModifier() *string {
	return s.Modifier
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) GetMonitorJson() *string {
	return s.MonitorJson
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) GetStatus() *int64 {
	return s.Status
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) GetType() *int64 {
	return s.Type
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) SetAndroidConfig(v *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppAndroidConfig) *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp {
	s.AndroidConfig = v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) SetAppDesc(v string) *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp {
	s.AppDesc = &v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) SetAppIcon(v string) *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp {
	s.AppIcon = &v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) SetAppId(v string) *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp {
	s.AppId = &v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) SetAppName(v string) *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp {
	s.AppName = &v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) SetAppSecret(v string) *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp {
	s.AppSecret = &v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) SetCreator(v string) *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp {
	s.Creator = &v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) SetGmtCreate(v string) *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp {
	s.GmtCreate = &v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) SetGmtModified(v string) *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp {
	s.GmtModified = &v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) SetId(v int64) *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp {
	s.Id = &v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) SetIosConfig(v *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppIosConfig) *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp {
	s.IosConfig = v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) SetModifier(v string) *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp {
	s.Modifier = &v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) SetMonitorJson(v string) *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp {
	s.MonitorJson = &v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) SetStatus(v int64) *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp {
	s.Status = &v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) SetTenantId(v string) *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp {
	s.TenantId = &v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) SetType(v int64) *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp {
	s.Type = &v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterApp) Validate() error {
	return dara.Validate(s)
}

type QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppAndroidConfig struct {
	CertRSA     *string `json:"CertRSA,omitempty" xml:"CertRSA,omitempty"`
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
}

func (s QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppAndroidConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppAndroidConfig) GoString() string {
	return s.String()
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppAndroidConfig) GetCertRSA() *string {
	return s.CertRSA
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppAndroidConfig) GetPackageName() *string {
	return s.PackageName
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppAndroidConfig) SetCertRSA(v string) *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppAndroidConfig {
	s.CertRSA = &v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppAndroidConfig) SetPackageName(v string) *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppAndroidConfig {
	s.PackageName = &v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppAndroidConfig) Validate() error {
	return dara.Validate(s)
}

type QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppIosConfig struct {
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
}

func (s QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppIosConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppIosConfig) GoString() string {
	return s.String()
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppIosConfig) GetBundleId() *string {
	return s.BundleId
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppIosConfig) SetBundleId(v string) *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppIosConfig {
	s.BundleId = &v
	return s
}

func (s *QueryMappCenterAppResponseBodyQueryMappCenterAppResultMappCenterAppIosConfig) Validate() error {
	return dara.Validate(s)
}
