// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdvancedQueryTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAdvancedQueryTemplateResponseBody
	GetRequestId() *string
	SetSimpleQuery(v bool) *GetAdvancedQueryTemplateResponseBody
	GetSimpleQuery() *bool
	SetTemplateId(v string) *GetAdvancedQueryTemplateResponseBody
	GetTemplateId() *string
	SetTemplateName(v string) *GetAdvancedQueryTemplateResponseBody
	GetTemplateName() *string
	SetTemplateSql(v string) *GetAdvancedQueryTemplateResponseBody
	GetTemplateSql() *string
}

type GetAdvancedQueryTemplateResponseBody struct {
	// example:
	//
	// 32110C73-0004-5141-9DA7-4B8045C8173A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	SimpleQuery *bool `json:"SimpleQuery,omitempty" xml:"SimpleQuery,omitempty"`
	// example:
	//
	// utpl-N9fpjnFBSWauSXhVNP3erw
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// event.userIdentity.type: root-account AND event.userIdentity.accessKeyId: *
	TemplateSql *string `json:"TemplateSql,omitempty" xml:"TemplateSql,omitempty"`
}

func (s GetAdvancedQueryTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAdvancedQueryTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetAdvancedQueryTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAdvancedQueryTemplateResponseBody) GetSimpleQuery() *bool {
	return s.SimpleQuery
}

func (s *GetAdvancedQueryTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetAdvancedQueryTemplateResponseBody) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetAdvancedQueryTemplateResponseBody) GetTemplateSql() *string {
	return s.TemplateSql
}

func (s *GetAdvancedQueryTemplateResponseBody) SetRequestId(v string) *GetAdvancedQueryTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAdvancedQueryTemplateResponseBody) SetSimpleQuery(v bool) *GetAdvancedQueryTemplateResponseBody {
	s.SimpleQuery = &v
	return s
}

func (s *GetAdvancedQueryTemplateResponseBody) SetTemplateId(v string) *GetAdvancedQueryTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetAdvancedQueryTemplateResponseBody) SetTemplateName(v string) *GetAdvancedQueryTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetAdvancedQueryTemplateResponseBody) SetTemplateSql(v string) *GetAdvancedQueryTemplateResponseBody {
	s.TemplateSql = &v
	return s
}

func (s *GetAdvancedQueryTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
