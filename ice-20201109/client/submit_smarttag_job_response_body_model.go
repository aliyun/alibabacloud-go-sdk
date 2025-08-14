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
	// The ID of the smart tagging job. We recommend that you save this ID for subsequent calls of other operations.
	//
	// example:
	//
	// ****d80e4e4044975745c14b****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
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
