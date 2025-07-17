// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableBusinessMetadataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UpdateTableBusinessMetadataRequest
	GetId() *string
	SetReadme(v string) *UpdateTableBusinessMetadataRequest
	GetReadme() *string
}

type UpdateTableBusinessMetadataRequest struct {
	// The data table ID. You can call the ListTables operation to query the ID.
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
	// ## introduction
	Readme *string `json:"Readme,omitempty" xml:"Readme,omitempty"`
}

func (s UpdateTableBusinessMetadataRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableBusinessMetadataRequest) GoString() string {
	return s.String()
}

func (s *UpdateTableBusinessMetadataRequest) GetId() *string {
	return s.Id
}

func (s *UpdateTableBusinessMetadataRequest) GetReadme() *string {
	return s.Readme
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
