// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptAgreementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AcceptAgreementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AcceptAgreementResponse
	GetStatusCode() *int32
	SetBody(v *AcceptAgreementResponseBody) *AcceptAgreementResponse
	GetBody() *AcceptAgreementResponseBody
}

type AcceptAgreementResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AcceptAgreementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AcceptAgreementResponse) String() string {
	return dara.Prettify(s)
}

func (s AcceptAgreementResponse) GoString() string {
	return s.String()
}

func (s *AcceptAgreementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AcceptAgreementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AcceptAgreementResponse) GetBody() *AcceptAgreementResponseBody {
	return s.Body
}

func (s *AcceptAgreementResponse) SetHeaders(v map[string]*string) *AcceptAgreementResponse {
	s.Headers = v
	return s
}

func (s *AcceptAgreementResponse) SetStatusCode(v int32) *AcceptAgreementResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptAgreementResponse) SetBody(v *AcceptAgreementResponseBody) *AcceptAgreementResponse {
	s.Body = v
	return s
}

func (s *AcceptAgreementResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
