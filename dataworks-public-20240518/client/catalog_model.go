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
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The creation time. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1722073854000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The catalog ID. Currently, only the DLF and StarRocks types are supported. For details, see [description of concepts related to metadata entities.](https://help.aliyun.com/document_detail/2880092.html)
	//
	// 	- For the DLF type, the format is `dlf-catalog::catalog_id`.
	//
	// 	- For the starrocks type, the format is `starrocks-catalog:(instance_id|encoded_jdbc_url):catalog_name`.
	//
	// > \\
	//
	// `catalog_id`: The the DLF catalog ID.\\
	//
	// `instance_id`: The instance ID, required if the data source is registered in instance mode.\\
	//
	// `encoded_jdbc_url`: The URL-encoded JDBC connection string, required if the data source is registered using a connection string.\\
	//
	// `catalog_name`: The StarRocks catalog name.
	//
	// example:
	//
	// dlf-catalog:123456XXX:test_catalog
	//
	// starrocks-catalog:c-abc123xxx:default_catalog
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The modification time. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1722073854000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The catalog name.
	//
	// example:
	//
	// default_catalog
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parent entity ID.
	//
	// 	- For the DLF type, the format of `ParentMetaEntityId` is `dlf`.
	//
	// 	- For the StarRocks type, the format of `ParentMetaEntityId` is `starrocks:(instance_id|encoded_jdbc_url)`.
	//
	// > \\
	//
	// `instance_id`: The instance ID, required when the data source is registered in instance mode.\\
	//
	// `encoded_jdbc_url`: The URL-encoded JDBC connection string, required if the data source is registered via a connection string.
	//
	// example:
	//
	// dlf
	//
	// starrocks:c-abc123xxx
	ParentMetaEntityId *string `json:"ParentMetaEntityId,omitempty" xml:"ParentMetaEntityId,omitempty"`
	// The catalog type. The value of this parameter varies based on the type of the metadata crawler.
	//
	// example:
	//
	// Internal
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
