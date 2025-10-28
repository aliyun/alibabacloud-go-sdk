// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUrlAsyncModerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UrlAsyncModerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UrlAsyncModerationResponse
	GetStatusCode() *int32
	SetBody(v *UrlAsyncModerationResponseBody) *UrlAsyncModerationResponse
	GetBody() *UrlAsyncModerationResponseBody
}

type UrlAsyncModerationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UrlAsyncModerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UrlAsyncModerationResponse) String() string {
	return dara.Prettify(s)
}

func (s UrlAsyncModerationResponse) GoString() string {
	return s.String()
}

func (s *UrlAsyncModerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UrlAsyncModerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UrlAsyncModerationResponse) GetBody() *UrlAsyncModerationResponseBody {
	return s.Body
}

func (s *UrlAsyncModerationResponse) SetHeaders(v map[string]*string) *UrlAsyncModerationResponse {
	s.Headers = v
	return s
}

func (s *UrlAsyncModerationResponse) SetStatusCode(v int32) *UrlAsyncModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *UrlAsyncModerationResponse) SetBody(v *UrlAsyncModerationResponseBody) *UrlAsyncModerationResponse {
	s.Body = v
	return s
}

func (s *UrlAsyncModerationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
