// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTrafficControlFlinkTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopTrafficControlFlinkTaskResponseBody
	GetRequestId() *string
}

type StopTrafficControlFlinkTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopTrafficControlFlinkTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopTrafficControlFlinkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopTrafficControlFlinkTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopTrafficControlFlinkTaskResponseBody) SetRequestId(v string) *StopTrafficControlFlinkTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTrafficControlFlinkTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
