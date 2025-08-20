// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateScratchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteTemplateScratchRequest
	GetRegionId() *string
	SetTemplateScratchId(v string) *DeleteTemplateScratchRequest
	GetTemplateScratchId() *string
}

type DeleteTemplateScratchRequest struct {
	// The region ID of the scenario.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// ts-4f83704400994409****
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
}

func (s DeleteTemplateScratchRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateScratchRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateScratchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteTemplateScratchRequest) GetTemplateScratchId() *string {
	return s.TemplateScratchId
}

func (s *DeleteTemplateScratchRequest) SetRegionId(v string) *DeleteTemplateScratchRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTemplateScratchRequest) SetTemplateScratchId(v string) *DeleteTemplateScratchRequest {
	s.TemplateScratchId = &v
	return s
}

func (s *DeleteTemplateScratchRequest) Validate() error {
	return dara.Validate(s)
}
