// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeMetricsBySignatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListComputeMetricsBySignatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListComputeMetricsBySignatureResponse
	GetStatusCode() *int32
	SetBody(v *ListComputeMetricsBySignatureResponseBody) *ListComputeMetricsBySignatureResponse
	GetBody() *ListComputeMetricsBySignatureResponseBody
}

type ListComputeMetricsBySignatureResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListComputeMetricsBySignatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListComputeMetricsBySignatureResponse) String() string {
	return dara.Prettify(s)
}

func (s ListComputeMetricsBySignatureResponse) GoString() string {
	return s.String()
}

func (s *ListComputeMetricsBySignatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListComputeMetricsBySignatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListComputeMetricsBySignatureResponse) GetBody() *ListComputeMetricsBySignatureResponseBody {
	return s.Body
}

func (s *ListComputeMetricsBySignatureResponse) SetHeaders(v map[string]*string) *ListComputeMetricsBySignatureResponse {
	s.Headers = v
	return s
}

func (s *ListComputeMetricsBySignatureResponse) SetStatusCode(v int32) *ListComputeMetricsBySignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *ListComputeMetricsBySignatureResponse) SetBody(v *ListComputeMetricsBySignatureResponseBody) *ListComputeMetricsBySignatureResponse {
	s.Body = v
	return s
}

func (s *ListComputeMetricsBySignatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
