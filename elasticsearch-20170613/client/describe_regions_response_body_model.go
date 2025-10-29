// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRegionsResponseBody
	GetRequestId() *string
	SetResult(v []*DescribeRegionsResponseBodyResult) *DescribeRegionsResponseBody
	GetResult() []*DescribeRegionsResponseBodyResult
}

type DescribeRegionsResponseBody struct {
	// The available status of the region.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1ADFG
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The endpoint of the region.
	Result []*DescribeRegionsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsResponseBody) GetResult() []*DescribeRegionsResponseBodyResult {
	return s.Result
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetResult(v []*DescribeRegionsResponseBodyResult) *DescribeRegionsResponseBody {
	s.Result = v
	return s
}

func (s *DescribeRegionsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRegionsResponseBodyResult struct {
	// example:
	//
	// https://elasticsearch-cn-hangzhou.console.aliyun.com
	ConsoleEndpoint *string `json:"consoleEndpoint,omitempty" xml:"consoleEndpoint,omitempty"`
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"localName,omitempty" xml:"localName,omitempty"`
	// example:
	//
	// elasticsearch.cn-hangzhou.aliyuncs.com
	RegionEndpoint *string `json:"regionEndpoint,omitempty" xml:"regionEndpoint,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The endpoint of the region that is exposed in the console.
	//
	// example:
	//
	// available
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeRegionsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyResult) GetConsoleEndpoint() *string {
	return s.ConsoleEndpoint
}

func (s *DescribeRegionsResponseBodyResult) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeRegionsResponseBodyResult) GetRegionEndpoint() *string {
	return s.RegionEndpoint
}

func (s *DescribeRegionsResponseBodyResult) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *DescribeRegionsResponseBodyResult) SetConsoleEndpoint(v string) *DescribeRegionsResponseBodyResult {
	s.ConsoleEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetLocalName(v string) *DescribeRegionsResponseBodyResult {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyResult {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetRegionId(v string) *DescribeRegionsResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetStatus(v string) *DescribeRegionsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
