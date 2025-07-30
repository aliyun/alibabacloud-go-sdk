// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *ResetPasswordRequest
	GetClientId() *string
	SetClientToken(v string) *ResetPasswordRequest
	GetClientToken() *string
	SetEmail(v string) *ResetPasswordRequest
	GetEmail() *string
	SetEndUserId(v string) *ResetPasswordRequest
	GetEndUserId() *string
	SetOfficeSiteId(v string) *ResetPasswordRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *ResetPasswordRequest
	GetRegionId() *string
	SetPhone(v string) *ResetPasswordRequest
	GetPhone() *string
}

type ResetPasswordRequest struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95e41934-383e-4c9f-824f-3b93b19b****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 2f00ab32-a473-4c90-9aae-dd8842ae****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// a***@example.edu
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// liming
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The office network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-899235****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The phone number of the user.
	//
	// example:
	//
	// 1827912****
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s ResetPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetPasswordRequest) GetClientId() *string {
	return s.ClientId
}

func (s *ResetPasswordRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ResetPasswordRequest) GetEmail() *string {
	return s.Email
}

func (s *ResetPasswordRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ResetPasswordRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ResetPasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetPasswordRequest) GetPhone() *string {
	return s.Phone
}

func (s *ResetPasswordRequest) SetClientId(v string) *ResetPasswordRequest {
	s.ClientId = &v
	return s
}

func (s *ResetPasswordRequest) SetClientToken(v string) *ResetPasswordRequest {
	s.ClientToken = &v
	return s
}

func (s *ResetPasswordRequest) SetEmail(v string) *ResetPasswordRequest {
	s.Email = &v
	return s
}

func (s *ResetPasswordRequest) SetEndUserId(v string) *ResetPasswordRequest {
	s.EndUserId = &v
	return s
}

func (s *ResetPasswordRequest) SetOfficeSiteId(v string) *ResetPasswordRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ResetPasswordRequest) SetRegionId(v string) *ResetPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ResetPasswordRequest) SetPhone(v string) *ResetPasswordRequest {
	s.Phone = &v
	return s
}

func (s *ResetPasswordRequest) Validate() error {
	return dara.Validate(s)
}
