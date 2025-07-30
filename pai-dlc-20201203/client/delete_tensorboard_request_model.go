// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTensorboardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *DeleteTensorboardRequest
	GetWorkspaceId() *string
}

type DeleteTensorboardRequest struct {
	// The workspace ID.
	//
	// example:
	//
	// 46099
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteTensorboardRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTensorboardRequest) GoString() string {
	return s.String()
}

func (s *DeleteTensorboardRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteTensorboardRequest) SetWorkspaceId(v string) *DeleteTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteTensorboardRequest) Validate() error {
	return dara.Validate(s)
}
