// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultimodalFineTuneDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteMultimodalFineTuneDatasetRequest
	GetDBClusterId() *string
	SetDatasetId(v string) *DeleteMultimodalFineTuneDatasetRequest
	GetDatasetId() *string
	SetSourceRegionId(v string) *DeleteMultimodalFineTuneDatasetRequest
	GetSourceRegionId() *string
}

type DeleteMultimodalFineTuneDatasetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ds-***
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// cn-beijing
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
}

func (s DeleteMultimodalFineTuneDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultimodalFineTuneDatasetRequest) GoString() string {
	return s.String()
}

func (s *DeleteMultimodalFineTuneDatasetRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteMultimodalFineTuneDatasetRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *DeleteMultimodalFineTuneDatasetRequest) GetSourceRegionId() *string {
	return s.SourceRegionId
}

func (s *DeleteMultimodalFineTuneDatasetRequest) SetDBClusterId(v string) *DeleteMultimodalFineTuneDatasetRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteMultimodalFineTuneDatasetRequest) SetDatasetId(v string) *DeleteMultimodalFineTuneDatasetRequest {
	s.DatasetId = &v
	return s
}

func (s *DeleteMultimodalFineTuneDatasetRequest) SetSourceRegionId(v string) *DeleteMultimodalFineTuneDatasetRequest {
	s.SourceRegionId = &v
	return s
}

func (s *DeleteMultimodalFineTuneDatasetRequest) Validate() error {
	return dara.Validate(s)
}
