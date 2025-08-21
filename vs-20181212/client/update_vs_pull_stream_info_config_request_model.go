// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVsPullStreamInfoConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlways(v string) *UpdateVsPullStreamInfoConfigRequest
	GetAlways() *string
	SetAppName(v string) *UpdateVsPullStreamInfoConfigRequest
	GetAppName() *string
	SetDomainName(v string) *UpdateVsPullStreamInfoConfigRequest
	GetDomainName() *string
	SetEndTime(v string) *UpdateVsPullStreamInfoConfigRequest
	GetEndTime() *string
	SetOwnerId(v int64) *UpdateVsPullStreamInfoConfigRequest
	GetOwnerId() *int64
	SetSourceUrl(v string) *UpdateVsPullStreamInfoConfigRequest
	GetSourceUrl() *string
	SetStartTime(v string) *UpdateVsPullStreamInfoConfigRequest
	GetStartTime() *string
	SetStreamName(v string) *UpdateVsPullStreamInfoConfigRequest
	GetStreamName() *string
}

type UpdateVsPullStreamInfoConfigRequest struct {
	Always *string `json:"Always,omitempty" xml:"Always,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 2018-12-10T18:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// up.xxx.com.cn
	SourceUrl *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	// example:
	//
	// 2021-12-10T10:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxStream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s UpdateVsPullStreamInfoConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVsPullStreamInfoConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateVsPullStreamInfoConfigRequest) GetAlways() *string {
	return s.Always
}

func (s *UpdateVsPullStreamInfoConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateVsPullStreamInfoConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateVsPullStreamInfoConfigRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateVsPullStreamInfoConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateVsPullStreamInfoConfigRequest) GetSourceUrl() *string {
	return s.SourceUrl
}

func (s *UpdateVsPullStreamInfoConfigRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateVsPullStreamInfoConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *UpdateVsPullStreamInfoConfigRequest) SetAlways(v string) *UpdateVsPullStreamInfoConfigRequest {
	s.Always = &v
	return s
}

func (s *UpdateVsPullStreamInfoConfigRequest) SetAppName(v string) *UpdateVsPullStreamInfoConfigRequest {
	s.AppName = &v
	return s
}

func (s *UpdateVsPullStreamInfoConfigRequest) SetDomainName(v string) *UpdateVsPullStreamInfoConfigRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateVsPullStreamInfoConfigRequest) SetEndTime(v string) *UpdateVsPullStreamInfoConfigRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateVsPullStreamInfoConfigRequest) SetOwnerId(v int64) *UpdateVsPullStreamInfoConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateVsPullStreamInfoConfigRequest) SetSourceUrl(v string) *UpdateVsPullStreamInfoConfigRequest {
	s.SourceUrl = &v
	return s
}

func (s *UpdateVsPullStreamInfoConfigRequest) SetStartTime(v string) *UpdateVsPullStreamInfoConfigRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateVsPullStreamInfoConfigRequest) SetStreamName(v string) *UpdateVsPullStreamInfoConfigRequest {
	s.StreamName = &v
	return s
}

func (s *UpdateVsPullStreamInfoConfigRequest) Validate() error {
	return dara.Validate(s)
}
