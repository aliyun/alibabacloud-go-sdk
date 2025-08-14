// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaLiveInputSecurityGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMediaLiveInputSecurityGroupResponseBody
	GetRequestId() *string
	SetSecurityGroupId(v string) *CreateMediaLiveInputSecurityGroupResponseBody
	GetSecurityGroupId() *string
}

type CreateMediaLiveInputSecurityGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s CreateMediaLiveInputSecurityGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveInputSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveInputSecurityGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMediaLiveInputSecurityGroupResponseBody) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateMediaLiveInputSecurityGroupResponseBody) SetRequestId(v string) *CreateMediaLiveInputSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMediaLiveInputSecurityGroupResponseBody) SetSecurityGroupId(v string) *CreateMediaLiveInputSecurityGroupResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateMediaLiveInputSecurityGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
