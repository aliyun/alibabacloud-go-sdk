// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkerBootstrapTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkerBootstrapTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkerBootstrapTokenResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkerBootstrapTokenResponseBody) *CreateWorkerBootstrapTokenResponse
	GetBody() *CreateWorkerBootstrapTokenResponseBody
}

type CreateWorkerBootstrapTokenResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkerBootstrapTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkerBootstrapTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerBootstrapTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkerBootstrapTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkerBootstrapTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkerBootstrapTokenResponse) GetBody() *CreateWorkerBootstrapTokenResponseBody {
	return s.Body
}

func (s *CreateWorkerBootstrapTokenResponse) SetHeaders(v map[string]*string) *CreateWorkerBootstrapTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkerBootstrapTokenResponse) SetStatusCode(v int32) *CreateWorkerBootstrapTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkerBootstrapTokenResponse) SetBody(v *CreateWorkerBootstrapTokenResponseBody) *CreateWorkerBootstrapTokenResponse {
	s.Body = v
	return s
}

func (s *CreateWorkerBootstrapTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
