// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForUpdateProhibitionLockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveBatchTaskForUpdateProhibitionLockResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveBatchTaskForUpdateProhibitionLockResponseBody
	GetTaskNo() *string
}

type SaveBatchTaskForUpdateProhibitionLockResponseBody struct {
	// example:
	//
	// F51977F9-2B40-462B-BCCD-CF5BB1E9DB56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// d3babb0a-c939-4c25-8c65-c47b65f5492a
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForUpdateProhibitionLockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForUpdateProhibitionLockResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdateProhibitionLockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveBatchTaskForUpdateProhibitionLockResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveBatchTaskForUpdateProhibitionLockResponseBody) SetRequestId(v string) *SaveBatchTaskForUpdateProhibitionLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForUpdateProhibitionLockResponseBody) SetTaskNo(v string) *SaveBatchTaskForUpdateProhibitionLockResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveBatchTaskForUpdateProhibitionLockResponseBody) Validate() error {
	return dara.Validate(s)
}
