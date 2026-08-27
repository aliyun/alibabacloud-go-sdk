// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalFeishuMinuteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialId(v string) *CreatePersonalFeishuMinuteRequest
	GetCredentialId() *string
	SetDescription(v string) *CreatePersonalFeishuMinuteRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalFeishuMinuteRequest
	GetDirectoryId() *string
	SetMinuteToken(v string) *CreatePersonalFeishuMinuteRequest
	GetMinuteToken() *string
	SetName(v string) *CreatePersonalFeishuMinuteRequest
	GetName() *string
	SetOperatingObjectName(v string) *CreatePersonalFeishuMinuteRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *CreatePersonalFeishuMinuteRequest
	GetTenantId() *string
}

type CreatePersonalFeishuMinuteRequest struct {
	// The credential ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleCredentialId
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// The resource description.
	//
	// example:
	//
	// created by eventbridge
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The Lark Minutes token (unique identifier of the minutes record, required).
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	MinuteToken *string `json:"minuteToken,omitempty" xml:"minuteToken,omitempty"`
	// The resource name.
	//
	// This parameter is required.
	//
	// example:
	//
	// _DevsAF_19df1a74-a740-449a-bd7a-9acb39e00f25
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the digital employee (operating object name, optional).
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1iSgnW4pARwoPUd5D5nuCNwiEiE
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalFeishuMinuteRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalFeishuMinuteRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalFeishuMinuteRequest) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CreatePersonalFeishuMinuteRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalFeishuMinuteRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalFeishuMinuteRequest) GetMinuteToken() *string {
	return s.MinuteToken
}

func (s *CreatePersonalFeishuMinuteRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalFeishuMinuteRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalFeishuMinuteRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalFeishuMinuteRequest) SetCredentialId(v string) *CreatePersonalFeishuMinuteRequest {
	s.CredentialId = &v
	return s
}

func (s *CreatePersonalFeishuMinuteRequest) SetDescription(v string) *CreatePersonalFeishuMinuteRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalFeishuMinuteRequest) SetDirectoryId(v string) *CreatePersonalFeishuMinuteRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalFeishuMinuteRequest) SetMinuteToken(v string) *CreatePersonalFeishuMinuteRequest {
	s.MinuteToken = &v
	return s
}

func (s *CreatePersonalFeishuMinuteRequest) SetName(v string) *CreatePersonalFeishuMinuteRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalFeishuMinuteRequest) SetOperatingObjectName(v string) *CreatePersonalFeishuMinuteRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalFeishuMinuteRequest) SetTenantId(v string) *CreatePersonalFeishuMinuteRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalFeishuMinuteRequest) Validate() error {
	return dara.Validate(s)
}
