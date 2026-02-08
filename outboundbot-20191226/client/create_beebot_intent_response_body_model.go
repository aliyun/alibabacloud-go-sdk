// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBeebotIntentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeebotRequestId(v string) *CreateBeebotIntentResponseBody
	GetBeebotRequestId() *string
	SetCode(v string) *CreateBeebotIntentResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateBeebotIntentResponseBody
	GetHttpStatusCode() *int32
	SetIntentId(v int64) *CreateBeebotIntentResponseBody
	GetIntentId() *int64
	SetMessage(v string) *CreateBeebotIntentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateBeebotIntentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBeebotIntentResponseBody
	GetSuccess() *bool
}

type CreateBeebotIntentResponseBody struct {
	// Internal request ID
	//
	// example:
	//
	// 497CFAFF-48CC-161A-AD2C-252DED569037
	BeebotRequestId *string `json:"BeebotRequestId,omitempty" xml:"BeebotRequestId,omitempty"`
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Intent ID
	//
	// example:
	//
	// 10717802
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBeebotIntentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBeebotIntentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBeebotIntentResponseBody) GetBeebotRequestId() *string {
	return s.BeebotRequestId
}

func (s *CreateBeebotIntentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateBeebotIntentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateBeebotIntentResponseBody) GetIntentId() *int64 {
	return s.IntentId
}

func (s *CreateBeebotIntentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateBeebotIntentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBeebotIntentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBeebotIntentResponseBody) SetBeebotRequestId(v string) *CreateBeebotIntentResponseBody {
	s.BeebotRequestId = &v
	return s
}

func (s *CreateBeebotIntentResponseBody) SetCode(v string) *CreateBeebotIntentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBeebotIntentResponseBody) SetHttpStatusCode(v int32) *CreateBeebotIntentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateBeebotIntentResponseBody) SetIntentId(v int64) *CreateBeebotIntentResponseBody {
	s.IntentId = &v
	return s
}

func (s *CreateBeebotIntentResponseBody) SetMessage(v string) *CreateBeebotIntentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBeebotIntentResponseBody) SetRequestId(v string) *CreateBeebotIntentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBeebotIntentResponseBody) SetSuccess(v bool) *CreateBeebotIntentResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBeebotIntentResponseBody) Validate() error {
	return dara.Validate(s)
}
