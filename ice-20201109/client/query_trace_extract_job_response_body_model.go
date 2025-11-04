// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTraceExtractJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryTraceExtractJobResponseBodyData) *QueryTraceExtractJobResponseBody
	GetData() *QueryTraceExtractJobResponseBodyData
	SetMessage(v string) *QueryTraceExtractJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryTraceExtractJobResponseBody
	GetRequestId() *string
	SetStatusCode(v int64) *QueryTraceExtractJobResponseBody
	GetStatusCode() *int64
}

type QueryTraceExtractJobResponseBody struct {
	// The data returned.
	Data *QueryTraceExtractJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// *****ACB-44F2-5F2D-88D7-1283E70*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code.
	//
	// example:
	//
	// 200
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QueryTraceExtractJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceExtractJobResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTraceExtractJobResponseBody) GetData() *QueryTraceExtractJobResponseBodyData {
	return s.Data
}

func (s *QueryTraceExtractJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryTraceExtractJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTraceExtractJobResponseBody) GetStatusCode() *int64 {
	return s.StatusCode
}

func (s *QueryTraceExtractJobResponseBody) SetData(v *QueryTraceExtractJobResponseBodyData) *QueryTraceExtractJobResponseBody {
	s.Data = v
	return s
}

func (s *QueryTraceExtractJobResponseBody) SetMessage(v string) *QueryTraceExtractJobResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTraceExtractJobResponseBody) SetRequestId(v string) *QueryTraceExtractJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTraceExtractJobResponseBody) SetStatusCode(v int64) *QueryTraceExtractJobResponseBody {
	s.StatusCode = &v
	return s
}

func (s *QueryTraceExtractJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTraceExtractJobResponseBodyData struct {
	// The trace watermark information.
	Trace *string `json:"Trace,omitempty" xml:"Trace,omitempty"`
}

func (s QueryTraceExtractJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceExtractJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTraceExtractJobResponseBodyData) GetTrace() *string {
	return s.Trace
}

func (s *QueryTraceExtractJobResponseBodyData) SetTrace(v string) *QueryTraceExtractJobResponseBodyData {
	s.Trace = &v
	return s
}

func (s *QueryTraceExtractJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
