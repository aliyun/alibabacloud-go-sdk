// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSolutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *SubmitSolutionRequest
	GetBizType() *string
	SetIntentionBizId(v string) *SubmitSolutionRequest
	GetIntentionBizId() *string
	SetOperateType(v string) *SubmitSolutionRequest
	GetOperateType() *string
	SetSolution(v string) *SubmitSolutionRequest
	GetSolution() *string
	SetUserId(v string) *SubmitSolutionRequest
	GetUserId() *string
}

type SubmitSolutionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// esp.wangwen
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// I20211223101045000001
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	OperateType    *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// This parameter is required.
	Solution *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	// example:
	//
	// 1219541161213057
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SubmitSolutionRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSolutionRequest) GoString() string {
	return s.String()
}

func (s *SubmitSolutionRequest) GetBizType() *string {
	return s.BizType
}

func (s *SubmitSolutionRequest) GetIntentionBizId() *string {
	return s.IntentionBizId
}

func (s *SubmitSolutionRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *SubmitSolutionRequest) GetSolution() *string {
	return s.Solution
}

func (s *SubmitSolutionRequest) GetUserId() *string {
	return s.UserId
}

func (s *SubmitSolutionRequest) SetBizType(v string) *SubmitSolutionRequest {
	s.BizType = &v
	return s
}

func (s *SubmitSolutionRequest) SetIntentionBizId(v string) *SubmitSolutionRequest {
	s.IntentionBizId = &v
	return s
}

func (s *SubmitSolutionRequest) SetOperateType(v string) *SubmitSolutionRequest {
	s.OperateType = &v
	return s
}

func (s *SubmitSolutionRequest) SetSolution(v string) *SubmitSolutionRequest {
	s.Solution = &v
	return s
}

func (s *SubmitSolutionRequest) SetUserId(v string) *SubmitSolutionRequest {
	s.UserId = &v
	return s
}

func (s *SubmitSolutionRequest) Validate() error {
	return dara.Validate(s)
}
