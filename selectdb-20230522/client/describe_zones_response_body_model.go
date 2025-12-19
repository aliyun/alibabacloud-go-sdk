// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeZonesResponseBodyItems) *DescribeZonesResponseBody
	GetItems() []*DescribeZonesResponseBodyItems
	SetMaxResults(v int32) *DescribeZonesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeZonesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeZonesResponseBody
	GetRequestId() *string
}

type DescribeZonesResponseBody struct {
	Items []*DescribeZonesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 40831b4f-d91d-4796-9589-ad306ec528d5
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ADF42B18-43FD-5100-83A9-BE81AB70C863
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) GetItems() []*DescribeZonesResponseBodyItems {
	return s.Items
}

func (s *DescribeZonesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeZonesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeZonesResponseBody) SetItems(v []*DescribeZonesResponseBodyItems) *DescribeZonesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeZonesResponseBody) SetMaxResults(v int32) *DescribeZonesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeZonesResponseBody) SetNextToken(v string) *DescribeZonesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeZonesResponseBodyItems struct {
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-beijing-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeZonesResponseBodyItems) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeZonesResponseBodyItems) SetRegionId(v string) *DescribeZonesResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeZonesResponseBodyItems) SetZoneId(v string) *DescribeZonesResponseBodyItems {
	s.ZoneId = &v
	return s
}

func (s *DescribeZonesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
