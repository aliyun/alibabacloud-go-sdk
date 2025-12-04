// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateErRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateErRequest
	GetDescription() *string
	SetErName(v string) *CreateErRequest
	GetErName() *string
	SetMasterZoneId(v string) *CreateErRequest
	GetMasterZoneId() *string
	SetRegionId(v string) *CreateErRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateErRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateErRequestTag) *CreateErRequest
	GetTag() []*CreateErRequestTag
}

type CreateErRequest struct {
	// The description of the document.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun HUB Name
	//
	// This parameter is required.
	//
	// example:
	//
	// er-wulanchabu-main
	ErName *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	// Primary Zone
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu-b
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmyuzlx2iihcy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// List of tags.
	Tag []*CreateErRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateErRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateErRequest) GoString() string {
	return s.String()
}

func (s *CreateErRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateErRequest) GetErName() *string {
	return s.ErName
}

func (s *CreateErRequest) GetMasterZoneId() *string {
	return s.MasterZoneId
}

func (s *CreateErRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateErRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateErRequest) GetTag() []*CreateErRequestTag {
	return s.Tag
}

func (s *CreateErRequest) SetDescription(v string) *CreateErRequest {
	s.Description = &v
	return s
}

func (s *CreateErRequest) SetErName(v string) *CreateErRequest {
	s.ErName = &v
	return s
}

func (s *CreateErRequest) SetMasterZoneId(v string) *CreateErRequest {
	s.MasterZoneId = &v
	return s
}

func (s *CreateErRequest) SetRegionId(v string) *CreateErRequest {
	s.RegionId = &v
	return s
}

func (s *CreateErRequest) SetResourceGroupId(v string) *CreateErRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateErRequest) SetTag(v []*CreateErRequestTag) *CreateErRequest {
	s.Tag = v
	return s
}

func (s *CreateErRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateErRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key-test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value-test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateErRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateErRequestTag) GoString() string {
	return s.String()
}

func (s *CreateErRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateErRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateErRequestTag) SetKey(v string) *CreateErRequestTag {
	s.Key = &v
	return s
}

func (s *CreateErRequestTag) SetValue(v string) *CreateErRequestTag {
	s.Value = &v
	return s
}

func (s *CreateErRequestTag) Validate() error {
	return dara.Validate(s)
}
