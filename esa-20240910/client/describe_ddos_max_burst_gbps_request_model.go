// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosMaxBurstGbpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDdosMaxBurstGbpsRequest
	GetInstanceId() *string
}

type DescribeDdosMaxBurstGbpsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// esa-site-a71k7bw19dz4
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDdosMaxBurstGbpsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosMaxBurstGbpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosMaxBurstGbpsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDdosMaxBurstGbpsRequest) SetInstanceId(v string) *DescribeDdosMaxBurstGbpsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDdosMaxBurstGbpsRequest) Validate() error {
	return dara.Validate(s)
}
