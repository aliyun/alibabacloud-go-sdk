// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitFactAuditUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitFactAuditUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitFactAuditUrlResponse
	GetStatusCode() *int32
	SetBody(v *SubmitFactAuditUrlResponseBody) *SubmitFactAuditUrlResponse
	GetBody() *SubmitFactAuditUrlResponseBody
}

type SubmitFactAuditUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitFactAuditUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitFactAuditUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitFactAuditUrlResponse) GoString() string {
	return s.String()
}

func (s *SubmitFactAuditUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitFactAuditUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitFactAuditUrlResponse) GetBody() *SubmitFactAuditUrlResponseBody {
	return s.Body
}

func (s *SubmitFactAuditUrlResponse) SetHeaders(v map[string]*string) *SubmitFactAuditUrlResponse {
	s.Headers = v
	return s
}

func (s *SubmitFactAuditUrlResponse) SetStatusCode(v int32) *SubmitFactAuditUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitFactAuditUrlResponse) SetBody(v *SubmitFactAuditUrlResponseBody) *SubmitFactAuditUrlResponse {
	s.Body = v
	return s
}

func (s *SubmitFactAuditUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
