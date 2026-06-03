// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForCreatingOrderRedeemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveBatchTaskForCreatingOrderRedeemResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveBatchTaskForCreatingOrderRedeemResponseBody
	GetTaskNo() *string
}

type SaveBatchTaskForCreatingOrderRedeemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderRedeemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderRedeemResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderRedeemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveBatchTaskForCreatingOrderRedeemResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveBatchTaskForCreatingOrderRedeemResponseBody) SetRequestId(v string) *SaveBatchTaskForCreatingOrderRedeemResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemResponseBody) SetTaskNo(v string) *SaveBatchTaskForCreatingOrderRedeemResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemResponseBody) Validate() error {
	return dara.Validate(s)
}
