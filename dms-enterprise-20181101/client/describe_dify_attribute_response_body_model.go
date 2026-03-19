// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDifyAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDifyAttributeResponseBody
	GetCode() *string
	SetErrorCode(v string) *DescribeDifyAttributeResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *DescribeDifyAttributeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeDifyAttributeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDifyAttributeResponseBody
	GetRequestId() *string
	SetRoot(v *DescribeDifyAttributeResponseBodyRoot) *DescribeDifyAttributeResponseBody
	GetRoot() *DescribeDifyAttributeResponseBodyRoot
	SetSuccess(v bool) *DescribeDifyAttributeResponseBody
	GetSuccess() *bool
	SetTags(v []*DescribeDifyAttributeResponseBodyTags) *DescribeDifyAttributeResponseBody
	GetTags() []*DescribeDifyAttributeResponseBodyTags
}

type DescribeDifyAttributeResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// UnknownError
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 427688B8-ADFB-4C4E-9D45-EF5C1FD6E23D
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Root      *DescribeDifyAttributeResponseBodyRoot `json:"Root,omitempty" xml:"Root,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	Tags    []*DescribeDifyAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeDifyAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDifyAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDifyAttributeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDifyAttributeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeDifyAttributeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDifyAttributeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDifyAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDifyAttributeResponseBody) GetRoot() *DescribeDifyAttributeResponseBodyRoot {
	return s.Root
}

func (s *DescribeDifyAttributeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDifyAttributeResponseBody) GetTags() []*DescribeDifyAttributeResponseBodyTags {
	return s.Tags
}

func (s *DescribeDifyAttributeResponseBody) SetCode(v string) *DescribeDifyAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDifyAttributeResponseBody) SetErrorCode(v string) *DescribeDifyAttributeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeDifyAttributeResponseBody) SetHttpStatusCode(v int32) *DescribeDifyAttributeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDifyAttributeResponseBody) SetMessage(v string) *DescribeDifyAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDifyAttributeResponseBody) SetRequestId(v string) *DescribeDifyAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDifyAttributeResponseBody) SetRoot(v *DescribeDifyAttributeResponseBodyRoot) *DescribeDifyAttributeResponseBody {
	s.Root = v
	return s
}

func (s *DescribeDifyAttributeResponseBody) SetSuccess(v bool) *DescribeDifyAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDifyAttributeResponseBody) SetTags(v []*DescribeDifyAttributeResponseBodyTags) *DescribeDifyAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeDifyAttributeResponseBody) Validate() error {
	if s.Root != nil {
		if err := s.Root.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDifyAttributeResponseBodyRoot struct {
	// example:
	//
	// DIFY
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// 92748163-af62-4ca4-ad85-1****
	AppUuid *string `json:"AppUuid,omitempty" xml:"AppUuid,omitempty"`
	// example:
	//
	// ABCD
	BillingInstanceId *string `json:"BillingInstanceId,omitempty" xml:"BillingInstanceId,omitempty"`
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// abc1-def2-ghi3-jkl4
	DifyInstanceId *string `json:"DifyInstanceId,omitempty" xml:"DifyInstanceId,omitempty"`
	// example:
	//
	// This is dify instance
	DifyInstanceName *string `json:"DifyInstanceName,omitempty" xml:"DifyInstanceName,omitempty"`
	// example:
	//
	// 20251201
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// cn-chengdu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1
	Replicas *string `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// example:
	//
	// 4CU
	ResourceQuota *string `json:"ResourceQuota,omitempty" xml:"ResourceQuota,omitempty"`
	// example:
	//
	// sg-bp1ik7t5d5f24b****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// DEPLOYED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// ESSD
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// example:
	//
	// vsw-bp1tzpv5jfsuoqy****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-bp1n16nsg8z1kn6****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// 339170706****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDifyAttributeResponseBodyRoot) String() string {
	return dara.Prettify(s)
}

func (s DescribeDifyAttributeResponseBodyRoot) GoString() string {
	return s.String()
}

func (s *DescribeDifyAttributeResponseBodyRoot) GetAppType() *string {
	return s.AppType
}

func (s *DescribeDifyAttributeResponseBodyRoot) GetAppUuid() *string {
	return s.AppUuid
}

func (s *DescribeDifyAttributeResponseBodyRoot) GetBillingInstanceId() *string {
	return s.BillingInstanceId
}

func (s *DescribeDifyAttributeResponseBodyRoot) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDifyAttributeResponseBodyRoot) GetDifyInstanceId() *string {
	return s.DifyInstanceId
}

func (s *DescribeDifyAttributeResponseBodyRoot) GetDifyInstanceName() *string {
	return s.DifyInstanceName
}

func (s *DescribeDifyAttributeResponseBodyRoot) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeDifyAttributeResponseBodyRoot) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDifyAttributeResponseBodyRoot) GetReplicas() *string {
	return s.Replicas
}

func (s *DescribeDifyAttributeResponseBodyRoot) GetResourceQuota() *string {
	return s.ResourceQuota
}

func (s *DescribeDifyAttributeResponseBodyRoot) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeDifyAttributeResponseBodyRoot) GetStatus() *string {
	return s.Status
}

func (s *DescribeDifyAttributeResponseBodyRoot) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDifyAttributeResponseBodyRoot) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDifyAttributeResponseBodyRoot) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDifyAttributeResponseBodyRoot) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeDifyAttributeResponseBodyRoot) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDifyAttributeResponseBodyRoot) SetAppType(v string) *DescribeDifyAttributeResponseBodyRoot {
	s.AppType = &v
	return s
}

func (s *DescribeDifyAttributeResponseBodyRoot) SetAppUuid(v string) *DescribeDifyAttributeResponseBodyRoot {
	s.AppUuid = &v
	return s
}

func (s *DescribeDifyAttributeResponseBodyRoot) SetBillingInstanceId(v string) *DescribeDifyAttributeResponseBodyRoot {
	s.BillingInstanceId = &v
	return s
}

func (s *DescribeDifyAttributeResponseBodyRoot) SetChargeType(v string) *DescribeDifyAttributeResponseBodyRoot {
	s.ChargeType = &v
	return s
}

func (s *DescribeDifyAttributeResponseBodyRoot) SetDifyInstanceId(v string) *DescribeDifyAttributeResponseBodyRoot {
	s.DifyInstanceId = &v
	return s
}

func (s *DescribeDifyAttributeResponseBodyRoot) SetDifyInstanceName(v string) *DescribeDifyAttributeResponseBodyRoot {
	s.DifyInstanceName = &v
	return s
}

func (s *DescribeDifyAttributeResponseBodyRoot) SetExpireTime(v int64) *DescribeDifyAttributeResponseBodyRoot {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDifyAttributeResponseBodyRoot) SetRegionId(v string) *DescribeDifyAttributeResponseBodyRoot {
	s.RegionId = &v
	return s
}

func (s *DescribeDifyAttributeResponseBodyRoot) SetReplicas(v string) *DescribeDifyAttributeResponseBodyRoot {
	s.Replicas = &v
	return s
}

func (s *DescribeDifyAttributeResponseBodyRoot) SetResourceQuota(v string) *DescribeDifyAttributeResponseBodyRoot {
	s.ResourceQuota = &v
	return s
}

func (s *DescribeDifyAttributeResponseBodyRoot) SetSecurityGroupId(v string) *DescribeDifyAttributeResponseBodyRoot {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeDifyAttributeResponseBodyRoot) SetStatus(v string) *DescribeDifyAttributeResponseBodyRoot {
	s.Status = &v
	return s
}

func (s *DescribeDifyAttributeResponseBodyRoot) SetStorageType(v string) *DescribeDifyAttributeResponseBodyRoot {
	s.StorageType = &v
	return s
}

func (s *DescribeDifyAttributeResponseBodyRoot) SetVSwitchId(v string) *DescribeDifyAttributeResponseBodyRoot {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDifyAttributeResponseBodyRoot) SetVpcId(v string) *DescribeDifyAttributeResponseBodyRoot {
	s.VpcId = &v
	return s
}

func (s *DescribeDifyAttributeResponseBodyRoot) SetWorkspaceId(v string) *DescribeDifyAttributeResponseBodyRoot {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeDifyAttributeResponseBodyRoot) SetZoneId(v string) *DescribeDifyAttributeResponseBodyRoot {
	s.ZoneId = &v
	return s
}

func (s *DescribeDifyAttributeResponseBodyRoot) Validate() error {
	return dara.Validate(s)
}

type DescribeDifyAttributeResponseBodyTags struct {
	// example:
	//
	// Key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// Value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDifyAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDifyAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeDifyAttributeResponseBodyTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeDifyAttributeResponseBodyTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeDifyAttributeResponseBodyTags) SetTagKey(v string) *DescribeDifyAttributeResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *DescribeDifyAttributeResponseBodyTags) SetTagValue(v string) *DescribeDifyAttributeResponseBodyTags {
	s.TagValue = &v
	return s
}

func (s *DescribeDifyAttributeResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
