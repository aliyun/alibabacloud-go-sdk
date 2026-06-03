// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForCreatingOrderRenewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForCreatingOrderRenewResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForCreatingOrderRenewResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForCreatingOrderRenewResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForCreatingOrderRenewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderRenewResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderRenewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForCreatingOrderRenewResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForCreatingOrderRenewResponseBody) SetRequestId(v string) *SaveSingleTaskForCreatingOrderRenewResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewResponseBody) SetTaskNo(v string) *SaveSingleTaskForCreatingOrderRenewResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewResponseBody) Validate() error {
	return dara.Validate(s)
}
