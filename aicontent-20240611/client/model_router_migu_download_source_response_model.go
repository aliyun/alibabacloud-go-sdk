// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterMiguDownloadSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterMiguDownloadSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterMiguDownloadSourceResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterMiguDownloadSourceResponseBody) *ModelRouterMiguDownloadSourceResponse
	GetBody() *ModelRouterMiguDownloadSourceResponseBody
}

type ModelRouterMiguDownloadSourceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterMiguDownloadSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterMiguDownloadSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterMiguDownloadSourceResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterMiguDownloadSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterMiguDownloadSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterMiguDownloadSourceResponse) GetBody() *ModelRouterMiguDownloadSourceResponseBody {
	return s.Body
}

func (s *ModelRouterMiguDownloadSourceResponse) SetHeaders(v map[string]*string) *ModelRouterMiguDownloadSourceResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterMiguDownloadSourceResponse) SetStatusCode(v int32) *ModelRouterMiguDownloadSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterMiguDownloadSourceResponse) SetBody(v *ModelRouterMiguDownloadSourceResponseBody) *ModelRouterMiguDownloadSourceResponse {
	s.Body = v
	return s
}

func (s *ModelRouterMiguDownloadSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
