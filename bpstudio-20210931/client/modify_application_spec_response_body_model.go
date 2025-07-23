// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyApplicationSpecResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ModifyApplicationSpecResponseBody
	GetCode() *string
	SetData(v string) *ModifyApplicationSpecResponseBody
	GetData() *string
	SetMessage(v string) *ModifyApplicationSpecResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyApplicationSpecResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifyApplicationSpecResponseBody
	GetSuccess() *string
}

type ModifyApplicationSpecResponseBody struct {
	// example:
	//
	// {
	//
	//     "PolicyType": "",
	//
	//     "AuthPrincipalOwnerId": "",
	//
	//     "EncodedDiagnosticMessage": "",
	//
	//     "AuthPrincipalType": "",
	//
	//     "AuthPrincipalDisplayName": "",
	//
	//     "NoPermissionType": "",
	//
	//     "AuthAction": ""
	//
	//   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// -1
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A3488F1D-xxxx-xxxx-xxxx-5374BA0F3562
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyApplicationSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApplicationSpecResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyApplicationSpecResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyApplicationSpecResponseBody) GetData() *string {
	return s.Data
}

func (s *ModifyApplicationSpecResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyApplicationSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApplicationSpecResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifyApplicationSpecResponseBody) SetAccessDeniedDetail(v string) *ModifyApplicationSpecResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyApplicationSpecResponseBody) SetCode(v string) *ModifyApplicationSpecResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyApplicationSpecResponseBody) SetData(v string) *ModifyApplicationSpecResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyApplicationSpecResponseBody) SetMessage(v string) *ModifyApplicationSpecResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyApplicationSpecResponseBody) SetRequestId(v string) *ModifyApplicationSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApplicationSpecResponseBody) SetSuccess(v string) *ModifyApplicationSpecResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyApplicationSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
