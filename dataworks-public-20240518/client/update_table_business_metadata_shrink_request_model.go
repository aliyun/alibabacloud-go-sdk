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
	// The custom attribute values. The key is the custom attribute identifier, and the value contains at most one element. An empty list indicates that the attribute value is deleted. Passing this parameter without Readme prevents the usage description from being cleared. An empty object indicates that custom attributes are not updated.
	//
	// example:
	//
	// {"biz_owner":["张三"]}
	CustomAttributesShrink *string `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty"`
	// The ID of the table. For the format, refer to the response of the ListTables operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-table:::project_name:[schema_name]:table_name
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The usage description. Rich text format is supported.
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
