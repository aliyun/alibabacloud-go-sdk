// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateColumnBusinessMetadataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomAttributesShrink(v string) *UpdateColumnBusinessMetadataShrinkRequest
	GetCustomAttributesShrink() *string
	SetDescription(v string) *UpdateColumnBusinessMetadataShrinkRequest
	GetDescription() *string
	SetId(v string) *UpdateColumnBusinessMetadataShrinkRequest
	GetId() *string
}

type UpdateColumnBusinessMetadataShrinkRequest struct {
	// The custom attributes of the column, specified as key-value pairs. The key is the attribute identifier, and the value is an array that can contain at most one element. An empty array deletes the attribute\\"s value. To avoid overwriting the column\\"s business description, omit the `Description` parameter from the request. An empty object (`{}`) indicates that no custom attributes are updated.
	//
	// example:
	//
	// {"biz_owner":["张三"]}
	CustomAttributesShrink *string `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty"`
	// The business description of the column.
	//
	// example:
	//
	// test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the column. You can obtain this ID from the response of the `ListColumns` operation. For more information, see [Metadata Entity Concepts](https://help.aliyun.com/document_detail/2880092.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-column:::project_name:[schema_name]:table_name:column_name
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateColumnBusinessMetadataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateColumnBusinessMetadataShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateColumnBusinessMetadataShrinkRequest) GetCustomAttributesShrink() *string {
	return s.CustomAttributesShrink
}

func (s *UpdateColumnBusinessMetadataShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateColumnBusinessMetadataShrinkRequest) GetId() *string {
	return s.Id
}

func (s *UpdateColumnBusinessMetadataShrinkRequest) SetCustomAttributesShrink(v string) *UpdateColumnBusinessMetadataShrinkRequest {
	s.CustomAttributesShrink = &v
	return s
}

func (s *UpdateColumnBusinessMetadataShrinkRequest) SetDescription(v string) *UpdateColumnBusinessMetadataShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateColumnBusinessMetadataShrinkRequest) SetId(v string) *UpdateColumnBusinessMetadataShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateColumnBusinessMetadataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
