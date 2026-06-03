// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForModifyingDSRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForModifyingDSRecordResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForModifyingDSRecordResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForModifyingDSRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForModifyingDSRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForModifyingDSRecordResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForModifyingDSRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForModifyingDSRecordResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForModifyingDSRecordResponseBody) SetRequestId(v string) *SaveSingleTaskForModifyingDSRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordResponseBody) SetTaskNo(v string) *SaveSingleTaskForModifyingDSRecordResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
