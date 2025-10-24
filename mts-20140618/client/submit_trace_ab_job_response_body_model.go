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
	Data *SubmitTraceAbJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 030E2671-806A-52AF-A93C-DA8E308603A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// bfb786c639894f4d80648792021e****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 437bd2b516ffda105d07b12a9a82****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
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

func (s *SubmitTraceAbJobResponseBodyData) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitTraceAbJobResponseBodyData) SetJobId(v string) *SubmitTraceAbJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *SubmitTraceAbJobResponseBodyData) SetMediaId(v string) *SubmitTraceAbJobResponseBodyData {
	s.MediaId = &v
	return s
}

func (s *SubmitTraceAbJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
