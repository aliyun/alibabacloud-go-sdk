// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentResourceDetail interface {
	dara.Model
	String() string
	GoString() string
	SetExtraId(v string) *IncidentResourceDetail
	GetExtraId() *string
	SetResourceId(v map[string]interface{}) *IncidentResourceDetail
	GetResourceId() map[string]interface{}
	SetType(v string) *IncidentResourceDetail
	GetType() *string
}

type IncidentResourceDetail struct {
	// 额外的资源标识符
	//
	// example:
	//
	// app-prod-web-01
	ExtraId *string `json:"extraId,omitempty" xml:"extraId,omitempty"`
	// 资源 ID 对象。
	//
	// example:
	//
	// { "type": "ecs", "id": "i-bp1234567890abcdef" }
	ResourceId map[string]interface{} `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// 资源类型，表示该资源属于哪种云产品。
	//
	// example:
	//
	// ecs
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s IncidentResourceDetail) String() string {
	return dara.Prettify(s)
}

func (s IncidentResourceDetail) GoString() string {
	return s.String()
}

func (s *IncidentResourceDetail) GetExtraId() *string {
	return s.ExtraId
}

func (s *IncidentResourceDetail) GetResourceId() map[string]interface{} {
	return s.ResourceId
}

func (s *IncidentResourceDetail) GetType() *string {
	return s.Type
}

func (s *IncidentResourceDetail) SetExtraId(v string) *IncidentResourceDetail {
	s.ExtraId = &v
	return s
}

func (s *IncidentResourceDetail) SetResourceId(v map[string]interface{}) *IncidentResourceDetail {
	s.ResourceId = v
	return s
}

func (s *IncidentResourceDetail) SetType(v string) *IncidentResourceDetail {
	s.Type = &v
	return s
}

func (s *IncidentResourceDetail) Validate() error {
	return dara.Validate(s)
}
