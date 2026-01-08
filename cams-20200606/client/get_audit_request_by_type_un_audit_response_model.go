// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditRequestByTypeUnAuditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuditRequestByTypeUnAuditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuditRequestByTypeUnAuditResponse
	GetStatusCode() *int32
	SetBody(v *GetAuditRequestByTypeUnAuditResponseBody) *GetAuditRequestByTypeUnAuditResponse
	GetBody() *GetAuditRequestByTypeUnAuditResponseBody
}

type GetAuditRequestByTypeUnAuditResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuditRequestByTypeUnAuditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuditRequestByTypeUnAuditResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuditRequestByTypeUnAuditResponse) GoString() string {
	return s.String()
}

func (s *GetAuditRequestByTypeUnAuditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuditRequestByTypeUnAuditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuditRequestByTypeUnAuditResponse) GetBody() *GetAuditRequestByTypeUnAuditResponseBody {
	return s.Body
}

func (s *GetAuditRequestByTypeUnAuditResponse) SetHeaders(v map[string]*string) *GetAuditRequestByTypeUnAuditResponse {
	s.Headers = v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponse) SetStatusCode(v int32) *GetAuditRequestByTypeUnAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponse) SetBody(v *GetAuditRequestByTypeUnAuditResponseBody) *GetAuditRequestByTypeUnAuditResponse {
	s.Body = v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
