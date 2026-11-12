// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartInstancesResponseBody
	GetRequestId() *string
}

type StartInstancesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartInstancesResponseBody) SetRequestId(v string) *StartInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
