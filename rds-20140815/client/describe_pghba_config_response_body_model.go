// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePGHbaConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribePGHbaConfigResponseBody
	GetDBInstanceId() *string
	SetDefaultHbaItems(v *DescribePGHbaConfigResponseBodyDefaultHbaItems) *DescribePGHbaConfigResponseBody
	GetDefaultHbaItems() *DescribePGHbaConfigResponseBodyDefaultHbaItems
	SetHbaModifyTime(v string) *DescribePGHbaConfigResponseBody
	GetHbaModifyTime() *string
	SetLastModifyStatus(v string) *DescribePGHbaConfigResponseBody
	GetLastModifyStatus() *string
	SetModifyStatusReason(v string) *DescribePGHbaConfigResponseBody
	GetModifyStatusReason() *string
	SetRequestId(v string) *DescribePGHbaConfigResponseBody
	GetRequestId() *string
	SetRunningHbaItems(v *DescribePGHbaConfigResponseBodyRunningHbaItems) *DescribePGHbaConfigResponseBody
	GetRunningHbaItems() *DescribePGHbaConfigResponseBodyRunningHbaItems
}

type DescribePGHbaConfigResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-bp1*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The default configuration items of the pg_hba.conf file.
	DefaultHbaItems *DescribePGHbaConfigResponseBodyDefaultHbaItems `json:"DefaultHbaItems,omitempty" xml:"DefaultHbaItems,omitempty" type:"Struct"`
	// The time when the previous modification was made to the pg_hba.conf file.
	//
	// example:
	//
	// 2021-11-25T06:00:40Z
	HbaModifyTime *string `json:"HbaModifyTime,omitempty" xml:"HbaModifyTime,omitempty"`
	// The status of the previous modification to the pg_hba.conf file. Valid values:
	//
	// 	- **success**
	//
	// 	- **setting**
	//
	// 	- **failed**
	//
	// example:
	//
	// success
	LastModifyStatus *string `json:"LastModifyStatus,omitempty" xml:"LastModifyStatus,omitempty"`
	// The reason why the previous modification was made to the pg_hba.conf file.
	//
	// example:
	//
	// The specified users (testuser) is not exist.
	ModifyStatusReason *string `json:"ModifyStatusReason,omitempty" xml:"ModifyStatusReason,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A147A124-A147-5CCF-9609-B73C028848DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The current configuration items of the pg_hba.conf file.
	RunningHbaItems *DescribePGHbaConfigResponseBodyRunningHbaItems `json:"RunningHbaItems,omitempty" xml:"RunningHbaItems,omitempty" type:"Struct"`
}

func (s DescribePGHbaConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePGHbaConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePGHbaConfigResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribePGHbaConfigResponseBody) GetDefaultHbaItems() *DescribePGHbaConfigResponseBodyDefaultHbaItems {
	return s.DefaultHbaItems
}

func (s *DescribePGHbaConfigResponseBody) GetHbaModifyTime() *string {
	return s.HbaModifyTime
}

func (s *DescribePGHbaConfigResponseBody) GetLastModifyStatus() *string {
	return s.LastModifyStatus
}

func (s *DescribePGHbaConfigResponseBody) GetModifyStatusReason() *string {
	return s.ModifyStatusReason
}

func (s *DescribePGHbaConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePGHbaConfigResponseBody) GetRunningHbaItems() *DescribePGHbaConfigResponseBodyRunningHbaItems {
	return s.RunningHbaItems
}

func (s *DescribePGHbaConfigResponseBody) SetDBInstanceId(v string) *DescribePGHbaConfigResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribePGHbaConfigResponseBody) SetDefaultHbaItems(v *DescribePGHbaConfigResponseBodyDefaultHbaItems) *DescribePGHbaConfigResponseBody {
	s.DefaultHbaItems = v
	return s
}

func (s *DescribePGHbaConfigResponseBody) SetHbaModifyTime(v string) *DescribePGHbaConfigResponseBody {
	s.HbaModifyTime = &v
	return s
}

func (s *DescribePGHbaConfigResponseBody) SetLastModifyStatus(v string) *DescribePGHbaConfigResponseBody {
	s.LastModifyStatus = &v
	return s
}

func (s *DescribePGHbaConfigResponseBody) SetModifyStatusReason(v string) *DescribePGHbaConfigResponseBody {
	s.ModifyStatusReason = &v
	return s
}

func (s *DescribePGHbaConfigResponseBody) SetRequestId(v string) *DescribePGHbaConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePGHbaConfigResponseBody) SetRunningHbaItems(v *DescribePGHbaConfigResponseBodyRunningHbaItems) *DescribePGHbaConfigResponseBody {
	s.RunningHbaItems = v
	return s
}

func (s *DescribePGHbaConfigResponseBody) Validate() error {
	if s.DefaultHbaItems != nil {
		if err := s.DefaultHbaItems.Validate(); err != nil {
			return err
		}
	}
	if s.RunningHbaItems != nil {
		if err := s.RunningHbaItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePGHbaConfigResponseBodyDefaultHbaItems struct {
	HbaItem []*DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem `json:"HbaItem,omitempty" xml:"HbaItem,omitempty" type:"Repeated"`
}

func (s DescribePGHbaConfigResponseBodyDefaultHbaItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePGHbaConfigResponseBodyDefaultHbaItems) GoString() string {
	return s.String()
}

func (s *DescribePGHbaConfigResponseBodyDefaultHbaItems) GetHbaItem() []*DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem {
	return s.HbaItem
}

func (s *DescribePGHbaConfigResponseBodyDefaultHbaItems) SetHbaItem(v []*DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem) *DescribePGHbaConfigResponseBodyDefaultHbaItems {
	s.HbaItem = v
	return s
}

func (s *DescribePGHbaConfigResponseBodyDefaultHbaItems) Validate() error {
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

type DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem struct {
	// The IP addresses from which the specified users can access the specified databases. The value is fixed as 0.0.0.0/0.
	//
	// example:
	//
	// 0.0.0.0/0
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The names of the databases that the specified users are allowed to access. The value is fixed as all or replication.
	//
	// example:
	//
	// all
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The mask of the instance. The value is fixed as null.
	//
	// example:
	//
	// null
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The authentication method. The value is fixed as md5.
	//
	// example:
	//
	// md5
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The value of this parameter is based on the value of the Method parameter. The value is fixed as null.
	//
	// example:
	//
	// null
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// The priority of the configuration items in the pg_hba.conf file. This value is automatically generated.
	//
	// example:
	//
	// 0
	PriorityId *int32 `json:"PriorityId,omitempty" xml:"PriorityId,omitempty"`
	// The type of connection to the instance. The value is fixed as host.
	//
	// example:
	//
	// host
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user that is allowed to access the instance. The value is fixed as all.
	//
	// example:
	//
	// all
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem) String() string {
	return dara.Prettify(s)
}

func (s DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem) GoString() string {
	return s.String()
}

func (s *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem) GetAddress() *string {
	return s.Address
}

func (s *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem) GetDatabase() *string {
	return s.Database
}

func (s *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem) GetMask() *string {
	return s.Mask
}

func (s *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem) GetMethod() *string {
	return s.Method
}

func (s *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem) GetOption() *string {
	return s.Option
}

func (s *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem) GetPriorityId() *int32 {
	return s.PriorityId
}

func (s *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem) GetType() *string {
	return s.Type
}

func (s *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem) GetUser() *string {
	return s.User
}

func (s *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem) SetAddress(v string) *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem {
	s.Address = &v
	return s
}

func (s *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem) SetDatabase(v string) *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem {
	s.Database = &v
	return s
}

func (s *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem) SetMask(v string) *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem {
	s.Mask = &v
	return s
}

func (s *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem) SetMethod(v string) *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem {
	s.Method = &v
	return s
}

func (s *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem) SetOption(v string) *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem {
	s.Option = &v
	return s
}

func (s *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem) SetPriorityId(v int32) *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem {
	s.PriorityId = &v
	return s
}

func (s *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem) SetType(v string) *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem {
	s.Type = &v
	return s
}

func (s *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem) SetUser(v string) *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem {
	s.User = &v
	return s
}

func (s *DescribePGHbaConfigResponseBodyDefaultHbaItemsHbaItem) Validate() error {
	return dara.Validate(s)
}

type DescribePGHbaConfigResponseBodyRunningHbaItems struct {
	HbaItem []*DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem `json:"HbaItem,omitempty" xml:"HbaItem,omitempty" type:"Repeated"`
}

func (s DescribePGHbaConfigResponseBodyRunningHbaItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePGHbaConfigResponseBodyRunningHbaItems) GoString() string {
	return s.String()
}

func (s *DescribePGHbaConfigResponseBodyRunningHbaItems) GetHbaItem() []*DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem {
	return s.HbaItem
}

func (s *DescribePGHbaConfigResponseBodyRunningHbaItems) SetHbaItem(v []*DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem) *DescribePGHbaConfigResponseBodyRunningHbaItems {
	s.HbaItem = v
	return s
}

func (s *DescribePGHbaConfigResponseBodyRunningHbaItems) Validate() error {
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

type DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem struct {
	// The IP address of the client.
	//
	// example:
	//
	// 0.0.0.0/0
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// all
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The mask of the IP address.
	//
	// example:
	//
	// null
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The authentication method.
	//
	// example:
	//
	// md5
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The value of this parameter varies based on the value of the Method parameter. The value is fixed as null.
	//
	// example:
	//
	// null
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// The priority.
	//
	// example:
	//
	// 3
	PriorityId *int32 `json:"PriorityId,omitempty" xml:"PriorityId,omitempty"`
	// The connection type. Valor:
	//
	// 	- **host**: The record matches TCP/IP connections, including SSL connections and non-SSL connections.
	//
	// 	- **hostssl**: The record matches only TCP/IP connections that are established over SSL.
	//
	// 	- **hostnossl**: The record matches only TCP/IP connections that are not established over SSL connections.
	//
	// example:
	//
	// host
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// all
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem) String() string {
	return dara.Prettify(s)
}

func (s DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem) GoString() string {
	return s.String()
}

func (s *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem) GetAddress() *string {
	return s.Address
}

func (s *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem) GetDatabase() *string {
	return s.Database
}

func (s *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem) GetMask() *string {
	return s.Mask
}

func (s *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem) GetMethod() *string {
	return s.Method
}

func (s *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem) GetOption() *string {
	return s.Option
}

func (s *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem) GetPriorityId() *int32 {
	return s.PriorityId
}

func (s *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem) GetType() *string {
	return s.Type
}

func (s *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem) GetUser() *string {
	return s.User
}

func (s *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem) SetAddress(v string) *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem {
	s.Address = &v
	return s
}

func (s *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem) SetDatabase(v string) *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem {
	s.Database = &v
	return s
}

func (s *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem) SetMask(v string) *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem {
	s.Mask = &v
	return s
}

func (s *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem) SetMethod(v string) *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem {
	s.Method = &v
	return s
}

func (s *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem) SetOption(v string) *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem {
	s.Option = &v
	return s
}

func (s *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem) SetPriorityId(v int32) *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem {
	s.PriorityId = &v
	return s
}

func (s *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem) SetType(v string) *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem {
	s.Type = &v
	return s
}

func (s *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem) SetUser(v string) *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem {
	s.User = &v
	return s
}

func (s *DescribePGHbaConfigResponseBodyRunningHbaItemsHbaItem) Validate() error {
	return dara.Validate(s)
}
