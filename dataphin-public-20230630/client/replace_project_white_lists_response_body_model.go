// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceProjectWhiteListsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReplaceProjectWhiteListsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ReplaceProjectWhiteListsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ReplaceProjectWhiteListsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReplaceProjectWhiteListsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReplaceProjectWhiteListsResponseBody
	GetSuccess() *bool
}

type ReplaceProjectWhiteListsResponseBody struct {
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

func (s ReplaceProjectWhiteListsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplaceProjectWhiteListsResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceProjectWhiteListsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReplaceProjectWhiteListsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ReplaceProjectWhiteListsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReplaceProjectWhiteListsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplaceProjectWhiteListsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReplaceProjectWhiteListsResponseBody) SetCode(v string) *ReplaceProjectWhiteListsResponseBody {
	s.Code = &v
	return s
}

func (s *ReplaceProjectWhiteListsResponseBody) SetHttpStatusCode(v int32) *ReplaceProjectWhiteListsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReplaceProjectWhiteListsResponseBody) SetMessage(v string) *ReplaceProjectWhiteListsResponseBody {
	s.Message = &v
	return s
}

func (s *ReplaceProjectWhiteListsResponseBody) SetRequestId(v string) *ReplaceProjectWhiteListsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceProjectWhiteListsResponseBody) SetSuccess(v bool) *ReplaceProjectWhiteListsResponseBody {
	s.Success = &v
	return s
}

func (s *ReplaceProjectWhiteListsResponseBody) Validate() error {
	return dara.Validate(s)
}
