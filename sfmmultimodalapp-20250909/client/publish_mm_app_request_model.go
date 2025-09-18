// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishMmAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *PublishMmAppRequest
	GetAppId() *string
	SetDescription(v string) *PublishMmAppRequest
	GetDescription() *string
	SetWorkspaceId(v string) *PublishMmAppRequest
	GetWorkspaceId() *string
}

type PublishMmAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_xxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// xxxxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PublishMmAppRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishMmAppRequest) GoString() string {
	return s.String()
}

func (s *PublishMmAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *PublishMmAppRequest) GetDescription() *string {
	return s.Description
}

func (s *PublishMmAppRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PublishMmAppRequest) SetAppId(v string) *PublishMmAppRequest {
	s.AppId = &v
	return s
}

func (s *PublishMmAppRequest) SetDescription(v string) *PublishMmAppRequest {
	s.Description = &v
	return s
}

func (s *PublishMmAppRequest) SetWorkspaceId(v string) *PublishMmAppRequest {
	s.WorkspaceId = &v
	return s
}

func (s *PublishMmAppRequest) Validate() error {
	return dara.Validate(s)
}
