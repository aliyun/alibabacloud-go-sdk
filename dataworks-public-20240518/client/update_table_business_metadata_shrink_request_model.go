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
	// The values of custom attributes. The key specifies the identifier of a custom attribute, and the value is an array that can contain at most one item. To delete the value for an attribute, pass an empty array. To update only custom attributes, omit the `Readme` parameter to prevent its existing value from being cleared. To leave the custom attributes unchanged, pass an empty object `{}`.
	//
	// example:
	//
	// {"biz_owner":["张三"]}
	CustomAttributesShrink *string `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty"`
	// The table ID. For the required format, see the response of the `ListTables` operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-table:::project_name:[schema_name]:table_name
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The Readme of the table, which supports rich text format.
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
