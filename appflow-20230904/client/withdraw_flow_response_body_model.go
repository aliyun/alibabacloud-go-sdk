// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *WithdrawFlowResponseBody
	GetData() *string
	SetRequestId(v string) *WithdrawFlowResponseBody
	GetRequestId() *string
}

type WithdrawFlowResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// None
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39B608FB-906F-51CF-AD82-7EFE46C0D56A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s WithdrawFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WithdrawFlowResponseBody) GoString() string {
	return s.String()
}

func (s *WithdrawFlowResponseBody) GetData() *string {
	return s.Data
}

func (s *WithdrawFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WithdrawFlowResponseBody) SetData(v string) *WithdrawFlowResponseBody {
	s.Data = &v
	return s
}

func (s *WithdrawFlowResponseBody) SetRequestId(v string) *WithdrawFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *WithdrawFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
