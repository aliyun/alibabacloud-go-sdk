// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstancePrice4ModifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryInstancePrice4ModifyResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryInstancePrice4ModifyResponseBody
	GetCode() *string
	SetData(v string) *QueryInstancePrice4ModifyResponseBody
	GetData() *string
	SetMessage(v string) *QueryInstancePrice4ModifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryInstancePrice4ModifyResponseBody
	GetRequestId() *string
	SetSuccess(v string) *QueryInstancePrice4ModifyResponseBody
	GetSuccess() *string
}

type QueryInstancePrice4ModifyResponseBody struct {
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
	// taskId
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9656C816-1E9A-58D2-86D5-710678D61AF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryInstancePrice4ModifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryInstancePrice4ModifyResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstancePrice4ModifyResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryInstancePrice4ModifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryInstancePrice4ModifyResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryInstancePrice4ModifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryInstancePrice4ModifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryInstancePrice4ModifyResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *QueryInstancePrice4ModifyResponseBody) SetAccessDeniedDetail(v string) *QueryInstancePrice4ModifyResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryInstancePrice4ModifyResponseBody) SetCode(v string) *QueryInstancePrice4ModifyResponseBody {
	s.Code = &v
	return s
}

func (s *QueryInstancePrice4ModifyResponseBody) SetData(v string) *QueryInstancePrice4ModifyResponseBody {
	s.Data = &v
	return s
}

func (s *QueryInstancePrice4ModifyResponseBody) SetMessage(v string) *QueryInstancePrice4ModifyResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInstancePrice4ModifyResponseBody) SetRequestId(v string) *QueryInstancePrice4ModifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInstancePrice4ModifyResponseBody) SetSuccess(v string) *QueryInstancePrice4ModifyResponseBody {
	s.Success = &v
	return s
}

func (s *QueryInstancePrice4ModifyResponseBody) Validate() error {
	return dara.Validate(s)
}
