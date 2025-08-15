// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAdvancedQueryTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAdvancedQueryTemplateResponseBody
	GetRequestId() *string
	SetSimpleQuery(v string) *CreateAdvancedQueryTemplateResponseBody
	GetSimpleQuery() *string
	SetTemplateId(v string) *CreateAdvancedQueryTemplateResponseBody
	GetTemplateId() *string
	SetTemplateName(v string) *CreateAdvancedQueryTemplateResponseBody
	GetTemplateName() *string
	SetTemplateSql(v string) *CreateAdvancedQueryTemplateResponseBody
	GetTemplateSql() *string
}

type CreateAdvancedQueryTemplateResponseBody struct {
	// example:
	//
	// 4ABAEA6E-C740-5CE2-A003-643E551964F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	SimpleQuery *string `json:"SimpleQuery,omitempty" xml:"SimpleQuery,omitempty"`
	// example:
	//
	// x4a0Tw5dQy2J6IRJxf4kng
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// event.eventName: ConsoleSignin AND event.userIdentity.type: root-account
	TemplateSql *string `json:"TemplateSql,omitempty" xml:"TemplateSql,omitempty"`
}

func (s CreateAdvancedQueryTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAdvancedQueryTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAdvancedQueryTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAdvancedQueryTemplateResponseBody) GetSimpleQuery() *string {
	return s.SimpleQuery
}

func (s *CreateAdvancedQueryTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateAdvancedQueryTemplateResponseBody) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateAdvancedQueryTemplateResponseBody) GetTemplateSql() *string {
	return s.TemplateSql
}

func (s *CreateAdvancedQueryTemplateResponseBody) SetRequestId(v string) *CreateAdvancedQueryTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAdvancedQueryTemplateResponseBody) SetSimpleQuery(v string) *CreateAdvancedQueryTemplateResponseBody {
	s.SimpleQuery = &v
	return s
}

func (s *CreateAdvancedQueryTemplateResponseBody) SetTemplateId(v string) *CreateAdvancedQueryTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *CreateAdvancedQueryTemplateResponseBody) SetTemplateName(v string) *CreateAdvancedQueryTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *CreateAdvancedQueryTemplateResponseBody) SetTemplateSql(v string) *CreateAdvancedQueryTemplateResponseBody {
	s.TemplateSql = &v
	return s
}

func (s *CreateAdvancedQueryTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
