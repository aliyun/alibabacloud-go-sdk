// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppViewTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAppViewTemplatesResponseBody
	GetRequestId() *string
	SetTemplates(v []*DescribeAppViewTemplatesResponseBodyTemplates) *DescribeAppViewTemplatesResponseBody
	GetTemplates() []*DescribeAppViewTemplatesResponseBodyTemplates
	SetTotalNum(v int64) *DescribeAppViewTemplatesResponseBody
	GetTotalNum() *int64
	SetTotalPage(v int64) *DescribeAppViewTemplatesResponseBody
	GetTotalPage() *int64
}

type DescribeAppViewTemplatesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B0A2FCBC-43A4-428F-BC1D-3F4F85837F76
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates []*DescribeAppViewTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeAppViewTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppViewTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppViewTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppViewTemplatesResponseBody) GetTemplates() []*DescribeAppViewTemplatesResponseBodyTemplates {
	return s.Templates
}

func (s *DescribeAppViewTemplatesResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *DescribeAppViewTemplatesResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *DescribeAppViewTemplatesResponseBody) SetRequestId(v string) *DescribeAppViewTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppViewTemplatesResponseBody) SetTemplates(v []*DescribeAppViewTemplatesResponseBodyTemplates) *DescribeAppViewTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *DescribeAppViewTemplatesResponseBody) SetTotalNum(v int64) *DescribeAppViewTemplatesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeAppViewTemplatesResponseBody) SetTotalPage(v int64) *DescribeAppViewTemplatesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeAppViewTemplatesResponseBody) Validate() error {
	if s.Templates != nil {
		for _, item := range s.Templates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAppViewTemplatesResponseBodyTemplates struct {
	// example:
	//
	// 2020-09-04T06:22:15Z
	CreateTime *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LayoutIds  []*string `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// example:
	//
	// 模版名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Bj6D****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeAppViewTemplatesResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppViewTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeAppViewTemplatesResponseBodyTemplates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAppViewTemplatesResponseBodyTemplates) GetLayoutIds() []*string {
	return s.LayoutIds
}

func (s *DescribeAppViewTemplatesResponseBodyTemplates) GetMediaEncode() *int32 {
	return s.MediaEncode
}

func (s *DescribeAppViewTemplatesResponseBodyTemplates) GetName() *string {
	return s.Name
}

func (s *DescribeAppViewTemplatesResponseBodyTemplates) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeAppViewTemplatesResponseBodyTemplates) SetCreateTime(v string) *DescribeAppViewTemplatesResponseBodyTemplates {
	s.CreateTime = &v
	return s
}

func (s *DescribeAppViewTemplatesResponseBodyTemplates) SetLayoutIds(v []*string) *DescribeAppViewTemplatesResponseBodyTemplates {
	s.LayoutIds = v
	return s
}

func (s *DescribeAppViewTemplatesResponseBodyTemplates) SetMediaEncode(v int32) *DescribeAppViewTemplatesResponseBodyTemplates {
	s.MediaEncode = &v
	return s
}

func (s *DescribeAppViewTemplatesResponseBodyTemplates) SetName(v string) *DescribeAppViewTemplatesResponseBodyTemplates {
	s.Name = &v
	return s
}

func (s *DescribeAppViewTemplatesResponseBodyTemplates) SetTemplateId(v string) *DescribeAppViewTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

func (s *DescribeAppViewTemplatesResponseBodyTemplates) Validate() error {
	return dara.Validate(s)
}
