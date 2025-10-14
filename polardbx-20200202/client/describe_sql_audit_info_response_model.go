// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlAuditInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSqlAuditInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSqlAuditInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSqlAuditInfoResponseBody) *DescribeSqlAuditInfoResponse
	GetBody() *DescribeSqlAuditInfoResponseBody
}

type DescribeSqlAuditInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSqlAuditInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSqlAuditInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlAuditInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeSqlAuditInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSqlAuditInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSqlAuditInfoResponse) GetBody() *DescribeSqlAuditInfoResponseBody {
	return s.Body
}

func (s *DescribeSqlAuditInfoResponse) SetHeaders(v map[string]*string) *DescribeSqlAuditInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeSqlAuditInfoResponse) SetStatusCode(v int32) *DescribeSqlAuditInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSqlAuditInfoResponse) SetBody(v *DescribeSqlAuditInfoResponseBody) *DescribeSqlAuditInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeSqlAuditInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
