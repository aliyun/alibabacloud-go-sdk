// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeGroupFeatureGateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetNodeGroupFeatureGateResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *GetNodeGroupFeatureGateResponseBodyData) *GetNodeGroupFeatureGateResponseBody
	GetData() *GetNodeGroupFeatureGateResponseBodyData
	SetErrCode(v string) *GetNodeGroupFeatureGateResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetNodeGroupFeatureGateResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *GetNodeGroupFeatureGateResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetNodeGroupFeatureGateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetNodeGroupFeatureGateResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *GetNodeGroupFeatureGateResponseBody
	GetTotal() *int32
}

type GetNodeGroupFeatureGateResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                                  `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *GetNodeGroupFeatureGateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetNodeGroupFeatureGateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodeGroupFeatureGateResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeGroupFeatureGateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetNodeGroupFeatureGateResponseBody) GetData() *GetNodeGroupFeatureGateResponseBodyData {
	return s.Data
}

func (s *GetNodeGroupFeatureGateResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetNodeGroupFeatureGateResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetNodeGroupFeatureGateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetNodeGroupFeatureGateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNodeGroupFeatureGateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNodeGroupFeatureGateResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *GetNodeGroupFeatureGateResponseBody) SetAccessDeniedDetail(v string) *GetNodeGroupFeatureGateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetNodeGroupFeatureGateResponseBody) SetData(v *GetNodeGroupFeatureGateResponseBodyData) *GetNodeGroupFeatureGateResponseBody {
	s.Data = v
	return s
}

func (s *GetNodeGroupFeatureGateResponseBody) SetErrCode(v string) *GetNodeGroupFeatureGateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetNodeGroupFeatureGateResponseBody) SetErrMessage(v string) *GetNodeGroupFeatureGateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetNodeGroupFeatureGateResponseBody) SetHttpStatusCode(v int32) *GetNodeGroupFeatureGateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetNodeGroupFeatureGateResponseBody) SetRequestId(v string) *GetNodeGroupFeatureGateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeGroupFeatureGateResponseBody) SetSuccess(v bool) *GetNodeGroupFeatureGateResponseBody {
	s.Success = &v
	return s
}

func (s *GetNodeGroupFeatureGateResponseBody) SetTotal(v int32) *GetNodeGroupFeatureGateResponseBody {
	s.Total = &v
	return s
}

func (s *GetNodeGroupFeatureGateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNodeGroupFeatureGateResponseBodyData struct {
	// example:
	//
	// true
	NeedRestartAfterModifyDiskSize *bool `json:"NeedRestartAfterModifyDiskSize,omitempty" xml:"NeedRestartAfterModifyDiskSize,omitempty"`
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
	SupportModifySpecType *bool `json:"SupportModifySpecType,omitempty" xml:"SupportModifySpecType,omitempty"`
}

func (s GetNodeGroupFeatureGateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetNodeGroupFeatureGateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNodeGroupFeatureGateResponseBodyData) GetNeedRestartAfterModifyDiskSize() *bool {
	return s.NeedRestartAfterModifyDiskSize
}

func (s *GetNodeGroupFeatureGateResponseBodyData) GetSupportFastModeModifyResource() *bool {
	return s.SupportFastModeModifyResource
}

func (s *GetNodeGroupFeatureGateResponseBodyData) GetSupportFastRestart() *bool {
	return s.SupportFastRestart
}

func (s *GetNodeGroupFeatureGateResponseBodyData) GetSupportModifySpecType() *bool {
	return s.SupportModifySpecType
}

func (s *GetNodeGroupFeatureGateResponseBodyData) SetNeedRestartAfterModifyDiskSize(v bool) *GetNodeGroupFeatureGateResponseBodyData {
	s.NeedRestartAfterModifyDiskSize = &v
	return s
}

func (s *GetNodeGroupFeatureGateResponseBodyData) SetSupportFastModeModifyResource(v bool) *GetNodeGroupFeatureGateResponseBodyData {
	s.SupportFastModeModifyResource = &v
	return s
}

func (s *GetNodeGroupFeatureGateResponseBodyData) SetSupportFastRestart(v bool) *GetNodeGroupFeatureGateResponseBodyData {
	s.SupportFastRestart = &v
	return s
}

func (s *GetNodeGroupFeatureGateResponseBodyData) SetSupportModifySpecType(v bool) *GetNodeGroupFeatureGateResponseBodyData {
	s.SupportModifySpecType = &v
	return s
}

func (s *GetNodeGroupFeatureGateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
