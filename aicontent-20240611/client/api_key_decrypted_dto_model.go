// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApiKeyDecryptedDTO interface {
	dara.Model
	String() string
	GoString() string
	SetClient(v *ClientDTO) *ApiKeyDecryptedDTO
	GetClient() *ClientDTO
	SetClientId(v int64) *ApiKeyDecryptedDTO
	GetClientId() *int64
	SetDeleteTag(v int32) *ApiKeyDecryptedDTO
	GetDeleteTag() *int32
	SetGmtCreate(v string) *ApiKeyDecryptedDTO
	GetGmtCreate() *string
	SetGmtModified(v string) *ApiKeyDecryptedDTO
	GetGmtModified() *string
	SetId(v int64) *ApiKeyDecryptedDTO
	GetId() *int64
	SetKey(v string) *ApiKeyDecryptedDTO
	GetKey() *string
	SetKeyPreview(v string) *ApiKeyDecryptedDTO
	GetKeyPreview() *string
	SetName(v string) *ApiKeyDecryptedDTO
	GetName() *string
}

type ApiKeyDecryptedDTO struct {
	Client *ClientDTO `json:"client,omitempty" xml:"client,omitempty"`
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// example:
	//
	// 0
	DeleteTag *int32 `json:"deleteTag,omitempty" xml:"deleteTag,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// sk-xxxxxxxxxxxxxxxxxxx
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// sk-xxx****xxx
	KeyPreview *string `json:"keyPreview,omitempty" xml:"keyPreview,omitempty"`
	// example:
	//
	// MyApiKey
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ApiKeyDecryptedDTO) String() string {
	return dara.Prettify(s)
}

func (s ApiKeyDecryptedDTO) GoString() string {
	return s.String()
}

func (s *ApiKeyDecryptedDTO) GetClient() *ClientDTO {
	return s.Client
}

func (s *ApiKeyDecryptedDTO) GetClientId() *int64 {
	return s.ClientId
}

func (s *ApiKeyDecryptedDTO) GetDeleteTag() *int32 {
	return s.DeleteTag
}

func (s *ApiKeyDecryptedDTO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ApiKeyDecryptedDTO) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ApiKeyDecryptedDTO) GetId() *int64 {
	return s.Id
}

func (s *ApiKeyDecryptedDTO) GetKey() *string {
	return s.Key
}

func (s *ApiKeyDecryptedDTO) GetKeyPreview() *string {
	return s.KeyPreview
}

func (s *ApiKeyDecryptedDTO) GetName() *string {
	return s.Name
}

func (s *ApiKeyDecryptedDTO) SetClient(v *ClientDTO) *ApiKeyDecryptedDTO {
	s.Client = v
	return s
}

func (s *ApiKeyDecryptedDTO) SetClientId(v int64) *ApiKeyDecryptedDTO {
	s.ClientId = &v
	return s
}

func (s *ApiKeyDecryptedDTO) SetDeleteTag(v int32) *ApiKeyDecryptedDTO {
	s.DeleteTag = &v
	return s
}

func (s *ApiKeyDecryptedDTO) SetGmtCreate(v string) *ApiKeyDecryptedDTO {
	s.GmtCreate = &v
	return s
}

func (s *ApiKeyDecryptedDTO) SetGmtModified(v string) *ApiKeyDecryptedDTO {
	s.GmtModified = &v
	return s
}

func (s *ApiKeyDecryptedDTO) SetId(v int64) *ApiKeyDecryptedDTO {
	s.Id = &v
	return s
}

func (s *ApiKeyDecryptedDTO) SetKey(v string) *ApiKeyDecryptedDTO {
	s.Key = &v
	return s
}

func (s *ApiKeyDecryptedDTO) SetKeyPreview(v string) *ApiKeyDecryptedDTO {
	s.KeyPreview = &v
	return s
}

func (s *ApiKeyDecryptedDTO) SetName(v string) *ApiKeyDecryptedDTO {
	s.Name = &v
	return s
}

func (s *ApiKeyDecryptedDTO) Validate() error {
	if s.Client != nil {
		if err := s.Client.Validate(); err != nil {
			return err
		}
	}
	return nil
}
