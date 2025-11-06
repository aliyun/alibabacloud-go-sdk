// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForTransferProhibitionLockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveBatchTaskForTransferProhibitionLockResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveBatchTaskForTransferProhibitionLockResponseBody
	GetTaskNo() *string
}

type SaveBatchTaskForTransferProhibitionLockResponseBody struct {
	// example:
	//
	// F51977F9-2B40-462B-BCCD-CF5BB1E9DB56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// d3babb0a-c939-4c25-8c65-c47b65f5492a
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForTransferProhibitionLockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForTransferProhibitionLockResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForTransferProhibitionLockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveBatchTaskForTransferProhibitionLockResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveBatchTaskForTransferProhibitionLockResponseBody) SetRequestId(v string) *SaveBatchTaskForTransferProhibitionLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForTransferProhibitionLockResponseBody) SetTaskNo(v string) *SaveBatchTaskForTransferProhibitionLockResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveBatchTaskForTransferProhibitionLockResponseBody) Validate() error {
	return dara.Validate(s)
}
