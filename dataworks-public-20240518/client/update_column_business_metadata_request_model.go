// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateColumnBusinessMetadataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomAttributes(v map[string][]*string) *UpdateColumnBusinessMetadataRequest
	GetCustomAttributes() map[string][]*string
	SetDescription(v string) *UpdateColumnBusinessMetadataRequest
	GetDescription() *string
	SetId(v string) *UpdateColumnBusinessMetadataRequest
	GetId() *string
}

type UpdateColumnBusinessMetadataRequest struct {
	// The custom attributes of the column, specified as key-value pairs. The key is the attribute identifier, and the value is an array that can contain at most one element. An empty array deletes the attribute\\"s value. To avoid overwriting the column\\"s business description, omit the `Description` parameter from the request. An empty object (`{}`) indicates that no custom attributes are updated.
	//
	// example:
	//
	// {"biz_owner":["张三"]}
	CustomAttributes map[string][]*string `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty"`
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

func (s UpdateColumnBusinessMetadataRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateColumnBusinessMetadataRequest) GoString() string {
	return s.String()
}

func (s *UpdateColumnBusinessMetadataRequest) GetCustomAttributes() map[string][]*string {
	return s.CustomAttributes
}

func (s *UpdateColumnBusinessMetadataRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateColumnBusinessMetadataRequest) GetId() *string {
	return s.Id
}

func (s *UpdateColumnBusinessMetadataRequest) SetCustomAttributes(v map[string][]*string) *UpdateColumnBusinessMetadataRequest {
	s.CustomAttributes = v
	return s
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
