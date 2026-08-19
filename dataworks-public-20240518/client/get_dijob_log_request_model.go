// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDIJobLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIJobId(v int64) *GetDIJobLogRequest
	GetDIJobId() *int64
	SetFailoverId(v int64) *GetDIJobLogRequest
	GetFailoverId() *int64
	SetId(v int64) *GetDIJobLogRequest
	GetId() *int64
	SetInstanceId(v int64) *GetDIJobLogRequest
	GetInstanceId() *int64
	SetNodeType(v string) *GetDIJobLogRequest
	GetNodeType() *string
	SetPageNumber(v int32) *GetDIJobLogRequest
	GetPageNumber() *int32
}

type GetDIJobLogRequest struct {
	// Deprecated
	//
	// **[Deprecated]*	- This parameter is deprecated. Use the Id parameter instead.
	//
	// example:
	//
	// 10000
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The failover ID.
	//
	// example:
	//
	// 10
	FailoverId *int64 `json:"FailoverId,omitempty" xml:"FailoverId,omitempty"`
	// The node ID.
	//
	// example:
	//
	// 10000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 6153616438
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node type. This parameter is applicable only to tasks that run on resource group 2.0. Valid values:
	//
	// 	- **MASTER**: retrieves the JobManager logs.
	//
	// 	- **WORKER**: retrieves the TaskManager logs.
	//
	// example:
	//
	// MASTER
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The page number for paging. The value must be a positive integer greater than or equal to 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s GetDIJobLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobLogRequest) GoString() string {
	return s.String()
}

func (s *GetDIJobLogRequest) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *GetDIJobLogRequest) GetFailoverId() *int64 {
	return s.FailoverId
}

func (s *GetDIJobLogRequest) GetId() *int64 {
	return s.Id
}

func (s *GetDIJobLogRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetDIJobLogRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *GetDIJobLogRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetDIJobLogRequest) SetDIJobId(v int64) *GetDIJobLogRequest {
	s.DIJobId = &v
	return s
}

func (s *GetDIJobLogRequest) SetFailoverId(v int64) *GetDIJobLogRequest {
	s.FailoverId = &v
	return s
}

func (s *GetDIJobLogRequest) SetId(v int64) *GetDIJobLogRequest {
	s.Id = &v
	return s
}

func (s *GetDIJobLogRequest) SetInstanceId(v int64) *GetDIJobLogRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDIJobLogRequest) SetNodeType(v string) *GetDIJobLogRequest {
	s.NodeType = &v
	return s
}

func (s *GetDIJobLogRequest) SetPageNumber(v int32) *GetDIJobLogRequest {
	s.PageNumber = &v
	return s
}

func (s *GetDIJobLogRequest) Validate() error {
	return dara.Validate(s)
}
