// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSeoTrendsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppSeoTrendsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppSeoTrendsResponse
	GetStatusCode() *int32
	SetBody(v *GetAppSeoTrendsResponseBody) *GetAppSeoTrendsResponse
	GetBody() *GetAppSeoTrendsResponseBody
}

type GetAppSeoTrendsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppSeoTrendsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppSeoTrendsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppSeoTrendsResponse) GoString() string {
	return s.String()
}

func (s *GetAppSeoTrendsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppSeoTrendsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppSeoTrendsResponse) GetBody() *GetAppSeoTrendsResponseBody {
	return s.Body
}

func (s *GetAppSeoTrendsResponse) SetHeaders(v map[string]*string) *GetAppSeoTrendsResponse {
	s.Headers = v
	return s
}

func (s *GetAppSeoTrendsResponse) SetStatusCode(v int32) *GetAppSeoTrendsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppSeoTrendsResponse) SetBody(v *GetAppSeoTrendsResponseBody) *GetAppSeoTrendsResponse {
	s.Body = v
	return s
}

func (s *GetAppSeoTrendsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
