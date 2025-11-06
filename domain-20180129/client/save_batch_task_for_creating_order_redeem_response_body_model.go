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
	// example:
	//
	// F51977F9-2B40-462B-BCCD-CF5BB1E9DB56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// d3babb0a-c939-4c25-8c65-c47b65f5492a
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
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
