// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForSubmittingDomainNameCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveTaskForSubmittingDomainNameCredentialResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveTaskForSubmittingDomainNameCredentialResponseBody
	GetTaskNo() *string
}

type SaveTaskForSubmittingDomainNameCredentialResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveTaskForSubmittingDomainNameCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForSubmittingDomainNameCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainNameCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveTaskForSubmittingDomainNameCredentialResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveTaskForSubmittingDomainNameCredentialResponseBody) SetRequestId(v string) *SaveTaskForSubmittingDomainNameCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialResponseBody) SetTaskNo(v string) *SaveTaskForSubmittingDomainNameCredentialResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
