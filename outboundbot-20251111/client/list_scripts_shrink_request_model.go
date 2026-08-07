// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListScriptsShrinkRequest
	GetInstanceId() *string
	SetName(v string) *ListScriptsShrinkRequest
	GetName() *string
	SetPageNumber(v int32) *ListScriptsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListScriptsShrinkRequest
	GetPageSize() *int32
	SetPublishOnly(v bool) *ListScriptsShrinkRequest
	GetPublishOnly() *bool
	SetScriptIdsShrink(v string) *ListScriptsShrinkRequest
	GetScriptIdsShrink() *string
}

type ListScriptsShrinkRequest struct {
	// 实例ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 搜索关键词
	//
	// example:
	//
	// 满意度调研
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 页码，从1开始
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页记录数
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 是否仅返回已发布的场景
	//
	// example:
	//
	// true
	PublishOnly *bool `json:"PublishOnly,omitempty" xml:"PublishOnly,omitempty"`
	// 场景ID列表
	ScriptIdsShrink *string `json:"ScriptIds,omitempty" xml:"ScriptIds,omitempty"`
}

func (s ListScriptsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListScriptsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListScriptsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListScriptsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListScriptsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScriptsShrinkRequest) GetPublishOnly() *bool {
	return s.PublishOnly
}

func (s *ListScriptsShrinkRequest) GetScriptIdsShrink() *string {
	return s.ScriptIdsShrink
}

func (s *ListScriptsShrinkRequest) SetInstanceId(v string) *ListScriptsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScriptsShrinkRequest) SetName(v string) *ListScriptsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListScriptsShrinkRequest) SetPageNumber(v int32) *ListScriptsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListScriptsShrinkRequest) SetPageSize(v int32) *ListScriptsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListScriptsShrinkRequest) SetPublishOnly(v bool) *ListScriptsShrinkRequest {
	s.PublishOnly = &v
	return s
}

func (s *ListScriptsShrinkRequest) SetScriptIdsShrink(v string) *ListScriptsShrinkRequest {
	s.ScriptIdsShrink = &v
	return s
}

func (s *ListScriptsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
