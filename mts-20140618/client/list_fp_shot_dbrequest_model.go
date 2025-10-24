// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFpShotDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFpDBIds(v string) *ListFpShotDBRequest
	GetFpDBIds() *string
	SetOwnerAccount(v string) *ListFpShotDBRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListFpShotDBRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListFpShotDBRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListFpShotDBRequest
	GetResourceOwnerId() *int64
}

type ListFpShotDBRequest struct {
	// The ID of the media fingerprint library. You can obtain the library ID from the response parameters of the [CreateFpShotDB](https://help.aliyun.com/document_detail/170149.html) operation. You can query up to 10 libraries at a time. Separate multiple library IDs with commas (,).
	//
	// example:
	//
	// 2288c6ca184c0e47098a5b665e2a12****,ae687c02fe944327ba9631e50da2****
	FpDBIds              *string `json:"FpDBIds,omitempty" xml:"FpDBIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListFpShotDBRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFpShotDBRequest) GoString() string {
	return s.String()
}

func (s *ListFpShotDBRequest) GetFpDBIds() *string {
	return s.FpDBIds
}

func (s *ListFpShotDBRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListFpShotDBRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListFpShotDBRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListFpShotDBRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListFpShotDBRequest) SetFpDBIds(v string) *ListFpShotDBRequest {
	s.FpDBIds = &v
	return s
}

func (s *ListFpShotDBRequest) SetOwnerAccount(v string) *ListFpShotDBRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListFpShotDBRequest) SetOwnerId(v int64) *ListFpShotDBRequest {
	s.OwnerId = &v
	return s
}

func (s *ListFpShotDBRequest) SetResourceOwnerAccount(v string) *ListFpShotDBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListFpShotDBRequest) SetResourceOwnerId(v int64) *ListFpShotDBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListFpShotDBRequest) Validate() error {
	return dara.Validate(s)
}
