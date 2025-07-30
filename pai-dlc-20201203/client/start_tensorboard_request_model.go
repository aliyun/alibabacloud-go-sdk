// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTensorboardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *StartTensorboardRequest
	GetWorkspaceId() *string
}

type StartTensorboardRequest struct {
	// The workspace ID.
	//
	// example:
	//
	// 380
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s StartTensorboardRequest) String() string {
	return dara.Prettify(s)
}

func (s StartTensorboardRequest) GoString() string {
	return s.String()
}

func (s *StartTensorboardRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *StartTensorboardRequest) SetWorkspaceId(v string) *StartTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

func (s *StartTensorboardRequest) Validate() error {
	return dara.Validate(s)
}
