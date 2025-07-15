// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v int64) *DeleteDatasetRequest
	GetDatasetId() *int64
	SetWorkspaceId(v string) *DeleteDatasetRequest
	GetWorkspaceId() *string
}

type DeleteDatasetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DatasetId *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *DeleteDatasetRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteDatasetRequest) SetDatasetId(v int64) *DeleteDatasetRequest {
	s.DatasetId = &v
	return s
}

func (s *DeleteDatasetRequest) SetWorkspaceId(v string) *DeleteDatasetRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteDatasetRequest) Validate() error {
	return dara.Validate(s)
}
