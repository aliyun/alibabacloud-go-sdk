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
	// example:
	//
	// F51977F9-2B40-462B-BCCD-CF5BB1E9DB56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// d3babb0a-c939-4c25-8c65-c47b65f5492a
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
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
