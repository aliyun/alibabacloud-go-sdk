// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppStreamingOutTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAppStreamingOutTemplatesResponseBody
	GetRequestId() *string
	SetTemplates(v []*DescribeAppStreamingOutTemplatesResponseBodyTemplates) *DescribeAppStreamingOutTemplatesResponseBody
	GetTemplates() []*DescribeAppStreamingOutTemplatesResponseBodyTemplates
	SetTotalNum(v int64) *DescribeAppStreamingOutTemplatesResponseBody
	GetTotalNum() *int64
	SetTotalPage(v int64) *DescribeAppStreamingOutTemplatesResponseBody
	GetTotalPage() *int64
}

type DescribeAppStreamingOutTemplatesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B0A2FCBC-43A4-428F-BC1D-3F4F85837F76
	RequestId *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates []*DescribeAppStreamingOutTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeAppStreamingOutTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppStreamingOutTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppStreamingOutTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppStreamingOutTemplatesResponseBody) GetTemplates() []*DescribeAppStreamingOutTemplatesResponseBodyTemplates {
	return s.Templates
}

func (s *DescribeAppStreamingOutTemplatesResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *DescribeAppStreamingOutTemplatesResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *DescribeAppStreamingOutTemplatesResponseBody) SetRequestId(v string) *DescribeAppStreamingOutTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponseBody) SetTemplates(v []*DescribeAppStreamingOutTemplatesResponseBodyTemplates) *DescribeAppStreamingOutTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponseBody) SetTotalNum(v int64) *DescribeAppStreamingOutTemplatesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponseBody) SetTotalPage(v int64) *DescribeAppStreamingOutTemplatesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponseBody) Validate() error {
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

type DescribeAppStreamingOutTemplatesResponseBodyTemplates struct {
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

func (s DescribeAppStreamingOutTemplatesResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppStreamingOutTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeAppStreamingOutTemplatesResponseBodyTemplates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAppStreamingOutTemplatesResponseBodyTemplates) GetLayoutIds() []*string {
	return s.LayoutIds
}

func (s *DescribeAppStreamingOutTemplatesResponseBodyTemplates) GetMediaEncode() *int32 {
	return s.MediaEncode
}

func (s *DescribeAppStreamingOutTemplatesResponseBodyTemplates) GetName() *string {
	return s.Name
}

func (s *DescribeAppStreamingOutTemplatesResponseBodyTemplates) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeAppStreamingOutTemplatesResponseBodyTemplates) SetCreateTime(v string) *DescribeAppStreamingOutTemplatesResponseBodyTemplates {
	s.CreateTime = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponseBodyTemplates) SetLayoutIds(v []*string) *DescribeAppStreamingOutTemplatesResponseBodyTemplates {
	s.LayoutIds = v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponseBodyTemplates) SetMediaEncode(v int32) *DescribeAppStreamingOutTemplatesResponseBodyTemplates {
	s.MediaEncode = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponseBodyTemplates) SetName(v string) *DescribeAppStreamingOutTemplatesResponseBodyTemplates {
	s.Name = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponseBodyTemplates) SetTemplateId(v string) *DescribeAppStreamingOutTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponseBodyTemplates) Validate() error {
	return dara.Validate(s)
}
