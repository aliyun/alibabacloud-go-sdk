// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseAccelerateTargetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnterpriseAccelerateTargetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnterpriseAccelerateTargetsResponse
	GetStatusCode() *int32
	SetBody(v *ListEnterpriseAccelerateTargetsResponseBody) *ListEnterpriseAccelerateTargetsResponse
	GetBody() *ListEnterpriseAccelerateTargetsResponseBody
}

type ListEnterpriseAccelerateTargetsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnterpriseAccelerateTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnterpriseAccelerateTargetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseAccelerateTargetsResponse) GoString() string {
	return s.String()
}

func (s *ListEnterpriseAccelerateTargetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnterpriseAccelerateTargetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnterpriseAccelerateTargetsResponse) GetBody() *ListEnterpriseAccelerateTargetsResponseBody {
	return s.Body
}

func (s *ListEnterpriseAccelerateTargetsResponse) SetHeaders(v map[string]*string) *ListEnterpriseAccelerateTargetsResponse {
	s.Headers = v
	return s
}

func (s *ListEnterpriseAccelerateTargetsResponse) SetStatusCode(v int32) *ListEnterpriseAccelerateTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnterpriseAccelerateTargetsResponse) SetBody(v *ListEnterpriseAccelerateTargetsResponseBody) *ListEnterpriseAccelerateTargetsResponse {
	s.Body = v
	return s
}

func (s *ListEnterpriseAccelerateTargetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
