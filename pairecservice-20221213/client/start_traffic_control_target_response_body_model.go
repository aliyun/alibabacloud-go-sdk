// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTrafficControlTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartTrafficControlTargetResponseBody
	GetRequestId() *string
}

type StartTrafficControlTargetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartTrafficControlTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartTrafficControlTargetResponseBody) GoString() string {
	return s.String()
}

func (s *StartTrafficControlTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartTrafficControlTargetResponseBody) SetRequestId(v string) *StartTrafficControlTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartTrafficControlTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
