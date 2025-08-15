// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTrailsResponseBody
	GetRequestId() *string
	SetTrailList(v []*DescribeTrailsResponseBodyTrailList) *DescribeTrailsResponseBody
	GetTrailList() []*DescribeTrailsResponseBodyTrailList
}

type DescribeTrailsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ED8BC689-69DA-42AC-855E-3B06C1271194
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The trails.
	TrailList []*DescribeTrailsResponseBodyTrailList `json:"TrailList,omitempty" xml:"TrailList,omitempty" type:"Repeated"`
}

func (s DescribeTrailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTrailsResponseBody) GetTrailList() []*DescribeTrailsResponseBodyTrailList {
	return s.TrailList
}

func (s *DescribeTrailsResponseBody) SetRequestId(v string) *DescribeTrailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTrailsResponseBody) SetTrailList(v []*DescribeTrailsResponseBodyTrailList) *DescribeTrailsResponseBody {
	s.TrailList = v
	return s
}

func (s *DescribeTrailsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTrailsResponseBodyTrailList struct {
	// The time when the trail was created.
	//
	// example:
	//
	// 2021-03-01T06:27:28Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The read/write type of the events that are delivered. Valid values:
	//
	// 	- Write: write events. This is the default value.
	//
	// 	- Read: read events.
	//
	// 	- All: read and write events.
	//
	// example:
	//
	// All
	EventRW *string `json:"EventRW,omitempty" xml:"EventRW,omitempty"`
	// The home region of the trail.
	//
	// example:
	//
	// cn-hangzhou
	HomeRegion *string `json:"HomeRegion,omitempty" xml:"HomeRegion,omitempty"`
	// Indicates whether the trail is a multi-account trail. Valid values:
	//
	// 	- false (default)
	//
	// 	- true
	//
	// example:
	//
	// false
	IsOrganizationTrail *bool `json:"IsOrganizationTrail,omitempty" xml:"IsOrganizationTrail,omitempty"`
	// The ARN of the MaxCompute project.
	//
	// example:
	//
	// acs:odps:cn-hangzhou:141266687691****:project/actiontrail_****
	MaxComputeProjectArn *string `json:"MaxComputeProjectArn,omitempty" xml:"MaxComputeProjectArn,omitempty"`
	// The ARN of the role that is assumed by ActionTrail to deliver events to the MaxCompute project.
	//
	// example:
	//
	// acs:ram::141266687691****:role/aliyunserviceroleforactiontrail
	MaxComputeWriteRoleArn *string `json:"MaxComputeWriteRoleArn,omitempty" xml:"MaxComputeWriteRoleArn,omitempty"`
	// The name of the trail.
	//
	// example:
	//
	// test-4
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource directory.
	//
	// >  This parameter is returned only when the trail is a multi-account trail.
	//
	// example:
	//
	// rd-EV****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// The region where the OSS bucket resides.
	//
	// example:
	//
	// oss-cn-hangzhou
	OssBucketLocation *string `json:"OssBucketLocation,omitempty" xml:"OssBucketLocation,omitempty"`
	// The name of the OSS bucket to which events are delivered.
	//
	// example:
	//
	// secloud
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The prefix of the files that are stored in the Object Storage Service (OSS) bucket.
	//
	// example:
	//
	// trail1
	OssKeyPrefix *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the RAM role that is assumed by ActionTrail to deliver events to the OSS bucket.
	//
	// example:
	//
	// acs:ram::***:role/aliyunserviceroleforactiontrail
	OssWriteRoleArn *string `json:"OssWriteRoleArn,omitempty" xml:"OssWriteRoleArn,omitempty"`
	// The region where the trail resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ARN of the Log Service project to which events are delivered.
	//
	// example:
	//
	// acs:log:cn-qingdao:159498693826****:project/zhengze-audit-log
	SlsProjectArn *string `json:"SlsProjectArn,omitempty" xml:"SlsProjectArn,omitempty"`
	// The ARN of the RAM role that is assumed by ActionTrail to deliver events to the Log Service project.
	//
	// example:
	//
	// acs:ram::159498693826****:role/aliyunserviceroleforactiontrail
	SlsWriteRoleArn *string `json:"SlsWriteRoleArn,omitempty" xml:"SlsWriteRoleArn,omitempty"`
	// The time when the trail was last enabled.
	//
	// example:
	//
	// 2021-04-06T02:08:38Z
	StartLoggingTime *string `json:"StartLoggingTime,omitempty" xml:"StartLoggingTime,omitempty"`
	// The status of the trail. Valid values:
	//
	// 	- Disable: disabled.
	//
	// 	- Enable: enabled.
	//
	// 	- Fresh: The trail is created but is not enabled.
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the trail was last disabled.
	//
	// example:
	//
	// 2021-04-06T02:09:04Z
	StopLoggingTime *string `json:"StopLoggingTime,omitempty" xml:"StopLoggingTime,omitempty"`
	// The ARN of the trail.
	//
	// example:
	//
	// acs:actiontrail:cn-hangzhou:159498693826****:trail/test-delivery-other
	TrailArn *string `json:"TrailArn,omitempty" xml:"TrailArn,omitempty"`
	// The region of the trail.
	//
	// example:
	//
	// All
	TrailRegion *string `json:"TrailRegion,omitempty" xml:"TrailRegion,omitempty"`
	// The time when the configurations of the trail were last updated.
	//
	// example:
	//
	// 2021-04-06T02:16:24Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeTrailsResponseBodyTrailList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrailsResponseBodyTrailList) GoString() string {
	return s.String()
}

func (s *DescribeTrailsResponseBodyTrailList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeTrailsResponseBodyTrailList) GetEventRW() *string {
	return s.EventRW
}

func (s *DescribeTrailsResponseBodyTrailList) GetHomeRegion() *string {
	return s.HomeRegion
}

func (s *DescribeTrailsResponseBodyTrailList) GetIsOrganizationTrail() *bool {
	return s.IsOrganizationTrail
}

func (s *DescribeTrailsResponseBodyTrailList) GetMaxComputeProjectArn() *string {
	return s.MaxComputeProjectArn
}

func (s *DescribeTrailsResponseBodyTrailList) GetMaxComputeWriteRoleArn() *string {
	return s.MaxComputeWriteRoleArn
}

func (s *DescribeTrailsResponseBodyTrailList) GetName() *string {
	return s.Name
}

func (s *DescribeTrailsResponseBodyTrailList) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DescribeTrailsResponseBodyTrailList) GetOssBucketLocation() *string {
	return s.OssBucketLocation
}

func (s *DescribeTrailsResponseBodyTrailList) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *DescribeTrailsResponseBodyTrailList) GetOssKeyPrefix() *string {
	return s.OssKeyPrefix
}

func (s *DescribeTrailsResponseBodyTrailList) GetOssWriteRoleArn() *string {
	return s.OssWriteRoleArn
}

func (s *DescribeTrailsResponseBodyTrailList) GetRegion() *string {
	return s.Region
}

func (s *DescribeTrailsResponseBodyTrailList) GetSlsProjectArn() *string {
	return s.SlsProjectArn
}

func (s *DescribeTrailsResponseBodyTrailList) GetSlsWriteRoleArn() *string {
	return s.SlsWriteRoleArn
}

func (s *DescribeTrailsResponseBodyTrailList) GetStartLoggingTime() *string {
	return s.StartLoggingTime
}

func (s *DescribeTrailsResponseBodyTrailList) GetStatus() *string {
	return s.Status
}

func (s *DescribeTrailsResponseBodyTrailList) GetStopLoggingTime() *string {
	return s.StopLoggingTime
}

func (s *DescribeTrailsResponseBodyTrailList) GetTrailArn() *string {
	return s.TrailArn
}

func (s *DescribeTrailsResponseBodyTrailList) GetTrailRegion() *string {
	return s.TrailRegion
}

func (s *DescribeTrailsResponseBodyTrailList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeTrailsResponseBodyTrailList) SetCreateTime(v string) *DescribeTrailsResponseBodyTrailList {
	s.CreateTime = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetEventRW(v string) *DescribeTrailsResponseBodyTrailList {
	s.EventRW = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetHomeRegion(v string) *DescribeTrailsResponseBodyTrailList {
	s.HomeRegion = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetIsOrganizationTrail(v bool) *DescribeTrailsResponseBodyTrailList {
	s.IsOrganizationTrail = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetMaxComputeProjectArn(v string) *DescribeTrailsResponseBodyTrailList {
	s.MaxComputeProjectArn = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetMaxComputeWriteRoleArn(v string) *DescribeTrailsResponseBodyTrailList {
	s.MaxComputeWriteRoleArn = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetName(v string) *DescribeTrailsResponseBodyTrailList {
	s.Name = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetOrganizationId(v string) *DescribeTrailsResponseBodyTrailList {
	s.OrganizationId = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetOssBucketLocation(v string) *DescribeTrailsResponseBodyTrailList {
	s.OssBucketLocation = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetOssBucketName(v string) *DescribeTrailsResponseBodyTrailList {
	s.OssBucketName = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetOssKeyPrefix(v string) *DescribeTrailsResponseBodyTrailList {
	s.OssKeyPrefix = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetOssWriteRoleArn(v string) *DescribeTrailsResponseBodyTrailList {
	s.OssWriteRoleArn = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetRegion(v string) *DescribeTrailsResponseBodyTrailList {
	s.Region = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetSlsProjectArn(v string) *DescribeTrailsResponseBodyTrailList {
	s.SlsProjectArn = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetSlsWriteRoleArn(v string) *DescribeTrailsResponseBodyTrailList {
	s.SlsWriteRoleArn = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetStartLoggingTime(v string) *DescribeTrailsResponseBodyTrailList {
	s.StartLoggingTime = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetStatus(v string) *DescribeTrailsResponseBodyTrailList {
	s.Status = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetStopLoggingTime(v string) *DescribeTrailsResponseBodyTrailList {
	s.StopLoggingTime = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetTrailArn(v string) *DescribeTrailsResponseBodyTrailList {
	s.TrailArn = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetTrailRegion(v string) *DescribeTrailsResponseBodyTrailList {
	s.TrailRegion = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetUpdateTime(v string) *DescribeTrailsResponseBodyTrailList {
	s.UpdateTime = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) Validate() error {
	return dara.Validate(s)
}
