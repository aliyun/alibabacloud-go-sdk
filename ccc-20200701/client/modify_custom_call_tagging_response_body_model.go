// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomCallTaggingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyCustomCallTaggingResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyCustomCallTaggingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyCustomCallTaggingResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyCustomCallTaggingResponseBody
	GetRequestId() *string
}

type ModifyCustomCallTaggingResponseBody struct {
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

func (s ModifyCustomCallTaggingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomCallTaggingResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCustomCallTaggingResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyCustomCallTaggingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyCustomCallTaggingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyCustomCallTaggingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCustomCallTaggingResponseBody) SetCode(v string) *ModifyCustomCallTaggingResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyCustomCallTaggingResponseBody) SetHttpStatusCode(v int32) *ModifyCustomCallTaggingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyCustomCallTaggingResponseBody) SetMessage(v string) *ModifyCustomCallTaggingResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyCustomCallTaggingResponseBody) SetRequestId(v string) *ModifyCustomCallTaggingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCustomCallTaggingResponseBody) Validate() error {
	return dara.Validate(s)
}
