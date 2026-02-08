// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawScriptReviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *WithdrawScriptReviewRequest
	GetInstanceId() *string
	SetScriptId(v string) *WithdrawScriptReviewRequest
	GetScriptId() *string
}

type WithdrawScriptReviewRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 38d2e8ed-04e9-4dac-83b5-a8e57642ef13
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Script ID
	//
	// This parameter is required.
	//
	// example:
	//
	// e4e2a770-b97b-465a-80d8-06dca008c503
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s WithdrawScriptReviewRequest) String() string {
	return dara.Prettify(s)
}

func (s WithdrawScriptReviewRequest) GoString() string {
	return s.String()
}

func (s *WithdrawScriptReviewRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *WithdrawScriptReviewRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *WithdrawScriptReviewRequest) SetInstanceId(v string) *WithdrawScriptReviewRequest {
	s.InstanceId = &v
	return s
}

func (s *WithdrawScriptReviewRequest) SetScriptId(v string) *WithdrawScriptReviewRequest {
	s.ScriptId = &v
	return s
}

func (s *WithdrawScriptReviewRequest) Validate() error {
	return dara.Validate(s)
}
