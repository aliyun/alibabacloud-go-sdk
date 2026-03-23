// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenSendifyTrialServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenSendifyTrialServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenSendifyTrialServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenSendifyTrialServiceResponseBody) *OpenSendifyTrialServiceResponse
	GetBody() *OpenSendifyTrialServiceResponseBody
}

type OpenSendifyTrialServiceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenSendifyTrialServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenSendifyTrialServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenSendifyTrialServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenSendifyTrialServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenSendifyTrialServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenSendifyTrialServiceResponse) GetBody() *OpenSendifyTrialServiceResponseBody {
	return s.Body
}

func (s *OpenSendifyTrialServiceResponse) SetHeaders(v map[string]*string) *OpenSendifyTrialServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenSendifyTrialServiceResponse) SetStatusCode(v int32) *OpenSendifyTrialServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenSendifyTrialServiceResponse) SetBody(v *OpenSendifyTrialServiceResponseBody) *OpenSendifyTrialServiceResponse {
	s.Body = v
	return s
}

func (s *OpenSendifyTrialServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
