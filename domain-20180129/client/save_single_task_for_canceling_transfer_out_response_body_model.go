// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForCancelingTransferOutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForCancelingTransferOutResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForCancelingTransferOutResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForCancelingTransferOutResponseBody struct {
	// example:
	//
	// 40F46D3D-F4F3-4CCB-AC30-2DD20E32E528
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3cb1adc3-20e8-44ae-9e76-e812fa6fc9d8
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForCancelingTransferOutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForCancelingTransferOutResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCancelingTransferOutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForCancelingTransferOutResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForCancelingTransferOutResponseBody) SetRequestId(v string) *SaveSingleTaskForCancelingTransferOutResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForCancelingTransferOutResponseBody) SetTaskNo(v string) *SaveSingleTaskForCancelingTransferOutResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForCancelingTransferOutResponseBody) Validate() error {
	return dara.Validate(s)
}
