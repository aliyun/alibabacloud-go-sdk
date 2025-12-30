// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultimodalDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateMultimodalDatasetRequest
	GetDBClusterId() *string
	SetDatasetDescription(v string) *CreateMultimodalDatasetRequest
	GetDatasetDescription() *string
	SetDatasetName(v string) *CreateMultimodalDatasetRequest
	GetDatasetName() *string
}

type CreateMultimodalDatasetRequest struct {
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

func (s CreateMultimodalDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMultimodalDatasetRequest) GoString() string {
	return s.String()
}

func (s *CreateMultimodalDatasetRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateMultimodalDatasetRequest) GetDatasetDescription() *string {
	return s.DatasetDescription
}

func (s *CreateMultimodalDatasetRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateMultimodalDatasetRequest) SetDBClusterId(v string) *CreateMultimodalDatasetRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateMultimodalDatasetRequest) SetDatasetDescription(v string) *CreateMultimodalDatasetRequest {
	s.DatasetDescription = &v
	return s
}

func (s *CreateMultimodalDatasetRequest) SetDatasetName(v string) *CreateMultimodalDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateMultimodalDatasetRequest) Validate() error {
	return dara.Validate(s)
}
