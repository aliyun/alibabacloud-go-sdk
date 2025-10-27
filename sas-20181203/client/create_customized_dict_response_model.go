// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomizedDictResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomizedDictResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomizedDictResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomizedDictResponseBody) *CreateCustomizedDictResponse
	GetBody() *CreateCustomizedDictResponseBody
}

type CreateCustomizedDictResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomizedDictResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomizedDictResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomizedDictResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomizedDictResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomizedDictResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomizedDictResponse) GetBody() *CreateCustomizedDictResponseBody {
	return s.Body
}

func (s *CreateCustomizedDictResponse) SetHeaders(v map[string]*string) *CreateCustomizedDictResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomizedDictResponse) SetStatusCode(v int32) *CreateCustomizedDictResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomizedDictResponse) SetBody(v *CreateCustomizedDictResponseBody) *CreateCustomizedDictResponse {
	s.Body = v
	return s
}

func (s *CreateCustomizedDictResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
