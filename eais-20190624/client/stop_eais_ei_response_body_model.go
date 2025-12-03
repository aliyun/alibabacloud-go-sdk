// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopEaisEiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopEaisEiResponseBody
	GetRequestId() *string
}

type StopEaisEiResponseBody struct {
	// example:
	//
	// F23AEEC7-4D98-4657-A104-02692701****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopEaisEiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopEaisEiResponseBody) GoString() string {
	return s.String()
}

func (s *StopEaisEiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopEaisEiResponseBody) SetRequestId(v string) *StopEaisEiResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopEaisEiResponseBody) Validate() error {
	return dara.Validate(s)
}
