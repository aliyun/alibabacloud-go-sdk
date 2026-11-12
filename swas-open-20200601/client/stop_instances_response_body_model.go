// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopInstancesResponseBody
	GetRequestId() *string
}

type StopInstancesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C2DE174B-7196-5778-A00D-6EA2601B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopInstancesResponseBody) SetRequestId(v string) *StopInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
