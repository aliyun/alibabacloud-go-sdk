// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMappCenterAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListMappCenterAppResult(v *ListMappCenterAppsResponseBodyListMappCenterAppResult) *ListMappCenterAppsResponseBody
	GetListMappCenterAppResult() *ListMappCenterAppsResponseBodyListMappCenterAppResult
	SetRequestId(v string) *ListMappCenterAppsResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListMappCenterAppsResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *ListMappCenterAppsResponseBody
	GetResultMessage() *string
}

type ListMappCenterAppsResponseBody struct {
	ListMappCenterAppResult *ListMappCenterAppsResponseBodyListMappCenterAppResult `json:"ListMappCenterAppResult,omitempty" xml:"ListMappCenterAppResult,omitempty" type:"Struct"`
	RequestId               *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode              *string                                                `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage           *string                                                `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMappCenterAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMappCenterAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMappCenterAppsResponseBody) GetListMappCenterAppResult() *ListMappCenterAppsResponseBodyListMappCenterAppResult {
	return s.ListMappCenterAppResult
}

func (s *ListMappCenterAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMappCenterAppsResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListMappCenterAppsResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListMappCenterAppsResponseBody) SetListMappCenterAppResult(v *ListMappCenterAppsResponseBodyListMappCenterAppResult) *ListMappCenterAppsResponseBody {
	s.ListMappCenterAppResult = v
	return s
}

func (s *ListMappCenterAppsResponseBody) SetRequestId(v string) *ListMappCenterAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMappCenterAppsResponseBody) SetResultCode(v string) *ListMappCenterAppsResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMappCenterAppsResponseBody) SetResultMessage(v string) *ListMappCenterAppsResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListMappCenterAppsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMappCenterAppsResponseBodyListMappCenterAppResult struct {
	MappCenterAppList []*ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList `json:"MappCenterAppList,omitempty" xml:"MappCenterAppList,omitempty" type:"Repeated"`
	ResultMsg         *string                                                                   `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success           *bool                                                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMappCenterAppsResponseBodyListMappCenterAppResult) String() string {
	return dara.Prettify(s)
}

func (s ListMappCenterAppsResponseBodyListMappCenterAppResult) GoString() string {
	return s.String()
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResult) GetMappCenterAppList() []*ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList {
	return s.MappCenterAppList
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResult) GetSuccess() *bool {
	return s.Success
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResult) SetMappCenterAppList(v []*ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) *ListMappCenterAppsResponseBodyListMappCenterAppResult {
	s.MappCenterAppList = v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResult) SetResultMsg(v string) *ListMappCenterAppsResponseBodyListMappCenterAppResult {
	s.ResultMsg = &v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResult) SetSuccess(v bool) *ListMappCenterAppsResponseBodyListMappCenterAppResult {
	s.Success = &v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResult) Validate() error {
	return dara.Validate(s)
}

type ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList struct {
	AndroidConfig *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListAndroidConfig `json:"AndroidConfig,omitempty" xml:"AndroidConfig,omitempty" type:"Struct"`
	AppDesc       *string                                                                              `json:"AppDesc,omitempty" xml:"AppDesc,omitempty"`
	AppIcon       *string                                                                              `json:"AppIcon,omitempty" xml:"AppIcon,omitempty"`
	AppId         *string                                                                              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName       *string                                                                              `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppSecret     *string                                                                              `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	Creator       *string                                                                              `json:"Creator,omitempty" xml:"Creator,omitempty"`
	GmtCreate     *string                                                                              `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified   *string                                                                              `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id            *int64                                                                               `json:"Id,omitempty" xml:"Id,omitempty"`
	IosConfig     *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListIosConfig     `json:"IosConfig,omitempty" xml:"IosConfig,omitempty" type:"Struct"`
	Modifier      *string                                                                              `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	MonitorJson   *string                                                                              `json:"MonitorJson,omitempty" xml:"MonitorJson,omitempty"`
	Status        *int64                                                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId      *string                                                                              `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Type          *int64                                                                               `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) String() string {
	return dara.Prettify(s)
}

func (s ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) GoString() string {
	return s.String()
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) GetAndroidConfig() *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListAndroidConfig {
	return s.AndroidConfig
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) GetAppDesc() *string {
	return s.AppDesc
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) GetAppIcon() *string {
	return s.AppIcon
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) GetAppId() *string {
	return s.AppId
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) GetAppName() *string {
	return s.AppName
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) GetAppSecret() *string {
	return s.AppSecret
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) GetCreator() *string {
	return s.Creator
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) GetId() *int64 {
	return s.Id
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) GetIosConfig() *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListIosConfig {
	return s.IosConfig
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) GetModifier() *string {
	return s.Modifier
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) GetMonitorJson() *string {
	return s.MonitorJson
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) GetStatus() *int64 {
	return s.Status
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) GetType() *int64 {
	return s.Type
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) SetAndroidConfig(v *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListAndroidConfig) *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList {
	s.AndroidConfig = v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) SetAppDesc(v string) *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList {
	s.AppDesc = &v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) SetAppIcon(v string) *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList {
	s.AppIcon = &v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) SetAppId(v string) *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList {
	s.AppId = &v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) SetAppName(v string) *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList {
	s.AppName = &v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) SetAppSecret(v string) *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList {
	s.AppSecret = &v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) SetCreator(v string) *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList {
	s.Creator = &v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) SetGmtCreate(v string) *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList {
	s.GmtCreate = &v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) SetGmtModified(v string) *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList {
	s.GmtModified = &v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) SetId(v int64) *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList {
	s.Id = &v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) SetIosConfig(v *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListIosConfig) *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList {
	s.IosConfig = v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) SetModifier(v string) *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList {
	s.Modifier = &v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) SetMonitorJson(v string) *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList {
	s.MonitorJson = &v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) SetStatus(v int64) *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList {
	s.Status = &v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) SetTenantId(v string) *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList {
	s.TenantId = &v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) SetType(v int64) *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList {
	s.Type = &v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppList) Validate() error {
	return dara.Validate(s)
}

type ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListAndroidConfig struct {
	CertRSA     *string `json:"CertRSA,omitempty" xml:"CertRSA,omitempty"`
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
}

func (s ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListAndroidConfig) String() string {
	return dara.Prettify(s)
}

func (s ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListAndroidConfig) GoString() string {
	return s.String()
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListAndroidConfig) GetCertRSA() *string {
	return s.CertRSA
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListAndroidConfig) GetPackageName() *string {
	return s.PackageName
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListAndroidConfig) SetCertRSA(v string) *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListAndroidConfig {
	s.CertRSA = &v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListAndroidConfig) SetPackageName(v string) *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListAndroidConfig {
	s.PackageName = &v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListAndroidConfig) Validate() error {
	return dara.Validate(s)
}

type ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListIosConfig struct {
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
}

func (s ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListIosConfig) String() string {
	return dara.Prettify(s)
}

func (s ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListIosConfig) GoString() string {
	return s.String()
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListIosConfig) GetBundleId() *string {
	return s.BundleId
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListIosConfig) SetBundleId(v string) *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListIosConfig {
	s.BundleId = &v
	return s
}

func (s *ListMappCenterAppsResponseBodyListMappCenterAppResultMappCenterAppListIosConfig) Validate() error {
	return dara.Validate(s)
}
