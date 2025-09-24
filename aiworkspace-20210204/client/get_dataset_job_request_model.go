// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetVersion(v string) *GetDatasetJobRequest
	GetDatasetVersion() *string
	SetWorkspaceId(v string) *GetDatasetJobRequest
	GetWorkspaceId() *string
}

type GetDatasetJobRequest struct {
	// The dataset version name.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The workspace ID. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
	//
	// example:
	//
	// 478**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDatasetJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetJobRequest) GoString() string {
	return s.String()
}

func (s *GetDatasetJobRequest) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *GetDatasetJobRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDatasetJobRequest) SetDatasetVersion(v string) *GetDatasetJobRequest {
	s.DatasetVersion = &v
	return s
}

func (s *GetDatasetJobRequest) SetWorkspaceId(v string) *GetDatasetJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetDatasetJobRequest) Validate() error {
	return dara.Validate(s)
}
