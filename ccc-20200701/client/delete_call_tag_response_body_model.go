// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCallTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCallTagResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteCallTagResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteCallTagResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCallTagResponseBody
	GetRequestId() *string
}

type DeleteCallTagResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCallTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCallTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCallTagResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCallTagResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteCallTagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCallTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCallTagResponseBody) SetCode(v string) *DeleteCallTagResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCallTagResponseBody) SetHttpStatusCode(v int32) *DeleteCallTagResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteCallTagResponseBody) SetMessage(v string) *DeleteCallTagResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCallTagResponseBody) SetRequestId(v string) *DeleteCallTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCallTagResponseBody) Validate() error {
	return dara.Validate(s)
}
