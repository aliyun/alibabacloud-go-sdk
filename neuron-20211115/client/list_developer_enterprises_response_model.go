// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeveloperEnterprisesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeveloperEnterprisesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeveloperEnterprisesResponse
	GetStatusCode() *int32
	SetBody(v *EnterprisePageResult) *ListDeveloperEnterprisesResponse
	GetBody() *EnterprisePageResult
}

type ListDeveloperEnterprisesResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterprisePageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeveloperEnterprisesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeveloperEnterprisesResponse) GoString() string {
	return s.String()
}

func (s *ListDeveloperEnterprisesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeveloperEnterprisesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeveloperEnterprisesResponse) GetBody() *EnterprisePageResult {
	return s.Body
}

func (s *ListDeveloperEnterprisesResponse) SetHeaders(v map[string]*string) *ListDeveloperEnterprisesResponse {
	s.Headers = v
	return s
}

func (s *ListDeveloperEnterprisesResponse) SetStatusCode(v int32) *ListDeveloperEnterprisesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeveloperEnterprisesResponse) SetBody(v *EnterprisePageResult) *ListDeveloperEnterprisesResponse {
	s.Body = v
	return s
}

func (s *ListDeveloperEnterprisesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
