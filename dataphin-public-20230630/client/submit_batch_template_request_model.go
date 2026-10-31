// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBatchTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *SubmitBatchTemplateRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *SubmitBatchTemplateRequest
	GetOpUserId() *string
	SetProjectId(v int64) *SubmitBatchTemplateRequest
	GetProjectId() *int64
	SetSubmitCommand(v *SubmitBatchTemplateRequestSubmitCommand) *SubmitBatchTemplateRequest
	GetSubmitCommand() *SubmitBatchTemplateRequestSubmitCommand
}

type SubmitBatchTemplateRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The template submission instruction.
	//
	// This parameter is required.
	SubmitCommand *SubmitBatchTemplateRequestSubmitCommand `json:"SubmitCommand,omitempty" xml:"SubmitCommand,omitempty" type:"Struct"`
}

func (s SubmitBatchTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTemplateRequest) GoString() string {
	return s.String()
}

func (s *SubmitBatchTemplateRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SubmitBatchTemplateRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *SubmitBatchTemplateRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *SubmitBatchTemplateRequest) GetSubmitCommand() *SubmitBatchTemplateRequestSubmitCommand {
	return s.SubmitCommand
}

func (s *SubmitBatchTemplateRequest) SetOpTenantId(v int64) *SubmitBatchTemplateRequest {
	s.OpTenantId = &v
	return s
}

func (s *SubmitBatchTemplateRequest) SetOpUserId(v string) *SubmitBatchTemplateRequest {
	s.OpUserId = &v
	return s
}

func (s *SubmitBatchTemplateRequest) SetProjectId(v int64) *SubmitBatchTemplateRequest {
	s.ProjectId = &v
	return s
}

func (s *SubmitBatchTemplateRequest) SetSubmitCommand(v *SubmitBatchTemplateRequestSubmitCommand) *SubmitBatchTemplateRequest {
	s.SubmitCommand = v
	return s
}

func (s *SubmitBatchTemplateRequest) Validate() error {
	if s.SubmitCommand != nil {
		if err := s.SubmitCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitBatchTemplateRequestSubmitCommand struct {
	// The submission comment.
	//
	// example:
	//
	// Initial submission
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1001
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitBatchTemplateRequestSubmitCommand) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTemplateRequestSubmitCommand) GoString() string {
	return s.String()
}

func (s *SubmitBatchTemplateRequestSubmitCommand) GetComment() *string {
	return s.Comment
}

func (s *SubmitBatchTemplateRequestSubmitCommand) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *SubmitBatchTemplateRequestSubmitCommand) SetComment(v string) *SubmitBatchTemplateRequestSubmitCommand {
	s.Comment = &v
	return s
}

func (s *SubmitBatchTemplateRequestSubmitCommand) SetTemplateId(v int64) *SubmitBatchTemplateRequestSubmitCommand {
	s.TemplateId = &v
	return s
}

func (s *SubmitBatchTemplateRequestSubmitCommand) Validate() error {
	return dara.Validate(s)
}
