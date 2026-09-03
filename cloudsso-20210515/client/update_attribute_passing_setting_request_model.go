// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAttributePassingSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *UpdateAttributePassingSettingRequest
	GetDirectoryId() *string
	SetSourceIdentityPassing(v string) *UpdateAttributePassingSettingRequest
	GetSourceIdentityPassing() *string
}

type UpdateAttributePassingSettingRequest struct {
	// The directory ID.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The SourceIdentity pass-through mode. Three pass-through modes are supported. If this parameter is not specified, the existing configuration is not modified.
	//
	// Valid values:
	//
	// - IdP: Uses the SourceIdentity attribute value carried in the SAML assertion from the external identity provider (IdP). The attribute value must be 2 to 64 characters in length and can contain only letters, digits, and the following special characters: =,.@-_.
	//
	// - UserName: Uses the CloudSSO username as the SourceIdentity. The system automatically adds a reserved prefix, resulting in the format acs:sso:<username>, with a total length of no more than 64 characters.
	//
	// - Disabled: Does not pass through SourceIdentity.
	//
	// example:
	//
	// Disabled
	SourceIdentityPassing *string `json:"SourceIdentityPassing,omitempty" xml:"SourceIdentityPassing,omitempty"`
}

func (s UpdateAttributePassingSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAttributePassingSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateAttributePassingSettingRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateAttributePassingSettingRequest) GetSourceIdentityPassing() *string {
	return s.SourceIdentityPassing
}

func (s *UpdateAttributePassingSettingRequest) SetDirectoryId(v string) *UpdateAttributePassingSettingRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateAttributePassingSettingRequest) SetSourceIdentityPassing(v string) *UpdateAttributePassingSettingRequest {
	s.SourceIdentityPassing = &v
	return s
}

func (s *UpdateAttributePassingSettingRequest) Validate() error {
	return dara.Validate(s)
}
