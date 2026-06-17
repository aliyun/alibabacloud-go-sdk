// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGatewayListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeGatewayListResponseBodyItems) *DescribeGatewayListResponseBody
	GetItems() []*DescribeGatewayListResponseBodyItems
	SetPageNumber(v string) *DescribeGatewayListResponseBody
	GetPageNumber() *string
	SetPageRecordCount(v string) *DescribeGatewayListResponseBody
	GetPageRecordCount() *string
	SetPageSize(v string) *DescribeGatewayListResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeGatewayListResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v string) *DescribeGatewayListResponseBody
	GetTotalRecordCount() *string
}

type DescribeGatewayListResponseBody struct {
	// A list of gateway instances.
	Items []*DescribeGatewayListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *string `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The number of entries per page.
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CED079B7-A408-41A1-BFF1-EC608E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalRecordCount *string `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeGatewayListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewayListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayListResponseBody) GetItems() []*DescribeGatewayListResponseBodyItems {
	return s.Items
}

func (s *DescribeGatewayListResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeGatewayListResponseBody) GetPageRecordCount() *string {
	return s.PageRecordCount
}

func (s *DescribeGatewayListResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeGatewayListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGatewayListResponseBody) GetTotalRecordCount() *string {
	return s.TotalRecordCount
}

func (s *DescribeGatewayListResponseBody) SetItems(v []*DescribeGatewayListResponseBodyItems) *DescribeGatewayListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeGatewayListResponseBody) SetPageNumber(v string) *DescribeGatewayListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGatewayListResponseBody) SetPageRecordCount(v string) *DescribeGatewayListResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeGatewayListResponseBody) SetPageSize(v string) *DescribeGatewayListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGatewayListResponseBody) SetRequestId(v string) *DescribeGatewayListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayListResponseBody) SetTotalRecordCount(v string) *DescribeGatewayListResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeGatewayListResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGatewayListResponseBodyItems struct {
	// The time when the gateway instance was created.
	//
	// example:
	//
	// 2020-08-14T05:58:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The database type.
	//
	// example:
	//
	// polardb_mysql
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The expiration time of the gateway instance.
	//
	// - For subscription instances, this parameter indicates the expiration time.
	//
	// - This parameter is empty for pay-as-you-go instances.
	//
	// example:
	//
	// 2028-09-01T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the gateway instance has expired. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// false
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The gateway instance ID.
	//
	// example:
	//
	// pg-xxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The description of the gateway instance.
	//
	// example:
	//
	// pg-xxxxxx
	GwDescription *string `json:"GwDescription,omitempty" xml:"GwDescription,omitempty"`
	// The time when the gateway instance was last modified.
	//
	// example:
	//
	// 2024-10-29T09:31:37Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The billing method. Valid values:
	//
	// - **Postpaid**: pay-as-you-go
	//
	// - **Prepaid**: subscription
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the gateway instance. Valid values:
	//
	// - **CREATE**: creating
	//
	// - **ACTIVATION**: running
	//
	// example:
	//
	// ACTIVATION
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-wz9u0v2yuskt1gth3uuju
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-****************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeGatewayListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewayListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeGatewayListResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGatewayListResponseBodyItems) GetDbType() *string {
	return s.DbType
}

func (s *DescribeGatewayListResponseBodyItems) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeGatewayListResponseBodyItems) GetExpired() *bool {
	return s.Expired
}

func (s *DescribeGatewayListResponseBodyItems) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeGatewayListResponseBodyItems) GetGwDescription() *string {
	return s.GwDescription
}

func (s *DescribeGatewayListResponseBodyItems) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeGatewayListResponseBodyItems) GetPayType() *string {
	return s.PayType
}

func (s *DescribeGatewayListResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGatewayListResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeGatewayListResponseBodyItems) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeGatewayListResponseBodyItems) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeGatewayListResponseBodyItems) SetCreateTime(v string) *DescribeGatewayListResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *DescribeGatewayListResponseBodyItems) SetDbType(v string) *DescribeGatewayListResponseBodyItems {
	s.DbType = &v
	return s
}

func (s *DescribeGatewayListResponseBodyItems) SetExpireTime(v string) *DescribeGatewayListResponseBodyItems {
	s.ExpireTime = &v
	return s
}

func (s *DescribeGatewayListResponseBodyItems) SetExpired(v bool) *DescribeGatewayListResponseBodyItems {
	s.Expired = &v
	return s
}

func (s *DescribeGatewayListResponseBodyItems) SetGwClusterId(v string) *DescribeGatewayListResponseBodyItems {
	s.GwClusterId = &v
	return s
}

func (s *DescribeGatewayListResponseBodyItems) SetGwDescription(v string) *DescribeGatewayListResponseBodyItems {
	s.GwDescription = &v
	return s
}

func (s *DescribeGatewayListResponseBodyItems) SetModifyTime(v string) *DescribeGatewayListResponseBodyItems {
	s.ModifyTime = &v
	return s
}

func (s *DescribeGatewayListResponseBodyItems) SetPayType(v string) *DescribeGatewayListResponseBodyItems {
	s.PayType = &v
	return s
}

func (s *DescribeGatewayListResponseBodyItems) SetRegionId(v string) *DescribeGatewayListResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeGatewayListResponseBodyItems) SetStatus(v string) *DescribeGatewayListResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeGatewayListResponseBodyItems) SetVSwitchId(v string) *DescribeGatewayListResponseBodyItems {
	s.VSwitchId = &v
	return s
}

func (s *DescribeGatewayListResponseBodyItems) SetVpcId(v string) *DescribeGatewayListResponseBodyItems {
	s.VpcId = &v
	return s
}

func (s *DescribeGatewayListResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
