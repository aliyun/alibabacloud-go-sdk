// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetListRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetListRecordResponseBody
	GetCode() *string
	SetData(v []*GetListRecordResponseBodyData) *GetListRecordResponseBody
	GetData() []*GetListRecordResponseBodyData
	SetMessage(v string) *GetListRecordResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetListRecordResponseBody
	GetRequestId() *string
	SetTotal(v int64) *GetListRecordResponseBody
	GetTotal() *int64
}

type GetListRecordResponseBody struct {
	// Status code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Returned data.
	Data []*GetListRecordResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Description of the status code.
	//
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Total number of records
	//
	// example:
	//
	// 19
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetListRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetListRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetListRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetListRecordResponseBody) GetData() []*GetListRecordResponseBodyData {
	return s.Data
}

func (s *GetListRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetListRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetListRecordResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *GetListRecordResponseBody) SetCode(v string) *GetListRecordResponseBody {
	s.Code = &v
	return s
}

func (s *GetListRecordResponseBody) SetData(v []*GetListRecordResponseBodyData) *GetListRecordResponseBody {
	s.Data = v
	return s
}

func (s *GetListRecordResponseBody) SetMessage(v string) *GetListRecordResponseBody {
	s.Message = &v
	return s
}

func (s *GetListRecordResponseBody) SetRequestId(v string) *GetListRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetListRecordResponseBody) SetTotal(v int64) *GetListRecordResponseBody {
	s.Total = &v
	return s
}

func (s *GetListRecordResponseBody) Validate() error {
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

type GetListRecordResponseBodyData struct {
	// AI analysis ID
	//
	// example:
	//
	// 16896fa8-37f6-4c70-bb32-67fa9817d426
	AnalysisId *string `json:"analysisId,omitempty" xml:"analysisId,omitempty"`
	// Analysis time
	//
	// example:
	//
	// 2024-12-24 12:02:05
	AnalysisTime *string `json:"analysisTime,omitempty" xml:"analysisTime,omitempty"`
	// Analysis parameters for the AI job
	//
	// example:
	//
	// timeout=2000 ms
	Arguments *string `json:"arguments,omitempty" xml:"arguments,omitempty"`
	// Analysis failure log
	//
	// example:
	//
	// 机器i-wz9dej066kii4goqpnze分析失败, 失败原因: Not get GPU trace data for \\"e59ce870-dbd4-4c44-a814-174ac6ab5bcf\\" \\"[\\"118534\\"]\\"!
	FailedLog *string `json:"failedLog,omitempty" xml:"failedLog,omitempty"`
	// Analysis status
	//
	// example:
	//
	// 已完成/分析失败
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetListRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetListRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetListRecordResponseBodyData) GetAnalysisId() *string {
	return s.AnalysisId
}

func (s *GetListRecordResponseBodyData) GetAnalysisTime() *string {
	return s.AnalysisTime
}

func (s *GetListRecordResponseBodyData) GetArguments() *string {
	return s.Arguments
}

func (s *GetListRecordResponseBodyData) GetFailedLog() *string {
	return s.FailedLog
}

func (s *GetListRecordResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetListRecordResponseBodyData) SetAnalysisId(v string) *GetListRecordResponseBodyData {
	s.AnalysisId = &v
	return s
}

func (s *GetListRecordResponseBodyData) SetAnalysisTime(v string) *GetListRecordResponseBodyData {
	s.AnalysisTime = &v
	return s
}

func (s *GetListRecordResponseBodyData) SetArguments(v string) *GetListRecordResponseBodyData {
	s.Arguments = &v
	return s
}

func (s *GetListRecordResponseBodyData) SetFailedLog(v string) *GetListRecordResponseBodyData {
	s.FailedLog = &v
	return s
}

func (s *GetListRecordResponseBodyData) SetStatus(v string) *GetListRecordResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetListRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
