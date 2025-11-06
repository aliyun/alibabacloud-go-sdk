// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForModifyingDomainDnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveBatchTaskForModifyingDomainDnsResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveBatchTaskForModifyingDomainDnsResponseBody
	GetTaskNo() *string
}

type SaveBatchTaskForModifyingDomainDnsResponseBody struct {
	// example:
	//
	// 6A862A8A-E7AB-4C4E-8946-A74122D9CC4B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 35fb2fb7-d4d6-4478-9408-22cb63696b86
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForModifyingDomainDnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForModifyingDomainDnsResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForModifyingDomainDnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveBatchTaskForModifyingDomainDnsResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveBatchTaskForModifyingDomainDnsResponseBody) SetRequestId(v string) *SaveBatchTaskForModifyingDomainDnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForModifyingDomainDnsResponseBody) SetTaskNo(v string) *SaveBatchTaskForModifyingDomainDnsResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveBatchTaskForModifyingDomainDnsResponseBody) Validate() error {
	return dara.Validate(s)
}
