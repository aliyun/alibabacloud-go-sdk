// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAclAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *UpdateAclAttributeRequest
	GetAclId() *string
	SetAclName(v string) *UpdateAclAttributeRequest
	GetAclName() *string
	SetClientToken(v string) *UpdateAclAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateAclAttributeRequest
	GetDryRun() *bool
}

type UpdateAclAttributeRequest struct {
	// The ACL ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// nacl-hp34s2h0xx1ht4nwo****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The ACL name. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-acl
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a `2xx` HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s UpdateAclAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAclAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateAclAttributeRequest) GetAclId() *string {
	return s.AclId
}

func (s *UpdateAclAttributeRequest) GetAclName() *string {
	return s.AclName
}

func (s *UpdateAclAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAclAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateAclAttributeRequest) SetAclId(v string) *UpdateAclAttributeRequest {
	s.AclId = &v
	return s
}

func (s *UpdateAclAttributeRequest) SetAclName(v string) *UpdateAclAttributeRequest {
	s.AclName = &v
	return s
}

func (s *UpdateAclAttributeRequest) SetClientToken(v string) *UpdateAclAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAclAttributeRequest) SetDryRun(v bool) *UpdateAclAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateAclAttributeRequest) Validate() error {
	return dara.Validate(s)
}
