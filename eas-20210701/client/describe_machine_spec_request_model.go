// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMachineSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceTypes(v []*string) *DescribeMachineSpecRequest
	GetInstanceTypes() []*string
}

type DescribeMachineSpecRequest struct {
	// Deprecated
	//
	// This parameter is deprecated.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
}

func (s DescribeMachineSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMachineSpecRequest) GoString() string {
	return s.String()
}

func (s *DescribeMachineSpecRequest) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *DescribeMachineSpecRequest) SetInstanceTypes(v []*string) *DescribeMachineSpecRequest {
	s.InstanceTypes = v
	return s
}

func (s *DescribeMachineSpecRequest) Validate() error {
	return dara.Validate(s)
}
