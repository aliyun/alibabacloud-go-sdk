// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpenUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOpenUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOpenUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetOpenUrlResponseBody) *GetOpenUrlResponse
	GetBody() *GetOpenUrlResponseBody
}

type GetOpenUrlResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpenUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpenUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOpenUrlResponse) GoString() string {
	return s.String()
}

func (s *GetOpenUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOpenUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOpenUrlResponse) GetBody() *GetOpenUrlResponseBody {
	return s.Body
}

func (s *GetOpenUrlResponse) SetHeaders(v map[string]*string) *GetOpenUrlResponse {
	s.Headers = v
	return s
}

func (s *GetOpenUrlResponse) SetStatusCode(v int32) *GetOpenUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpenUrlResponse) SetBody(v *GetOpenUrlResponseBody) *GetOpenUrlResponse {
	s.Body = v
	return s
}

func (s *GetOpenUrlResponse) Validate() error {
	return dara.Validate(s)
}
