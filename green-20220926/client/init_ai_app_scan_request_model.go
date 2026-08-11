// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitAiAppScanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *InitAiAppScanRequest
	GetChannel() *string
	SetCommodityCode(v string) *InitAiAppScanRequest
	GetCommodityCode() *string
	SetRegionId(v string) *InitAiAppScanRequest
	GetRegionId() *string
}

type InitAiAppScanRequest struct {
	// The channel type.
	//
	// example:
	//
	// bailian
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The commodity code.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InitAiAppScanRequest) String() string {
	return dara.Prettify(s)
}

func (s InitAiAppScanRequest) GoString() string {
	return s.String()
}

func (s *InitAiAppScanRequest) GetChannel() *string {
	return s.Channel
}

func (s *InitAiAppScanRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *InitAiAppScanRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InitAiAppScanRequest) SetChannel(v string) *InitAiAppScanRequest {
	s.Channel = &v
	return s
}

func (s *InitAiAppScanRequest) SetCommodityCode(v string) *InitAiAppScanRequest {
	s.CommodityCode = &v
	return s
}

func (s *InitAiAppScanRequest) SetRegionId(v string) *InitAiAppScanRequest {
	s.RegionId = &v
	return s
}

func (s *InitAiAppScanRequest) Validate() error {
	return dara.Validate(s)
}
