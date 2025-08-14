// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecognitionLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *DeleteRecognitionLibRequest
	GetAlgorithm() *string
	SetLibId(v string) *DeleteRecognitionLibRequest
	GetLibId() *string
	SetOwnerAccount(v string) *DeleteRecognitionLibRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteRecognitionLibRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteRecognitionLibRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteRecognitionLibRequest
	GetResourceOwnerId() *int64
}

type DeleteRecognitionLibRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// landmark
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// *************24b47865c6**************
	LibId                *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteRecognitionLibRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecognitionLibRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecognitionLibRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *DeleteRecognitionLibRequest) GetLibId() *string {
	return s.LibId
}

func (s *DeleteRecognitionLibRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteRecognitionLibRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteRecognitionLibRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteRecognitionLibRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteRecognitionLibRequest) SetAlgorithm(v string) *DeleteRecognitionLibRequest {
	s.Algorithm = &v
	return s
}

func (s *DeleteRecognitionLibRequest) SetLibId(v string) *DeleteRecognitionLibRequest {
	s.LibId = &v
	return s
}

func (s *DeleteRecognitionLibRequest) SetOwnerAccount(v string) *DeleteRecognitionLibRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteRecognitionLibRequest) SetOwnerId(v int64) *DeleteRecognitionLibRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRecognitionLibRequest) SetResourceOwnerAccount(v string) *DeleteRecognitionLibRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteRecognitionLibRequest) SetResourceOwnerId(v int64) *DeleteRecognitionLibRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteRecognitionLibRequest) Validate() error {
	return dara.Validate(s)
}
