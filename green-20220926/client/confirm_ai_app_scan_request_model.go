// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmAiAppScanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *ConfirmAiAppScanRequest
	GetCommodityCode() *string
	SetRegionId(v string) *ConfirmAiAppScanRequest
	GetRegionId() *string
}

type ConfirmAiAppScanRequest struct {
	// The commodity code.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The ID of the region where the application resides.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ConfirmAiAppScanRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfirmAiAppScanRequest) GoString() string {
	return s.String()
}

func (s *ConfirmAiAppScanRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ConfirmAiAppScanRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfirmAiAppScanRequest) SetCommodityCode(v string) *ConfirmAiAppScanRequest {
	s.CommodityCode = &v
	return s
}

func (s *ConfirmAiAppScanRequest) SetRegionId(v string) *ConfirmAiAppScanRequest {
	s.RegionId = &v
	return s
}

func (s *ConfirmAiAppScanRequest) Validate() error {
	return dara.Validate(s)
}
