// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComponentAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentAssetName(v string) *UpdateComponentAssetRequest
	GetComponentAssetName() *string
	SetComponentAssetUuid(v string) *UpdateComponentAssetRequest
	GetComponentAssetUuid() *string
	SetComponentAssetValues(v []*UpdateComponentAssetRequestComponentAssetValues) *UpdateComponentAssetRequest
	GetComponentAssetValues() []*UpdateComponentAssetRequestComponentAssetValues
	SetLang(v string) *UpdateComponentAssetRequest
	GetLang() *string
	SetRoleFor(v int64) *UpdateComponentAssetRequest
	GetRoleFor() *int64
}

type UpdateComponentAssetRequest struct {
	// Asset name.
	//
	// example:
	//
	// test_asset
	ComponentAssetName *string `json:"ComponentAssetName,omitempty" xml:"ComponentAssetName,omitempty"`
	// Asset UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1C5F11E9-****-51F0-****-43BB312A0557
	ComponentAssetUuid *string `json:"ComponentAssetUuid,omitempty" xml:"ComponentAssetUuid,omitempty"`
	// Configuration information of the asset.
	ComponentAssetValues []*UpdateComponentAssetRequestComponentAssetValues `json:"ComponentAssetValues,omitempty" xml:"ComponentAssetValues,omitempty" type:"Repeated"`
	// The language type for the request and response. Values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
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

func (s UpdateComponentAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateComponentAssetRequest) GoString() string {
	return s.String()
}

func (s *UpdateComponentAssetRequest) GetComponentAssetName() *string {
	return s.ComponentAssetName
}

func (s *UpdateComponentAssetRequest) GetComponentAssetUuid() *string {
	return s.ComponentAssetUuid
}

func (s *UpdateComponentAssetRequest) GetComponentAssetValues() []*UpdateComponentAssetRequestComponentAssetValues {
	return s.ComponentAssetValues
}

func (s *UpdateComponentAssetRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateComponentAssetRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateComponentAssetRequest) SetComponentAssetName(v string) *UpdateComponentAssetRequest {
	s.ComponentAssetName = &v
	return s
}

func (s *UpdateComponentAssetRequest) SetComponentAssetUuid(v string) *UpdateComponentAssetRequest {
	s.ComponentAssetUuid = &v
	return s
}

func (s *UpdateComponentAssetRequest) SetComponentAssetValues(v []*UpdateComponentAssetRequestComponentAssetValues) *UpdateComponentAssetRequest {
	s.ComponentAssetValues = v
	return s
}

func (s *UpdateComponentAssetRequest) SetLang(v string) *UpdateComponentAssetRequest {
	s.Lang = &v
	return s
}

func (s *UpdateComponentAssetRequest) SetRoleFor(v int64) *UpdateComponentAssetRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateComponentAssetRequest) Validate() error {
	if s.ComponentAssetValues != nil {
		for _, item := range s.ComponentAssetValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateComponentAssetRequestComponentAssetValues struct {
	// Field name.
	//
	// example:
	//
	// lh_source
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// Field value.
	//
	// example:
	//
	// device
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
}

func (s UpdateComponentAssetRequestComponentAssetValues) String() string {
	return dara.Prettify(s)
}

func (s UpdateComponentAssetRequestComponentAssetValues) GoString() string {
	return s.String()
}

func (s *UpdateComponentAssetRequestComponentAssetValues) GetFieldName() *string {
	return s.FieldName
}

func (s *UpdateComponentAssetRequestComponentAssetValues) GetFieldValue() *string {
	return s.FieldValue
}

func (s *UpdateComponentAssetRequestComponentAssetValues) SetFieldName(v string) *UpdateComponentAssetRequestComponentAssetValues {
	s.FieldName = &v
	return s
}

func (s *UpdateComponentAssetRequestComponentAssetValues) SetFieldValue(v string) *UpdateComponentAssetRequestComponentAssetValues {
	s.FieldValue = &v
	return s
}

func (s *UpdateComponentAssetRequestComponentAssetValues) Validate() error {
	return dara.Validate(s)
}
