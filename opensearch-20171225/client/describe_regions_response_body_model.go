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
	// The request ID.
	//
	// example:
	//
	// 3B7E42BD-1D63-2F6B-C8E0-41BACEA76EEB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The results returned.
	Result []*DescribeRegionsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
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
	// The console URL.
	//
	// example:
	//
	// https://opensearch-cn-hangzhou.console.aliyun.com
	ConsoleUrl *string `json:"consoleUrl,omitempty" xml:"consoleUrl,omitempty"`
	// The endpoint.
	//
	// example:
	//
	// opensearch.cn-hangzhou.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The region name.
	//
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"localName,omitempty" xml:"localName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DescribeRegionsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyResult) GetConsoleUrl() *string {
	return s.ConsoleUrl
}

func (s *DescribeRegionsResponseBodyResult) GetEndpoint() *string {
	return s.Endpoint
}

func (s *DescribeRegionsResponseBodyResult) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeRegionsResponseBodyResult) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsResponseBodyResult) SetConsoleUrl(v string) *DescribeRegionsResponseBodyResult {
	s.ConsoleUrl = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetEndpoint(v string) *DescribeRegionsResponseBodyResult {
	s.Endpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetLocalName(v string) *DescribeRegionsResponseBodyResult {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetRegionId(v string) *DescribeRegionsResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
