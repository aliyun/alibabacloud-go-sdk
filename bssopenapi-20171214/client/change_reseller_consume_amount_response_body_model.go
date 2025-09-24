// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResellerConsumeAmountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ChangeResellerConsumeAmountResponseBody
	GetCode() *string
	SetData(v string) *ChangeResellerConsumeAmountResponseBody
	GetData() *string
	SetMessage(v string) *ChangeResellerConsumeAmountResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChangeResellerConsumeAmountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChangeResellerConsumeAmountResponseBody
	GetSuccess() *bool
}

type ChangeResellerConsumeAmountResponseBody struct {
	// The error code returned if the call failed. For more information, see the "Error codes" section of the topic.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The consumption quota for the quota ledger after adjustment.
	//
	// example:
	//
	// 300.00
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D6E068C3-25BC-455A-85FE-45F0B22ECB1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates that the call is successful. A value of false indicates that the call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeResellerConsumeAmountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeResellerConsumeAmountResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResellerConsumeAmountResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChangeResellerConsumeAmountResponseBody) GetData() *string {
	return s.Data
}

func (s *ChangeResellerConsumeAmountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeResellerConsumeAmountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeResellerConsumeAmountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeResellerConsumeAmountResponseBody) SetCode(v string) *ChangeResellerConsumeAmountResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeResellerConsumeAmountResponseBody) SetData(v string) *ChangeResellerConsumeAmountResponseBody {
	s.Data = &v
	return s
}

func (s *ChangeResellerConsumeAmountResponseBody) SetMessage(v string) *ChangeResellerConsumeAmountResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeResellerConsumeAmountResponseBody) SetRequestId(v string) *ChangeResellerConsumeAmountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeResellerConsumeAmountResponseBody) SetSuccess(v bool) *ChangeResellerConsumeAmountResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeResellerConsumeAmountResponseBody) Validate() error {
	return dara.Validate(s)
}
