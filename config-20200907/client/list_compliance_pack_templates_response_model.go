// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCompliancePackTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCompliancePackTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCompliancePackTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListCompliancePackTemplatesResponseBody) *ListCompliancePackTemplatesResponse
	GetBody() *ListCompliancePackTemplatesResponseBody
}

type ListCompliancePackTemplatesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCompliancePackTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCompliancePackTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCompliancePackTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListCompliancePackTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCompliancePackTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCompliancePackTemplatesResponse) GetBody() *ListCompliancePackTemplatesResponseBody {
	return s.Body
}

func (s *ListCompliancePackTemplatesResponse) SetHeaders(v map[string]*string) *ListCompliancePackTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListCompliancePackTemplatesResponse) SetStatusCode(v int32) *ListCompliancePackTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCompliancePackTemplatesResponse) SetBody(v *ListCompliancePackTemplatesResponseBody) *ListCompliancePackTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListCompliancePackTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
