// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteQuotaResponse
	GetStatusCode() *int32
	SetBody(v *DeleteQuotaResponseBody) *DeleteQuotaResponse
	GetBody() *DeleteQuotaResponseBody
}

type DeleteQuotaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteQuotaResponse) GoString() string {
	return s.String()
}

func (s *DeleteQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteQuotaResponse) GetBody() *DeleteQuotaResponseBody {
	return s.Body
}

func (s *DeleteQuotaResponse) SetHeaders(v map[string]*string) *DeleteQuotaResponse {
	s.Headers = v
	return s
}

func (s *DeleteQuotaResponse) SetStatusCode(v int32) *DeleteQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQuotaResponse) SetBody(v *DeleteQuotaResponseBody) *DeleteQuotaResponse {
	s.Body = v
	return s
}

func (s *DeleteQuotaResponse) Validate() error {
	return dara.Validate(s)
}
