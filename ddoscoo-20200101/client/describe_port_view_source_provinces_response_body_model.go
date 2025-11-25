// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortViewSourceProvincesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePortViewSourceProvincesResponseBody
	GetRequestId() *string
	SetSourceProvinces(v []*DescribePortViewSourceProvincesResponseBodySourceProvinces) *DescribePortViewSourceProvincesResponseBody
	GetSourceProvinces() []*DescribePortViewSourceProvincesResponseBodySourceProvinces
}

type DescribePortViewSourceProvincesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the administrative region in China from which the requests are sent.
	SourceProvinces []*DescribePortViewSourceProvincesResponseBodySourceProvinces `json:"SourceProvinces,omitempty" xml:"SourceProvinces,omitempty" type:"Repeated"`
}

func (s DescribePortViewSourceProvincesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePortViewSourceProvincesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceProvincesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePortViewSourceProvincesResponseBody) GetSourceProvinces() []*DescribePortViewSourceProvincesResponseBodySourceProvinces {
	return s.SourceProvinces
}

func (s *DescribePortViewSourceProvincesResponseBody) SetRequestId(v string) *DescribePortViewSourceProvincesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortViewSourceProvincesResponseBody) SetSourceProvinces(v []*DescribePortViewSourceProvincesResponseBodySourceProvinces) *DescribePortViewSourceProvincesResponseBody {
	s.SourceProvinces = v
	return s
}

func (s *DescribePortViewSourceProvincesResponseBody) Validate() error {
	if s.SourceProvinces != nil {
		for _, item := range s.SourceProvinces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePortViewSourceProvincesResponseBodySourceProvinces struct {
	// The total number of requests that are sent from the ISP.
	//
	// > This parameter does not indicate the accurate number of requests. You can use this parameter to calculate the proportion of requests from different administrative regions in China.
	//
	// example:
	//
	// 3390671
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the administrative region in China from which the requests are sent. For example, **110000*	- indicates Beijing, and **120000*	- indicates Tianjin.
	//
	// > For more information, see [Location parameters](https://help.aliyun.com/document_detail/167926.html).
	//
	// example:
	//
	// 440000
	ProvinceId *string `json:"ProvinceId,omitempty" xml:"ProvinceId,omitempty"`
}

func (s DescribePortViewSourceProvincesResponseBodySourceProvinces) String() string {
	return dara.Prettify(s)
}

func (s DescribePortViewSourceProvincesResponseBodySourceProvinces) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceProvincesResponseBodySourceProvinces) GetCount() *int64 {
	return s.Count
}

func (s *DescribePortViewSourceProvincesResponseBodySourceProvinces) GetProvinceId() *string {
	return s.ProvinceId
}

func (s *DescribePortViewSourceProvincesResponseBodySourceProvinces) SetCount(v int64) *DescribePortViewSourceProvincesResponseBodySourceProvinces {
	s.Count = &v
	return s
}

func (s *DescribePortViewSourceProvincesResponseBodySourceProvinces) SetProvinceId(v string) *DescribePortViewSourceProvincesResponseBodySourceProvinces {
	s.ProvinceId = &v
	return s
}

func (s *DescribePortViewSourceProvincesResponseBodySourceProvinces) Validate() error {
	return dara.Validate(s)
}
