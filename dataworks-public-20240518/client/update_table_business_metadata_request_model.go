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
	// The values of custom attributes. The key specifies the identifier of a custom attribute, and the value is an array that can contain at most one item. To delete the value for an attribute, pass an empty array. To update only custom attributes, omit the `Readme` parameter to prevent its existing value from being cleared. To leave the custom attributes unchanged, pass an empty object `{}`.
	//
	// example:
	//
	// {"biz_owner":["张三"]}
	CustomAttributes map[string][]*string `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty"`
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
