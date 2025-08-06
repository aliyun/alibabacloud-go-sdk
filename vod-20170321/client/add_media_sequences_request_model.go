// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaSequencesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *AddMediaSequencesRequest
	GetMediaId() *string
	SetMediaSequences(v string) *AddMediaSequencesRequest
	GetMediaSequences() *string
	SetMediaType(v string) *AddMediaSequencesRequest
	GetMediaType() *string
	SetMediaURL(v string) *AddMediaSequencesRequest
	GetMediaURL() *string
	SetOwnerAccount(v string) *AddMediaSequencesRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *AddMediaSequencesRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *AddMediaSequencesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *AddMediaSequencesRequest
	GetResourceOwnerId() *string
}

type AddMediaSequencesRequest struct {
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	MediaSequences       *string `json:"MediaSequences,omitempty" xml:"MediaSequences,omitempty"`
	MediaType            *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	MediaURL             *string `json:"MediaURL,omitempty" xml:"MediaURL,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddMediaSequencesRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMediaSequencesRequest) GoString() string {
	return s.String()
}

func (s *AddMediaSequencesRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *AddMediaSequencesRequest) GetMediaSequences() *string {
	return s.MediaSequences
}

func (s *AddMediaSequencesRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *AddMediaSequencesRequest) GetMediaURL() *string {
	return s.MediaURL
}

func (s *AddMediaSequencesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddMediaSequencesRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *AddMediaSequencesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddMediaSequencesRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *AddMediaSequencesRequest) SetMediaId(v string) *AddMediaSequencesRequest {
	s.MediaId = &v
	return s
}

func (s *AddMediaSequencesRequest) SetMediaSequences(v string) *AddMediaSequencesRequest {
	s.MediaSequences = &v
	return s
}

func (s *AddMediaSequencesRequest) SetMediaType(v string) *AddMediaSequencesRequest {
	s.MediaType = &v
	return s
}

func (s *AddMediaSequencesRequest) SetMediaURL(v string) *AddMediaSequencesRequest {
	s.MediaURL = &v
	return s
}

func (s *AddMediaSequencesRequest) SetOwnerAccount(v string) *AddMediaSequencesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddMediaSequencesRequest) SetOwnerId(v string) *AddMediaSequencesRequest {
	s.OwnerId = &v
	return s
}

func (s *AddMediaSequencesRequest) SetResourceOwnerAccount(v string) *AddMediaSequencesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddMediaSequencesRequest) SetResourceOwnerId(v string) *AddMediaSequencesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddMediaSequencesRequest) Validate() error {
	return dara.Validate(s)
}
