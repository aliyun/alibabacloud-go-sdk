// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApiKeyDTO interface {
	dara.Model
	String() string
	GoString() string
	SetClient(v *ClientDTO) *ApiKeyDTO
	GetClient() *ClientDTO
	SetClientId(v int64) *ApiKeyDTO
	GetClientId() *int64
	SetGmtCreate(v string) *ApiKeyDTO
	GetGmtCreate() *string
	SetGmtModified(v string) *ApiKeyDTO
	GetGmtModified() *string
	SetId(v int64) *ApiKeyDTO
	GetId() *int64
	SetKeyPreview(v string) *ApiKeyDTO
	GetKeyPreview() *string
	SetName(v string) *ApiKeyDTO
	GetName() *string
}

type ApiKeyDTO struct {
	Client *ClientDTO `json:"client,omitempty" xml:"client,omitempty"`
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
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
	// sk-xxx****xxx
	KeyPreview *string `json:"keyPreview,omitempty" xml:"keyPreview,omitempty"`
	// example:
	//
	// MyApiKey
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ApiKeyDTO) String() string {
	return dara.Prettify(s)
}

func (s ApiKeyDTO) GoString() string {
	return s.String()
}

func (s *ApiKeyDTO) GetClient() *ClientDTO {
	return s.Client
}

func (s *ApiKeyDTO) GetClientId() *int64 {
	return s.ClientId
}

func (s *ApiKeyDTO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ApiKeyDTO) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ApiKeyDTO) GetId() *int64 {
	return s.Id
}

func (s *ApiKeyDTO) GetKeyPreview() *string {
	return s.KeyPreview
}

func (s *ApiKeyDTO) GetName() *string {
	return s.Name
}

func (s *ApiKeyDTO) SetClient(v *ClientDTO) *ApiKeyDTO {
	s.Client = v
	return s
}

func (s *ApiKeyDTO) SetClientId(v int64) *ApiKeyDTO {
	s.ClientId = &v
	return s
}

func (s *ApiKeyDTO) SetGmtCreate(v string) *ApiKeyDTO {
	s.GmtCreate = &v
	return s
}

func (s *ApiKeyDTO) SetGmtModified(v string) *ApiKeyDTO {
	s.GmtModified = &v
	return s
}

func (s *ApiKeyDTO) SetId(v int64) *ApiKeyDTO {
	s.Id = &v
	return s
}

func (s *ApiKeyDTO) SetKeyPreview(v string) *ApiKeyDTO {
	s.KeyPreview = &v
	return s
}

func (s *ApiKeyDTO) SetName(v string) *ApiKeyDTO {
	s.Name = &v
	return s
}

func (s *ApiKeyDTO) Validate() error {
	if s.Client != nil {
		if err := s.Client.Validate(); err != nil {
			return err
		}
	}
	return nil
}
