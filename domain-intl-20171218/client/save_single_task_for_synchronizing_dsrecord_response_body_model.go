// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForSynchronizingDSRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForSynchronizingDSRecordResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForSynchronizingDSRecordResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForSynchronizingDSRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForSynchronizingDSRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForSynchronizingDSRecordResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForSynchronizingDSRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForSynchronizingDSRecordResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForSynchronizingDSRecordResponseBody) SetRequestId(v string) *SaveSingleTaskForSynchronizingDSRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForSynchronizingDSRecordResponseBody) SetTaskNo(v string) *SaveSingleTaskForSynchronizingDSRecordResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForSynchronizingDSRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
