// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomCallTaggingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCustomCallTaggingResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteCustomCallTaggingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteCustomCallTaggingResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCustomCallTaggingResponseBody
	GetRequestId() *string
}

type DeleteCustomCallTaggingResponseBody struct {
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
	// BA03159C-E808-4FF1-B27E-A61B6E888D7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomCallTaggingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomCallTaggingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomCallTaggingResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCustomCallTaggingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteCustomCallTaggingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCustomCallTaggingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomCallTaggingResponseBody) SetCode(v string) *DeleteCustomCallTaggingResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomCallTaggingResponseBody) SetHttpStatusCode(v int32) *DeleteCustomCallTaggingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteCustomCallTaggingResponseBody) SetMessage(v string) *DeleteCustomCallTaggingResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomCallTaggingResponseBody) SetRequestId(v string) *DeleteCustomCallTaggingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomCallTaggingResponseBody) Validate() error {
	return dara.Validate(s)
}
