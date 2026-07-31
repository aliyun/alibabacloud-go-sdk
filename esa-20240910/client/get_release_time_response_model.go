// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReleaseTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetReleaseTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetReleaseTimeResponse
	GetStatusCode() *int32
	SetBody(v *GetReleaseTimeResponseBody) *GetReleaseTimeResponse
	GetBody() *GetReleaseTimeResponseBody
}

type GetReleaseTimeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReleaseTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReleaseTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetReleaseTimeResponse) GoString() string {
	return s.String()
}

func (s *GetReleaseTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetReleaseTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetReleaseTimeResponse) GetBody() *GetReleaseTimeResponseBody {
	return s.Body
}

func (s *GetReleaseTimeResponse) SetHeaders(v map[string]*string) *GetReleaseTimeResponse {
	s.Headers = v
	return s
}

func (s *GetReleaseTimeResponse) SetStatusCode(v int32) *GetReleaseTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReleaseTimeResponse) SetBody(v *GetReleaseTimeResponseBody) *GetReleaseTimeResponse {
	s.Body = v
	return s
}

func (s *GetReleaseTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
