// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoRenderJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SubmitVideoRenderJobResponseBody
	GetErrorCode() *string
	SetJobId(v string) *SubmitVideoRenderJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitVideoRenderJobResponseBody
	GetRequestId() *string
}

type SubmitVideoRenderJobResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 68ca759e798b40b4903b255*******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitVideoRenderJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoRenderJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitVideoRenderJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitVideoRenderJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitVideoRenderJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitVideoRenderJobResponseBody) SetErrorCode(v string) *SubmitVideoRenderJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitVideoRenderJobResponseBody) SetJobId(v string) *SubmitVideoRenderJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitVideoRenderJobResponseBody) SetRequestId(v string) *SubmitVideoRenderJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitVideoRenderJobResponseBody) Validate() error {
	return dara.Validate(s)
}
