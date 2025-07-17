// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDIJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIJobId(v int64) *GetDIJobRequest
	GetDIJobId() *int64
	SetId(v int64) *GetDIJobRequest
	GetId() *int64
	SetProjectId(v int64) *GetDIJobRequest
	GetProjectId() *int64
	SetWithDetails(v bool) *GetDIJobRequest
	GetWithDetails() *bool
}

type GetDIJobRequest struct {
	// Deprecated
	//
	// This parameter is deprecated. Use the Id parameter instead.
	//
	// example:
	//
	// 11588
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// 11588
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Specifies whether to return detailed configuration information, including TransformationRules, TableMappings, and JobSettings. Valid values: true and false. Default value: true.
	//
	// example:
	//
	// true
	WithDetails *bool `json:"WithDetails,omitempty" xml:"WithDetails,omitempty"`
}

func (s GetDIJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobRequest) GoString() string {
	return s.String()
}

func (s *GetDIJobRequest) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *GetDIJobRequest) GetId() *int64 {
	return s.Id
}

func (s *GetDIJobRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDIJobRequest) GetWithDetails() *bool {
	return s.WithDetails
}

func (s *GetDIJobRequest) SetDIJobId(v int64) *GetDIJobRequest {
	s.DIJobId = &v
	return s
}

func (s *GetDIJobRequest) SetId(v int64) *GetDIJobRequest {
	s.Id = &v
	return s
}

func (s *GetDIJobRequest) SetProjectId(v int64) *GetDIJobRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDIJobRequest) SetWithDetails(v bool) *GetDIJobRequest {
	s.WithDetails = &v
	return s
}

func (s *GetDIJobRequest) Validate() error {
	return dara.Validate(s)
}
