// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSampleConsistencyJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopSampleConsistencyJobResponseBody
	GetRequestId() *string
}

type StopSampleConsistencyJobResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopSampleConsistencyJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopSampleConsistencyJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopSampleConsistencyJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopSampleConsistencyJobResponseBody) SetRequestId(v string) *StopSampleConsistencyJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopSampleConsistencyJobResponseBody) Validate() error {
	return dara.Validate(s)
}
