// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditResultDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuditResultDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuditResultDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetAuditResultDetailResponseBody) *GetAuditResultDetailResponse
	GetBody() *GetAuditResultDetailResponseBody
}

type GetAuditResultDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuditResultDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuditResultDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuditResultDetailResponse) GoString() string {
	return s.String()
}

func (s *GetAuditResultDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuditResultDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuditResultDetailResponse) GetBody() *GetAuditResultDetailResponseBody {
	return s.Body
}

func (s *GetAuditResultDetailResponse) SetHeaders(v map[string]*string) *GetAuditResultDetailResponse {
	s.Headers = v
	return s
}

func (s *GetAuditResultDetailResponse) SetStatusCode(v int32) *GetAuditResultDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuditResultDetailResponse) SetBody(v *GetAuditResultDetailResponseBody) *GetAuditResultDetailResponse {
	s.Body = v
	return s
}

func (s *GetAuditResultDetailResponse) Validate() error {
	return dara.Validate(s)
}
