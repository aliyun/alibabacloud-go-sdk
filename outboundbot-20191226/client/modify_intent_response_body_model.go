// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIntentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyIntentResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyIntentResponseBody
	GetHttpStatusCode() *int32
	SetIntentId(v string) *ModifyIntentResponseBody
	GetIntentId() *string
	SetMessage(v string) *ModifyIntentResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyIntentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyIntentResponseBody
	GetSuccess() *bool
}

type ModifyIntentResponseBody struct {
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
	// 3b9a2b33-50d4-4576-8c68-22498f4bf731
	IntentId *string `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyIntentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIntentResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIntentResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyIntentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyIntentResponseBody) GetIntentId() *string {
	return s.IntentId
}

func (s *ModifyIntentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyIntentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIntentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyIntentResponseBody) SetCode(v string) *ModifyIntentResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyIntentResponseBody) SetHttpStatusCode(v int32) *ModifyIntentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyIntentResponseBody) SetIntentId(v string) *ModifyIntentResponseBody {
	s.IntentId = &v
	return s
}

func (s *ModifyIntentResponseBody) SetMessage(v string) *ModifyIntentResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyIntentResponseBody) SetRequestId(v string) *ModifyIntentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIntentResponseBody) SetSuccess(v bool) *ModifyIntentResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyIntentResponseBody) Validate() error {
	return dara.Validate(s)
}
