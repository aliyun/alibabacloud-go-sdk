// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSystemConfigsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigsShrink(v string) *UpdateSystemConfigsShrinkRequest
	GetConfigsShrink() *string
	SetObjectId(v string) *UpdateSystemConfigsShrinkRequest
	GetObjectId() *string
	SetObjectType(v string) *UpdateSystemConfigsShrinkRequest
	GetObjectType() *string
}

type UpdateSystemConfigsShrinkRequest struct {
	// 配置列表
	ConfigsShrink *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	// 对象ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// 外呼开发时补充参数限制
	//
	// example:
	//
	// INSTANCE
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
}

func (s UpdateSystemConfigsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSystemConfigsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSystemConfigsShrinkRequest) GetConfigsShrink() *string {
	return s.ConfigsShrink
}

func (s *UpdateSystemConfigsShrinkRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *UpdateSystemConfigsShrinkRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *UpdateSystemConfigsShrinkRequest) SetConfigsShrink(v string) *UpdateSystemConfigsShrinkRequest {
	s.ConfigsShrink = &v
	return s
}

func (s *UpdateSystemConfigsShrinkRequest) SetObjectId(v string) *UpdateSystemConfigsShrinkRequest {
	s.ObjectId = &v
	return s
}

func (s *UpdateSystemConfigsShrinkRequest) SetObjectType(v string) *UpdateSystemConfigsShrinkRequest {
	s.ObjectType = &v
	return s
}

func (s *UpdateSystemConfigsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
