// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiGroupInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ModifyApiGroupInstanceRequest
	GetGroupId() *string
	SetRemark(v string) *ModifyApiGroupInstanceRequest
	GetRemark() *string
	SetSecurityToken(v string) *ModifyApiGroupInstanceRequest
	GetSecurityToken() *string
	SetTag(v []*ModifyApiGroupInstanceRequestTag) *ModifyApiGroupInstanceRequest
	GetTag() []*ModifyApiGroupInstanceRequestTag
	SetTargetInstanceId(v string) *ModifyApiGroupInstanceRequest
	GetTargetInstanceId() *string
}

type ModifyApiGroupInstanceRequest struct {
	// The ID of the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 01c97ed08a614118849b00079753d1e2
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The remarks.
	//
	// example:
	//
	// migrate
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The tag of objects that match the rule. You can specify multiple tags.
	Tag []*ModifyApiGroupInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the instance to which you want to migrate the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// apigateway-bj-c325375b1ebe
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
}

func (s ModifyApiGroupInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiGroupInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupInstanceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyApiGroupInstanceRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModifyApiGroupInstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyApiGroupInstanceRequest) GetTag() []*ModifyApiGroupInstanceRequestTag {
	return s.Tag
}

func (s *ModifyApiGroupInstanceRequest) GetTargetInstanceId() *string {
	return s.TargetInstanceId
}

func (s *ModifyApiGroupInstanceRequest) SetGroupId(v string) *ModifyApiGroupInstanceRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyApiGroupInstanceRequest) SetRemark(v string) *ModifyApiGroupInstanceRequest {
	s.Remark = &v
	return s
}

func (s *ModifyApiGroupInstanceRequest) SetSecurityToken(v string) *ModifyApiGroupInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyApiGroupInstanceRequest) SetTag(v []*ModifyApiGroupInstanceRequestTag) *ModifyApiGroupInstanceRequest {
	s.Tag = v
	return s
}

func (s *ModifyApiGroupInstanceRequest) SetTargetInstanceId(v string) *ModifyApiGroupInstanceRequest {
	s.TargetInstanceId = &v
	return s
}

func (s *ModifyApiGroupInstanceRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyApiGroupInstanceRequestTag struct {
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyApiGroupInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiGroupInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *ModifyApiGroupInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *ModifyApiGroupInstanceRequestTag) SetKey(v string) *ModifyApiGroupInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *ModifyApiGroupInstanceRequestTag) SetValue(v string) *ModifyApiGroupInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *ModifyApiGroupInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
