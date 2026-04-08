// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneTrafficControlTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CloneTrafficControlTaskResponseBody
	GetRequestId() *string
	SetTrafficControlTaskId(v string) *CloneTrafficControlTaskResponseBody
	GetTrafficControlTaskId() *string
}

type CloneTrafficControlTaskResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	TrafficControlTaskId *string `json:"TrafficControlTaskId,omitempty" xml:"TrafficControlTaskId,omitempty"`
}

func (s CloneTrafficControlTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloneTrafficControlTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CloneTrafficControlTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloneTrafficControlTaskResponseBody) GetTrafficControlTaskId() *string {
	return s.TrafficControlTaskId
}

func (s *CloneTrafficControlTaskResponseBody) SetRequestId(v string) *CloneTrafficControlTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneTrafficControlTaskResponseBody) SetTrafficControlTaskId(v string) *CloneTrafficControlTaskResponseBody {
	s.TrafficControlTaskId = &v
	return s
}

func (s *CloneTrafficControlTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
