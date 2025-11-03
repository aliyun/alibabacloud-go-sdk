// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceResponseBody
	GetCode() *string
	SetCreateTime(v int64) *GetInstanceResponseBody
	GetCreateTime() *int64
	SetInstanceId(v string) *GetInstanceResponseBody
	GetInstanceId() *string
	SetInstanceIssue(v string) *GetInstanceResponseBody
	GetInstanceIssue() *string
	SetInstanceName(v string) *GetInstanceResponseBody
	GetInstanceName() *string
	SetInstanceSpecification(v string) *GetInstanceResponseBody
	GetInstanceSpecification() *string
	SetInstanceStatus(v string) *GetInstanceResponseBody
	GetInstanceStatus() *string
	SetIsSuccess(v bool) *GetInstanceResponseBody
	GetIsSuccess() *bool
	SetModifiedTime(v int64) *GetInstanceResponseBody
	GetModifiedTime() *int64
	SetRequestId(v string) *GetInstanceResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetInstanceResponseBody
	GetResourceGroupId() *string
	SetTags(v []*GetInstanceResponseBodyTags) *GetInstanceResponseBody
	GetTags() []*GetInstanceResponseBodyTags
}

type GetInstanceResponseBody struct {
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 1571926439000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the Container Registry instance.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The issue occurs on the instance.
	//
	// example:
	//
	// The issue occurs on the instance. Valid values: OSS_TOO_MANY_BUCKETS: The number of Object Storage Service (OSS) buckets exceeds the upper limit. OSS_BUCKET_ALREADY_EXISTS: An OSS bucket that has the duplicate name already exists. OSS_SERVICE_ROLE_UNAUTHORIZED: The OSS service-linked role is not granted permissions. USER_NOT_REGISTERED_BY_REAL_NAME: The Alibaba Cloud account has not passed a real name verification.
	InstanceIssue *string `json:"InstanceIssue,omitempty" xml:"InstanceIssue,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// shanghai-instance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The edition of the instance. Valid values: Enterprise_Basic: Basic Edition instances Enterprise_Standard: Standard Edition instances Enterprise_Advanced: Advanced Edition instances
	//
	// example:
	//
	// Enterprise_Basic
	InstanceSpecification *string `json:"InstanceSpecification,omitempty" xml:"InstanceSpecification,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- `PENDING`: The instance is being initialized.
	//
	// 	- `INIT_ERROR`: The instance failed to be initialized.
	//
	// 	- `STARTING`: The instance is being started.
	//
	// 	- `RUNNING`: The instance is running.
	//
	// 	- `STOPPING`: The instance is being stopped.
	//
	// 	- `STOPPED`: The instance is stopped.
	//
	// 	- `DELETING`: The instance is being deleted.
	//
	// 	- `DELETED`: The instance is deleted.
	//
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The time when the instance was last modified.
	//
	// example:
	//
	// 1571926560000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6EF34B18-4228-470C-860C-D28597CF010E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmv36i4isx****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the instance.
	Tags []*GetInstanceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceResponseBody) GetInstanceIssue() *string {
	return s.InstanceIssue
}

func (s *GetInstanceResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetInstanceResponseBody) GetInstanceSpecification() *string {
	return s.InstanceSpecification
}

func (s *GetInstanceResponseBody) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *GetInstanceResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetInstanceResponseBody) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *GetInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetInstanceResponseBody) GetTags() []*GetInstanceResponseBodyTags {
	return s.Tags
}

func (s *GetInstanceResponseBody) SetCode(v string) *GetInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceResponseBody) SetCreateTime(v int64) *GetInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceId(v string) *GetInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceIssue(v string) *GetInstanceResponseBody {
	s.InstanceIssue = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceName(v string) *GetInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceSpecification(v string) *GetInstanceResponseBody {
	s.InstanceSpecification = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceStatus(v string) *GetInstanceResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstanceResponseBody) SetIsSuccess(v bool) *GetInstanceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetInstanceResponseBody) SetModifiedTime(v int64) *GetInstanceResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetResourceGroupId(v string) *GetInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetInstanceResponseBody) SetTags(v []*GetInstanceResponseBodyTags) *GetInstanceResponseBody {
	s.Tags = v
	return s
}

func (s *GetInstanceResponseBody) Validate() error {
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

type GetInstanceResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// test_key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test_value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetInstanceResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetInstanceResponseBodyTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetInstanceResponseBodyTags) SetTagKey(v string) *GetInstanceResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *GetInstanceResponseBodyTags) SetTagValue(v string) *GetInstanceResponseBodyTags {
	s.TagValue = &v
	return s
}

func (s *GetInstanceResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
