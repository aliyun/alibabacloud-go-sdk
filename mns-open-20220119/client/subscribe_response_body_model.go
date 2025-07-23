// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *SubscribeResponseBody
	GetCode() *int64
	SetData(v string) *SubscribeResponseBody
	GetData() *string
	SetMessage(v string) *SubscribeResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubscribeResponseBody
	GetRequestId() *string
	SetStatus(v string) *SubscribeResponseBody
	GetStatus() *string
	SetSuccess(v bool) *SubscribeResponseBody
	GetSuccess() *bool
}

type SubscribeResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// {\\"Code\\": 200, \\"Success\\": True}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubscribeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubscribeResponseBody) GoString() string {
	return s.String()
}

func (s *SubscribeResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *SubscribeResponseBody) GetData() *string {
	return s.Data
}

func (s *SubscribeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubscribeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubscribeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SubscribeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubscribeResponseBody) SetCode(v int64) *SubscribeResponseBody {
	s.Code = &v
	return s
}

func (s *SubscribeResponseBody) SetData(v string) *SubscribeResponseBody {
	s.Data = &v
	return s
}

func (s *SubscribeResponseBody) SetMessage(v string) *SubscribeResponseBody {
	s.Message = &v
	return s
}

func (s *SubscribeResponseBody) SetRequestId(v string) *SubscribeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubscribeResponseBody) SetStatus(v string) *SubscribeResponseBody {
	s.Status = &v
	return s
}

func (s *SubscribeResponseBody) SetSuccess(v bool) *SubscribeResponseBody {
	s.Success = &v
	return s
}

func (s *SubscribeResponseBody) Validate() error {
	return dara.Validate(s)
}
