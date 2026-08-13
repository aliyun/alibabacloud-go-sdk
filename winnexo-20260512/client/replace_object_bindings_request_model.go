// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceObjectBindingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetObjectBindings(v []*ReplaceObjectBindingsRequestObjectBindings) *ReplaceObjectBindingsRequest
	GetObjectBindings() []*ReplaceObjectBindingsRequestObjectBindings
	SetSourceId(v string) *ReplaceObjectBindingsRequest
	GetSourceId() *string
	SetTenantId(v string) *ReplaceObjectBindingsRequest
	GetTenantId() *string
}

type ReplaceObjectBindingsRequest struct {
	// 新的对象绑定列表（全量替换；传空列表表示清空所有绑定）
	//
	// This parameter is required.
	ObjectBindings []*ReplaceObjectBindingsRequestObjectBindings `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
	// 数据源 ID（租户内唯一）
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ReplaceObjectBindingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceObjectBindingsRequest) GoString() string {
	return s.String()
}

func (s *ReplaceObjectBindingsRequest) GetObjectBindings() []*ReplaceObjectBindingsRequestObjectBindings {
	return s.ObjectBindings
}

func (s *ReplaceObjectBindingsRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ReplaceObjectBindingsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReplaceObjectBindingsRequest) SetObjectBindings(v []*ReplaceObjectBindingsRequestObjectBindings) *ReplaceObjectBindingsRequest {
	s.ObjectBindings = v
	return s
}

func (s *ReplaceObjectBindingsRequest) SetSourceId(v string) *ReplaceObjectBindingsRequest {
	s.SourceId = &v
	return s
}

func (s *ReplaceObjectBindingsRequest) SetTenantId(v string) *ReplaceObjectBindingsRequest {
	s.TenantId = &v
	return s
}

func (s *ReplaceObjectBindingsRequest) Validate() error {
	if s.ObjectBindings != nil {
		for _, item := range s.ObjectBindings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReplaceObjectBindingsRequestObjectBindings struct {
	// 绑定对象归属的语义图谱名（object_id 在该 graph 下唯一，必填）
	//
	// example:
	//
	// string_value
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// 绑定对象 ID
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// 绑定对象类型（如 customer / project）
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s ReplaceObjectBindingsRequestObjectBindings) String() string {
	return dara.Prettify(s)
}

func (s ReplaceObjectBindingsRequestObjectBindings) GoString() string {
	return s.String()
}

func (s *ReplaceObjectBindingsRequestObjectBindings) GetGraphName() *string {
	return s.GraphName
}

func (s *ReplaceObjectBindingsRequestObjectBindings) GetObjectId() *string {
	return s.ObjectId
}

func (s *ReplaceObjectBindingsRequestObjectBindings) GetObjectType() *string {
	return s.ObjectType
}

func (s *ReplaceObjectBindingsRequestObjectBindings) SetGraphName(v string) *ReplaceObjectBindingsRequestObjectBindings {
	s.GraphName = &v
	return s
}

func (s *ReplaceObjectBindingsRequestObjectBindings) SetObjectId(v string) *ReplaceObjectBindingsRequestObjectBindings {
	s.ObjectId = &v
	return s
}

func (s *ReplaceObjectBindingsRequestObjectBindings) SetObjectType(v string) *ReplaceObjectBindingsRequestObjectBindings {
	s.ObjectType = &v
	return s
}

func (s *ReplaceObjectBindingsRequestObjectBindings) Validate() error {
	return dara.Validate(s)
}
