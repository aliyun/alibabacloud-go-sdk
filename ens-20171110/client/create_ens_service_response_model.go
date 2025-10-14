// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnsServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEnsServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEnsServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateEnsServiceResponseBody) *CreateEnsServiceResponse
	GetBody() *CreateEnsServiceResponseBody
}

type CreateEnsServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEnsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEnsServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEnsServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateEnsServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEnsServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEnsServiceResponse) GetBody() *CreateEnsServiceResponseBody {
	return s.Body
}

func (s *CreateEnsServiceResponse) SetHeaders(v map[string]*string) *CreateEnsServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateEnsServiceResponse) SetStatusCode(v int32) *CreateEnsServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEnsServiceResponse) SetBody(v *CreateEnsServiceResponseBody) *CreateEnsServiceResponse {
	s.Body = v
	return s
}

func (s *CreateEnsServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
