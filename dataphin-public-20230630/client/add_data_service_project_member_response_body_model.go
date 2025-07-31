// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataServiceProjectMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddDataServiceProjectMemberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddDataServiceProjectMemberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddDataServiceProjectMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddDataServiceProjectMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddDataServiceProjectMemberResponseBody
	GetSuccess() *bool
}

type AddDataServiceProjectMemberResponseBody struct {
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

func (s AddDataServiceProjectMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDataServiceProjectMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddDataServiceProjectMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddDataServiceProjectMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddDataServiceProjectMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddDataServiceProjectMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDataServiceProjectMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddDataServiceProjectMemberResponseBody) SetCode(v string) *AddDataServiceProjectMemberResponseBody {
	s.Code = &v
	return s
}

func (s *AddDataServiceProjectMemberResponseBody) SetHttpStatusCode(v int32) *AddDataServiceProjectMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddDataServiceProjectMemberResponseBody) SetMessage(v string) *AddDataServiceProjectMemberResponseBody {
	s.Message = &v
	return s
}

func (s *AddDataServiceProjectMemberResponseBody) SetRequestId(v string) *AddDataServiceProjectMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDataServiceProjectMemberResponseBody) SetSuccess(v bool) *AddDataServiceProjectMemberResponseBody {
	s.Success = &v
	return s
}

func (s *AddDataServiceProjectMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
