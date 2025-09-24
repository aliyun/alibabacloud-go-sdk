// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetJobConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *DeleteDatasetJobConfigRequest
	GetWorkspaceId() *string
}

type DeleteDatasetJobConfigRequest struct {
	// The workspace ID.
	//
	// example:
	//
	// 513663
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteDatasetJobConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetJobConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetJobConfigRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteDatasetJobConfigRequest) SetWorkspaceId(v string) *DeleteDatasetJobConfigRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteDatasetJobConfigRequest) Validate() error {
	return dara.Validate(s)
}
