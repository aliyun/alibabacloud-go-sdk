// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNetworkAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNetworkAccessResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNetworkAccessResponseBody) *UpdateNetworkAccessResponse
	GetBody() *UpdateNetworkAccessResponseBody
}

type UpdateNetworkAccessResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNetworkAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNetworkAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkAccessResponse) GoString() string {
	return s.String()
}

func (s *UpdateNetworkAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNetworkAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNetworkAccessResponse) GetBody() *UpdateNetworkAccessResponseBody {
	return s.Body
}

func (s *UpdateNetworkAccessResponse) SetHeaders(v map[string]*string) *UpdateNetworkAccessResponse {
	s.Headers = v
	return s
}

func (s *UpdateNetworkAccessResponse) SetStatusCode(v int32) *UpdateNetworkAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNetworkAccessResponse) SetBody(v *UpdateNetworkAccessResponseBody) *UpdateNetworkAccessResponse {
	s.Body = v
	return s
}

func (s *UpdateNetworkAccessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
