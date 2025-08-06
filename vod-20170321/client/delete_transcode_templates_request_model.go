// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTranscodeTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteTranscodeTemplatesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteTranscodeTemplatesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTranscodeTemplatesRequest
	GetResourceOwnerId() *int64
	SetTranscodeTemplateGroupId(v string) *DeleteTranscodeTemplatesRequest
	GetTranscodeTemplateGroupId() *string
	SetTranscodeTemplateIdList(v string) *DeleteTranscodeTemplatesRequest
	GetTranscodeTemplateIdList() *string
}

type DeleteTranscodeTemplatesRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
	// This parameter is required.
	TranscodeTemplateIdList *string `json:"TranscodeTemplateIdList,omitempty" xml:"TranscodeTemplateIdList,omitempty"`
}

func (s DeleteTranscodeTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTranscodeTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DeleteTranscodeTemplatesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTranscodeTemplatesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTranscodeTemplatesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTranscodeTemplatesRequest) GetTranscodeTemplateGroupId() *string {
	return s.TranscodeTemplateGroupId
}

func (s *DeleteTranscodeTemplatesRequest) GetTranscodeTemplateIdList() *string {
	return s.TranscodeTemplateIdList
}

func (s *DeleteTranscodeTemplatesRequest) SetOwnerId(v int64) *DeleteTranscodeTemplatesRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTranscodeTemplatesRequest) SetResourceOwnerAccount(v string) *DeleteTranscodeTemplatesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTranscodeTemplatesRequest) SetResourceOwnerId(v int64) *DeleteTranscodeTemplatesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTranscodeTemplatesRequest) SetTranscodeTemplateGroupId(v string) *DeleteTranscodeTemplatesRequest {
	s.TranscodeTemplateGroupId = &v
	return s
}

func (s *DeleteTranscodeTemplatesRequest) SetTranscodeTemplateIdList(v string) *DeleteTranscodeTemplatesRequest {
	s.TranscodeTemplateIdList = &v
	return s
}

func (s *DeleteTranscodeTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
