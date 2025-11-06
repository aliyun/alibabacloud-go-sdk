// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForReserveDropListDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForReserveDropListDomainResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForReserveDropListDomainResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForReserveDropListDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForReserveDropListDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForReserveDropListDomainResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForReserveDropListDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForReserveDropListDomainResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForReserveDropListDomainResponseBody) SetRequestId(v string) *SaveSingleTaskForReserveDropListDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForReserveDropListDomainResponseBody) SetTaskNo(v string) *SaveSingleTaskForReserveDropListDomainResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForReserveDropListDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
