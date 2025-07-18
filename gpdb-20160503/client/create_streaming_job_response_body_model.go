// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int32) *CreateStreamingJobResponseBody
	GetJobId() *int32
	SetRequestId(v string) *CreateStreamingJobResponseBody
	GetRequestId() *string
}

type CreateStreamingJobResponseBody struct {
	// Job ID.
	//
	// example:
	//
	// 1
	JobId *int32 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateStreamingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamingJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStreamingJobResponseBody) GetJobId() *int32 {
	return s.JobId
}

func (s *CreateStreamingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStreamingJobResponseBody) SetJobId(v int32) *CreateStreamingJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateStreamingJobResponseBody) SetRequestId(v string) *CreateStreamingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStreamingJobResponseBody) Validate() error {
	return dara.Validate(s)
}
