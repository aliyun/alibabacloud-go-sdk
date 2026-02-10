// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultimodalFineTuneDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateMultimodalFineTuneDatasetRequest
	GetDBClusterId() *string
	SetDatasetDescription(v string) *CreateMultimodalFineTuneDatasetRequest
	GetDatasetDescription() *string
	SetDatasetName(v string) *CreateMultimodalFineTuneDatasetRequest
	GetDatasetName() *string
}

type CreateMultimodalFineTuneDatasetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 用户输入数据集描述
	DatasetDescription *string `json:"DatasetDescription,omitempty" xml:"DatasetDescription,omitempty"`
	// example:
	//
	// 用户输入数据集名
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
}

func (s CreateMultimodalFineTuneDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMultimodalFineTuneDatasetRequest) GoString() string {
	return s.String()
}

func (s *CreateMultimodalFineTuneDatasetRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateMultimodalFineTuneDatasetRequest) GetDatasetDescription() *string {
	return s.DatasetDescription
}

func (s *CreateMultimodalFineTuneDatasetRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateMultimodalFineTuneDatasetRequest) SetDBClusterId(v string) *CreateMultimodalFineTuneDatasetRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateMultimodalFineTuneDatasetRequest) SetDatasetDescription(v string) *CreateMultimodalFineTuneDatasetRequest {
	s.DatasetDescription = &v
	return s
}

func (s *CreateMultimodalFineTuneDatasetRequest) SetDatasetName(v string) *CreateMultimodalFineTuneDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateMultimodalFineTuneDatasetRequest) Validate() error {
	return dara.Validate(s)
}
