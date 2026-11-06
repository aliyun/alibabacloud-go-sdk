// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemAuditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkItemAuditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkItemAuditResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkItemAuditResponseBody) *GetWorkItemAuditResponse
	GetBody() *GetWorkItemAuditResponseBody
}

type GetWorkItemAuditResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkItemAuditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkItemAuditResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemAuditResponse) GoString() string {
	return s.String()
}

func (s *GetWorkItemAuditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkItemAuditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkItemAuditResponse) GetBody() *GetWorkItemAuditResponseBody {
	return s.Body
}

func (s *GetWorkItemAuditResponse) SetHeaders(v map[string]*string) *GetWorkItemAuditResponse {
	s.Headers = v
	return s
}

func (s *GetWorkItemAuditResponse) SetStatusCode(v int32) *GetWorkItemAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkItemAuditResponse) SetBody(v *GetWorkItemAuditResponseBody) *GetWorkItemAuditResponse {
	s.Body = v
	return s
}

func (s *GetWorkItemAuditResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
