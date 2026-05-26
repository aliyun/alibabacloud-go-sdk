// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListConfigsResponseBodyConfigs) *ListConfigsResponseBody
	GetConfigs() []*ListConfigsResponseBodyConfigs
	SetHasMore(v bool) *ListConfigsResponseBody
	GetHasMore() *bool
	SetName(v string) *ListConfigsResponseBody
	GetName() *string
	SetPage(v int64) *ListConfigsResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListConfigsResponseBody
	GetPageSize() *int64
	SetTotal(v int64) *ListConfigsResponseBody
	GetTotal() *int64
	SetType(v string) *ListConfigsResponseBody
	GetType() *string
}

type ListConfigsResponseBody struct {
	// 配置项列表
	//
	// example:
	//
	// [{"Key": "rate_limit", "Value": "{\\"limit\\": 100}", "CreatedAt": "2024-01-15T10:30:00Z", "UpdatedAt": "2024-01-15T10:30:00Z"}]
	Configs []*ListConfigsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// 是否有更多数据
	//
	// example:
	//
	// true
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// 服务名称
	//
	// example:
	//
	// my-service
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 当前页码
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// 每页数量
	//
	// example:
	//
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总数量
	//
	// example:
	//
	// 150
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// 配置类型
	//
	// example:
	//
	// Service
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigsResponseBody) GetConfigs() []*ListConfigsResponseBodyConfigs {
	return s.Configs
}

func (s *ListConfigsResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListConfigsResponseBody) GetName() *string {
	return s.Name
}

func (s *ListConfigsResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListConfigsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListConfigsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListConfigsResponseBody) GetType() *string {
	return s.Type
}

func (s *ListConfigsResponseBody) SetConfigs(v []*ListConfigsResponseBodyConfigs) *ListConfigsResponseBody {
	s.Configs = v
	return s
}

func (s *ListConfigsResponseBody) SetHasMore(v bool) *ListConfigsResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListConfigsResponseBody) SetName(v string) *ListConfigsResponseBody {
	s.Name = &v
	return s
}

func (s *ListConfigsResponseBody) SetPage(v int64) *ListConfigsResponseBody {
	s.Page = &v
	return s
}

func (s *ListConfigsResponseBody) SetPageSize(v int64) *ListConfigsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListConfigsResponseBody) SetTotal(v int64) *ListConfigsResponseBody {
	s.Total = &v
	return s
}

func (s *ListConfigsResponseBody) SetType(v string) *ListConfigsResponseBody {
	s.Type = &v
	return s
}

func (s *ListConfigsResponseBody) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListConfigsResponseBodyConfigs struct {
	// 创建时间
	//
	// example:
	//
	// 2024-01-01T10:00:00Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// 配置项键名
	//
	// example:
	//
	// rate_limit
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 更新时间
	//
	// example:
	//
	// 2024-01-01T10:00:00Z
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	// 配置值
	//
	// example:
	//
	// 100
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListConfigsResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListConfigsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListConfigsResponseBodyConfigs) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListConfigsResponseBodyConfigs) GetKey() *string {
	return s.Key
}

func (s *ListConfigsResponseBodyConfigs) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListConfigsResponseBodyConfigs) GetValue() *string {
	return s.Value
}

func (s *ListConfigsResponseBodyConfigs) SetCreatedAt(v string) *ListConfigsResponseBodyConfigs {
	s.CreatedAt = &v
	return s
}

func (s *ListConfigsResponseBodyConfigs) SetKey(v string) *ListConfigsResponseBodyConfigs {
	s.Key = &v
	return s
}

func (s *ListConfigsResponseBodyConfigs) SetUpdatedAt(v string) *ListConfigsResponseBodyConfigs {
	s.UpdatedAt = &v
	return s
}

func (s *ListConfigsResponseBodyConfigs) SetValue(v string) *ListConfigsResponseBodyConfigs {
	s.Value = &v
	return s
}

func (s *ListConfigsResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
