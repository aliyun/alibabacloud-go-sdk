// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse
	GetStatusCode() *int32
	SetBody(v *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse
	GetBody() *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody
}

type VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse) GoString() string {
	return s.String()
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse) GetBody() *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody {
	return s.Body
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse) SetHeaders(v map[string]*string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse {
	s.Headers = v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse) SetStatusCode(v int32) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse) SetBody(v *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse {
	s.Body = v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
