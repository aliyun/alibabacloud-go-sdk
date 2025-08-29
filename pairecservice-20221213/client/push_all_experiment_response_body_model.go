// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushAllExperimentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PushAllExperimentResponseBody
	GetRequestId() *string
}

type PushAllExperimentResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 74D958EF-3598-56FA-8296-FF1575CE43DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushAllExperimentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushAllExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *PushAllExperimentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushAllExperimentResponseBody) SetRequestId(v string) *PushAllExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushAllExperimentResponseBody) Validate() error {
	return dara.Validate(s)
}
