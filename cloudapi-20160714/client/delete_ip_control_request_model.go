// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpControlId(v string) *DeleteIpControlRequest
	GetIpControlId() *string
	SetSecurityToken(v string) *DeleteIpControlRequest
	GetSecurityToken() *string
}

type DeleteIpControlRequest struct {
	// The ID of the ACL. The ID is unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7ea91319a34d48a09b5c9c871d9768b1
	IpControlId   *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteIpControlRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpControlRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpControlRequest) GetIpControlId() *string {
	return s.IpControlId
}

func (s *DeleteIpControlRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteIpControlRequest) SetIpControlId(v string) *DeleteIpControlRequest {
	s.IpControlId = &v
	return s
}

func (s *DeleteIpControlRequest) SetSecurityToken(v string) *DeleteIpControlRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteIpControlRequest) Validate() error {
	return dara.Validate(s)
}
