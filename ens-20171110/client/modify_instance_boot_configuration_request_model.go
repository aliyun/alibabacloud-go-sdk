// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceBootConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBootSet(v string) *ModifyInstanceBootConfigurationRequest
	GetBootSet() *string
	SetBootType(v string) *ModifyInstanceBootConfigurationRequest
	GetBootType() *string
	SetDiskSet(v string) *ModifyInstanceBootConfigurationRequest
	GetDiskSet() *string
	SetInstanceId(v string) *ModifyInstanceBootConfigurationRequest
	GetInstanceId() *string
}

type ModifyInstanceBootConfigurationRequest struct {
	// example:
	//
	// legacy
	BootSet *string `json:"BootSet,omitempty" xml:"BootSet,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxe
	BootType *string `json:"BootType,omitempty" xml:"BootType,omitempty"`
	// example:
	//
	// on
	DiskSet *string `json:"DiskSet,omitempty" xml:"DiskSet,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// i-instance****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyInstanceBootConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceBootConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceBootConfigurationRequest) GetBootSet() *string {
	return s.BootSet
}

func (s *ModifyInstanceBootConfigurationRequest) GetBootType() *string {
	return s.BootType
}

func (s *ModifyInstanceBootConfigurationRequest) GetDiskSet() *string {
	return s.DiskSet
}

func (s *ModifyInstanceBootConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceBootConfigurationRequest) SetBootSet(v string) *ModifyInstanceBootConfigurationRequest {
	s.BootSet = &v
	return s
}

func (s *ModifyInstanceBootConfigurationRequest) SetBootType(v string) *ModifyInstanceBootConfigurationRequest {
	s.BootType = &v
	return s
}

func (s *ModifyInstanceBootConfigurationRequest) SetDiskSet(v string) *ModifyInstanceBootConfigurationRequest {
	s.DiskSet = &v
	return s
}

func (s *ModifyInstanceBootConfigurationRequest) SetInstanceId(v string) *ModifyInstanceBootConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceBootConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
