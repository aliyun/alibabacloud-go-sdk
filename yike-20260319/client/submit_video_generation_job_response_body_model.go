// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoGenerationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitVideoGenerationJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitVideoGenerationJobResponseBody
	GetRequestId() *string
}

type SubmitVideoGenerationJobResponseBody struct {
	// The task ID.
	//
	// example:
	//
	// ****3e761e9d11edba640c42a1b7****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitVideoGenerationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoGenerationJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitVideoGenerationJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitVideoGenerationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitVideoGenerationJobResponseBody) SetJobId(v string) *SubmitVideoGenerationJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitVideoGenerationJobResponseBody) SetRequestId(v string) *SubmitVideoGenerationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitVideoGenerationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
