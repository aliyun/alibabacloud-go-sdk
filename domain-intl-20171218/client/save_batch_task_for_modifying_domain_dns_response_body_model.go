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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
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
