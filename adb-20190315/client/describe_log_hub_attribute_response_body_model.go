// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogHubAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoghubInfo(v *DescribeLogHubAttributeResponseBodyLoghubInfo) *DescribeLogHubAttributeResponseBody
	GetLoghubInfo() *DescribeLogHubAttributeResponseBodyLoghubInfo
	SetRequestId(v string) *DescribeLogHubAttributeResponseBody
	GetRequestId() *string
}

type DescribeLogHubAttributeResponseBody struct {
	// The log collection information.
	LoghubInfo *DescribeLogHubAttributeResponseBodyLoghubInfo `json:"LoghubInfo,omitempty" xml:"LoghubInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 8A564B7F-8C00-43C0-8EC5-919FBB70573
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLogHubAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogHubAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogHubAttributeResponseBody) GetLoghubInfo() *DescribeLogHubAttributeResponseBodyLoghubInfo {
	return s.LoghubInfo
}

func (s *DescribeLogHubAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogHubAttributeResponseBody) SetLoghubInfo(v *DescribeLogHubAttributeResponseBodyLoghubInfo) *DescribeLogHubAttributeResponseBody {
	s.LoghubInfo = v
	return s
}

func (s *DescribeLogHubAttributeResponseBody) SetRequestId(v string) *DescribeLogHubAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBody) Validate() error {
	if s.LoghubInfo != nil {
		if err := s.LoghubInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLogHubAttributeResponseBodyLoghubInfo struct {
	// The cluster ID.
	//
	// example:
	//
	// am-8vbs48m7553du1gz2
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database type.
	//
	// example:
	//
	// rds
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The synchronization latency, which is the latency between the latest update time of the synchronization job and the current system time. Unit: seconds.
	//
	// example:
	//
	// 361
	Delay *int64 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The name of the log shipping job.
	//
	// example:
	//
	// loghub-web-login-new
	DeliverName *string `json:"DeliverName,omitempty" xml:"DeliverName,omitempty"`
	// The log shipping time.
	//
	// example:
	//
	// 2024-12-06\\"T\\"10:15\\"Z\\"
	DeliverTime *string `json:"DeliverTime,omitempty" xml:"DeliverTime,omitempty"`
	// The description.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether dirty data is filtered.
	//
	// example:
	//
	// true
	FilterDirtyData *bool `json:"FilterDirtyData,omitempty" xml:"FilterDirtyData,omitempty"`
	// The names of the Logstores.
	LogHubStores *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStores `json:"LogHubStores,omitempty" xml:"LogHubStores,omitempty" type:"Struct"`
	// The name of the Logstore.
	//
	// example:
	//
	// device_login
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The returned message.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// test-adb
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// cbd_bi
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The synchronization status.
	//
	// example:
	//
	// processing
	SyncStatus *string `json:"SyncStatus,omitempty" xml:"SyncStatus,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test2
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// aaa
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeLogHubAttributeResponseBodyLoghubInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogHubAttributeResponseBodyLoghubInfo) GoString() string {
	return s.String()
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) GetDBType() *string {
	return s.DBType
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) GetDelay() *int64 {
	return s.Delay
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) GetDeliverName() *string {
	return s.DeliverName
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) GetDeliverTime() *string {
	return s.DeliverTime
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) GetFilterDirtyData() *bool {
	return s.FilterDirtyData
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) GetLogHubStores() *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStores {
	return s.LogHubStores
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) GetMessage() *string {
	return s.Message
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) GetSyncStatus() *string {
	return s.SyncStatus
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) GetTableName() *string {
	return s.TableName
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) GetUserName() *string {
	return s.UserName
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetDBClusterId(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetDBType(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.DBType = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetDelay(v int64) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.Delay = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetDeliverName(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.DeliverName = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetDeliverTime(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.DeliverTime = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetDescription(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.Description = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetFilterDirtyData(v bool) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.FilterDirtyData = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetLogHubStores(v *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStores) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.LogHubStores = v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetLogStoreName(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.LogStoreName = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetMessage(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.Message = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetProjectName(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.ProjectName = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetRegionId(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetSchemaName(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.SchemaName = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetSyncStatus(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.SyncStatus = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetTableName(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.TableName = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetUserName(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.UserName = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetZoneId(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.ZoneId = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) Validate() error {
	if s.LogHubStores != nil {
		if err := s.LogHubStores.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStores struct {
	LogHubStore []*DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore `json:"LogHubStore,omitempty" xml:"LogHubStore,omitempty" type:"Repeated"`
}

func (s DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStores) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStores) GoString() string {
	return s.String()
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStores) GetLogHubStore() []*DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore {
	return s.LogHubStore
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStores) SetLogHubStore(v []*DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore) *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStores {
	s.LogHubStore = v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStores) Validate() error {
	if s.LogHubStore != nil {
		for _, item := range s.LogHubStore {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore struct {
	// The value of the log keyword.
	//
	// example:
	//
	// 255311
	FieldKey *string `json:"FieldKey,omitempty" xml:"FieldKey,omitempty"`
	// The log keyword.
	//
	// example:
	//
	// taskid
	LogKey *string `json:"LogKey,omitempty" xml:"LogKey,omitempty"`
}

func (s DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore) GoString() string {
	return s.String()
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore) GetFieldKey() *string {
	return s.FieldKey
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore) GetLogKey() *string {
	return s.LogKey
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore) SetFieldKey(v string) *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore {
	s.FieldKey = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore) SetLogKey(v string) *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore {
	s.LogKey = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore) Validate() error {
	return dara.Validate(s)
}
