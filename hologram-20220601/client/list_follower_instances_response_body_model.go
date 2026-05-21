// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFollowerInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListFollowerInstancesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListFollowerInstancesResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *ListFollowerInstancesResponseBody
	GetHttpStatusCode() *string
	SetInstanceList(v []*ListFollowerInstancesResponseBodyInstanceList) *ListFollowerInstancesResponseBody
	GetInstanceList() []*ListFollowerInstancesResponseBodyInstanceList
	SetRequestId(v string) *ListFollowerInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ListFollowerInstancesResponseBody
	GetSuccess() *string
}

type ListFollowerInstancesResponseBody struct {
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string                                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceList   []*ListFollowerInstancesResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// D3AE84AB-0873-5FC7-A4C4-8CF869D2FA70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFollowerInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFollowerInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFollowerInstancesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListFollowerInstancesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListFollowerInstancesResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ListFollowerInstancesResponseBody) GetInstanceList() []*ListFollowerInstancesResponseBodyInstanceList {
	return s.InstanceList
}

func (s *ListFollowerInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFollowerInstancesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ListFollowerInstancesResponseBody) SetErrorCode(v string) *ListFollowerInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListFollowerInstancesResponseBody) SetErrorMessage(v string) *ListFollowerInstancesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListFollowerInstancesResponseBody) SetHttpStatusCode(v string) *ListFollowerInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListFollowerInstancesResponseBody) SetInstanceList(v []*ListFollowerInstancesResponseBodyInstanceList) *ListFollowerInstancesResponseBody {
	s.InstanceList = v
	return s
}

func (s *ListFollowerInstancesResponseBody) SetRequestId(v string) *ListFollowerInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFollowerInstancesResponseBody) SetSuccess(v string) *ListFollowerInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListFollowerInstancesResponseBody) Validate() error {
	if s.InstanceList != nil {
		for _, item := range s.InstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFollowerInstancesResponseBodyInstanceList struct {
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// hologram_postpay_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// 2022-12-16T02:24:05Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// 2023-05-04T16:00:00.000Z
	ExpirationTime *string `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// example:
	//
	// hgpostcn-cn-aaab9ad2d8fb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// test_instance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// Running
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// example:
	//
	// Standard
	InstanceType *string                                              `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Tags         []*ListFollowerInstancesResponseBodyInstanceListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListFollowerInstancesResponseBodyInstanceList) String() string {
	return dara.Prettify(s)
}

func (s ListFollowerInstancesResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *ListFollowerInstancesResponseBodyInstanceList) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListFollowerInstancesResponseBodyInstanceList) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListFollowerInstancesResponseBodyInstanceList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListFollowerInstancesResponseBodyInstanceList) GetExpirationTime() *string {
	return s.ExpirationTime
}

func (s *ListFollowerInstancesResponseBodyInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFollowerInstancesResponseBodyInstanceList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListFollowerInstancesResponseBodyInstanceList) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *ListFollowerInstancesResponseBodyInstanceList) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListFollowerInstancesResponseBodyInstanceList) GetTags() []*ListFollowerInstancesResponseBodyInstanceListTags {
	return s.Tags
}

func (s *ListFollowerInstancesResponseBodyInstanceList) SetChargeType(v string) *ListFollowerInstancesResponseBodyInstanceList {
	s.ChargeType = &v
	return s
}

func (s *ListFollowerInstancesResponseBodyInstanceList) SetCommodityCode(v string) *ListFollowerInstancesResponseBodyInstanceList {
	s.CommodityCode = &v
	return s
}

func (s *ListFollowerInstancesResponseBodyInstanceList) SetCreationTime(v string) *ListFollowerInstancesResponseBodyInstanceList {
	s.CreationTime = &v
	return s
}

func (s *ListFollowerInstancesResponseBodyInstanceList) SetExpirationTime(v string) *ListFollowerInstancesResponseBodyInstanceList {
	s.ExpirationTime = &v
	return s
}

func (s *ListFollowerInstancesResponseBodyInstanceList) SetInstanceId(v string) *ListFollowerInstancesResponseBodyInstanceList {
	s.InstanceId = &v
	return s
}

func (s *ListFollowerInstancesResponseBodyInstanceList) SetInstanceName(v string) *ListFollowerInstancesResponseBodyInstanceList {
	s.InstanceName = &v
	return s
}

func (s *ListFollowerInstancesResponseBodyInstanceList) SetInstanceStatus(v string) *ListFollowerInstancesResponseBodyInstanceList {
	s.InstanceStatus = &v
	return s
}

func (s *ListFollowerInstancesResponseBodyInstanceList) SetInstanceType(v string) *ListFollowerInstancesResponseBodyInstanceList {
	s.InstanceType = &v
	return s
}

func (s *ListFollowerInstancesResponseBodyInstanceList) SetTags(v []*ListFollowerInstancesResponseBodyInstanceListTags) *ListFollowerInstancesResponseBodyInstanceList {
	s.Tags = v
	return s
}

func (s *ListFollowerInstancesResponseBodyInstanceList) Validate() error {
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

type ListFollowerInstancesResponseBodyInstanceListTags struct {
	// example:
	//
	// tag
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFollowerInstancesResponseBodyInstanceListTags) String() string {
	return dara.Prettify(s)
}

func (s ListFollowerInstancesResponseBodyInstanceListTags) GoString() string {
	return s.String()
}

func (s *ListFollowerInstancesResponseBodyInstanceListTags) GetKey() *string {
	return s.Key
}

func (s *ListFollowerInstancesResponseBodyInstanceListTags) GetValue() *string {
	return s.Value
}

func (s *ListFollowerInstancesResponseBodyInstanceListTags) SetKey(v string) *ListFollowerInstancesResponseBodyInstanceListTags {
	s.Key = &v
	return s
}

func (s *ListFollowerInstancesResponseBodyInstanceListTags) SetValue(v string) *ListFollowerInstancesResponseBodyInstanceListTags {
	s.Value = &v
	return s
}

func (s *ListFollowerInstancesResponseBodyInstanceListTags) Validate() error {
	return dara.Validate(s)
}
