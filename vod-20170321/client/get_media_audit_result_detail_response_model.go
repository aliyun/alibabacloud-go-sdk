// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaAuditResultDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaAuditResultDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaAuditResultDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaAuditResultDetailResponseBody) *GetMediaAuditResultDetailResponse
	GetBody() *GetMediaAuditResultDetailResponseBody
}

type GetMediaAuditResultDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaAuditResultDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaAuditResultDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultDetailResponse) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaAuditResultDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaAuditResultDetailResponse) GetBody() *GetMediaAuditResultDetailResponseBody {
	return s.Body
}

func (s *GetMediaAuditResultDetailResponse) SetHeaders(v map[string]*string) *GetMediaAuditResultDetailResponse {
	s.Headers = v
	return s
}

func (s *GetMediaAuditResultDetailResponse) SetStatusCode(v int32) *GetMediaAuditResultDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaAuditResultDetailResponse) SetBody(v *GetMediaAuditResultDetailResponseBody) *GetMediaAuditResultDetailResponse {
	s.Body = v
	return s
}

func (s *GetMediaAuditResultDetailResponse) Validate() error {
	return dara.Validate(s)
}
