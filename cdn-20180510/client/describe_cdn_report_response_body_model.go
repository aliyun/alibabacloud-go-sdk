// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *DescribeCdnReportResponseBodyContent) *DescribeCdnReportResponseBody
	GetContent() *DescribeCdnReportResponseBodyContent
	SetRequestId(v string) *DescribeCdnReportResponseBody
	GetRequestId() *string
}

type DescribeCdnReportResponseBody struct {
	// The content of the operations report.
	//
	// example:
	//
	// "data":[{"deliver":{"report":{"title":"TopUrlByAcc","format":"table","shape":"","header":["url","traf","traf_rate","acc","acc_rate"]}},"data":[{"acc":440,"acc_rate":"0.200%","traf":22,"url":"http://demo.com","traf_rate":"0.100%"},{"acc":440,"acc_rate":"0.200%","traf":22,"url":"http://demo.com","traf_rate":"0.100%"}]}]}}
	Content *DescribeCdnReportResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportResponseBody) GetContent() *DescribeCdnReportResponseBodyContent {
	return s.Content
}

func (s *DescribeCdnReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnReportResponseBody) SetContent(v *DescribeCdnReportResponseBodyContent) *DescribeCdnReportResponseBody {
	s.Content = v
	return s
}

func (s *DescribeCdnReportResponseBody) SetRequestId(v string) *DescribeCdnReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnReportResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnReportResponseBodyContent struct {
	Data []*DescribeCdnReportResponseBodyContentData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeCdnReportResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnReportResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportResponseBodyContent) GetData() []*DescribeCdnReportResponseBodyContentData {
	return s.Data
}

func (s *DescribeCdnReportResponseBodyContent) SetData(v []*DescribeCdnReportResponseBodyContentData) *DescribeCdnReportResponseBodyContent {
	s.Data = v
	return s
}

func (s *DescribeCdnReportResponseBodyContent) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdnReportResponseBodyContentData struct {
	Data    []map[string]*string                             `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Deliver *DescribeCdnReportResponseBodyContentDataDeliver `json:"deliver,omitempty" xml:"deliver,omitempty" type:"Struct"`
}

func (s DescribeCdnReportResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnReportResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportResponseBodyContentData) GetData() []map[string]*string {
	return s.Data
}

func (s *DescribeCdnReportResponseBodyContentData) GetDeliver() *DescribeCdnReportResponseBodyContentDataDeliver {
	return s.Deliver
}

func (s *DescribeCdnReportResponseBodyContentData) SetData(v []map[string]*string) *DescribeCdnReportResponseBodyContentData {
	s.Data = v
	return s
}

func (s *DescribeCdnReportResponseBodyContentData) SetDeliver(v *DescribeCdnReportResponseBodyContentDataDeliver) *DescribeCdnReportResponseBodyContentData {
	s.Deliver = v
	return s
}

func (s *DescribeCdnReportResponseBodyContentData) Validate() error {
	if s.Deliver != nil {
		if err := s.Deliver.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnReportResponseBodyContentDataDeliver struct {
	Report *DescribeCdnReportResponseBodyContentDataDeliverReport `json:"report,omitempty" xml:"report,omitempty" type:"Struct"`
}

func (s DescribeCdnReportResponseBodyContentDataDeliver) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnReportResponseBodyContentDataDeliver) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportResponseBodyContentDataDeliver) GetReport() *DescribeCdnReportResponseBodyContentDataDeliverReport {
	return s.Report
}

func (s *DescribeCdnReportResponseBodyContentDataDeliver) SetReport(v *DescribeCdnReportResponseBodyContentDataDeliverReport) *DescribeCdnReportResponseBodyContentDataDeliver {
	s.Report = v
	return s
}

func (s *DescribeCdnReportResponseBodyContentDataDeliver) Validate() error {
	if s.Report != nil {
		if err := s.Report.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnReportResponseBodyContentDataDeliverReport struct {
	// example:
	//
	// table
	Format *string   `json:"format,omitempty" xml:"format,omitempty"`
	Header []*string `json:"header,omitempty" xml:"header,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	OutLine *int64 `json:"outLine,omitempty" xml:"outLine,omitempty"`
	// example:
	//
	// 0
	OutSize *int64 `json:"outSize,omitempty" xml:"outSize,omitempty"`
	// example:
	//
	// line
	Shape *string `json:"shape,omitempty" xml:"shape,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s DescribeCdnReportResponseBodyContentDataDeliverReport) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnReportResponseBodyContentDataDeliverReport) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportResponseBodyContentDataDeliverReport) GetFormat() *string {
	return s.Format
}

func (s *DescribeCdnReportResponseBodyContentDataDeliverReport) GetHeader() []*string {
	return s.Header
}

func (s *DescribeCdnReportResponseBodyContentDataDeliverReport) GetOutLine() *int64 {
	return s.OutLine
}

func (s *DescribeCdnReportResponseBodyContentDataDeliverReport) GetOutSize() *int64 {
	return s.OutSize
}

func (s *DescribeCdnReportResponseBodyContentDataDeliverReport) GetShape() *string {
	return s.Shape
}

func (s *DescribeCdnReportResponseBodyContentDataDeliverReport) GetTitle() *string {
	return s.Title
}

func (s *DescribeCdnReportResponseBodyContentDataDeliverReport) SetFormat(v string) *DescribeCdnReportResponseBodyContentDataDeliverReport {
	s.Format = &v
	return s
}

func (s *DescribeCdnReportResponseBodyContentDataDeliverReport) SetHeader(v []*string) *DescribeCdnReportResponseBodyContentDataDeliverReport {
	s.Header = v
	return s
}

func (s *DescribeCdnReportResponseBodyContentDataDeliverReport) SetOutLine(v int64) *DescribeCdnReportResponseBodyContentDataDeliverReport {
	s.OutLine = &v
	return s
}

func (s *DescribeCdnReportResponseBodyContentDataDeliverReport) SetOutSize(v int64) *DescribeCdnReportResponseBodyContentDataDeliverReport {
	s.OutSize = &v
	return s
}

func (s *DescribeCdnReportResponseBodyContentDataDeliverReport) SetShape(v string) *DescribeCdnReportResponseBodyContentDataDeliverReport {
	s.Shape = &v
	return s
}

func (s *DescribeCdnReportResponseBodyContentDataDeliverReport) SetTitle(v string) *DescribeCdnReportResponseBodyContentDataDeliverReport {
	s.Title = &v
	return s
}

func (s *DescribeCdnReportResponseBodyContentDataDeliverReport) Validate() error {
	return dara.Validate(s)
}
