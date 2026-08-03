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
	// Specifies the read/write type of events that the trail delivers. Valid values:
	//
	// - Write: Write events.
	//
	// - Read: Read events.
	//
	// - All (default): All read and write events.
	//
	// example:
	//
	// Write
	EventRW *string `json:"EventRW,omitempty" xml:"EventRW,omitempty"`
	// Specifies whether the trail is a multi-account trail. Valid values:
	//
	// - true
	//
	// - false (default)
	//
	// To create a trail for an organization, set this parameter to `true`. The trail will collect events from all member accounts in the organization.
	//
	// example:
	//
	// false
	IsOrganizationTrail *bool `json:"IsOrganizationTrail,omitempty" xml:"IsOrganizationTrail,omitempty"`
	// The ARN of the MaxCompute project to which ActionTrail delivers events.
	//
	// > You must specify a destination for the trail by providing at least one of the following parameters: `OssBucketName`, `SlsProjectArn`, or `MaxComputeProjectArn`.
	//
	// > The project name in the ARN must start with `actiontrail_`.
	//
	// example:
	//
	// acs:odps:cn-hangzhou:15127787691****:project/actiontrail_****
	MaxComputeProjectArn *string `json:"MaxComputeProjectArn,omitempty" xml:"MaxComputeProjectArn,omitempty"`
	// The ARN of the RAM role that ActionTrail assumes to deliver events to the MaxCompute project.
	//
	// - If this parameter is not specified, ActionTrail creates a service-linked role to deliver events. For more information, see [ActionTrail service-linked role](https://help.aliyun.com/document_detail/169244.html).
	//
	// - If you specify a role, it must be a RAM role that you created. This role must have a trust policy that allows the ActionTrail service (\\`actiontrail.aliyuncs.com\\`) to assume it. The role\\"s permission policy must grant permissions to write to the specified MaxCompute project. For more information about cross-account delivery, see [Deliver events from multiple Alibaba Cloud accounts to the same account](https://help.aliyun.com/document_detail/207462.html).
	//
	// example:
	//
	// acs:ram::15127787691****:role/aliyunserviceroleforactiontrail
	MaxComputeWriteRoleArn *string `json:"MaxComputeWriteRoleArn,omitempty" xml:"MaxComputeWriteRoleArn,omitempty"`
	// The name of the trail.
	//
	// > - Length: 6 to 36 characters.
	//
	// >
	//
	// > - Characters: Lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// >
	//
	// > - Must start with a lowercase letter.
	//
	// >
	//
	// > - Must be uniquewithin an Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// trail-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the OSS bucket to which ActionTrail delivers events.
	//
	// - Length: 3 to 63 characters.
	//
	// - Characters: Lowercase letters, digits, and hyphens (-).
	//
	// - Must start with a lowercase letter or a digit.
	//
	// > You must specify a destination for the trail by providing at least one of the following parameters: `OssBucketName`, `SlsProjectArn`, or `MaxComputeProjectArn`.
	//
	// example:
	//
	// audit-log
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The prefix for the names of log files that ActionTrail delivers to your OSS bucket.
	//
	// - Length: 6 to 32 characters.
	//
	// - Characters: Letters, digits, hyphens (-), forward slashes (/), and underscores (_).
	//
	// - Must start with a letter.
	//
	// example:
	//
	// at-product-account-audit-B
	OssKeyPrefix *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the RAM role that ActionTrail assumes to deliver events to the OSS bucket.
	//
	// - If you do not specify this parameter, ActionTrail creates a service-linked role to deliver events. For more information, see [ActionTrail service-linked role](https://help.aliyun.com/document_detail/169244.html).
	//
	// - If you specify a role, it must be a RAM role that you created. This role must have a trust policy that allows the ActionTrail service (actiontrail.aliyuncs.com) to assume it. The role\\"s RAM policy must grant permissions to write to the specified OSS bucket. For more information about cross-account delivery, see [Deliver events from multiple Alibaba Cloud accounts to the same account](https://help.aliyun.com/document_detail/207462.html).
	//
	// example:
	//
	// acs:ram::15127787691****:role/aliyunserviceroleforactiontrail
	OssWriteRoleArn *string `json:"OssWriteRoleArn,omitempty" xml:"OssWriteRoleArn,omitempty"`
	// The ARN of the SLS project to which ActionTrail delivers events.
	//
	// > You must specify a destination for the trail by providing at least one of the following parameters: `OssBucketName`, `SlsProjectArn`, or `MaxComputeProjectArn`.
	//
	// example:
	//
	// acs:log:cn-shanghai:151266687691****:project/test-project
	SlsProjectArn *string `json:"SlsProjectArn,omitempty" xml:"SlsProjectArn,omitempty"`
	// The ARN of the RAM role that ActionTrail assumes to deliver events to the SLS project.
	//
	// - If this parameter is not specified, ActionTrail creates a service-linked role to deliver events. For more information, see [ActionTrail service-linked role](https://help.aliyun.com/document_detail/169244.html).
	//
	// - If you specify a role, it must be a RAM role that you created. This role must have a trust policy that allows the ActionTrail service (actiontrail.aliyuncs.com) to assume it. The role\\"s permission policy must grant permissions to write to the specified SLS project. For more information about cross-account delivery, see [Deliver events from multiple Alibaba Cloud accounts to the same account](https://help.aliyun.com/document_detail/207462.html).
	//
	// example:
	//
	// acs:ram::151266687691****:role/aliyunserviceroleforactiontrail
	SlsWriteRoleArn *string `json:"SlsWriteRoleArn,omitempty" xml:"SlsWriteRoleArn,omitempty"`
	// The region in which the trail is created. By default, a trail is created in all regions and this parameter is set to `All`. To create a trail in a specific region, provide the region ID. For more information about regions, call the [DescribeRegions](https://help.aliyun.com/document_detail/213597.html) operation.
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
