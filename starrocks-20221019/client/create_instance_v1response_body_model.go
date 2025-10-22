// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceV1ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateInstanceV1ResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *CreateInstanceV1ResponseBodyData) *CreateInstanceV1ResponseBody
	GetData() *CreateInstanceV1ResponseBodyData
	SetErrCode(v string) *CreateInstanceV1ResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateInstanceV1ResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CreateInstanceV1ResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateInstanceV1ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateInstanceV1ResponseBody
	GetSuccess() *bool
}

type CreateInstanceV1ResponseBody struct {
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                           `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *CreateInstanceV1ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [Region id should be select from set [cn-beijing, cn-hangzhou]]
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
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceV1ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceV1ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceV1ResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateInstanceV1ResponseBody) GetData() *CreateInstanceV1ResponseBodyData {
	return s.Data
}

func (s *CreateInstanceV1ResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateInstanceV1ResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateInstanceV1ResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateInstanceV1ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceV1ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateInstanceV1ResponseBody) SetAccessDeniedDetail(v string) *CreateInstanceV1ResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateInstanceV1ResponseBody) SetData(v *CreateInstanceV1ResponseBodyData) *CreateInstanceV1ResponseBody {
	s.Data = v
	return s
}

func (s *CreateInstanceV1ResponseBody) SetErrCode(v string) *CreateInstanceV1ResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateInstanceV1ResponseBody) SetErrMessage(v string) *CreateInstanceV1ResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateInstanceV1ResponseBody) SetHttpStatusCode(v int32) *CreateInstanceV1ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateInstanceV1ResponseBody) SetRequestId(v string) *CreateInstanceV1ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceV1ResponseBody) SetSuccess(v bool) *CreateInstanceV1ResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInstanceV1ResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateInstanceV1ResponseBodyData struct {
	// example:
	//
	// c-b25e21e243889XXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 241526000650XXX
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateInstanceV1ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceV1ResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateInstanceV1ResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceV1ResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateInstanceV1ResponseBodyData) SetInstanceId(v string) *CreateInstanceV1ResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceV1ResponseBodyData) SetOrderId(v int64) *CreateInstanceV1ResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateInstanceV1ResponseBodyData) Validate() error {
	return dara.Validate(s)
}
