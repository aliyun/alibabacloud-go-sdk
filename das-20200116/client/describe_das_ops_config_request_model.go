// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDasOpsConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDasOpsConfigRequest
	GetInstanceId() *string
}

type DescribeDasOpsConfigRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1u5mas9exx7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDasOpsConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDasOpsConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDasOpsConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDasOpsConfigRequest) SetInstanceId(v string) *DescribeDasOpsConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDasOpsConfigRequest) Validate() error {
	return dara.Validate(s)
}
