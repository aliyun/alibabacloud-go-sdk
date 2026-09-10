// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckReportInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDataCheckReportInstanceResponseBodyData) *ListDataCheckReportInstanceResponseBody
	GetData() []*ListDataCheckReportInstanceResponseBodyData
	SetErrCode(v string) *ListDataCheckReportInstanceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListDataCheckReportInstanceResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ListDataCheckReportInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataCheckReportInstanceResponseBody
	GetSuccess() *bool
}

type ListDataCheckReportInstanceResponseBody struct {
	// The list of report historical instances.
	Data []*ListDataCheckReportInstanceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error code. An empty string is returned if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. An empty string is returned if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues for this call.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call failed. Troubleshoot by using errCode and errMessage.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListDataCheckReportInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckReportInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataCheckReportInstanceResponseBody) GetData() []*ListDataCheckReportInstanceResponseBodyData {
	return s.Data
}

func (s *ListDataCheckReportInstanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListDataCheckReportInstanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListDataCheckReportInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataCheckReportInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataCheckReportInstanceResponseBody) SetData(v []*ListDataCheckReportInstanceResponseBodyData) *ListDataCheckReportInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ListDataCheckReportInstanceResponseBody) SetErrCode(v string) *ListDataCheckReportInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListDataCheckReportInstanceResponseBody) SetErrMessage(v string) *ListDataCheckReportInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListDataCheckReportInstanceResponseBody) SetRequestId(v string) *ListDataCheckReportInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataCheckReportInstanceResponseBody) SetSuccess(v bool) *ListDataCheckReportInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataCheckReportInstanceResponseBody) Validate() error {
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

type ListDataCheckReportInstanceResponseBodyData struct {
	// The batch ID.
	//
	// example:
	//
	// 20001
	BatchId *string `json:"batchId,omitempty" xml:"batchId,omitempty"`
	// The report label.
	//
	// example:
	//
	// daily_check
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// The report generation time.
	//
	// example:
	//
	// 2024-01-01 12:00:00
	ReportTime *string `json:"reportTime,omitempty" xml:"reportTime,omitempty"`
}

func (s ListDataCheckReportInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckReportInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataCheckReportInstanceResponseBodyData) GetBatchId() *string {
	return s.BatchId
}

func (s *ListDataCheckReportInstanceResponseBodyData) GetLabel() *string {
	return s.Label
}

func (s *ListDataCheckReportInstanceResponseBodyData) GetReportTime() *string {
	return s.ReportTime
}

func (s *ListDataCheckReportInstanceResponseBodyData) SetBatchId(v string) *ListDataCheckReportInstanceResponseBodyData {
	s.BatchId = &v
	return s
}

func (s *ListDataCheckReportInstanceResponseBodyData) SetLabel(v string) *ListDataCheckReportInstanceResponseBodyData {
	s.Label = &v
	return s
}

func (s *ListDataCheckReportInstanceResponseBodyData) SetReportTime(v string) *ListDataCheckReportInstanceResponseBodyData {
	s.ReportTime = &v
	return s
}

func (s *ListDataCheckReportInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
