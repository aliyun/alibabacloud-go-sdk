// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParameterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateParameterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateParameterResponse
	GetStatusCode() *int32
	SetBody(v *CreateParameterResponseBody) *CreateParameterResponse
	GetBody() *CreateParameterResponseBody
}

type CreateParameterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateParameterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateParameterResponse) GoString() string {
	return s.String()
}

func (s *CreateParameterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateParameterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateParameterResponse) GetBody() *CreateParameterResponseBody {
	return s.Body
}

func (s *CreateParameterResponse) SetHeaders(v map[string]*string) *CreateParameterResponse {
	s.Headers = v
	return s
}

func (s *CreateParameterResponse) SetStatusCode(v int32) *CreateParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateParameterResponse) SetBody(v *CreateParameterResponseBody) *CreateParameterResponse {
	s.Body = v
	return s
}

func (s *CreateParameterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
