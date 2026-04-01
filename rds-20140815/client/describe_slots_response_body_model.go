// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSlotsResponseBody
	GetRequestId() *string
	SetSlots(v []*DescribeSlotsResponseBodySlots) *DescribeSlotsResponseBody
	GetSlots() []*DescribeSlotsResponseBodySlots
}

type DescribeSlotsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 76AF0609-4195-5DFC-BC78-3AD76FF872BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the replication slot.
	Slots []*DescribeSlotsResponseBodySlots `json:"Slots,omitempty" xml:"Slots,omitempty" type:"Repeated"`
}

func (s DescribeSlotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlotsResponseBody) GetSlots() []*DescribeSlotsResponseBodySlots {
	return s.Slots
}

func (s *DescribeSlotsResponseBody) SetRequestId(v string) *DescribeSlotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlotsResponseBody) SetSlots(v []*DescribeSlotsResponseBodySlots) *DescribeSlotsResponseBody {
	s.Slots = v
	return s
}

func (s *DescribeSlotsResponseBody) Validate() error {
	if s.Slots != nil {
		for _, item := range s.Slots {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSlotsResponseBodySlots struct {
	// The name of the database in which the replication slot resides.
	//
	// example:
	//
	// db_test01
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The extension used by the replication slot.
	//
	// example:
	//
	// test_decoding
	Plugin *string `json:"Plugin,omitempty" xml:"Plugin,omitempty"`
	// The replication slot name.
	//
	// example:
	//
	// slot_test01
	SlotName *string `json:"SlotName,omitempty" xml:"SlotName,omitempty"`
	// The replication slot status. Valid values:
	//
	// 	- ACTIVE
	//
	// 	- INACTIVE
	//
	// example:
	//
	// INACTIVE
	SlotStatus *string `json:"SlotStatus,omitempty" xml:"SlotStatus,omitempty"`
	// The replication slot type. Valid values:
	//
	// 	- physical
	//
	// 	- logical
	//
	// example:
	//
	// logical
	SlotType *string `json:"SlotType,omitempty" xml:"SlotType,omitempty"`
	// The latency of the logical subscription on the subscriber node that corresponds to the current replication slot. Unit: seconds.
	//
	// example:
	//
	// 0
	SubReplayLag *string `json:"SubReplayLag,omitempty" xml:"SubReplayLag,omitempty"`
	// Indicates whether the replication slot is a temporary replication slot. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Temporary *string `json:"Temporary,omitempty" xml:"Temporary,omitempty"`
	// The number of logs accumulated in the replication slot.
	//
	// example:
	//
	// 16 MB
	WalDelay *string `json:"WalDelay,omitempty" xml:"WalDelay,omitempty"`
}

func (s DescribeSlotsResponseBodySlots) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlotsResponseBodySlots) GoString() string {
	return s.String()
}

func (s *DescribeSlotsResponseBodySlots) GetDatabase() *string {
	return s.Database
}

func (s *DescribeSlotsResponseBodySlots) GetPlugin() *string {
	return s.Plugin
}

func (s *DescribeSlotsResponseBodySlots) GetSlotName() *string {
	return s.SlotName
}

func (s *DescribeSlotsResponseBodySlots) GetSlotStatus() *string {
	return s.SlotStatus
}

func (s *DescribeSlotsResponseBodySlots) GetSlotType() *string {
	return s.SlotType
}

func (s *DescribeSlotsResponseBodySlots) GetSubReplayLag() *string {
	return s.SubReplayLag
}

func (s *DescribeSlotsResponseBodySlots) GetTemporary() *string {
	return s.Temporary
}

func (s *DescribeSlotsResponseBodySlots) GetWalDelay() *string {
	return s.WalDelay
}

func (s *DescribeSlotsResponseBodySlots) SetDatabase(v string) *DescribeSlotsResponseBodySlots {
	s.Database = &v
	return s
}

func (s *DescribeSlotsResponseBodySlots) SetPlugin(v string) *DescribeSlotsResponseBodySlots {
	s.Plugin = &v
	return s
}

func (s *DescribeSlotsResponseBodySlots) SetSlotName(v string) *DescribeSlotsResponseBodySlots {
	s.SlotName = &v
	return s
}

func (s *DescribeSlotsResponseBodySlots) SetSlotStatus(v string) *DescribeSlotsResponseBodySlots {
	s.SlotStatus = &v
	return s
}

func (s *DescribeSlotsResponseBodySlots) SetSlotType(v string) *DescribeSlotsResponseBodySlots {
	s.SlotType = &v
	return s
}

func (s *DescribeSlotsResponseBodySlots) SetSubReplayLag(v string) *DescribeSlotsResponseBodySlots {
	s.SubReplayLag = &v
	return s
}

func (s *DescribeSlotsResponseBodySlots) SetTemporary(v string) *DescribeSlotsResponseBodySlots {
	s.Temporary = &v
	return s
}

func (s *DescribeSlotsResponseBodySlots) SetWalDelay(v string) *DescribeSlotsResponseBodySlots {
	s.WalDelay = &v
	return s
}

func (s *DescribeSlotsResponseBodySlots) Validate() error {
	return dara.Validate(s)
}
