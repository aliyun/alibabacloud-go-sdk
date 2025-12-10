// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *CreateCloudAccountRequest
	GetDisplayName() *string
	SetEmail(v string) *CreateCloudAccountRequest
	GetEmail() *string
	SetParentFolderId(v string) *CreateCloudAccountRequest
	GetParentFolderId() *string
	SetPayerAccountId(v string) *CreateCloudAccountRequest
	GetPayerAccountId() *string
}

type CreateCloudAccountRequest struct {
	// The display name of the member account.
	//
	// The name must be 2 to 50 characters in length and can contain letters, digits, underscores (_), periods (.), and hyphens (-).
	//
	// The name must be unique in the current resource directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// admin-****
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address used to log on to the cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// someone@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The ID of the parent folder.
	//
	// example:
	//
	// fd-bVaRIG****
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// The ID of the settlement account. If you do not specify this parameter, the current account is used for settlement.
	//
	// example:
	//
	// 12323344****
	PayerAccountId *string `json:"PayerAccountId,omitempty" xml:"PayerAccountId,omitempty"`
}

func (s CreateCloudAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudAccountRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateCloudAccountRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateCloudAccountRequest) GetParentFolderId() *string {
	return s.ParentFolderId
}

func (s *CreateCloudAccountRequest) GetPayerAccountId() *string {
	return s.PayerAccountId
}

func (s *CreateCloudAccountRequest) SetDisplayName(v string) *CreateCloudAccountRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateCloudAccountRequest) SetEmail(v string) *CreateCloudAccountRequest {
	s.Email = &v
	return s
}

func (s *CreateCloudAccountRequest) SetParentFolderId(v string) *CreateCloudAccountRequest {
	s.ParentFolderId = &v
	return s
}

func (s *CreateCloudAccountRequest) SetPayerAccountId(v string) *CreateCloudAccountRequest {
	s.PayerAccountId = &v
	return s
}

func (s *CreateCloudAccountRequest) Validate() error {
	return dara.Validate(s)
}
