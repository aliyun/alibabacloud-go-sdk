// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrialDataChannelCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTrialDataChannelCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTrialDataChannelCredentialResponse
	GetStatusCode() *int32
	SetBody(v *GetTrialDataChannelCredentialResponseBody) *GetTrialDataChannelCredentialResponse
	GetBody() *GetTrialDataChannelCredentialResponseBody
}

type GetTrialDataChannelCredentialResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrialDataChannelCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrialDataChannelCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTrialDataChannelCredentialResponse) GoString() string {
	return s.String()
}

func (s *GetTrialDataChannelCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTrialDataChannelCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTrialDataChannelCredentialResponse) GetBody() *GetTrialDataChannelCredentialResponseBody {
	return s.Body
}

func (s *GetTrialDataChannelCredentialResponse) SetHeaders(v map[string]*string) *GetTrialDataChannelCredentialResponse {
	s.Headers = v
	return s
}

func (s *GetTrialDataChannelCredentialResponse) SetStatusCode(v int32) *GetTrialDataChannelCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrialDataChannelCredentialResponse) SetBody(v *GetTrialDataChannelCredentialResponseBody) *GetTrialDataChannelCredentialResponse {
	s.Body = v
	return s
}

func (s *GetTrialDataChannelCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
