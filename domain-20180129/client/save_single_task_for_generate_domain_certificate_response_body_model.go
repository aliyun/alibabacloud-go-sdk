// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForGenerateDomainCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForGenerateDomainCertificateResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForGenerateDomainCertificateResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForGenerateDomainCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForGenerateDomainCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForGenerateDomainCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForGenerateDomainCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForGenerateDomainCertificateResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForGenerateDomainCertificateResponseBody) SetRequestId(v string) *SaveSingleTaskForGenerateDomainCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForGenerateDomainCertificateResponseBody) SetTaskNo(v string) *SaveSingleTaskForGenerateDomainCertificateResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForGenerateDomainCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
