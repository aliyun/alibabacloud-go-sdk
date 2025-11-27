// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSlotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSlotResponseBody
	GetRequestId() *string
	SetSlotName(v string) *DeleteSlotResponseBody
	GetSlotName() *string
}

type DeleteSlotResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2875D608-A228-53D7-B8C9-35F13EDCF36D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the replication slot.
	//
	// example:
	//
	// slot_test01
	SlotName *string `json:"SlotName,omitempty" xml:"SlotName,omitempty"`
}

func (s DeleteSlotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSlotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSlotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSlotResponseBody) GetSlotName() *string {
	return s.SlotName
}

func (s *DeleteSlotResponseBody) SetRequestId(v string) *DeleteSlotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSlotResponseBody) SetSlotName(v string) *DeleteSlotResponseBody {
	s.SlotName = &v
	return s
}

func (s *DeleteSlotResponseBody) Validate() error {
	return dara.Validate(s)
}
