// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopWuyingServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopWuyingServerResponseBody
	GetRequestId() *string
}

type StopWuyingServerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AD2D0761-1FE5-549D-B169******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopWuyingServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopWuyingServerResponseBody) GoString() string {
	return s.String()
}

func (s *StopWuyingServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopWuyingServerResponseBody) SetRequestId(v string) *StopWuyingServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopWuyingServerResponseBody) Validate() error {
	return dara.Validate(s)
}
