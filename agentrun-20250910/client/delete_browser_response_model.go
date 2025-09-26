// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBrowserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBrowserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBrowserResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBrowserResult) *DeleteBrowserResponse
	GetBody() *DeleteBrowserResult
}

type DeleteBrowserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBrowserResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBrowserResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBrowserResponse) GoString() string {
	return s.String()
}

func (s *DeleteBrowserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBrowserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBrowserResponse) GetBody() *DeleteBrowserResult {
	return s.Body
}

func (s *DeleteBrowserResponse) SetHeaders(v map[string]*string) *DeleteBrowserResponse {
	s.Headers = v
	return s
}

func (s *DeleteBrowserResponse) SetStatusCode(v int32) *DeleteBrowserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBrowserResponse) SetBody(v *DeleteBrowserResult) *DeleteBrowserResponse {
	s.Body = v
	return s
}

func (s *DeleteBrowserResponse) Validate() error {
	return dara.Validate(s)
}
