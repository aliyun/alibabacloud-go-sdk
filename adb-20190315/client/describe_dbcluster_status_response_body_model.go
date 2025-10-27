// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDBClusterStatusResponseBody
	GetRequestId() *string
	SetStatus(v []*string) *DescribeDBClusterStatusResponseBody
	GetStatus() []*string
}

type DescribeDBClusterStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEAU
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of clusters.
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterStatusResponseBody) GetStatus() []*string {
	return s.Status
}

func (s *DescribeDBClusterStatusResponseBody) SetRequestId(v string) *DescribeDBClusterStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterStatusResponseBody) SetStatus(v []*string) *DescribeDBClusterStatusResponseBody {
	s.Status = v
	return s
}

func (s *DescribeDBClusterStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
