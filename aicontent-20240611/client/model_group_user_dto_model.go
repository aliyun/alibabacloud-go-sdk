// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelGroupUserDTO interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeys(v []*ModelGroupClientKeyItemDTO) *ModelGroupUserDTO
	GetApiKeys() []*ModelGroupClientKeyItemDTO
	SetUserId(v int64) *ModelGroupUserDTO
	GetUserId() *int64
	SetUserName(v string) *ModelGroupUserDTO
	GetUserName() *string
}

type ModelGroupUserDTO struct {
	// example:
	//
	// []
	ApiKeys []*ModelGroupClientKeyItemDTO `json:"apiKeys,omitempty" xml:"apiKeys,omitempty" type:"Repeated"`
	// example:
	//
	// 30001
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// Zhang San
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s ModelGroupUserDTO) String() string {
	return dara.Prettify(s)
}

func (s ModelGroupUserDTO) GoString() string {
	return s.String()
}

func (s *ModelGroupUserDTO) GetApiKeys() []*ModelGroupClientKeyItemDTO {
	return s.ApiKeys
}

func (s *ModelGroupUserDTO) GetUserId() *int64 {
	return s.UserId
}

func (s *ModelGroupUserDTO) GetUserName() *string {
	return s.UserName
}

func (s *ModelGroupUserDTO) SetApiKeys(v []*ModelGroupClientKeyItemDTO) *ModelGroupUserDTO {
	s.ApiKeys = v
	return s
}

func (s *ModelGroupUserDTO) SetUserId(v int64) *ModelGroupUserDTO {
	s.UserId = &v
	return s
}

func (s *ModelGroupUserDTO) SetUserName(v string) *ModelGroupUserDTO {
	s.UserName = &v
	return s
}

func (s *ModelGroupUserDTO) Validate() error {
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
