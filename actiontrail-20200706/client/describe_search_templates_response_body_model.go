// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSearchTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSearchTemplatesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSearchTemplatesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSearchTemplatesResponseBody
	GetRequestId() *string
	SetTemplateList(v []*DescribeSearchTemplatesResponseBodyTemplateList) *DescribeSearchTemplatesResponseBody
	GetTemplateList() []*DescribeSearchTemplatesResponseBodyTemplateList
}

type DescribeSearchTemplatesResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 787DD24A-E322-5C0D-A730-057FE62B****
	RequestId    *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateList []*DescribeSearchTemplatesResponseBodyTemplateList `json:"TemplateList,omitempty" xml:"TemplateList,omitempty" type:"Repeated"`
}

func (s DescribeSearchTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSearchTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSearchTemplatesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSearchTemplatesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSearchTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSearchTemplatesResponseBody) GetTemplateList() []*DescribeSearchTemplatesResponseBodyTemplateList {
	return s.TemplateList
}

func (s *DescribeSearchTemplatesResponseBody) SetPageNumber(v int32) *DescribeSearchTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSearchTemplatesResponseBody) SetPageSize(v int32) *DescribeSearchTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSearchTemplatesResponseBody) SetRequestId(v string) *DescribeSearchTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSearchTemplatesResponseBody) SetTemplateList(v []*DescribeSearchTemplatesResponseBodyTemplateList) *DescribeSearchTemplatesResponseBody {
	s.TemplateList = v
	return s
}

func (s *DescribeSearchTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSearchTemplatesResponseBodyTemplateList struct {
	// example:
	//
	// []
	Charts      *string `json:"Charts,omitempty" xml:"Charts,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Params      *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// example:
	//
	// sc-lpYrjKouRfy3MK-wteJW_Q
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// select \\"event.userIdentity.accountId\\" as account_id, count(1) as cnt group by account_id limit 1000
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// example:
	//
	// tpl-wCZAFWx3Spq6CO9Ymp****
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// identity.rootLogin
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// example:
	//
	// audit
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSearchTemplatesResponseBodyTemplateList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSearchTemplatesResponseBodyTemplateList) GoString() string {
	return s.String()
}

func (s *DescribeSearchTemplatesResponseBodyTemplateList) GetCharts() *string {
	return s.Charts
}

func (s *DescribeSearchTemplatesResponseBodyTemplateList) GetDescription() *string {
	return s.Description
}

func (s *DescribeSearchTemplatesResponseBodyTemplateList) GetParams() *string {
	return s.Params
}

func (s *DescribeSearchTemplatesResponseBodyTemplateList) GetSceneId() *string {
	return s.SceneId
}

func (s *DescribeSearchTemplatesResponseBodyTemplateList) GetSql() *string {
	return s.Sql
}

func (s *DescribeSearchTemplatesResponseBodyTemplateList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeSearchTemplatesResponseBodyTemplateList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeSearchTemplatesResponseBodyTemplateList) GetToken() *string {
	return s.Token
}

func (s *DescribeSearchTemplatesResponseBodyTemplateList) GetType() *string {
	return s.Type
}

func (s *DescribeSearchTemplatesResponseBodyTemplateList) SetCharts(v string) *DescribeSearchTemplatesResponseBodyTemplateList {
	s.Charts = &v
	return s
}

func (s *DescribeSearchTemplatesResponseBodyTemplateList) SetDescription(v string) *DescribeSearchTemplatesResponseBodyTemplateList {
	s.Description = &v
	return s
}

func (s *DescribeSearchTemplatesResponseBodyTemplateList) SetParams(v string) *DescribeSearchTemplatesResponseBodyTemplateList {
	s.Params = &v
	return s
}

func (s *DescribeSearchTemplatesResponseBodyTemplateList) SetSceneId(v string) *DescribeSearchTemplatesResponseBodyTemplateList {
	s.SceneId = &v
	return s
}

func (s *DescribeSearchTemplatesResponseBodyTemplateList) SetSql(v string) *DescribeSearchTemplatesResponseBodyTemplateList {
	s.Sql = &v
	return s
}

func (s *DescribeSearchTemplatesResponseBodyTemplateList) SetTemplateId(v string) *DescribeSearchTemplatesResponseBodyTemplateList {
	s.TemplateId = &v
	return s
}

func (s *DescribeSearchTemplatesResponseBodyTemplateList) SetTemplateName(v string) *DescribeSearchTemplatesResponseBodyTemplateList {
	s.TemplateName = &v
	return s
}

func (s *DescribeSearchTemplatesResponseBodyTemplateList) SetToken(v string) *DescribeSearchTemplatesResponseBodyTemplateList {
	s.Token = &v
	return s
}

func (s *DescribeSearchTemplatesResponseBodyTemplateList) SetType(v string) *DescribeSearchTemplatesResponseBodyTemplateList {
	s.Type = &v
	return s
}

func (s *DescribeSearchTemplatesResponseBodyTemplateList) Validate() error {
	return dara.Validate(s)
}
