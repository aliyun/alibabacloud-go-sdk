// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditCountDistributionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuditCountDistributionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuditCountDistributionResponse
	GetStatusCode() *int32
	SetBody(v *GetAuditCountDistributionResponseBody) *GetAuditCountDistributionResponse
	GetBody() *GetAuditCountDistributionResponseBody
}

type GetAuditCountDistributionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuditCountDistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuditCountDistributionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuditCountDistributionResponse) GoString() string {
	return s.String()
}

func (s *GetAuditCountDistributionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuditCountDistributionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuditCountDistributionResponse) GetBody() *GetAuditCountDistributionResponseBody {
	return s.Body
}

func (s *GetAuditCountDistributionResponse) SetHeaders(v map[string]*string) *GetAuditCountDistributionResponse {
	s.Headers = v
	return s
}

func (s *GetAuditCountDistributionResponse) SetStatusCode(v int32) *GetAuditCountDistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuditCountDistributionResponse) SetBody(v *GetAuditCountDistributionResponseBody) *GetAuditCountDistributionResponse {
	s.Body = v
	return s
}

func (s *GetAuditCountDistributionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
