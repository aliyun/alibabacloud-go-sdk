// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstancesQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstancesQuotaResponse
	GetStatusCode() *int32
	SetBody(v *ListInstancesQuotaResponseBody) *ListInstancesQuotaResponse
	GetBody() *ListInstancesQuotaResponseBody
}

type ListInstancesQuotaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesQuotaResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstancesQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstancesQuotaResponse) GetBody() *ListInstancesQuotaResponseBody {
	return s.Body
}

func (s *ListInstancesQuotaResponse) SetHeaders(v map[string]*string) *ListInstancesQuotaResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesQuotaResponse) SetStatusCode(v int32) *ListInstancesQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesQuotaResponse) SetBody(v *ListInstancesQuotaResponseBody) *ListInstancesQuotaResponse {
	s.Body = v
	return s
}

func (s *ListInstancesQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
