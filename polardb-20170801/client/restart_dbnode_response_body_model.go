// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestartDBNodeResponseBody
	GetRequestId() *string
}

type RestartDBNodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD86******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartDBNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartDBNodeResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDBNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartDBNodeResponseBody) SetRequestId(v string) *RestartDBNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartDBNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
