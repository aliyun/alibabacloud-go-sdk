// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDNADBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateDNADBRequest
	GetDescription() *string
	SetModel(v string) *CreateDNADBRequest
	GetModel() *string
	SetName(v string) *CreateDNADBRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateDNADBRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDNADBRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateDNADBRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDNADBRequest
	GetResourceOwnerId() *int64
}

type CreateDNADBRequest struct {
	// The description of the media fingerprint library.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The model of the media fingerprint library. Valid values:
	//
	// 	- **Video**
	//
	// 	- **Audio**
	//
	// 	- **Image**
	//
	// 	- **Text*	- (supported only in the China (Shanghai) region)
	//
	// example:
	//
	// Video
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The name of the media fingerprint library.
	//
	// This parameter is required.
	//
	// example:
	//
	// example name
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateDNADBRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDNADBRequest) GoString() string {
	return s.String()
}

func (s *CreateDNADBRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDNADBRequest) GetModel() *string {
	return s.Model
}

func (s *CreateDNADBRequest) GetName() *string {
	return s.Name
}

func (s *CreateDNADBRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDNADBRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDNADBRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDNADBRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDNADBRequest) SetDescription(v string) *CreateDNADBRequest {
	s.Description = &v
	return s
}

func (s *CreateDNADBRequest) SetModel(v string) *CreateDNADBRequest {
	s.Model = &v
	return s
}

func (s *CreateDNADBRequest) SetName(v string) *CreateDNADBRequest {
	s.Name = &v
	return s
}

func (s *CreateDNADBRequest) SetOwnerAccount(v string) *CreateDNADBRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDNADBRequest) SetOwnerId(v int64) *CreateDNADBRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDNADBRequest) SetResourceOwnerAccount(v string) *CreateDNADBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDNADBRequest) SetResourceOwnerId(v int64) *CreateDNADBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDNADBRequest) Validate() error {
	return dara.Validate(s)
}
