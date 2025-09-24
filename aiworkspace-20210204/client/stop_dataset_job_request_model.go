// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDatasetJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetVersion(v string) *StopDatasetJobRequest
	GetDatasetVersion() *string
	SetWorkspaceId(v string) *StopDatasetJobRequest
	GetWorkspaceId() *string
}

type StopDatasetJobRequest struct {
	// The dataset version.
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

func (s StopDatasetJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDatasetJobRequest) GoString() string {
	return s.String()
}

func (s *StopDatasetJobRequest) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *StopDatasetJobRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *StopDatasetJobRequest) SetDatasetVersion(v string) *StopDatasetJobRequest {
	s.DatasetVersion = &v
	return s
}

func (s *StopDatasetJobRequest) SetWorkspaceId(v string) *StopDatasetJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *StopDatasetJobRequest) Validate() error {
	return dara.Validate(s)
}
