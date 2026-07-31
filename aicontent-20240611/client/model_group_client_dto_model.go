// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelGroupClientDTO interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeys(v []*ModelGroupClientKeyItemDTO) *ModelGroupClientDTO
	GetApiKeys() []*ModelGroupClientKeyItemDTO
	SetClientId(v int64) *ModelGroupClientDTO
	GetClientId() *int64
	SetClientName(v string) *ModelGroupClientDTO
	GetClientName() *string
}

type ModelGroupClientDTO struct {
	// example:
	//
	// []
	ApiKeys []*ModelGroupClientKeyItemDTO `json:"apiKeys,omitempty" xml:"apiKeys,omitempty" type:"Repeated"`
	// example:
	//
	// 1001
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// example:
	//
	// UserA-Professional
	ClientName *string `json:"clientName,omitempty" xml:"clientName,omitempty"`
}

func (s ModelGroupClientDTO) String() string {
	return dara.Prettify(s)
}

func (s ModelGroupClientDTO) GoString() string {
	return s.String()
}

func (s *ModelGroupClientDTO) GetApiKeys() []*ModelGroupClientKeyItemDTO {
	return s.ApiKeys
}

func (s *ModelGroupClientDTO) GetClientId() *int64 {
	return s.ClientId
}

func (s *ModelGroupClientDTO) GetClientName() *string {
	return s.ClientName
}

func (s *ModelGroupClientDTO) SetApiKeys(v []*ModelGroupClientKeyItemDTO) *ModelGroupClientDTO {
	s.ApiKeys = v
	return s
}

func (s *ModelGroupClientDTO) SetClientId(v int64) *ModelGroupClientDTO {
	s.ClientId = &v
	return s
}

func (s *ModelGroupClientDTO) SetClientName(v string) *ModelGroupClientDTO {
	s.ClientName = &v
	return s
}

func (s *ModelGroupClientDTO) Validate() error {
	if s.ApiKeys != nil {
		for _, item := range s.ApiKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
