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
	SetIncludeExtendedProperties(v bool) *GetTableRequest
	GetIncludeExtendedProperties() *bool
}

type GetTableRequest struct {
	// The ID of the table. You can obtain this value from the response of the ListTables operation. For more information, see [Metadata entity concepts](https://help.aliyun.com/document_detail/2880092.html).
	//
	//
	// The format is `${EntityType}:${InstanceID or encoded URL}:${DataCatalogIdentifier}:${DatabaseName}:${SchemaName}:${TableName}`. Use an empty string as a placeholder for levels that do not exist.
	//
	// > For MaxCompute and DLF types, use an empty string as a placeholder for the instance ID.
	//
	// > For StarRocks, the data catalog identifier is the catalog name. For DLF, the data catalog identifier is the catalog ID. Other types do not support the catalog level. Use an empty string as a placeholder.
	//
	// > For MaxCompute, the database name is the MaxCompute project name. Projects with the three-layer model enabled require a schema name. For projects without the three-layer model enabled, use an empty string as a placeholder for the schema name.
	//
	// The following are ID format examples for common types:
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
	// > Where
	//
	// `instance_id`: The instance ID. Required when the data source is registered in instance mode.
	//
	// `encoded_jdbc_url`: The URL-encoded JDBC connection string. Required when the data source is registered by using a connection string.
	//
	// `catalog_id`: The DLF catalog ID.
	//
	// `project_name`: The MaxCompute project name.
	//
	// `database_name`: The database name.
	//
	// `schema_name`: The schema name. For MaxCompute, this is required only when the three-layer model is enabled for the project. Otherwise, use an empty string as a placeholder.
	//
	// `table_name`: The table name.
	//
	// Recommended procedure for obtaining this parameter: First, call ListCrawlers to obtain the MetaEntityId of the metadata crawler. For types that include a data catalog level, such as DLF and StarRocks, call ListCatalogs to obtain the catalog ID. Then, call ListDatabases to obtain the database ID. If necessary, call ListSchemas to obtain the schema ID. Finally, call ListTables to obtain the target table ID, and use the returned table ID as the Id for this operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-table:::project_name:[schema_name]:table_name
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Specifies whether to include business metadata. Default value: false.
	//
	// example:
	//
	// true
	IncludeBusinessMetadata *bool `json:"IncludeBusinessMetadata,omitempty" xml:"IncludeBusinessMetadata,omitempty"`
	// Specifies whether to return extended properties. Set this parameter to `true` to return extended properties, or `false` to not return them.
	IncludeExtendedProperties *bool `json:"IncludeExtendedProperties,omitempty" xml:"IncludeExtendedProperties,omitempty"`
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

func (s *GetTableRequest) GetIncludeExtendedProperties() *bool {
	return s.IncludeExtendedProperties
}

func (s *GetTableRequest) SetId(v string) *GetTableRequest {
	s.Id = &v
	return s
}

func (s *GetTableRequest) SetIncludeBusinessMetadata(v bool) *GetTableRequest {
	s.IncludeBusinessMetadata = &v
	return s
}

func (s *GetTableRequest) SetIncludeExtendedProperties(v bool) *GetTableRequest {
	s.IncludeExtendedProperties = &v
	return s
}

func (s *GetTableRequest) Validate() error {
	return dara.Validate(s)
}
