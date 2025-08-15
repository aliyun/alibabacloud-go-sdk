// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdvancedQueryTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAdvancedQueryTemplateResponseBody
	GetRequestId() *string
	SetSimpleQuery(v string) *UpdateAdvancedQueryTemplateResponseBody
	GetSimpleQuery() *string
	SetTemplateId(v string) *UpdateAdvancedQueryTemplateResponseBody
	GetTemplateId() *string
	SetTemplateName(v string) *UpdateAdvancedQueryTemplateResponseBody
	GetTemplateName() *string
	SetTemplateSql(v string) *UpdateAdvancedQueryTemplateResponseBody
	GetTemplateSql() *string
}

type UpdateAdvancedQueryTemplateResponseBody struct {
	// example:
	//
	// 145318BE-DEE1-4C57-AA7C-5BE7D34A6AE0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	SimpleQuery *string `json:"SimpleQuery,omitempty" xml:"SimpleQuery,omitempty"`
	// example:
	//
	// utpl-QNL3dpYkQcyjZxrIQCciqQ
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// event.userIdentity.type: root-account AND event.userIdentity.accessKeyId: *
	TemplateSql *string `json:"TemplateSql,omitempty" xml:"TemplateSql,omitempty"`
}

func (s UpdateAdvancedQueryTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdvancedQueryTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAdvancedQueryTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAdvancedQueryTemplateResponseBody) GetSimpleQuery() *string {
	return s.SimpleQuery
}

func (s *UpdateAdvancedQueryTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateAdvancedQueryTemplateResponseBody) GetTemplateName() *string {
	return s.TemplateName
}

func (s *UpdateAdvancedQueryTemplateResponseBody) GetTemplateSql() *string {
	return s.TemplateSql
}

func (s *UpdateAdvancedQueryTemplateResponseBody) SetRequestId(v string) *UpdateAdvancedQueryTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAdvancedQueryTemplateResponseBody) SetSimpleQuery(v string) *UpdateAdvancedQueryTemplateResponseBody {
	s.SimpleQuery = &v
	return s
}

func (s *UpdateAdvancedQueryTemplateResponseBody) SetTemplateId(v string) *UpdateAdvancedQueryTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *UpdateAdvancedQueryTemplateResponseBody) SetTemplateName(v string) *UpdateAdvancedQueryTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *UpdateAdvancedQueryTemplateResponseBody) SetTemplateSql(v string) *UpdateAdvancedQueryTemplateResponseBody {
	s.TemplateSql = &v
	return s
}

func (s *UpdateAdvancedQueryTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
