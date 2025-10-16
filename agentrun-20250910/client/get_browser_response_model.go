// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBrowserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBrowserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBrowserResponse
	GetStatusCode() *int32
	SetBody(v *BrowserResult) *GetBrowserResponse
	GetBody() *BrowserResult
}

type GetBrowserResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BrowserResult     `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBrowserResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserResponse) GoString() string {
	return s.String()
}

func (s *GetBrowserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBrowserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBrowserResponse) GetBody() *BrowserResult {
	return s.Body
}

func (s *GetBrowserResponse) SetHeaders(v map[string]*string) *GetBrowserResponse {
	s.Headers = v
	return s
}

func (s *GetBrowserResponse) SetStatusCode(v int32) *GetBrowserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBrowserResponse) SetBody(v *BrowserResult) *GetBrowserResponse {
	s.Body = v
	return s
}

func (s *GetBrowserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
