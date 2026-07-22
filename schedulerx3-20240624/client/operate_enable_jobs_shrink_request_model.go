// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateEnableJobsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppGroupId(v int64) *OperateEnableJobsShrinkRequest
	GetAppGroupId() *int64
	SetAppName(v string) *OperateEnableJobsShrinkRequest
	GetAppName() *string
	SetClusterId(v string) *OperateEnableJobsShrinkRequest
	GetClusterId() *string
	SetJobIdsShrink(v string) *OperateEnableJobsShrinkRequest
	GetJobIdsShrink() *string
}

type OperateEnableJobsShrinkRequest struct {
	AppGroupId *int64 `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
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
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	JobIdsShrink *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
}

func (s OperateEnableJobsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateEnableJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *OperateEnableJobsShrinkRequest) GetAppGroupId() *int64 {
	return s.AppGroupId
}

func (s *OperateEnableJobsShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateEnableJobsShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateEnableJobsShrinkRequest) GetJobIdsShrink() *string {
	return s.JobIdsShrink
}

func (s *OperateEnableJobsShrinkRequest) SetAppGroupId(v int64) *OperateEnableJobsShrinkRequest {
	s.AppGroupId = &v
	return s
}

func (s *OperateEnableJobsShrinkRequest) SetAppName(v string) *OperateEnableJobsShrinkRequest {
	s.AppName = &v
	return s
}

func (s *OperateEnableJobsShrinkRequest) SetClusterId(v string) *OperateEnableJobsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateEnableJobsShrinkRequest) SetJobIdsShrink(v string) *OperateEnableJobsShrinkRequest {
	s.JobIdsShrink = &v
	return s
}

func (s *OperateEnableJobsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
