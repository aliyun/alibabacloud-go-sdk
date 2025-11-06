// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForTransferProhibitionLockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForTransferProhibitionLockResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForTransferProhibitionLockResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForTransferProhibitionLockResponseBody struct {
	// example:
	//
	// F51977F9-2B40-462B-BCCD-CF5BB1E9DB56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// d3babb0a-c939-4c25-8c65-c47b65f5492a
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForTransferProhibitionLockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForTransferProhibitionLockResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForTransferProhibitionLockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForTransferProhibitionLockResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForTransferProhibitionLockResponseBody) SetRequestId(v string) *SaveSingleTaskForTransferProhibitionLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForTransferProhibitionLockResponseBody) SetTaskNo(v string) *SaveSingleTaskForTransferProhibitionLockResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForTransferProhibitionLockResponseBody) Validate() error {
	return dara.Validate(s)
}
