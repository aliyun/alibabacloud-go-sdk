// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataServiceAppMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDataServiceAppMemberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateDataServiceAppMemberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateDataServiceAppMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDataServiceAppMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataServiceAppMemberResponseBody
	GetSuccess() *bool
}

type UpdateDataServiceAppMemberResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataServiceAppMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataServiceAppMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataServiceAppMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDataServiceAppMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateDataServiceAppMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDataServiceAppMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataServiceAppMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataServiceAppMemberResponseBody) SetCode(v string) *UpdateDataServiceAppMemberResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDataServiceAppMemberResponseBody) SetHttpStatusCode(v int32) *UpdateDataServiceAppMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDataServiceAppMemberResponseBody) SetMessage(v string) *UpdateDataServiceAppMemberResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDataServiceAppMemberResponseBody) SetRequestId(v string) *UpdateDataServiceAppMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataServiceAppMemberResponseBody) SetSuccess(v bool) *UpdateDataServiceAppMemberResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataServiceAppMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
