// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAtiAgentRegisterInfoAcmeChallengeRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse
	GetStatusCode() *int32
	SetBody(v *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse
	GetBody() *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponseBody
}

type CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse) GoString() string {
	return s.String()
}

func (s *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse) GetBody() *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponseBody {
	return s.Body
}

func (s *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse) SetHeaders(v map[string]*string) *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse {
	s.Headers = v
	return s
}

func (s *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse) SetStatusCode(v int32) *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse) SetBody(v *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse {
	s.Body = v
	return s
}

func (s *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
