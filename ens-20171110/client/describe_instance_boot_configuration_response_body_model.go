// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceBootConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v *DescribeInstanceBootConfigurationResponseBodyInstances) *DescribeInstanceBootConfigurationResponseBody
	GetInstances() *DescribeInstanceBootConfigurationResponseBodyInstances
	SetRequestId(v string) *DescribeInstanceBootConfigurationResponseBody
	GetRequestId() *string
}

type DescribeInstanceBootConfigurationResponseBody struct {
	// Schema of Response
	Instances *DescribeInstanceBootConfigurationResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// F3B261DD-3858-4D3C-877D-303ADF374600
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceBootConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceBootConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceBootConfigurationResponseBody) GetInstances() *DescribeInstanceBootConfigurationResponseBodyInstances {
	return s.Instances
}

func (s *DescribeInstanceBootConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceBootConfigurationResponseBody) SetInstances(v *DescribeInstanceBootConfigurationResponseBodyInstances) *DescribeInstanceBootConfigurationResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstanceBootConfigurationResponseBody) SetRequestId(v string) *DescribeInstanceBootConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceBootConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceBootConfigurationResponseBodyInstances struct {
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
	// disk
	BootType *string `json:"BootType,omitempty" xml:"BootType,omitempty"`
	// Specifies whether the startup depends on the disk.
	//
	// example:
	//
	// off
	DiskSet *string `json:"DiskSet,omitempty" xml:"DiskSet,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceBootConfigurationResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceBootConfigurationResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstanceBootConfigurationResponseBodyInstances) GetBootSet() *string {
	return s.BootSet
}

func (s *DescribeInstanceBootConfigurationResponseBodyInstances) GetBootType() *string {
	return s.BootType
}

func (s *DescribeInstanceBootConfigurationResponseBodyInstances) GetDiskSet() *string {
	return s.DiskSet
}

func (s *DescribeInstanceBootConfigurationResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceBootConfigurationResponseBodyInstances) SetBootSet(v string) *DescribeInstanceBootConfigurationResponseBodyInstances {
	s.BootSet = &v
	return s
}

func (s *DescribeInstanceBootConfigurationResponseBodyInstances) SetBootType(v string) *DescribeInstanceBootConfigurationResponseBodyInstances {
	s.BootType = &v
	return s
}

func (s *DescribeInstanceBootConfigurationResponseBodyInstances) SetDiskSet(v string) *DescribeInstanceBootConfigurationResponseBodyInstances {
	s.DiskSet = &v
	return s
}

func (s *DescribeInstanceBootConfigurationResponseBodyInstances) SetInstanceId(v string) *DescribeInstanceBootConfigurationResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceBootConfigurationResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}
