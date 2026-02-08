// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBeebotIntentLgfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeebotRequestId(v string) *CreateBeebotIntentLgfResponseBody
	GetBeebotRequestId() *string
	SetCode(v string) *CreateBeebotIntentLgfResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateBeebotIntentLgfResponseBody
	GetHttpStatusCode() *int32
	SetLgfId(v int64) *CreateBeebotIntentLgfResponseBody
	GetLgfId() *int64
	SetMessage(v string) *CreateBeebotIntentLgfResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateBeebotIntentLgfResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBeebotIntentLgfResponseBody
	GetSuccess() *bool
}

type CreateBeebotIntentLgfResponseBody struct {
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
	// HTTP error code
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
	// API prompt message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the operation succeeded
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBeebotIntentLgfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBeebotIntentLgfResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBeebotIntentLgfResponseBody) GetBeebotRequestId() *string {
	return s.BeebotRequestId
}

func (s *CreateBeebotIntentLgfResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateBeebotIntentLgfResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateBeebotIntentLgfResponseBody) GetLgfId() *int64 {
	return s.LgfId
}

func (s *CreateBeebotIntentLgfResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateBeebotIntentLgfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBeebotIntentLgfResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBeebotIntentLgfResponseBody) SetBeebotRequestId(v string) *CreateBeebotIntentLgfResponseBody {
	s.BeebotRequestId = &v
	return s
}

func (s *CreateBeebotIntentLgfResponseBody) SetCode(v string) *CreateBeebotIntentLgfResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBeebotIntentLgfResponseBody) SetHttpStatusCode(v int32) *CreateBeebotIntentLgfResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateBeebotIntentLgfResponseBody) SetLgfId(v int64) *CreateBeebotIntentLgfResponseBody {
	s.LgfId = &v
	return s
}

func (s *CreateBeebotIntentLgfResponseBody) SetMessage(v string) *CreateBeebotIntentLgfResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBeebotIntentLgfResponseBody) SetRequestId(v string) *CreateBeebotIntentLgfResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBeebotIntentLgfResponseBody) SetSuccess(v bool) *CreateBeebotIntentLgfResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBeebotIntentLgfResponseBody) Validate() error {
	return dara.Validate(s)
}
