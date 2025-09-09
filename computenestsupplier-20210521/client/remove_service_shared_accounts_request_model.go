// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveServiceSharedAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RemoveServiceSharedAccountsRequest
	GetClientToken() *string
	SetRegionId(v string) *RemoveServiceSharedAccountsRequest
	GetRegionId() *string
	SetServiceId(v string) *RemoveServiceSharedAccountsRequest
	GetServiceId() *string
	SetType(v string) *RemoveServiceSharedAccountsRequest
	GetType() *string
	SetUserAliUids(v []*int64) *RemoveServiceSharedAccountsRequest
	GetUserAliUids() []*int64
}

type RemoveServiceSharedAccountsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The share type of the service. Default value: SharedAccount. Valid values:
	//
	// 	- SharedAccount: The service is shared by multiple accounts.
	//
	// 	- Reseller: The service is distributed.
	//
	// example:
	//
	// SharedAccount
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Whitelist accounts for service sharing.
	//
	// This parameter is required.
	UserAliUids []*int64 `json:"UserAliUids,omitempty" xml:"UserAliUids,omitempty" type:"Repeated"`
}

func (s RemoveServiceSharedAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveServiceSharedAccountsRequest) GoString() string {
	return s.String()
}

func (s *RemoveServiceSharedAccountsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveServiceSharedAccountsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveServiceSharedAccountsRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *RemoveServiceSharedAccountsRequest) GetType() *string {
	return s.Type
}

func (s *RemoveServiceSharedAccountsRequest) GetUserAliUids() []*int64 {
	return s.UserAliUids
}

func (s *RemoveServiceSharedAccountsRequest) SetClientToken(v string) *RemoveServiceSharedAccountsRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveServiceSharedAccountsRequest) SetRegionId(v string) *RemoveServiceSharedAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveServiceSharedAccountsRequest) SetServiceId(v string) *RemoveServiceSharedAccountsRequest {
	s.ServiceId = &v
	return s
}

func (s *RemoveServiceSharedAccountsRequest) SetType(v string) *RemoveServiceSharedAccountsRequest {
	s.Type = &v
	return s
}

func (s *RemoveServiceSharedAccountsRequest) SetUserAliUids(v []*int64) *RemoveServiceSharedAccountsRequest {
	s.UserAliUids = v
	return s
}

func (s *RemoveServiceSharedAccountsRequest) Validate() error {
	return dara.Validate(s)
}
