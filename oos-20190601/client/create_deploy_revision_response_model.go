// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeployRevisionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDeployRevisionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDeployRevisionResponse
	GetStatusCode() *int32
	SetBody(v *CreateDeployRevisionResponseBody) *CreateDeployRevisionResponse
	GetBody() *CreateDeployRevisionResponseBody
}

type CreateDeployRevisionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeployRevisionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeployRevisionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDeployRevisionResponse) GoString() string {
	return s.String()
}

func (s *CreateDeployRevisionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDeployRevisionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDeployRevisionResponse) GetBody() *CreateDeployRevisionResponseBody {
	return s.Body
}

func (s *CreateDeployRevisionResponse) SetHeaders(v map[string]*string) *CreateDeployRevisionResponse {
	s.Headers = v
	return s
}

func (s *CreateDeployRevisionResponse) SetStatusCode(v int32) *CreateDeployRevisionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeployRevisionResponse) SetBody(v *CreateDeployRevisionResponseBody) *CreateDeployRevisionResponse {
	s.Body = v
	return s
}

func (s *CreateDeployRevisionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
