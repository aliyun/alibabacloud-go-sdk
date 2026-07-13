// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateModelProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateModelProviderResponse
	GetStatusCode() *int32
	SetBody(v *CreateModelProviderResponseBody) *CreateModelProviderResponse
	GetBody() *CreateModelProviderResponseBody
}

type CreateModelProviderResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateModelProviderResponse) GoString() string {
	return s.String()
}

func (s *CreateModelProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateModelProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateModelProviderResponse) GetBody() *CreateModelProviderResponseBody {
	return s.Body
}

func (s *CreateModelProviderResponse) SetHeaders(v map[string]*string) *CreateModelProviderResponse {
	s.Headers = v
	return s
}

func (s *CreateModelProviderResponse) SetStatusCode(v int32) *CreateModelProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelProviderResponse) SetBody(v *CreateModelProviderResponseBody) *CreateModelProviderResponse {
	s.Body = v
	return s
}

func (s *CreateModelProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
