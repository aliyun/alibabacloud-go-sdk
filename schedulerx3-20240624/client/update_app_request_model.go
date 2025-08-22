// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *UpdateAppRequest
	GetAccessToken() *string
	SetAppName(v string) *UpdateAppRequest
	GetAppName() *string
	SetClusterId(v string) *UpdateAppRequest
	GetClusterId() *string
	SetEnableLog(v bool) *UpdateAppRequest
	GetEnableLog() *bool
	SetLabelRouteStrategy(v int32) *UpdateAppRequest
	GetLabelRouteStrategy() *int32
	SetMaxConcurrency(v int32) *UpdateAppRequest
	GetMaxConcurrency() *int32
	SetTitle(v string) *UpdateAppRequest
	GetTitle() *string
}

type UpdateAppRequest struct {
	// example:
	//
	// f312159702f4469585586ed5a6904163v3
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
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

func (s UpdateAppRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *UpdateAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateAppRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateAppRequest) GetEnableLog() *bool {
	return s.EnableLog
}

func (s *UpdateAppRequest) GetLabelRouteStrategy() *int32 {
	return s.LabelRouteStrategy
}

func (s *UpdateAppRequest) GetMaxConcurrency() *int32 {
	return s.MaxConcurrency
}

func (s *UpdateAppRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateAppRequest) SetAccessToken(v string) *UpdateAppRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateAppRequest) SetAppName(v string) *UpdateAppRequest {
	s.AppName = &v
	return s
}

func (s *UpdateAppRequest) SetClusterId(v string) *UpdateAppRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateAppRequest) SetEnableLog(v bool) *UpdateAppRequest {
	s.EnableLog = &v
	return s
}

func (s *UpdateAppRequest) SetLabelRouteStrategy(v int32) *UpdateAppRequest {
	s.LabelRouteStrategy = &v
	return s
}

func (s *UpdateAppRequest) SetMaxConcurrency(v int32) *UpdateAppRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *UpdateAppRequest) SetTitle(v string) *UpdateAppRequest {
	s.Title = &v
	return s
}

func (s *UpdateAppRequest) Validate() error {
	return dara.Validate(s)
}
