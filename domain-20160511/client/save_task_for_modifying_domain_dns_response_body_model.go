// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForModifyingDomainDnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveTaskForModifyingDomainDnsResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveTaskForModifyingDomainDnsResponseBody
	GetTaskNo() *string
}

type SaveTaskForModifyingDomainDnsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveTaskForModifyingDomainDnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForModifyingDomainDnsResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTaskForModifyingDomainDnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveTaskForModifyingDomainDnsResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveTaskForModifyingDomainDnsResponseBody) SetRequestId(v string) *SaveTaskForModifyingDomainDnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTaskForModifyingDomainDnsResponseBody) SetTaskNo(v string) *SaveTaskForModifyingDomainDnsResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveTaskForModifyingDomainDnsResponseBody) Validate() error {
	return dara.Validate(s)
}
