// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSlotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointIds(v string) *CreateSlotResponseBody
	GetEndpointIds() *string
	SetRequestId(v string) *CreateSlotResponseBody
	GetRequestId() *string
	SetSlotId(v string) *CreateSlotResponseBody
	GetSlotId() *string
}

type CreateSlotResponseBody struct {
	// example:
	//
	// end-5zk866779me51jgu3w,end-5zk866779me51jgu3w
	EndpointIds *string `json:"EndpointIds,omitempty" xml:"EndpointIds,omitempty"`
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// slot-5zk866779me51jgu3w
	SlotId *string `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
}

func (s CreateSlotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSlotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSlotResponseBody) GetEndpointIds() *string {
	return s.EndpointIds
}

func (s *CreateSlotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSlotResponseBody) GetSlotId() *string {
	return s.SlotId
}

func (s *CreateSlotResponseBody) SetEndpointIds(v string) *CreateSlotResponseBody {
	s.EndpointIds = &v
	return s
}

func (s *CreateSlotResponseBody) SetRequestId(v string) *CreateSlotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSlotResponseBody) SetSlotId(v string) *CreateSlotResponseBody {
	s.SlotId = &v
	return s
}

func (s *CreateSlotResponseBody) Validate() error {
	return dara.Validate(s)
}
