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
	// example:
	//
	// {"biz_owner":["张三"]}
	CustomAttributesShrink *string `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty"`
	// The field business description.
	//
	// example:
	//
	// test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The field ID. You can refer to the response from the ListColumns operation. You can also refer to the [Concepts related to metadata entities](https://help.aliyun.com/document_detail/2880092.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-column:11075xxxx::test_project:test_schema:test_table:test_column
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
