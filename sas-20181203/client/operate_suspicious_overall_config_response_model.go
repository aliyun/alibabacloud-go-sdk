// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateSuspiciousOverallConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateSuspiciousOverallConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateSuspiciousOverallConfigResponse
	GetStatusCode() *int32
	SetBody(v *OperateSuspiciousOverallConfigResponseBody) *OperateSuspiciousOverallConfigResponse
	GetBody() *OperateSuspiciousOverallConfigResponseBody
}

type OperateSuspiciousOverallConfigResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateSuspiciousOverallConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateSuspiciousOverallConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateSuspiciousOverallConfigResponse) GoString() string {
	return s.String()
}

func (s *OperateSuspiciousOverallConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateSuspiciousOverallConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateSuspiciousOverallConfigResponse) GetBody() *OperateSuspiciousOverallConfigResponseBody {
	return s.Body
}

func (s *OperateSuspiciousOverallConfigResponse) SetHeaders(v map[string]*string) *OperateSuspiciousOverallConfigResponse {
	s.Headers = v
	return s
}

func (s *OperateSuspiciousOverallConfigResponse) SetStatusCode(v int32) *OperateSuspiciousOverallConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateSuspiciousOverallConfigResponse) SetBody(v *OperateSuspiciousOverallConfigResponseBody) *OperateSuspiciousOverallConfigResponse {
	s.Body = v
	return s
}

func (s *OperateSuspiciousOverallConfigResponse) Validate() error {
	return dara.Validate(s)
}
