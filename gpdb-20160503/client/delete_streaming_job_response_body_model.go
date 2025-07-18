// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStreamingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteStreamingJobResponseBody
	GetRequestId() *string
}

type DeleteStreamingJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStreamingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStreamingJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStreamingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStreamingJobResponseBody) SetRequestId(v string) *DeleteStreamingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStreamingJobResponseBody) Validate() error {
	return dara.Validate(s)
}
