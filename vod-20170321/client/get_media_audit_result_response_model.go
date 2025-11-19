// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaAuditResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaAuditResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaAuditResultResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaAuditResultResponseBody) *GetMediaAuditResultResponse
	GetBody() *GetMediaAuditResultResponseBody
}

type GetMediaAuditResultResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaAuditResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaAuditResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponse) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaAuditResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaAuditResultResponse) GetBody() *GetMediaAuditResultResponseBody {
	return s.Body
}

func (s *GetMediaAuditResultResponse) SetHeaders(v map[string]*string) *GetMediaAuditResultResponse {
	s.Headers = v
	return s
}

func (s *GetMediaAuditResultResponse) SetStatusCode(v int32) *GetMediaAuditResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaAuditResultResponse) SetBody(v *GetMediaAuditResultResponseBody) *GetMediaAuditResultResponse {
	s.Body = v
	return s
}

func (s *GetMediaAuditResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
