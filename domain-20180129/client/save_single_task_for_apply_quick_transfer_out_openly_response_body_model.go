// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForApplyQuickTransferOutOpenlyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForApplyQuickTransferOutOpenlyResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForApplyQuickTransferOutOpenlyResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForApplyQuickTransferOutOpenlyResponseBody struct {
	// example:
	//
	// D200000-C0B9-4CD3-B92A-9B44A000000
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3cb1adc3-20e8-44ae-9e76-e812fa6fc9d8
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForApplyQuickTransferOutOpenlyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForApplyQuickTransferOutOpenlyResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForApplyQuickTransferOutOpenlyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForApplyQuickTransferOutOpenlyResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForApplyQuickTransferOutOpenlyResponseBody) SetRequestId(v string) *SaveSingleTaskForApplyQuickTransferOutOpenlyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForApplyQuickTransferOutOpenlyResponseBody) SetTaskNo(v string) *SaveSingleTaskForApplyQuickTransferOutOpenlyResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForApplyQuickTransferOutOpenlyResponseBody) Validate() error {
	return dara.Validate(s)
}
