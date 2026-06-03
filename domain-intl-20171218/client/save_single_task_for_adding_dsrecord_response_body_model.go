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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
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
