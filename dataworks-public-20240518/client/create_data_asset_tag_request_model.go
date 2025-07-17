// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAssetTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateDataAssetTagRequest
	GetDescription() *string
	SetKey(v string) *CreateDataAssetTagRequest
	GetKey() *string
	SetManagers(v []*string) *CreateDataAssetTagRequest
	GetManagers() []*string
	SetValueType(v string) *CreateDataAssetTagRequest
	GetValueType() *string
	SetValues(v []*string) *CreateDataAssetTagRequest
	GetValues() []*string
}

type CreateDataAssetTagRequest struct {
	// The description of the tag.
	//
	// example:
	//
	// This is a description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag administrators.
	Managers []*string `json:"Managers,omitempty" xml:"Managers,omitempty" type:"Repeated"`
	// The type of the tag value. Valid values:
	//
	// 	- Boolean
	//
	// 	- Int
	//
	// 	- String
	//
	// 	- Double
	//
	// example:
	//
	// String
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
	// The tag values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateDataAssetTagRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAssetTagRequest) GoString() string {
	return s.String()
}

func (s *CreateDataAssetTagRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataAssetTagRequest) GetKey() *string {
	return s.Key
}

func (s *CreateDataAssetTagRequest) GetManagers() []*string {
	return s.Managers
}

func (s *CreateDataAssetTagRequest) GetValueType() *string {
	return s.ValueType
}

func (s *CreateDataAssetTagRequest) GetValues() []*string {
	return s.Values
}

func (s *CreateDataAssetTagRequest) SetDescription(v string) *CreateDataAssetTagRequest {
	s.Description = &v
	return s
}

func (s *CreateDataAssetTagRequest) SetKey(v string) *CreateDataAssetTagRequest {
	s.Key = &v
	return s
}

func (s *CreateDataAssetTagRequest) SetManagers(v []*string) *CreateDataAssetTagRequest {
	s.Managers = v
	return s
}

func (s *CreateDataAssetTagRequest) SetValueType(v string) *CreateDataAssetTagRequest {
	s.ValueType = &v
	return s
}

func (s *CreateDataAssetTagRequest) SetValues(v []*string) *CreateDataAssetTagRequest {
	s.Values = v
	return s
}

func (s *CreateDataAssetTagRequest) Validate() error {
	return dara.Validate(s)
}
