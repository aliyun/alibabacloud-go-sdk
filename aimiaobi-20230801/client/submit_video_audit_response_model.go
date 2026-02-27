// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoAuditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitVideoAuditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitVideoAuditResponse
	GetStatusCode() *int32
	SetBody(v *SubmitVideoAuditResponseBody) *SubmitVideoAuditResponse
	GetBody() *SubmitVideoAuditResponseBody
}

type SubmitVideoAuditResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitVideoAuditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitVideoAuditResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoAuditResponse) GoString() string {
	return s.String()
}

func (s *SubmitVideoAuditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitVideoAuditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitVideoAuditResponse) GetBody() *SubmitVideoAuditResponseBody {
	return s.Body
}

func (s *SubmitVideoAuditResponse) SetHeaders(v map[string]*string) *SubmitVideoAuditResponse {
	s.Headers = v
	return s
}

func (s *SubmitVideoAuditResponse) SetStatusCode(v int32) *SubmitVideoAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitVideoAuditResponse) SetBody(v *SubmitVideoAuditResponseBody) *SubmitVideoAuditResponse {
	s.Body = v
	return s
}

func (s *SubmitVideoAuditResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
