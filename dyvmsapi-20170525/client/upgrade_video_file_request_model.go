// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeVideoFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *UpgradeVideoFileRequest
	GetCallId() *string
	SetCalledNumber(v string) *UpgradeVideoFileRequest
	GetCalledNumber() *string
	SetMediaType(v string) *UpgradeVideoFileRequest
	GetMediaType() *string
	SetOutId(v string) *UpgradeVideoFileRequest
	GetOutId() *string
	SetOwnerId(v int64) *UpgradeVideoFileRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpgradeVideoFileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpgradeVideoFileRequest
	GetResourceOwnerId() *int64
}

type UpgradeVideoFileRequest struct {
	// example:
	//
	// 116012354148^10281378****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1590****000
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// 225869*****
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpgradeVideoFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeVideoFileRequest) GoString() string {
	return s.String()
}

func (s *UpgradeVideoFileRequest) GetCallId() *string {
	return s.CallId
}

func (s *UpgradeVideoFileRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *UpgradeVideoFileRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *UpgradeVideoFileRequest) GetOutId() *string {
	return s.OutId
}

func (s *UpgradeVideoFileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeVideoFileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpgradeVideoFileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeVideoFileRequest) SetCallId(v string) *UpgradeVideoFileRequest {
	s.CallId = &v
	return s
}

func (s *UpgradeVideoFileRequest) SetCalledNumber(v string) *UpgradeVideoFileRequest {
	s.CalledNumber = &v
	return s
}

func (s *UpgradeVideoFileRequest) SetMediaType(v string) *UpgradeVideoFileRequest {
	s.MediaType = &v
	return s
}

func (s *UpgradeVideoFileRequest) SetOutId(v string) *UpgradeVideoFileRequest {
	s.OutId = &v
	return s
}

func (s *UpgradeVideoFileRequest) SetOwnerId(v int64) *UpgradeVideoFileRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeVideoFileRequest) SetResourceOwnerAccount(v string) *UpgradeVideoFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeVideoFileRequest) SetResourceOwnerId(v int64) *UpgradeVideoFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeVideoFileRequest) Validate() error {
	return dara.Validate(s)
}
