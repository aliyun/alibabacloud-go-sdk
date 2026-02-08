// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBeebotIntentLgfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeebotRequestId(v string) *DeleteBeebotIntentLgfResponseBody
	GetBeebotRequestId() *string
	SetCode(v string) *DeleteBeebotIntentLgfResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteBeebotIntentLgfResponseBody
	GetHttpStatusCode() *int32
	SetLgfId(v int64) *DeleteBeebotIntentLgfResponseBody
	GetLgfId() *int64
	SetMessage(v string) *DeleteBeebotIntentLgfResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteBeebotIntentLgfResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteBeebotIntentLgfResponseBody
	GetSuccess() *bool
}

type DeleteBeebotIntentLgfResponseBody struct {
	// Internal request ID
	//
	// example:
	//
	// 0B219FCB-EC71-1F08-BB1B-0E87C20158C8
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
	// Utterance template ID
	//
	// example:
	//
	// 5666117
	LgfId *int64 `json:"LgfId,omitempty" xml:"LgfId,omitempty"`
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

func (s DeleteBeebotIntentLgfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBeebotIntentLgfResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBeebotIntentLgfResponseBody) GetBeebotRequestId() *string {
	return s.BeebotRequestId
}

func (s *DeleteBeebotIntentLgfResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteBeebotIntentLgfResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteBeebotIntentLgfResponseBody) GetLgfId() *int64 {
	return s.LgfId
}

func (s *DeleteBeebotIntentLgfResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteBeebotIntentLgfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBeebotIntentLgfResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteBeebotIntentLgfResponseBody) SetBeebotRequestId(v string) *DeleteBeebotIntentLgfResponseBody {
	s.BeebotRequestId = &v
	return s
}

func (s *DeleteBeebotIntentLgfResponseBody) SetCode(v string) *DeleteBeebotIntentLgfResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBeebotIntentLgfResponseBody) SetHttpStatusCode(v int32) *DeleteBeebotIntentLgfResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteBeebotIntentLgfResponseBody) SetLgfId(v int64) *DeleteBeebotIntentLgfResponseBody {
	s.LgfId = &v
	return s
}

func (s *DeleteBeebotIntentLgfResponseBody) SetMessage(v string) *DeleteBeebotIntentLgfResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBeebotIntentLgfResponseBody) SetRequestId(v string) *DeleteBeebotIntentLgfResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBeebotIntentLgfResponseBody) SetSuccess(v bool) *DeleteBeebotIntentLgfResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBeebotIntentLgfResponseBody) Validate() error {
	return dara.Validate(s)
}
