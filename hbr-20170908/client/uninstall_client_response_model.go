// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallClientResponse
	GetStatusCode() *int32
	SetBody(v *UninstallClientResponseBody) *UninstallClientResponse
	GetBody() *UninstallClientResponseBody
}

type UninstallClientResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallClientResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallClientResponse) GoString() string {
	return s.String()
}

func (s *UninstallClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallClientResponse) GetBody() *UninstallClientResponseBody {
	return s.Body
}

func (s *UninstallClientResponse) SetHeaders(v map[string]*string) *UninstallClientResponse {
	s.Headers = v
	return s
}

func (s *UninstallClientResponse) SetStatusCode(v int32) *UninstallClientResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallClientResponse) SetBody(v *UninstallClientResponseBody) *UninstallClientResponse {
	s.Body = v
	return s
}

func (s *UninstallClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
