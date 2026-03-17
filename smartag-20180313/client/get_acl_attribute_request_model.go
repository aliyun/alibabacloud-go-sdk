// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAclAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *GetAclAttributeRequest
	GetAclId() *string
	SetRegionId(v string) *GetAclAttributeRequest
	GetRegionId() *string
}

type GetAclAttributeRequest struct {
	// The ID of the ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// acl-xhwhyuo43l0n*****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The ID of the region where the ACL is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAclAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAclAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetAclAttributeRequest) GetAclId() *string {
	return s.AclId
}

func (s *GetAclAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAclAttributeRequest) SetAclId(v string) *GetAclAttributeRequest {
	s.AclId = &v
	return s
}

func (s *GetAclAttributeRequest) SetRegionId(v string) *GetAclAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *GetAclAttributeRequest) Validate() error {
	return dara.Validate(s)
}
