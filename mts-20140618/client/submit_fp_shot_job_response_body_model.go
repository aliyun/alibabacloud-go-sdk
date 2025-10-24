// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitFpShotJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitFpShotJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitFpShotJobResponseBody
	GetRequestId() *string
}

type SubmitFpShotJobResponseBody struct {
	// The ID of the media fingerprint analysis job. We recommend that you keep this ID for subsequent operation calls.
	//
	// example:
	//
	// 2a0697e35a7342859f733a9190c4****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitFpShotJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitFpShotJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitFpShotJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitFpShotJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitFpShotJobResponseBody) SetJobId(v string) *SubmitFpShotJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitFpShotJobResponseBody) SetRequestId(v string) *SubmitFpShotJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitFpShotJobResponseBody) Validate() error {
	return dara.Validate(s)
}
