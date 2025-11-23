// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableTopologyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetTableTopologyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetTableTopologyResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetTableTopologyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTableTopologyResponseBody
	GetSuccess() *bool
	SetTableTopology(v *GetTableTopologyResponseBodyTableTopology) *GetTableTopologyResponseBody
	GetTableTopology() *GetTableTopologyResponseBodyTableTopology
}

type GetTableTopologyResponseBody struct {
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F5ED6A40-F344-4C7D-A8F0-5685CA584CA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The topology information.
	TableTopology *GetTableTopologyResponseBodyTableTopology `json:"TableTopology,omitempty" xml:"TableTopology,omitempty" type:"Struct"`
}

func (s GetTableTopologyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableTopologyResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableTopologyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTableTopologyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTableTopologyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableTopologyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTableTopologyResponseBody) GetTableTopology() *GetTableTopologyResponseBodyTableTopology {
	return s.TableTopology
}

func (s *GetTableTopologyResponseBody) SetErrorCode(v string) *GetTableTopologyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTableTopologyResponseBody) SetErrorMessage(v string) *GetTableTopologyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTableTopologyResponseBody) SetRequestId(v string) *GetTableTopologyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableTopologyResponseBody) SetSuccess(v bool) *GetTableTopologyResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableTopologyResponseBody) SetTableTopology(v *GetTableTopologyResponseBodyTableTopology) *GetTableTopologyResponseBody {
	s.TableTopology = v
	return s
}

func (s *GetTableTopologyResponseBody) Validate() error {
	if s.TableTopology != nil {
		if err := s.TableTopology.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTableTopologyResponseBodyTableTopology struct {
	// Indicates whether the table is a logical table. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The GUID of the table in DMS.
	//
	// example:
	//
	// IDB_L_308302.yuyang_test.test_ch
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test_ch
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// Information of the topology of the table.
	TableTopologyInfoList []*GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList `json:"TableTopologyInfoList,omitempty" xml:"TableTopologyInfoList,omitempty" type:"Repeated"`
}

func (s GetTableTopologyResponseBodyTableTopology) String() string {
	return dara.Prettify(s)
}

func (s GetTableTopologyResponseBodyTableTopology) GoString() string {
	return s.String()
}

func (s *GetTableTopologyResponseBodyTableTopology) GetLogic() *bool {
	return s.Logic
}

func (s *GetTableTopologyResponseBodyTableTopology) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetTableTopologyResponseBodyTableTopology) GetTableName() *string {
	return s.TableName
}

func (s *GetTableTopologyResponseBodyTableTopology) GetTableTopologyInfoList() []*GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	return s.TableTopologyInfoList
}

func (s *GetTableTopologyResponseBodyTableTopology) SetLogic(v bool) *GetTableTopologyResponseBodyTableTopology {
	s.Logic = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopology) SetTableGuid(v string) *GetTableTopologyResponseBodyTableTopology {
	s.TableGuid = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopology) SetTableName(v string) *GetTableTopologyResponseBodyTableTopology {
	s.TableName = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopology) SetTableTopologyInfoList(v []*GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) *GetTableTopologyResponseBodyTableTopology {
	s.TableTopologyInfoList = v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopology) Validate() error {
	if s.TableTopologyInfoList != nil {
		for _, item := range s.TableTopologyInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList struct {
	// The ID of the physical database.
	//
	// example:
	//
	// 43215
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// yuyang_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The name that is used to search for the database.
	//
	// > We recommend that you do not use this parameter for business development. The format of the parameter value may be modified in later versions.
	//
	// example:
	//
	// yuyang_test@localhost:3306
	//
	// [yuyang_test_dev]
	DbSearchName *string `json:"DbSearchName,omitempty" xml:"DbSearchName,omitempty"`
	// The database engine.
	//
	// example:
	//
	// mysql
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The ID of the instance to which the physical database belongs.
	//
	// example:
	//
	// 4325325
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the resource related to the instance. The resource corresponds with the database instance type returned in the InstanceSource parameter.
	//
	// 	- **RDS**:The ID of the ApsaraDB RDS instance.
	//
	// 	- **ECS_OWN**: The ID of the Elastic Compute Service (ECS) instance.
	//
	// 	- **PUBLIC_OWN**: This parameter is left empty for self-managed database instances that are connected over the Internet.
	//
	// 	- **VPC_ID**:The ID of the virtual private cloud (VPC).
	//
	// 	- **GATEWAY**: The ID of the database gateway.
	//
	// example:
	//
	// rm-xxx
	InstanceResourceId *string `json:"InstanceResourceId,omitempty" xml:"InstanceResourceId,omitempty"`
	// The type of the database instance. Valid values:
	//
	// 	- **RDS**: an ApsaraDB RDS instance.
	//
	// 	- **ECS_OWN**: a self-managed database that is deployed on an ECS instance
	//
	// 	- **PUBLIC_OWN**: a self-managed database instance that is connected over the Internet.
	//
	// 	- **VPC_ID**: a self-managed database instance in a VPC that is connected over Express Connect circuits.
	//
	// 	- **GATEWAY**: a database instance connected by using a database gateway.
	//
	// example:
	//
	// RDS
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of tables.
	//
	// example:
	//
	// 6
	TableCount *int64 `json:"TableCount,omitempty" xml:"TableCount,omitempty"`
	// The expression of the names of logical tables.
	//
	// **
	//
	// **Description*	- This parameter is not returned for physical tables.
	//
	// example:
	//
	// test_ch_[0000-0005]
	TableNameExpr *string `json:"TableNameExpr,omitempty" xml:"TableNameExpr,omitempty"`
	// The names of tables.
	//
	// > The table names are separated by commas (,).
	//
	// example:
	//
	// test_ch_0000,test_ch_0001,test_ch_0002,test_ch_0003,test_ch_0004,test_ch_0005
	TableNameList *string `json:"TableNameList,omitempty" xml:"TableNameList,omitempty"`
}

func (s GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) GoString() string {
	return s.String()
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) GetDbId() *int64 {
	return s.DbId
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) GetDbName() *string {
	return s.DbName
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) GetDbSearchName() *string {
	return s.DbSearchName
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) GetDbType() *string {
	return s.DbType
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) GetInstanceResourceId() *string {
	return s.InstanceResourceId
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) GetInstanceSource() *string {
	return s.InstanceSource
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) GetTableCount() *int64 {
	return s.TableCount
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) GetTableNameExpr() *string {
	return s.TableNameExpr
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) GetTableNameList() *string {
	return s.TableNameList
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetDbId(v int64) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.DbId = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetDbName(v string) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.DbName = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetDbSearchName(v string) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.DbSearchName = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetDbType(v string) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.DbType = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetInstanceId(v int64) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.InstanceId = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetInstanceResourceId(v string) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.InstanceResourceId = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetInstanceSource(v string) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.InstanceSource = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetRegionId(v string) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.RegionId = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetTableCount(v int64) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.TableCount = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetTableNameExpr(v string) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.TableNameExpr = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetTableNameList(v string) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.TableNameList = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) Validate() error {
	return dara.Validate(s)
}
