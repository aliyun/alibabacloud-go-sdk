// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosMaxBurstGbpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDdosMaxBurstGbpsResponseBody
	GetInstanceId() *string
	SetMaxBurstGbps(v string) *DescribeDdosMaxBurstGbpsResponseBody
	GetMaxBurstGbps() *string
	SetRequestId(v string) *DescribeDdosMaxBurstGbpsResponseBody
	GetRequestId() *string
}

type DescribeDdosMaxBurstGbpsResponseBody struct {
	// example:
	//
	// esa-site-a71k7bw19dz4
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 300
	MaxBurstGbps *string `json:"MaxBurstGbps,omitempty" xml:"MaxBurstGbps,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B5D71671-B074-5702-A0F5-B923920FDDD4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDdosMaxBurstGbpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosMaxBurstGbpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosMaxBurstGbpsResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDdosMaxBurstGbpsResponseBody) GetMaxBurstGbps() *string {
	return s.MaxBurstGbps
}

func (s *DescribeDdosMaxBurstGbpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDdosMaxBurstGbpsResponseBody) SetInstanceId(v string) *DescribeDdosMaxBurstGbpsResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDdosMaxBurstGbpsResponseBody) SetMaxBurstGbps(v string) *DescribeDdosMaxBurstGbpsResponseBody {
	s.MaxBurstGbps = &v
	return s
}

func (s *DescribeDdosMaxBurstGbpsResponseBody) SetRequestId(v string) *DescribeDdosMaxBurstGbpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosMaxBurstGbpsResponseBody) Validate() error {
	return dara.Validate(s)
}
