// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComfyProductionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductionId(v string) *DeleteComfyProductionRequest
	GetProductionId() *string
}

type DeleteComfyProductionRequest struct {
	// example:
	//
	// 3e5bda20-5cd4-4d55-8d23-88d624a18caa
	ProductionId *string `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
}

func (s DeleteComfyProductionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteComfyProductionRequest) GoString() string {
	return s.String()
}

func (s *DeleteComfyProductionRequest) GetProductionId() *string {
	return s.ProductionId
}

func (s *DeleteComfyProductionRequest) SetProductionId(v string) *DeleteComfyProductionRequest {
	s.ProductionId = &v
	return s
}

func (s *DeleteComfyProductionRequest) Validate() error {
	return dara.Validate(s)
}
