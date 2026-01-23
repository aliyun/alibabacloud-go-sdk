// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishStandardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *PublishStandardRequest
	GetOpTenantId() *int64
	SetPublishCommand(v *PublishStandardRequestPublishCommand) *PublishStandardRequest
	GetPublishCommand() *PublishStandardRequestPublishCommand
}

type PublishStandardRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	PublishCommand *PublishStandardRequestPublishCommand `json:"PublishCommand,omitempty" xml:"PublishCommand,omitempty" type:"Struct"`
}

func (s PublishStandardRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishStandardRequest) GoString() string {
	return s.String()
}

func (s *PublishStandardRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *PublishStandardRequest) GetPublishCommand() *PublishStandardRequestPublishCommand {
	return s.PublishCommand
}

func (s *PublishStandardRequest) SetOpTenantId(v int64) *PublishStandardRequest {
	s.OpTenantId = &v
	return s
}

func (s *PublishStandardRequest) SetPublishCommand(v *PublishStandardRequestPublishCommand) *PublishStandardRequest {
	s.PublishCommand = v
	return s
}

func (s *PublishStandardRequest) Validate() error {
	if s.PublishCommand != nil {
		if err := s.PublishCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PublishStandardRequestPublishCommand struct {
	AutoPublishAfterApproval *bool `json:"AutoPublishAfterApproval,omitempty" xml:"AutoPublishAfterApproval,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id             *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	ReviewerIdList []*string `json:"ReviewerIdList,omitempty" xml:"ReviewerIdList,omitempty" type:"Repeated"`
	// example:
	//
	// DEV
	StandardStage *string `json:"StandardStage,omitempty" xml:"StandardStage,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s PublishStandardRequestPublishCommand) String() string {
	return dara.Prettify(s)
}

func (s PublishStandardRequestPublishCommand) GoString() string {
	return s.String()
}

func (s *PublishStandardRequestPublishCommand) GetAutoPublishAfterApproval() *bool {
	return s.AutoPublishAfterApproval
}

func (s *PublishStandardRequestPublishCommand) GetComment() *string {
	return s.Comment
}

func (s *PublishStandardRequestPublishCommand) GetId() *int64 {
	return s.Id
}

func (s *PublishStandardRequestPublishCommand) GetReviewerIdList() []*string {
	return s.ReviewerIdList
}

func (s *PublishStandardRequestPublishCommand) GetStandardStage() *string {
	return s.StandardStage
}

func (s *PublishStandardRequestPublishCommand) GetVersion() *int32 {
	return s.Version
}

func (s *PublishStandardRequestPublishCommand) SetAutoPublishAfterApproval(v bool) *PublishStandardRequestPublishCommand {
	s.AutoPublishAfterApproval = &v
	return s
}

func (s *PublishStandardRequestPublishCommand) SetComment(v string) *PublishStandardRequestPublishCommand {
	s.Comment = &v
	return s
}

func (s *PublishStandardRequestPublishCommand) SetId(v int64) *PublishStandardRequestPublishCommand {
	s.Id = &v
	return s
}

func (s *PublishStandardRequestPublishCommand) SetReviewerIdList(v []*string) *PublishStandardRequestPublishCommand {
	s.ReviewerIdList = v
	return s
}

func (s *PublishStandardRequestPublishCommand) SetStandardStage(v string) *PublishStandardRequestPublishCommand {
	s.StandardStage = &v
	return s
}

func (s *PublishStandardRequestPublishCommand) SetVersion(v int32) *PublishStandardRequestPublishCommand {
	s.Version = &v
	return s
}

func (s *PublishStandardRequestPublishCommand) Validate() error {
	return dara.Validate(s)
}
