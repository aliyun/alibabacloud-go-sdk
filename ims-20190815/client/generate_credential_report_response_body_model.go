// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCredentialReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GenerateCredentialReportResponseBody
	GetRequestId() *string
	SetState(v string) *GenerateCredentialReportResponseBody
	GetState() *string
}

type GenerateCredentialReportResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GenerateCredentialReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateCredentialReportResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCredentialReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateCredentialReportResponseBody) GetState() *string {
	return s.State
}

func (s *GenerateCredentialReportResponseBody) SetRequestId(v string) *GenerateCredentialReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateCredentialReportResponseBody) SetState(v string) *GenerateCredentialReportResponseBody {
	s.State = &v
	return s
}

func (s *GenerateCredentialReportResponseBody) Validate() error {
	return dara.Validate(s)
}
