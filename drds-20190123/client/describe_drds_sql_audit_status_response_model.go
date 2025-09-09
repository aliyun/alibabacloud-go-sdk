// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsSqlAuditStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDrdsSqlAuditStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDrdsSqlAuditStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDrdsSqlAuditStatusResponseBody) *DescribeDrdsSqlAuditStatusResponse
	GetBody() *DescribeDrdsSqlAuditStatusResponseBody
}

type DescribeDrdsSqlAuditStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsSqlAuditStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsSqlAuditStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsSqlAuditStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSqlAuditStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDrdsSqlAuditStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDrdsSqlAuditStatusResponse) GetBody() *DescribeDrdsSqlAuditStatusResponseBody {
	return s.Body
}

func (s *DescribeDrdsSqlAuditStatusResponse) SetHeaders(v map[string]*string) *DescribeDrdsSqlAuditStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponse) SetStatusCode(v int32) *DescribeDrdsSqlAuditStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponse) SetBody(v *DescribeDrdsSqlAuditStatusResponseBody) *DescribeDrdsSqlAuditStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponse) Validate() error {
	return dara.Validate(s)
}
