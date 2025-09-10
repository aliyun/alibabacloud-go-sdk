// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQuotaApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateQuotaApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateQuotaApplicationResponse
	GetStatusCode() *int32
	SetBody(v *CreateQuotaApplicationResponseBody) *CreateQuotaApplicationResponse
	GetBody() *CreateQuotaApplicationResponseBody
}

type CreateQuotaApplicationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQuotaApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQuotaApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateQuotaApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateQuotaApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateQuotaApplicationResponse) GetBody() *CreateQuotaApplicationResponseBody {
	return s.Body
}

func (s *CreateQuotaApplicationResponse) SetHeaders(v map[string]*string) *CreateQuotaApplicationResponse {
	s.Headers = v
	return s
}

func (s *CreateQuotaApplicationResponse) SetStatusCode(v int32) *CreateQuotaApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQuotaApplicationResponse) SetBody(v *CreateQuotaApplicationResponseBody) *CreateQuotaApplicationResponse {
	s.Body = v
	return s
}

func (s *CreateQuotaApplicationResponse) Validate() error {
	return dara.Validate(s)
}
