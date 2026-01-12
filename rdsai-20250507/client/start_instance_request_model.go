// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *StartInstanceRequest
	GetInstanceName() *string
	SetRegionId(v string) *StartInstanceRequest
	GetRegionId() *string
}

type StartInstanceRequest struct {
	// The region ID.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The operation that you want to perform. Set the value to **StartInstance**.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *StartInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartInstanceRequest) SetInstanceName(v string) *StartInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *StartInstanceRequest) SetRegionId(v string) *StartInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartInstanceRequest) Validate() error {
	return dara.Validate(s)
}
