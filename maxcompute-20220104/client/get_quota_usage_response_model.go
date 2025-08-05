// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQuotaUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQuotaUsageResponse
	GetStatusCode() *int32
	SetBody(v *GetQuotaUsageResponseBody) *GetQuotaUsageResponse
	GetBody() *GetQuotaUsageResponseBody
}

type GetQuotaUsageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaUsageResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQuotaUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQuotaUsageResponse) GetBody() *GetQuotaUsageResponseBody {
	return s.Body
}

func (s *GetQuotaUsageResponse) SetHeaders(v map[string]*string) *GetQuotaUsageResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaUsageResponse) SetStatusCode(v int32) *GetQuotaUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaUsageResponse) SetBody(v *GetQuotaUsageResponseBody) *GetQuotaUsageResponse {
	s.Body = v
	return s
}

func (s *GetQuotaUsageResponse) Validate() error {
	return dara.Validate(s)
}
