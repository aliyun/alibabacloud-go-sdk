// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGuestClusterPodsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPodList(v []*string) *DescribeGuestClusterPodsResponseBody
	GetPodList() []*string
	SetRequestId(v string) *DescribeGuestClusterPodsResponseBody
	GetRequestId() *string
}

type DescribeGuestClusterPodsResponseBody struct {
	// The list of the names of the queried pods.
	PodList []*string `json:"PodList,omitempty" xml:"PodList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// EDDC0D86-2FC3-56FB-9213-96EB0A3523F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGuestClusterPodsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGuestClusterPodsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGuestClusterPodsResponseBody) GetPodList() []*string {
	return s.PodList
}

func (s *DescribeGuestClusterPodsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGuestClusterPodsResponseBody) SetPodList(v []*string) *DescribeGuestClusterPodsResponseBody {
	s.PodList = v
	return s
}

func (s *DescribeGuestClusterPodsResponseBody) SetRequestId(v string) *DescribeGuestClusterPodsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGuestClusterPodsResponseBody) Validate() error {
	return dara.Validate(s)
}
