// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopApmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopApmResponseBody
	GetRequestId() *string
	SetResult(v bool) *StopApmResponseBody
	GetResult() *bool
}

type StopApmResponseBody struct {
	// example:
	//
	// FEC32FE6-4697-5110-9668-C6016EAEB5DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s StopApmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopApmResponseBody) GoString() string {
	return s.String()
}

func (s *StopApmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopApmResponseBody) GetResult() *bool {
	return s.Result
}

func (s *StopApmResponseBody) SetRequestId(v string) *StopApmResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopApmResponseBody) SetResult(v bool) *StopApmResponseBody {
	s.Result = &v
	return s
}

func (s *StopApmResponseBody) Validate() error {
	return dara.Validate(s)
}
