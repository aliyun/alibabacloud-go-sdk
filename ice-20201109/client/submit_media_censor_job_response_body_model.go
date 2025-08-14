// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaCensorJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitMediaCensorJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitMediaCensorJobResponseBody
	GetRequestId() *string
}

type SubmitMediaCensorJobResponseBody struct {
	// The ID of the content moderation job. We recommend that you save this ID for subsequent calls of other operations.
	//
	// example:
	//
	// 88c6ca184c0e47098a5b665e2****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitMediaCensorJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaCensorJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitMediaCensorJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitMediaCensorJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitMediaCensorJobResponseBody) SetJobId(v string) *SubmitMediaCensorJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitMediaCensorJobResponseBody) SetRequestId(v string) *SubmitMediaCensorJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitMediaCensorJobResponseBody) Validate() error {
	return dara.Validate(s)
}
