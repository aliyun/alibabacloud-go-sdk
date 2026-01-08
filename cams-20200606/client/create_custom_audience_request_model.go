// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAudienceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdAccountId(v string) *CreateCustomAudienceRequest
	GetAdAccountId() *string
	SetCustSpaceId(v string) *CreateCustomAudienceRequest
	GetCustSpaceId() *string
	SetDescription(v string) *CreateCustomAudienceRequest
	GetDescription() *string
	SetFilePath(v string) *CreateCustomAudienceRequest
	GetFilePath() *string
	SetName(v string) *CreateCustomAudienceRequest
	GetName() *string
	SetOwnerId(v int64) *CreateCustomAudienceRequest
	GetOwnerId() *int64
	SetPageId(v string) *CreateCustomAudienceRequest
	GetPageId() *string
	SetResourceOwnerAccount(v string) *CreateCustomAudienceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateCustomAudienceRequest
	GetResourceOwnerId() *int64
	SetUploadType(v string) *CreateCustomAudienceRequest
	GetUploadType() *string
}

type CreateCustomAudienceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 23**
	AdAccountId *string `json:"AdAccountId,omitempty" xml:"AdAccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cams-***
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// bucket/file.xlsx
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// audience name
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 239***
	PageId               *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// excel
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
}

func (s CreateCustomAudienceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAudienceRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomAudienceRequest) GetAdAccountId() *string {
	return s.AdAccountId
}

func (s *CreateCustomAudienceRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *CreateCustomAudienceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCustomAudienceRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *CreateCustomAudienceRequest) GetName() *string {
	return s.Name
}

func (s *CreateCustomAudienceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCustomAudienceRequest) GetPageId() *string {
	return s.PageId
}

func (s *CreateCustomAudienceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCustomAudienceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCustomAudienceRequest) GetUploadType() *string {
	return s.UploadType
}

func (s *CreateCustomAudienceRequest) SetAdAccountId(v string) *CreateCustomAudienceRequest {
	s.AdAccountId = &v
	return s
}

func (s *CreateCustomAudienceRequest) SetCustSpaceId(v string) *CreateCustomAudienceRequest {
	s.CustSpaceId = &v
	return s
}

func (s *CreateCustomAudienceRequest) SetDescription(v string) *CreateCustomAudienceRequest {
	s.Description = &v
	return s
}

func (s *CreateCustomAudienceRequest) SetFilePath(v string) *CreateCustomAudienceRequest {
	s.FilePath = &v
	return s
}

func (s *CreateCustomAudienceRequest) SetName(v string) *CreateCustomAudienceRequest {
	s.Name = &v
	return s
}

func (s *CreateCustomAudienceRequest) SetOwnerId(v int64) *CreateCustomAudienceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCustomAudienceRequest) SetPageId(v string) *CreateCustomAudienceRequest {
	s.PageId = &v
	return s
}

func (s *CreateCustomAudienceRequest) SetResourceOwnerAccount(v string) *CreateCustomAudienceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCustomAudienceRequest) SetResourceOwnerId(v int64) *CreateCustomAudienceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCustomAudienceRequest) SetUploadType(v string) *CreateCustomAudienceRequest {
	s.UploadType = &v
	return s
}

func (s *CreateCustomAudienceRequest) Validate() error {
	return dara.Validate(s)
}
