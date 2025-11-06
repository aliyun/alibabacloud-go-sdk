// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportServicesResponse
	GetStatusCode() *int32
	SetBody(v *ImportServicesResponseBody) *ImportServicesResponse
	GetBody() *ImportServicesResponseBody
}

type ImportServicesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportServicesResponse) GoString() string {
	return s.String()
}

func (s *ImportServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportServicesResponse) GetBody() *ImportServicesResponseBody {
	return s.Body
}

func (s *ImportServicesResponse) SetHeaders(v map[string]*string) *ImportServicesResponse {
	s.Headers = v
	return s
}

func (s *ImportServicesResponse) SetStatusCode(v int32) *ImportServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportServicesResponse) SetBody(v *ImportServicesResponseBody) *ImportServicesResponse {
	s.Body = v
	return s
}

func (s *ImportServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
