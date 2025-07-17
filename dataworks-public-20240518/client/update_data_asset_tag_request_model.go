// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataAssetTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateDataAssetTagRequest
	GetDescription() *string
	SetKey(v string) *UpdateDataAssetTagRequest
	GetKey() *string
	SetManagers(v []*string) *UpdateDataAssetTagRequest
	GetManagers() []*string
	SetValues(v []*string) *UpdateDataAssetTagRequest
	GetValues() []*string
}

type UpdateDataAssetTagRequest struct {
	// The description of the tag.
	//
	// example:
	//
	// This is a description.
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
	// The tag values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateDataAssetTagRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAssetTagRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataAssetTagRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDataAssetTagRequest) GetKey() *string {
	return s.Key
}

func (s *UpdateDataAssetTagRequest) GetManagers() []*string {
	return s.Managers
}

func (s *UpdateDataAssetTagRequest) GetValues() []*string {
	return s.Values
}

func (s *UpdateDataAssetTagRequest) SetDescription(v string) *UpdateDataAssetTagRequest {
	s.Description = &v
	return s
}

func (s *UpdateDataAssetTagRequest) SetKey(v string) *UpdateDataAssetTagRequest {
	s.Key = &v
	return s
}

func (s *UpdateDataAssetTagRequest) SetManagers(v []*string) *UpdateDataAssetTagRequest {
	s.Managers = v
	return s
}

func (s *UpdateDataAssetTagRequest) SetValues(v []*string) *UpdateDataAssetTagRequest {
	s.Values = v
	return s
}

func (s *UpdateDataAssetTagRequest) Validate() error {
	return dara.Validate(s)
}
