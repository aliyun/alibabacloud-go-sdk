// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCatalog interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *Catalog
	GetComment() *string
	SetCreateTime(v int64) *Catalog
	GetCreateTime() *int64
	SetId(v string) *Catalog
	GetId() *string
	SetModifyTime(v int64) *Catalog
	GetModifyTime() *int64
	SetName(v string) *Catalog
	GetName() *string
	SetParentMetaEntityId(v string) *Catalog
	GetParentMetaEntityId() *string
	SetType(v string) *Catalog
	GetType() *string
}

type Catalog struct {
	// The comment.
	//
	// example:
	//
	// Default catalog
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The creation time. The value is a millisecond-level timestamp.
	//
	// example:
	//
	// 1725330777000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the data catalog entity. Currently, the DLF and StarRocks types are supported. For more information, see [Concepts of metadata entities](https://help.aliyun.com/document_detail/2880092.html).
	//
	// - For the DLF type, the format is `dlf-catalog::catalog_id`.
	//
	// - For the StarRocks type, the format is `starrocks-catalog:(instance_id|encoded_jdbc_url):catalog_name`.
	//
	// > Where:
	//
	// `catalog_id`: The DLF catalog ID.
	//
	// `instance_id`: The instance ID, which is required when the data source is registered in instance mode.
	//
	// `encoded_jdbc_url`: the URL-encoded JDBC connection string, which is required when the data source is registered by using a connection string.
	//
	// `catalog_name`: the name of the StarRocks catalog.
	//
	// example:
	//
	// dlf-catalog::catalog1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The modification time. The value is a millisecond-level timestamp.
	//
	// example:
	//
	// 1725330777000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name.
	//
	// example:
	//
	// catalog1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the parent entity.
	//
	// - For the DLF type, the format of `ParentMetaEntityId` is `dlf`.
	//
	// - For the StarRocks type, the format of `ParentMetaEntityId` is `starrocks:(instance_id|encoded_jdbc_url)`.
	//
	// > Where:
	//
	// `instance_id`: The instance ID, which is required when the data source is registered in instance mode.
	//
	// `encoded_jdbc_url`: the URL-encoded JDBC connection string, which is required when the data source is registered by using a connection string.
	//
	// example:
	//
	// dlf
	ParentMetaEntityId *string `json:"ParentMetaEntityId,omitempty" xml:"ParentMetaEntityId,omitempty"`
	// The catalog type. The valid values depend on the metadata collector type.
	//
	// example:
	//
	// HIVE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s Catalog) String() string {
	return dara.Prettify(s)
}

func (s Catalog) GoString() string {
	return s.String()
}

func (s *Catalog) GetComment() *string {
	return s.Comment
}

func (s *Catalog) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *Catalog) GetId() *string {
	return s.Id
}

func (s *Catalog) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *Catalog) GetName() *string {
	return s.Name
}

func (s *Catalog) GetParentMetaEntityId() *string {
	return s.ParentMetaEntityId
}

func (s *Catalog) GetType() *string {
	return s.Type
}

func (s *Catalog) SetComment(v string) *Catalog {
	s.Comment = &v
	return s
}

func (s *Catalog) SetCreateTime(v int64) *Catalog {
	s.CreateTime = &v
	return s
}

func (s *Catalog) SetId(v string) *Catalog {
	s.Id = &v
	return s
}

func (s *Catalog) SetModifyTime(v int64) *Catalog {
	s.ModifyTime = &v
	return s
}

func (s *Catalog) SetName(v string) *Catalog {
	s.Name = &v
	return s
}

func (s *Catalog) SetParentMetaEntityId(v string) *Catalog {
	s.ParentMetaEntityId = &v
	return s
}

func (s *Catalog) SetType(v string) *Catalog {
	s.Type = &v
	return s
}

func (s *Catalog) Validate() error {
	return dara.Validate(s)
}
