// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultimodalFineTuneDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *UpdateMultimodalFineTuneDatasetRequest
	GetDBClusterId() *string
	SetDatasetDescription(v string) *UpdateMultimodalFineTuneDatasetRequest
	GetDatasetDescription() *string
	SetDatasetId(v string) *UpdateMultimodalFineTuneDatasetRequest
	GetDatasetId() *string
	SetDatasetName(v string) *UpdateMultimodalFineTuneDatasetRequest
	GetDatasetName() *string
}

type UpdateMultimodalFineTuneDatasetRequest struct {
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
	// ds-****
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// 车辆
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
}

func (s UpdateMultimodalFineTuneDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultimodalFineTuneDatasetRequest) GoString() string {
	return s.String()
}

func (s *UpdateMultimodalFineTuneDatasetRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UpdateMultimodalFineTuneDatasetRequest) GetDatasetDescription() *string {
	return s.DatasetDescription
}

func (s *UpdateMultimodalFineTuneDatasetRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *UpdateMultimodalFineTuneDatasetRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *UpdateMultimodalFineTuneDatasetRequest) SetDBClusterId(v string) *UpdateMultimodalFineTuneDatasetRequest {
	s.DBClusterId = &v
	return s
}

func (s *UpdateMultimodalFineTuneDatasetRequest) SetDatasetDescription(v string) *UpdateMultimodalFineTuneDatasetRequest {
	s.DatasetDescription = &v
	return s
}

func (s *UpdateMultimodalFineTuneDatasetRequest) SetDatasetId(v string) *UpdateMultimodalFineTuneDatasetRequest {
	s.DatasetId = &v
	return s
}

func (s *UpdateMultimodalFineTuneDatasetRequest) SetDatasetName(v string) *UpdateMultimodalFineTuneDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateMultimodalFineTuneDatasetRequest) Validate() error {
	return dara.Validate(s)
}
