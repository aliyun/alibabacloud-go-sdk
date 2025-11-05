// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnReportListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *DescribeCdnReportListResponseBodyContent) *DescribeCdnReportListResponseBody
	GetContent() *DescribeCdnReportListResponseBodyContent
	SetRequestId(v string) *DescribeCdnReportListResponseBody
	GetRequestId() *string
}

type DescribeCdnReportListResponseBody struct {
	// The information about the report that is queried.
	//
	// example:
	//
	// "data":[{"reportId":1,"deliver":{"report":{"title":"DomainPvUv","format":"chart","shape":"line","xAxis":"ds","yAxis":"cnt","legend":"cnt_type","header":["ds","cnt_type","cnt"]}}}
	Content *DescribeCdnReportListResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnReportListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnReportListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportListResponseBody) GetContent() *DescribeCdnReportListResponseBodyContent {
	return s.Content
}

func (s *DescribeCdnReportListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnReportListResponseBody) SetContent(v *DescribeCdnReportListResponseBodyContent) *DescribeCdnReportListResponseBody {
	s.Content = v
	return s
}

func (s *DescribeCdnReportListResponseBody) SetRequestId(v string) *DescribeCdnReportListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnReportListResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnReportListResponseBodyContent struct {
	Data []*DescribeCdnReportListResponseBodyContentData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeCdnReportListResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnReportListResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportListResponseBodyContent) GetData() []*DescribeCdnReportListResponseBodyContentData {
	return s.Data
}

func (s *DescribeCdnReportListResponseBodyContent) SetData(v []*DescribeCdnReportListResponseBodyContentData) *DescribeCdnReportListResponseBodyContent {
	s.Data = v
	return s
}

func (s *DescribeCdnReportListResponseBodyContent) Validate() error {
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

type DescribeCdnReportListResponseBodyContentData struct {
	Deliver *DescribeCdnReportListResponseBodyContentDataDeliver `json:"deliver,omitempty" xml:"deliver,omitempty" type:"Struct"`
	// example:
	//
	// 1
	ReportId *int64 `json:"reportId,omitempty" xml:"reportId,omitempty"`
}

func (s DescribeCdnReportListResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnReportListResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportListResponseBodyContentData) GetDeliver() *DescribeCdnReportListResponseBodyContentDataDeliver {
	return s.Deliver
}

func (s *DescribeCdnReportListResponseBodyContentData) GetReportId() *int64 {
	return s.ReportId
}

func (s *DescribeCdnReportListResponseBodyContentData) SetDeliver(v *DescribeCdnReportListResponseBodyContentDataDeliver) *DescribeCdnReportListResponseBodyContentData {
	s.Deliver = v
	return s
}

func (s *DescribeCdnReportListResponseBodyContentData) SetReportId(v int64) *DescribeCdnReportListResponseBodyContentData {
	s.ReportId = &v
	return s
}

func (s *DescribeCdnReportListResponseBodyContentData) Validate() error {
	if s.Deliver != nil {
		if err := s.Deliver.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnReportListResponseBodyContentDataDeliver struct {
	Report *DescribeCdnReportListResponseBodyContentDataDeliverReport `json:"report,omitempty" xml:"report,omitempty" type:"Struct"`
}

func (s DescribeCdnReportListResponseBodyContentDataDeliver) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnReportListResponseBodyContentDataDeliver) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportListResponseBodyContentDataDeliver) GetReport() *DescribeCdnReportListResponseBodyContentDataDeliverReport {
	return s.Report
}

func (s *DescribeCdnReportListResponseBodyContentDataDeliver) SetReport(v *DescribeCdnReportListResponseBodyContentDataDeliverReport) *DescribeCdnReportListResponseBodyContentDataDeliver {
	s.Report = v
	return s
}

func (s *DescribeCdnReportListResponseBodyContentDataDeliver) Validate() error {
	if s.Report != nil {
		if err := s.Report.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnReportListResponseBodyContentDataDeliverReport struct {
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
	// TopUrlByAcc
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s DescribeCdnReportListResponseBodyContentDataDeliverReport) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnReportListResponseBodyContentDataDeliverReport) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportListResponseBodyContentDataDeliverReport) GetFormat() *string {
	return s.Format
}

func (s *DescribeCdnReportListResponseBodyContentDataDeliverReport) GetHeader() []*string {
	return s.Header
}

func (s *DescribeCdnReportListResponseBodyContentDataDeliverReport) GetOutLine() *int64 {
	return s.OutLine
}

func (s *DescribeCdnReportListResponseBodyContentDataDeliverReport) GetOutSize() *int64 {
	return s.OutSize
}

func (s *DescribeCdnReportListResponseBodyContentDataDeliverReport) GetShape() *string {
	return s.Shape
}

func (s *DescribeCdnReportListResponseBodyContentDataDeliverReport) GetTitle() *string {
	return s.Title
}

func (s *DescribeCdnReportListResponseBodyContentDataDeliverReport) SetFormat(v string) *DescribeCdnReportListResponseBodyContentDataDeliverReport {
	s.Format = &v
	return s
}

func (s *DescribeCdnReportListResponseBodyContentDataDeliverReport) SetHeader(v []*string) *DescribeCdnReportListResponseBodyContentDataDeliverReport {
	s.Header = v
	return s
}

func (s *DescribeCdnReportListResponseBodyContentDataDeliverReport) SetOutLine(v int64) *DescribeCdnReportListResponseBodyContentDataDeliverReport {
	s.OutLine = &v
	return s
}

func (s *DescribeCdnReportListResponseBodyContentDataDeliverReport) SetOutSize(v int64) *DescribeCdnReportListResponseBodyContentDataDeliverReport {
	s.OutSize = &v
	return s
}

func (s *DescribeCdnReportListResponseBodyContentDataDeliverReport) SetShape(v string) *DescribeCdnReportListResponseBodyContentDataDeliverReport {
	s.Shape = &v
	return s
}

func (s *DescribeCdnReportListResponseBodyContentDataDeliverReport) SetTitle(v string) *DescribeCdnReportListResponseBodyContentDataDeliverReport {
	s.Title = &v
	return s
}

func (s *DescribeCdnReportListResponseBodyContentDataDeliverReport) Validate() error {
	return dara.Validate(s)
}
