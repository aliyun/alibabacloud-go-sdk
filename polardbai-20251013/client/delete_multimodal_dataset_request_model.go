// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultimodalDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteMultimodalDatasetRequest
	GetDBClusterId() *string
	SetDatasetId(v string) *DeleteMultimodalDatasetRequest
	GetDatasetId() *string
}

type DeleteMultimodalDatasetRequest struct {
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
	// ds-*****ab0
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
}

func (s DeleteMultimodalDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultimodalDatasetRequest) GoString() string {
	return s.String()
}

func (s *DeleteMultimodalDatasetRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteMultimodalDatasetRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *DeleteMultimodalDatasetRequest) SetDBClusterId(v string) *DeleteMultimodalDatasetRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteMultimodalDatasetRequest) SetDatasetId(v string) *DeleteMultimodalDatasetRequest {
	s.DatasetId = &v
	return s
}

func (s *DeleteMultimodalDatasetRequest) Validate() error {
	return dara.Validate(s)
}
