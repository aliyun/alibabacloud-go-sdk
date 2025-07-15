// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaTipResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQuotaTipResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQuotaTipResponse
	GetStatusCode() *int32
	SetBody(v *GetQuotaTipResponseBody) *GetQuotaTipResponse
	GetBody() *GetQuotaTipResponseBody
}

type GetQuotaTipResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaTipResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaTipResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaTipResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaTipResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQuotaTipResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQuotaTipResponse) GetBody() *GetQuotaTipResponseBody {
	return s.Body
}

func (s *GetQuotaTipResponse) SetHeaders(v map[string]*string) *GetQuotaTipResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaTipResponse) SetStatusCode(v int32) *GetQuotaTipResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaTipResponse) SetBody(v *GetQuotaTipResponseBody) *GetQuotaTipResponse {
	s.Body = v
	return s
}

func (s *GetQuotaTipResponse) Validate() error {
	return dara.Validate(s)
}
