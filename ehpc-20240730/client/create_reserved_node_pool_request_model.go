// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReservedNodePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateReservedNodePoolRequest
	GetClusterId() *string
	SetCount(v int32) *CreateReservedNodePoolRequest
	GetCount() *int32
	SetDescription(v string) *CreateReservedNodePoolRequest
	GetDescription() *string
	SetHostPostfix(v string) *CreateReservedNodePoolRequest
	GetHostPostfix() *string
	SetHostPrefix(v string) *CreateReservedNodePoolRequest
	GetHostPrefix() *string
	SetName(v string) *CreateReservedNodePoolRequest
	GetName() *string
	SetVSwitchId(v string) *CreateReservedNodePoolRequest
	GetVSwitchId() *string
}

type CreateReservedNodePoolRequest struct {
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// ed
	HostPostfix *string `json:"HostPostfix,omitempty" xml:"HostPostfix,omitempty"`
	// example:
	//
	// cloud
	HostPrefix *string `json:"HostPrefix,omitempty" xml:"HostPrefix,omitempty"`
	// example:
	//
	// nodepool
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// vsw-bp1lfcjbfb099rrjn****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateReservedNodePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateReservedNodePoolRequest) GoString() string {
	return s.String()
}

func (s *CreateReservedNodePoolRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateReservedNodePoolRequest) GetCount() *int32 {
	return s.Count
}

func (s *CreateReservedNodePoolRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateReservedNodePoolRequest) GetHostPostfix() *string {
	return s.HostPostfix
}

func (s *CreateReservedNodePoolRequest) GetHostPrefix() *string {
	return s.HostPrefix
}

func (s *CreateReservedNodePoolRequest) GetName() *string {
	return s.Name
}

func (s *CreateReservedNodePoolRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateReservedNodePoolRequest) SetClusterId(v string) *CreateReservedNodePoolRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateReservedNodePoolRequest) SetCount(v int32) *CreateReservedNodePoolRequest {
	s.Count = &v
	return s
}

func (s *CreateReservedNodePoolRequest) SetDescription(v string) *CreateReservedNodePoolRequest {
	s.Description = &v
	return s
}

func (s *CreateReservedNodePoolRequest) SetHostPostfix(v string) *CreateReservedNodePoolRequest {
	s.HostPostfix = &v
	return s
}

func (s *CreateReservedNodePoolRequest) SetHostPrefix(v string) *CreateReservedNodePoolRequest {
	s.HostPrefix = &v
	return s
}

func (s *CreateReservedNodePoolRequest) SetName(v string) *CreateReservedNodePoolRequest {
	s.Name = &v
	return s
}

func (s *CreateReservedNodePoolRequest) SetVSwitchId(v string) *CreateReservedNodePoolRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateReservedNodePoolRequest) Validate() error {
	return dara.Validate(s)
}
