// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferTaskApprovalCallbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOssBucketName(v string) *TransferTaskApprovalCallbackRequest
	GetOssBucketName() *string
	SetOssBucketRegionId(v string) *TransferTaskApprovalCallbackRequest
	GetOssBucketRegionId() *string
	SetResult(v string) *TransferTaskApprovalCallbackRequest
	GetResult() *string
	SetTaskId(v string) *TransferTaskApprovalCallbackRequest
	GetTaskId() *string
}

type TransferTaskApprovalCallbackRequest struct {
	// example:
	//
	// ed****-17337752804***
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// example:
	//
	// cn-hangzhou
	OssBucketRegionId *string `json:"OssBucketRegionId,omitempty" xml:"OssBucketRegionId,omitempty"`
	// example:
	//
	// Approved
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// trt-msndfksm18fs****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TransferTaskApprovalCallbackRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferTaskApprovalCallbackRequest) GoString() string {
	return s.String()
}

func (s *TransferTaskApprovalCallbackRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *TransferTaskApprovalCallbackRequest) GetOssBucketRegionId() *string {
	return s.OssBucketRegionId
}

func (s *TransferTaskApprovalCallbackRequest) GetResult() *string {
	return s.Result
}

func (s *TransferTaskApprovalCallbackRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *TransferTaskApprovalCallbackRequest) SetOssBucketName(v string) *TransferTaskApprovalCallbackRequest {
	s.OssBucketName = &v
	return s
}

func (s *TransferTaskApprovalCallbackRequest) SetOssBucketRegionId(v string) *TransferTaskApprovalCallbackRequest {
	s.OssBucketRegionId = &v
	return s
}

func (s *TransferTaskApprovalCallbackRequest) SetResult(v string) *TransferTaskApprovalCallbackRequest {
	s.Result = &v
	return s
}

func (s *TransferTaskApprovalCallbackRequest) SetTaskId(v string) *TransferTaskApprovalCallbackRequest {
	s.TaskId = &v
	return s
}

func (s *TransferTaskApprovalCallbackRequest) Validate() error {
	return dara.Validate(s)
}
