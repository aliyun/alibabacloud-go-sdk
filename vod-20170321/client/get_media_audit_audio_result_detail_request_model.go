// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaAuditAudioResultDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *GetMediaAuditAudioResultDetailRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *GetMediaAuditAudioResultDetailRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetMediaAuditAudioResultDetailRequest
	GetOwnerId() *string
	SetPageNo(v int32) *GetMediaAuditAudioResultDetailRequest
	GetPageNo() *int32
	SetResourceOwnerAccount(v string) *GetMediaAuditAudioResultDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetMediaAuditAudioResultDetailRequest
	GetResourceOwnerId() *string
}

type GetMediaAuditAudioResultDetailRequest struct {
	// The ID of the video. You can query the video ID by using the ApsaraVideo VOD console or calling the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 93ab850b4f6f*****54b6e91d24d81d4
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. This parameter is optional. If you do not specify this parameter, all results are returned without pagination.
	//
	// example:
	//
	// 1
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetMediaAuditAudioResultDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditAudioResultDetailRequest) GoString() string {
	return s.String()
}

func (s *GetMediaAuditAudioResultDetailRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetMediaAuditAudioResultDetailRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetMediaAuditAudioResultDetailRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetMediaAuditAudioResultDetailRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetMediaAuditAudioResultDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetMediaAuditAudioResultDetailRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetMediaAuditAudioResultDetailRequest) SetMediaId(v string) *GetMediaAuditAudioResultDetailRequest {
	s.MediaId = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailRequest) SetOwnerAccount(v string) *GetMediaAuditAudioResultDetailRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailRequest) SetOwnerId(v string) *GetMediaAuditAudioResultDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailRequest) SetPageNo(v int32) *GetMediaAuditAudioResultDetailRequest {
	s.PageNo = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailRequest) SetResourceOwnerAccount(v string) *GetMediaAuditAudioResultDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailRequest) SetResourceOwnerId(v string) *GetMediaAuditAudioResultDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailRequest) Validate() error {
	return dara.Validate(s)
}
