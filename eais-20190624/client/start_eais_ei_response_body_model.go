// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartEaisEiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartEaisEiResponseBody
	GetRequestId() *string
}

type StartEaisEiResponseBody struct {
	// example:
	//
	// F23AEEC7-4D98-4657-A104-02692701****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartEaisEiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartEaisEiResponseBody) GoString() string {
	return s.String()
}

func (s *StartEaisEiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartEaisEiResponseBody) SetRequestId(v string) *StartEaisEiResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartEaisEiResponseBody) Validate() error {
	return dara.Validate(s)
}
