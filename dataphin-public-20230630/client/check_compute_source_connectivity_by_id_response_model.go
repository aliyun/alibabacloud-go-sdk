// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckComputeSourceConnectivityByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckComputeSourceConnectivityByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckComputeSourceConnectivityByIdResponse
	GetStatusCode() *int32
	SetBody(v *CheckComputeSourceConnectivityByIdResponseBody) *CheckComputeSourceConnectivityByIdResponse
	GetBody() *CheckComputeSourceConnectivityByIdResponseBody
}

type CheckComputeSourceConnectivityByIdResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckComputeSourceConnectivityByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckComputeSourceConnectivityByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckComputeSourceConnectivityByIdResponse) GoString() string {
	return s.String()
}

func (s *CheckComputeSourceConnectivityByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckComputeSourceConnectivityByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckComputeSourceConnectivityByIdResponse) GetBody() *CheckComputeSourceConnectivityByIdResponseBody {
	return s.Body
}

func (s *CheckComputeSourceConnectivityByIdResponse) SetHeaders(v map[string]*string) *CheckComputeSourceConnectivityByIdResponse {
	s.Headers = v
	return s
}

func (s *CheckComputeSourceConnectivityByIdResponse) SetStatusCode(v int32) *CheckComputeSourceConnectivityByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckComputeSourceConnectivityByIdResponse) SetBody(v *CheckComputeSourceConnectivityByIdResponseBody) *CheckComputeSourceConnectivityByIdResponse {
	s.Body = v
	return s
}

func (s *CheckComputeSourceConnectivityByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
