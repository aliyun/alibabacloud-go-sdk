// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForCreatingOrderRedeemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForCreatingOrderRedeemResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForCreatingOrderRedeemResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForCreatingOrderRedeemResponseBody struct {
	// example:
	//
	// 40F46D3D-F4F3-4CCB-AC30-2DD20E32E528
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3cb1adc3-20e8-44ae-9e76-e812fa6fc9d8
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForCreatingOrderRedeemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderRedeemResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderRedeemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForCreatingOrderRedeemResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForCreatingOrderRedeemResponseBody) SetRequestId(v string) *SaveSingleTaskForCreatingOrderRedeemResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemResponseBody) SetTaskNo(v string) *SaveSingleTaskForCreatingOrderRedeemResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemResponseBody) Validate() error {
	return dara.Validate(s)
}
