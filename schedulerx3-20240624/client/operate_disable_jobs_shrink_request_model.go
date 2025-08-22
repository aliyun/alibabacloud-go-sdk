// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateDisableJobsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateDisableJobsShrinkRequest
	GetAppName() *string
	SetClusterId(v string) *OperateDisableJobsShrinkRequest
	GetClusterId() *string
	SetJobIdsShrink(v string) *OperateDisableJobsShrinkRequest
	GetJobIdsShrink() *string
}

type OperateDisableJobsShrinkRequest struct {
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
	// -
	JobIdsShrink *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
}

func (s OperateDisableJobsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateDisableJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *OperateDisableJobsShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateDisableJobsShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateDisableJobsShrinkRequest) GetJobIdsShrink() *string {
	return s.JobIdsShrink
}

func (s *OperateDisableJobsShrinkRequest) SetAppName(v string) *OperateDisableJobsShrinkRequest {
	s.AppName = &v
	return s
}

func (s *OperateDisableJobsShrinkRequest) SetClusterId(v string) *OperateDisableJobsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateDisableJobsShrinkRequest) SetJobIdsShrink(v string) *OperateDisableJobsShrinkRequest {
	s.JobIdsShrink = &v
	return s
}

func (s *OperateDisableJobsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
