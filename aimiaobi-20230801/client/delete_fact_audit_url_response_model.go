// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFactAuditUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFactAuditUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFactAuditUrlResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFactAuditUrlResponseBody) *DeleteFactAuditUrlResponse
	GetBody() *DeleteFactAuditUrlResponseBody
}

type DeleteFactAuditUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFactAuditUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFactAuditUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFactAuditUrlResponse) GoString() string {
	return s.String()
}

func (s *DeleteFactAuditUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFactAuditUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFactAuditUrlResponse) GetBody() *DeleteFactAuditUrlResponseBody {
	return s.Body
}

func (s *DeleteFactAuditUrlResponse) SetHeaders(v map[string]*string) *DeleteFactAuditUrlResponse {
	s.Headers = v
	return s
}

func (s *DeleteFactAuditUrlResponse) SetStatusCode(v int32) *DeleteFactAuditUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFactAuditUrlResponse) SetBody(v *DeleteFactAuditUrlResponseBody) *DeleteFactAuditUrlResponse {
	s.Body = v
	return s
}

func (s *DeleteFactAuditUrlResponse) Validate() error {
	return dara.Validate(s)
}
