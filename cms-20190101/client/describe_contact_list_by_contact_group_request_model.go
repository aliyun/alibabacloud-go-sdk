// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContactListByContactGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupName(v string) *DescribeContactListByContactGroupRequest
	GetContactGroupName() *string
	SetRegionId(v string) *DescribeContactListByContactGroupRequest
	GetRegionId() *string
}

type DescribeContactListByContactGroupRequest struct {
	// The name of the alert contact group.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudMonitor
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeContactListByContactGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactListByContactGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeContactListByContactGroupRequest) GetContactGroupName() *string {
	return s.ContactGroupName
}

func (s *DescribeContactListByContactGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeContactListByContactGroupRequest) SetContactGroupName(v string) *DescribeContactListByContactGroupRequest {
	s.ContactGroupName = &v
	return s
}

func (s *DescribeContactListByContactGroupRequest) SetRegionId(v string) *DescribeContactListByContactGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContactListByContactGroupRequest) Validate() error {
	return dara.Validate(s)
}
