// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableResources(v *DescribeAvailableResourceResponseBodyAvailableResources) *DescribeAvailableResourceResponseBody
	GetAvailableResources() *DescribeAvailableResourceResponseBodyAvailableResources
	SetRequestId(v string) *DescribeAvailableResourceResponseBody
	GetRequestId() *string
}

type DescribeAvailableResourceResponseBody struct {
	// The zones and the supported resources.
	AvailableResources *DescribeAvailableResourceResponseBodyAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 173B0EEA-22ED-4EE2-91F9-3A1CDDFFBBBA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBody) GetAvailableResources() *DescribeAvailableResourceResponseBodyAvailableResources {
	return s.AvailableResources
}

func (s *DescribeAvailableResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvailableResourceResponseBody) SetAvailableResources(v *DescribeAvailableResourceResponseBodyAvailableResources) *DescribeAvailableResourceResponseBody {
	s.AvailableResources = v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetRequestId(v string) *DescribeAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBody) Validate() error {
	if s.AvailableResources != nil {
		if err := s.AvailableResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodyAvailableResources struct {
	AvailableResource []*DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource `json:"AvailableResource,omitempty" xml:"AvailableResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableResources) GetAvailableResource() []*DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource {
	return s.AvailableResource
}

func (s *DescribeAvailableResourceResponseBodyAvailableResources) SetAvailableResource(v []*DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource) *DescribeAvailableResourceResponseBodyAvailableResources {
	s.AvailableResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableResources) Validate() error {
	if s.AvailableResource != nil {
		for _, item := range s.AvailableResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource struct {
	// The primary zone.
	//
	// example:
	//
	// cn-shanghai-a
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	// The secondary zone.
	//
	// example:
	//
	// cn-shanghai-b
	SlaveZoneId *string `json:"SlaveZoneId,omitempty" xml:"SlaveZoneId,omitempty"`
	// The supported resources.
	SupportResources *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResources `json:"SupportResources,omitempty" xml:"SupportResources,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource) GetMasterZoneId() *string {
	return s.MasterZoneId
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource) GetSlaveZoneId() *string {
	return s.SlaveZoneId
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource) GetSupportResources() *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResources {
	return s.SupportResources
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource) SetMasterZoneId(v string) *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource {
	s.MasterZoneId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource) SetSlaveZoneId(v string) *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource {
	s.SlaveZoneId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource) SetSupportResources(v *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResources) *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource {
	s.SupportResources = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource) Validate() error {
	if s.SupportResources != nil {
		if err := s.SupportResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResources struct {
	SupportResource []*DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource `json:"SupportResource,omitempty" xml:"SupportResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResources) GetSupportResource() []*DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource {
	return s.SupportResource
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResources) SetSupportResource(v []*DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource) *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResources {
	s.SupportResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResources) Validate() error {
	if s.SupportResource != nil {
		for _, item := range s.SupportResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource struct {
	// The type of the IP address.
	//
	// Valid values: **ipv4 and ipv6**.
	//
	// example:
	//
	// ipv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	// The network type.
	//
	// Valid values: **vpc, classic-internet, and classic-intranet**.
	//
	// example:
	//
	// classic_internet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource) GetAddressType() *string {
	return s.AddressType
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource) SetAddressIPVersion(v string) *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource) SetAddressType(v string) *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource {
	s.AddressType = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource) Validate() error {
	return dara.Validate(s)
}
