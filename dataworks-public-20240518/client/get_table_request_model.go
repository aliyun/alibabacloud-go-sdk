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
	// The table ID. For more information, see [Concepts related to metadata entities](https://help.aliyun.com/document_detail/2880092.html).
	//
	// The common format of this parameter is `${Entity type}:${Instance ID or escaped URL}:${Catalog identifier}:${Database name}:${Schema name}:${Table name}`. If a level does not exist, specify an empty string as a placeholder.
	//
	// >  For MaxCompute and DLF data sources, specify an empty string at the Instance ID level.
	//
	// >  For StarRocks data sources, specify a catalog name at the Catalog identifier level. For DLF data sources, specify a catalog ID at the Catalog identifier level. Other types of data sources do not support the Catalog identifier level. You can specify an empty string as a placeholder.
	//
	// >  For MaxCompute data sources, specify a MaxCompute project name at the Database name level. If the three-layer model is enabled for your MaxCompute project, you must specify a schema name at the Schema name level. Otherwise, you can specify an empty string as a placeholder.
	//
	// You can configure this parameter in one of the following formats based on your data source type:
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
	// `instance_id`: the ID of an instance. If the related data source is added to DataWorks in Alibaba Cloud instance mode, you must configure this parameter.\\
	//
	// `encoded_jdbc_url`: the JDBC connection string that is URL-encoded. If the related data source is added to DataWorks in connection string mode, you must configure this parameter.\\
	//
	// `catalog_id`: the ID of a DLF catalog.\\
	//
	// `project_name`: the name of a MaxCompute project.\\
	//
	// `database_name`: the name of a database.\\
	//
	// `schema_name`: the name of a schema. For a MaxCompute table, if the three-layer model is enabled for the MaxCompute project to which the table belongs, you must configure this parameter. Otherwise, you can specify an empty string for schema_name as a placeholder.\\
	//
	// `table_name`: the name of a table.
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
