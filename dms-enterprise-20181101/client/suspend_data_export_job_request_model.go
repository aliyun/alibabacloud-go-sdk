// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendDataExportJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int64) *SuspendDataExportJobRequest
	GetJobId() *int64
	SetOrderId(v int64) *SuspendDataExportJobRequest
	GetOrderId() *int64
	SetTid(v int64) *SuspendDataExportJobRequest
	GetTid() *int64
}

type SuspendDataExportJobRequest struct {
	// The ID of the SQL result set export task. You can call the [GetDataExportOrderDetail](https://help.aliyun.com/document_detail/465911.html) operation to obtain the value of this parameter. If you set this parameter to Null, no SQL result set export task is suspended.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1276****
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ticket ID. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to query the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 903****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the DMS console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s SuspendDataExportJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendDataExportJobRequest) GoString() string {
	return s.String()
}

func (s *SuspendDataExportJobRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *SuspendDataExportJobRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *SuspendDataExportJobRequest) GetTid() *int64 {
	return s.Tid
}

func (s *SuspendDataExportJobRequest) SetJobId(v int64) *SuspendDataExportJobRequest {
	s.JobId = &v
	return s
}

func (s *SuspendDataExportJobRequest) SetOrderId(v int64) *SuspendDataExportJobRequest {
	s.OrderId = &v
	return s
}

func (s *SuspendDataExportJobRequest) SetTid(v int64) *SuspendDataExportJobRequest {
	s.Tid = &v
	return s
}

func (s *SuspendDataExportJobRequest) Validate() error {
	return dara.Validate(s)
}
