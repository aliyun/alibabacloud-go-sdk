// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v string) *CreateAccessTokenRequest
	GetCount() *string
	SetDescription(v string) *CreateAccessTokenRequest
	GetDescription() *string
	SetName(v string) *CreateAccessTokenRequest
	GetName() *string
	SetOwnerId(v int64) *CreateAccessTokenRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateAccessTokenRequest
	GetResourceOwnerAccount() *string
	SetTimeToLiveInDays(v string) *CreateAccessTokenRequest
	GetTimeToLiveInDays() *string
}

type CreateAccessTokenRequest struct {
	// The maximum number of times that the activation code can be used to import the information of migration sources. Valid values: 1 to 1000.
	//
	// Default value: 100.
	//
	// example:
	//
	// 10
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The description of the activation code.
	//
	// example:
	//
	// The description of the activation code.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the activation code. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// test_name
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The validity period of the activation code. The activation code can no longer be used to import the information of migration sources after the code expires. Unit: day. Valid values: 1 to 90.
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	TimeToLiveInDays *string `json:"TimeToLiveInDays,omitempty" xml:"TimeToLiveInDays,omitempty"`
}

func (s CreateAccessTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessTokenRequest) GetCount() *string {
	return s.Count
}

func (s *CreateAccessTokenRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAccessTokenRequest) GetName() *string {
	return s.Name
}

func (s *CreateAccessTokenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAccessTokenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAccessTokenRequest) GetTimeToLiveInDays() *string {
	return s.TimeToLiveInDays
}

func (s *CreateAccessTokenRequest) SetCount(v string) *CreateAccessTokenRequest {
	s.Count = &v
	return s
}

func (s *CreateAccessTokenRequest) SetDescription(v string) *CreateAccessTokenRequest {
	s.Description = &v
	return s
}

func (s *CreateAccessTokenRequest) SetName(v string) *CreateAccessTokenRequest {
	s.Name = &v
	return s
}

func (s *CreateAccessTokenRequest) SetOwnerId(v int64) *CreateAccessTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAccessTokenRequest) SetResourceOwnerAccount(v string) *CreateAccessTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAccessTokenRequest) SetTimeToLiveInDays(v string) *CreateAccessTokenRequest {
	s.TimeToLiveInDays = &v
	return s
}

func (s *CreateAccessTokenRequest) Validate() error {
	return dara.Validate(s)
}
