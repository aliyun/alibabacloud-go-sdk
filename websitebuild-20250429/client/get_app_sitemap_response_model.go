// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSitemapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppSitemapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppSitemapResponse
	GetStatusCode() *int32
	SetBody(v *GetAppSitemapResponseBody) *GetAppSitemapResponse
	GetBody() *GetAppSitemapResponseBody
}

type GetAppSitemapResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppSitemapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppSitemapResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppSitemapResponse) GoString() string {
	return s.String()
}

func (s *GetAppSitemapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppSitemapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppSitemapResponse) GetBody() *GetAppSitemapResponseBody {
	return s.Body
}

func (s *GetAppSitemapResponse) SetHeaders(v map[string]*string) *GetAppSitemapResponse {
	s.Headers = v
	return s
}

func (s *GetAppSitemapResponse) SetStatusCode(v int32) *GetAppSitemapResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppSitemapResponse) SetBody(v *GetAppSitemapResponseBody) *GetAppSitemapResponse {
	s.Body = v
	return s
}

func (s *GetAppSitemapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
