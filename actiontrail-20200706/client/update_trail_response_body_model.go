// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventRW(v string) *UpdateTrailResponseBody
	GetEventRW() *string
	SetHomeRegion(v string) *UpdateTrailResponseBody
	GetHomeRegion() *string
	SetMaxComputeProjectArn(v string) *UpdateTrailResponseBody
	GetMaxComputeProjectArn() *string
	SetMaxComputeWriteRoleArn(v string) *UpdateTrailResponseBody
	GetMaxComputeWriteRoleArn() *string
	SetName(v string) *UpdateTrailResponseBody
	GetName() *string
	SetOssBucketName(v string) *UpdateTrailResponseBody
	GetOssBucketName() *string
	SetOssKeyPrefix(v string) *UpdateTrailResponseBody
	GetOssKeyPrefix() *string
	SetOssWriteRoleArn(v string) *UpdateTrailResponseBody
	GetOssWriteRoleArn() *string
	SetRequestId(v string) *UpdateTrailResponseBody
	GetRequestId() *string
	SetSlsProjectArn(v string) *UpdateTrailResponseBody
	GetSlsProjectArn() *string
	SetSlsWriteRoleArn(v string) *UpdateTrailResponseBody
	GetSlsWriteRoleArn() *string
	SetTrailRegion(v string) *UpdateTrailResponseBody
	GetTrailRegion() *string
}

type UpdateTrailResponseBody struct {
	// The read/write type of the events to be delivered.
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
	// ARN of the Big Data Compute Service project for tracking delivery.
	//
	// example:
	//
	// acs:odps:cn-hangzhou:151266687691****:project/actiontrail_****
	MaxComputeProjectArn *string `json:"MaxComputeProjectArn,omitempty" xml:"MaxComputeProjectArn,omitempty"`
	// The ARN of the role that Operation Audit assumes when delivering operation events to the Big Data Compute Service project.
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
	// The name of the OSS bucket.
	//
	// example:
	//
	// audit-log
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The prefix of the log files to be stored in the destination OSS bucket.
	//
	// example:
	//
	// at-product-account-audit-B
	OssKeyPrefix *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	// The ARN of the RAM role that is assumed by ActionTrail to deliver events to the OSS bucket.
	//
	// example:
	//
	// acs:ram::***:role/aliyunserviceroleforactiontrail
	OssWriteRoleArn *string `json:"OssWriteRoleArn,omitempty" xml:"OssWriteRoleArn,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2599A180-5236-44D8-9490-50B6F4F8BA35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ARN of the Log Service project to which events are to be delivered.
	//
	// example:
	//
	// acs:log:cn-hangzhou:151266687691****:project/test-project
	SlsProjectArn *string `json:"SlsProjectArn,omitempty" xml:"SlsProjectArn,omitempty"`
	// The ARN of the RAM role that is assumed by ActionTrail is to deliver events to the Log Service project.
	//
	// example:
	//
	// acs:ram::***:role/aliyunserviceroleforactiontrail
	SlsWriteRoleArn *string `json:"SlsWriteRoleArn,omitempty" xml:"SlsWriteRoleArn,omitempty"`
	// The one or more regions from which the trail delivers events.
	//
	// example:
	//
	// All
	TrailRegion *string `json:"TrailRegion,omitempty" xml:"TrailRegion,omitempty"`
}

func (s UpdateTrailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrailResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTrailResponseBody) GetEventRW() *string {
	return s.EventRW
}

func (s *UpdateTrailResponseBody) GetHomeRegion() *string {
	return s.HomeRegion
}

func (s *UpdateTrailResponseBody) GetMaxComputeProjectArn() *string {
	return s.MaxComputeProjectArn
}

func (s *UpdateTrailResponseBody) GetMaxComputeWriteRoleArn() *string {
	return s.MaxComputeWriteRoleArn
}

func (s *UpdateTrailResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateTrailResponseBody) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *UpdateTrailResponseBody) GetOssKeyPrefix() *string {
	return s.OssKeyPrefix
}

func (s *UpdateTrailResponseBody) GetOssWriteRoleArn() *string {
	return s.OssWriteRoleArn
}

func (s *UpdateTrailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTrailResponseBody) GetSlsProjectArn() *string {
	return s.SlsProjectArn
}

func (s *UpdateTrailResponseBody) GetSlsWriteRoleArn() *string {
	return s.SlsWriteRoleArn
}

func (s *UpdateTrailResponseBody) GetTrailRegion() *string {
	return s.TrailRegion
}

func (s *UpdateTrailResponseBody) SetEventRW(v string) *UpdateTrailResponseBody {
	s.EventRW = &v
	return s
}

func (s *UpdateTrailResponseBody) SetHomeRegion(v string) *UpdateTrailResponseBody {
	s.HomeRegion = &v
	return s
}

func (s *UpdateTrailResponseBody) SetMaxComputeProjectArn(v string) *UpdateTrailResponseBody {
	s.MaxComputeProjectArn = &v
	return s
}

func (s *UpdateTrailResponseBody) SetMaxComputeWriteRoleArn(v string) *UpdateTrailResponseBody {
	s.MaxComputeWriteRoleArn = &v
	return s
}

func (s *UpdateTrailResponseBody) SetName(v string) *UpdateTrailResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateTrailResponseBody) SetOssBucketName(v string) *UpdateTrailResponseBody {
	s.OssBucketName = &v
	return s
}

func (s *UpdateTrailResponseBody) SetOssKeyPrefix(v string) *UpdateTrailResponseBody {
	s.OssKeyPrefix = &v
	return s
}

func (s *UpdateTrailResponseBody) SetOssWriteRoleArn(v string) *UpdateTrailResponseBody {
	s.OssWriteRoleArn = &v
	return s
}

func (s *UpdateTrailResponseBody) SetRequestId(v string) *UpdateTrailResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTrailResponseBody) SetSlsProjectArn(v string) *UpdateTrailResponseBody {
	s.SlsProjectArn = &v
	return s
}

func (s *UpdateTrailResponseBody) SetSlsWriteRoleArn(v string) *UpdateTrailResponseBody {
	s.SlsWriteRoleArn = &v
	return s
}

func (s *UpdateTrailResponseBody) SetTrailRegion(v string) *UpdateTrailResponseBody {
	s.TrailRegion = &v
	return s
}

func (s *UpdateTrailResponseBody) Validate() error {
	return dara.Validate(s)
}
