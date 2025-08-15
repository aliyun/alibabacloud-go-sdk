// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLoggingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopLoggingResponseBody
	GetRequestId() *string
}

type StopLoggingResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1C488B66-B819-4D14-8711-C4EAAA13AC01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopLoggingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopLoggingResponseBody) GoString() string {
	return s.String()
}

func (s *StopLoggingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopLoggingResponseBody) SetRequestId(v string) *StopLoggingResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopLoggingResponseBody) Validate() error {
	return dara.Validate(s)
}
