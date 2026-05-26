// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *UpdateConfigResponseBody
	GetCreatedAt() *string
	SetKey(v string) *UpdateConfigResponseBody
	GetKey() *string
	SetUpdatedAt(v string) *UpdateConfigResponseBody
	GetUpdatedAt() *string
	SetValue(v string) *UpdateConfigResponseBody
	GetValue() *string
}

type UpdateConfigResponseBody struct {
	// 创建时间
	//
	// example:
	//
	// 2024-01-15T10:30:00Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// 配置键名
	//
	// example:
	//
	// llm_gateway.route_policy
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 更新时间
	//
	// example:
	//
	// 2024-01-15T11:00:00Z
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	// 配置值
	//
	// example:
	//
	// {"model": "gpt-4-turbo"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateConfigResponseBody) GetKey() *string {
	return s.Key
}

func (s *UpdateConfigResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateConfigResponseBody) GetValue() *string {
	return s.Value
}

func (s *UpdateConfigResponseBody) SetCreatedAt(v string) *UpdateConfigResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *UpdateConfigResponseBody) SetKey(v string) *UpdateConfigResponseBody {
	s.Key = &v
	return s
}

func (s *UpdateConfigResponseBody) SetUpdatedAt(v string) *UpdateConfigResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateConfigResponseBody) SetValue(v string) *UpdateConfigResponseBody {
	s.Value = &v
	return s
}

func (s *UpdateConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
