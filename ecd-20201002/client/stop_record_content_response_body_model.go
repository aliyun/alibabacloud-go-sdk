// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRecordContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopRecordContentResponseBody
	GetRequestId() *string
}

type StopRecordContentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopRecordContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopRecordContentResponseBody) GoString() string {
	return s.String()
}

func (s *StopRecordContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopRecordContentResponseBody) SetRequestId(v string) *StopRecordContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopRecordContentResponseBody) Validate() error {
	return dara.Validate(s)
}
