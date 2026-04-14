// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *CreateDomainRequest
	GetDomainName() *string
	SetOwnerId(v int64) *CreateDomainRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateDomainRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDomainRequest
	GetResourceOwnerId() *int64
	SetDkimSelector(v string) *CreateDomainRequest
	GetDkimSelector() *string
}

type CreateDomainRequest struct {
	// Domain name, length 1-50, can include numbers, uppercase letters, lowercase letters, ., and -.
	//
	// This parameter is required.
	//
	// example:
	//
	// sub.example.com
	DomainName           *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The selector field in the DKIM protocol, used to identify a specific public key. It is recommended to leave it blank, as the system will automatically generate it based on cluster information. If the user specifies it manually, for example, if the sending domain is "sub.example.com" and dkimSelector is set to "default", then the host record will be "default._domainkey.sub"
	//
	// Constraints:
	//
	// 1. The length must not exceed 60 characters.
	//
	// 2. It must consist of visible characters only.
	//
	// 3. It cannot start with a hyphen (-).
	//
	// 4. It cannot end with a hyphen (-).
	//
	// 5. It cannot contain any of the following characters: _ :;/!*~.@#$%^&()+=[{]}|?<>,\\""
	//
	// example:
	//
	// default
	DkimSelector *string `json:"dkimSelector,omitempty" xml:"dkimSelector,omitempty"`
}

func (s CreateDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDomainRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDomainRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDomainRequest) GetDkimSelector() *string {
	return s.DkimSelector
}

func (s *CreateDomainRequest) SetDomainName(v string) *CreateDomainRequest {
	s.DomainName = &v
	return s
}

func (s *CreateDomainRequest) SetOwnerId(v int64) *CreateDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDomainRequest) SetResourceOwnerAccount(v string) *CreateDomainRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDomainRequest) SetResourceOwnerId(v int64) *CreateDomainRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDomainRequest) SetDkimSelector(v string) *CreateDomainRequest {
	s.DkimSelector = &v
	return s
}

func (s *CreateDomainRequest) Validate() error {
	return dara.Validate(s)
}
