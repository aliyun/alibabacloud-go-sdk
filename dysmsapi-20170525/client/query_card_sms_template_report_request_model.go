// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCardSmsTemplateReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *QueryCardSmsTemplateReportRequest
	GetEndDate() *string
	SetStartDate(v string) *QueryCardSmsTemplateReportRequest
	GetStartDate() *string
	SetTemplateCodes(v []*string) *QueryCardSmsTemplateReportRequest
	GetTemplateCodes() []*string
}

type QueryCardSmsTemplateReportRequest struct {
	// The end of the time range to query. Specify the time in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2020-10-11 00:00:01
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2020-10-10 00:00:01
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The array of message templates.
	//
	// This parameter is required.
	TemplateCodes []*string `json:"TemplateCodes,omitempty" xml:"TemplateCodes,omitempty" type:"Repeated"`
}

func (s QueryCardSmsTemplateReportRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCardSmsTemplateReportRequest) GoString() string {
	return s.String()
}

func (s *QueryCardSmsTemplateReportRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryCardSmsTemplateReportRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryCardSmsTemplateReportRequest) GetTemplateCodes() []*string {
	return s.TemplateCodes
}

func (s *QueryCardSmsTemplateReportRequest) SetEndDate(v string) *QueryCardSmsTemplateReportRequest {
	s.EndDate = &v
	return s
}

func (s *QueryCardSmsTemplateReportRequest) SetStartDate(v string) *QueryCardSmsTemplateReportRequest {
	s.StartDate = &v
	return s
}

func (s *QueryCardSmsTemplateReportRequest) SetTemplateCodes(v []*string) *QueryCardSmsTemplateReportRequest {
	s.TemplateCodes = v
	return s
}

func (s *QueryCardSmsTemplateReportRequest) Validate() error {
	return dara.Validate(s)
}
