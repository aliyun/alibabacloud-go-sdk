// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckComputeClusterConnectivityByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckComputeClusterConnectivityByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckComputeClusterConnectivityByIdResponse
	GetStatusCode() *int32
	SetBody(v *CheckComputeClusterConnectivityByIdResponseBody) *CheckComputeClusterConnectivityByIdResponse
	GetBody() *CheckComputeClusterConnectivityByIdResponseBody
}

type CheckComputeClusterConnectivityByIdResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckComputeClusterConnectivityByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckComputeClusterConnectivityByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckComputeClusterConnectivityByIdResponse) GoString() string {
	return s.String()
}

func (s *CheckComputeClusterConnectivityByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckComputeClusterConnectivityByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckComputeClusterConnectivityByIdResponse) GetBody() *CheckComputeClusterConnectivityByIdResponseBody {
	return s.Body
}

func (s *CheckComputeClusterConnectivityByIdResponse) SetHeaders(v map[string]*string) *CheckComputeClusterConnectivityByIdResponse {
	s.Headers = v
	return s
}

func (s *CheckComputeClusterConnectivityByIdResponse) SetStatusCode(v int32) *CheckComputeClusterConnectivityByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckComputeClusterConnectivityByIdResponse) SetBody(v *CheckComputeClusterConnectivityByIdResponseBody) *CheckComputeClusterConnectivityByIdResponse {
	s.Body = v
	return s
}

func (s *CheckComputeClusterConnectivityByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
