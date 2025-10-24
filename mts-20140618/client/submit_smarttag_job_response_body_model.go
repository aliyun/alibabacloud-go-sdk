// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmarttagJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitSmarttagJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitSmarttagJobResponseBody
	GetRequestId() *string
}

type SubmitSmarttagJobResponseBody struct {
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 7B117AF5-2A16-412C-B127-FA6175ED1AD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSmarttagJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmarttagJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSmarttagJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitSmarttagJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSmarttagJobResponseBody) SetJobId(v string) *SubmitSmarttagJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitSmarttagJobResponseBody) SetRequestId(v string) *SubmitSmarttagJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSmarttagJobResponseBody) Validate() error {
	return dara.Validate(s)
}
