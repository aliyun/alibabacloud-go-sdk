// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIMediaAuditJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAIMediaAuditJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAIMediaAuditJobResponse
	GetStatusCode() *int32
	SetBody(v *GetAIMediaAuditJobResponseBody) *GetAIMediaAuditJobResponse
	GetBody() *GetAIMediaAuditJobResponseBody
}

type GetAIMediaAuditJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAIMediaAuditJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAIMediaAuditJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponse) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAIMediaAuditJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAIMediaAuditJobResponse) GetBody() *GetAIMediaAuditJobResponseBody {
	return s.Body
}

func (s *GetAIMediaAuditJobResponse) SetHeaders(v map[string]*string) *GetAIMediaAuditJobResponse {
	s.Headers = v
	return s
}

func (s *GetAIMediaAuditJobResponse) SetStatusCode(v int32) *GetAIMediaAuditJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAIMediaAuditJobResponse) SetBody(v *GetAIMediaAuditJobResponseBody) *GetAIMediaAuditJobResponse {
	s.Body = v
	return s
}

func (s *GetAIMediaAuditJobResponse) Validate() error {
	return dara.Validate(s)
}
