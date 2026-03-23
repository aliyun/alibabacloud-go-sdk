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
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Slots     []*DescribeSlotsResponseBodySlots `json:"Slots,omitempty" xml:"Slots,omitempty" type:"Repeated"`
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
	Database     *string `json:"Database,omitempty" xml:"Database,omitempty"`
	Plugin       *string `json:"Plugin,omitempty" xml:"Plugin,omitempty"`
	SlotName     *string `json:"SlotName,omitempty" xml:"SlotName,omitempty"`
	SlotStatus   *string `json:"SlotStatus,omitempty" xml:"SlotStatus,omitempty"`
	SlotType     *string `json:"SlotType,omitempty" xml:"SlotType,omitempty"`
	SubReplayLag *string `json:"SubReplayLag,omitempty" xml:"SubReplayLag,omitempty"`
	Temporary    *string `json:"Temporary,omitempty" xml:"Temporary,omitempty"`
	WalDelay     *string `json:"WalDelay,omitempty" xml:"WalDelay,omitempty"`
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
