// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopAppResponseBody
	GetRequestId() *string
}

type StopAppResponseBody struct {
	// example:
	//
	// 83A9075B-C646-59A9-8232-CAE41AF4B9E3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopAppResponseBody) GoString() string {
	return s.String()
}

func (s *StopAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopAppResponseBody) SetRequestId(v string) *StopAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopAppResponseBody) Validate() error {
	return dara.Validate(s)
}
