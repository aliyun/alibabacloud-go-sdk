// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossCloudRegionMappingToAliyunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunRegionId(v string) *DescribeCrossCloudRegionMappingToAliyunRequest
	GetAliyunRegionId() *string
	SetCloudProvider(v string) *DescribeCrossCloudRegionMappingToAliyunRequest
	GetCloudProvider() *string
	SetCrossCloudRegionId(v string) *DescribeCrossCloudRegionMappingToAliyunRequest
	GetCrossCloudRegionId() *string
}

type DescribeCrossCloudRegionMappingToAliyunRequest struct {
	// example:
	//
	// cn-beijing
	AliyunRegionId *string `json:"AliyunRegionId,omitempty" xml:"AliyunRegionId,omitempty"`
	// example:
	//
	// HuaweiCloud
	CloudProvider *string `json:"CloudProvider,omitempty" xml:"CloudProvider,omitempty"`
	// example:
	//
	// cn-east-3
	CrossCloudRegionId *string `json:"CrossCloudRegionId,omitempty" xml:"CrossCloudRegionId,omitempty"`
}

func (s DescribeCrossCloudRegionMappingToAliyunRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossCloudRegionMappingToAliyunRequest) GoString() string {
	return s.String()
}

func (s *DescribeCrossCloudRegionMappingToAliyunRequest) GetAliyunRegionId() *string {
	return s.AliyunRegionId
}

func (s *DescribeCrossCloudRegionMappingToAliyunRequest) GetCloudProvider() *string {
	return s.CloudProvider
}

func (s *DescribeCrossCloudRegionMappingToAliyunRequest) GetCrossCloudRegionId() *string {
	return s.CrossCloudRegionId
}

func (s *DescribeCrossCloudRegionMappingToAliyunRequest) SetAliyunRegionId(v string) *DescribeCrossCloudRegionMappingToAliyunRequest {
	s.AliyunRegionId = &v
	return s
}

func (s *DescribeCrossCloudRegionMappingToAliyunRequest) SetCloudProvider(v string) *DescribeCrossCloudRegionMappingToAliyunRequest {
	s.CloudProvider = &v
	return s
}

func (s *DescribeCrossCloudRegionMappingToAliyunRequest) SetCrossCloudRegionId(v string) *DescribeCrossCloudRegionMappingToAliyunRequest {
	s.CrossCloudRegionId = &v
	return s
}

func (s *DescribeCrossCloudRegionMappingToAliyunRequest) Validate() error {
	return dara.Validate(s)
}
