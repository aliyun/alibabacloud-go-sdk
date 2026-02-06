// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBeebotIntentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeebotRequestId(v string) *DeleteBeebotIntentResponseBody
	GetBeebotRequestId() *string
	SetCode(v string) *DeleteBeebotIntentResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteBeebotIntentResponseBody
	GetHttpStatusCode() *int32
	SetIntentId(v int64) *DeleteBeebotIntentResponseBody
	GetIntentId() *int64
	SetMessage(v string) *DeleteBeebotIntentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteBeebotIntentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteBeebotIntentResponseBody
	GetSuccess() *bool
}

type DeleteBeebotIntentResponseBody struct {
	// example:
	//
	// 497CFAFF-48CC-161A-AD2C-252DED569037
	BeebotRequestId *string `json:"BeebotRequestId,omitempty" xml:"BeebotRequestId,omitempty"`
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
	// 10717802
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
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

func (s DeleteBeebotIntentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBeebotIntentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBeebotIntentResponseBody) GetBeebotRequestId() *string {
	return s.BeebotRequestId
}

func (s *DeleteBeebotIntentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteBeebotIntentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteBeebotIntentResponseBody) GetIntentId() *int64 {
	return s.IntentId
}

func (s *DeleteBeebotIntentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteBeebotIntentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBeebotIntentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteBeebotIntentResponseBody) SetBeebotRequestId(v string) *DeleteBeebotIntentResponseBody {
	s.BeebotRequestId = &v
	return s
}

func (s *DeleteBeebotIntentResponseBody) SetCode(v string) *DeleteBeebotIntentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBeebotIntentResponseBody) SetHttpStatusCode(v int32) *DeleteBeebotIntentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteBeebotIntentResponseBody) SetIntentId(v int64) *DeleteBeebotIntentResponseBody {
	s.IntentId = &v
	return s
}

func (s *DeleteBeebotIntentResponseBody) SetMessage(v string) *DeleteBeebotIntentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBeebotIntentResponseBody) SetRequestId(v string) *DeleteBeebotIntentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBeebotIntentResponseBody) SetSuccess(v bool) *DeleteBeebotIntentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBeebotIntentResponseBody) Validate() error {
	return dara.Validate(s)
}
