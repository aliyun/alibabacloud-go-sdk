// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQuotaApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQuotaApplicationResponse
	GetStatusCode() *int32
	SetBody(v *GetQuotaApplicationResponseBody) *GetQuotaApplicationResponse
	GetBody() *GetQuotaApplicationResponseBody
}

type GetQuotaApplicationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQuotaApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQuotaApplicationResponse) GetBody() *GetQuotaApplicationResponseBody {
	return s.Body
}

func (s *GetQuotaApplicationResponse) SetHeaders(v map[string]*string) *GetQuotaApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaApplicationResponse) SetStatusCode(v int32) *GetQuotaApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaApplicationResponse) SetBody(v *GetQuotaApplicationResponseBody) *GetQuotaApplicationResponse {
	s.Body = v
	return s
}

func (s *GetQuotaApplicationResponse) Validate() error {
	return dara.Validate(s)
}
