// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileResponse
	GetStatusCode() *int32
	SetBody(v *GetFileResponseBody) *GetFileResponse
	GetBody() *GetFileResponseBody
}

type GetFileResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileResponse) GoString() string {
	return s.String()
}

func (s *GetFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileResponse) GetBody() *GetFileResponseBody {
	return s.Body
}

func (s *GetFileResponse) SetHeaders(v map[string]*string) *GetFileResponse {
	s.Headers = v
	return s
}

func (s *GetFileResponse) SetStatusCode(v int32) *GetFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileResponse) SetBody(v *GetFileResponseBody) *GetFileResponse {
	s.Body = v
	return s
}

func (s *GetFileResponse) Validate() error {
	return dara.Validate(s)
}
