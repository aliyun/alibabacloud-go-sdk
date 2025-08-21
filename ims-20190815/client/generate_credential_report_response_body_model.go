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
	// The request ID.
	//
	// example:
	//
	// BBCCA90A-A1F0-4B16-B355-692247197805
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The generation status of the user credential report. Valid values:
	//
	// 	- STARTED: The system starts to generate the user credential report.
	//
	// 	- INPROGRESS: The user credential report is being generated.
	//
	// 	- COMPLETED: The user credential report is generated.
	//
	// example:
	//
	// STARTED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
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
