// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSampleApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateSampleApiResponseBody
	GetCode() *int32
	SetMessage(v string) *CreateSampleApiResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSampleApiResponseBody
	GetRequestId() *string
}

type CreateSampleApiResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Error message.
	//
	// example:
	//
	// ConsolePocQueryServiceImpl.queryServiceCodeName.arg0.tab tab Type error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSampleApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSampleApiResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSampleApiResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateSampleApiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSampleApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSampleApiResponseBody) SetCode(v int32) *CreateSampleApiResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSampleApiResponseBody) SetMessage(v string) *CreateSampleApiResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSampleApiResponseBody) SetRequestId(v string) *CreateSampleApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSampleApiResponseBody) Validate() error {
	return dara.Validate(s)
}
