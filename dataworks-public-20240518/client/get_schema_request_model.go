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
	//
	//
	//
	// The format is `${EntityType}:${Instance ID or escaped URL}:${Catalog ID}:${Database name}:${Schema name}`. Use empty strings as placeholders for missing levels.
	//
	//
	//
	//
	// > For the MaxCompute type, use an empty string as the placeholder for the instance ID level. The database name is the MaxCompute project name, and the project must have the three-level model enabled.
	//
	//
	//
	//
	// Examples:
	//
	//
	//
	//
	// `maxcompute-schema:::project_name:schema_name` (The three-level model is enabled for the MaxCompute project.)
	//
	//
	//
	//
	// `holo-schema:instance_id::database_name:schema_name`
	//
	//
	//
	//
	// > &lt;br&gt;`instance_id`: The Hologres instance ID&lt;br&gt;
	//
	// > . `database_name`: The database name&lt;br&gt;
	//
	// > . `project_name`: The MaxCompute project name&lt;br&gt;
	//
	// > . `schema_name`: The schema name.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-schema:::project_name:schema_name
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
