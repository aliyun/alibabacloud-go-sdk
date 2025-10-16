// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBrowserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBrowserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBrowserResponse
	GetStatusCode() *int32
	SetBody(v *BrowserResult) *CreateBrowserResponse
	GetBody() *BrowserResult
}

type CreateBrowserResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BrowserResult     `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBrowserResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserResponse) GoString() string {
	return s.String()
}

func (s *CreateBrowserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBrowserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBrowserResponse) GetBody() *BrowserResult {
	return s.Body
}

func (s *CreateBrowserResponse) SetHeaders(v map[string]*string) *CreateBrowserResponse {
	s.Headers = v
	return s
}

func (s *CreateBrowserResponse) SetStatusCode(v int32) *CreateBrowserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBrowserResponse) SetBody(v *BrowserResult) *CreateBrowserResponse {
	s.Body = v
	return s
}

func (s *CreateBrowserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
