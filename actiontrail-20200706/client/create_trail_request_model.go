// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventRW(v string) *CreateTrailRequest
	GetEventRW() *string
	SetIsOrganizationTrail(v bool) *CreateTrailRequest
	GetIsOrganizationTrail() *bool
	SetMaxComputeProjectArn(v string) *CreateTrailRequest
	GetMaxComputeProjectArn() *string
	SetMaxComputeWriteRoleArn(v string) *CreateTrailRequest
	GetMaxComputeWriteRoleArn() *string
	SetName(v string) *CreateTrailRequest
	GetName() *string
	SetOssBucketName(v string) *CreateTrailRequest
	GetOssBucketName() *string
	SetOssKeyPrefix(v string) *CreateTrailRequest
	GetOssKeyPrefix() *string
	SetOssWriteRoleArn(v string) *CreateTrailRequest
	GetOssWriteRoleArn() *string
	SetSlsProjectArn(v string) *CreateTrailRequest
	GetSlsProjectArn() *string
	SetSlsWriteRoleArn(v string) *CreateTrailRequest
	GetSlsWriteRoleArn() *string
	SetTrailRegion(v string) *CreateTrailRequest
	GetTrailRegion() *string
}

type CreateTrailRequest struct {
	// The read/write type of the events to be delivered. Valid values:
	//
	// 	- Write: write events. It is the default value.
	//
	// 	- Read: read events.
	//
	// 	- All: read and write events.
	//
	// example:
	//
	// Write
	EventRW *string `json:"EventRW,omitempty" xml:"EventRW,omitempty"`
	// Specifies whether to create a multi-account trail. Valid values:
	//
	// 	- true: creates a multi-account trail.
	//
	// 	- false (default): creates a single-account trail.
	//
	// example:
	//
	// false
	IsOrganizationTrail *bool `json:"IsOrganizationTrail,omitempty" xml:"IsOrganizationTrail,omitempty"`
	// The ARN of the MaxCompute project to which you want to deliver events.
	//
	// >  You must specify at least one of the following parameters: OssBucketName, SlsProjectArn, and MaxComputeProjectArn.
	//
	// >  The name of the MaxCompute project must be prefixed with actiontrail_.
	//
	// example:
	//
	// acs:odps:cn-hangzhou:15127787691****:project/actiontrail_****
	MaxComputeProjectArn *string `json:"MaxComputeProjectArn,omitempty" xml:"MaxComputeProjectArn,omitempty"`
	// The ARN of the role that is assumed by ActionTrail to deliver events to the MaxCompute project.
	//
	// 	- If you do not specify this parameter, ActionTrail creates a service-linked role to create the corresponding resource. For more information, see [Manage the service-linked role](https://help.aliyun.com/document_detail/169244.html).
	//
	// 	- If you specify this parameter and deliver events to the current account, you must grant the RAM role the permissions on the service-linked role for ActionTrail. If you want to deliver events to other accounts, you must attach a system policy to the RAM role. For more information about how to deliver events across Alibaba Cloud accounts, see [Deliver events across Alibaba Cloud accounts](https://help.aliyun.com/document_detail/207462.html).
	//
	// example:
	//
	// acs:ram::15127787691****:role/aliyunserviceroleforactiontrail
	MaxComputeWriteRoleArn *string `json:"MaxComputeWriteRoleArn,omitempty" xml:"MaxComputeWriteRoleArn,omitempty"`
	// The name of the trail to be created.
	//
	// The name must be 6 to 36 characters in length. The name must start with a lowercase letter and can contain lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// > The name must be unique within your Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// trail-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the OSS bucket to which events are to be delivered.
	//
	// The name must be 3 to 63 characters in length. The name must start with a lowercase letter or a digit and can contain lowercase letters, digits, and hyphens (-).
	//
	// > You must specify at least one of the OssBucketName and SlsProjectArn parameters.
	//
	// example:
	//
	// audit-log
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The prefix of the log files to be stored in the destination OSS bucket. This parameter can be left empty.
	//
	// The prefix must be 6 to 32 characters in length. The prefix must start with a letter and can contain letters, digits, hyphens (-), forward slashes (/), and underscores (_).
	//
	// example:
	//
	// at-product-account-audit-B
	OssKeyPrefix *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the service-linked role that is assumed by ActionTrail to deliver events to the destination Object Storage Service (OSS) bucket.
	//
	// 	- If you do not specify this parameter, ActionTrail creates a service-linked role to create the corresponding resource. For more information, see [Manage the service-linked role](https://help.aliyun.com/document_detail/169244.html).
	//
	// 	- If you specify this parameter and deliver events to the current account, you must grant the RAM role the permissions on the service-linked role for ActionTrail. If you want to deliver events to other accounts, you must attach a system policy to the RAM role. For more information about how to deliver events across Alibaba Cloud accounts, see [Deliver events across Alibaba Cloud accounts](https://help.aliyun.com/document_detail/207462.html).
	//
	// example:
	//
	// acs:ram::***:role/aliyunserviceroleforactiontrail
	OssWriteRoleArn *string `json:"OssWriteRoleArn,omitempty" xml:"OssWriteRoleArn,omitempty"`
	// The ARN of the Log Service project to which events are to be delivered.
	//
	// > You must specify at least one of the OssBucketName and SlsProjectArn parameters.
	//
	// example:
	//
	// acs:log:cn-shanghai::project/***
	SlsProjectArn *string `json:"SlsProjectArn,omitempty" xml:"SlsProjectArn,omitempty"`
	// The ARN of the RAM role that is assumed by ActionTrail to deliver events to the Log Service project.
	//
	// 	- If you do not specify this parameter, ActionTrail creates a service-linked role to create the corresponding resource. For more information, see [Manage the service-linked role](https://help.aliyun.com/document_detail/169244.html).
	//
	// 	- If you specify this parameter, you must grant the permissions of the service-linked role that is assumed by ActionTrail to the RAM role before you can deliver events to your Alibaba Cloud account. If you need to deliver events to other Alibaba Cloud accounts, you must attach the permission policy that is used to grant permissions related to event delivery to the RAM role. For more information about how to deliver events across Alibaba Cloud accounts, see [Deliver events across Alibaba Cloud accounts](https://help.aliyun.com/document_detail/207462.html).
	//
	// example:
	//
	// acs:ram::***:role/aliyunserviceroleforactiontrail
	SlsWriteRoleArn *string `json:"SlsWriteRoleArn,omitempty" xml:"SlsWriteRoleArn,omitempty"`
	// The one or more regions from which the trail delivers events.
	//
	// The default value is All, which indicates that the trail delivers events from all regions.
	//
	// You can also specify specific regions. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/213597.html) operation to query all the supported regions.
	//
	// example:
	//
	// All
	TrailRegion *string `json:"TrailRegion,omitempty" xml:"TrailRegion,omitempty"`
}

func (s CreateTrailRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTrailRequest) GoString() string {
	return s.String()
}

func (s *CreateTrailRequest) GetEventRW() *string {
	return s.EventRW
}

func (s *CreateTrailRequest) GetIsOrganizationTrail() *bool {
	return s.IsOrganizationTrail
}

func (s *CreateTrailRequest) GetMaxComputeProjectArn() *string {
	return s.MaxComputeProjectArn
}

func (s *CreateTrailRequest) GetMaxComputeWriteRoleArn() *string {
	return s.MaxComputeWriteRoleArn
}

func (s *CreateTrailRequest) GetName() *string {
	return s.Name
}

func (s *CreateTrailRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *CreateTrailRequest) GetOssKeyPrefix() *string {
	return s.OssKeyPrefix
}

func (s *CreateTrailRequest) GetOssWriteRoleArn() *string {
	return s.OssWriteRoleArn
}

func (s *CreateTrailRequest) GetSlsProjectArn() *string {
	return s.SlsProjectArn
}

func (s *CreateTrailRequest) GetSlsWriteRoleArn() *string {
	return s.SlsWriteRoleArn
}

func (s *CreateTrailRequest) GetTrailRegion() *string {
	return s.TrailRegion
}

func (s *CreateTrailRequest) SetEventRW(v string) *CreateTrailRequest {
	s.EventRW = &v
	return s
}

func (s *CreateTrailRequest) SetIsOrganizationTrail(v bool) *CreateTrailRequest {
	s.IsOrganizationTrail = &v
	return s
}

func (s *CreateTrailRequest) SetMaxComputeProjectArn(v string) *CreateTrailRequest {
	s.MaxComputeProjectArn = &v
	return s
}

func (s *CreateTrailRequest) SetMaxComputeWriteRoleArn(v string) *CreateTrailRequest {
	s.MaxComputeWriteRoleArn = &v
	return s
}

func (s *CreateTrailRequest) SetName(v string) *CreateTrailRequest {
	s.Name = &v
	return s
}

func (s *CreateTrailRequest) SetOssBucketName(v string) *CreateTrailRequest {
	s.OssBucketName = &v
	return s
}

func (s *CreateTrailRequest) SetOssKeyPrefix(v string) *CreateTrailRequest {
	s.OssKeyPrefix = &v
	return s
}

func (s *CreateTrailRequest) SetOssWriteRoleArn(v string) *CreateTrailRequest {
	s.OssWriteRoleArn = &v
	return s
}

func (s *CreateTrailRequest) SetSlsProjectArn(v string) *CreateTrailRequest {
	s.SlsProjectArn = &v
	return s
}

func (s *CreateTrailRequest) SetSlsWriteRoleArn(v string) *CreateTrailRequest {
	s.SlsWriteRoleArn = &v
	return s
}

func (s *CreateTrailRequest) SetTrailRegion(v string) *CreateTrailRequest {
	s.TrailRegion = &v
	return s
}

func (s *CreateTrailRequest) Validate() error {
	return dara.Validate(s)
}
