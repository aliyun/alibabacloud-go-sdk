// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTransferStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopTransferStreamResponseBody
	GetRequestId() *string
}

type StopTransferStreamResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopTransferStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopTransferStreamResponseBody) GoString() string {
	return s.String()
}

func (s *StopTransferStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopTransferStreamResponseBody) SetRequestId(v string) *StopTransferStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTransferStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
