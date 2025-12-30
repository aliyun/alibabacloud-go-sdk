// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultimodalDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *UpdateMultimodalDatasetRequest
	GetDBClusterId() *string
	SetDatasetDescription(v string) *UpdateMultimodalDatasetRequest
	GetDatasetDescription() *string
	SetDatasetId(v string) *UpdateMultimodalDatasetRequest
	GetDatasetId() *string
	SetDatasetName(v string) *UpdateMultimodalDatasetRequest
	GetDatasetName() *string
}

type UpdateMultimodalDatasetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 车辆图片
	DatasetDescription *string `json:"DatasetDescription,omitempty" xml:"DatasetDescription,omitempty"`
	// example:
	//
	// 车辆图片
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// 车辆图片
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
}

func (s UpdateMultimodalDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultimodalDatasetRequest) GoString() string {
	return s.String()
}

func (s *UpdateMultimodalDatasetRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UpdateMultimodalDatasetRequest) GetDatasetDescription() *string {
	return s.DatasetDescription
}

func (s *UpdateMultimodalDatasetRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *UpdateMultimodalDatasetRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *UpdateMultimodalDatasetRequest) SetDBClusterId(v string) *UpdateMultimodalDatasetRequest {
	s.DBClusterId = &v
	return s
}

func (s *UpdateMultimodalDatasetRequest) SetDatasetDescription(v string) *UpdateMultimodalDatasetRequest {
	s.DatasetDescription = &v
	return s
}

func (s *UpdateMultimodalDatasetRequest) SetDatasetId(v string) *UpdateMultimodalDatasetRequest {
	s.DatasetId = &v
	return s
}

func (s *UpdateMultimodalDatasetRequest) SetDatasetName(v string) *UpdateMultimodalDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateMultimodalDatasetRequest) Validate() error {
	return dara.Validate(s)
}
