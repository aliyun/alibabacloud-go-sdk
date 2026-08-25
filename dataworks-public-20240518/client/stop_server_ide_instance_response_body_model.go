// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopServerIdeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StopServerIdeInstanceResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *StopServerIdeInstanceResponseBody
	GetRequestId() *string
}

type StopServerIdeInstanceResponseBody struct {
	// The ID of the personal development environment instance.
	//
	// example:
	//
	// 699573
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopServerIdeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopServerIdeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopServerIdeInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopServerIdeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopServerIdeInstanceResponseBody) SetInstanceId(v string) *StopServerIdeInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *StopServerIdeInstanceResponseBody) SetRequestId(v string) *StopServerIdeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopServerIdeInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
