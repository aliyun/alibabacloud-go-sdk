// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTensorboardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *StopTensorboardRequest
	GetWorkspaceId() *string
}

type StopTensorboardRequest struct {
	// The workspace ID.
	//
	// example:
	//
	// 380
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s StopTensorboardRequest) String() string {
	return dara.Prettify(s)
}

func (s StopTensorboardRequest) GoString() string {
	return s.String()
}

func (s *StopTensorboardRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *StopTensorboardRequest) SetWorkspaceId(v string) *StopTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

func (s *StopTensorboardRequest) Validate() error {
	return dara.Validate(s)
}
