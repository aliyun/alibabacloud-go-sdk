// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAclRelationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclIds(v []*string) *ListAclRelationsRequest
	GetAclIds() []*string
}

type ListAclRelationsRequest struct {
	// The access control list (ACL) IDs. You can query at most five ACLs in each call.
	//
	// This parameter is required.
	AclIds []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
}

func (s ListAclRelationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAclRelationsRequest) GoString() string {
	return s.String()
}

func (s *ListAclRelationsRequest) GetAclIds() []*string {
	return s.AclIds
}

func (s *ListAclRelationsRequest) SetAclIds(v []*string) *ListAclRelationsRequest {
	s.AclIds = v
	return s
}

func (s *ListAclRelationsRequest) Validate() error {
	return dara.Validate(s)
}
