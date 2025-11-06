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
	// example:
	//
	// B7AB5469-5E38-4AA9-A920-C65B7A9C8E6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3cb1adc3-20e8-44ae-9e76-e812fa6fc9d8
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
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
