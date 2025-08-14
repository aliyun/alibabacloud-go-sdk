// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaLiveInputSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityGroupId(v string) *DeleteMediaLiveInputSecurityGroupRequest
	GetSecurityGroupId() *string
}

type DeleteMediaLiveInputSecurityGroupRequest struct {
	// The ID of the security group.
	//
	// This parameter is required.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s DeleteMediaLiveInputSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaLiveInputSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteMediaLiveInputSecurityGroupRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DeleteMediaLiveInputSecurityGroupRequest) SetSecurityGroupId(v string) *DeleteMediaLiveInputSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DeleteMediaLiveInputSecurityGroupRequest) Validate() error {
	return dara.Validate(s)
}
