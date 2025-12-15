// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSecurityIpsResponseBody
	GetRequestId() *string
	SetSecurityIpGroups(v *DescribeSecurityIpsResponseBodySecurityIpGroups) *DescribeSecurityIpsResponseBody
	GetSecurityIpGroups() *DescribeSecurityIpsResponseBodySecurityIpGroups
}

type DescribeSecurityIpsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EFC9161F-15E3-4A6E-8A99-C09916D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The whitelists of the instance.
	SecurityIpGroups *DescribeSecurityIpsResponseBodySecurityIpGroups `json:"SecurityIpGroups,omitempty" xml:"SecurityIpGroups,omitempty" type:"Struct"`
}

func (s DescribeSecurityIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityIpsResponseBody) GetSecurityIpGroups() *DescribeSecurityIpsResponseBodySecurityIpGroups {
	return s.SecurityIpGroups
}

func (s *DescribeSecurityIpsResponseBody) SetRequestId(v string) *DescribeSecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityIpsResponseBody) SetSecurityIpGroups(v *DescribeSecurityIpsResponseBodySecurityIpGroups) *DescribeSecurityIpsResponseBody {
	s.SecurityIpGroups = v
	return s
}

func (s *DescribeSecurityIpsResponseBody) Validate() error {
	if s.SecurityIpGroups != nil {
		if err := s.SecurityIpGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityIpsResponseBodySecurityIpGroups struct {
	SecurityIpGroup []*DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup `json:"SecurityIpGroup,omitempty" xml:"SecurityIpGroup,omitempty" type:"Repeated"`
}

func (s DescribeSecurityIpsResponseBodySecurityIpGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIpsResponseBodySecurityIpGroups) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsResponseBodySecurityIpGroups) GetSecurityIpGroup() []*DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup {
	return s.SecurityIpGroup
}

func (s *DescribeSecurityIpsResponseBodySecurityIpGroups) SetSecurityIpGroup(v []*DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup) *DescribeSecurityIpsResponseBodySecurityIpGroups {
	s.SecurityIpGroup = v
	return s
}

func (s *DescribeSecurityIpsResponseBodySecurityIpGroups) Validate() error {
	if s.SecurityIpGroup != nil {
		for _, item := range s.SecurityIpGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup struct {
	// The attribute of the whitelist. This parameter is empty by default.
	//
	// >  If the instance is authorized to use a service such as Database Autonomy Service (DAS), Data Management (DMS), or Data Transmission Service (DTS), this service automatically generates a **hidden*	- whitelist for the instance. This type of whitelists cannot be modified or deleted.
	//
	// example:
	//
	// hidden
	SecurityIpGroupAttribute *string `json:"SecurityIpGroupAttribute,omitempty" xml:"SecurityIpGroupAttribute,omitempty"`
	// The name of the whitelist.
	//
	// example:
	//
	// default
	SecurityIpGroupName *string `json:"SecurityIpGroupName,omitempty" xml:"SecurityIpGroupName,omitempty"`
	// The IP addresses in the whitelist. A maximum of 1,000 IP addresses can be specified in a whitelist.
	SecurityIpList *string `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty"`
}

func (s DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup) GetSecurityIpGroupAttribute() *string {
	return s.SecurityIpGroupAttribute
}

func (s *DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup) GetSecurityIpGroupName() *string {
	return s.SecurityIpGroupName
}

func (s *DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup) GetSecurityIpList() *string {
	return s.SecurityIpList
}

func (s *DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup) SetSecurityIpGroupAttribute(v string) *DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup {
	s.SecurityIpGroupAttribute = &v
	return s
}

func (s *DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup) SetSecurityIpGroupName(v string) *DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup {
	s.SecurityIpGroupName = &v
	return s
}

func (s *DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup) SetSecurityIpList(v string) *DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup {
	s.SecurityIpList = &v
	return s
}

func (s *DescribeSecurityIpsResponseBodySecurityIpGroupsSecurityIpGroup) Validate() error {
	return dara.Validate(s)
}
