// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventRW(v string) *CreateTrailResponseBody
	GetEventRW() *string
	SetHomeRegion(v string) *CreateTrailResponseBody
	GetHomeRegion() *string
	SetMaxComputeProjectArn(v string) *CreateTrailResponseBody
	GetMaxComputeProjectArn() *string
	SetMaxComputeWriteRoleArn(v string) *CreateTrailResponseBody
	GetMaxComputeWriteRoleArn() *string
	SetName(v string) *CreateTrailResponseBody
	GetName() *string
	SetOssBucketName(v string) *CreateTrailResponseBody
	GetOssBucketName() *string
	SetOssKeyPrefix(v string) *CreateTrailResponseBody
	GetOssKeyPrefix() *string
	SetOssWriteRoleArn(v string) *CreateTrailResponseBody
	GetOssWriteRoleArn() *string
	SetRequestId(v string) *CreateTrailResponseBody
	GetRequestId() *string
	SetSlsProjectArn(v string) *CreateTrailResponseBody
	GetSlsProjectArn() *string
	SetSlsWriteRoleArn(v string) *CreateTrailResponseBody
	GetSlsWriteRoleArn() *string
	SetTrailRegion(v string) *CreateTrailResponseBody
	GetTrailRegion() *string
}

type CreateTrailResponseBody struct {
	// The read/write type of events that the trail delivers.
	//
	// example:
	//
	// Write
	EventRW *string `json:"EventRW,omitempty" xml:"EventRW,omitempty"`
	// The home region of the trail.
	//
	// example:
	//
	// cn-hangzhou
	HomeRegion *string `json:"HomeRegion,omitempty" xml:"HomeRegion,omitempty"`
	// The ARN of the MaxCompute project to which the trail delivers events.
	//
	// example:
	//
	// acs:odps:cn-hangzhou:151266687691****:project/actiontrail_****
	MaxComputeProjectArn *string `json:"MaxComputeProjectArn,omitempty" xml:"MaxComputeProjectArn,omitempty"`
	// The ARN of the RAM role that ActionTrail assumes to deliver events to the MaxCompute project.
	//
	// example:
	//
	// acs:ram::151266687691****:role/aliyunserviceroleforactiontrail
	MaxComputeWriteRoleArn *string `json:"MaxComputeWriteRoleArn,omitempty" xml:"MaxComputeWriteRoleArn,omitempty"`
	// The name of the trail.
	//
	// example:
	//
	// trail-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the destination OSS bucket.
	//
	// example:
	//
	// audit-log
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The prefix for the names of log files in the OSS bucket.
	//
	// example:
	//
	// at-product-account-audit-B
	OssKeyPrefix *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	// The ARN of the RAM role that ActionTrail assumes to deliver events to the OSS bucket.
	//
	// example:
	//
	// acs:ram::151266687691****:role/aliyunserviceroleforactiontrail
	OssWriteRoleArn *string `json:"OssWriteRoleArn,omitempty" xml:"OssWriteRoleArn,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 442DDADF-DA58-4029-8E8B-82C73E9A7A70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ARN of the SLS project to which the trail delivers events.
	//
	// example:
	//
	// acs:log:cn-hangzhou:151266687691****:project/test-project
	SlsProjectArn *string `json:"SlsProjectArn,omitempty" xml:"SlsProjectArn,omitempty"`
	// The ARN of the RAM role that ActionTrail assumes to deliver events to the SLS project.
	//
	// example:
	//
	// acs:ram::151266687691****:role/aliyunserviceroleforactiontrail
	SlsWriteRoleArn *string `json:"SlsWriteRoleArn,omitempty" xml:"SlsWriteRoleArn,omitempty"`
	// The region in which the trail is created. A value of `All` indicates that the trail processes events from all regions.
	//
	// example:
	//
	// All
	TrailRegion *string `json:"TrailRegion,omitempty" xml:"TrailRegion,omitempty"`
}

func (s CreateTrailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTrailResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrailResponseBody) GetEventRW() *string {
	return s.EventRW
}

func (s *CreateTrailResponseBody) GetHomeRegion() *string {
	return s.HomeRegion
}

func (s *CreateTrailResponseBody) GetMaxComputeProjectArn() *string {
	return s.MaxComputeProjectArn
}

func (s *CreateTrailResponseBody) GetMaxComputeWriteRoleArn() *string {
	return s.MaxComputeWriteRoleArn
}

func (s *CreateTrailResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateTrailResponseBody) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *CreateTrailResponseBody) GetOssKeyPrefix() *string {
	return s.OssKeyPrefix
}

func (s *CreateTrailResponseBody) GetOssWriteRoleArn() *string {
	return s.OssWriteRoleArn
}

func (s *CreateTrailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTrailResponseBody) GetSlsProjectArn() *string {
	return s.SlsProjectArn
}

func (s *CreateTrailResponseBody) GetSlsWriteRoleArn() *string {
	return s.SlsWriteRoleArn
}

func (s *CreateTrailResponseBody) GetTrailRegion() *string {
	return s.TrailRegion
}

func (s *CreateTrailResponseBody) SetEventRW(v string) *CreateTrailResponseBody {
	s.EventRW = &v
	return s
}

func (s *CreateTrailResponseBody) SetHomeRegion(v string) *CreateTrailResponseBody {
	s.HomeRegion = &v
	return s
}

func (s *CreateTrailResponseBody) SetMaxComputeProjectArn(v string) *CreateTrailResponseBody {
	s.MaxComputeProjectArn = &v
	return s
}

func (s *CreateTrailResponseBody) SetMaxComputeWriteRoleArn(v string) *CreateTrailResponseBody {
	s.MaxComputeWriteRoleArn = &v
	return s
}

func (s *CreateTrailResponseBody) SetName(v string) *CreateTrailResponseBody {
	s.Name = &v
	return s
}

func (s *CreateTrailResponseBody) SetOssBucketName(v string) *CreateTrailResponseBody {
	s.OssBucketName = &v
	return s
}

func (s *CreateTrailResponseBody) SetOssKeyPrefix(v string) *CreateTrailResponseBody {
	s.OssKeyPrefix = &v
	return s
}

func (s *CreateTrailResponseBody) SetOssWriteRoleArn(v string) *CreateTrailResponseBody {
	s.OssWriteRoleArn = &v
	return s
}

func (s *CreateTrailResponseBody) SetRequestId(v string) *CreateTrailResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrailResponseBody) SetSlsProjectArn(v string) *CreateTrailResponseBody {
	s.SlsProjectArn = &v
	return s
}

func (s *CreateTrailResponseBody) SetSlsWriteRoleArn(v string) *CreateTrailResponseBody {
	s.SlsWriteRoleArn = &v
	return s
}

func (s *CreateTrailResponseBody) SetTrailRegion(v string) *CreateTrailResponseBody {
	s.TrailRegion = &v
	return s
}

func (s *CreateTrailResponseBody) Validate() error {
	return dara.Validate(s)
}
