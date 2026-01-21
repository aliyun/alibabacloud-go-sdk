// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionResourceTypeAutoEnableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegionResourceAutoEnable(v map[string]map[string]interface{}) *DescribeRegionResourceTypeAutoEnableResponseBody
	GetRegionResourceAutoEnable() map[string]map[string]interface{}
	SetRequestId(v string) *DescribeRegionResourceTypeAutoEnableResponseBody
	GetRequestId() *string
}

type DescribeRegionResourceTypeAutoEnableResponseBody struct {
	RegionResourceAutoEnable map[string]map[string]interface{} `json:"RegionResourceAutoEnable,omitempty" xml:"RegionResourceAutoEnable,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B2841452-CB8D-4F7D-B247-38E1CF7334F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionResourceTypeAutoEnableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceTypeAutoEnableResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceTypeAutoEnableResponseBody) GetRegionResourceAutoEnable() map[string]map[string]interface{} {
	return s.RegionResourceAutoEnable
}

func (s *DescribeRegionResourceTypeAutoEnableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionResourceTypeAutoEnableResponseBody) SetRegionResourceAutoEnable(v map[string]map[string]interface{}) *DescribeRegionResourceTypeAutoEnableResponseBody {
	s.RegionResourceAutoEnable = v
	return s
}

func (s *DescribeRegionResourceTypeAutoEnableResponseBody) SetRequestId(v string) *DescribeRegionResourceTypeAutoEnableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionResourceTypeAutoEnableResponseBody) Validate() error {
	return dara.Validate(s)
}
