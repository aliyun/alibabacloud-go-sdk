// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAckClusterNamespacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAckNamespaces(v []*string) *DescribeAckClusterNamespacesResponseBody
	GetAckNamespaces() []*string
	SetRequestId(v string) *DescribeAckClusterNamespacesResponseBody
	GetRequestId() *string
}

type DescribeAckClusterNamespacesResponseBody struct {
	AckNamespaces []*string `json:"AckNamespaces,omitempty" xml:"AckNamespaces,omitempty" type:"Repeated"`
	// example:
	//
	// 133173B9-8010-5DF5-8B93-********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAckClusterNamespacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckClusterNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAckClusterNamespacesResponseBody) GetAckNamespaces() []*string {
	return s.AckNamespaces
}

func (s *DescribeAckClusterNamespacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAckClusterNamespacesResponseBody) SetAckNamespaces(v []*string) *DescribeAckClusterNamespacesResponseBody {
	s.AckNamespaces = v
	return s
}

func (s *DescribeAckClusterNamespacesResponseBody) SetRequestId(v string) *DescribeAckClusterNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAckClusterNamespacesResponseBody) Validate() error {
	return dara.Validate(s)
}
