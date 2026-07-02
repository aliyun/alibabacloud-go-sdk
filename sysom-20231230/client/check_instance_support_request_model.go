// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceSupportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*string) *CheckInstanceSupportRequest
	GetInstances() []*string
	SetRegion(v string) *CheckInstanceSupportRequest
	GetRegion() *string
}

type CheckInstanceSupportRequest struct {
	// The list of instance IDs to check.
	//
	// example:
	//
	// ["i-2zxxxxxx"]
	Instances []*string `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	// The region to which the instances belong. All instance IDs specified in instances must belong to the same region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s CheckInstanceSupportRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceSupportRequest) GoString() string {
	return s.String()
}

func (s *CheckInstanceSupportRequest) GetInstances() []*string {
	return s.Instances
}

func (s *CheckInstanceSupportRequest) GetRegion() *string {
	return s.Region
}

func (s *CheckInstanceSupportRequest) SetInstances(v []*string) *CheckInstanceSupportRequest {
	s.Instances = v
	return s
}

func (s *CheckInstanceSupportRequest) SetRegion(v string) *CheckInstanceSupportRequest {
	s.Region = &v
	return s
}

func (s *CheckInstanceSupportRequest) Validate() error {
	return dara.Validate(s)
}
