// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnterpriseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnterpriseInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListEnterpriseInfoResponseBody) *ListEnterpriseInfoResponse
	GetBody() *ListEnterpriseInfoResponseBody
}

type ListEnterpriseInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnterpriseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnterpriseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseInfoResponse) GoString() string {
	return s.String()
}

func (s *ListEnterpriseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnterpriseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnterpriseInfoResponse) GetBody() *ListEnterpriseInfoResponseBody {
	return s.Body
}

func (s *ListEnterpriseInfoResponse) SetHeaders(v map[string]*string) *ListEnterpriseInfoResponse {
	s.Headers = v
	return s
}

func (s *ListEnterpriseInfoResponse) SetStatusCode(v int32) *ListEnterpriseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnterpriseInfoResponse) SetBody(v *ListEnterpriseInfoResponseBody) *ListEnterpriseInfoResponse {
	s.Body = v
	return s
}

func (s *ListEnterpriseInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
