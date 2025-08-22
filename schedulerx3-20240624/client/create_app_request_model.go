// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *CreateAppRequest
	GetAccessToken() *string
	SetAppName(v string) *CreateAppRequest
	GetAppName() *string
	SetAppType(v int32) *CreateAppRequest
	GetAppType() *int32
	SetClusterId(v string) *CreateAppRequest
	GetClusterId() *string
	SetEnableLog(v bool) *CreateAppRequest
	GetEnableLog() *bool
	SetLabelRouteStrategy(v int32) *CreateAppRequest
	GetLabelRouteStrategy() *int32
	SetMaxConcurrency(v int32) *CreateAppRequest
	GetMaxConcurrency() *int32
	SetTitle(v string) *CreateAppRequest
	GetTitle() *string
}

type CreateAppRequest struct {
	// example:
	//
	// ltk1ZXHv6LvibZypFkPHzRA
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	AppType *int32 `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// true
	EnableLog          *bool  `json:"EnableLog,omitempty" xml:"EnableLog,omitempty"`
	LabelRouteStrategy *int32 `json:"LabelRouteStrategy,omitempty" xml:"LabelRouteStrategy,omitempty"`
	// example:
	//
	// 10
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateAppRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateAppRequest) GetAppType() *int32 {
	return s.AppType
}

func (s *CreateAppRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateAppRequest) GetEnableLog() *bool {
	return s.EnableLog
}

func (s *CreateAppRequest) GetLabelRouteStrategy() *int32 {
	return s.LabelRouteStrategy
}

func (s *CreateAppRequest) GetMaxConcurrency() *int32 {
	return s.MaxConcurrency
}

func (s *CreateAppRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateAppRequest) SetAccessToken(v string) *CreateAppRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateAppRequest) SetAppName(v string) *CreateAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppRequest) SetAppType(v int32) *CreateAppRequest {
	s.AppType = &v
	return s
}

func (s *CreateAppRequest) SetClusterId(v string) *CreateAppRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateAppRequest) SetEnableLog(v bool) *CreateAppRequest {
	s.EnableLog = &v
	return s
}

func (s *CreateAppRequest) SetLabelRouteStrategy(v int32) *CreateAppRequest {
	s.LabelRouteStrategy = &v
	return s
}

func (s *CreateAppRequest) SetMaxConcurrency(v int32) *CreateAppRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *CreateAppRequest) SetTitle(v string) *CreateAppRequest {
	s.Title = &v
	return s
}

func (s *CreateAppRequest) Validate() error {
	return dara.Validate(s)
}
