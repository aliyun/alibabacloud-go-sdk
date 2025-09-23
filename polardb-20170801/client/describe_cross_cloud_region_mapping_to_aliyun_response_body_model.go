// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossCloudRegionMappingToAliyunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCrossCloudRegionMappingList(v []*DescribeCrossCloudRegionMappingToAliyunResponseBodyCrossCloudRegionMappingList) *DescribeCrossCloudRegionMappingToAliyunResponseBody
	GetCrossCloudRegionMappingList() []*DescribeCrossCloudRegionMappingToAliyunResponseBodyCrossCloudRegionMappingList
	SetRequestId(v string) *DescribeCrossCloudRegionMappingToAliyunResponseBody
	GetRequestId() *string
}

type DescribeCrossCloudRegionMappingToAliyunResponseBody struct {
	CrossCloudRegionMappingList []*DescribeCrossCloudRegionMappingToAliyunResponseBodyCrossCloudRegionMappingList `json:"CrossCloudRegionMappingList,omitempty" xml:"CrossCloudRegionMappingList,omitempty" type:"Repeated"`
	// example:
	//
	// E56531A4-E552-40BA-9C58-137B80******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCrossCloudRegionMappingToAliyunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossCloudRegionMappingToAliyunResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCrossCloudRegionMappingToAliyunResponseBody) GetCrossCloudRegionMappingList() []*DescribeCrossCloudRegionMappingToAliyunResponseBodyCrossCloudRegionMappingList {
	return s.CrossCloudRegionMappingList
}

func (s *DescribeCrossCloudRegionMappingToAliyunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCrossCloudRegionMappingToAliyunResponseBody) SetCrossCloudRegionMappingList(v []*DescribeCrossCloudRegionMappingToAliyunResponseBodyCrossCloudRegionMappingList) *DescribeCrossCloudRegionMappingToAliyunResponseBody {
	s.CrossCloudRegionMappingList = v
	return s
}

func (s *DescribeCrossCloudRegionMappingToAliyunResponseBody) SetRequestId(v string) *DescribeCrossCloudRegionMappingToAliyunResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCrossCloudRegionMappingToAliyunResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCrossCloudRegionMappingToAliyunResponseBodyCrossCloudRegionMappingList struct {
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

func (s DescribeCrossCloudRegionMappingToAliyunResponseBodyCrossCloudRegionMappingList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossCloudRegionMappingToAliyunResponseBodyCrossCloudRegionMappingList) GoString() string {
	return s.String()
}

func (s *DescribeCrossCloudRegionMappingToAliyunResponseBodyCrossCloudRegionMappingList) GetAliyunRegionId() *string {
	return s.AliyunRegionId
}

func (s *DescribeCrossCloudRegionMappingToAliyunResponseBodyCrossCloudRegionMappingList) GetCloudProvider() *string {
	return s.CloudProvider
}

func (s *DescribeCrossCloudRegionMappingToAliyunResponseBodyCrossCloudRegionMappingList) GetCrossCloudRegionId() *string {
	return s.CrossCloudRegionId
}

func (s *DescribeCrossCloudRegionMappingToAliyunResponseBodyCrossCloudRegionMappingList) SetAliyunRegionId(v string) *DescribeCrossCloudRegionMappingToAliyunResponseBodyCrossCloudRegionMappingList {
	s.AliyunRegionId = &v
	return s
}

func (s *DescribeCrossCloudRegionMappingToAliyunResponseBodyCrossCloudRegionMappingList) SetCloudProvider(v string) *DescribeCrossCloudRegionMappingToAliyunResponseBodyCrossCloudRegionMappingList {
	s.CloudProvider = &v
	return s
}

func (s *DescribeCrossCloudRegionMappingToAliyunResponseBodyCrossCloudRegionMappingList) SetCrossCloudRegionId(v string) *DescribeCrossCloudRegionMappingToAliyunResponseBodyCrossCloudRegionMappingList {
	s.CrossCloudRegionId = &v
	return s
}

func (s *DescribeCrossCloudRegionMappingToAliyunResponseBodyCrossCloudRegionMappingList) Validate() error {
	return dara.Validate(s)
}
