// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReadDBInstanceDelayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeReadDBInstanceDelayResponseBody
	GetDBInstanceId() *string
	SetDelayTime(v int32) *DescribeReadDBInstanceDelayResponseBody
	GetDelayTime() *int32
	SetItems(v *DescribeReadDBInstanceDelayResponseBodyItems) *DescribeReadDBInstanceDelayResponseBody
	GetItems() *DescribeReadDBInstanceDelayResponseBodyItems
	SetReadDBInstanceId(v string) *DescribeReadDBInstanceDelayResponseBody
	GetReadDBInstanceId() *string
	SetRequestId(v string) *DescribeReadDBInstanceDelayResponseBody
	GetRequestId() *string
}

type DescribeReadDBInstanceDelayResponseBody struct {
	// The primary instance ID.
	//
	// example:
	//
	// rm-bp*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The latency of data replication. Unit: seconds.
	//
	// example:
	//
	// 0
	DelayTime *int32                                        `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	Items     *DescribeReadDBInstanceDelayResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The read-only instance ID.
	//
	// example:
	//
	// rr-bp*****
	ReadDBInstanceId *string `json:"ReadDBInstanceId,omitempty" xml:"ReadDBInstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F1BDDEA8-452D-450B-AB10-CD5C5BAFC5DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeReadDBInstanceDelayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeReadDBInstanceDelayResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeReadDBInstanceDelayResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeReadDBInstanceDelayResponseBody) GetDelayTime() *int32 {
	return s.DelayTime
}

func (s *DescribeReadDBInstanceDelayResponseBody) GetItems() *DescribeReadDBInstanceDelayResponseBodyItems {
	return s.Items
}

func (s *DescribeReadDBInstanceDelayResponseBody) GetReadDBInstanceId() *string {
	return s.ReadDBInstanceId
}

func (s *DescribeReadDBInstanceDelayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeReadDBInstanceDelayResponseBody) SetDBInstanceId(v string) *DescribeReadDBInstanceDelayResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBody) SetDelayTime(v int32) *DescribeReadDBInstanceDelayResponseBody {
	s.DelayTime = &v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBody) SetItems(v *DescribeReadDBInstanceDelayResponseBodyItems) *DescribeReadDBInstanceDelayResponseBody {
	s.Items = v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBody) SetReadDBInstanceId(v string) *DescribeReadDBInstanceDelayResponseBody {
	s.ReadDBInstanceId = &v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBody) SetRequestId(v string) *DescribeReadDBInstanceDelayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeReadDBInstanceDelayResponseBodyItems struct {
	Items []*DescribeReadDBInstanceDelayResponseBodyItemsItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeReadDBInstanceDelayResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeReadDBInstanceDelayResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeReadDBInstanceDelayResponseBodyItems) GetItems() []*DescribeReadDBInstanceDelayResponseBodyItemsItems {
	return s.Items
}

func (s *DescribeReadDBInstanceDelayResponseBodyItems) SetItems(v []*DescribeReadDBInstanceDelayResponseBodyItemsItems) *DescribeReadDBInstanceDelayResponseBodyItems {
	s.Items = v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBodyItems) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeReadDBInstanceDelayResponseBodyItemsItems struct {
	DBInstanceId          *string                                                                 `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ReadDBInstanceNames   *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadDBInstanceNames   `json:"ReadDBInstanceNames,omitempty" xml:"ReadDBInstanceNames,omitempty" type:"Struct"`
	ReadDelayTimes        *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadDelayTimes        `json:"ReadDelayTimes,omitempty" xml:"ReadDelayTimes,omitempty" type:"Struct"`
	ReadonlyInstanceDelay *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelay `json:"ReadonlyInstanceDelay,omitempty" xml:"ReadonlyInstanceDelay,omitempty" type:"Struct"`
}

func (s DescribeReadDBInstanceDelayResponseBodyItemsItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeReadDBInstanceDelayResponseBodyItemsItems) GoString() string {
	return s.String()
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItems) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItems) GetReadDBInstanceNames() *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadDBInstanceNames {
	return s.ReadDBInstanceNames
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItems) GetReadDelayTimes() *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadDelayTimes {
	return s.ReadDelayTimes
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItems) GetReadonlyInstanceDelay() *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelay {
	return s.ReadonlyInstanceDelay
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItems) SetDBInstanceId(v string) *DescribeReadDBInstanceDelayResponseBodyItemsItems {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItems) SetReadDBInstanceNames(v *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadDBInstanceNames) *DescribeReadDBInstanceDelayResponseBodyItemsItems {
	s.ReadDBInstanceNames = v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItems) SetReadDelayTimes(v *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadDelayTimes) *DescribeReadDBInstanceDelayResponseBodyItemsItems {
	s.ReadDelayTimes = v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItems) SetReadonlyInstanceDelay(v *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelay) *DescribeReadDBInstanceDelayResponseBodyItemsItems {
	s.ReadonlyInstanceDelay = v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItems) Validate() error {
	if s.ReadDBInstanceNames != nil {
		if err := s.ReadDBInstanceNames.Validate(); err != nil {
			return err
		}
	}
	if s.ReadDelayTimes != nil {
		if err := s.ReadDelayTimes.Validate(); err != nil {
			return err
		}
	}
	if s.ReadonlyInstanceDelay != nil {
		if err := s.ReadonlyInstanceDelay.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeReadDBInstanceDelayResponseBodyItemsItemsReadDBInstanceNames struct {
	ReadDBInstanceName []*string `json:"ReadDBInstanceName,omitempty" xml:"ReadDBInstanceName,omitempty" type:"Repeated"`
}

func (s DescribeReadDBInstanceDelayResponseBodyItemsItemsReadDBInstanceNames) String() string {
	return dara.Prettify(s)
}

func (s DescribeReadDBInstanceDelayResponseBodyItemsItemsReadDBInstanceNames) GoString() string {
	return s.String()
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadDBInstanceNames) GetReadDBInstanceName() []*string {
	return s.ReadDBInstanceName
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadDBInstanceNames) SetReadDBInstanceName(v []*string) *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadDBInstanceNames {
	s.ReadDBInstanceName = v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadDBInstanceNames) Validate() error {
	return dara.Validate(s)
}

type DescribeReadDBInstanceDelayResponseBodyItemsItemsReadDelayTimes struct {
	ReadDelayTime []*string `json:"ReadDelayTime,omitempty" xml:"ReadDelayTime,omitempty" type:"Repeated"`
}

func (s DescribeReadDBInstanceDelayResponseBodyItemsItemsReadDelayTimes) String() string {
	return dara.Prettify(s)
}

func (s DescribeReadDBInstanceDelayResponseBodyItemsItemsReadDelayTimes) GoString() string {
	return s.String()
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadDelayTimes) GetReadDelayTime() []*string {
	return s.ReadDelayTime
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadDelayTimes) SetReadDelayTime(v []*string) *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadDelayTimes {
	s.ReadDelayTime = v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadDelayTimes) Validate() error {
	return dara.Validate(s)
}

type DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelay struct {
	ReadonlyInstanceDelay []*DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay `json:"ReadonlyInstanceDelay,omitempty" xml:"ReadonlyInstanceDelay,omitempty" type:"Repeated"`
}

func (s DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelay) String() string {
	return dara.Prettify(s)
}

func (s DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelay) GoString() string {
	return s.String()
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelay) GetReadonlyInstanceDelay() []*DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay {
	return s.ReadonlyInstanceDelay
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelay) SetReadonlyInstanceDelay(v []*DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay) *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelay {
	s.ReadonlyInstanceDelay = v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelay) Validate() error {
	if s.ReadonlyInstanceDelay != nil {
		for _, item := range s.ReadonlyInstanceDelay {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay struct {
	FlushLag           *string `json:"FlushLag,omitempty" xml:"FlushLag,omitempty"`
	FlushLatency       *string `json:"FlushLatency,omitempty" xml:"FlushLatency,omitempty"`
	ReadDBInstanceName *string `json:"ReadDBInstanceName,omitempty" xml:"ReadDBInstanceName,omitempty"`
	ReplayLag          *string `json:"ReplayLag,omitempty" xml:"ReplayLag,omitempty"`
	ReplayLatency      *string `json:"ReplayLatency,omitempty" xml:"ReplayLatency,omitempty"`
	SendLatency        *string `json:"SendLatency,omitempty" xml:"SendLatency,omitempty"`
	WriteLag           *string `json:"WriteLag,omitempty" xml:"WriteLag,omitempty"`
	WriteLatency       *string `json:"WriteLatency,omitempty" xml:"WriteLatency,omitempty"`
}

func (s DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay) String() string {
	return dara.Prettify(s)
}

func (s DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay) GoString() string {
	return s.String()
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay) GetFlushLag() *string {
	return s.FlushLag
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay) GetFlushLatency() *string {
	return s.FlushLatency
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay) GetReadDBInstanceName() *string {
	return s.ReadDBInstanceName
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay) GetReplayLag() *string {
	return s.ReplayLag
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay) GetReplayLatency() *string {
	return s.ReplayLatency
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay) GetSendLatency() *string {
	return s.SendLatency
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay) GetWriteLag() *string {
	return s.WriteLag
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay) GetWriteLatency() *string {
	return s.WriteLatency
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay) SetFlushLag(v string) *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay {
	s.FlushLag = &v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay) SetFlushLatency(v string) *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay {
	s.FlushLatency = &v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay) SetReadDBInstanceName(v string) *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay {
	s.ReadDBInstanceName = &v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay) SetReplayLag(v string) *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay {
	s.ReplayLag = &v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay) SetReplayLatency(v string) *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay {
	s.ReplayLatency = &v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay) SetSendLatency(v string) *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay {
	s.SendLatency = &v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay) SetWriteLag(v string) *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay {
	s.WriteLag = &v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay) SetWriteLatency(v string) *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay {
	s.WriteLatency = &v
	return s
}

func (s *DescribeReadDBInstanceDelayResponseBodyItemsItemsReadonlyInstanceDelayReadonlyInstanceDelay) Validate() error {
	return dara.Validate(s)
}
