// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *DescribeDBClusterConfigResponseBody
	GetConfig() *string
	SetRequestId(v string) *DescribeDBClusterConfigResponseBody
	GetRequestId() *string
}

type DescribeDBClusterConfigResponseBody struct {
	// The information about the parameter settings of the cluster.
	//
	// example:
	//
	// [ { "name": "keep_alive_timeout", "defaultValue": 300, "currentValue": 300, "restart": true, "valueRange": ">0", "desc": "The number of seconds that ClickHouse waits for incoming requests before closing the connection." }, ... ,{ "name": "max_partition_size_to_drop", "defaultValue": 0, "currentValue": 0, "restart": true, "valueRange": ">=0", "desc": "If the size of a MergeTree partition exceeds max_partition_size_to_drop (in bytes), you can’t delete it using a DROP query." } ]
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9A23C87D-87DF-4DA0-A50E-CB13F4F7923D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigResponseBody) GetConfig() *string {
	return s.Config
}

func (s *DescribeDBClusterConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterConfigResponseBody) SetConfig(v string) *DescribeDBClusterConfigResponseBody {
	s.Config = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBody) SetRequestId(v string) *DescribeDBClusterConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
