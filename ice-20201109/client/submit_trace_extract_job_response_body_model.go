// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTraceExtractJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitTraceExtractJobResponseBodyData) *SubmitTraceExtractJobResponseBody
	GetData() *SubmitTraceExtractJobResponseBodyData
	SetMessage(v string) *SubmitTraceExtractJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitTraceExtractJobResponseBody
	GetRequestId() *string
	SetStatusCode(v int64) *SubmitTraceExtractJobResponseBody
	GetStatusCode() *int64
}

type SubmitTraceExtractJobResponseBody struct {
	Data *SubmitTraceExtractJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s SubmitTraceExtractJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitTraceExtractJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTraceExtractJobResponseBody) GetData() *SubmitTraceExtractJobResponseBodyData {
	return s.Data
}

func (s *SubmitTraceExtractJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitTraceExtractJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitTraceExtractJobResponseBody) GetStatusCode() *int64 {
	return s.StatusCode
}

func (s *SubmitTraceExtractJobResponseBody) SetData(v *SubmitTraceExtractJobResponseBodyData) *SubmitTraceExtractJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitTraceExtractJobResponseBody) SetMessage(v string) *SubmitTraceExtractJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitTraceExtractJobResponseBody) SetRequestId(v string) *SubmitTraceExtractJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitTraceExtractJobResponseBody) SetStatusCode(v int64) *SubmitTraceExtractJobResponseBody {
	s.StatusCode = &v
	return s
}

func (s *SubmitTraceExtractJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitTraceExtractJobResponseBodyData struct {
	// example:
	//
	// bfb786c639894f4d80648792021e****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitTraceExtractJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitTraceExtractJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitTraceExtractJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *SubmitTraceExtractJobResponseBodyData) SetJobId(v string) *SubmitTraceExtractJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *SubmitTraceExtractJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
