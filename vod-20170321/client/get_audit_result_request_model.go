// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *GetAuditResultRequest
	GetMediaId() *string
	SetMediaType(v string) *GetAuditResultRequest
	GetMediaType() *string
	SetOwnerAccount(v string) *GetAuditResultRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetAuditResultRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetAuditResultRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetAuditResultRequest
	GetResourceOwnerId() *string
	SetVideoId(v string) *GetAuditResultRequest
	GetVideoId() *string
}

type GetAuditResultRequest struct {
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	MediaType            *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VideoId              *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetAuditResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuditResultRequest) GoString() string {
	return s.String()
}

func (s *GetAuditResultRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetAuditResultRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *GetAuditResultRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetAuditResultRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetAuditResultRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetAuditResultRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetAuditResultRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *GetAuditResultRequest) SetMediaId(v string) *GetAuditResultRequest {
	s.MediaId = &v
	return s
}

func (s *GetAuditResultRequest) SetMediaType(v string) *GetAuditResultRequest {
	s.MediaType = &v
	return s
}

func (s *GetAuditResultRequest) SetOwnerAccount(v string) *GetAuditResultRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetAuditResultRequest) SetOwnerId(v string) *GetAuditResultRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAuditResultRequest) SetResourceOwnerAccount(v string) *GetAuditResultRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAuditResultRequest) SetResourceOwnerId(v string) *GetAuditResultRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAuditResultRequest) SetVideoId(v string) *GetAuditResultRequest {
	s.VideoId = &v
	return s
}

func (s *GetAuditResultRequest) Validate() error {
	return dara.Validate(s)
}
