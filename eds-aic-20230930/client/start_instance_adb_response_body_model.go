// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstanceAdbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartInstanceAdbResponseBody
	GetRequestId() *string
}

type StartInstanceAdbResponseBody struct {
	// example:
	//
	// 69BCBBE4-FCF2-59B8-AD9D-531EB422****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartInstanceAdbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceAdbResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceAdbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartInstanceAdbResponseBody) SetRequestId(v string) *StartInstanceAdbResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartInstanceAdbResponseBody) Validate() error {
	return dara.Validate(s)
}
