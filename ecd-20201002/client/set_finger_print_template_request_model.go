// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetFingerPrintTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *SetFingerPrintTemplateRequest
	GetClientId() *string
	SetClientToken(v string) *SetFingerPrintTemplateRequest
	GetClientToken() *string
	SetDescription(v string) *SetFingerPrintTemplateRequest
	GetDescription() *string
	SetEncryptedFingerPrintTemplate(v string) *SetFingerPrintTemplateRequest
	GetEncryptedFingerPrintTemplate() *string
	SetEncryptedKey(v string) *SetFingerPrintTemplateRequest
	GetEncryptedKey() *string
	SetFingerPrintTemplate(v string) *SetFingerPrintTemplateRequest
	GetFingerPrintTemplate() *string
	SetLoginToken(v string) *SetFingerPrintTemplateRequest
	GetLoginToken() *string
	SetPassword(v string) *SetFingerPrintTemplateRequest
	GetPassword() *string
	SetRegionId(v string) *SetFingerPrintTemplateRequest
	GetRegionId() *string
	SetSessionId(v string) *SetFingerPrintTemplateRequest
	GetSessionId() *string
}

type SetFingerPrintTemplateRequest struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// 347431a9-90f6-448e-82c4-42bc84a9****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client token to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the node.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The encrypted fingerprint template.
	//
	// example:
	//
	// AAAAAAAAAAAAAA
	EncryptedFingerPrintTemplate *string `json:"EncryptedFingerPrintTemplate,omitempty" xml:"EncryptedFingerPrintTemplate,omitempty"`
	// The encryption key.
	//
	// example:
	//
	// drjfspchj
	EncryptedKey *string `json:"EncryptedKey,omitempty" xml:"EncryptedKey,omitempty"`
	// The fingerprint template.
	//
	// example:
	//
	// goG3gG8AAABhujtscn
	FingerPrintTemplate *string `json:"FingerPrintTemplate,omitempty" xml:"FingerPrintTemplate,omitempty"`
	// The logon credentials.
	//
	// This parameter is required.
	//
	// example:
	//
	// v11c73e7af0cb43ff39301651142485099ffb447085d76c4147519dbaa21c3bd90d53045e327c1f525ee6331c52556****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The password that you want to encrypt.
	//
	// This parameter is required.
	//
	// example:
	//
	// As53328794
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8b42538a-246e-45a1-95ea-e5c65b09****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s SetFingerPrintTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s SetFingerPrintTemplateRequest) GoString() string {
	return s.String()
}

func (s *SetFingerPrintTemplateRequest) GetClientId() *string {
	return s.ClientId
}

func (s *SetFingerPrintTemplateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SetFingerPrintTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *SetFingerPrintTemplateRequest) GetEncryptedFingerPrintTemplate() *string {
	return s.EncryptedFingerPrintTemplate
}

func (s *SetFingerPrintTemplateRequest) GetEncryptedKey() *string {
	return s.EncryptedKey
}

func (s *SetFingerPrintTemplateRequest) GetFingerPrintTemplate() *string {
	return s.FingerPrintTemplate
}

func (s *SetFingerPrintTemplateRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *SetFingerPrintTemplateRequest) GetPassword() *string {
	return s.Password
}

func (s *SetFingerPrintTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetFingerPrintTemplateRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *SetFingerPrintTemplateRequest) SetClientId(v string) *SetFingerPrintTemplateRequest {
	s.ClientId = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetClientToken(v string) *SetFingerPrintTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetDescription(v string) *SetFingerPrintTemplateRequest {
	s.Description = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetEncryptedFingerPrintTemplate(v string) *SetFingerPrintTemplateRequest {
	s.EncryptedFingerPrintTemplate = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetEncryptedKey(v string) *SetFingerPrintTemplateRequest {
	s.EncryptedKey = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetFingerPrintTemplate(v string) *SetFingerPrintTemplateRequest {
	s.FingerPrintTemplate = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetLoginToken(v string) *SetFingerPrintTemplateRequest {
	s.LoginToken = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetPassword(v string) *SetFingerPrintTemplateRequest {
	s.Password = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetRegionId(v string) *SetFingerPrintTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetSessionId(v string) *SetFingerPrintTemplateRequest {
	s.SessionId = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) Validate() error {
	return dara.Validate(s)
}
