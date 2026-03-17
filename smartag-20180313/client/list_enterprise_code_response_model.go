// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnterpriseCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnterpriseCodeResponse
	GetStatusCode() *int32
	SetBody(v *ListEnterpriseCodeResponseBody) *ListEnterpriseCodeResponse
	GetBody() *ListEnterpriseCodeResponseBody
}

type ListEnterpriseCodeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnterpriseCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnterpriseCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseCodeResponse) GoString() string {
	return s.String()
}

func (s *ListEnterpriseCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnterpriseCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnterpriseCodeResponse) GetBody() *ListEnterpriseCodeResponseBody {
	return s.Body
}

func (s *ListEnterpriseCodeResponse) SetHeaders(v map[string]*string) *ListEnterpriseCodeResponse {
	s.Headers = v
	return s
}

func (s *ListEnterpriseCodeResponse) SetStatusCode(v int32) *ListEnterpriseCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnterpriseCodeResponse) SetBody(v *ListEnterpriseCodeResponseBody) *ListEnterpriseCodeResponse {
	s.Body = v
	return s
}

func (s *ListEnterpriseCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
