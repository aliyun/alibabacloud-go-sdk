// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoDNAResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *GetVideoDNAResultRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *GetVideoDNAResultRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetVideoDNAResultRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetVideoDNAResultRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetVideoDNAResultRequest
	GetResourceOwnerId() *string
}

type GetVideoDNAResultRequest struct {
	// This parameter is required.
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetVideoDNAResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDNAResultRequest) GoString() string {
	return s.String()
}

func (s *GetVideoDNAResultRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetVideoDNAResultRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetVideoDNAResultRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetVideoDNAResultRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetVideoDNAResultRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetVideoDNAResultRequest) SetMediaId(v string) *GetVideoDNAResultRequest {
	s.MediaId = &v
	return s
}

func (s *GetVideoDNAResultRequest) SetOwnerAccount(v string) *GetVideoDNAResultRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetVideoDNAResultRequest) SetOwnerId(v string) *GetVideoDNAResultRequest {
	s.OwnerId = &v
	return s
}

func (s *GetVideoDNAResultRequest) SetResourceOwnerAccount(v string) *GetVideoDNAResultRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetVideoDNAResultRequest) SetResourceOwnerId(v string) *GetVideoDNAResultRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetVideoDNAResultRequest) Validate() error {
	return dara.Validate(s)
}
