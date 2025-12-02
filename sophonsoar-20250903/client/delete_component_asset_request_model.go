// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComponentAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentAssetUuid(v string) *DeleteComponentAssetRequest
	GetComponentAssetUuid() *string
	SetLang(v string) *DeleteComponentAssetRequest
	GetLang() *string
	SetRoleFor(v int64) *DeleteComponentAssetRequest
	GetRoleFor() *int64
}

type DeleteComponentAssetRequest struct {
	// Asset UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1C5F11E9-****-51F0-****-43BB312A0557
	ComponentAssetUuid *string `json:"ComponentAssetUuid,omitempty" xml:"ComponentAssetUuid,omitempty"`
	// Set the language type for requests and received messages, default is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Resource directory member account ID.
	//
	// example:
	//
	// 13760*****718726
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s DeleteComponentAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteComponentAssetRequest) GoString() string {
	return s.String()
}

func (s *DeleteComponentAssetRequest) GetComponentAssetUuid() *string {
	return s.ComponentAssetUuid
}

func (s *DeleteComponentAssetRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteComponentAssetRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DeleteComponentAssetRequest) SetComponentAssetUuid(v string) *DeleteComponentAssetRequest {
	s.ComponentAssetUuid = &v
	return s
}

func (s *DeleteComponentAssetRequest) SetLang(v string) *DeleteComponentAssetRequest {
	s.Lang = &v
	return s
}

func (s *DeleteComponentAssetRequest) SetRoleFor(v int64) *DeleteComponentAssetRequest {
	s.RoleFor = &v
	return s
}

func (s *DeleteComponentAssetRequest) Validate() error {
	return dara.Validate(s)
}
