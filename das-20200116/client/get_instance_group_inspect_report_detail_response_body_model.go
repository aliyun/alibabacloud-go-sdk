// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceGroupInspectReportDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceGroupInspectReportDetailResponseBody
	GetCode() *string
	SetData(v *GetInstanceGroupInspectReportDetailResponseBodyData) *GetInstanceGroupInspectReportDetailResponseBody
	GetData() *GetInstanceGroupInspectReportDetailResponseBodyData
	SetMessage(v string) *GetInstanceGroupInspectReportDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceGroupInspectReportDetailResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetInstanceGroupInspectReportDetailResponseBody
	GetSuccess() *string
}

type GetInstanceGroupInspectReportDetailResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// ReportDetail
	Data *GetInstanceGroupInspectReportDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceGroupInspectReportDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceGroupInspectReportDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceGroupInspectReportDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceGroupInspectReportDetailResponseBody) GetData() *GetInstanceGroupInspectReportDetailResponseBodyData {
	return s.Data
}

func (s *GetInstanceGroupInspectReportDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceGroupInspectReportDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceGroupInspectReportDetailResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetInstanceGroupInspectReportDetailResponseBody) SetCode(v string) *GetInstanceGroupInspectReportDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceGroupInspectReportDetailResponseBody) SetData(v *GetInstanceGroupInspectReportDetailResponseBodyData) *GetInstanceGroupInspectReportDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceGroupInspectReportDetailResponseBody) SetMessage(v string) *GetInstanceGroupInspectReportDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceGroupInspectReportDetailResponseBody) SetRequestId(v string) *GetInstanceGroupInspectReportDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceGroupInspectReportDetailResponseBody) SetSuccess(v string) *GetInstanceGroupInspectReportDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceGroupInspectReportDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceGroupInspectReportDetailResponseBodyData struct {
	ReportDetail *string `json:"ReportDetail,omitempty" xml:"ReportDetail,omitempty"`
	// example:
	//
	// 13f52040-5a6e-42c3-bb84-051f5d6d****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s GetInstanceGroupInspectReportDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceGroupInspectReportDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceGroupInspectReportDetailResponseBodyData) GetReportDetail() *string {
	return s.ReportDetail
}

func (s *GetInstanceGroupInspectReportDetailResponseBodyData) GetReportId() *string {
	return s.ReportId
}

func (s *GetInstanceGroupInspectReportDetailResponseBodyData) SetReportDetail(v string) *GetInstanceGroupInspectReportDetailResponseBodyData {
	s.ReportDetail = &v
	return s
}

func (s *GetInstanceGroupInspectReportDetailResponseBodyData) SetReportId(v string) *GetInstanceGroupInspectReportDetailResponseBodyData {
	s.ReportId = &v
	return s
}

func (s *GetInstanceGroupInspectReportDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}
