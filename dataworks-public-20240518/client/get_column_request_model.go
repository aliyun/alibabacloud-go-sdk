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
	// The ID. You can obtain this value from the response of the ListColumns operation. For more information, see [Metadata entity concepts](https://help.aliyun.com/document_detail/2880092.html).
	//
	// The format is `${EntityType}:${InstanceID or encoded URL}:${DataCatalogIdentifier}:${DatabaseName}:${SchemaName}:${TableName}:${ColumnName}`. Use an empty string as a placeholder for levels that do not exist.
	//
	// > For MaxCompute and DLF types, use an empty string as a placeholder for the instance ID. For MaxCompute, the database name is the MaxCompute project name. Projects with the three-layer model enabled must include the schema name. For projects without the three-layer model, use an empty string as a placeholder for the schema name.
	//
	// > For StarRocks, the data catalog identifier is the catalog name. For DLF, the data catalog identifier is the catalog ID. Other types do not support the catalog level. Use an empty string as a placeholder.
	//
	// The following are ID format examples for common types:
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
	// > Other parameters:
	//
	// `instance_id`: The instance ID. This parameter is required when the data source is registered in instance mode.
	//
	// `encoded_jdbc_url`: The URL-encoded JDBC connection string. This parameter is required when the data source is registered by using a connection string.
	//
	// `catalog_id`: The DLF catalog ID.
	//
	// `project_name`: The MaxCompute project name.
	//
	// `database_name`: The database name.
	//
	// `schema_name`: The schema name. For MaxCompute, this parameter is required only when the three-layer model is enabled for the project. If the three-layer model is not enabled, use an empty string as a placeholder.
	//
	// `table_name`: The table name.
	//
	// `column_name`: The column name.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-column:::project_name:[schema_name]:table_name:column_name
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
