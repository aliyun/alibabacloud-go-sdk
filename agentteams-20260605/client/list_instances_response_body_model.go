// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstancesResponseBody
	GetCode() *string
	SetItems(v []*ListInstancesResponseBodyItems) *ListInstancesResponseBody
	GetItems() []*ListInstancesResponseBodyItems
	SetMaxResults(v int32) *ListInstancesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListInstancesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListInstancesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstancesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListInstancesResponseBody
	GetTotalCount() *int32
}

type ListInstancesResponseBody struct {
	// example:
	//
	// SUCCESS
	Code  *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Items []*ListInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 10
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// request-1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 23
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstancesResponseBody) GetItems() []*ListInstancesResponseBodyItems {
	return s.Items
}

func (s *ListInstancesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInstancesResponseBody) SetCode(v string) *ListInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesResponseBody) SetItems(v []*ListInstancesResponseBodyItems) *ListInstancesResponseBody {
	s.Items = v
	return s
}

func (s *ListInstancesResponseBody) SetMaxResults(v int32) *ListInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInstancesResponseBody) SetMessage(v string) *ListInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesResponseBody) SetNextToken(v string) *ListInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetSuccess(v bool) *ListInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int32) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
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

type ListInstancesResponseBodyItems struct {
	InstanceId    *string                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName  *string                                `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceSpec  *string                                `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	OssBucketName *string                                `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	RegionId      *string                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroup *string                                `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	Status        *string                                `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId         *string                                `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Zones         []*ListInstancesResponseBodyItemsZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyItems) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstancesResponseBodyItems) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *ListInstancesResponseBodyItems) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *ListInstancesResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancesResponseBodyItems) GetSecurityGroup() *string {
	return s.SecurityGroup
}

func (s *ListInstancesResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyItems) GetVpcId() *string {
	return s.VpcId
}

func (s *ListInstancesResponseBodyItems) GetZones() []*ListInstancesResponseBodyItemsZones {
	return s.Zones
}

func (s *ListInstancesResponseBodyItems) SetInstanceId(v string) *ListInstancesResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyItems) SetInstanceName(v string) *ListInstancesResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyItems) SetInstanceSpec(v string) *ListInstancesResponseBodyItems {
	s.InstanceSpec = &v
	return s
}

func (s *ListInstancesResponseBodyItems) SetOssBucketName(v string) *ListInstancesResponseBodyItems {
	s.OssBucketName = &v
	return s
}

func (s *ListInstancesResponseBodyItems) SetRegionId(v string) *ListInstancesResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *ListInstancesResponseBodyItems) SetSecurityGroup(v string) *ListInstancesResponseBodyItems {
	s.SecurityGroup = &v
	return s
}

func (s *ListInstancesResponseBodyItems) SetStatus(v string) *ListInstancesResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyItems) SetVpcId(v string) *ListInstancesResponseBodyItems {
	s.VpcId = &v
	return s
}

func (s *ListInstancesResponseBodyItems) SetZones(v []*ListInstancesResponseBodyItemsZones) *ListInstancesResponseBodyItems {
	s.Zones = v
	return s
}

func (s *ListInstancesResponseBodyItems) Validate() error {
	if s.Zones != nil {
		for _, item := range s.Zones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyItemsZones struct {
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListInstancesResponseBodyItemsZones) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyItemsZones) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyItemsZones) GetVswitchId() *string {
	return s.VswitchId
}

func (s *ListInstancesResponseBodyItemsZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListInstancesResponseBodyItemsZones) SetVswitchId(v string) *ListInstancesResponseBodyItemsZones {
	s.VswitchId = &v
	return s
}

func (s *ListInstancesResponseBodyItemsZones) SetZoneId(v string) *ListInstancesResponseBodyItemsZones {
	s.ZoneId = &v
	return s
}

func (s *ListInstancesResponseBodyItemsZones) Validate() error {
	return dara.Validate(s)
}
