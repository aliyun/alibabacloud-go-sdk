// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvancedQueryTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAdvancedQueryTemplateResponseBody
	GetRequestId() *string
	SetTemplatePage(v *DescribeAdvancedQueryTemplateResponseBodyTemplatePage) *DescribeAdvancedQueryTemplateResponseBody
	GetTemplatePage() *DescribeAdvancedQueryTemplateResponseBodyTemplatePage
}

type DescribeAdvancedQueryTemplateResponseBody struct {
	// example:
	//
	// 1EC1FDC7-6D01-559F-852C-30D86E9EEB3F
	RequestId    *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplatePage *DescribeAdvancedQueryTemplateResponseBodyTemplatePage `json:"TemplatePage,omitempty" xml:"TemplatePage,omitempty" type:"Struct"`
}

func (s DescribeAdvancedQueryTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvancedQueryTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdvancedQueryTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAdvancedQueryTemplateResponseBody) GetTemplatePage() *DescribeAdvancedQueryTemplateResponseBodyTemplatePage {
	return s.TemplatePage
}

func (s *DescribeAdvancedQueryTemplateResponseBody) SetRequestId(v string) *DescribeAdvancedQueryTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdvancedQueryTemplateResponseBody) SetTemplatePage(v *DescribeAdvancedQueryTemplateResponseBodyTemplatePage) *DescribeAdvancedQueryTemplateResponseBody {
	s.TemplatePage = v
	return s
}

func (s *DescribeAdvancedQueryTemplateResponseBody) Validate() error {
	if s.TemplatePage != nil {
		if err := s.TemplatePage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAdvancedQueryTemplateResponseBodyTemplatePage struct {
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize     *string                                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TemplateList []*DescribeAdvancedQueryTemplateResponseBodyTemplatePageTemplateList `json:"TemplateList,omitempty" xml:"TemplateList,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAdvancedQueryTemplateResponseBodyTemplatePage) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvancedQueryTemplateResponseBodyTemplatePage) GoString() string {
	return s.String()
}

func (s *DescribeAdvancedQueryTemplateResponseBodyTemplatePage) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeAdvancedQueryTemplateResponseBodyTemplatePage) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAdvancedQueryTemplateResponseBodyTemplatePage) GetTemplateList() []*DescribeAdvancedQueryTemplateResponseBodyTemplatePageTemplateList {
	return s.TemplateList
}

func (s *DescribeAdvancedQueryTemplateResponseBodyTemplatePage) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeAdvancedQueryTemplateResponseBodyTemplatePage) SetPageNumber(v string) *DescribeAdvancedQueryTemplateResponseBodyTemplatePage {
	s.PageNumber = &v
	return s
}

func (s *DescribeAdvancedQueryTemplateResponseBodyTemplatePage) SetPageSize(v string) *DescribeAdvancedQueryTemplateResponseBodyTemplatePage {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvancedQueryTemplateResponseBodyTemplatePage) SetTemplateList(v []*DescribeAdvancedQueryTemplateResponseBodyTemplatePageTemplateList) *DescribeAdvancedQueryTemplateResponseBodyTemplatePage {
	s.TemplateList = v
	return s
}

func (s *DescribeAdvancedQueryTemplateResponseBodyTemplatePage) SetTotal(v int64) *DescribeAdvancedQueryTemplateResponseBodyTemplatePage {
	s.Total = &v
	return s
}

func (s *DescribeAdvancedQueryTemplateResponseBodyTemplatePage) Validate() error {
	if s.TemplateList != nil {
		for _, item := range s.TemplateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAdvancedQueryTemplateResponseBodyTemplatePageTemplateList struct {
	// example:
	//
	// false
	SimpleQuery *bool `json:"SimpleQuery,omitempty" xml:"SimpleQuery,omitempty"`
	// example:
	//
	// utpl-7OaxbyJATDaoLOgZRcV5RQ
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// event.userIdentity.type: root-account AND event.userIdentity.accessKeyId: *
	TemplateSql *string `json:"TemplateSql,omitempty" xml:"TemplateSql,omitempty"`
}

func (s DescribeAdvancedQueryTemplateResponseBodyTemplatePageTemplateList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvancedQueryTemplateResponseBodyTemplatePageTemplateList) GoString() string {
	return s.String()
}

func (s *DescribeAdvancedQueryTemplateResponseBodyTemplatePageTemplateList) GetSimpleQuery() *bool {
	return s.SimpleQuery
}

func (s *DescribeAdvancedQueryTemplateResponseBodyTemplatePageTemplateList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeAdvancedQueryTemplateResponseBodyTemplatePageTemplateList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeAdvancedQueryTemplateResponseBodyTemplatePageTemplateList) GetTemplateSql() *string {
	return s.TemplateSql
}

func (s *DescribeAdvancedQueryTemplateResponseBodyTemplatePageTemplateList) SetSimpleQuery(v bool) *DescribeAdvancedQueryTemplateResponseBodyTemplatePageTemplateList {
	s.SimpleQuery = &v
	return s
}

func (s *DescribeAdvancedQueryTemplateResponseBodyTemplatePageTemplateList) SetTemplateId(v string) *DescribeAdvancedQueryTemplateResponseBodyTemplatePageTemplateList {
	s.TemplateId = &v
	return s
}

func (s *DescribeAdvancedQueryTemplateResponseBodyTemplatePageTemplateList) SetTemplateName(v string) *DescribeAdvancedQueryTemplateResponseBodyTemplatePageTemplateList {
	s.TemplateName = &v
	return s
}

func (s *DescribeAdvancedQueryTemplateResponseBodyTemplatePageTemplateList) SetTemplateSql(v string) *DescribeAdvancedQueryTemplateResponseBodyTemplatePageTemplateList {
	s.TemplateSql = &v
	return s
}

func (s *DescribeAdvancedQueryTemplateResponseBodyTemplatePageTemplateList) Validate() error {
	return dara.Validate(s)
}
