// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRCInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartRCInstancesResponseBody
	GetRequestId() *string
}

type StartRCInstancesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 847BA085-B377-4BFA-8267-F82345ECE1D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartRCInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartRCInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *StartRCInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartRCInstancesResponseBody) SetRequestId(v string) *StartRCInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartRCInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
