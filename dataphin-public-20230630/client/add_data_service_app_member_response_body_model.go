// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataServiceAppMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddDataServiceAppMemberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddDataServiceAppMemberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddDataServiceAppMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddDataServiceAppMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddDataServiceAppMemberResponseBody
	GetSuccess() *bool
}

type AddDataServiceAppMemberResponseBody struct {
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

func (s AddDataServiceAppMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDataServiceAppMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddDataServiceAppMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddDataServiceAppMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddDataServiceAppMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddDataServiceAppMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDataServiceAppMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddDataServiceAppMemberResponseBody) SetCode(v string) *AddDataServiceAppMemberResponseBody {
	s.Code = &v
	return s
}

func (s *AddDataServiceAppMemberResponseBody) SetHttpStatusCode(v int32) *AddDataServiceAppMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddDataServiceAppMemberResponseBody) SetMessage(v string) *AddDataServiceAppMemberResponseBody {
	s.Message = &v
	return s
}

func (s *AddDataServiceAppMemberResponseBody) SetRequestId(v string) *AddDataServiceAppMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDataServiceAppMemberResponseBody) SetSuccess(v bool) *AddDataServiceAppMemberResponseBody {
	s.Success = &v
	return s
}

func (s *AddDataServiceAppMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
