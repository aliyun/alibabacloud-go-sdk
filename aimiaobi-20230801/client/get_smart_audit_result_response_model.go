// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmartAuditResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSmartAuditResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSmartAuditResultResponse
	GetStatusCode() *int32
	SetBody(v *GetSmartAuditResultResponseBody) *GetSmartAuditResultResponse
	GetBody() *GetSmartAuditResultResponseBody
}

type GetSmartAuditResultResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSmartAuditResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSmartAuditResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSmartAuditResultResponse) GoString() string {
	return s.String()
}

func (s *GetSmartAuditResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSmartAuditResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSmartAuditResultResponse) GetBody() *GetSmartAuditResultResponseBody {
	return s.Body
}

func (s *GetSmartAuditResultResponse) SetHeaders(v map[string]*string) *GetSmartAuditResultResponse {
	s.Headers = v
	return s
}

func (s *GetSmartAuditResultResponse) SetStatusCode(v int32) *GetSmartAuditResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSmartAuditResultResponse) SetBody(v *GetSmartAuditResultResponseBody) *GetSmartAuditResultResponse {
	s.Body = v
	return s
}

func (s *GetSmartAuditResultResponse) Validate() error {
	return dara.Validate(s)
}
