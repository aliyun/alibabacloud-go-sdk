// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *GetAiReportRequest
	GetCallId() *string
	SetReportId(v string) *GetAiReportRequest
	GetReportId() *string
}

type GetAiReportRequest struct {
	// The data call ID, which is obtained from the callId returned by InvokeDataProduct. You must specify either this parameter or reportId. You cannot specify both or leave both empty. The value must start with call- and be 1 to 64 characters in length. It can contain only ASCII letters, digits, underscores (_), and hyphens (-). The value is case-sensitive and leading and trailing spaces are not trimmed. If the format is invalid, InvalidParameter.CallId is returned. If both parameters are specified or left empty, InvalidParameter.ReportReference is returned. This parameter is used to query the status of per-use reports. Querying the status of periodic reports by this ID always returns WAITING_PERIOD. Use the ReportId from the list instead.
	//
	// example:
	//
	// call-3f2b8c1e-6d4a-4b7e-9a1c-2e5f7d8b9c0a
	CallId *string `json:"callId,omitempty" xml:"callId,omitempty"`
	// The report ID, which is obtained from the ReportId returned by ListProductReports or the Report.ReportId in the previous available result of GetAiReport. You must specify either this parameter or callId. You cannot specify both or leave both empty. The value must start with report- and be 1 to 64 characters in length. It can contain only ASCII letters, digits, underscores (_), and hyphens (-). The value is case-sensitive and leading and trailing spaces are not trimmed. If the format is invalid, InvalidParameter.ReportId is returned. If both parameters are specified or left empty, InvalidParameter.ReportReference is returned. You can only query reports that belong to the current Alibaba Cloud account.
	//
	// example:
	//
	// report-3f2b8c1e-6d4a-4b7e-9a1c-2e5f7d8b9c0a
	ReportId *string `json:"reportId,omitempty" xml:"reportId,omitempty"`
}

func (s GetAiReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiReportRequest) GoString() string {
	return s.String()
}

func (s *GetAiReportRequest) GetCallId() *string {
	return s.CallId
}

func (s *GetAiReportRequest) GetReportId() *string {
	return s.ReportId
}

func (s *GetAiReportRequest) SetCallId(v string) *GetAiReportRequest {
	s.CallId = &v
	return s
}

func (s *GetAiReportRequest) SetReportId(v string) *GetAiReportRequest {
	s.ReportId = &v
	return s
}

func (s *GetAiReportRequest) Validate() error {
	return dara.Validate(s)
}
