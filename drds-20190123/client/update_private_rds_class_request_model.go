// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateRdsClassRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoUseCoupon(v bool) *UpdatePrivateRdsClassRequest
	GetAutoUseCoupon() *bool
	SetDBInstanceId(v string) *UpdatePrivateRdsClassRequest
	GetDBInstanceId() *string
	SetDrdsInstanceId(v string) *UpdatePrivateRdsClassRequest
	GetDrdsInstanceId() *string
	SetPrePayDuration(v int32) *UpdatePrivateRdsClassRequest
	GetPrePayDuration() *int32
	SetRdsClass(v string) *UpdatePrivateRdsClassRequest
	GetRdsClass() *string
	SetStorage(v string) *UpdatePrivateRdsClassRequest
	GetStorage() *string
}

type UpdatePrivateRdsClassRequest struct {
	// Specifies whether to use vouchers to offset the purchase fees. Valid values: **true*	- and **false**. Default value: false.
	//
	// > If you downgrade the specifications of an instance after you use the vouchers, the vouchers used for the purchase cannot be refunded.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The ID of the custom ApsaraDB RDS instance at the storage layer.
	//
	// > You can call the [DescribeDrdsRdsInstances](~~xxxx~~) operation to query the details of all ApsaraDB RDS instances at the storage layer of a PolarDB-X 1.0 instance, including the IDs of the ApsaraDB RDS instances.
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
	// This parameter is discontinued.
	//
	// example:
	//
	// 12
	PrePayDuration *int32 `json:"PrePayDuration,omitempty" xml:"PrePayDuration,omitempty"`
	// The new instance type of the custom ApsaraDB RDS instance at the storage layer.
	//
	// > You can call the [DescribeAvailableClasses](https://help.aliyun.com/document_detail/196546.html) operation to view the specifications that are supported for a custom ApsaraDB RDS instance. The specifications include the instance type and the storage capacity.
	//
	// example:
	//
	// rds.mysql.c1.xlarge
	RdsClass *string `json:"RdsClass,omitempty" xml:"RdsClass,omitempty"`
	// The new storage capacity of the custom ApsaraDB RDS instance at the storage layer.
	//
	// > You can call the [DescribeAvailableClasses](https://help.aliyun.com/document_detail/196546.html) operation to view the specifications that are supported for a custom ApsaraDB RDS instance. The specifications include the instance type and the storage capacity.
	//
	// example:
	//
	// 50
	Storage *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
}

func (s UpdatePrivateRdsClassRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateRdsClassRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrivateRdsClassRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *UpdatePrivateRdsClassRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpdatePrivateRdsClassRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *UpdatePrivateRdsClassRequest) GetPrePayDuration() *int32 {
	return s.PrePayDuration
}

func (s *UpdatePrivateRdsClassRequest) GetRdsClass() *string {
	return s.RdsClass
}

func (s *UpdatePrivateRdsClassRequest) GetStorage() *string {
	return s.Storage
}

func (s *UpdatePrivateRdsClassRequest) SetAutoUseCoupon(v bool) *UpdatePrivateRdsClassRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) SetDBInstanceId(v string) *UpdatePrivateRdsClassRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) SetDrdsInstanceId(v string) *UpdatePrivateRdsClassRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) SetPrePayDuration(v int32) *UpdatePrivateRdsClassRequest {
	s.PrePayDuration = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) SetRdsClass(v string) *UpdatePrivateRdsClassRequest {
	s.RdsClass = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) SetStorage(v string) *UpdatePrivateRdsClassRequest {
	s.Storage = &v
	return s
}

func (s *UpdatePrivateRdsClassRequest) Validate() error {
	return dara.Validate(s)
}
