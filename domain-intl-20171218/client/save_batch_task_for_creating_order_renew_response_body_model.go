// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForCreatingOrderRenewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveBatchTaskForCreatingOrderRenewResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveBatchTaskForCreatingOrderRenewResponseBody
	GetTaskNo() *string
}

type SaveBatchTaskForCreatingOrderRenewResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderRenewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderRenewResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderRenewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveBatchTaskForCreatingOrderRenewResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveBatchTaskForCreatingOrderRenewResponseBody) SetRequestId(v string) *SaveBatchTaskForCreatingOrderRenewResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewResponseBody) SetTaskNo(v string) *SaveBatchTaskForCreatingOrderRenewResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewResponseBody) Validate() error {
	return dara.Validate(s)
}
