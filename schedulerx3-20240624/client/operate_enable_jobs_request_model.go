// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateEnableJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateEnableJobsRequest
	GetAppName() *string
	SetClusterId(v string) *OperateEnableJobsRequest
	GetClusterId() *string
	SetJobIds(v []*int64) *OperateEnableJobsRequest
	GetJobIds() []*int64
}

type OperateEnableJobsRequest struct {
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

func (s OperateEnableJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateEnableJobsRequest) GoString() string {
	return s.String()
}

func (s *OperateEnableJobsRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateEnableJobsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateEnableJobsRequest) GetJobIds() []*int64 {
	return s.JobIds
}

func (s *OperateEnableJobsRequest) SetAppName(v string) *OperateEnableJobsRequest {
	s.AppName = &v
	return s
}

func (s *OperateEnableJobsRequest) SetClusterId(v string) *OperateEnableJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateEnableJobsRequest) SetJobIds(v []*int64) *OperateEnableJobsRequest {
	s.JobIds = v
	return s
}

func (s *OperateEnableJobsRequest) Validate() error {
	return dara.Validate(s)
}
