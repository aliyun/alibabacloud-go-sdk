// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceFeatureGateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetInstanceFeatureGateResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *GetInstanceFeatureGateResponseBodyData) *GetInstanceFeatureGateResponseBody
	GetData() *GetInstanceFeatureGateResponseBodyData
	SetErrCode(v string) *GetInstanceFeatureGateResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetInstanceFeatureGateResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *GetInstanceFeatureGateResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetInstanceFeatureGateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceFeatureGateResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *GetInstanceFeatureGateResponseBody
	GetTotal() *int32
}

type GetInstanceFeatureGateResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                                 `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *GetInstanceFeatureGateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// B67D142D-D54E-184F-A306-22BDC01B2XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetInstanceFeatureGateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceFeatureGateResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceFeatureGateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetInstanceFeatureGateResponseBody) GetData() *GetInstanceFeatureGateResponseBodyData {
	return s.Data
}

func (s *GetInstanceFeatureGateResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetInstanceFeatureGateResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetInstanceFeatureGateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceFeatureGateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceFeatureGateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceFeatureGateResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *GetInstanceFeatureGateResponseBody) SetAccessDeniedDetail(v string) *GetInstanceFeatureGateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBody) SetData(v *GetInstanceFeatureGateResponseBodyData) *GetInstanceFeatureGateResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceFeatureGateResponseBody) SetErrCode(v string) *GetInstanceFeatureGateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBody) SetErrMessage(v string) *GetInstanceFeatureGateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBody) SetHttpStatusCode(v int32) *GetInstanceFeatureGateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBody) SetRequestId(v string) *GetInstanceFeatureGateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBody) SetSuccess(v bool) *GetInstanceFeatureGateResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBody) SetTotal(v int32) *GetInstanceFeatureGateResponseBody {
	s.Total = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceFeatureGateResponseBodyData struct {
	// example:
	//
	// true
	ConsoleControlRestart *bool `json:"ConsoleControlRestart,omitempty" xml:"ConsoleControlRestart,omitempty"`
	// example:
	//
	// true
	EnableManageMv *bool `json:"EnableManageMv,omitempty" xml:"EnableManageMv,omitempty"`
	// example:
	//
	// true
	FullManagedSecurityGroup *bool `json:"FullManagedSecurityGroup,omitempty" xml:"FullManagedSecurityGroup,omitempty"`
	// example:
	//
	// true
	MountDlfMetaToken     *bool     `json:"MountDlfMetaToken,omitempty" xml:"MountDlfMetaToken,omitempty"`
	SupportAddConfigTypes []*string `json:"SupportAddConfigTypes,omitempty" xml:"SupportAddConfigTypes,omitempty" type:"Repeated"`
	// example:
	//
	// true
	SupportBackup *bool `json:"SupportBackup,omitempty" xml:"SupportBackup,omitempty"`
	// example:
	//
	// true
	SupportCreateAgent *bool `json:"SupportCreateAgent,omitempty" xml:"SupportCreateAgent,omitempty"`
	// example:
	//
	// true
	SupportCreateNonStandardNodeGroup *bool `json:"SupportCreateNonStandardNodeGroup,omitempty" xml:"SupportCreateNonStandardNodeGroup,omitempty"`
	// example:
	//
	// true
	SupportEed *bool `json:"SupportEed,omitempty" xml:"SupportEed,omitempty"`
	// example:
	//
	// true
	SupportEnableAi *bool `json:"SupportEnableAi,omitempty" xml:"SupportEnableAi,omitempty"`
	// example:
	//
	// true
	SupportEnableSSL *bool `json:"SupportEnableSSL,omitempty" xml:"SupportEnableSSL,omitempty"`
	// example:
	//
	// true
	SupportFastModeModifyConfig *bool `json:"SupportFastModeModifyConfig,omitempty" xml:"SupportFastModeModifyConfig,omitempty"`
	// example:
	//
	// true
	SupportFastModeModifyResource *bool `json:"SupportFastModeModifyResource,omitempty" xml:"SupportFastModeModifyResource,omitempty"`
	// example:
	//
	// true
	SupportFastRestart *bool `json:"SupportFastRestart,omitempty" xml:"SupportFastRestart,omitempty"`
	// example:
	//
	// true
	SupportFeGateway *bool `json:"SupportFeGateway,omitempty" xml:"SupportFeGateway,omitempty"`
	// example:
	//
	// true
	SupportHostAlias *bool `json:"SupportHostAlias,omitempty" xml:"SupportHostAlias,omitempty"`
	// example:
	//
	// true
	SupportModifyTimezone *bool `json:"SupportModifyTimezone,omitempty" xml:"SupportModifyTimezone,omitempty"`
	// example:
	//
	// true
	SupportMultiAZ *bool `json:"SupportMultiAZ,omitempty" xml:"SupportMultiAZ,omitempty"`
	// example:
	//
	// true
	UseComputeNode           *bool `json:"UseComputeNode,omitempty" xml:"UseComputeNode,omitempty"`
	SupportCompactionService *bool `json:"supportCompactionService,omitempty" xml:"supportCompactionService,omitempty"`
}

func (s GetInstanceFeatureGateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceFeatureGateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceFeatureGateResponseBodyData) GetConsoleControlRestart() *bool {
	return s.ConsoleControlRestart
}

func (s *GetInstanceFeatureGateResponseBodyData) GetEnableManageMv() *bool {
	return s.EnableManageMv
}

func (s *GetInstanceFeatureGateResponseBodyData) GetFullManagedSecurityGroup() *bool {
	return s.FullManagedSecurityGroup
}

func (s *GetInstanceFeatureGateResponseBodyData) GetMountDlfMetaToken() *bool {
	return s.MountDlfMetaToken
}

func (s *GetInstanceFeatureGateResponseBodyData) GetSupportAddConfigTypes() []*string {
	return s.SupportAddConfigTypes
}

func (s *GetInstanceFeatureGateResponseBodyData) GetSupportBackup() *bool {
	return s.SupportBackup
}

func (s *GetInstanceFeatureGateResponseBodyData) GetSupportCreateAgent() *bool {
	return s.SupportCreateAgent
}

func (s *GetInstanceFeatureGateResponseBodyData) GetSupportCreateNonStandardNodeGroup() *bool {
	return s.SupportCreateNonStandardNodeGroup
}

func (s *GetInstanceFeatureGateResponseBodyData) GetSupportEed() *bool {
	return s.SupportEed
}

func (s *GetInstanceFeatureGateResponseBodyData) GetSupportEnableAi() *bool {
	return s.SupportEnableAi
}

func (s *GetInstanceFeatureGateResponseBodyData) GetSupportEnableSSL() *bool {
	return s.SupportEnableSSL
}

func (s *GetInstanceFeatureGateResponseBodyData) GetSupportFastModeModifyConfig() *bool {
	return s.SupportFastModeModifyConfig
}

func (s *GetInstanceFeatureGateResponseBodyData) GetSupportFastModeModifyResource() *bool {
	return s.SupportFastModeModifyResource
}

func (s *GetInstanceFeatureGateResponseBodyData) GetSupportFastRestart() *bool {
	return s.SupportFastRestart
}

func (s *GetInstanceFeatureGateResponseBodyData) GetSupportFeGateway() *bool {
	return s.SupportFeGateway
}

func (s *GetInstanceFeatureGateResponseBodyData) GetSupportHostAlias() *bool {
	return s.SupportHostAlias
}

func (s *GetInstanceFeatureGateResponseBodyData) GetSupportModifyTimezone() *bool {
	return s.SupportModifyTimezone
}

func (s *GetInstanceFeatureGateResponseBodyData) GetSupportMultiAZ() *bool {
	return s.SupportMultiAZ
}

func (s *GetInstanceFeatureGateResponseBodyData) GetUseComputeNode() *bool {
	return s.UseComputeNode
}

func (s *GetInstanceFeatureGateResponseBodyData) GetSupportCompactionService() *bool {
	return s.SupportCompactionService
}

func (s *GetInstanceFeatureGateResponseBodyData) SetConsoleControlRestart(v bool) *GetInstanceFeatureGateResponseBodyData {
	s.ConsoleControlRestart = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBodyData) SetEnableManageMv(v bool) *GetInstanceFeatureGateResponseBodyData {
	s.EnableManageMv = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBodyData) SetFullManagedSecurityGroup(v bool) *GetInstanceFeatureGateResponseBodyData {
	s.FullManagedSecurityGroup = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBodyData) SetMountDlfMetaToken(v bool) *GetInstanceFeatureGateResponseBodyData {
	s.MountDlfMetaToken = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBodyData) SetSupportAddConfigTypes(v []*string) *GetInstanceFeatureGateResponseBodyData {
	s.SupportAddConfigTypes = v
	return s
}

func (s *GetInstanceFeatureGateResponseBodyData) SetSupportBackup(v bool) *GetInstanceFeatureGateResponseBodyData {
	s.SupportBackup = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBodyData) SetSupportCreateAgent(v bool) *GetInstanceFeatureGateResponseBodyData {
	s.SupportCreateAgent = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBodyData) SetSupportCreateNonStandardNodeGroup(v bool) *GetInstanceFeatureGateResponseBodyData {
	s.SupportCreateNonStandardNodeGroup = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBodyData) SetSupportEed(v bool) *GetInstanceFeatureGateResponseBodyData {
	s.SupportEed = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBodyData) SetSupportEnableAi(v bool) *GetInstanceFeatureGateResponseBodyData {
	s.SupportEnableAi = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBodyData) SetSupportEnableSSL(v bool) *GetInstanceFeatureGateResponseBodyData {
	s.SupportEnableSSL = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBodyData) SetSupportFastModeModifyConfig(v bool) *GetInstanceFeatureGateResponseBodyData {
	s.SupportFastModeModifyConfig = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBodyData) SetSupportFastModeModifyResource(v bool) *GetInstanceFeatureGateResponseBodyData {
	s.SupportFastModeModifyResource = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBodyData) SetSupportFastRestart(v bool) *GetInstanceFeatureGateResponseBodyData {
	s.SupportFastRestart = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBodyData) SetSupportFeGateway(v bool) *GetInstanceFeatureGateResponseBodyData {
	s.SupportFeGateway = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBodyData) SetSupportHostAlias(v bool) *GetInstanceFeatureGateResponseBodyData {
	s.SupportHostAlias = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBodyData) SetSupportModifyTimezone(v bool) *GetInstanceFeatureGateResponseBodyData {
	s.SupportModifyTimezone = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBodyData) SetSupportMultiAZ(v bool) *GetInstanceFeatureGateResponseBodyData {
	s.SupportMultiAZ = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBodyData) SetUseComputeNode(v bool) *GetInstanceFeatureGateResponseBodyData {
	s.UseComputeNode = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBodyData) SetSupportCompactionService(v bool) *GetInstanceFeatureGateResponseBodyData {
	s.SupportCompactionService = &v
	return s
}

func (s *GetInstanceFeatureGateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
