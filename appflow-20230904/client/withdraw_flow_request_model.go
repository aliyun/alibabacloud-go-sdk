// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v string) *WithdrawFlowRequest
	GetFlowId() *string
	SetFlowVersion(v int32) *WithdrawFlowRequest
	GetFlowVersion() *int32
}

type WithdrawFlowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// flow-6b03788c25244c93b254
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// example:
	//
	// 1
	FlowVersion *int32 `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
}

func (s WithdrawFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s WithdrawFlowRequest) GoString() string {
	return s.String()
}

func (s *WithdrawFlowRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *WithdrawFlowRequest) GetFlowVersion() *int32 {
	return s.FlowVersion
}

func (s *WithdrawFlowRequest) SetFlowId(v string) *WithdrawFlowRequest {
	s.FlowId = &v
	return s
}

func (s *WithdrawFlowRequest) SetFlowVersion(v int32) *WithdrawFlowRequest {
	s.FlowVersion = &v
	return s
}

func (s *WithdrawFlowRequest) Validate() error {
	return dara.Validate(s)
}
