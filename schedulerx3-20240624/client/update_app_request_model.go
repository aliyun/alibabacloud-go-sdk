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
	SetWorkerId(v int32) *UpdateAppRequest
	GetWorkerId() *int32
}

type UpdateAppRequest struct {
	// The access token.
	//
	// example:
	//
	// demo-app-key-20261109
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to enable logging.
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	EnableLog *bool `json:"EnableLog,omitempty" xml:"EnableLog,omitempty"`
	// The label routing strategy.
	//
	// example:
	//
	// 1
	LabelRouteStrategy *int32 `json:"LabelRouteStrategy,omitempty" xml:"LabelRouteStrategy,omitempty"`
	// The size of the task execution queue.
	//
	// > The maximum number of task instances that can run concurrently for an application. Task instances that exceed the concurrency limit are not discarded, but are placed in a queue to wait for execution.
	//
	// example:
	//
	// 10
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// The title.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test application
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The primary key ID of the worker group.
	//
	// example:
	//
	// 1
	WorkerId *int32 `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
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

func (s *UpdateAppRequest) GetWorkerId() *int32 {
	return s.WorkerId
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

func (s *UpdateAppRequest) SetWorkerId(v int32) *UpdateAppRequest {
	s.WorkerId = &v
	return s
}

func (s *UpdateAppRequest) Validate() error {
	return dara.Validate(s)
}
