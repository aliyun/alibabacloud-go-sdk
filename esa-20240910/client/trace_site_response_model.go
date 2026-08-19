// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTraceSiteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TraceSiteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TraceSiteResponse
	GetStatusCode() *int32
	SetBody(v *TraceSiteResponseBody) *TraceSiteResponse
	GetBody() *TraceSiteResponseBody
}

type TraceSiteResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TraceSiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TraceSiteResponse) String() string {
	return dara.Prettify(s)
}

func (s TraceSiteResponse) GoString() string {
	return s.String()
}

func (s *TraceSiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TraceSiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TraceSiteResponse) GetBody() *TraceSiteResponseBody {
	return s.Body
}

func (s *TraceSiteResponse) SetHeaders(v map[string]*string) *TraceSiteResponse {
	s.Headers = v
	return s
}

func (s *TraceSiteResponse) SetStatusCode(v int32) *TraceSiteResponse {
	s.StatusCode = &v
	return s
}

func (s *TraceSiteResponse) SetBody(v *TraceSiteResponseBody) *TraceSiteResponse {
	s.Body = v
	return s
}

func (s *TraceSiteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
