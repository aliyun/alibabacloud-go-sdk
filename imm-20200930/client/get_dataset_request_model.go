// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *GetDatasetRequest
	GetDatasetName() *string
	SetProjectName(v string) *GetDatasetRequest
	GetProjectName() *string
	SetWithStatistics(v bool) *GetDatasetRequest
	GetWithStatistics() *bool
}

type GetDatasetRequest struct {
	// The name of the dataset. You can obtain the name of the dataset from the response of the [CreateDataset](https://help.aliyun.com/document_detail/478160.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// dataset001
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Specifies whether to enable real-time retrieval of file statistics. Default value: false.
	//
	// 	- If you set the value to true, FileCount and TotalFileSize in the response return true and valid values.
	//
	// 	- If you set the value to false, FileCount and TotalFileSize in the response return invalid values or 0.
	//
	// example:
	//
	// true
	WithStatistics *bool `json:"WithStatistics,omitempty" xml:"WithStatistics,omitempty"`
}

func (s GetDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetRequest) GoString() string {
	return s.String()
}

func (s *GetDatasetRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetDatasetRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetDatasetRequest) GetWithStatistics() *bool {
	return s.WithStatistics
}

func (s *GetDatasetRequest) SetDatasetName(v string) *GetDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *GetDatasetRequest) SetProjectName(v string) *GetDatasetRequest {
	s.ProjectName = &v
	return s
}

func (s *GetDatasetRequest) SetWithStatistics(v bool) *GetDatasetRequest {
	s.WithStatistics = &v
	return s
}

func (s *GetDatasetRequest) Validate() error {
	return dara.Validate(s)
}
