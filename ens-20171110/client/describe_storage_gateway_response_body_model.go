// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeStorageGatewayResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeStorageGatewayResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeStorageGatewayResponseBody
	GetRequestId() *string
	SetStorageGateways(v []*DescribeStorageGatewayResponseBodyStorageGateways) *DescribeStorageGatewayResponseBody
	GetStorageGateways() []*DescribeStorageGatewayResponseBodyStorageGateways
	SetTotalCount(v int32) *DescribeStorageGatewayResponseBody
	GetTotalCount() *int32
}

type DescribeStorageGatewayResponseBody struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6666C5A5-75ED-422E-A022-7121FA18C968
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list information.
	StorageGateways []*DescribeStorageGatewayResponseBodyStorageGateways `json:"StorageGateways,omitempty" xml:"StorageGateways,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 16
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeStorageGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStorageGatewayResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeStorageGatewayResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeStorageGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStorageGatewayResponseBody) GetStorageGateways() []*DescribeStorageGatewayResponseBodyStorageGateways {
	return s.StorageGateways
}

func (s *DescribeStorageGatewayResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeStorageGatewayResponseBody) SetPageNumber(v int32) *DescribeStorageGatewayResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeStorageGatewayResponseBody) SetPageSize(v int32) *DescribeStorageGatewayResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStorageGatewayResponseBody) SetRequestId(v string) *DescribeStorageGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStorageGatewayResponseBody) SetStorageGateways(v []*DescribeStorageGatewayResponseBodyStorageGateways) *DescribeStorageGatewayResponseBody {
	s.StorageGateways = v
	return s
}

func (s *DescribeStorageGatewayResponseBody) SetTotalCount(v int32) *DescribeStorageGatewayResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeStorageGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeStorageGatewayResponseBodyStorageGateways struct {
	// The internal CIDR block.
	//
	// example:
	//
	// 192.168.2.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The time when the storage gateway was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-05-14T03:07:47Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the storage gateway.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// cn-beijing-cmcc
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The IP address of the service.
	//
	// example:
	//
	// *.*.*.*
	ServiceIp *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	// The status of the storage gateway. Valid values:
	//
	// 	- creating
	//
	// 	- available
	//
	// 	- deleting
	//
	// 	- deleted
	//
	// example:
	//
	// available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the storage gateway.
	//
	// example:
	//
	// sgw-***
	StorageGatewayId *string `json:"StorageGatewayId,omitempty" xml:"StorageGatewayId,omitempty"`
	// The name of the storage gateway.
	//
	// example:
	//
	// testGateway
	StorageGatewayName *string `json:"StorageGatewayName,omitempty" xml:"StorageGatewayName,omitempty"`
	// The type of the storage gateway. Default value: 1, which indicates iSCSI.
	//
	// example:
	//
	// 1
	StorageGatewayType *int32 `json:"StorageGatewayType,omitempty" xml:"StorageGatewayType,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// n-***
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeStorageGatewayResponseBodyStorageGateways) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageGatewayResponseBodyStorageGateways) GoString() string {
	return s.String()
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) GetDescription() *string {
	return s.Description
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) GetServiceIp() *string {
	return s.ServiceIp
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) GetStatus() *string {
	return s.Status
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) GetStorageGatewayId() *string {
	return s.StorageGatewayId
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) GetStorageGatewayName() *string {
	return s.StorageGatewayName
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) GetStorageGatewayType() *int32 {
	return s.StorageGatewayType
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) SetCidrBlock(v string) *DescribeStorageGatewayResponseBodyStorageGateways {
	s.CidrBlock = &v
	return s
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) SetCreationTime(v string) *DescribeStorageGatewayResponseBodyStorageGateways {
	s.CreationTime = &v
	return s
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) SetDescription(v string) *DescribeStorageGatewayResponseBodyStorageGateways {
	s.Description = &v
	return s
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) SetEnsRegionId(v string) *DescribeStorageGatewayResponseBodyStorageGateways {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) SetServiceIp(v string) *DescribeStorageGatewayResponseBodyStorageGateways {
	s.ServiceIp = &v
	return s
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) SetStatus(v string) *DescribeStorageGatewayResponseBodyStorageGateways {
	s.Status = &v
	return s
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) SetStorageGatewayId(v string) *DescribeStorageGatewayResponseBodyStorageGateways {
	s.StorageGatewayId = &v
	return s
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) SetStorageGatewayName(v string) *DescribeStorageGatewayResponseBodyStorageGateways {
	s.StorageGatewayName = &v
	return s
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) SetStorageGatewayType(v int32) *DescribeStorageGatewayResponseBodyStorageGateways {
	s.StorageGatewayType = &v
	return s
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) SetVpcId(v string) *DescribeStorageGatewayResponseBodyStorageGateways {
	s.VpcId = &v
	return s
}

func (s *DescribeStorageGatewayResponseBodyStorageGateways) Validate() error {
	return dara.Validate(s)
}
