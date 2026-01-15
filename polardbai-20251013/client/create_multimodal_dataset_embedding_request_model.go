// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultimodalDatasetEmbeddingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateMultimodalDatasetEmbeddingRequest
	GetDBClusterId() *string
	SetDatasetId(v string) *CreateMultimodalDatasetEmbeddingRequest
	GetDatasetId() *string
	SetModel(v string) *CreateMultimodalDatasetEmbeddingRequest
	GetModel() *string
	SetModelMode(v string) *CreateMultimodalDatasetEmbeddingRequest
	GetModelMode() *string
}

type CreateMultimodalDatasetEmbeddingRequest struct {
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
	// ds-3h6bm*****af60
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// default
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// flash
	ModelMode *string `json:"ModelMode,omitempty" xml:"ModelMode,omitempty"`
}

func (s CreateMultimodalDatasetEmbeddingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMultimodalDatasetEmbeddingRequest) GoString() string {
	return s.String()
}

func (s *CreateMultimodalDatasetEmbeddingRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateMultimodalDatasetEmbeddingRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *CreateMultimodalDatasetEmbeddingRequest) GetModel() *string {
	return s.Model
}

func (s *CreateMultimodalDatasetEmbeddingRequest) GetModelMode() *string {
	return s.ModelMode
}

func (s *CreateMultimodalDatasetEmbeddingRequest) SetDBClusterId(v string) *CreateMultimodalDatasetEmbeddingRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateMultimodalDatasetEmbeddingRequest) SetDatasetId(v string) *CreateMultimodalDatasetEmbeddingRequest {
	s.DatasetId = &v
	return s
}

func (s *CreateMultimodalDatasetEmbeddingRequest) SetModel(v string) *CreateMultimodalDatasetEmbeddingRequest {
	s.Model = &v
	return s
}

func (s *CreateMultimodalDatasetEmbeddingRequest) SetModelMode(v string) *CreateMultimodalDatasetEmbeddingRequest {
	s.ModelMode = &v
	return s
}

func (s *CreateMultimodalDatasetEmbeddingRequest) Validate() error {
	return dara.Validate(s)
}
