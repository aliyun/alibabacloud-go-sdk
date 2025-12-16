// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQueryProcessorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateQueryProcessorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateQueryProcessorResponse
	GetStatusCode() *int32
	SetBody(v *CreateQueryProcessorResponseBody) *CreateQueryProcessorResponse
	GetBody() *CreateQueryProcessorResponseBody
}

type CreateQueryProcessorResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQueryProcessorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQueryProcessorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateQueryProcessorResponse) GoString() string {
	return s.String()
}

func (s *CreateQueryProcessorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateQueryProcessorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateQueryProcessorResponse) GetBody() *CreateQueryProcessorResponseBody {
	return s.Body
}

func (s *CreateQueryProcessorResponse) SetHeaders(v map[string]*string) *CreateQueryProcessorResponse {
	s.Headers = v
	return s
}

func (s *CreateQueryProcessorResponse) SetStatusCode(v int32) *CreateQueryProcessorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQueryProcessorResponse) SetBody(v *CreateQueryProcessorResponseBody) *CreateQueryProcessorResponse {
	s.Body = v
	return s
}

func (s *CreateQueryProcessorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
