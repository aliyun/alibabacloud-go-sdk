// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConnectableClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeConnectableClustersResponseBody
	GetRequestId() *string
	SetResult(v []*DescribeConnectableClustersResponseBodyResult) *DescribeConnectableClustersResponseBody
	GetResult() []*DescribeConnectableClustersResponseBodyResult
}

type DescribeConnectableClustersResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeConnectableClustersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeConnectableClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConnectableClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConnectableClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConnectableClustersResponseBody) GetResult() []*DescribeConnectableClustersResponseBodyResult {
	return s.Result
}

func (s *DescribeConnectableClustersResponseBody) SetRequestId(v string) *DescribeConnectableClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConnectableClustersResponseBody) SetResult(v []*DescribeConnectableClustersResponseBodyResult) *DescribeConnectableClustersResponseBody {
	s.Result = v
	return s
}

func (s *DescribeConnectableClustersResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeConnectableClustersResponseBodyResult struct {
	// example:
	//
	// es-cn-xxx
	Instances *string `json:"instances,omitempty" xml:"instances,omitempty"`
	// example:
	//
	// vpc
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
}

func (s DescribeConnectableClustersResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeConnectableClustersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeConnectableClustersResponseBodyResult) GetInstances() *string {
	return s.Instances
}

func (s *DescribeConnectableClustersResponseBodyResult) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeConnectableClustersResponseBodyResult) SetInstances(v string) *DescribeConnectableClustersResponseBodyResult {
	s.Instances = &v
	return s
}

func (s *DescribeConnectableClustersResponseBodyResult) SetNetworkType(v string) *DescribeConnectableClustersResponseBodyResult {
	s.NetworkType = &v
	return s
}

func (s *DescribeConnectableClustersResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
