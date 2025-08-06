// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuditHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuditHistoryResponse
	GetStatusCode() *int32
	SetBody(v *GetAuditHistoryResponseBody) *GetAuditHistoryResponse
	GetBody() *GetAuditHistoryResponseBody
}

type GetAuditHistoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuditHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuditHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuditHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetAuditHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuditHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuditHistoryResponse) GetBody() *GetAuditHistoryResponseBody {
	return s.Body
}

func (s *GetAuditHistoryResponse) SetHeaders(v map[string]*string) *GetAuditHistoryResponse {
	s.Headers = v
	return s
}

func (s *GetAuditHistoryResponse) SetStatusCode(v int32) *GetAuditHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuditHistoryResponse) SetBody(v *GetAuditHistoryResponseBody) *GetAuditHistoryResponse {
	s.Body = v
	return s
}

func (s *GetAuditHistoryResponse) Validate() error {
	return dara.Validate(s)
}
