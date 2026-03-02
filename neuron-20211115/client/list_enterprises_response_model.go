// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterprisesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnterprisesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnterprisesResponse
	GetStatusCode() *int32
	SetBody(v *EnterprisePageResult) *ListEnterprisesResponse
	GetBody() *EnterprisePageResult
}

type ListEnterprisesResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterprisePageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnterprisesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnterprisesResponse) GoString() string {
	return s.String()
}

func (s *ListEnterprisesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnterprisesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnterprisesResponse) GetBody() *EnterprisePageResult {
	return s.Body
}

func (s *ListEnterprisesResponse) SetHeaders(v map[string]*string) *ListEnterprisesResponse {
	s.Headers = v
	return s
}

func (s *ListEnterprisesResponse) SetStatusCode(v int32) *ListEnterprisesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnterprisesResponse) SetBody(v *EnterprisePageResult) *ListEnterprisesResponse {
	s.Body = v
	return s
}

func (s *ListEnterprisesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
