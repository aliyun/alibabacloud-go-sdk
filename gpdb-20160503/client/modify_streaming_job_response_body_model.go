// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStreamingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyStreamingJobResponseBody
	GetRequestId() *string
}

type ModifyStreamingJobResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyStreamingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyStreamingJobResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyStreamingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyStreamingJobResponseBody) SetRequestId(v string) *ModifyStreamingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyStreamingJobResponseBody) Validate() error {
	return dara.Validate(s)
}
