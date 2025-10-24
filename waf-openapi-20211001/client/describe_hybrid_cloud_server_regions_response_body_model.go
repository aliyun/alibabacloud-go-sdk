// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudServerRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v []*DescribeHybridCloudServerRegionsResponseBodyRegions) *DescribeHybridCloudServerRegionsResponseBody
	GetRegions() []*DescribeHybridCloudServerRegionsResponseBodyRegions
	SetRequestId(v string) *DescribeHybridCloudServerRegionsResponseBody
	GetRequestId() *string
}

type DescribeHybridCloudServerRegionsResponseBody struct {
	// The information about the regions.
	Regions []*DescribeHybridCloudServerRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0FBBDE11-C35F-531B-96BA-64CA****C875
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHybridCloudServerRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudServerRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudServerRegionsResponseBody) GetRegions() []*DescribeHybridCloudServerRegionsResponseBodyRegions {
	return s.Regions
}

func (s *DescribeHybridCloudServerRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridCloudServerRegionsResponseBody) SetRegions(v []*DescribeHybridCloudServerRegionsResponseBodyRegions) *DescribeHybridCloudServerRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeHybridCloudServerRegionsResponseBody) SetRequestId(v string) *DescribeHybridCloudServerRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridCloudServerRegionsResponseBody) Validate() error {
	if s.Regions != nil {
		for _, item := range s.Regions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHybridCloudServerRegionsResponseBodyRegions struct {
	// The code of the region.
	//
	// example:
	//
	// 1
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// aliyun
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeHybridCloudServerRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudServerRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudServerRegionsResponseBodyRegions) GetCode() *int32 {
	return s.Code
}

func (s *DescribeHybridCloudServerRegionsResponseBodyRegions) GetName() *string {
	return s.Name
}

func (s *DescribeHybridCloudServerRegionsResponseBodyRegions) SetCode(v int32) *DescribeHybridCloudServerRegionsResponseBodyRegions {
	s.Code = &v
	return s
}

func (s *DescribeHybridCloudServerRegionsResponseBodyRegions) SetName(v string) *DescribeHybridCloudServerRegionsResponseBodyRegions {
	s.Name = &v
	return s
}

func (s *DescribeHybridCloudServerRegionsResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
