// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDentryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyDentryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyDentryResponse
	GetStatusCode() *int32
	SetBody(v *CopyDentryResponseBody) *CopyDentryResponse
	GetBody() *CopyDentryResponseBody
}

type CopyDentryResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyDentryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyDentryResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryResponse) GoString() string {
	return s.String()
}

func (s *CopyDentryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyDentryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyDentryResponse) GetBody() *CopyDentryResponseBody {
	return s.Body
}

func (s *CopyDentryResponse) SetHeaders(v map[string]*string) *CopyDentryResponse {
	s.Headers = v
	return s
}

func (s *CopyDentryResponse) SetStatusCode(v int32) *CopyDentryResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyDentryResponse) SetBody(v *CopyDentryResponseBody) *CopyDentryResponse {
	s.Body = v
	return s
}

func (s *CopyDentryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
