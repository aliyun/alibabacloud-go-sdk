// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReActivateAuditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReActivateAuditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReActivateAuditResponse
	GetStatusCode() *int32
	SetBody(v *ReActivateAuditResponseBody) *ReActivateAuditResponse
	GetBody() *ReActivateAuditResponseBody
}

type ReActivateAuditResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReActivateAuditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReActivateAuditResponse) String() string {
	return dara.Prettify(s)
}

func (s ReActivateAuditResponse) GoString() string {
	return s.String()
}

func (s *ReActivateAuditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReActivateAuditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReActivateAuditResponse) GetBody() *ReActivateAuditResponseBody {
	return s.Body
}

func (s *ReActivateAuditResponse) SetHeaders(v map[string]*string) *ReActivateAuditResponse {
	s.Headers = v
	return s
}

func (s *ReActivateAuditResponse) SetStatusCode(v int32) *ReActivateAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *ReActivateAuditResponse) SetBody(v *ReActivateAuditResponseBody) *ReActivateAuditResponse {
	s.Body = v
	return s
}

func (s *ReActivateAuditResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
