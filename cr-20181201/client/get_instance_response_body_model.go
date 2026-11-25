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
	// The return code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The creation time, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1571926439000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance issue.
	//
	// example:
	//
	// The instance issue. Valid values:
	//
	// OSS_TOO_MANY_BUCKETS: The number of OSS buckets exceeds the limit.
	//
	// OSS_BUCKET_ALREADY_EXISTS: An OSS bucket with the same name already exists.
	//
	// OSS_SERVICE_ROLE_UNAUTHORIZED: OSS permissions are not authorized.
	//
	// USER_NOT_REGISTERED_BY_REAL_NAME: The China site account has not completed real-name verification when purchasing an OSS bucket on the China site
	InstanceIssue *string `json:"InstanceIssue,omitempty" xml:"InstanceIssue,omitempty"`
	// The instance name.
	//
	// example:
	//
	// shanghai-instance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance specification. Valid values:
	//
	// Enterprise_Basic: basic instance.
	//
	// Enterprise_Standard: standard instance.
	//
	// Enterprise_Advanced: advanced instance.
	//
	// example:
	//
	// Enterprise_Basic
	InstanceSpecification *string `json:"InstanceSpecification,omitempty" xml:"InstanceSpecification,omitempty"`
	// The instance status. Valid values:
	//
	// - `PENDING`: Initializing.
	//
	// - `INIT_ERROR`: Initialization failed.
	//
	// - `STARTING`: Starting.
	//
	// - `RUNNING`: Running.
	//
	// - `STOPPING`: Stopping.
	//
	// - `STOPPED`: Stopped.
	//
	// - `DELETING`: Deleting.
	//
	// - `DELETED`: Deleted.
	//
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - `true`: The call was successful.
	//
	// - `false`: The call failed.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The modification time, in UNIX timestamp format. Unit: milliseconds.
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
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmv36i4isx****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag collection of the instance.
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
