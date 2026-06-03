// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForApprovingTransferOutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForApprovingTransferOutResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForApprovingTransferOutResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForApprovingTransferOutResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForApprovingTransferOutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForApprovingTransferOutResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForApprovingTransferOutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForApprovingTransferOutResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForApprovingTransferOutResponseBody) SetRequestId(v string) *SaveSingleTaskForApprovingTransferOutResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForApprovingTransferOutResponseBody) SetTaskNo(v string) *SaveSingleTaskForApprovingTransferOutResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForApprovingTransferOutResponseBody) Validate() error {
	return dara.Validate(s)
}
