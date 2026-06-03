// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForCreatingOrderTransferResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveBatchTaskForCreatingOrderTransferResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveBatchTaskForCreatingOrderTransferResponseBody
	GetTaskNo() *string
}

type SaveBatchTaskForCreatingOrderTransferResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderTransferResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderTransferResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderTransferResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveBatchTaskForCreatingOrderTransferResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveBatchTaskForCreatingOrderTransferResponseBody) SetRequestId(v string) *SaveBatchTaskForCreatingOrderTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferResponseBody) SetTaskNo(v string) *SaveBatchTaskForCreatingOrderTransferResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferResponseBody) Validate() error {
	return dara.Validate(s)
}
