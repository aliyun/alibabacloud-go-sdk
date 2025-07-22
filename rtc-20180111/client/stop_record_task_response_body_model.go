// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRecordTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopRecordTaskResponseBody
	GetRequestId() *string
}

type StopRecordTaskResponseBody struct {
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopRecordTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopRecordTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopRecordTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopRecordTaskResponseBody) SetRequestId(v string) *StopRecordTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopRecordTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
