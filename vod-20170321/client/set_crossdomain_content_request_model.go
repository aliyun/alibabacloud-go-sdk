// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCrossdomainContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *SetCrossdomainContentRequest
	GetContent() *string
	SetOwnerAccount(v string) *SetCrossdomainContentRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SetCrossdomainContentRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SetCrossdomainContentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SetCrossdomainContentRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v string) *SetCrossdomainContentRequest
	GetResourceRealOwnerId() *string
	SetStorageLocation(v string) *SetCrossdomainContentRequest
	GetStorageLocation() *string
}

type SetCrossdomainContentRequest struct {
	// The content of the cross-domain policy file. The file must be in the XML format and can contain up to 2,048 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// &lt;cross-domain-policy&gt;&lt;allow-access-from domain="*"/&gt;&lt;allow-http-request-headers-from domain="*" headers="*" secure="false"/&gt;&lt;/cross-domain-policy&gt;
	Content              *string `json:"Content,omitempty" xml:"Content,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the resource owner.
	//
	// example:
	//
	// 3461111
	ResourceRealOwnerId *string `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	// The URL of the Object Storage Service (OSS) bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// outin-67870fd5b****1e98a3900163e1c35d5.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
}

func (s SetCrossdomainContentRequest) String() string {
	return dara.Prettify(s)
}

func (s SetCrossdomainContentRequest) GoString() string {
	return s.String()
}

func (s *SetCrossdomainContentRequest) GetContent() *string {
	return s.Content
}

func (s *SetCrossdomainContentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetCrossdomainContentRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SetCrossdomainContentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetCrossdomainContentRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SetCrossdomainContentRequest) GetResourceRealOwnerId() *string {
	return s.ResourceRealOwnerId
}

func (s *SetCrossdomainContentRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *SetCrossdomainContentRequest) SetContent(v string) *SetCrossdomainContentRequest {
	s.Content = &v
	return s
}

func (s *SetCrossdomainContentRequest) SetOwnerAccount(v string) *SetCrossdomainContentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetCrossdomainContentRequest) SetOwnerId(v string) *SetCrossdomainContentRequest {
	s.OwnerId = &v
	return s
}

func (s *SetCrossdomainContentRequest) SetResourceOwnerAccount(v string) *SetCrossdomainContentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetCrossdomainContentRequest) SetResourceOwnerId(v string) *SetCrossdomainContentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetCrossdomainContentRequest) SetResourceRealOwnerId(v string) *SetCrossdomainContentRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *SetCrossdomainContentRequest) SetStorageLocation(v string) *SetCrossdomainContentRequest {
	s.StorageLocation = &v
	return s
}

func (s *SetCrossdomainContentRequest) Validate() error {
	return dara.Validate(s)
}
