// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableServiceResponse
	GetStatusCode() *int32
	SetBody(v *DisableServiceResponseBody) *DisableServiceResponse
	GetBody() *DisableServiceResponseBody
}

type DisableServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableServiceResponse) GoString() string {
	return s.String()
}

func (s *DisableServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableServiceResponse) GetBody() *DisableServiceResponseBody {
	return s.Body
}

func (s *DisableServiceResponse) SetHeaders(v map[string]*string) *DisableServiceResponse {
	s.Headers = v
	return s
}

func (s *DisableServiceResponse) SetStatusCode(v int32) *DisableServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableServiceResponse) SetBody(v *DisableServiceResponseBody) *DisableServiceResponse {
	s.Body = v
	return s
}

func (s *DisableServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
