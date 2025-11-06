// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOperationAuditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelOperationAuditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelOperationAuditResponse
	GetStatusCode() *int32
	SetBody(v *CancelOperationAuditResponseBody) *CancelOperationAuditResponse
	GetBody() *CancelOperationAuditResponseBody
}

type CancelOperationAuditResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelOperationAuditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelOperationAuditResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelOperationAuditResponse) GoString() string {
	return s.String()
}

func (s *CancelOperationAuditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelOperationAuditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelOperationAuditResponse) GetBody() *CancelOperationAuditResponseBody {
	return s.Body
}

func (s *CancelOperationAuditResponse) SetHeaders(v map[string]*string) *CancelOperationAuditResponse {
	s.Headers = v
	return s
}

func (s *CancelOperationAuditResponse) SetStatusCode(v int32) *CancelOperationAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelOperationAuditResponse) SetBody(v *CancelOperationAuditResponseBody) *CancelOperationAuditResponse {
	s.Body = v
	return s
}

func (s *CancelOperationAuditResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
