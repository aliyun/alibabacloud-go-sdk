// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikeAIAppJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitYikeAIAppJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitYikeAIAppJobResponseBody
	GetRequestId() *string
}

type SubmitYikeAIAppJobResponseBody struct {
	// example:
	//
	// ab4802364a2e49208c99efab82df****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 4E84BE44-58A7-****-****-FBEBEA16EF94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitYikeAIAppJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikeAIAppJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitYikeAIAppJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitYikeAIAppJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitYikeAIAppJobResponseBody) SetJobId(v string) *SubmitYikeAIAppJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitYikeAIAppJobResponseBody) SetRequestId(v string) *SubmitYikeAIAppJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitYikeAIAppJobResponseBody) Validate() error {
	return dara.Validate(s)
}
