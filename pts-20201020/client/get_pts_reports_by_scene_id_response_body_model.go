// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsReportsBySceneIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPtsReportsBySceneIdResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetPtsReportsBySceneIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPtsReportsBySceneIdResponseBody
	GetMessage() *string
	SetReportOverViewList(v []*GetPtsReportsBySceneIdResponseBodyReportOverViewList) *GetPtsReportsBySceneIdResponseBody
	GetReportOverViewList() []*GetPtsReportsBySceneIdResponseBodyReportOverViewList
	SetRequestId(v string) *GetPtsReportsBySceneIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPtsReportsBySceneIdResponseBody
	GetSuccess() *bool
}

type GetPtsReportsBySceneIdResponseBody struct {
	// The system status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. If the request was successful, this parameter is left empty.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reports.
	ReportOverViewList []*GetPtsReportsBySceneIdResponseBodyReportOverViewList `json:"ReportOverViewList,omitempty" xml:"ReportOverViewList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// DC4E3177-6745-4925-B423-4E89VV34221A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPtsReportsBySceneIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportsBySceneIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetPtsReportsBySceneIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPtsReportsBySceneIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPtsReportsBySceneIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPtsReportsBySceneIdResponseBody) GetReportOverViewList() []*GetPtsReportsBySceneIdResponseBodyReportOverViewList {
	return s.ReportOverViewList
}

func (s *GetPtsReportsBySceneIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPtsReportsBySceneIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPtsReportsBySceneIdResponseBody) SetCode(v string) *GetPtsReportsBySceneIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBody) SetHttpStatusCode(v int32) *GetPtsReportsBySceneIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBody) SetMessage(v string) *GetPtsReportsBySceneIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBody) SetReportOverViewList(v []*GetPtsReportsBySceneIdResponseBodyReportOverViewList) *GetPtsReportsBySceneIdResponseBody {
	s.ReportOverViewList = v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBody) SetRequestId(v string) *GetPtsReportsBySceneIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBody) SetSuccess(v bool) *GetPtsReportsBySceneIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBody) Validate() error {
	if s.ReportOverViewList != nil {
		for _, item := range s.ReportOverViewList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPtsReportsBySceneIdResponseBodyReportOverViewList struct {
	// The number of stress testers.
	//
	// example:
	//
	// 1
	AgentCount *int32 `json:"AgentCount,omitempty" xml:"AgentCount,omitempty"`
	// The end time of the stress testing.
	//
	// example:
	//
	// 2021-02-26 16:38:30
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The report ID.
	//
	// example:
	//
	// NGGB5FV
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The title of the report.
	//
	// example:
	//
	// PTS-test-20240920094710
	ReportName *string `json:"ReportName,omitempty" xml:"ReportName,omitempty"`
	// The start time of the stress testing.
	//
	// example:
	//
	// 2021-02-26 16:28:30
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The consumed Virtual User Minutes (VUM).
	//
	// example:
	//
	// 100
	Vum *int64 `json:"Vum,omitempty" xml:"Vum,omitempty"`
}

func (s GetPtsReportsBySceneIdResponseBodyReportOverViewList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportsBySceneIdResponseBodyReportOverViewList) GoString() string {
	return s.String()
}

func (s *GetPtsReportsBySceneIdResponseBodyReportOverViewList) GetAgentCount() *int32 {
	return s.AgentCount
}

func (s *GetPtsReportsBySceneIdResponseBodyReportOverViewList) GetEndTime() *string {
	return s.EndTime
}

func (s *GetPtsReportsBySceneIdResponseBodyReportOverViewList) GetReportId() *string {
	return s.ReportId
}

func (s *GetPtsReportsBySceneIdResponseBodyReportOverViewList) GetReportName() *string {
	return s.ReportName
}

func (s *GetPtsReportsBySceneIdResponseBodyReportOverViewList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetPtsReportsBySceneIdResponseBodyReportOverViewList) GetVum() *int64 {
	return s.Vum
}

func (s *GetPtsReportsBySceneIdResponseBodyReportOverViewList) SetAgentCount(v int32) *GetPtsReportsBySceneIdResponseBodyReportOverViewList {
	s.AgentCount = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBodyReportOverViewList) SetEndTime(v string) *GetPtsReportsBySceneIdResponseBodyReportOverViewList {
	s.EndTime = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBodyReportOverViewList) SetReportId(v string) *GetPtsReportsBySceneIdResponseBodyReportOverViewList {
	s.ReportId = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBodyReportOverViewList) SetReportName(v string) *GetPtsReportsBySceneIdResponseBodyReportOverViewList {
	s.ReportName = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBodyReportOverViewList) SetStartTime(v string) *GetPtsReportsBySceneIdResponseBodyReportOverViewList {
	s.StartTime = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBodyReportOverViewList) SetVum(v int64) *GetPtsReportsBySceneIdResponseBodyReportOverViewList {
	s.Vum = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBodyReportOverViewList) Validate() error {
	return dara.Validate(s)
}
