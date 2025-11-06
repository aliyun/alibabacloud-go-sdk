// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForApplyQuickTransferOutOpenlyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveBatchTaskForApplyQuickTransferOutOpenlyResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveBatchTaskForApplyQuickTransferOutOpenlyResponseBody
	GetTaskNo() *string
}

type SaveBatchTaskForApplyQuickTransferOutOpenlyResponseBody struct {
	// example:
	//
	// D6CB3623-4726-4947-AC2B-2C6E673B447C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// d3babb0a-c939-4c25-8c65-c47b65f5492a
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForApplyQuickTransferOutOpenlyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForApplyQuickTransferOutOpenlyResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForApplyQuickTransferOutOpenlyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveBatchTaskForApplyQuickTransferOutOpenlyResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveBatchTaskForApplyQuickTransferOutOpenlyResponseBody) SetRequestId(v string) *SaveBatchTaskForApplyQuickTransferOutOpenlyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForApplyQuickTransferOutOpenlyResponseBody) SetTaskNo(v string) *SaveBatchTaskForApplyQuickTransferOutOpenlyResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveBatchTaskForApplyQuickTransferOutOpenlyResponseBody) Validate() error {
	return dara.Validate(s)
}
