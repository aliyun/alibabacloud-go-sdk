// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComponentAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentAssetName(v string) *CreateComponentAssetRequest
	GetComponentAssetName() *string
	SetComponentAssetValues(v []*CreateComponentAssetRequestComponentAssetValues) *CreateComponentAssetRequest
	GetComponentAssetValues() []*CreateComponentAssetRequestComponentAssetValues
	SetComponentName(v string) *CreateComponentAssetRequest
	GetComponentName() *string
	SetLang(v string) *CreateComponentAssetRequest
	GetLang() *string
	SetRoleFor(v int64) *CreateComponentAssetRequest
	GetRoleFor() *int64
}

type CreateComponentAssetRequest struct {
	// The name of the asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// shanghai-log-config
	ComponentAssetName *string `json:"ComponentAssetName,omitempty" xml:"ComponentAssetName,omitempty"`
	// Configuration information of the asset.
	//
	// This parameter is required.
	ComponentAssetValues []*CreateComponentAssetRequestComponentAssetValues `json:"ComponentAssetValues,omitempty" xml:"ComponentAssetValues,omitempty" type:"Repeated"`
	// The name of the component.
	//
	// This parameter is required.
	//
	// example:
	//
	// SLS
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The language type for receiving messages. Values:
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

func (s CreateComponentAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateComponentAssetRequest) GoString() string {
	return s.String()
}

func (s *CreateComponentAssetRequest) GetComponentAssetName() *string {
	return s.ComponentAssetName
}

func (s *CreateComponentAssetRequest) GetComponentAssetValues() []*CreateComponentAssetRequestComponentAssetValues {
	return s.ComponentAssetValues
}

func (s *CreateComponentAssetRequest) GetComponentName() *string {
	return s.ComponentName
}

func (s *CreateComponentAssetRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateComponentAssetRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *CreateComponentAssetRequest) SetComponentAssetName(v string) *CreateComponentAssetRequest {
	s.ComponentAssetName = &v
	return s
}

func (s *CreateComponentAssetRequest) SetComponentAssetValues(v []*CreateComponentAssetRequestComponentAssetValues) *CreateComponentAssetRequest {
	s.ComponentAssetValues = v
	return s
}

func (s *CreateComponentAssetRequest) SetComponentName(v string) *CreateComponentAssetRequest {
	s.ComponentName = &v
	return s
}

func (s *CreateComponentAssetRequest) SetLang(v string) *CreateComponentAssetRequest {
	s.Lang = &v
	return s
}

func (s *CreateComponentAssetRequest) SetRoleFor(v int64) *CreateComponentAssetRequest {
	s.RoleFor = &v
	return s
}

func (s *CreateComponentAssetRequest) Validate() error {
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

type CreateComponentAssetRequestComponentAssetValues struct {
	// Field name.
	//
	// This parameter is required.
	//
	// example:
	//
	// endpoint
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// Field value.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://logs.aliyuncs.com
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
}

func (s CreateComponentAssetRequestComponentAssetValues) String() string {
	return dara.Prettify(s)
}

func (s CreateComponentAssetRequestComponentAssetValues) GoString() string {
	return s.String()
}

func (s *CreateComponentAssetRequestComponentAssetValues) GetFieldName() *string {
	return s.FieldName
}

func (s *CreateComponentAssetRequestComponentAssetValues) GetFieldValue() *string {
	return s.FieldValue
}

func (s *CreateComponentAssetRequestComponentAssetValues) SetFieldName(v string) *CreateComponentAssetRequestComponentAssetValues {
	s.FieldName = &v
	return s
}

func (s *CreateComponentAssetRequestComponentAssetValues) SetFieldValue(v string) *CreateComponentAssetRequestComponentAssetValues {
	s.FieldValue = &v
	return s
}

func (s *CreateComponentAssetRequestComponentAssetValues) Validate() error {
	return dara.Validate(s)
}
