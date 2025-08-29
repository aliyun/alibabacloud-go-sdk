// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTrafficControlTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartTrafficControlTaskResponseBody
	GetRequestId() *string
}

type StartTrafficControlTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartTrafficControlTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartTrafficControlTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartTrafficControlTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartTrafficControlTaskResponseBody) SetRequestId(v string) *StartTrafficControlTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartTrafficControlTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
