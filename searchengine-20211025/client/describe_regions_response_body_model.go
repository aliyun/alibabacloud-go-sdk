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
	// id of request
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
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
	return dara.Validate(s)
}

type DescribeRegionsResponseBodyResult struct {
	// The endpoint of the region.
	//
	// example:
	//
	// endpoint
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"localName,omitempty" xml:"localName,omitempty"`
	// The ID of the region. Valid values:
	//
	// cn-hangzhou: China (Hangzhou)
	//
	// cn-shanghai: China (Shanghai)
	//
	// cn-qingdao: China (Qingdao)
	//
	// cn-beijing: China (Beijing)
	//
	// cn-zhangjiakou: China (Zhangjiakou)
	//
	// cn-shenzhen: China (Shenzhen)
	//
	// ap-southeast-1: Singapore (Singapore)
	//
	// cn-internal: Internal Center
	//
	// cn-zhangbei-in: Internal Center (Zhangjiakou)
	//
	// us-west-1-in: Internal Center (US)
	//
	// rus-west-1-in: Internal Center (Russia)
	//
	// cn-daily: Daily Environment
	//
	// cn-test: Joint Debugging
	//
	// pre-hangzhou: China (Hangzhou)-Staging
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

func (s *DescribeRegionsResponseBodyResult) GetEndpoint() *string {
	return s.Endpoint
}

func (s *DescribeRegionsResponseBodyResult) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeRegionsResponseBodyResult) GetRegionId() *string {
	return s.RegionId
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
