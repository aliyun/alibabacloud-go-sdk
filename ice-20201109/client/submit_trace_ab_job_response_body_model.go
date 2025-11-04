// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTraceAbJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitTraceAbJobResponseBodyData) *SubmitTraceAbJobResponseBody
	GetData() *SubmitTraceAbJobResponseBodyData
	SetMessage(v string) *SubmitTraceAbJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitTraceAbJobResponseBody
	GetRequestId() *string
	SetStatusCode(v int64) *SubmitTraceAbJobResponseBody
	GetStatusCode() *int64
}

type SubmitTraceAbJobResponseBody struct {
	// The data returned.
	Data *SubmitTraceAbJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******36-3C1E-4417-BDB2-1E034F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code.
	//
	// example:
	//
	// 200
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s SubmitTraceAbJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitTraceAbJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTraceAbJobResponseBody) GetData() *SubmitTraceAbJobResponseBodyData {
	return s.Data
}

func (s *SubmitTraceAbJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitTraceAbJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitTraceAbJobResponseBody) GetStatusCode() *int64 {
	return s.StatusCode
}

func (s *SubmitTraceAbJobResponseBody) SetData(v *SubmitTraceAbJobResponseBodyData) *SubmitTraceAbJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitTraceAbJobResponseBody) SetMessage(v string) *SubmitTraceAbJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitTraceAbJobResponseBody) SetRequestId(v string) *SubmitTraceAbJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitTraceAbJobResponseBody) SetStatusCode(v int64) *SubmitTraceAbJobResponseBody {
	s.StatusCode = &v
	return s
}

func (s *SubmitTraceAbJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTraceAbJobResponseBodyData struct {
	// The job ID.
	//
	// example:
	//
	// bfb786c639894f4d80648792021e****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The media ID.
	//
	// example:
	//
	// bf53333264f4d80648792021e****
	TraceMediaId *string `json:"TraceMediaId,omitempty" xml:"TraceMediaId,omitempty"`
}

func (s SubmitTraceAbJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitTraceAbJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitTraceAbJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *SubmitTraceAbJobResponseBodyData) GetTraceMediaId() *string {
	return s.TraceMediaId
}

func (s *SubmitTraceAbJobResponseBodyData) SetJobId(v string) *SubmitTraceAbJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *SubmitTraceAbJobResponseBodyData) SetTraceMediaId(v string) *SubmitTraceAbJobResponseBodyData {
	s.TraceMediaId = &v
	return s
}

func (s *SubmitTraceAbJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
