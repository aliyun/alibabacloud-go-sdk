// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *CreateTableRequest
	GetCatalog() *string
	SetClientToken(v string) *CreateTableRequest
	GetClientToken() *string
	SetColumns(v []*CreateTableRequestColumns) *CreateTableRequest
	GetColumns() []*CreateTableRequestColumns
	SetComment(v string) *CreateTableRequest
	GetComment() *string
	SetName(v string) *CreateTableRequest
	GetName() *string
	SetNamespace(v string) *CreateTableRequest
	GetNamespace() *string
	SetRetentionPolicy(v *CreateTableRequestRetentionPolicy) *CreateTableRequest
	GetRetentionPolicy() *CreateTableRequestRetentionPolicy
}

type CreateTableRequest struct {
	// 表所属的数据目录名称。可通过 ListCatalogs 获取已有目录列表
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// 用于保证请求幂等性的Token，防止因网络重试导致重复创建。建议使用 UUID
	//
	// example:
	//
	// 1e9b8f60-3a2c-4d7e-9f1b-8c3d5e7a2b4f
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 表的列定义（JSON 数组）。每列包含 Name（列名，必填）、Type（数据类型，必填，如 STRING、INT32、INT64、FLOAT、DOUBLE、BOOLEAN、TIMESTAMP）、Comment（列备注，选填）
	//
	// example:
	//
	// [{"Name":"id","Type":"bigint","Comment":"主键"}]
	Columns []*CreateTableRequestColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// 表的备注描述信息，无格式限制
	//
	// example:
	//
	// 测试事件表
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// 事件表名称。以字母或数字开头，支持字母、数字、下划线和短横线，长度1~127。在同一命名空间下唯一
	//
	// This parameter is required.
	//
	// example:
	//
	// my_table
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 表所属的命名空间名称。可通过 ListNamespaces 获取已有命名空间列表
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 数据保留策略（JSON 对象）。包含 HotTTL（热数据保留天数，高性能查询）和 ColdTTL（冷数据保留天数，低成本存储）。不传则使用系统默认值
	//
	// example:
	//
	// {"HotTTL":7,"ColdTTL":30}
	RetentionPolicy *CreateTableRequestRetentionPolicy `json:"RetentionPolicy,omitempty" xml:"RetentionPolicy,omitempty" type:"Struct"`
}

func (s CreateTableRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTableRequest) GoString() string {
	return s.String()
}

func (s *CreateTableRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *CreateTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTableRequest) GetColumns() []*CreateTableRequestColumns {
	return s.Columns
}

func (s *CreateTableRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateTableRequest) GetName() *string {
	return s.Name
}

func (s *CreateTableRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateTableRequest) GetRetentionPolicy() *CreateTableRequestRetentionPolicy {
	return s.RetentionPolicy
}

func (s *CreateTableRequest) SetCatalog(v string) *CreateTableRequest {
	s.Catalog = &v
	return s
}

func (s *CreateTableRequest) SetClientToken(v string) *CreateTableRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTableRequest) SetColumns(v []*CreateTableRequestColumns) *CreateTableRequest {
	s.Columns = v
	return s
}

func (s *CreateTableRequest) SetComment(v string) *CreateTableRequest {
	s.Comment = &v
	return s
}

func (s *CreateTableRequest) SetName(v string) *CreateTableRequest {
	s.Name = &v
	return s
}

func (s *CreateTableRequest) SetNamespace(v string) *CreateTableRequest {
	s.Namespace = &v
	return s
}

func (s *CreateTableRequest) SetRetentionPolicy(v *CreateTableRequestRetentionPolicy) *CreateTableRequest {
	s.RetentionPolicy = v
	return s
}

func (s *CreateTableRequest) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RetentionPolicy != nil {
		if err := s.RetentionPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTableRequestColumns struct {
	// example:
	//
	// Isp
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// kafka-default-agent-alikafka_pre-cn-28t3sfzno003
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ehpc_cluster
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateTableRequestColumns) String() string {
	return dara.Prettify(s)
}

func (s CreateTableRequestColumns) GoString() string {
	return s.String()
}

func (s *CreateTableRequestColumns) GetComment() *string {
	return s.Comment
}

func (s *CreateTableRequestColumns) GetName() *string {
	return s.Name
}

func (s *CreateTableRequestColumns) GetType() *string {
	return s.Type
}

func (s *CreateTableRequestColumns) SetComment(v string) *CreateTableRequestColumns {
	s.Comment = &v
	return s
}

func (s *CreateTableRequestColumns) SetName(v string) *CreateTableRequestColumns {
	s.Name = &v
	return s
}

func (s *CreateTableRequestColumns) SetType(v string) *CreateTableRequestColumns {
	s.Type = &v
	return s
}

func (s *CreateTableRequestColumns) Validate() error {
	return dara.Validate(s)
}

type CreateTableRequestRetentionPolicy struct {
	// example:
	//
	// 30
	ColdTTL *int32 `json:"ColdTTL,omitempty" xml:"ColdTTL,omitempty"`
	// example:
	//
	// 30
	HotTTL *int32 `json:"HotTTL,omitempty" xml:"HotTTL,omitempty"`
}

func (s CreateTableRequestRetentionPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateTableRequestRetentionPolicy) GoString() string {
	return s.String()
}

func (s *CreateTableRequestRetentionPolicy) GetColdTTL() *int32 {
	return s.ColdTTL
}

func (s *CreateTableRequestRetentionPolicy) GetHotTTL() *int32 {
	return s.HotTTL
}

func (s *CreateTableRequestRetentionPolicy) SetColdTTL(v int32) *CreateTableRequestRetentionPolicy {
	s.ColdTTL = &v
	return s
}

func (s *CreateTableRequestRetentionPolicy) SetHotTTL(v int32) *CreateTableRequestRetentionPolicy {
	s.HotTTL = &v
	return s
}

func (s *CreateTableRequestRetentionPolicy) Validate() error {
	return dara.Validate(s)
}
