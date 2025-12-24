// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNotifyTemplateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeNotifyTemplateListResponseBodyData) *DescribeNotifyTemplateListResponseBody
	GetData() []*DescribeNotifyTemplateListResponseBodyData
	SetPage(v *DescribeNotifyTemplateListResponseBodyPage) *DescribeNotifyTemplateListResponseBody
	GetPage() *DescribeNotifyTemplateListResponseBodyPage
	SetRequestId(v string) *DescribeNotifyTemplateListResponseBody
	GetRequestId() *string
}

type DescribeNotifyTemplateListResponseBody struct {
	// The data returned.
	Data []*DescribeNotifyTemplateListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The pagination information.
	Page *DescribeNotifyTemplateListResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B3FED5B9-190A-5952-93A4-24FBF0F0C573
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNotifyTemplateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNotifyTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNotifyTemplateListResponseBody) GetData() []*DescribeNotifyTemplateListResponseBodyData {
	return s.Data
}

func (s *DescribeNotifyTemplateListResponseBody) GetPage() *DescribeNotifyTemplateListResponseBodyPage {
	return s.Page
}

func (s *DescribeNotifyTemplateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNotifyTemplateListResponseBody) SetData(v []*DescribeNotifyTemplateListResponseBodyData) *DescribeNotifyTemplateListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeNotifyTemplateListResponseBody) SetPage(v *DescribeNotifyTemplateListResponseBodyPage) *DescribeNotifyTemplateListResponseBody {
	s.Page = v
	return s
}

func (s *DescribeNotifyTemplateListResponseBody) SetRequestId(v string) *DescribeNotifyTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNotifyTemplateListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNotifyTemplateListResponseBodyData struct {
	// The body of the template.
	//
	// example:
	//
	// Dear $aliyunUID : Cloud Security Center Threat Analysis and Response has detected a newly discovered security incident $incidentName(Incident id :$incidentID) in $startTime, Please go to Cloud Security Center Console View.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the message event corresponding to the template.
	//
	// example:
	//
	// yundun_soar_incident_generate
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The parameters of the template.
	//
	// example:
	//
	// [\\"aliyunUID\\",\\"incidentName\\",\\"incidentID\\",\\"startTime\\"]
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The title of the template.
	//
	// example:
	//
	// [Alibaba Cloud Threat Analysis and Response] has detected a newly discovered security incident $incidentName($incidentID)
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// incident generate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DescribeNotifyTemplateListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeNotifyTemplateListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeNotifyTemplateListResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *DescribeNotifyTemplateListResponseBodyData) GetEventId() *string {
	return s.EventId
}

func (s *DescribeNotifyTemplateListResponseBodyData) GetParams() *string {
	return s.Params
}

func (s *DescribeNotifyTemplateListResponseBodyData) GetSubject() *string {
	return s.Subject
}

func (s *DescribeNotifyTemplateListResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeNotifyTemplateListResponseBodyData) SetContent(v string) *DescribeNotifyTemplateListResponseBodyData {
	s.Content = &v
	return s
}

func (s *DescribeNotifyTemplateListResponseBodyData) SetEventId(v string) *DescribeNotifyTemplateListResponseBodyData {
	s.EventId = &v
	return s
}

func (s *DescribeNotifyTemplateListResponseBodyData) SetParams(v string) *DescribeNotifyTemplateListResponseBodyData {
	s.Params = &v
	return s
}

func (s *DescribeNotifyTemplateListResponseBodyData) SetSubject(v string) *DescribeNotifyTemplateListResponseBodyData {
	s.Subject = &v
	return s
}

func (s *DescribeNotifyTemplateListResponseBodyData) SetTemplateName(v string) *DescribeNotifyTemplateListResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *DescribeNotifyTemplateListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeNotifyTemplateListResponseBodyPage struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNotifyTemplateListResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s DescribeNotifyTemplateListResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribeNotifyTemplateListResponseBodyPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeNotifyTemplateListResponseBodyPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNotifyTemplateListResponseBodyPage) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeNotifyTemplateListResponseBodyPage) SetPageNumber(v int32) *DescribeNotifyTemplateListResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeNotifyTemplateListResponseBodyPage) SetPageSize(v int32) *DescribeNotifyTemplateListResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribeNotifyTemplateListResponseBodyPage) SetTotalCount(v int32) *DescribeNotifyTemplateListResponseBodyPage {
	s.TotalCount = &v
	return s
}

func (s *DescribeNotifyTemplateListResponseBodyPage) Validate() error {
	return dara.Validate(s)
}
