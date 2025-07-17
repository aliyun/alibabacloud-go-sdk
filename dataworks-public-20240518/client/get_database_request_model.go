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
	// Database entity ID. You can refer to the response of the ListDatabases operation and [the description of metadata entity concepts.](https://help.aliyun.com/document_detail/2880092.html)
	//
	// The format is `${EntityType}:${Instance ID or encoded URL}:${Catalog identifier}:${Database name}`. Use empty strings as placeholders for non-existent levels.
	//
	// >  For StarRocks, the catalog identifier is the catalog name. For DLF, the catalog identifier is the catalog ID. For other types, catalog hierarchy is not supported, and an empty string can be used as a placeholder.
	//
	// Examples of common ID formats
	//
	// `dlf-database::catalog_id:database_name`
	//
	// `holo-database:instance_id::database_name`
	//
	// `mysql-database:(instance_id|encoded_jdbc_url)::database_name`
	//
	// >  Parameter descriptions\\
	//
	// `catalog_id`: The DLF catalog ID.\\
	//
	// `instance_id`: The instance ID, required for a data source registered in instance mode.\\
	//
	// `encoded_jdbc_url`: The JDBC connection string that has been URL encoded. This parameter is required for the data source registered via a connection string.\\
	//
	// `database_name`: The database name.
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
