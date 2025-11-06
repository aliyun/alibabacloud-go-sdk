// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForAddingDSRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForAddingDSRecordResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForAddingDSRecordResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForAddingDSRecordResponseBody struct {
	// example:
	//
	// E2598CAF-DBFE-494E-95EF-B42A33C178AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// e893148f-6343-4ae1-9eba-6e2a4116e142
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForAddingDSRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForAddingDSRecordResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForAddingDSRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForAddingDSRecordResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForAddingDSRecordResponseBody) SetRequestId(v string) *SaveSingleTaskForAddingDSRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordResponseBody) SetTaskNo(v string) *SaveSingleTaskForAddingDSRecordResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
