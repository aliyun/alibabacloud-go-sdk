// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyFormalServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyFormalServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyFormalServiceResponse
	GetStatusCode() *int32
	SetBody(v *ApplyFormalServiceResponseBody) *ApplyFormalServiceResponse
	GetBody() *ApplyFormalServiceResponseBody
}

type ApplyFormalServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyFormalServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyFormalServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyFormalServiceResponse) GoString() string {
	return s.String()
}

func (s *ApplyFormalServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyFormalServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyFormalServiceResponse) GetBody() *ApplyFormalServiceResponseBody {
	return s.Body
}

func (s *ApplyFormalServiceResponse) SetHeaders(v map[string]*string) *ApplyFormalServiceResponse {
	s.Headers = v
	return s
}

func (s *ApplyFormalServiceResponse) SetStatusCode(v int32) *ApplyFormalServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyFormalServiceResponse) SetBody(v *ApplyFormalServiceResponseBody) *ApplyFormalServiceResponse {
	s.Body = v
	return s
}

func (s *ApplyFormalServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
