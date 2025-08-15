// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventRW(v string) *UpdateTrailRequest
	GetEventRW() *string
	SetMaxComputeProjectArn(v string) *UpdateTrailRequest
	GetMaxComputeProjectArn() *string
	SetMaxComputeWriteRoleArn(v string) *UpdateTrailRequest
	GetMaxComputeWriteRoleArn() *string
	SetName(v string) *UpdateTrailRequest
	GetName() *string
	SetOssBucketName(v string) *UpdateTrailRequest
	GetOssBucketName() *string
	SetOssKeyPrefix(v string) *UpdateTrailRequest
	GetOssKeyPrefix() *string
	SetOssWriteRoleArn(v string) *UpdateTrailRequest
	GetOssWriteRoleArn() *string
	SetSlsProjectArn(v string) *UpdateTrailRequest
	GetSlsProjectArn() *string
	SetSlsWriteRoleArn(v string) *UpdateTrailRequest
	GetSlsWriteRoleArn() *string
	SetTrailRegion(v string) *UpdateTrailRequest
	GetTrailRegion() *string
}

type UpdateTrailRequest struct {
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
	// All
	EventRW *string `json:"EventRW,omitempty" xml:"EventRW,omitempty"`
	// The ARN of the MaxCompute project to which you want to deliver events.
	//
	// >  The name of the MaxCompute project must be prefixed with actiontrail_.
	//
	// example:
	//
	// acs:odps:cn-hangzhou:、151277687691****:project/actiontrail_****
	MaxComputeProjectArn *string `json:"MaxComputeProjectArn,omitempty" xml:"MaxComputeProjectArn,omitempty"`
	// The ARN of the role that is assumed by ActionTrail to deliver events to the destination Simple Log Service project.
	//
	// 	- If you do not specify this parameter, ActionTrail creates a service-linked role to create the required resources. For more information, see [Manage the service-linked role](https://help.aliyun.com/document_detail/169244.html).
	//
	// 	- If you specify this parameter and deliver events to the current account, you must grant the RAM role the permissions on the service-linked role for ActionTrail. If you want to deliver events to other accounts, you must attach a system policy to the RAM role. For more information about how to deliver events across Alibaba Cloud accounts, see [Deliver events across Alibaba Cloud accounts](https://help.aliyun.com/document_detail/207462.html).
	//
	// example:
	//
	// acs:ram::151277687691****:role/aliyunserviceroleforactiontrail
	MaxComputeWriteRoleArn *string `json:"MaxComputeWriteRoleArn,omitempty" xml:"MaxComputeWriteRoleArn,omitempty"`
	// The name of the trail whose configurations you want to update.
	//
	// The name must be 6 to 36 characters in length and can contain lowercase letters, digits, hyphens (-), and underscores (_). It must start with a lowercase letter.
	//
	// >  The name must be unique within an Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// trail-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the Object Storage Service (OSS) bucket to which you want to deliver events.
	//
	// The name must be 3 to 63 characters in length. The name must start with a lowercase letter or a digit and can contain lowercase letters, digits, and hyphens (-).
	//
	// >  Make sure that the bucket exists before you update the configuration of the trail.
	//
	// example:
	//
	// audit-log
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The prefix of the files that are stored in the OSS bucket.
	//
	// The prefix must be 6 to 32 characters in length. The prefix must start with a letter and can contain letters, digits, hyphens (-), forward slashes (/), and underscores (_).
	//
	// example:
	//
	// at-product-account-audit-B
	OssKeyPrefix *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the RAM role that is assumed by ActionTrail to deliver events to the OSS bucket.
	//
	// 	- If you do not specify this parameter, ActionTrail creates a service-linked role to create the required resources. For more information, see [Manage the service-linked role](https://help.aliyun.com/document_detail/169244.html).
	//
	// 	- If you specify this parameter, you must grant the permissions of the service-linked role that is assumed by ActionTrail to the RAM role before you can deliver events to your Alibaba Cloud account. If you need to deliver events to other Alibaba Cloud accounts, you must attach the permission policy that is used to grant permissions related to event delivery to the RAM role. For more information about how to deliver events across Alibaba Cloud accounts, see [Deliver events across Alibaba Cloud accounts](https://help.aliyun.com/document_detail/207462.html).
	//
	// example:
	//
	// acs:ram::***:role/aliyunserviceroleforactiontrail
	OssWriteRoleArn *string `json:"OssWriteRoleArn,omitempty" xml:"OssWriteRoleArn,omitempty"`
	// The ARN of the Log Service project to which you want to deliver events.
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
	// The region of the trail.
	//
	// 	- The default value is All, which indicates that the trail delivers events from all regions.
	//
	// You can also specify specific regions. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/213597.html) operation to query all the supported regions.
	//
	// example:
	//
	// All
	TrailRegion *string `json:"TrailRegion,omitempty" xml:"TrailRegion,omitempty"`
}

func (s UpdateTrailRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrailRequest) GoString() string {
	return s.String()
}

func (s *UpdateTrailRequest) GetEventRW() *string {
	return s.EventRW
}

func (s *UpdateTrailRequest) GetMaxComputeProjectArn() *string {
	return s.MaxComputeProjectArn
}

func (s *UpdateTrailRequest) GetMaxComputeWriteRoleArn() *string {
	return s.MaxComputeWriteRoleArn
}

func (s *UpdateTrailRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTrailRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *UpdateTrailRequest) GetOssKeyPrefix() *string {
	return s.OssKeyPrefix
}

func (s *UpdateTrailRequest) GetOssWriteRoleArn() *string {
	return s.OssWriteRoleArn
}

func (s *UpdateTrailRequest) GetSlsProjectArn() *string {
	return s.SlsProjectArn
}

func (s *UpdateTrailRequest) GetSlsWriteRoleArn() *string {
	return s.SlsWriteRoleArn
}

func (s *UpdateTrailRequest) GetTrailRegion() *string {
	return s.TrailRegion
}

func (s *UpdateTrailRequest) SetEventRW(v string) *UpdateTrailRequest {
	s.EventRW = &v
	return s
}

func (s *UpdateTrailRequest) SetMaxComputeProjectArn(v string) *UpdateTrailRequest {
	s.MaxComputeProjectArn = &v
	return s
}

func (s *UpdateTrailRequest) SetMaxComputeWriteRoleArn(v string) *UpdateTrailRequest {
	s.MaxComputeWriteRoleArn = &v
	return s
}

func (s *UpdateTrailRequest) SetName(v string) *UpdateTrailRequest {
	s.Name = &v
	return s
}

func (s *UpdateTrailRequest) SetOssBucketName(v string) *UpdateTrailRequest {
	s.OssBucketName = &v
	return s
}

func (s *UpdateTrailRequest) SetOssKeyPrefix(v string) *UpdateTrailRequest {
	s.OssKeyPrefix = &v
	return s
}

func (s *UpdateTrailRequest) SetOssWriteRoleArn(v string) *UpdateTrailRequest {
	s.OssWriteRoleArn = &v
	return s
}

func (s *UpdateTrailRequest) SetSlsProjectArn(v string) *UpdateTrailRequest {
	s.SlsProjectArn = &v
	return s
}

func (s *UpdateTrailRequest) SetSlsWriteRoleArn(v string) *UpdateTrailRequest {
	s.SlsWriteRoleArn = &v
	return s
}

func (s *UpdateTrailRequest) SetTrailRegion(v string) *UpdateTrailRequest {
	s.TrailRegion = &v
	return s
}

func (s *UpdateTrailRequest) Validate() error {
	return dara.Validate(s)
}
