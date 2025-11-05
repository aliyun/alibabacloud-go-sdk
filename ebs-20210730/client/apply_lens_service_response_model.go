// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyLensServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyLensServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyLensServiceResponse
	GetStatusCode() *int32
	SetBody(v *ApplyLensServiceResponseBody) *ApplyLensServiceResponse
	GetBody() *ApplyLensServiceResponseBody
}

type ApplyLensServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyLensServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyLensServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyLensServiceResponse) GoString() string {
	return s.String()
}

func (s *ApplyLensServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyLensServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyLensServiceResponse) GetBody() *ApplyLensServiceResponseBody {
	return s.Body
}

func (s *ApplyLensServiceResponse) SetHeaders(v map[string]*string) *ApplyLensServiceResponse {
	s.Headers = v
	return s
}

func (s *ApplyLensServiceResponse) SetStatusCode(v int32) *ApplyLensServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyLensServiceResponse) SetBody(v *ApplyLensServiceResponseBody) *ApplyLensServiceResponse {
	s.Body = v
	return s
}

func (s *ApplyLensServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
