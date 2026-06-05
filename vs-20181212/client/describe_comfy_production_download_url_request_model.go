// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyProductionDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductionId(v string) *DescribeComfyProductionDownloadUrlRequest
	GetProductionId() *string
}

type DescribeComfyProductionDownloadUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3e5bda20-5cd4-4d55-8d23-88d624a18caa
	ProductionId *string `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
}

func (s DescribeComfyProductionDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyProductionDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeComfyProductionDownloadUrlRequest) GetProductionId() *string {
	return s.ProductionId
}

func (s *DescribeComfyProductionDownloadUrlRequest) SetProductionId(v string) *DescribeComfyProductionDownloadUrlRequest {
	s.ProductionId = &v
	return s
}

func (s *DescribeComfyProductionDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
