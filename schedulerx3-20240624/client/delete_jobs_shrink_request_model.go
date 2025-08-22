// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteJobsShrinkRequest
	GetAppName() *string
	SetClusterId(v string) *DeleteJobsShrinkRequest
	GetClusterId() *string
	SetJobIdsShrink(v string) *DeleteJobsShrinkRequest
	GetJobIdsShrink() *string
}

type DeleteJobsShrinkRequest struct {
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

func (s DeleteJobsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobsShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteJobsShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteJobsShrinkRequest) GetJobIdsShrink() *string {
	return s.JobIdsShrink
}

func (s *DeleteJobsShrinkRequest) SetAppName(v string) *DeleteJobsShrinkRequest {
	s.AppName = &v
	return s
}

func (s *DeleteJobsShrinkRequest) SetClusterId(v string) *DeleteJobsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteJobsShrinkRequest) SetJobIdsShrink(v string) *DeleteJobsShrinkRequest {
	s.JobIdsShrink = &v
	return s
}

func (s *DeleteJobsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
