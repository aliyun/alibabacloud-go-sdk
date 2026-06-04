// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableBusinessMetadataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomAttributes(v map[string][]*string) *UpdateTableBusinessMetadataRequest
	GetCustomAttributes() map[string][]*string
	SetId(v string) *UpdateTableBusinessMetadataRequest
	GetId() *string
	SetReadme(v string) *UpdateTableBusinessMetadataRequest
	GetReadme() *string
}

type UpdateTableBusinessMetadataRequest struct {
	// example:
	//
	// {"biz_owner":["张三"]}
	CustomAttributes map[string][]*string `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty"`
	// The table ID. You can refer to the format of the table ID returned by the ListTables operation.
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
	// The usage notes. The rich text format is supported.
	//
	// example:
	//
	// ## introduction
	Readme *string `json:"Readme,omitempty" xml:"Readme,omitempty"`
}

func (s UpdateTableBusinessMetadataRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableBusinessMetadataRequest) GoString() string {
	return s.String()
}

func (s *UpdateTableBusinessMetadataRequest) GetCustomAttributes() map[string][]*string {
	return s.CustomAttributes
}

func (s *UpdateTableBusinessMetadataRequest) GetId() *string {
	return s.Id
}

func (s *UpdateTableBusinessMetadataRequest) GetReadme() *string {
	return s.Readme
}

func (s *UpdateTableBusinessMetadataRequest) SetCustomAttributes(v map[string][]*string) *UpdateTableBusinessMetadataRequest {
	s.CustomAttributes = v
	return s
}

func (s *UpdateTableBusinessMetadataRequest) SetId(v string) *UpdateTableBusinessMetadataRequest {
	s.Id = &v
	return s
}

func (s *UpdateTableBusinessMetadataRequest) SetReadme(v string) *UpdateTableBusinessMetadataRequest {
	s.Readme = &v
	return s
}

func (s *UpdateTableBusinessMetadataRequest) Validate() error {
	return dara.Validate(s)
}
