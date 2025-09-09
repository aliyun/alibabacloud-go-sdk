// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMaskingProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopMaskingProcessResponseBody
	GetRequestId() *string
}

type StopMaskingProcessResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopMaskingProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopMaskingProcessResponseBody) GoString() string {
	return s.String()
}

func (s *StopMaskingProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopMaskingProcessResponseBody) SetRequestId(v string) *StopMaskingProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopMaskingProcessResponseBody) Validate() error {
	return dara.Validate(s)
}
