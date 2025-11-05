// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscriptionBillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubscriptionBillResponseBody
	GetCode() *string
	SetData(v bool) *SubscriptionBillResponseBody
	GetData() *bool
	SetMessage(v string) *SubscriptionBillResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubscriptionBillResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubscriptionBillResponseBody
	GetSuccess() *bool
}

type SubscriptionBillResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubscriptionBillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionBillResponseBody) GoString() string {
	return s.String()
}

func (s *SubscriptionBillResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubscriptionBillResponseBody) GetData() *bool {
	return s.Data
}

func (s *SubscriptionBillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubscriptionBillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubscriptionBillResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubscriptionBillResponseBody) SetCode(v string) *SubscriptionBillResponseBody {
	s.Code = &v
	return s
}

func (s *SubscriptionBillResponseBody) SetData(v bool) *SubscriptionBillResponseBody {
	s.Data = &v
	return s
}

func (s *SubscriptionBillResponseBody) SetMessage(v string) *SubscriptionBillResponseBody {
	s.Message = &v
	return s
}

func (s *SubscriptionBillResponseBody) SetRequestId(v string) *SubscriptionBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubscriptionBillResponseBody) SetSuccess(v bool) *SubscriptionBillResponseBody {
	s.Success = &v
	return s
}

func (s *SubscriptionBillResponseBody) Validate() error {
	return dara.Validate(s)
}
