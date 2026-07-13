// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateModelProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateModelProviderResponse
	GetStatusCode() *int32
	SetBody(v *UpdateModelProviderResponseBody) *UpdateModelProviderResponse
	GetBody() *UpdateModelProviderResponseBody
}

type UpdateModelProviderResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModelProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelProviderResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateModelProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateModelProviderResponse) GetBody() *UpdateModelProviderResponseBody {
	return s.Body
}

func (s *UpdateModelProviderResponse) SetHeaders(v map[string]*string) *UpdateModelProviderResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelProviderResponse) SetStatusCode(v int32) *UpdateModelProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelProviderResponse) SetBody(v *UpdateModelProviderResponseBody) *UpdateModelProviderResponse {
	s.Body = v
	return s
}

func (s *UpdateModelProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
