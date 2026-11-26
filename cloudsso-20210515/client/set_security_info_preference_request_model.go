// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSecurityInfoPreferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowUserManageSecurityEmail(v bool) *SetSecurityInfoPreferenceRequest
	GetAllowUserManageSecurityEmail() *bool
	SetDirectoryId(v string) *SetSecurityInfoPreferenceRequest
	GetDirectoryId() *string
}

type SetSecurityInfoPreferenceRequest struct {
	// Specifies whether to allow CloudSSO users in the directory to manage their security email addresses in the CloudSSO user portal, including setting, modifying, and deleting security email addresses, and sending verification emails. Valid values:
	//
	// - true: allowed.
	//
	// - false: not allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	AllowUserManageSecurityEmail *bool `json:"AllowUserManageSecurityEmail,omitempty" xml:"AllowUserManageSecurityEmail,omitempty"`
	// The directory ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s SetSecurityInfoPreferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityInfoPreferenceRequest) GoString() string {
	return s.String()
}

func (s *SetSecurityInfoPreferenceRequest) GetAllowUserManageSecurityEmail() *bool {
	return s.AllowUserManageSecurityEmail
}

func (s *SetSecurityInfoPreferenceRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SetSecurityInfoPreferenceRequest) SetAllowUserManageSecurityEmail(v bool) *SetSecurityInfoPreferenceRequest {
	s.AllowUserManageSecurityEmail = &v
	return s
}

func (s *SetSecurityInfoPreferenceRequest) SetDirectoryId(v string) *SetSecurityInfoPreferenceRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetSecurityInfoPreferenceRequest) Validate() error {
	return dara.Validate(s)
}
