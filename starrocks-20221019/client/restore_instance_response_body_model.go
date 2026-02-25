// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RestoreInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *RestoreInstanceResponseBodyData) *RestoreInstanceResponseBody
	GetData() *RestoreInstanceResponseBodyData
	SetErrCode(v string) *RestoreInstanceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *RestoreInstanceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *RestoreInstanceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *RestoreInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RestoreInstanceResponseBody
	GetSuccess() *bool
}

type RestoreInstanceResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                          `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *RestoreInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// null
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE74XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RestoreInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestoreInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RestoreInstanceResponseBody) GetData() *RestoreInstanceResponseBodyData {
	return s.Data
}

func (s *RestoreInstanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *RestoreInstanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *RestoreInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RestoreInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestoreInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RestoreInstanceResponseBody) SetAccessDeniedDetail(v string) *RestoreInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RestoreInstanceResponseBody) SetData(v *RestoreInstanceResponseBodyData) *RestoreInstanceResponseBody {
	s.Data = v
	return s
}

func (s *RestoreInstanceResponseBody) SetErrCode(v string) *RestoreInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RestoreInstanceResponseBody) SetErrMessage(v string) *RestoreInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *RestoreInstanceResponseBody) SetHttpStatusCode(v int32) *RestoreInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RestoreInstanceResponseBody) SetRequestId(v string) *RestoreInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestoreInstanceResponseBody) SetSuccess(v bool) *RestoreInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *RestoreInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RestoreInstanceResponseBodyData struct {
	// example:
	//
	// c-b25e21e243889XXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 241526000650XXX
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s RestoreInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RestoreInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RestoreInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RestoreInstanceResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *RestoreInstanceResponseBodyData) SetInstanceId(v string) *RestoreInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *RestoreInstanceResponseBodyData) SetOrderId(v int64) *RestoreInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *RestoreInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
