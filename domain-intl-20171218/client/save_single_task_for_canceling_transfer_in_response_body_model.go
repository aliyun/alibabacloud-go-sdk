// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForCancelingTransferInResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForCancelingTransferInResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForCancelingTransferInResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForCancelingTransferInResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForCancelingTransferInResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForCancelingTransferInResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCancelingTransferInResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForCancelingTransferInResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForCancelingTransferInResponseBody) SetRequestId(v string) *SaveSingleTaskForCancelingTransferInResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForCancelingTransferInResponseBody) SetTaskNo(v string) *SaveSingleTaskForCancelingTransferInResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForCancelingTransferInResponseBody) Validate() error {
	return dara.Validate(s)
}
