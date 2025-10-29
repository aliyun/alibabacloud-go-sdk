// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetTableRequest
	GetId() *string
	SetIncludeBusinessMetadata(v bool) *GetTableRequest
	GetIncludeBusinessMetadata() *bool
}

type GetTableRequest struct {
	// The ID. You can refer to the response of the ListTables operation and the [concepts related to metadata entities.](https://help.aliyun.com/document_detail/2880092.html)
	//
	// The format: `${EntityType}:${Instance ID or escaped URL}:${Catalog identifier}:${Database name}:${Table name}`. Use empty strings as placeholders for levels that do not exist.
	//
	// >  For the MaxCompute and DLF types, use an empty string as the placeholder for the instance ID.
	//
	// >  The catalog identifier of the StarRocks is the catalog name, and the catalog identifier of the DLF type is the catalog ID. Other types do not support the catalog level. Use an empty string as a placeholder.
	//
	// >  For MaxCompute, the database name refers to the MaxCompute project name. If the project has schema enabled, you must specify the schema name. Otherwise, use an empty string as the placeholder for the schema name.
	//
	// Examples of common ID formats
	//
	// `maxcompute-table:::project_name:[schema_name]:table_name`
	//
	// `dlf-table::catalog_id:database_name::table_name`
	//
	// `hms-table:instance_id::database_name::table_name`
	//
	// `holo-table:instance_id::database_name:schema_name:table_name`
	//
	// `mysql-table:(instance_id|encoded_jdbc_url)::database_name::table_name`
	//
	// > \\
	//
	// `instance_id`: The instance ID, required when the data source is registered in instance mode.\\
	//
	// `encoded_jdbc_url`: The URL-encoded JDBC connection string, which is required when the data source is registered via a connection string.\\
	//
	// `catalog_id`: The DLF catalog ID.\\
	//
	// `project_name`: The MaxCompute project name.\\
	//
	// `database_name`: The database name.\\
	//
	// `schema_name`: The schema name. For the MaxCompute type, this is required only if the project has enabled schema. Otherwise, use an empty string as a placeholder.\\
	//
	// `table_name`: The table name.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-table:123456XXX::test_project::test_tbl
	//
	// dlf-table:123456XXX:test_catalog:test_db::test_tbl
	//
	// hms-table:c-abc123xxx::test_db::test_tbl
	//
	// holo-table:h-abc123xxx::test_db:test_schema:test_tbl
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Specifies whether to include metadata. Default: false.
	//
	// example:
	//
	// true
	IncludeBusinessMetadata *bool `json:"IncludeBusinessMetadata,omitempty" xml:"IncludeBusinessMetadata,omitempty"`
}

func (s GetTableRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableRequest) GoString() string {
	return s.String()
}

func (s *GetTableRequest) GetId() *string {
	return s.Id
}

func (s *GetTableRequest) GetIncludeBusinessMetadata() *bool {
	return s.IncludeBusinessMetadata
}

func (s *GetTableRequest) SetId(v string) *GetTableRequest {
	s.Id = &v
	return s
}

func (s *GetTableRequest) SetIncludeBusinessMetadata(v bool) *GetTableRequest {
	s.IncludeBusinessMetadata = &v
	return s
}

func (s *GetTableRequest) Validate() error {
	return dara.Validate(s)
}
