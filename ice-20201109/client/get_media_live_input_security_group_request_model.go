// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaLiveInputSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityGroupId(v string) *GetMediaLiveInputSecurityGroupRequest
	GetSecurityGroupId() *string
}

type GetMediaLiveInputSecurityGroupRequest struct {
	// The ID of the security group.
	//
	// This parameter is required.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s GetMediaLiveInputSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveInputSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *GetMediaLiveInputSecurityGroupRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetMediaLiveInputSecurityGroupRequest) SetSecurityGroupId(v string) *GetMediaLiveInputSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *GetMediaLiveInputSecurityGroupRequest) Validate() error {
	return dara.Validate(s)
}
