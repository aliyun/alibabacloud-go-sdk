// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFactAuditUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFactAuditUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFactAuditUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetFactAuditUrlResponseBody) *GetFactAuditUrlResponse
	GetBody() *GetFactAuditUrlResponseBody
}

type GetFactAuditUrlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFactAuditUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFactAuditUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFactAuditUrlResponse) GoString() string {
	return s.String()
}

func (s *GetFactAuditUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFactAuditUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFactAuditUrlResponse) GetBody() *GetFactAuditUrlResponseBody {
	return s.Body
}

func (s *GetFactAuditUrlResponse) SetHeaders(v map[string]*string) *GetFactAuditUrlResponse {
	s.Headers = v
	return s
}

func (s *GetFactAuditUrlResponse) SetStatusCode(v int32) *GetFactAuditUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFactAuditUrlResponse) SetBody(v *GetFactAuditUrlResponseBody) *GetFactAuditUrlResponse {
	s.Body = v
	return s
}

func (s *GetFactAuditUrlResponse) Validate() error {
	return dara.Validate(s)
}
