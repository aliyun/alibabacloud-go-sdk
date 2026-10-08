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
	// The name of the dataset. For more information about how to obtain the dataset name, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// dataset001
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The name of the project. For more information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Specifies whether to collect file statistics. Valid values:
	//
	// - true: File statistics are collected. The FileCount and TotalFileSize fields in the Dataset struct are valid.
	//
	// - false: File statistics are not collected. The FileCount and TotalFileSize fields in the Dataset struct may be incorrect or both 0.
	//
	// Default value: false.
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
