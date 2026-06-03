// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForDomainNameProxyServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveBatchTaskForDomainNameProxyServiceResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveBatchTaskForDomainNameProxyServiceResponseBody
	GetTaskNo() *string
}

type SaveBatchTaskForDomainNameProxyServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForDomainNameProxyServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForDomainNameProxyServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForDomainNameProxyServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveBatchTaskForDomainNameProxyServiceResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveBatchTaskForDomainNameProxyServiceResponseBody) SetRequestId(v string) *SaveBatchTaskForDomainNameProxyServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForDomainNameProxyServiceResponseBody) SetTaskNo(v string) *SaveBatchTaskForDomainNameProxyServiceResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveBatchTaskForDomainNameProxyServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
