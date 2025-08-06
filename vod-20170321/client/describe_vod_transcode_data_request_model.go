// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodTranscodeDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeVodTranscodeDataRequest
	GetAppId() *string
	SetEndTime(v string) *DescribeVodTranscodeDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVodTranscodeDataRequest
	GetInterval() *string
	SetOwnerId(v int64) *DescribeVodTranscodeDataRequest
	GetOwnerId() *int64
	SetRegion(v string) *DescribeVodTranscodeDataRequest
	GetRegion() *string
	SetSpecification(v string) *DescribeVodTranscodeDataRequest
	GetSpecification() *string
	SetStartTime(v string) *DescribeVodTranscodeDataRequest
	GetStartTime() *string
	SetStorage(v string) *DescribeVodTranscodeDataRequest
	GetStorage() *string
}

type DescribeVodTranscodeDataRequest struct {
	// The ID of the application. You can specify this parameter to query the transcoding statistics of a specific application. By default, the transcoding statistics of all applications is returned. You can obtain the application ID from the `AppId` parameter in the response to the [CreateAppInfo](~~CreateAppInfo~~) operation.
	//
	// example:
	//
	// app-1000001
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-02-01T15:59:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The interval at which you want to query data. Valid values:
	//
	// 	- **day**: days
	//
	// 	- **hour**: hours
	//
	// example:
	//
	// day
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region in which you want to query data. If you leave this parameter empty, data in all regions is returned. Separate multiple regions with commas (,). Valid values:
	//
	// 	- **cn-shanghai**: China (Shanghai)
	//
	// 	- **cn-beijing**: China (Beijing)
	//
	// 	- **eu-central-1**: Germany (Frankfurt)
	//
	// 	- **ap-southeast-1**: Singapore
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The transcoding specification. If you leave this parameter empty, data of all transcoding specifications is returned. Separate multiple transcoding specifications with commas (,). Valid values:
	//
	// 	- **Audio**: audio transcoding
	//
	// 	- **Segmentation**: container format conversion
	//
	// 	- **H264.LD**, **H264.SD**, **H264.HD**, **H264.2K**, **H264.4K**, and more
	//
	// example:
	//
	// Audio
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-02-01T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the Object Storage Service (OSS) bucket. If you leave this parameter empty, data of all buckets is returned. Separate multiple bucket names with commas (,).
	//
	// example:
	//
	// bucket01
	Storage *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
}

func (s DescribeVodTranscodeDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTranscodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodTranscodeDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeVodTranscodeDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodTranscodeDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodTranscodeDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodTranscodeDataRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeVodTranscodeDataRequest) GetSpecification() *string {
	return s.Specification
}

func (s *DescribeVodTranscodeDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodTranscodeDataRequest) GetStorage() *string {
	return s.Storage
}

func (s *DescribeVodTranscodeDataRequest) SetAppId(v string) *DescribeVodTranscodeDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeVodTranscodeDataRequest) SetEndTime(v string) *DescribeVodTranscodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodTranscodeDataRequest) SetInterval(v string) *DescribeVodTranscodeDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodTranscodeDataRequest) SetOwnerId(v int64) *DescribeVodTranscodeDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodTranscodeDataRequest) SetRegion(v string) *DescribeVodTranscodeDataRequest {
	s.Region = &v
	return s
}

func (s *DescribeVodTranscodeDataRequest) SetSpecification(v string) *DescribeVodTranscodeDataRequest {
	s.Specification = &v
	return s
}

func (s *DescribeVodTranscodeDataRequest) SetStartTime(v string) *DescribeVodTranscodeDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodTranscodeDataRequest) SetStorage(v string) *DescribeVodTranscodeDataRequest {
	s.Storage = &v
	return s
}

func (s *DescribeVodTranscodeDataRequest) Validate() error {
	return dara.Validate(s)
}
