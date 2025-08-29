// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficControlTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTrafficControlTargetResponseBody
	GetRequestId() *string
	SetTrafficControlTargetId(v string) *CreateTrafficControlTargetResponseBody
	GetTrafficControlTargetId() *string
}

type CreateTrafficControlTargetResponseBody struct {
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrafficControlTargetId *string `json:"TrafficControlTargetId,omitempty" xml:"TrafficControlTargetId,omitempty"`
}

func (s CreateTrafficControlTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficControlTargetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrafficControlTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTrafficControlTargetResponseBody) GetTrafficControlTargetId() *string {
	return s.TrafficControlTargetId
}

func (s *CreateTrafficControlTargetResponseBody) SetRequestId(v string) *CreateTrafficControlTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrafficControlTargetResponseBody) SetTrafficControlTargetId(v string) *CreateTrafficControlTargetResponseBody {
	s.TrafficControlTargetId = &v
	return s
}

func (s *CreateTrafficControlTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
