// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitRectifyFlowForReviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *SubmitRectifyFlowForReviewRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *SubmitRectifyFlowForReviewRequest
	GetOpUserId() *string
	SetSubmitCommand(v *SubmitRectifyFlowForReviewRequestSubmitCommand) *SubmitRectifyFlowForReviewRequest
	GetSubmitCommand() *SubmitRectifyFlowForReviewRequestSubmitCommand
}

type SubmitRectifyFlowForReviewRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The operator user ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The business request parameters.
	//
	// This parameter is required.
	SubmitCommand *SubmitRectifyFlowForReviewRequestSubmitCommand `json:"SubmitCommand,omitempty" xml:"SubmitCommand,omitempty" type:"Struct"`
}

func (s SubmitRectifyFlowForReviewRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitRectifyFlowForReviewRequest) GoString() string {
	return s.String()
}

func (s *SubmitRectifyFlowForReviewRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SubmitRectifyFlowForReviewRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *SubmitRectifyFlowForReviewRequest) GetSubmitCommand() *SubmitRectifyFlowForReviewRequestSubmitCommand {
	return s.SubmitCommand
}

func (s *SubmitRectifyFlowForReviewRequest) SetOpTenantId(v int64) *SubmitRectifyFlowForReviewRequest {
	s.OpTenantId = &v
	return s
}

func (s *SubmitRectifyFlowForReviewRequest) SetOpUserId(v string) *SubmitRectifyFlowForReviewRequest {
	s.OpUserId = &v
	return s
}

func (s *SubmitRectifyFlowForReviewRequest) SetSubmitCommand(v *SubmitRectifyFlowForReviewRequestSubmitCommand) *SubmitRectifyFlowForReviewRequest {
	s.SubmitCommand = v
	return s
}

func (s *SubmitRectifyFlowForReviewRequest) Validate() error {
	if s.SubmitCommand != nil {
		if err := s.SubmitCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitRectifyFlowForReviewRequestSubmitCommand struct {
	// The description of the rectification result. This parameter is required and can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the rectification process. Submission is allowed only when the process is in the rectifying or re-rectifying state.
	//
	// This parameter is required.
	//
	// example:
	//
	// 358710058950
	RectifyId *int64 `json:"RectifyId,omitempty" xml:"RectifyId,omitempty"`
}

func (s SubmitRectifyFlowForReviewRequestSubmitCommand) String() string {
	return dara.Prettify(s)
}

func (s SubmitRectifyFlowForReviewRequestSubmitCommand) GoString() string {
	return s.String()
}

func (s *SubmitRectifyFlowForReviewRequestSubmitCommand) GetComment() *string {
	return s.Comment
}

func (s *SubmitRectifyFlowForReviewRequestSubmitCommand) GetRectifyId() *int64 {
	return s.RectifyId
}

func (s *SubmitRectifyFlowForReviewRequestSubmitCommand) SetComment(v string) *SubmitRectifyFlowForReviewRequestSubmitCommand {
	s.Comment = &v
	return s
}

func (s *SubmitRectifyFlowForReviewRequestSubmitCommand) SetRectifyId(v int64) *SubmitRectifyFlowForReviewRequestSubmitCommand {
	s.RectifyId = &v
	return s
}

func (s *SubmitRectifyFlowForReviewRequestSubmitCommand) Validate() error {
	return dara.Validate(s)
}
