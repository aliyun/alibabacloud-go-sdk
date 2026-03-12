// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody
	GetSuccess() *bool
	SetTaskNo(v string) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody
	GetTaskNo() *string
}

type SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody) SetRequestId(v string) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody) SetSuccess(v bool) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody {
	s.Success = &v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody) SetTaskNo(v string) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody) Validate() error {
	return dara.Validate(s)
}
