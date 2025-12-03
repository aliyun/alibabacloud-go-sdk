// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeployKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDeployKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDeployKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreateDeployKeyResponseBody) *CreateDeployKeyResponse
	GetBody() *CreateDeployKeyResponseBody
}

type CreateDeployKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeployKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeployKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDeployKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateDeployKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDeployKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDeployKeyResponse) GetBody() *CreateDeployKeyResponseBody {
	return s.Body
}

func (s *CreateDeployKeyResponse) SetHeaders(v map[string]*string) *CreateDeployKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateDeployKeyResponse) SetStatusCode(v int32) *CreateDeployKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeployKeyResponse) SetBody(v *CreateDeployKeyResponseBody) *CreateDeployKeyResponse {
	s.Body = v
	return s
}

func (s *CreateDeployKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
