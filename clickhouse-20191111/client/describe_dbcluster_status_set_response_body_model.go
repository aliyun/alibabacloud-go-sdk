// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterStatusSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDBClusterStatusSetResponseBody
	GetRequestId() *string
	SetStatusSet(v []*string) *DescribeDBClusterStatusSetResponseBody
	GetStatusSet() []*string
}

type DescribeDBClusterStatusSetResponseBody struct {
	// example:
	//
	// DE309AA1-BD83-5E1F-9945-8A4D336E0829
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatusSet []*string `json:"StatusSet,omitempty" xml:"StatusSet,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterStatusSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterStatusSetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterStatusSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterStatusSetResponseBody) GetStatusSet() []*string {
	return s.StatusSet
}

func (s *DescribeDBClusterStatusSetResponseBody) SetRequestId(v string) *DescribeDBClusterStatusSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterStatusSetResponseBody) SetStatusSet(v []*string) *DescribeDBClusterStatusSetResponseBody {
	s.StatusSet = v
	return s
}

func (s *DescribeDBClusterStatusSetResponseBody) Validate() error {
	return dara.Validate(s)
}
