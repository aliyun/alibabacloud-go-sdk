// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopStackGroupOperationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopStackGroupOperationResponseBody
	GetRequestId() *string
}

type StopStackGroupOperationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopStackGroupOperationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopStackGroupOperationResponseBody) GoString() string {
	return s.String()
}

func (s *StopStackGroupOperationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopStackGroupOperationResponseBody) SetRequestId(v string) *StopStackGroupOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopStackGroupOperationResponseBody) Validate() error {
	return dara.Validate(s)
}
