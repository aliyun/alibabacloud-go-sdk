// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCVClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRCVClusterResponseBody
	GetRequestId() *string
	SetVClusterStatus(v string) *DescribeRCVClusterResponseBody
	GetVClusterStatus() *string
}

type DescribeRCVClusterResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VClusterStatus *string `json:"VClusterStatus,omitempty" xml:"VClusterStatus,omitempty"`
}

func (s DescribeRCVClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCVClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCVClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCVClusterResponseBody) GetVClusterStatus() *string {
	return s.VClusterStatus
}

func (s *DescribeRCVClusterResponseBody) SetRequestId(v string) *DescribeRCVClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCVClusterResponseBody) SetVClusterStatus(v string) *DescribeRCVClusterResponseBody {
	s.VClusterStatus = &v
	return s
}

func (s *DescribeRCVClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
