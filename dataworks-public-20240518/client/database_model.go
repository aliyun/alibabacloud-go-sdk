// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatabase interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *Database
	GetComment() *string
	SetCreateTime(v int64) *Database
	GetCreateTime() *int64
	SetId(v string) *Database
	GetId() *string
	SetLocationUri(v string) *Database
	GetLocationUri() *string
	SetModifyTime(v int64) *Database
	GetModifyTime() *int64
	SetName(v string) *Database
	GetName() *string
	SetParentMetaEntityId(v string) *Database
	GetParentMetaEntityId() *string
}

type Database struct {
	// The comments.
	//
	// example:
	//
	// test comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The creation time. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1736852168000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The database ID. For more information, see [Concepts related to metadata entities](https://help.aliyun.com/document_detail/2880092.html).
	//
	// The common format of this parameter is `${Entity type}:${Instance ID or escaped URL}:${Catalog identifier}:${Database name}`. If a level does not exist, specify an empty string as a placeholder.
	//
	// >  For StarRocks data sources, specify a catalog name at the Catalog identifier level. For DLF data sources, specify a catalog ID at the Catalog identifier level. Other types of data sources do not support the Catalog identifier level. You can specify an empty string as a placeholder.
	//
	// You can configure this parameter in one of the following formats based on your data source type:
	//
	// `dlf-database::catalog_id:database_name`
	//
	// `holo-database:instance_id::database_name`
	//
	// `mysql-database:(instance_id|encoded_jdbc_url)::database_name`
	//
	// > \\
	//
	// `catalog_id`: the ID of a DLF catalog.\\
	//
	// `instance_id`: the ID of an instance. If the related data source is added to DataWorks in Alibaba Cloud instance mode, you must configure this parameter.\\
	//
	// `encoded_jdbc_url`: the JDBC connection string that is URL-encoded. If the related data source is added to DataWorks in connection string mode, you must configure this parameter.\\
	//
	// `database_name`: the name of a database.
	//
	// example:
	//
	// holo-database:h-xxxx::test_db
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The uniform resource identifier (URI) of the storage location.
	//
	// example:
	//
	// oss://test-bucket/test_db
	LocationUri *string `json:"LocationUri,omitempty" xml:"LocationUri,omitempty"`
	// The update time. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1736852168000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The database name.
	//
	// example:
	//
	// test_db
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parent entity ID. For more information, see [Concepts related to metadata entities](https://help.aliyun.com/document_detail/2880092.html).
	//
	// You can call the ListCrawlerTypes operation to query the parent entity types.
	//
	// 	- If the parent entity is a catalog, you can configure the `ParentMetaEntityId` parameter by referring to the `Catalog` topic.
	//
	// 	- If the parent entity is a metadata crawler, you can configure the `ParentMetaEntityId` parameter in the `${Crawler type}:${Instance ID or escaped URL}` format.
	//
	// You can configure the ParentMetaEntityId parameter in one of the following formats based on your data source type:
	//
	// `dlf-catalog::catalog_id`
	//
	// `holo:instance_id`
	//
	// `mysql:(instance_id|encoded_jdbc_url)`
	//
	// > \\
	//
	// `catalog_id`: the ID of a DLF catalog.\\
	//
	// `instance_id`: the ID of an instance. If the related data source is added to DataWorks in Alibaba Cloud instance mode, you must configure this parameter.\\
	//
	// `encoded_jdbc_url`: the JDBC connection string that is URL-encoded. If the related data source is added to DataWorks in connection string mode, you must configure this parameter.
	//
	// example:
	//
	// holo:h-xxxx
	ParentMetaEntityId *string `json:"ParentMetaEntityId,omitempty" xml:"ParentMetaEntityId,omitempty"`
}

func (s Database) String() string {
	return dara.Prettify(s)
}

func (s Database) GoString() string {
	return s.String()
}

func (s *Database) GetComment() *string {
	return s.Comment
}

func (s *Database) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *Database) GetId() *string {
	return s.Id
}

func (s *Database) GetLocationUri() *string {
	return s.LocationUri
}

func (s *Database) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *Database) GetName() *string {
	return s.Name
}

func (s *Database) GetParentMetaEntityId() *string {
	return s.ParentMetaEntityId
}

func (s *Database) SetComment(v string) *Database {
	s.Comment = &v
	return s
}

func (s *Database) SetCreateTime(v int64) *Database {
	s.CreateTime = &v
	return s
}

func (s *Database) SetId(v string) *Database {
	s.Id = &v
	return s
}

func (s *Database) SetLocationUri(v string) *Database {
	s.LocationUri = &v
	return s
}

func (s *Database) SetModifyTime(v int64) *Database {
	s.ModifyTime = &v
	return s
}

func (s *Database) SetName(v string) *Database {
	s.Name = &v
	return s
}

func (s *Database) SetParentMetaEntityId(v string) *Database {
	s.ParentMetaEntityId = &v
	return s
}

func (s *Database) Validate() error {
	return dara.Validate(s)
}
