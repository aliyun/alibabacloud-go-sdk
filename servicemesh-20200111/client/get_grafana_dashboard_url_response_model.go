// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGrafanaDashboardUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGrafanaDashboardUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGrafanaDashboardUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetGrafanaDashboardUrlResponseBody) *GetGrafanaDashboardUrlResponse
	GetBody() *GetGrafanaDashboardUrlResponseBody
}

type GetGrafanaDashboardUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGrafanaDashboardUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGrafanaDashboardUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGrafanaDashboardUrlResponse) GoString() string {
	return s.String()
}

func (s *GetGrafanaDashboardUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGrafanaDashboardUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGrafanaDashboardUrlResponse) GetBody() *GetGrafanaDashboardUrlResponseBody {
	return s.Body
}

func (s *GetGrafanaDashboardUrlResponse) SetHeaders(v map[string]*string) *GetGrafanaDashboardUrlResponse {
	s.Headers = v
	return s
}

func (s *GetGrafanaDashboardUrlResponse) SetStatusCode(v int32) *GetGrafanaDashboardUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGrafanaDashboardUrlResponse) SetBody(v *GetGrafanaDashboardUrlResponseBody) *GetGrafanaDashboardUrlResponse {
	s.Body = v
	return s
}

func (s *GetGrafanaDashboardUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
