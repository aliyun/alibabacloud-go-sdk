// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForGenerateDomainCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveBatchTaskForGenerateDomainCertificateResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveBatchTaskForGenerateDomainCertificateResponseBody
	GetTaskNo() *string
}

type SaveBatchTaskForGenerateDomainCertificateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 40F46D3D-F4F3-4CCB-AC30-2DD20E32E528
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 8b1cd755-4928-4b02-adee-e5d41d7b1939
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForGenerateDomainCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForGenerateDomainCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForGenerateDomainCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveBatchTaskForGenerateDomainCertificateResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveBatchTaskForGenerateDomainCertificateResponseBody) SetRequestId(v string) *SaveBatchTaskForGenerateDomainCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForGenerateDomainCertificateResponseBody) SetTaskNo(v string) *SaveBatchTaskForGenerateDomainCertificateResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveBatchTaskForGenerateDomainCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
