// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageCount(v int64) *DescribeStreamsResponseBody
	GetPageCount() *int64
	SetPageNum(v int64) *DescribeStreamsResponseBody
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeStreamsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeStreamsResponseBody
	GetRequestId() *string
	SetStreams(v []*DescribeStreamsResponseBodyStreams) *DescribeStreamsResponseBody
	GetStreams() []*DescribeStreamsResponseBodyStreams
	SetTotalCount(v int64) *DescribeStreamsResponseBody
	GetTotalCount() *int64
}

type DescribeStreamsResponseBody struct {
	// example:
	//
	// 5
	PageCount *int64 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Streams   []*DescribeStreamsResponseBodyStreams `json:"Streams,omitempty" xml:"Streams,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeStreamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStreamsResponseBody) GetPageCount() *int64 {
	return s.PageCount
}

func (s *DescribeStreamsResponseBody) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeStreamsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeStreamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStreamsResponseBody) GetStreams() []*DescribeStreamsResponseBodyStreams {
	return s.Streams
}

func (s *DescribeStreamsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeStreamsResponseBody) SetPageCount(v int64) *DescribeStreamsResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeStreamsResponseBody) SetPageNum(v int64) *DescribeStreamsResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeStreamsResponseBody) SetPageSize(v int64) *DescribeStreamsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStreamsResponseBody) SetRequestId(v string) *DescribeStreamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStreamsResponseBody) SetStreams(v []*DescribeStreamsResponseBodyStreams) *DescribeStreamsResponseBody {
	s.Streams = v
	return s
}

func (s *DescribeStreamsResponseBody) SetTotalCount(v int64) *DescribeStreamsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeStreamsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeStreamsResponseBodyStreams struct {
	// example:
	//
	// live
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// example:
	//
	// 2018-12-10T17:00:00Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 348*****380-cn-qingdao
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 348*****174-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 720
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 323*****997-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 31000000*****0000002
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// demo.aliyundoc.com
	PlayDomain *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	// example:
	//
	// gb28181
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// example.aliyundoc.com
	PushDomain *string `json:"PushDomain,omitempty" xml:"PushDomain,omitempty"`
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1280
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DescribeStreamsResponseBodyStreams) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamsResponseBodyStreams) GoString() string {
	return s.String()
}

func (s *DescribeStreamsResponseBodyStreams) GetApp() *string {
	return s.App
}

func (s *DescribeStreamsResponseBodyStreams) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeStreamsResponseBodyStreams) GetDeviceId() *string {
	return s.DeviceId
}

func (s *DescribeStreamsResponseBodyStreams) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeStreamsResponseBodyStreams) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeStreamsResponseBodyStreams) GetHeight() *int32 {
	return s.Height
}

func (s *DescribeStreamsResponseBodyStreams) GetId() *string {
	return s.Id
}

func (s *DescribeStreamsResponseBodyStreams) GetName() *string {
	return s.Name
}

func (s *DescribeStreamsResponseBodyStreams) GetPlayDomain() *string {
	return s.PlayDomain
}

func (s *DescribeStreamsResponseBodyStreams) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeStreamsResponseBodyStreams) GetPushDomain() *string {
	return s.PushDomain
}

func (s *DescribeStreamsResponseBodyStreams) GetStatus() *string {
	return s.Status
}

func (s *DescribeStreamsResponseBodyStreams) GetWidth() *int32 {
	return s.Width
}

func (s *DescribeStreamsResponseBodyStreams) SetApp(v string) *DescribeStreamsResponseBodyStreams {
	s.App = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetCreatedTime(v string) *DescribeStreamsResponseBodyStreams {
	s.CreatedTime = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetDeviceId(v string) *DescribeStreamsResponseBodyStreams {
	s.DeviceId = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetEnabled(v bool) *DescribeStreamsResponseBodyStreams {
	s.Enabled = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetGroupId(v string) *DescribeStreamsResponseBodyStreams {
	s.GroupId = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetHeight(v int32) *DescribeStreamsResponseBodyStreams {
	s.Height = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetId(v string) *DescribeStreamsResponseBodyStreams {
	s.Id = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetName(v string) *DescribeStreamsResponseBodyStreams {
	s.Name = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetPlayDomain(v string) *DescribeStreamsResponseBodyStreams {
	s.PlayDomain = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetProtocol(v string) *DescribeStreamsResponseBodyStreams {
	s.Protocol = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetPushDomain(v string) *DescribeStreamsResponseBodyStreams {
	s.PushDomain = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetStatus(v string) *DescribeStreamsResponseBodyStreams {
	s.Status = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetWidth(v int32) *DescribeStreamsResponseBodyStreams {
	s.Width = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) Validate() error {
	return dara.Validate(s)
}
