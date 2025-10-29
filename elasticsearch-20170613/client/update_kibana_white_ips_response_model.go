// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKibanaWhiteIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKibanaWhiteIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKibanaWhiteIpsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKibanaWhiteIpsResponseBody) *UpdateKibanaWhiteIpsResponse
	GetBody() *UpdateKibanaWhiteIpsResponseBody
}

type UpdateKibanaWhiteIpsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKibanaWhiteIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKibanaWhiteIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKibanaWhiteIpsResponse) GoString() string {
	return s.String()
}

func (s *UpdateKibanaWhiteIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKibanaWhiteIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKibanaWhiteIpsResponse) GetBody() *UpdateKibanaWhiteIpsResponseBody {
	return s.Body
}

func (s *UpdateKibanaWhiteIpsResponse) SetHeaders(v map[string]*string) *UpdateKibanaWhiteIpsResponse {
	s.Headers = v
	return s
}

func (s *UpdateKibanaWhiteIpsResponse) SetStatusCode(v int32) *UpdateKibanaWhiteIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKibanaWhiteIpsResponse) SetBody(v *UpdateKibanaWhiteIpsResponseBody) *UpdateKibanaWhiteIpsResponse {
	s.Body = v
	return s
}

func (s *UpdateKibanaWhiteIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
