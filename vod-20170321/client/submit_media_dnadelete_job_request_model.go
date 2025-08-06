// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaDNADeleteJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *SubmitMediaDNADeleteJobRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *SubmitMediaDNADeleteJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SubmitMediaDNADeleteJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SubmitMediaDNADeleteJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SubmitMediaDNADeleteJobRequest
	GetResourceOwnerId() *string
}

type SubmitMediaDNADeleteJobRequest struct {
	// The ID of the video.
	//
	// This parameter is required.
	//
	// example:
	//
	// 656eaaa8c43a4597******1f09a36
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SubmitMediaDNADeleteJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaDNADeleteJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitMediaDNADeleteJobRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitMediaDNADeleteJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitMediaDNADeleteJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SubmitMediaDNADeleteJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitMediaDNADeleteJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SubmitMediaDNADeleteJobRequest) SetMediaId(v string) *SubmitMediaDNADeleteJobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitMediaDNADeleteJobRequest) SetOwnerAccount(v string) *SubmitMediaDNADeleteJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitMediaDNADeleteJobRequest) SetOwnerId(v string) *SubmitMediaDNADeleteJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitMediaDNADeleteJobRequest) SetResourceOwnerAccount(v string) *SubmitMediaDNADeleteJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitMediaDNADeleteJobRequest) SetResourceOwnerId(v string) *SubmitMediaDNADeleteJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitMediaDNADeleteJobRequest) Validate() error {
	return dara.Validate(s)
}
