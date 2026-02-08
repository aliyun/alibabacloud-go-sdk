// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBeebotIntentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeebotRequestId(v string) *ModifyBeebotIntentResponseBody
	GetBeebotRequestId() *string
	SetCode(v string) *ModifyBeebotIntentResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyBeebotIntentResponseBody
	GetHttpStatusCode() *int32
	SetIntentId(v int64) *ModifyBeebotIntentResponseBody
	GetIntentId() *int64
	SetMessage(v string) *ModifyBeebotIntentResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyBeebotIntentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyBeebotIntentResponseBody
	GetSuccess() *bool
}

type ModifyBeebotIntentResponseBody struct {
	// Internal request ID
	//
	// example:
	//
	// A1F21BF2-CB21-1968-8039-C74699E7DDEB
	BeebotRequestId *string `json:"BeebotRequestId,omitempty" xml:"BeebotRequestId,omitempty"`
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP return code
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
	// 8785D26A-7406-50A1-9653-1313C292E23B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyBeebotIntentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBeebotIntentResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBeebotIntentResponseBody) GetBeebotRequestId() *string {
	return s.BeebotRequestId
}

func (s *ModifyBeebotIntentResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyBeebotIntentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyBeebotIntentResponseBody) GetIntentId() *int64 {
	return s.IntentId
}

func (s *ModifyBeebotIntentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyBeebotIntentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBeebotIntentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyBeebotIntentResponseBody) SetBeebotRequestId(v string) *ModifyBeebotIntentResponseBody {
	s.BeebotRequestId = &v
	return s
}

func (s *ModifyBeebotIntentResponseBody) SetCode(v string) *ModifyBeebotIntentResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyBeebotIntentResponseBody) SetHttpStatusCode(v int32) *ModifyBeebotIntentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyBeebotIntentResponseBody) SetIntentId(v int64) *ModifyBeebotIntentResponseBody {
	s.IntentId = &v
	return s
}

func (s *ModifyBeebotIntentResponseBody) SetMessage(v string) *ModifyBeebotIntentResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyBeebotIntentResponseBody) SetRequestId(v string) *ModifyBeebotIntentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBeebotIntentResponseBody) SetSuccess(v bool) *ModifyBeebotIntentResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyBeebotIntentResponseBody) Validate() error {
	return dara.Validate(s)
}
