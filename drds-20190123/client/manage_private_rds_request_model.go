// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManagePrivateRdsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ManagePrivateRdsRequest
	GetDBInstanceId() *string
	SetDrdsInstanceId(v string) *ManagePrivateRdsRequest
	GetDrdsInstanceId() *string
	SetParams(v string) *ManagePrivateRdsRequest
	GetParams() *string
	SetRdsAction(v string) *ManagePrivateRdsRequest
	GetRdsAction() *string
	SetRegionId(v string) *ManagePrivateRdsRequest
	GetRegionId() *string
}

type ManagePrivateRdsRequest struct {
	// The ID of the custom ApsaraDB RDS instance at the storage layer.
	//
	// > You can call the [DescribeDrdsRdsInstances](https://help.aliyun.com/document_detail/215526.html) operation to query the details of all ApsaraDB RDS instances, including the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// > You can call the [DescribeDrdsInstances](https://help.aliyun.com/document_detail/139284.html) operation to query the details of all PolarDB-X 1.0 instances within an Alibaba Cloud account, including the IDs of the instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds*************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The JSON string that consists of request parameters and the values of the request parameters of an operation that you need to call for the custom ApsaraDB RDS instance. The value of a request parameter is of the STRING type. Example: `{NodeId:"1797****"}`.
	//
	// For more information about the request parameters and valid values of the request parameters of each operation, see the request parameter sections in the following topics:
	//
	// 	- [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/26231.html)
	//
	// 	- [DescribeAvailableClasses](https://help.aliyun.com/document_detail/196546.html)
	//
	// 	- [DescribeSQLCollectorPolicy](https://help.aliyun.com/document_detail/26292.html)
	//
	// 	- [ModifySQLCollectorPolicy](https://help.aliyun.com/document_detail/26293.html)
	//
	// 	- [DescribeParameters](https://help.aliyun.com/document_detail/26285.html)
	//
	// 	- [ModifyParameter](https://help.aliyun.com/document_detail/26286.html)
	//
	// 	- [DescribeDBInstanceHAConfig](https://help.aliyun.com/document_detail/26244.html)
	//
	// 	- [SwitchDBInstanceHA](https://help.aliyun.com/document_detail/26251.html)
	//
	// > Among the required request parameters of the preceding operations, you do not need to specify the `Action` and `DBInstanceId` parameters. You must specify all the other required request parameters.
	//
	// example:
	//
	// {NodeId:"1797****"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The operation that you want to perform on the custom ApsaraDB RDS instance. Valid values:
	//
	// 	- **DescribeDBInstanceAttribute**: queries the details of the custom ApsaraDB RDS instance.
	//
	// 	- **DescribeAvailableClasses**: queries the specifications that are supported for a custom ApsaraDB RDS instance. The specifications include the instance type and the storage capacity.
	//
	// 	- **DescribeSQLCollectorPolicy**: queries whether the SQL Explorer (SQL Audit) feature is enabled for custom ApsaraDB RDS instance.
	//
	// 	- **ModifySQLCollectorPolicy**: enables or disables the SQL Explorer (SQL Audit) feature for the custom ApsaraDB RDS instance.
	//
	// 	- **DescribeParameters**: queries the parameter settings of the custom ApsaraDB RDS instance.
	//
	// 	- **ModifyParameter**: modifies the parameters of the custom ApsaraDB RDS instance.
	//
	// 	- **DescribeDBInstanceHAConfig**: queries the high availability mode and data replication mode of the custom ApsaraDB RDS instance.
	//
	// 	- **SwitchDBInstanceHA**: switches workloads between the primary and secondary custom ApsaraDB RDS instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// SwitchDBInstanceHA
	RdsAction *string `json:"RdsAction,omitempty" xml:"RdsAction,omitempty"`
	// The ID of the region in which the PolarDB-X 1.0 instance resides.
	//
	// > You can call the [DescribeDrdsInstances](https://help.aliyun.com/document_detail/139284.html) operation to query the details of all PolarDB-X 1.0 instances within an Alibaba Cloud account, including the IDs of regions in which the instances reside.
	//
	// example:
	//
	// cn-hanzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ManagePrivateRdsRequest) String() string {
	return dara.Prettify(s)
}

func (s ManagePrivateRdsRequest) GoString() string {
	return s.String()
}

func (s *ManagePrivateRdsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ManagePrivateRdsRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *ManagePrivateRdsRequest) GetParams() *string {
	return s.Params
}

func (s *ManagePrivateRdsRequest) GetRdsAction() *string {
	return s.RdsAction
}

func (s *ManagePrivateRdsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ManagePrivateRdsRequest) SetDBInstanceId(v string) *ManagePrivateRdsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ManagePrivateRdsRequest) SetDrdsInstanceId(v string) *ManagePrivateRdsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ManagePrivateRdsRequest) SetParams(v string) *ManagePrivateRdsRequest {
	s.Params = &v
	return s
}

func (s *ManagePrivateRdsRequest) SetRdsAction(v string) *ManagePrivateRdsRequest {
	s.RdsAction = &v
	return s
}

func (s *ManagePrivateRdsRequest) SetRegionId(v string) *ManagePrivateRdsRequest {
	s.RegionId = &v
	return s
}

func (s *ManagePrivateRdsRequest) Validate() error {
	return dara.Validate(s)
}
