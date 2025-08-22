// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateDisableJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateDisableJobsRequest
	GetAppName() *string
	SetClusterId(v string) *OperateDisableJobsRequest
	GetClusterId() *string
	SetJobIds(v []*int64) *OperateDisableJobsRequest
	GetJobIds() []*int64
}

type OperateDisableJobsRequest struct {
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

func (s OperateDisableJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateDisableJobsRequest) GoString() string {
	return s.String()
}

func (s *OperateDisableJobsRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateDisableJobsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateDisableJobsRequest) GetJobIds() []*int64 {
	return s.JobIds
}

func (s *OperateDisableJobsRequest) SetAppName(v string) *OperateDisableJobsRequest {
	s.AppName = &v
	return s
}

func (s *OperateDisableJobsRequest) SetClusterId(v string) *OperateDisableJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateDisableJobsRequest) SetJobIds(v []*int64) *OperateDisableJobsRequest {
	s.JobIds = v
	return s
}

func (s *OperateDisableJobsRequest) Validate() error {
	return dara.Validate(s)
}
