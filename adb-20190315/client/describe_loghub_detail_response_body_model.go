// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoghubDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoghubInfo(v *DescribeLoghubDetailResponseBodyLoghubInfo) *DescribeLoghubDetailResponseBody
	GetLoghubInfo() *DescribeLoghubDetailResponseBodyLoghubInfo
	SetRequestId(v string) *DescribeLoghubDetailResponseBody
	GetRequestId() *string
}

type DescribeLoghubDetailResponseBody struct {
	// The queried log collection information.
	LoghubInfo *DescribeLoghubDetailResponseBodyLoghubInfo `json:"LoghubInfo,omitempty" xml:"LoghubInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLoghubDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoghubDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoghubDetailResponseBody) GetLoghubInfo() *DescribeLoghubDetailResponseBodyLoghubInfo {
	return s.LoghubInfo
}

func (s *DescribeLoghubDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoghubDetailResponseBody) SetLoghubInfo(v *DescribeLoghubDetailResponseBodyLoghubInfo) *DescribeLoghubDetailResponseBody {
	s.LoghubInfo = v
	return s
}

func (s *DescribeLoghubDetailResponseBody) SetRequestId(v string) *DescribeLoghubDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoghubDetailResponseBody) Validate() error {
	if s.LoghubInfo != nil {
		if err := s.LoghubInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLoghubDetailResponseBodyLoghubInfo struct {
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
	// The name of the log shipping job.
	//
	// example:
	//
	// hub-pay-callback-ykt
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
	// aaa
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The endpoint.
	//
	// example:
	//
	// am-bp1p38m2c26s7hpk690650.ads.aliyuncs.com
	DomainUrl *string `json:"DomainUrl,omitempty" xml:"DomainUrl,omitempty"`
	// Indicates whether dirty data is filtered.
	//
	// example:
	//
	// true
	FilterDirtyData *bool `json:"FilterDirtyData,omitempty" xml:"FilterDirtyData,omitempty"`
	// The log keywords.
	LogHubStores *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStores `json:"LogHubStores,omitempty" xml:"LogHubStores,omitempty" type:"Struct"`
	// The name of the Logstore.
	//
	// example:
	//
	// device_login
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
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
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// cbd_bi
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
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
	// test
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeLoghubDetailResponseBodyLoghubInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoghubDetailResponseBodyLoghubInfo) GoString() string {
	return s.String()
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) GetDBType() *string {
	return s.DBType
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) GetDeliverName() *string {
	return s.DeliverName
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) GetDeliverTime() *string {
	return s.DeliverTime
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) GetDomainUrl() *string {
	return s.DomainUrl
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) GetFilterDirtyData() *bool {
	return s.FilterDirtyData
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) GetLogHubStores() *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStores {
	return s.LogHubStores
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) GetTableName() *string {
	return s.TableName
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) GetUserName() *string {
	return s.UserName
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetDBClusterId(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetDBType(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.DBType = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetDeliverName(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.DeliverName = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetDeliverTime(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.DeliverTime = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetDescription(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.Description = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetDomainUrl(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.DomainUrl = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetFilterDirtyData(v bool) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.FilterDirtyData = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetLogHubStores(v *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStores) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.LogHubStores = v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetLogStoreName(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.LogStoreName = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetProjectName(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.ProjectName = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetRegionId(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetSchemaName(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.SchemaName = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetTableName(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.TableName = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetUserName(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.UserName = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetZoneId(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.ZoneId = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) Validate() error {
	if s.LogHubStores != nil {
		if err := s.LogHubStores.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLoghubDetailResponseBodyLoghubInfoLogHubStores struct {
	LogHubStore []*DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore `json:"LogHubStore,omitempty" xml:"LogHubStore,omitempty" type:"Repeated"`
}

func (s DescribeLoghubDetailResponseBodyLoghubInfoLogHubStores) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoghubDetailResponseBodyLoghubInfoLogHubStores) GoString() string {
	return s.String()
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStores) GetLogHubStore() []*DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore {
	return s.LogHubStore
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStores) SetLogHubStore(v []*DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore) *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStores {
	s.LogHubStore = v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStores) Validate() error {
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

type DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore struct {
	// The value of the log keyword.
	//
	// example:
	//
	// channel_id
	FieldKey *string `json:"FieldKey,omitempty" xml:"FieldKey,omitempty"`
	// The log keyword.
	//
	// example:
	//
	// put_request
	LogKey *string `json:"LogKey,omitempty" xml:"LogKey,omitempty"`
}

func (s DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore) GoString() string {
	return s.String()
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore) GetFieldKey() *string {
	return s.FieldKey
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore) GetLogKey() *string {
	return s.LogKey
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore) SetFieldKey(v string) *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore {
	s.FieldKey = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore) SetLogKey(v string) *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore {
	s.LogKey = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore) Validate() error {
	return dara.Validate(s)
}
