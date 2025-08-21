// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRecordStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopRecordStreamResponseBody
	GetRequestId() *string
}

type StopRecordStreamResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopRecordStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopRecordStreamResponseBody) GoString() string {
	return s.String()
}

func (s *StopRecordStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopRecordStreamResponseBody) SetRequestId(v string) *StopRecordStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopRecordStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
