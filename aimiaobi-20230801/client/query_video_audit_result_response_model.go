// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVideoAuditResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryVideoAuditResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryVideoAuditResultResponse
	GetStatusCode() *int32
	SetBody(v *QueryVideoAuditResultResponseBody) *QueryVideoAuditResultResponse
	GetBody() *QueryVideoAuditResultResponseBody
}

type QueryVideoAuditResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVideoAuditResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVideoAuditResultResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryVideoAuditResultResponse) GoString() string {
	return s.String()
}

func (s *QueryVideoAuditResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryVideoAuditResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryVideoAuditResultResponse) GetBody() *QueryVideoAuditResultResponseBody {
	return s.Body
}

func (s *QueryVideoAuditResultResponse) SetHeaders(v map[string]*string) *QueryVideoAuditResultResponse {
	s.Headers = v
	return s
}

func (s *QueryVideoAuditResultResponse) SetStatusCode(v int32) *QueryVideoAuditResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVideoAuditResultResponse) SetBody(v *QueryVideoAuditResultResponseBody) *QueryVideoAuditResultResponse {
	s.Body = v
	return s
}

func (s *QueryVideoAuditResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
