// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiscoveredResourceBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDiscoveredResourceBatchResponseBody
	GetRequestId() *string
}

type DescribeDiscoveredResourceBatchResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6CE4ABA1-9A57-41A9-8EA9-E8B17D4671CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiscoveredResourceBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiscoveredResourceBatchResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiscoveredResourceBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiscoveredResourceBatchResponseBody) SetRequestId(v string) *DescribeDiscoveredResourceBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiscoveredResourceBatchResponseBody) Validate() error {
	return dara.Validate(s)
}
