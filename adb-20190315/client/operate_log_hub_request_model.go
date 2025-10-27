// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateLogHubRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreate(v bool) *OperateLogHubRequest
	GetCreate() *bool
	SetDBClusterId(v string) *OperateLogHubRequest
	GetDBClusterId() *string
	SetDeliverName(v string) *OperateLogHubRequest
	GetDeliverName() *string
	SetDeliverTime(v string) *OperateLogHubRequest
	GetDeliverTime() *string
	SetDescription(v string) *OperateLogHubRequest
	GetDescription() *string
	SetFilterDirtyData(v bool) *OperateLogHubRequest
	GetFilterDirtyData() *bool
	SetLogHubStores(v []*OperateLogHubRequestLogHubStores) *OperateLogHubRequest
	GetLogHubStores() []*OperateLogHubRequestLogHubStores
	SetLogStoreName(v string) *OperateLogHubRequest
	GetLogStoreName() *string
	SetOwnerAccount(v string) *OperateLogHubRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *OperateLogHubRequest
	GetOwnerId() *int64
	SetPassword(v string) *OperateLogHubRequest
	GetPassword() *string
	SetProjectName(v string) *OperateLogHubRequest
	GetProjectName() *string
	SetProvider(v string) *OperateLogHubRequest
	GetProvider() *string
	SetResourceOwnerAccount(v string) *OperateLogHubRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *OperateLogHubRequest
	GetResourceOwnerId() *int64
	SetSchemaName(v string) *OperateLogHubRequest
	GetSchemaName() *string
	SetTableName(v string) *OperateLogHubRequest
	GetTableName() *string
	SetUserName(v string) *OperateLogHubRequest
	GetUserName() *string
}

type OperateLogHubRequest struct {
	// Specifies whether to create the log shipping job.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Create *bool `json:"Create,omitempty" xml:"Create,omitempty"`
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-uf6rtqaj25491628z
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the log shipping job.
	//
	// This parameter is required.
	//
	// example:
	//
	// xhxsblz_limited_gift_pay
	DeliverName *string `json:"DeliverName,omitempty" xml:"DeliverName,omitempty"`
	// The shipping time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-10-24\\"T\\"13:10\\"Z\\"
	DeliverTime *string `json:"DeliverTime,omitempty" xml:"DeliverTime,omitempty"`
	// The description of the log shipping job.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to filter dirty data.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	FilterDirtyData *bool `json:"FilterDirtyData,omitempty" xml:"FilterDirtyData,omitempty"`
	// The log keywords.
	//
	// This parameter is required.
	LogHubStores []*OperateLogHubRequestLogHubStores `json:"LogHubStores,omitempty" xml:"LogHubStores,omitempty" type:"Repeated"`
	// The name of the Logstore.
	//
	// This parameter is required.
	//
	// example:
	//
	// beta-game-mjxb-ham-pool-export
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ads_123
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The name of the Simple Log Service project.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-adb
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The channel of the log shipping job.
	//
	// example:
	//
	// SLS
	Provider             *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// wddata
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// rest_action_latest
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// admin
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s OperateLogHubRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateLogHubRequest) GoString() string {
	return s.String()
}

func (s *OperateLogHubRequest) GetCreate() *bool {
	return s.Create
}

func (s *OperateLogHubRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *OperateLogHubRequest) GetDeliverName() *string {
	return s.DeliverName
}

func (s *OperateLogHubRequest) GetDeliverTime() *string {
	return s.DeliverTime
}

func (s *OperateLogHubRequest) GetDescription() *string {
	return s.Description
}

func (s *OperateLogHubRequest) GetFilterDirtyData() *bool {
	return s.FilterDirtyData
}

func (s *OperateLogHubRequest) GetLogHubStores() []*OperateLogHubRequestLogHubStores {
	return s.LogHubStores
}

func (s *OperateLogHubRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *OperateLogHubRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *OperateLogHubRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OperateLogHubRequest) GetPassword() *string {
	return s.Password
}

func (s *OperateLogHubRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *OperateLogHubRequest) GetProvider() *string {
	return s.Provider
}

func (s *OperateLogHubRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OperateLogHubRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *OperateLogHubRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *OperateLogHubRequest) GetTableName() *string {
	return s.TableName
}

func (s *OperateLogHubRequest) GetUserName() *string {
	return s.UserName
}

func (s *OperateLogHubRequest) SetCreate(v bool) *OperateLogHubRequest {
	s.Create = &v
	return s
}

func (s *OperateLogHubRequest) SetDBClusterId(v string) *OperateLogHubRequest {
	s.DBClusterId = &v
	return s
}

func (s *OperateLogHubRequest) SetDeliverName(v string) *OperateLogHubRequest {
	s.DeliverName = &v
	return s
}

func (s *OperateLogHubRequest) SetDeliverTime(v string) *OperateLogHubRequest {
	s.DeliverTime = &v
	return s
}

func (s *OperateLogHubRequest) SetDescription(v string) *OperateLogHubRequest {
	s.Description = &v
	return s
}

func (s *OperateLogHubRequest) SetFilterDirtyData(v bool) *OperateLogHubRequest {
	s.FilterDirtyData = &v
	return s
}

func (s *OperateLogHubRequest) SetLogHubStores(v []*OperateLogHubRequestLogHubStores) *OperateLogHubRequest {
	s.LogHubStores = v
	return s
}

func (s *OperateLogHubRequest) SetLogStoreName(v string) *OperateLogHubRequest {
	s.LogStoreName = &v
	return s
}

func (s *OperateLogHubRequest) SetOwnerAccount(v string) *OperateLogHubRequest {
	s.OwnerAccount = &v
	return s
}

func (s *OperateLogHubRequest) SetOwnerId(v int64) *OperateLogHubRequest {
	s.OwnerId = &v
	return s
}

func (s *OperateLogHubRequest) SetPassword(v string) *OperateLogHubRequest {
	s.Password = &v
	return s
}

func (s *OperateLogHubRequest) SetProjectName(v string) *OperateLogHubRequest {
	s.ProjectName = &v
	return s
}

func (s *OperateLogHubRequest) SetProvider(v string) *OperateLogHubRequest {
	s.Provider = &v
	return s
}

func (s *OperateLogHubRequest) SetResourceOwnerAccount(v string) *OperateLogHubRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OperateLogHubRequest) SetResourceOwnerId(v int64) *OperateLogHubRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OperateLogHubRequest) SetSchemaName(v string) *OperateLogHubRequest {
	s.SchemaName = &v
	return s
}

func (s *OperateLogHubRequest) SetTableName(v string) *OperateLogHubRequest {
	s.TableName = &v
	return s
}

func (s *OperateLogHubRequest) SetUserName(v string) *OperateLogHubRequest {
	s.UserName = &v
	return s
}

func (s *OperateLogHubRequest) Validate() error {
	if s.LogHubStores != nil {
		for _, item := range s.LogHubStores {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OperateLogHubRequestLogHubStores struct {
	// The value of the log keyword.
	//
	// This parameter is required.
	//
	// example:
	//
	// testValue
	FieldKey *string `json:"FieldKey,omitempty" xml:"FieldKey,omitempty"`
	// The log keyword.
	//
	// This parameter is required.
	//
	// example:
	//
	// testKey
	LogKey *string `json:"LogKey,omitempty" xml:"LogKey,omitempty"`
}

func (s OperateLogHubRequestLogHubStores) String() string {
	return dara.Prettify(s)
}

func (s OperateLogHubRequestLogHubStores) GoString() string {
	return s.String()
}

func (s *OperateLogHubRequestLogHubStores) GetFieldKey() *string {
	return s.FieldKey
}

func (s *OperateLogHubRequestLogHubStores) GetLogKey() *string {
	return s.LogKey
}

func (s *OperateLogHubRequestLogHubStores) SetFieldKey(v string) *OperateLogHubRequestLogHubStores {
	s.FieldKey = &v
	return s
}

func (s *OperateLogHubRequestLogHubStores) SetLogKey(v string) *OperateLogHubRequestLogHubStores {
	s.LogKey = &v
	return s
}

func (s *OperateLogHubRequestLogHubStores) Validate() error {
	return dara.Validate(s)
}
