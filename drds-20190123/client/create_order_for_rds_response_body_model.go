// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderForRdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateOrderForRdsResponseBody
	GetData() *string
	SetRequestId(v string) *CreateOrderForRdsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateOrderForRdsResponseBody
	GetSuccess() *bool
}

type CreateOrderForRdsResponseBody struct {
	// The ID of the purchased RDS instance.
	//
	// example:
	//
	// [rm-***********]
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9819BC51-D33D-4EB1-B80F-A89A20******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateOrderForRdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderForRdsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderForRdsResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateOrderForRdsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrderForRdsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateOrderForRdsResponseBody) SetData(v string) *CreateOrderForRdsResponseBody {
	s.Data = &v
	return s
}

func (s *CreateOrderForRdsResponseBody) SetRequestId(v string) *CreateOrderForRdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrderForRdsResponseBody) SetSuccess(v bool) *CreateOrderForRdsResponseBody {
	s.Success = &v
	return s
}

func (s *CreateOrderForRdsResponseBody) Validate() error {
	return dara.Validate(s)
}
