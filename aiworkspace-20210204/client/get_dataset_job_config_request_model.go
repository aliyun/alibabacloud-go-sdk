// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetJobConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *GetDatasetJobConfigRequest
	GetWorkspaceId() *string
}

type GetDatasetJobConfigRequest struct {
	// The workspace ID.
	//
	// example:
	//
	// 114243
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDatasetJobConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetJobConfigRequest) GoString() string {
	return s.String()
}

func (s *GetDatasetJobConfigRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDatasetJobConfigRequest) SetWorkspaceId(v string) *GetDatasetJobConfigRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetDatasetJobConfigRequest) Validate() error {
	return dara.Validate(s)
}
