// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttributePassingSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttributePassingSetting(v *GetAttributePassingSettingResponseBodyAttributePassingSetting) *GetAttributePassingSettingResponseBody
	GetAttributePassingSetting() *GetAttributePassingSettingResponseBodyAttributePassingSetting
	SetRequestId(v string) *GetAttributePassingSettingResponseBody
	GetRequestId() *string
}

type GetAttributePassingSettingResponseBody struct {
	// The attribute passing settings.
	AttributePassingSetting *GetAttributePassingSettingResponseBodyAttributePassingSetting `json:"AttributePassingSetting,omitempty" xml:"AttributePassingSetting,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F6F90F3D-4502-5877-B80B-97476F6AE2CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAttributePassingSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAttributePassingSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetAttributePassingSettingResponseBody) GetAttributePassingSetting() *GetAttributePassingSettingResponseBodyAttributePassingSetting {
	return s.AttributePassingSetting
}

func (s *GetAttributePassingSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAttributePassingSettingResponseBody) SetAttributePassingSetting(v *GetAttributePassingSettingResponseBodyAttributePassingSetting) *GetAttributePassingSettingResponseBody {
	s.AttributePassingSetting = v
	return s
}

func (s *GetAttributePassingSettingResponseBody) SetRequestId(v string) *GetAttributePassingSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAttributePassingSettingResponseBody) Validate() error {
	if s.AttributePassingSetting != nil {
		if err := s.AttributePassingSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAttributePassingSettingResponseBodyAttributePassingSetting struct {
	// The SourceIdentity pass-through mode. If not configured, Disabled is returned by default. Valid values:
	//
	// - IdP: Uses the SourceIdentity attribute value carried in the SAML assertion from the external identity provider (IdP). The attribute value must be 2 to 64 characters in length and can contain only letters, digits, and the following special characters: =,.@-_.
	//
	// - UserName: Uses the CloudSSO username as the SourceIdentity. The system automatically adds a reserved prefix, resulting in the format acs:sso:<username>. The total length cannot exceed 64 characters.
	//
	// - Disabled: Does not pass through the SourceIdentity.
	//
	// example:
	//
	// Disabled
	SourceIdentityPassing *string `json:"SourceIdentityPassing,omitempty" xml:"SourceIdentityPassing,omitempty"`
}

func (s GetAttributePassingSettingResponseBodyAttributePassingSetting) String() string {
	return dara.Prettify(s)
}

func (s GetAttributePassingSettingResponseBodyAttributePassingSetting) GoString() string {
	return s.String()
}

func (s *GetAttributePassingSettingResponseBodyAttributePassingSetting) GetSourceIdentityPassing() *string {
	return s.SourceIdentityPassing
}

func (s *GetAttributePassingSettingResponseBodyAttributePassingSetting) SetSourceIdentityPassing(v string) *GetAttributePassingSettingResponseBodyAttributePassingSetting {
	s.SourceIdentityPassing = &v
	return s
}

func (s *GetAttributePassingSettingResponseBodyAttributePassingSetting) Validate() error {
	return dara.Validate(s)
}
