// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCompliancePacksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCompliancePacksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCompliancePacksResponse
	GetStatusCode() *int32
	SetBody(v *ListCompliancePacksResponseBody) *ListCompliancePacksResponse
	GetBody() *ListCompliancePacksResponseBody
}

type ListCompliancePacksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCompliancePacksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCompliancePacksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCompliancePacksResponse) GoString() string {
	return s.String()
}

func (s *ListCompliancePacksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCompliancePacksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCompliancePacksResponse) GetBody() *ListCompliancePacksResponseBody {
	return s.Body
}

func (s *ListCompliancePacksResponse) SetHeaders(v map[string]*string) *ListCompliancePacksResponse {
	s.Headers = v
	return s
}

func (s *ListCompliancePacksResponse) SetStatusCode(v int32) *ListCompliancePacksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCompliancePacksResponse) SetBody(v *ListCompliancePacksResponseBody) *ListCompliancePacksResponse {
	s.Body = v
	return s
}

func (s *ListCompliancePacksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
