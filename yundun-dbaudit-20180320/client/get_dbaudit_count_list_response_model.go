// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDBAuditCountListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDBAuditCountListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDBAuditCountListResponse
	GetStatusCode() *int32
	SetBody(v *GetDBAuditCountListResponseBody) *GetDBAuditCountListResponse
	GetBody() *GetDBAuditCountListResponseBody
}

type GetDBAuditCountListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDBAuditCountListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDBAuditCountListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDBAuditCountListResponse) GoString() string {
	return s.String()
}

func (s *GetDBAuditCountListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDBAuditCountListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDBAuditCountListResponse) GetBody() *GetDBAuditCountListResponseBody {
	return s.Body
}

func (s *GetDBAuditCountListResponse) SetHeaders(v map[string]*string) *GetDBAuditCountListResponse {
	s.Headers = v
	return s
}

func (s *GetDBAuditCountListResponse) SetStatusCode(v int32) *GetDBAuditCountListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDBAuditCountListResponse) SetBody(v *GetDBAuditCountListResponseBody) *GetDBAuditCountListResponse {
	s.Body = v
	return s
}

func (s *GetDBAuditCountListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
