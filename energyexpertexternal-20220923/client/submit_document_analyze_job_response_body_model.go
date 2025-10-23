// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocumentAnalyzeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitDocumentAnalyzeJobResponseBodyData) *SubmitDocumentAnalyzeJobResponseBody
	GetData() *SubmitDocumentAnalyzeJobResponseBodyData
	SetRequestId(v string) *SubmitDocumentAnalyzeJobResponseBody
	GetRequestId() *string
}

type SubmitDocumentAnalyzeJobResponseBody struct {
	// The data returned.
	Data *SubmitDocumentAnalyzeJobResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 4A0AEC56-5C9A-5D47-93DF-7227836FFF82
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SubmitDocumentAnalyzeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocumentAnalyzeJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDocumentAnalyzeJobResponseBody) GetData() *SubmitDocumentAnalyzeJobResponseBodyData {
	return s.Data
}

func (s *SubmitDocumentAnalyzeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitDocumentAnalyzeJobResponseBody) SetData(v *SubmitDocumentAnalyzeJobResponseBodyData) *SubmitDocumentAnalyzeJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitDocumentAnalyzeJobResponseBody) SetRequestId(v string) *SubmitDocumentAnalyzeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDocumentAnalyzeJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitDocumentAnalyzeJobResponseBodyData struct {
	// The job ID.
	//
	// example:
	//
	// adkc-kk2k41-kk2ol-222424
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
}

func (s SubmitDocumentAnalyzeJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocumentAnalyzeJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitDocumentAnalyzeJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *SubmitDocumentAnalyzeJobResponseBodyData) SetJobId(v string) *SubmitDocumentAnalyzeJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *SubmitDocumentAnalyzeJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
