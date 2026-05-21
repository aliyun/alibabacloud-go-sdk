// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLeaderInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListLeaderInstancesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListLeaderInstancesResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *ListLeaderInstancesResponseBody
	GetHttpStatusCode() *string
	SetInstanceList(v []*ListLeaderInstancesResponseBodyInstanceList) *ListLeaderInstancesResponseBody
	GetInstanceList() []*ListLeaderInstancesResponseBodyInstanceList
	SetRequestId(v string) *ListLeaderInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ListLeaderInstancesResponseBody
	GetSuccess() *string
}

type ListLeaderInstancesResponseBody struct {
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceList   []*ListLeaderInstancesResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 5600196E-78E2-50F2-B2A1-C44D3B665438
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLeaderInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLeaderInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLeaderInstancesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListLeaderInstancesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListLeaderInstancesResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ListLeaderInstancesResponseBody) GetInstanceList() []*ListLeaderInstancesResponseBodyInstanceList {
	return s.InstanceList
}

func (s *ListLeaderInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLeaderInstancesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ListLeaderInstancesResponseBody) SetErrorCode(v string) *ListLeaderInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListLeaderInstancesResponseBody) SetErrorMessage(v string) *ListLeaderInstancesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListLeaderInstancesResponseBody) SetHttpStatusCode(v string) *ListLeaderInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListLeaderInstancesResponseBody) SetInstanceList(v []*ListLeaderInstancesResponseBodyInstanceList) *ListLeaderInstancesResponseBody {
	s.InstanceList = v
	return s
}

func (s *ListLeaderInstancesResponseBody) SetRequestId(v string) *ListLeaderInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLeaderInstancesResponseBody) SetSuccess(v string) *ListLeaderInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListLeaderInstancesResponseBody) Validate() error {
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

type ListLeaderInstancesResponseBodyInstanceList struct {
	// example:
	//
	// 0
	BindingCode *string `json:"BindingCode,omitempty" xml:"BindingCode,omitempty"`
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// hologram_combo_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// 2022-07-25T02:15:35Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// 2022-07-25T02:15:35Z
	ExpirationTime *string `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// example:
	//
	// hgpostcn-cn-tl32vsdir00h
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// user_defined_name
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// Running
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// example:
	//
	// Standard
	InstanceType *string                                            `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Tags         []*ListLeaderInstancesResponseBodyInstanceListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListLeaderInstancesResponseBodyInstanceList) String() string {
	return dara.Prettify(s)
}

func (s ListLeaderInstancesResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *ListLeaderInstancesResponseBodyInstanceList) GetBindingCode() *string {
	return s.BindingCode
}

func (s *ListLeaderInstancesResponseBodyInstanceList) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListLeaderInstancesResponseBodyInstanceList) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListLeaderInstancesResponseBodyInstanceList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListLeaderInstancesResponseBodyInstanceList) GetExpirationTime() *string {
	return s.ExpirationTime
}

func (s *ListLeaderInstancesResponseBodyInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListLeaderInstancesResponseBodyInstanceList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListLeaderInstancesResponseBodyInstanceList) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *ListLeaderInstancesResponseBodyInstanceList) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListLeaderInstancesResponseBodyInstanceList) GetTags() []*ListLeaderInstancesResponseBodyInstanceListTags {
	return s.Tags
}

func (s *ListLeaderInstancesResponseBodyInstanceList) SetBindingCode(v string) *ListLeaderInstancesResponseBodyInstanceList {
	s.BindingCode = &v
	return s
}

func (s *ListLeaderInstancesResponseBodyInstanceList) SetChargeType(v string) *ListLeaderInstancesResponseBodyInstanceList {
	s.ChargeType = &v
	return s
}

func (s *ListLeaderInstancesResponseBodyInstanceList) SetCommodityCode(v string) *ListLeaderInstancesResponseBodyInstanceList {
	s.CommodityCode = &v
	return s
}

func (s *ListLeaderInstancesResponseBodyInstanceList) SetCreationTime(v string) *ListLeaderInstancesResponseBodyInstanceList {
	s.CreationTime = &v
	return s
}

func (s *ListLeaderInstancesResponseBodyInstanceList) SetExpirationTime(v string) *ListLeaderInstancesResponseBodyInstanceList {
	s.ExpirationTime = &v
	return s
}

func (s *ListLeaderInstancesResponseBodyInstanceList) SetInstanceId(v string) *ListLeaderInstancesResponseBodyInstanceList {
	s.InstanceId = &v
	return s
}

func (s *ListLeaderInstancesResponseBodyInstanceList) SetInstanceName(v string) *ListLeaderInstancesResponseBodyInstanceList {
	s.InstanceName = &v
	return s
}

func (s *ListLeaderInstancesResponseBodyInstanceList) SetInstanceStatus(v string) *ListLeaderInstancesResponseBodyInstanceList {
	s.InstanceStatus = &v
	return s
}

func (s *ListLeaderInstancesResponseBodyInstanceList) SetInstanceType(v string) *ListLeaderInstancesResponseBodyInstanceList {
	s.InstanceType = &v
	return s
}

func (s *ListLeaderInstancesResponseBodyInstanceList) SetTags(v []*ListLeaderInstancesResponseBodyInstanceListTags) *ListLeaderInstancesResponseBodyInstanceList {
	s.Tags = v
	return s
}

func (s *ListLeaderInstancesResponseBodyInstanceList) Validate() error {
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

type ListLeaderInstancesResponseBodyInstanceListTags struct {
	// example:
	//
	// tagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListLeaderInstancesResponseBodyInstanceListTags) String() string {
	return dara.Prettify(s)
}

func (s ListLeaderInstancesResponseBodyInstanceListTags) GoString() string {
	return s.String()
}

func (s *ListLeaderInstancesResponseBodyInstanceListTags) GetKey() *string {
	return s.Key
}

func (s *ListLeaderInstancesResponseBodyInstanceListTags) GetValue() *string {
	return s.Value
}

func (s *ListLeaderInstancesResponseBodyInstanceListTags) SetKey(v string) *ListLeaderInstancesResponseBodyInstanceListTags {
	s.Key = &v
	return s
}

func (s *ListLeaderInstancesResponseBodyInstanceListTags) SetValue(v string) *ListLeaderInstancesResponseBodyInstanceListTags {
	s.Value = &v
	return s
}

func (s *ListLeaderInstancesResponseBodyInstanceListTags) Validate() error {
	return dara.Validate(s)
}
