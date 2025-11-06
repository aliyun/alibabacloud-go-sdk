// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody
	GetTaskNo() *string
}

type SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody) SetRequestId(v string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody) SetTaskNo(v string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody) Validate() error {
	return dara.Validate(s)
}
