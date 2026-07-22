// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateRerunJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int64) *OperateRerunJobRequest
	GetAppId() *int64
	SetAppName(v string) *OperateRerunJobRequest
	GetAppName() *string
	SetClusterId(v string) *OperateRerunJobRequest
	GetClusterId() *string
	SetDataTime(v string) *OperateRerunJobRequest
	GetDataTime() *string
	SetEndDate(v int64) *OperateRerunJobRequest
	GetEndDate() *int64
	SetJobId(v int64) *OperateRerunJobRequest
	GetJobId() *int64
	SetStartDate(v int64) *OperateRerunJobRequest
	GetStartDate() *int64
}

type OperateRerunJobRequest struct {
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
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
	// The data timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 14:11:10
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// The end time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1698458024000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1698458024000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s OperateRerunJobRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateRerunJobRequest) GoString() string {
	return s.String()
}

func (s *OperateRerunJobRequest) GetAppId() *int64 {
	return s.AppId
}

func (s *OperateRerunJobRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateRerunJobRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateRerunJobRequest) GetDataTime() *string {
	return s.DataTime
}

func (s *OperateRerunJobRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *OperateRerunJobRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *OperateRerunJobRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *OperateRerunJobRequest) SetAppId(v int64) *OperateRerunJobRequest {
	s.AppId = &v
	return s
}

func (s *OperateRerunJobRequest) SetAppName(v string) *OperateRerunJobRequest {
	s.AppName = &v
	return s
}

func (s *OperateRerunJobRequest) SetClusterId(v string) *OperateRerunJobRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateRerunJobRequest) SetDataTime(v string) *OperateRerunJobRequest {
	s.DataTime = &v
	return s
}

func (s *OperateRerunJobRequest) SetEndDate(v int64) *OperateRerunJobRequest {
	s.EndDate = &v
	return s
}

func (s *OperateRerunJobRequest) SetJobId(v int64) *OperateRerunJobRequest {
	s.JobId = &v
	return s
}

func (s *OperateRerunJobRequest) SetStartDate(v int64) *OperateRerunJobRequest {
	s.StartDate = &v
	return s
}

func (s *OperateRerunJobRequest) Validate() error {
	return dara.Validate(s)
}
