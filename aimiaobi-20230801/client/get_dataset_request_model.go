// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v int64) *GetDatasetRequest
	GetDatasetId() *int64
	SetDatasetName(v string) *GetDatasetRequest
	GetDatasetName() *string
	SetWorkspaceId(v string) *GetDatasetRequest
	GetWorkspaceId() *string
}

type GetDatasetRequest struct {
	// example:
	//
	// 1
	DatasetId *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// businessDataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetRequest) GoString() string {
	return s.String()
}

func (s *GetDatasetRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *GetDatasetRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetDatasetRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDatasetRequest) SetDatasetId(v int64) *GetDatasetRequest {
	s.DatasetId = &v
	return s
}

func (s *GetDatasetRequest) SetDatasetName(v string) *GetDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *GetDatasetRequest) SetWorkspaceId(v string) *GetDatasetRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetDatasetRequest) Validate() error {
	return dara.Validate(s)
}
