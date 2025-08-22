// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteJobsRequest
	GetAppName() *string
	SetClusterId(v string) *DeleteJobsRequest
	GetClusterId() *string
	SetJobIds(v []*int64) *DeleteJobsRequest
	GetJobIds() []*int64
}

type DeleteJobsRequest struct {
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
	JobIds []*int64 `json:"JobIds,omitempty" xml:"JobIds,omitempty" type:"Repeated"`
}

func (s DeleteJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobsRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobsRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteJobsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteJobsRequest) GetJobIds() []*int64 {
	return s.JobIds
}

func (s *DeleteJobsRequest) SetAppName(v string) *DeleteJobsRequest {
	s.AppName = &v
	return s
}

func (s *DeleteJobsRequest) SetClusterId(v string) *DeleteJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteJobsRequest) SetJobIds(v []*int64) *DeleteJobsRequest {
	s.JobIds = v
	return s
}

func (s *DeleteJobsRequest) Validate() error {
	return dara.Validate(s)
}
