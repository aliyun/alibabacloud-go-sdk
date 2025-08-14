// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecognitionSampleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *DeleteRecognitionSampleRequest
	GetAlgorithm() *string
	SetEntityId(v string) *DeleteRecognitionSampleRequest
	GetEntityId() *string
	SetLibId(v string) *DeleteRecognitionSampleRequest
	GetLibId() *string
	SetOwnerAccount(v string) *DeleteRecognitionSampleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteRecognitionSampleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteRecognitionSampleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteRecognitionSampleRequest
	GetResourceOwnerId() *int64
	SetSampleId(v string) *DeleteRecognitionSampleRequest
	GetSampleId() *string
}

type DeleteRecognitionSampleRequest struct {
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
	// **************544cb84754************
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// **************4d2ba728e2f**************
	SampleId *string `json:"SampleId,omitempty" xml:"SampleId,omitempty"`
}

func (s DeleteRecognitionSampleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecognitionSampleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecognitionSampleRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *DeleteRecognitionSampleRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *DeleteRecognitionSampleRequest) GetLibId() *string {
	return s.LibId
}

func (s *DeleteRecognitionSampleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteRecognitionSampleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteRecognitionSampleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteRecognitionSampleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteRecognitionSampleRequest) GetSampleId() *string {
	return s.SampleId
}

func (s *DeleteRecognitionSampleRequest) SetAlgorithm(v string) *DeleteRecognitionSampleRequest {
	s.Algorithm = &v
	return s
}

func (s *DeleteRecognitionSampleRequest) SetEntityId(v string) *DeleteRecognitionSampleRequest {
	s.EntityId = &v
	return s
}

func (s *DeleteRecognitionSampleRequest) SetLibId(v string) *DeleteRecognitionSampleRequest {
	s.LibId = &v
	return s
}

func (s *DeleteRecognitionSampleRequest) SetOwnerAccount(v string) *DeleteRecognitionSampleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteRecognitionSampleRequest) SetOwnerId(v int64) *DeleteRecognitionSampleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRecognitionSampleRequest) SetResourceOwnerAccount(v string) *DeleteRecognitionSampleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteRecognitionSampleRequest) SetResourceOwnerId(v int64) *DeleteRecognitionSampleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteRecognitionSampleRequest) SetSampleId(v string) *DeleteRecognitionSampleRequest {
	s.SampleId = &v
	return s
}

func (s *DeleteRecognitionSampleRequest) Validate() error {
	return dara.Validate(s)
}
