// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayConsumerForPolarDBXRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateGatewayConsumerForPolarDBXRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *CreateGatewayConsumerForPolarDBXRequest
	GetRegionId() *string
}

type CreateGatewayConsumerForPolarDBXRequest struct {
	// The instance ID. > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the details of all instances in the specified region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateGatewayConsumerForPolarDBXRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayConsumerForPolarDBXRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayConsumerForPolarDBXRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateGatewayConsumerForPolarDBXRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateGatewayConsumerForPolarDBXRequest) SetDBInstanceName(v string) *CreateGatewayConsumerForPolarDBXRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateGatewayConsumerForPolarDBXRequest) SetRegionId(v string) *CreateGatewayConsumerForPolarDBXRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGatewayConsumerForPolarDBXRequest) Validate() error {
	return dara.Validate(s)
}
