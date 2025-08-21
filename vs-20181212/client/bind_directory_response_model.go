// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *BindDirectoryResponseBody) *BindDirectoryResponse
	GetBody() *BindDirectoryResponseBody
}

type BindDirectoryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s BindDirectoryResponse) GoString() string {
	return s.String()
}

func (s *BindDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindDirectoryResponse) GetBody() *BindDirectoryResponseBody {
	return s.Body
}

func (s *BindDirectoryResponse) SetHeaders(v map[string]*string) *BindDirectoryResponse {
	s.Headers = v
	return s
}

func (s *BindDirectoryResponse) SetStatusCode(v int32) *BindDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *BindDirectoryResponse) SetBody(v *BindDirectoryResponseBody) *BindDirectoryResponse {
	s.Body = v
	return s
}

func (s *BindDirectoryResponse) Validate() error {
	return dara.Validate(s)
}
