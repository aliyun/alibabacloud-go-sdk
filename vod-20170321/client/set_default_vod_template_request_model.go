// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultVodTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *SetDefaultVodTemplateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SetDefaultVodTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetDefaultVodTemplateRequest
	GetResourceOwnerId() *int64
	SetVodTemplateId(v string) *SetDefaultVodTemplateRequest
	GetVodTemplateId() *string
}

type SetDefaultVodTemplateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	VodTemplateId *string `json:"VodTemplateId,omitempty" xml:"VodTemplateId,omitempty"`
}

func (s SetDefaultVodTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultVodTemplateRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultVodTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetDefaultVodTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetDefaultVodTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetDefaultVodTemplateRequest) GetVodTemplateId() *string {
	return s.VodTemplateId
}

func (s *SetDefaultVodTemplateRequest) SetOwnerId(v int64) *SetDefaultVodTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDefaultVodTemplateRequest) SetResourceOwnerAccount(v string) *SetDefaultVodTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetDefaultVodTemplateRequest) SetResourceOwnerId(v int64) *SetDefaultVodTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetDefaultVodTemplateRequest) SetVodTemplateId(v string) *SetDefaultVodTemplateRequest {
	s.VodTemplateId = &v
	return s
}

func (s *SetDefaultVodTemplateRequest) Validate() error {
	return dara.Validate(s)
}
