// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTraceLocationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegionConfigs(v []*DescribeTraceLocationResponseBodyRegionConfigs) *DescribeTraceLocationResponseBody
	GetRegionConfigs() []*DescribeTraceLocationResponseBodyRegionConfigs
	SetRequestId(v string) *DescribeTraceLocationResponseBody
	GetRequestId() *string
}

type DescribeTraceLocationResponseBody struct {
	RegionConfigs []*DescribeTraceLocationResponseBodyRegionConfigs `json:"RegionConfigs,omitempty" xml:"RegionConfigs,omitempty" type:"Repeated"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTraceLocationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceLocationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTraceLocationResponseBody) GetRegionConfigs() []*DescribeTraceLocationResponseBodyRegionConfigs {
	return s.RegionConfigs
}

func (s *DescribeTraceLocationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTraceLocationResponseBody) SetRegionConfigs(v []*DescribeTraceLocationResponseBodyRegionConfigs) *DescribeTraceLocationResponseBody {
	s.RegionConfigs = v
	return s
}

func (s *DescribeTraceLocationResponseBody) SetRequestId(v string) *DescribeTraceLocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTraceLocationResponseBody) Validate() error {
	if s.RegionConfigs != nil {
		for _, item := range s.RegionConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTraceLocationResponseBodyRegionConfigs struct {
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeTraceLocationResponseBodyRegionConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceLocationResponseBodyRegionConfigs) GoString() string {
	return s.String()
}

func (s *DescribeTraceLocationResponseBodyRegionConfigs) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeTraceLocationResponseBodyRegionConfigs) GetUrl() *string {
	return s.Url
}

func (s *DescribeTraceLocationResponseBodyRegionConfigs) SetRegionNo(v string) *DescribeTraceLocationResponseBodyRegionConfigs {
	s.RegionNo = &v
	return s
}

func (s *DescribeTraceLocationResponseBodyRegionConfigs) SetUrl(v string) *DescribeTraceLocationResponseBodyRegionConfigs {
	s.Url = &v
	return s
}

func (s *DescribeTraceLocationResponseBodyRegionConfigs) Validate() error {
	return dara.Validate(s)
}
