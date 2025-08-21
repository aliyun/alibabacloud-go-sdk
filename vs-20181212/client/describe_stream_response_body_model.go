// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *DescribeStreamResponseBody
	GetApp() *string
	SetCreatedTime(v string) *DescribeStreamResponseBody
	GetCreatedTime() *string
	SetDeviceId(v string) *DescribeStreamResponseBody
	GetDeviceId() *string
	SetEnabled(v bool) *DescribeStreamResponseBody
	GetEnabled() *bool
	SetGroupId(v string) *DescribeStreamResponseBody
	GetGroupId() *string
	SetHeight(v int32) *DescribeStreamResponseBody
	GetHeight() *int32
	SetId(v string) *DescribeStreamResponseBody
	GetId() *string
	SetName(v string) *DescribeStreamResponseBody
	GetName() *string
	SetPlayDomain(v string) *DescribeStreamResponseBody
	GetPlayDomain() *string
	SetProtocol(v string) *DescribeStreamResponseBody
	GetProtocol() *string
	SetPushDomain(v string) *DescribeStreamResponseBody
	GetPushDomain() *string
	SetRequestId(v string) *DescribeStreamResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeStreamResponseBody
	GetStatus() *string
	SetWidth(v int32) *DescribeStreamResponseBody
	GetWidth() *int32
}

type DescribeStreamResponseBody struct {
	// example:
	//
	// live
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// example:
	//
	// 2019-02-28T17:00:17Z
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
	// example.aliyundoc.com
	PlayDomain *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	// example:
	//
	// gb28181
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// demo.aliyundoc.com
	PushDomain *string `json:"PushDomain,omitempty" xml:"PushDomain,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1280
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DescribeStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStreamResponseBody) GetApp() *string {
	return s.App
}

func (s *DescribeStreamResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeStreamResponseBody) GetDeviceId() *string {
	return s.DeviceId
}

func (s *DescribeStreamResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeStreamResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeStreamResponseBody) GetHeight() *int32 {
	return s.Height
}

func (s *DescribeStreamResponseBody) GetId() *string {
	return s.Id
}

func (s *DescribeStreamResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeStreamResponseBody) GetPlayDomain() *string {
	return s.PlayDomain
}

func (s *DescribeStreamResponseBody) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeStreamResponseBody) GetPushDomain() *string {
	return s.PushDomain
}

func (s *DescribeStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStreamResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeStreamResponseBody) GetWidth() *int32 {
	return s.Width
}

func (s *DescribeStreamResponseBody) SetApp(v string) *DescribeStreamResponseBody {
	s.App = &v
	return s
}

func (s *DescribeStreamResponseBody) SetCreatedTime(v string) *DescribeStreamResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeStreamResponseBody) SetDeviceId(v string) *DescribeStreamResponseBody {
	s.DeviceId = &v
	return s
}

func (s *DescribeStreamResponseBody) SetEnabled(v bool) *DescribeStreamResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeStreamResponseBody) SetGroupId(v string) *DescribeStreamResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeStreamResponseBody) SetHeight(v int32) *DescribeStreamResponseBody {
	s.Height = &v
	return s
}

func (s *DescribeStreamResponseBody) SetId(v string) *DescribeStreamResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeStreamResponseBody) SetName(v string) *DescribeStreamResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeStreamResponseBody) SetPlayDomain(v string) *DescribeStreamResponseBody {
	s.PlayDomain = &v
	return s
}

func (s *DescribeStreamResponseBody) SetProtocol(v string) *DescribeStreamResponseBody {
	s.Protocol = &v
	return s
}

func (s *DescribeStreamResponseBody) SetPushDomain(v string) *DescribeStreamResponseBody {
	s.PushDomain = &v
	return s
}

func (s *DescribeStreamResponseBody) SetRequestId(v string) *DescribeStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStreamResponseBody) SetStatus(v string) *DescribeStreamResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeStreamResponseBody) SetWidth(v int32) *DescribeStreamResponseBody {
	s.Width = &v
	return s
}

func (s *DescribeStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
