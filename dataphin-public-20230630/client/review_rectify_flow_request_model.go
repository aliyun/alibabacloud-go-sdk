// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReviewRectifyFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *ReviewRectifyFlowRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ReviewRectifyFlowRequest
	GetOpUserId() *string
	SetReviewCommand(v *ReviewRectifyFlowRequestReviewCommand) *ReviewRectifyFlowRequest
	GetReviewCommand() *ReviewRectifyFlowRequestReviewCommand
}

type ReviewRectifyFlowRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The business request parameters.
	//
	// This parameter is required.
	ReviewCommand *ReviewRectifyFlowRequestReviewCommand `json:"ReviewCommand,omitempty" xml:"ReviewCommand,omitempty" type:"Struct"`
}

func (s ReviewRectifyFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s ReviewRectifyFlowRequest) GoString() string {
	return s.String()
}

func (s *ReviewRectifyFlowRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ReviewRectifyFlowRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ReviewRectifyFlowRequest) GetReviewCommand() *ReviewRectifyFlowRequestReviewCommand {
	return s.ReviewCommand
}

func (s *ReviewRectifyFlowRequest) SetOpTenantId(v int64) *ReviewRectifyFlowRequest {
	s.OpTenantId = &v
	return s
}

func (s *ReviewRectifyFlowRequest) SetOpUserId(v string) *ReviewRectifyFlowRequest {
	s.OpUserId = &v
	return s
}

func (s *ReviewRectifyFlowRequest) SetReviewCommand(v *ReviewRectifyFlowRequestReviewCommand) *ReviewRectifyFlowRequest {
	s.ReviewCommand = v
	return s
}

func (s *ReviewRectifyFlowRequest) Validate() error {
	if s.ReviewCommand != nil {
		if err := s.ReviewCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReviewRectifyFlowRequestReviewCommand struct {
	// The acceptance description. This parameter is optional and can contain up to 128 characters.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the rectification process. This operation is allowed only when the process is pending acceptance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 616543434760
	RectifyId *int64 `json:"RectifyId,omitempty" xml:"RectifyId,omitempty"`
	// The acceptance result. Valid values: PASSED and REJECTED.
	//
	// This parameter is required.
	//
	// example:
	//
	// PASSED
	ReviewResult *string `json:"ReviewResult,omitempty" xml:"ReviewResult,omitempty"`
}

func (s ReviewRectifyFlowRequestReviewCommand) String() string {
	return dara.Prettify(s)
}

func (s ReviewRectifyFlowRequestReviewCommand) GoString() string {
	return s.String()
}

func (s *ReviewRectifyFlowRequestReviewCommand) GetComment() *string {
	return s.Comment
}

func (s *ReviewRectifyFlowRequestReviewCommand) GetRectifyId() *int64 {
	return s.RectifyId
}

func (s *ReviewRectifyFlowRequestReviewCommand) GetReviewResult() *string {
	return s.ReviewResult
}

func (s *ReviewRectifyFlowRequestReviewCommand) SetComment(v string) *ReviewRectifyFlowRequestReviewCommand {
	s.Comment = &v
	return s
}

func (s *ReviewRectifyFlowRequestReviewCommand) SetRectifyId(v int64) *ReviewRectifyFlowRequestReviewCommand {
	s.RectifyId = &v
	return s
}

func (s *ReviewRectifyFlowRequestReviewCommand) SetReviewResult(v string) *ReviewRectifyFlowRequestReviewCommand {
	s.ReviewResult = &v
	return s
}

func (s *ReviewRectifyFlowRequestReviewCommand) Validate() error {
	return dara.Validate(s)
}
