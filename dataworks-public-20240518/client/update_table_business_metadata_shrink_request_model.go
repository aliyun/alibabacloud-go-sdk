// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableBusinessMetadataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomAttributesShrink(v string) *UpdateTableBusinessMetadataShrinkRequest
	GetCustomAttributesShrink() *string
	SetId(v string) *UpdateTableBusinessMetadataShrinkRequest
	GetId() *string
	SetReadme(v string) *UpdateTableBusinessMetadataShrinkRequest
	GetReadme() *string
}

type UpdateTableBusinessMetadataShrinkRequest struct {
	// example:
	//
	// {"biz_owner":["张三"]}
	CustomAttributesShrink *string `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty"`
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

func (s UpdateTableBusinessMetadataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableBusinessMetadataShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTableBusinessMetadataShrinkRequest) GetCustomAttributesShrink() *string {
	return s.CustomAttributesShrink
}

func (s *UpdateTableBusinessMetadataShrinkRequest) GetId() *string {
	return s.Id
}

func (s *UpdateTableBusinessMetadataShrinkRequest) GetReadme() *string {
	return s.Readme
}

func (s *UpdateTableBusinessMetadataShrinkRequest) SetCustomAttributesShrink(v string) *UpdateTableBusinessMetadataShrinkRequest {
	s.CustomAttributesShrink = &v
	return s
}

func (s *UpdateTableBusinessMetadataShrinkRequest) SetId(v string) *UpdateTableBusinessMetadataShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateTableBusinessMetadataShrinkRequest) SetReadme(v string) *UpdateTableBusinessMetadataShrinkRequest {
	s.Readme = &v
	return s
}

func (s *UpdateTableBusinessMetadataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
