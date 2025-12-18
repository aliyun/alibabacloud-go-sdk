// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceGroupInspectReportListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceGroupInspectReportListResponseBody
	GetCode() *string
	SetData(v []*GetInstanceGroupInspectReportListResponseBodyData) *GetInstanceGroupInspectReportListResponseBody
	GetData() []*GetInstanceGroupInspectReportListResponseBodyData
	SetMessage(v string) *GetInstanceGroupInspectReportListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceGroupInspectReportListResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetInstanceGroupInspectReportListResponseBody
	GetSuccess() *string
}

type GetInstanceGroupInspectReportListResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// List<ReportStatus>
	Data []*GetInstanceGroupInspectReportListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s GetInstanceGroupInspectReportListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceGroupInspectReportListResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceGroupInspectReportListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceGroupInspectReportListResponseBody) GetData() []*GetInstanceGroupInspectReportListResponseBodyData {
	return s.Data
}

func (s *GetInstanceGroupInspectReportListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceGroupInspectReportListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceGroupInspectReportListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetInstanceGroupInspectReportListResponseBody) SetCode(v string) *GetInstanceGroupInspectReportListResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceGroupInspectReportListResponseBody) SetData(v []*GetInstanceGroupInspectReportListResponseBodyData) *GetInstanceGroupInspectReportListResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceGroupInspectReportListResponseBody) SetMessage(v string) *GetInstanceGroupInspectReportListResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceGroupInspectReportListResponseBody) SetRequestId(v string) *GetInstanceGroupInspectReportListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceGroupInspectReportListResponseBody) SetSuccess(v string) *GetInstanceGroupInspectReportListResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceGroupInspectReportListResponseBody) Validate() error {
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

type GetInstanceGroupInspectReportListResponseBodyData struct {
	// example:
	//
	// 2025-12-11 00:39:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2025-12-10
	ReportDate *string `json:"ReportDate,omitempty" xml:"ReportDate,omitempty"`
	// example:
	//
	// 13f52040-5a6e-42c3-bb84-051f5d6d****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// example:
	//
	// 3
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetInstanceGroupInspectReportListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceGroupInspectReportListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceGroupInspectReportListResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetInstanceGroupInspectReportListResponseBodyData) GetReportDate() *string {
	return s.ReportDate
}

func (s *GetInstanceGroupInspectReportListResponseBodyData) GetReportId() *string {
	return s.ReportId
}

func (s *GetInstanceGroupInspectReportListResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceGroupInspectReportListResponseBodyData) SetCreateTime(v string) *GetInstanceGroupInspectReportListResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceGroupInspectReportListResponseBodyData) SetReportDate(v string) *GetInstanceGroupInspectReportListResponseBodyData {
	s.ReportDate = &v
	return s
}

func (s *GetInstanceGroupInspectReportListResponseBodyData) SetReportId(v string) *GetInstanceGroupInspectReportListResponseBodyData {
	s.ReportId = &v
	return s
}

func (s *GetInstanceGroupInspectReportListResponseBodyData) SetStatus(v string) *GetInstanceGroupInspectReportListResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetInstanceGroupInspectReportListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
