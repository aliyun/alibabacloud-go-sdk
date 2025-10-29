// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetSchemaRequest
	GetId() *string
}

type GetSchemaRequest struct {
	// The ID. You can refer to the ListSchemas operation and [Concepts related to metadata entities](https://help.aliyun.com/document_detail/2880092.html).
	//
	// The format is `${EntityType}:${Instance ID or escaped URL}:${Catalog ID}:${Database name}:${Schema name}</code>`. Use empty strings as placeholders for missing levels.
	//
	// >  For the MaxCompute type, use an empty string as the placeholder for the instance ID level. The database name is the MaxCompute project name, and the project must have the three-level model enabled.
	//
	// Examples:
	//
	// `maxcompute-schema:::project_name:schema_name` (The three-level model is enabled for the MaxCompute project.)
	//
	// `holo-schema:instance_id::database_name:schema_name`
	//
	// > \\
	//
	// `instance_id`: The Hologres instance ID\\
	//
	// . `database_name`: The database name\\
	//
	// . `database_name`: The MaxCompute project name\\
	//
	// . `schema_name`: The schema name.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-schema:123456XXX::test_project:default
	//
	// holo-schema:h-abc123xxx::test_db:test_schema
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSchemaRequest) GoString() string {
	return s.String()
}

func (s *GetSchemaRequest) GetId() *string {
	return s.Id
}

func (s *GetSchemaRequest) SetId(v string) *GetSchemaRequest {
	s.Id = &v
	return s
}

func (s *GetSchemaRequest) Validate() error {
	return dara.Validate(s)
}
