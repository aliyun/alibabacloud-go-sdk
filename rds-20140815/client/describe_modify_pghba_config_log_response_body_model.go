// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModifyPGHbaConfigLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeModifyPGHbaConfigLogResponseBody
	GetDBInstanceId() *string
	SetHbaLogItems(v *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItems) *DescribeModifyPGHbaConfigLogResponseBody
	GetHbaLogItems() *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItems
	SetLogItemCount(v int32) *DescribeModifyPGHbaConfigLogResponseBody
	GetLogItemCount() *int32
	SetRequestId(v string) *DescribeModifyPGHbaConfigLogResponseBody
	GetRequestId() *string
}

type DescribeModifyPGHbaConfigLogResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// pgm-bp1lymyn1v3i****
	DBInstanceId *string                                              `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	HbaLogItems  *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItems `json:"HbaLogItems,omitempty" xml:"HbaLogItems,omitempty" type:"Struct"`
	// The number of modification records.
	//
	// example:
	//
	// 1
	LogItemCount *int32 `json:"LogItemCount,omitempty" xml:"LogItemCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6D797E6B-E157-510C-A27F-6F9E6DA40633
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeModifyPGHbaConfigLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModifyPGHbaConfigLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModifyPGHbaConfigLogResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeModifyPGHbaConfigLogResponseBody) GetHbaLogItems() *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItems {
	return s.HbaLogItems
}

func (s *DescribeModifyPGHbaConfigLogResponseBody) GetLogItemCount() *int32 {
	return s.LogItemCount
}

func (s *DescribeModifyPGHbaConfigLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModifyPGHbaConfigLogResponseBody) SetDBInstanceId(v string) *DescribeModifyPGHbaConfigLogResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBody) SetHbaLogItems(v *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItems) *DescribeModifyPGHbaConfigLogResponseBody {
	s.HbaLogItems = v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBody) SetLogItemCount(v int32) *DescribeModifyPGHbaConfigLogResponseBody {
	s.LogItemCount = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBody) SetRequestId(v string) *DescribeModifyPGHbaConfigLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBody) Validate() error {
	if s.HbaLogItems != nil {
		if err := s.HbaLogItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeModifyPGHbaConfigLogResponseBodyHbaLogItems struct {
	HbaLogItem []*DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem `json:"HbaLogItem,omitempty" xml:"HbaLogItem,omitempty" type:"Repeated"`
}

func (s DescribeModifyPGHbaConfigLogResponseBodyHbaLogItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeModifyPGHbaConfigLogResponseBodyHbaLogItems) GoString() string {
	return s.String()
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItems) GetHbaLogItem() []*DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem {
	return s.HbaLogItem
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItems) SetHbaLogItem(v []*DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItems {
	s.HbaLogItem = v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItems) Validate() error {
	if s.HbaLogItem != nil {
		for _, item := range s.HbaLogItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem struct {
	AfterHbaItems  *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItems  `json:"AfterHbaItems,omitempty" xml:"AfterHbaItems,omitempty" type:"Struct"`
	BeforeHbaItems *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItems `json:"BeforeHbaItems,omitempty" xml:"BeforeHbaItems,omitempty" type:"Struct"`
	ModifyStatus   *string                                                                      `json:"ModifyStatus,omitempty" xml:"ModifyStatus,omitempty"`
	ModifyTime     *string                                                                      `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	StatusReason   *string                                                                      `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
}

func (s DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem) GoString() string {
	return s.String()
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem) GetAfterHbaItems() *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItems {
	return s.AfterHbaItems
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem) GetBeforeHbaItems() *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItems {
	return s.BeforeHbaItems
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem) GetModifyStatus() *string {
	return s.ModifyStatus
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem) GetStatusReason() *string {
	return s.StatusReason
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem) SetAfterHbaItems(v *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItems) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem {
	s.AfterHbaItems = v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem) SetBeforeHbaItems(v *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItems) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem {
	s.BeforeHbaItems = v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem) SetModifyStatus(v string) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem {
	s.ModifyStatus = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem) SetModifyTime(v string) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem {
	s.ModifyTime = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem) SetStatusReason(v string) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem {
	s.StatusReason = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItem) Validate() error {
	if s.AfterHbaItems != nil {
		if err := s.AfterHbaItems.Validate(); err != nil {
			return err
		}
	}
	if s.BeforeHbaItems != nil {
		if err := s.BeforeHbaItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItems struct {
	HbaItem []*DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem `json:"HbaItem,omitempty" xml:"HbaItem,omitempty" type:"Repeated"`
}

func (s DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItems) GoString() string {
	return s.String()
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItems) GetHbaItem() []*DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem {
	return s.HbaItem
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItems) SetHbaItem(v []*DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItems {
	s.HbaItem = v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItems) Validate() error {
	if s.HbaItem != nil {
		for _, item := range s.HbaItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem struct {
	Address    *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Database   *string `json:"Database,omitempty" xml:"Database,omitempty"`
	Mask       *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	Method     *string `json:"Method,omitempty" xml:"Method,omitempty"`
	Option     *string `json:"Option,omitempty" xml:"Option,omitempty"`
	PriorityId *int32  `json:"PriorityId,omitempty" xml:"PriorityId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	User       *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem) GoString() string {
	return s.String()
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem) GetAddress() *string {
	return s.Address
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem) GetDatabase() *string {
	return s.Database
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem) GetMask() *string {
	return s.Mask
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem) GetMethod() *string {
	return s.Method
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem) GetOption() *string {
	return s.Option
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem) GetPriorityId() *int32 {
	return s.PriorityId
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem) GetType() *string {
	return s.Type
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem) GetUser() *string {
	return s.User
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem) SetAddress(v string) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem {
	s.Address = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem) SetDatabase(v string) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem {
	s.Database = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem) SetMask(v string) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem {
	s.Mask = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem) SetMethod(v string) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem {
	s.Method = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem) SetOption(v string) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem {
	s.Option = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem) SetPriorityId(v int32) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem {
	s.PriorityId = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem) SetType(v string) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem {
	s.Type = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem) SetUser(v string) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem {
	s.User = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemAfterHbaItemsHbaItem) Validate() error {
	return dara.Validate(s)
}

type DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItems struct {
	HbaItem []*DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem `json:"HbaItem,omitempty" xml:"HbaItem,omitempty" type:"Repeated"`
}

func (s DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItems) GoString() string {
	return s.String()
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItems) GetHbaItem() []*DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem {
	return s.HbaItem
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItems) SetHbaItem(v []*DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItems {
	s.HbaItem = v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItems) Validate() error {
	if s.HbaItem != nil {
		for _, item := range s.HbaItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem struct {
	Address    *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Database   *string `json:"Database,omitempty" xml:"Database,omitempty"`
	Mask       *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	Method     *string `json:"Method,omitempty" xml:"Method,omitempty"`
	Option     *string `json:"Option,omitempty" xml:"Option,omitempty"`
	PriorityId *int32  `json:"PriorityId,omitempty" xml:"PriorityId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	User       *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem) GoString() string {
	return s.String()
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem) GetAddress() *string {
	return s.Address
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem) GetDatabase() *string {
	return s.Database
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem) GetMask() *string {
	return s.Mask
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem) GetMethod() *string {
	return s.Method
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem) GetOption() *string {
	return s.Option
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem) GetPriorityId() *int32 {
	return s.PriorityId
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem) GetType() *string {
	return s.Type
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem) GetUser() *string {
	return s.User
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem) SetAddress(v string) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem {
	s.Address = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem) SetDatabase(v string) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem {
	s.Database = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem) SetMask(v string) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem {
	s.Mask = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem) SetMethod(v string) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem {
	s.Method = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem) SetOption(v string) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem {
	s.Option = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem) SetPriorityId(v int32) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem {
	s.PriorityId = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem) SetType(v string) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem {
	s.Type = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem) SetUser(v string) *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem {
	s.User = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponseBodyHbaLogItemsHbaLogItemBeforeHbaItemsHbaItem) Validate() error {
	return dara.Validate(s)
}
