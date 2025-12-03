// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *CreateDatasetRequest
	GetDatasetName() *string
	SetDatasetType(v string) *CreateDatasetRequest
	GetDatasetType() *string
	SetDescription(v string) *CreateDatasetRequest
	GetDescription() *string
	SetSecurityToken(v string) *CreateDatasetRequest
	GetSecurityToken() *string
	SetTag(v []*CreateDatasetRequestTag) *CreateDatasetRequest
	GetTag() []*CreateDatasetRequestTag
}

type CreateDatasetRequest struct {
	// The name of the dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// DatasetName
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The type of the dataset. Valid values:
	//
	// 	- JWT_BLOCKING: a JSON Web Token (JWT) blacklist
	//
	// 	- IP_WHITELIST_CIDR : an IP address whitelist
	//
	// 	- PARAMETER_ACCESS : parameter-based access control
	//
	// This parameter is required.
	//
	// example:
	//
	// JWT_BLOCKING
	DatasetType   *string `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The tag of objects that match the rule. You can specify multiple tags.
	Tag []*CreateDatasetRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateDatasetRequest) GetDatasetType() *string {
	return s.DatasetType
}

func (s *CreateDatasetRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDatasetRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateDatasetRequest) GetTag() []*CreateDatasetRequestTag {
	return s.Tag
}

func (s *CreateDatasetRequest) SetDatasetName(v string) *CreateDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateDatasetRequest) SetDatasetType(v string) *CreateDatasetRequest {
	s.DatasetType = &v
	return s
}

func (s *CreateDatasetRequest) SetDescription(v string) *CreateDatasetRequest {
	s.Description = &v
	return s
}

func (s *CreateDatasetRequest) SetSecurityToken(v string) *CreateDatasetRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateDatasetRequest) SetTag(v []*CreateDatasetRequestTag) *CreateDatasetRequest {
	s.Tag = v
	return s
}

func (s *CreateDatasetRequest) Validate() error {
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

type CreateDatasetRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// 123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDatasetRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDatasetRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDatasetRequestTag) SetKey(v string) *CreateDatasetRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDatasetRequestTag) SetValue(v string) *CreateDatasetRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDatasetRequestTag) Validate() error {
	return dara.Validate(s)
}
