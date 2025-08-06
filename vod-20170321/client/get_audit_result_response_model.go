// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuditResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuditResultResponse
	GetStatusCode() *int32
	SetBody(v *GetAuditResultResponseBody) *GetAuditResultResponse
	GetBody() *GetAuditResultResponseBody
}

type GetAuditResultResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuditResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuditResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuditResultResponse) GoString() string {
	return s.String()
}

func (s *GetAuditResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuditResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuditResultResponse) GetBody() *GetAuditResultResponseBody {
	return s.Body
}

func (s *GetAuditResultResponse) SetHeaders(v map[string]*string) *GetAuditResultResponse {
	s.Headers = v
	return s
}

func (s *GetAuditResultResponse) SetStatusCode(v int32) *GetAuditResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuditResultResponse) SetBody(v *GetAuditResultResponseBody) *GetAuditResultResponse {
	s.Body = v
	return s
}

func (s *GetAuditResultResponse) Validate() error {
	return dara.Validate(s)
}
