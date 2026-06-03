// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForCreatingOrderActivateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveBatchTaskForCreatingOrderActivateResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveBatchTaskForCreatingOrderActivateResponseBody
	GetTaskNo() *string
}

type SaveBatchTaskForCreatingOrderActivateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderActivateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderActivateResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderActivateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveBatchTaskForCreatingOrderActivateResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveBatchTaskForCreatingOrderActivateResponseBody) SetRequestId(v string) *SaveBatchTaskForCreatingOrderActivateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateResponseBody) SetTaskNo(v string) *SaveBatchTaskForCreatingOrderActivateResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateResponseBody) Validate() error {
	return dara.Validate(s)
}
