// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficControlTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTrafficControlTaskResponseBody
	GetRequestId() *string
	SetTrafficControlTaskId(v string) *CreateTrafficControlTaskResponseBody
	GetTrafficControlTaskId() *string
}

type CreateTrafficControlTaskResponseBody struct {
	// example:
	//
	// 42391E6D-822C-58F8-9F7E-D991BB86D6AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TrafficControlTaskId *string `json:"TrafficControlTaskId,omitempty" xml:"TrafficControlTaskId,omitempty"`
}

func (s CreateTrafficControlTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficControlTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrafficControlTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTrafficControlTaskResponseBody) GetTrafficControlTaskId() *string {
	return s.TrafficControlTaskId
}

func (s *CreateTrafficControlTaskResponseBody) SetRequestId(v string) *CreateTrafficControlTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrafficControlTaskResponseBody) SetTrafficControlTaskId(v string) *CreateTrafficControlTaskResponseBody {
	s.TrafficControlTaskId = &v
	return s
}

func (s *CreateTrafficControlTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
