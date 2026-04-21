// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateYikeProductionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTitle(v string) *CreateYikeProductionRequest
	GetTitle() *string
	SetWorkspaceId(v string) *CreateYikeProductionRequest
	GetWorkspaceId() *string
}

type CreateYikeProductionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 581236
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateYikeProductionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateYikeProductionRequest) GoString() string {
	return s.String()
}

func (s *CreateYikeProductionRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateYikeProductionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateYikeProductionRequest) SetTitle(v string) *CreateYikeProductionRequest {
	s.Title = &v
	return s
}

func (s *CreateYikeProductionRequest) SetWorkspaceId(v string) *CreateYikeProductionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateYikeProductionRequest) Validate() error {
	return dara.Validate(s)
}
