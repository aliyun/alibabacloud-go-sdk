// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmartAuditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSmartAuditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSmartAuditResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSmartAuditResponseBody) *SubmitSmartAuditResponse
	GetBody() *SubmitSmartAuditResponseBody
}

type SubmitSmartAuditResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSmartAuditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSmartAuditResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartAuditResponse) GoString() string {
	return s.String()
}

func (s *SubmitSmartAuditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSmartAuditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSmartAuditResponse) GetBody() *SubmitSmartAuditResponseBody {
	return s.Body
}

func (s *SubmitSmartAuditResponse) SetHeaders(v map[string]*string) *SubmitSmartAuditResponse {
	s.Headers = v
	return s
}

func (s *SubmitSmartAuditResponse) SetStatusCode(v int32) *SubmitSmartAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSmartAuditResponse) SetBody(v *SubmitSmartAuditResponseBody) *SubmitSmartAuditResponse {
	s.Body = v
	return s
}

func (s *SubmitSmartAuditResponse) Validate() error {
	return dara.Validate(s)
}
