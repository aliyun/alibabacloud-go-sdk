// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateColumnBusinessMetadataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateColumnBusinessMetadataRequest
	GetDescription() *string
	SetId(v string) *UpdateColumnBusinessMetadataRequest
	GetId() *string
}

type UpdateColumnBusinessMetadataRequest struct {
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

func (s UpdateColumnBusinessMetadataRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateColumnBusinessMetadataRequest) GoString() string {
	return s.String()
}

func (s *UpdateColumnBusinessMetadataRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateColumnBusinessMetadataRequest) GetId() *string {
	return s.Id
}

func (s *UpdateColumnBusinessMetadataRequest) SetDescription(v string) *UpdateColumnBusinessMetadataRequest {
	s.Description = &v
	return s
}

func (s *UpdateColumnBusinessMetadataRequest) SetId(v string) *UpdateColumnBusinessMetadataRequest {
	s.Id = &v
	return s
}

func (s *UpdateColumnBusinessMetadataRequest) Validate() error {
	return dara.Validate(s)
}
