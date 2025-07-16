// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopTaskResponseBody
	GetRequestId() *string
}

type StopTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 514F82AF-3C04-5C3D-8F38-A11261BF37B0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StopTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopTaskResponseBody) SetRequestId(v string) *StopTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
