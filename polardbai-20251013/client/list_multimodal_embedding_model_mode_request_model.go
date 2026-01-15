// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalEmbeddingModelModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListMultimodalEmbeddingModelModeRequest
	GetDBClusterId() *string
}

type ListMultimodalEmbeddingModelModeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s ListMultimodalEmbeddingModelModeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalEmbeddingModelModeRequest) GoString() string {
	return s.String()
}

func (s *ListMultimodalEmbeddingModelModeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListMultimodalEmbeddingModelModeRequest) SetDBClusterId(v string) *ListMultimodalEmbeddingModelModeRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListMultimodalEmbeddingModelModeRequest) Validate() error {
	return dara.Validate(s)
}
