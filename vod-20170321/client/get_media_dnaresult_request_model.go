// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaDNAResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *GetMediaDNAResultRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *GetMediaDNAResultRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetMediaDNAResultRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetMediaDNAResultRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetMediaDNAResultRequest
	GetResourceOwnerId() *string
}

type GetMediaDNAResultRequest struct {
	// The ID of the video.
	//
	// This parameter is required.
	//
	// example:
	//
	// 88c6ca184c0e*****a5b665e2a126797
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetMediaDNAResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaDNAResultRequest) GoString() string {
	return s.String()
}

func (s *GetMediaDNAResultRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetMediaDNAResultRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetMediaDNAResultRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetMediaDNAResultRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetMediaDNAResultRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetMediaDNAResultRequest) SetMediaId(v string) *GetMediaDNAResultRequest {
	s.MediaId = &v
	return s
}

func (s *GetMediaDNAResultRequest) SetOwnerAccount(v string) *GetMediaDNAResultRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetMediaDNAResultRequest) SetOwnerId(v string) *GetMediaDNAResultRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMediaDNAResultRequest) SetResourceOwnerAccount(v string) *GetMediaDNAResultRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetMediaDNAResultRequest) SetResourceOwnerId(v string) *GetMediaDNAResultRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetMediaDNAResultRequest) Validate() error {
	return dara.Validate(s)
}
