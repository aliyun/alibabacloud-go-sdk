// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceBootConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBootSet(v string) *DescribeInstanceBootConfigurationRequest
	GetBootSet() *string
	SetBootType(v string) *DescribeInstanceBootConfigurationRequest
	GetBootType() *string
	SetDiskSet(v string) *DescribeInstanceBootConfigurationRequest
	GetDiskSet() *string
	SetInstanceId(v string) *DescribeInstanceBootConfigurationRequest
	GetInstanceId() *string
}

type DescribeInstanceBootConfigurationRequest struct {
	// The startup method.
	//
	// example:
	//
	// legacy
	BootSet *string `json:"BootSet,omitempty" xml:"BootSet,omitempty"`
	// The startup type.
	//
	// example:
	//
	// pxe
	BootType *string `json:"BootType,omitempty" xml:"BootType,omitempty"`
	// Specifies whether the startup depends on the disk.
	//
	// example:
	//
	// on
	DiskSet *string `json:"DiskSet,omitempty" xml:"DiskSet,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceBootConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceBootConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceBootConfigurationRequest) GetBootSet() *string {
	return s.BootSet
}

func (s *DescribeInstanceBootConfigurationRequest) GetBootType() *string {
	return s.BootType
}

func (s *DescribeInstanceBootConfigurationRequest) GetDiskSet() *string {
	return s.DiskSet
}

func (s *DescribeInstanceBootConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceBootConfigurationRequest) SetBootSet(v string) *DescribeInstanceBootConfigurationRequest {
	s.BootSet = &v
	return s
}

func (s *DescribeInstanceBootConfigurationRequest) SetBootType(v string) *DescribeInstanceBootConfigurationRequest {
	s.BootType = &v
	return s
}

func (s *DescribeInstanceBootConfigurationRequest) SetDiskSet(v string) *DescribeInstanceBootConfigurationRequest {
	s.DiskSet = &v
	return s
}

func (s *DescribeInstanceBootConfigurationRequest) SetInstanceId(v string) *DescribeInstanceBootConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceBootConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
