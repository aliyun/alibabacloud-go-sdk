// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDataServiceAppMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveDataServiceAppMemberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RemoveDataServiceAppMemberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RemoveDataServiceAppMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveDataServiceAppMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveDataServiceAppMemberResponseBody
	GetSuccess() *bool
}

type RemoveDataServiceAppMemberResponseBody struct {
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

func (s RemoveDataServiceAppMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveDataServiceAppMemberResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDataServiceAppMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveDataServiceAppMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemoveDataServiceAppMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveDataServiceAppMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveDataServiceAppMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveDataServiceAppMemberResponseBody) SetCode(v string) *RemoveDataServiceAppMemberResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveDataServiceAppMemberResponseBody) SetHttpStatusCode(v int32) *RemoveDataServiceAppMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveDataServiceAppMemberResponseBody) SetMessage(v string) *RemoveDataServiceAppMemberResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveDataServiceAppMemberResponseBody) SetRequestId(v string) *RemoveDataServiceAppMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDataServiceAppMemberResponseBody) SetSuccess(v bool) *RemoveDataServiceAppMemberResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveDataServiceAppMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
