// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVsPullStreamInfoConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlways(v string) *AddVsPullStreamInfoConfigRequest
	GetAlways() *string
	SetAppName(v string) *AddVsPullStreamInfoConfigRequest
	GetAppName() *string
	SetDomainName(v string) *AddVsPullStreamInfoConfigRequest
	GetDomainName() *string
	SetEndTime(v string) *AddVsPullStreamInfoConfigRequest
	GetEndTime() *string
	SetOwnerId(v int64) *AddVsPullStreamInfoConfigRequest
	GetOwnerId() *int64
	SetSourceUrl(v string) *AddVsPullStreamInfoConfigRequest
	GetSourceUrl() *string
	SetStartTime(v string) *AddVsPullStreamInfoConfigRequest
	GetStartTime() *string
	SetStreamName(v string) *AddVsPullStreamInfoConfigRequest
	GetStreamName() *string
}

type AddVsPullStreamInfoConfigRequest struct {
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
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 2017-08-28T09:30:30Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// up.******.com.cn
	SourceUrl *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	// example:
	//
	// 2017-08-28T07:30:30Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxStream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s AddVsPullStreamInfoConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s AddVsPullStreamInfoConfigRequest) GoString() string {
	return s.String()
}

func (s *AddVsPullStreamInfoConfigRequest) GetAlways() *string {
	return s.Always
}

func (s *AddVsPullStreamInfoConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *AddVsPullStreamInfoConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddVsPullStreamInfoConfigRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *AddVsPullStreamInfoConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddVsPullStreamInfoConfigRequest) GetSourceUrl() *string {
	return s.SourceUrl
}

func (s *AddVsPullStreamInfoConfigRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *AddVsPullStreamInfoConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *AddVsPullStreamInfoConfigRequest) SetAlways(v string) *AddVsPullStreamInfoConfigRequest {
	s.Always = &v
	return s
}

func (s *AddVsPullStreamInfoConfigRequest) SetAppName(v string) *AddVsPullStreamInfoConfigRequest {
	s.AppName = &v
	return s
}

func (s *AddVsPullStreamInfoConfigRequest) SetDomainName(v string) *AddVsPullStreamInfoConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddVsPullStreamInfoConfigRequest) SetEndTime(v string) *AddVsPullStreamInfoConfigRequest {
	s.EndTime = &v
	return s
}

func (s *AddVsPullStreamInfoConfigRequest) SetOwnerId(v int64) *AddVsPullStreamInfoConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *AddVsPullStreamInfoConfigRequest) SetSourceUrl(v string) *AddVsPullStreamInfoConfigRequest {
	s.SourceUrl = &v
	return s
}

func (s *AddVsPullStreamInfoConfigRequest) SetStartTime(v string) *AddVsPullStreamInfoConfigRequest {
	s.StartTime = &v
	return s
}

func (s *AddVsPullStreamInfoConfigRequest) SetStreamName(v string) *AddVsPullStreamInfoConfigRequest {
	s.StreamName = &v
	return s
}

func (s *AddVsPullStreamInfoConfigRequest) Validate() error {
	return dara.Validate(s)
}
