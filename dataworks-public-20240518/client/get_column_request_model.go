// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetColumnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetColumnRequest
	GetId() *string
}

type GetColumnRequest struct {
	// The ID. You can refer to the response of the ListColumns operation and the [description of concepts related to metadata entities.](https://help.aliyun.com/document_detail/2880092.html)
	//
	// The format: `${EntityType}:${Instance ID or escaped URL}:${Catalog identifier}:${Database name}:${Schema name}:${Table name}:${Column name}`. Use empty strings as placeholders for levels that do not exist.
	//
	// >  For the MaxCompute and DLF types, use an empty string as the placeholder for the instance ID. For MaxCompute, the database name refers to the MaxCompute project name. If the project has schema enabled, you must specify the schema name. Otherwise, use an empty string as the placeholder for the schema name.
	//
	// >  The catalog identifier of the StarRocks is the catalog name, and the catalog identifier of the DLF type is the catalog ID. Other types do not support catalog levels. Use empty strings as placeholders.
	//
	// Examples of common ID formats
	//
	// `maxcompute-column:::project_name:[schema_name]:table_name:column_name`
	//
	// `dlf-column::catalog_id:database_name::table_name:column_name`
	//
	// `hms-column:instance_id::database_name::table_name:column_name`
	//
	// `holo-column:instance_id::database_name:schema_name:table_name:column_name`
	//
	// `mysql-column:(instance_id|encoded_jdbc_url)::database_name::table_name:column_name`
	//
	// > \\
	//
	// `instance_id`: the ID of the instance, which is required when the data source is registered in instance mode.\\
	//
	// `encoded_jdbc_url`: The URL-encoded JDBC connection string, which is required when the data source is registered via a connection string.\\
	//
	// `catalog_id`: The DLF catalog ID.\\
	//
	// `project_name`: The MaxCompute project name.\\
	//
	// `database_name`: The database name.\\
	//
	// `schema_name`: The schema name. For the MaxCompute type, this is required only if the project has enabled schema; otherwise, use an empty string as a placeholder.\\
	//
	// `table_name`: The table name.\\
	//
	// `column_name`: The field name.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-column:11075xxxx::test_project:test_schema:test_table:test_column
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetColumnRequest) String() string {
	return dara.Prettify(s)
}

func (s GetColumnRequest) GoString() string {
	return s.String()
}

func (s *GetColumnRequest) GetId() *string {
	return s.Id
}

func (s *GetColumnRequest) SetId(v string) *GetColumnRequest {
	s.Id = &v
	return s
}

func (s *GetColumnRequest) Validate() error {
	return dara.Validate(s)
}
