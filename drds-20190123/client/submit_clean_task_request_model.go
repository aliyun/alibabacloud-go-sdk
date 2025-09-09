// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCleanTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *SubmitCleanTaskRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *SubmitCleanTaskRequest
	GetDrdsInstanceId() *string
	SetExpandType(v string) *SubmitCleanTaskRequest
	GetExpandType() *string
	SetJobId(v string) *SubmitCleanTaskRequest
	GetJobId() *string
	SetParentJobId(v string) *SubmitCleanTaskRequest
	GetParentJobId() *string
}

type SubmitCleanTaskRequest struct {
	// The name of the database that is scaled out.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds*********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The scale-out type. Valid values:
	//
	// 	- smooth_expand: smooth scale-out
	//
	// 	- hot_expand: hot-spot scale-out
	//
	// This parameter is required.
	//
	// example:
	//
	// smooth_expand
	ExpandType *string `json:"ExpandType,omitempty" xml:"ExpandType,omitempty"`
	// The job ID of the scale-out task. The value of this parameter is the same as that of the ParentJobId parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the scale-out task. This parameter is returned if you send a request for the smooth scale-out task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ParentJobId *string `json:"ParentJobId,omitempty" xml:"ParentJobId,omitempty"`
}

func (s SubmitCleanTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitCleanTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitCleanTaskRequest) GetDbName() *string {
	return s.DbName
}

func (s *SubmitCleanTaskRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *SubmitCleanTaskRequest) GetExpandType() *string {
	return s.ExpandType
}

func (s *SubmitCleanTaskRequest) GetJobId() *string {
	return s.JobId
}

func (s *SubmitCleanTaskRequest) GetParentJobId() *string {
	return s.ParentJobId
}

func (s *SubmitCleanTaskRequest) SetDbName(v string) *SubmitCleanTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitCleanTaskRequest) SetDrdsInstanceId(v string) *SubmitCleanTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitCleanTaskRequest) SetExpandType(v string) *SubmitCleanTaskRequest {
	s.ExpandType = &v
	return s
}

func (s *SubmitCleanTaskRequest) SetJobId(v string) *SubmitCleanTaskRequest {
	s.JobId = &v
	return s
}

func (s *SubmitCleanTaskRequest) SetParentJobId(v string) *SubmitCleanTaskRequest {
	s.ParentJobId = &v
	return s
}

func (s *SubmitCleanTaskRequest) Validate() error {
	return dara.Validate(s)
}
