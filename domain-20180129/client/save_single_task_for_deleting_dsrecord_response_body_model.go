// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForDeletingDSRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForDeletingDSRecordResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForDeletingDSRecordResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForDeletingDSRecordResponseBody struct {
	// example:
	//
	// E2598CAF-DBFE-494E-95EF-B42A33C178AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// e893148f-6343-4ae1-9eba-6e2a4116e142
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForDeletingDSRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForDeletingDSRecordResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDeletingDSRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForDeletingDSRecordResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForDeletingDSRecordResponseBody) SetRequestId(v string) *SaveSingleTaskForDeletingDSRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForDeletingDSRecordResponseBody) SetTaskNo(v string) *SaveSingleTaskForDeletingDSRecordResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForDeletingDSRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
