// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuditCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuditCountResponse
	GetStatusCode() *int32
	SetBody(v *GetAuditCountResponseBody) *GetAuditCountResponse
	GetBody() *GetAuditCountResponseBody
}

type GetAuditCountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuditCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuditCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuditCountResponse) GoString() string {
	return s.String()
}

func (s *GetAuditCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuditCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuditCountResponse) GetBody() *GetAuditCountResponseBody {
	return s.Body
}

func (s *GetAuditCountResponse) SetHeaders(v map[string]*string) *GetAuditCountResponse {
	s.Headers = v
	return s
}

func (s *GetAuditCountResponse) SetStatusCode(v int32) *GetAuditCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuditCountResponse) SetBody(v *GetAuditCountResponseBody) *GetAuditCountResponse {
	s.Body = v
	return s
}

func (s *GetAuditCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
