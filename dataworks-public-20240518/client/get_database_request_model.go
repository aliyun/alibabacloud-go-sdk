// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetDatabaseRequest
	GetId() *string
}

type GetDatabaseRequest struct {
	// Database entity ID. You can refer to the response of the ListDatabases operation and [Concepts related to metadata entities](https://help.aliyun.com/document_detail/2880092.html).
	//
	// The format: `${EntityType}:${Instance ID or escaped URL}:${Catalog identifier}:${Database name}`. Use empty strings as placeholders for levels that do not exist.
	//
	// >  For StarRocks, the catalog identifier is the catalog name. For DLF, the catalog identifier is the catalog ID. For other types, the catalog-level hierarchy is not supported, and an empty string can be used as a placeholder.
	//
	// Examples of common ID formats
	//
	// 	- `dlf-database::catalog_id:database_name`
	//
	// 	- `holo-database:instance_id::database_name`
	//
	// 	- `mysql-database:(instance_id|encoded_jdbc_url)::database_name`
	//
	// >
	//
	// 	- `catalog_id`: The ID of the DLF catalog.
	//
	// 	- `instance_id`: The instance ID, which is required when the data source is registered in instance mode.
	//
	// 	- `encoded_jdbc_url`: The encoded JDBC connection string, which is required when the data source is registered in connection-string mode.
	//
	// 	- `database_name`: The database name.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql-database:rm-abc123xxx::test_db
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseRequest) GoString() string {
	return s.String()
}

func (s *GetDatabaseRequest) GetId() *string {
	return s.Id
}

func (s *GetDatabaseRequest) SetId(v string) *GetDatabaseRequest {
	s.Id = &v
	return s
}

func (s *GetDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
