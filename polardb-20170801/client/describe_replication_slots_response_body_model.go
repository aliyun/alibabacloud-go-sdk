// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReplicationSlotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeReplicationSlotsResponseBody
	GetRequestId() *string
	SetSlots(v []*DescribeReplicationSlotsResponseBodySlots) *DescribeReplicationSlotsResponseBody
	GetSlots() []*DescribeReplicationSlotsResponseBodySlots
}

type DescribeReplicationSlotsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of replication slots.
	Slots []*DescribeReplicationSlotsResponseBodySlots `json:"Slots,omitempty" xml:"Slots,omitempty" type:"Repeated"`
}

func (s DescribeReplicationSlotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationSlotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeReplicationSlotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeReplicationSlotsResponseBody) GetSlots() []*DescribeReplicationSlotsResponseBodySlots {
	return s.Slots
}

func (s *DescribeReplicationSlotsResponseBody) SetRequestId(v string) *DescribeReplicationSlotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeReplicationSlotsResponseBody) SetSlots(v []*DescribeReplicationSlotsResponseBodySlots) *DescribeReplicationSlotsResponseBody {
	s.Slots = v
	return s
}

func (s *DescribeReplicationSlotsResponseBody) Validate() error {
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

type DescribeReplicationSlotsResponseBodySlots struct {
	// The ID of the PolarDB cluster to which the replication slot belongs.
	//
	// example:
	//
	// pc-test123456
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database associated with the replication slot.
	//
	// example:
	//
	// testDB
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The name of the output plugin used by the logical replication slot.
	//
	// example:
	//
	// testPlugin
	Plugin *string `json:"Plugin,omitempty" xml:"Plugin,omitempty"`
	// The name of the replication slot.
	//
	// example:
	//
	// testSlot
	SlotName *string `json:"SlotName,omitempty" xml:"SlotName,omitempty"`
	// The status of the replication slot.
	//
	// example:
	//
	// active
	SlotStatus *string `json:"SlotStatus,omitempty" xml:"SlotStatus,omitempty"`
	// The type of the replication slot.
	//
	// example:
	//
	// logical
	SlotType *string `json:"SlotType,omitempty" xml:"SlotType,omitempty"`
	// The name of the child node of the distributed instance.
	//
	// example:
	//
	// testGroup
	SubGroupName *string `json:"SubGroupName,omitempty" xml:"SubGroupName,omitempty"`
	// The replay lag of the subscription.
	//
	// example:
	//
	// 100
	SubReplayLag *string `json:"SubReplayLag,omitempty" xml:"SubReplayLag,omitempty"`
	// Indicates whether the replication slot is temporary. The value is a string: true or false.
	//
	// example:
	//
	// false
	Temporary *string `json:"Temporary,omitempty" xml:"Temporary,omitempty"`
	// The WAL latency.
	//
	// example:
	//
	// 50
	WalDelay *string `json:"WalDelay,omitempty" xml:"WalDelay,omitempty"`
}

func (s DescribeReplicationSlotsResponseBodySlots) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationSlotsResponseBodySlots) GoString() string {
	return s.String()
}

func (s *DescribeReplicationSlotsResponseBodySlots) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeReplicationSlotsResponseBodySlots) GetDatabase() *string {
	return s.Database
}

func (s *DescribeReplicationSlotsResponseBodySlots) GetPlugin() *string {
	return s.Plugin
}

func (s *DescribeReplicationSlotsResponseBodySlots) GetSlotName() *string {
	return s.SlotName
}

func (s *DescribeReplicationSlotsResponseBodySlots) GetSlotStatus() *string {
	return s.SlotStatus
}

func (s *DescribeReplicationSlotsResponseBodySlots) GetSlotType() *string {
	return s.SlotType
}

func (s *DescribeReplicationSlotsResponseBodySlots) GetSubGroupName() *string {
	return s.SubGroupName
}

func (s *DescribeReplicationSlotsResponseBodySlots) GetSubReplayLag() *string {
	return s.SubReplayLag
}

func (s *DescribeReplicationSlotsResponseBodySlots) GetTemporary() *string {
	return s.Temporary
}

func (s *DescribeReplicationSlotsResponseBodySlots) GetWalDelay() *string {
	return s.WalDelay
}

func (s *DescribeReplicationSlotsResponseBodySlots) SetDBClusterId(v string) *DescribeReplicationSlotsResponseBodySlots {
	s.DBClusterId = &v
	return s
}

func (s *DescribeReplicationSlotsResponseBodySlots) SetDatabase(v string) *DescribeReplicationSlotsResponseBodySlots {
	s.Database = &v
	return s
}

func (s *DescribeReplicationSlotsResponseBodySlots) SetPlugin(v string) *DescribeReplicationSlotsResponseBodySlots {
	s.Plugin = &v
	return s
}

func (s *DescribeReplicationSlotsResponseBodySlots) SetSlotName(v string) *DescribeReplicationSlotsResponseBodySlots {
	s.SlotName = &v
	return s
}

func (s *DescribeReplicationSlotsResponseBodySlots) SetSlotStatus(v string) *DescribeReplicationSlotsResponseBodySlots {
	s.SlotStatus = &v
	return s
}

func (s *DescribeReplicationSlotsResponseBodySlots) SetSlotType(v string) *DescribeReplicationSlotsResponseBodySlots {
	s.SlotType = &v
	return s
}

func (s *DescribeReplicationSlotsResponseBodySlots) SetSubGroupName(v string) *DescribeReplicationSlotsResponseBodySlots {
	s.SubGroupName = &v
	return s
}

func (s *DescribeReplicationSlotsResponseBodySlots) SetSubReplayLag(v string) *DescribeReplicationSlotsResponseBodySlots {
	s.SubReplayLag = &v
	return s
}

func (s *DescribeReplicationSlotsResponseBodySlots) SetTemporary(v string) *DescribeReplicationSlotsResponseBodySlots {
	s.Temporary = &v
	return s
}

func (s *DescribeReplicationSlotsResponseBodySlots) SetWalDelay(v string) *DescribeReplicationSlotsResponseBodySlots {
	s.WalDelay = &v
	return s
}

func (s *DescribeReplicationSlotsResponseBodySlots) Validate() error {
	return dara.Validate(s)
}
