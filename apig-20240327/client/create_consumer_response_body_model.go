// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateConsumerResponseBody
	GetCode() *string
	SetData(v *CreateConsumerResponseBodyData) *CreateConsumerResponseBody
	GetData() *CreateConsumerResponseBodyData
	SetMessage(v string) *CreateConsumerResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateConsumerResponseBody
	GetRequestId() *string
}

type CreateConsumerResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response payload.
	Data *CreateConsumerResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The status message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 479AE38F-A574-52F7-87EA-E91199999F9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateConsumerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConsumerResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateConsumerResponseBody) GetData() *CreateConsumerResponseBodyData {
	return s.Data
}

func (s *CreateConsumerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateConsumerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConsumerResponseBody) SetCode(v string) *CreateConsumerResponseBody {
	s.Code = &v
	return s
}

func (s *CreateConsumerResponseBody) SetData(v *CreateConsumerResponseBodyData) *CreateConsumerResponseBody {
	s.Data = v
	return s
}

func (s *CreateConsumerResponseBody) SetMessage(v string) *CreateConsumerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateConsumerResponseBody) SetRequestId(v string) *CreateConsumerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConsumerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateConsumerResponseBodyData struct {
	// The consumer ID.
	//
	// example:
	//
	// cs-cvnjramm1hks1r3fmmug
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
}

func (s CreateConsumerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateConsumerResponseBodyData) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *CreateConsumerResponseBodyData) SetConsumerId(v string) *CreateConsumerResponseBodyData {
	s.ConsumerId = &v
	return s
}

func (s *CreateConsumerResponseBodyData) Validate() error {
	return dara.Validate(s)
}
