// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFormationInstsByTaskIDRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *QueryFormationInstsByTaskIDRequest
	GetDBClusterId() *string
	SetRegionId(v string) *QueryFormationInstsByTaskIDRequest
	GetRegionId() *string
	SetTaskId(v string) *QueryFormationInstsByTaskIDRequest
	GetTaskId() *string
}

type QueryFormationInstsByTaskIDRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) to query available region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryFormationInstsByTaskIDRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFormationInstsByTaskIDRequest) GoString() string {
	return s.String()
}

func (s *QueryFormationInstsByTaskIDRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *QueryFormationInstsByTaskIDRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryFormationInstsByTaskIDRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryFormationInstsByTaskIDRequest) SetDBClusterId(v string) *QueryFormationInstsByTaskIDRequest {
	s.DBClusterId = &v
	return s
}

func (s *QueryFormationInstsByTaskIDRequest) SetRegionId(v string) *QueryFormationInstsByTaskIDRequest {
	s.RegionId = &v
	return s
}

func (s *QueryFormationInstsByTaskIDRequest) SetTaskId(v string) *QueryFormationInstsByTaskIDRequest {
	s.TaskId = &v
	return s
}

func (s *QueryFormationInstsByTaskIDRequest) Validate() error {
	return dara.Validate(s)
}
