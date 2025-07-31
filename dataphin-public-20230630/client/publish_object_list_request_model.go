// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishObjectListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *PublishObjectListRequest
	GetOpTenantId() *int64
	SetPublishCommand(v *PublishObjectListRequestPublishCommand) *PublishObjectListRequest
	GetPublishCommand() *PublishObjectListRequestPublishCommand
}

type PublishObjectListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	PublishCommand *PublishObjectListRequestPublishCommand `json:"PublishCommand,omitempty" xml:"PublishCommand,omitempty" type:"Struct"`
}

func (s PublishObjectListRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishObjectListRequest) GoString() string {
	return s.String()
}

func (s *PublishObjectListRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *PublishObjectListRequest) GetPublishCommand() *PublishObjectListRequestPublishCommand {
	return s.PublishCommand
}

func (s *PublishObjectListRequest) SetOpTenantId(v int64) *PublishObjectListRequest {
	s.OpTenantId = &v
	return s
}

func (s *PublishObjectListRequest) SetPublishCommand(v *PublishObjectListRequestPublishCommand) *PublishObjectListRequest {
	s.PublishCommand = v
	return s
}

func (s *PublishObjectListRequest) Validate() error {
	return dara.Validate(s)
}

type PublishObjectListRequestPublishCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 发布
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	SubmitIdList []*int64 `json:"SubmitIdList,omitempty" xml:"SubmitIdList,omitempty" type:"Repeated"`
}

func (s PublishObjectListRequestPublishCommand) String() string {
	return dara.Prettify(s)
}

func (s PublishObjectListRequestPublishCommand) GoString() string {
	return s.String()
}

func (s *PublishObjectListRequestPublishCommand) GetComment() *string {
	return s.Comment
}

func (s *PublishObjectListRequestPublishCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *PublishObjectListRequestPublishCommand) GetSubmitIdList() []*int64 {
	return s.SubmitIdList
}

func (s *PublishObjectListRequestPublishCommand) SetComment(v string) *PublishObjectListRequestPublishCommand {
	s.Comment = &v
	return s
}

func (s *PublishObjectListRequestPublishCommand) SetProjectId(v int64) *PublishObjectListRequestPublishCommand {
	s.ProjectId = &v
	return s
}

func (s *PublishObjectListRequestPublishCommand) SetSubmitIdList(v []*int64) *PublishObjectListRequestPublishCommand {
	s.SubmitIdList = v
	return s
}

func (s *PublishObjectListRequestPublishCommand) Validate() error {
	return dara.Validate(s)
}
