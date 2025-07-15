// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopFailoverTestJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopFailoverTestJobResponseBody
	GetRequestId() *string
}

type StopFailoverTestJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C44F62BE-9CE7-4277-B117-69243F3988BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopFailoverTestJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopFailoverTestJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopFailoverTestJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopFailoverTestJobResponseBody) SetRequestId(v string) *StopFailoverTestJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopFailoverTestJobResponseBody) Validate() error {
	return dara.Validate(s)
}
