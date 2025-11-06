// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForTransferOutByAuthorizationCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveBatchTaskForTransferOutByAuthorizationCodeResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveBatchTaskForTransferOutByAuthorizationCodeResponseBody
	GetTaskNo() *string
}

type SaveBatchTaskForTransferOutByAuthorizationCodeResponseBody struct {
	// example:
	//
	// E2598CAF-DBFE-494E-95EF-B42A33C178AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3cb1adc3-20e8-44ae-9e76-e812fa6fc9d8
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForTransferOutByAuthorizationCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForTransferOutByAuthorizationCodeResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForTransferOutByAuthorizationCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveBatchTaskForTransferOutByAuthorizationCodeResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveBatchTaskForTransferOutByAuthorizationCodeResponseBody) SetRequestId(v string) *SaveBatchTaskForTransferOutByAuthorizationCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForTransferOutByAuthorizationCodeResponseBody) SetTaskNo(v string) *SaveBatchTaskForTransferOutByAuthorizationCodeResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveBatchTaskForTransferOutByAuthorizationCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
