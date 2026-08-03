// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAdvancedQueryTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSimpleQuery(v bool) *CreateAdvancedQueryTemplateRequest
	GetSimpleQuery() *bool
	SetTemplateName(v string) *CreateAdvancedQueryTemplateRequest
	GetTemplateName() *string
	SetTemplateSql(v string) *CreateAdvancedQueryTemplateRequest
	GetTemplateSql() *string
}

type CreateAdvancedQueryTemplateRequest struct {
	// Specifies whether to enable the simple query mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	SimpleQuery *bool `json:"SimpleQuery,omitempty" xml:"SimpleQuery,omitempty"`
	// The name of the template, which can contain a maximum of 64 characters. Uniqueness is not required.
	//
	// example:
	//
	// test1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The query statement of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// event.errorCode: 	- AND event.userIdentity.accessKeyId: *
	TemplateSql *string `json:"TemplateSql,omitempty" xml:"TemplateSql,omitempty"`
}

func (s CreateAdvancedQueryTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAdvancedQueryTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateAdvancedQueryTemplateRequest) GetSimpleQuery() *bool {
	return s.SimpleQuery
}

func (s *CreateAdvancedQueryTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateAdvancedQueryTemplateRequest) GetTemplateSql() *string {
	return s.TemplateSql
}

func (s *CreateAdvancedQueryTemplateRequest) SetSimpleQuery(v bool) *CreateAdvancedQueryTemplateRequest {
	s.SimpleQuery = &v
	return s
}

func (s *CreateAdvancedQueryTemplateRequest) SetTemplateName(v string) *CreateAdvancedQueryTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateAdvancedQueryTemplateRequest) SetTemplateSql(v string) *CreateAdvancedQueryTemplateRequest {
	s.TemplateSql = &v
	return s
}

func (s *CreateAdvancedQueryTemplateRequest) Validate() error {
	return dara.Validate(s)
}
