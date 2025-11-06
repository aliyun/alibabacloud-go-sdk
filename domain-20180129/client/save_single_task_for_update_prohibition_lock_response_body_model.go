// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForUpdateProhibitionLockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForUpdateProhibitionLockResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForUpdateProhibitionLockResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForUpdateProhibitionLockResponseBody struct {
	// example:
	//
	// F51977F9-2B40-462B-BCCD-CF5BB1E9DB56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// d3babb0a-c939-4c25-8c65-c47b65f5492a
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForUpdateProhibitionLockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForUpdateProhibitionLockResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForUpdateProhibitionLockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForUpdateProhibitionLockResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForUpdateProhibitionLockResponseBody) SetRequestId(v string) *SaveSingleTaskForUpdateProhibitionLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForUpdateProhibitionLockResponseBody) SetTaskNo(v string) *SaveSingleTaskForUpdateProhibitionLockResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForUpdateProhibitionLockResponseBody) Validate() error {
	return dara.Validate(s)
}
