// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSubscriptionBillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelSubscriptionBillResponseBody
	GetCode() *string
	SetData(v bool) *CancelSubscriptionBillResponseBody
	GetData() *bool
	SetMessage(v string) *CancelSubscriptionBillResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelSubscriptionBillResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelSubscriptionBillResponseBody
	GetSuccess() *bool
}

type CancelSubscriptionBillResponseBody struct {
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
	// 210e876f16704666020714468dab35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelSubscriptionBillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelSubscriptionBillResponseBody) GoString() string {
	return s.String()
}

func (s *CancelSubscriptionBillResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelSubscriptionBillResponseBody) GetData() *bool {
	return s.Data
}

func (s *CancelSubscriptionBillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelSubscriptionBillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelSubscriptionBillResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelSubscriptionBillResponseBody) SetCode(v string) *CancelSubscriptionBillResponseBody {
	s.Code = &v
	return s
}

func (s *CancelSubscriptionBillResponseBody) SetData(v bool) *CancelSubscriptionBillResponseBody {
	s.Data = &v
	return s
}

func (s *CancelSubscriptionBillResponseBody) SetMessage(v string) *CancelSubscriptionBillResponseBody {
	s.Message = &v
	return s
}

func (s *CancelSubscriptionBillResponseBody) SetRequestId(v string) *CancelSubscriptionBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelSubscriptionBillResponseBody) SetSuccess(v bool) *CancelSubscriptionBillResponseBody {
	s.Success = &v
	return s
}

func (s *CancelSubscriptionBillResponseBody) Validate() error {
	return dara.Validate(s)
}
