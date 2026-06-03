// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForReserveDropListDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveBatchTaskForReserveDropListDomainResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveBatchTaskForReserveDropListDomainResponseBody
	GetTaskNo() *string
}

type SaveBatchTaskForReserveDropListDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForReserveDropListDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForReserveDropListDomainResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForReserveDropListDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveBatchTaskForReserveDropListDomainResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveBatchTaskForReserveDropListDomainResponseBody) SetRequestId(v string) *SaveBatchTaskForReserveDropListDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForReserveDropListDomainResponseBody) SetTaskNo(v string) *SaveBatchTaskForReserveDropListDomainResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveBatchTaskForReserveDropListDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
