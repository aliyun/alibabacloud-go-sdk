// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForDomainNameProxyServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForDomainNameProxyServiceResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForDomainNameProxyServiceResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForDomainNameProxyServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForDomainNameProxyServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForDomainNameProxyServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDomainNameProxyServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForDomainNameProxyServiceResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForDomainNameProxyServiceResponseBody) SetRequestId(v string) *SaveSingleTaskForDomainNameProxyServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForDomainNameProxyServiceResponseBody) SetTaskNo(v string) *SaveSingleTaskForDomainNameProxyServiceResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForDomainNameProxyServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
