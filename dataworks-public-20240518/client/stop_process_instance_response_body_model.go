// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopProcessInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopProcessInstanceResponseBody
	GetRequestId() *string
}

type StopProcessInstanceResponseBody struct {
	// example:
	//
	// 0bc5df3a17****903790e8e8a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopProcessInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopProcessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopProcessInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopProcessInstanceResponseBody) SetRequestId(v string) *StopProcessInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopProcessInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
