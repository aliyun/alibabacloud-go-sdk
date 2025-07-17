// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOnlineDDLProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobDetailId(v int64) *GetOnlineDDLProgressRequest
	GetJobDetailId() *int64
	SetTid(v int64) *GetOnlineDDLProgressRequest
	GetTid() *int64
}

type GetOnlineDDLProgressRequest struct {
	// The ID of the OnlineDDL SQL task details. You can call the [ListDBTaskSQLJobDetail](https://help.aliyun.com/document_detail/207073.html) operation to obtain the task detail ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15***
	JobDetailId *int64 `json:"JobDetailId,omitempty" xml:"JobDetailId,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the "View information about the current tenant" section of the [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html) topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetOnlineDDLProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOnlineDDLProgressRequest) GoString() string {
	return s.String()
}

func (s *GetOnlineDDLProgressRequest) GetJobDetailId() *int64 {
	return s.JobDetailId
}

func (s *GetOnlineDDLProgressRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetOnlineDDLProgressRequest) SetJobDetailId(v int64) *GetOnlineDDLProgressRequest {
	s.JobDetailId = &v
	return s
}

func (s *GetOnlineDDLProgressRequest) SetTid(v int64) *GetOnlineDDLProgressRequest {
	s.Tid = &v
	return s
}

func (s *GetOnlineDDLProgressRequest) Validate() error {
	return dara.Validate(s)
}
