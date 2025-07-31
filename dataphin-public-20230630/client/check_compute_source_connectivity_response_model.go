// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckComputeSourceConnectivityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckComputeSourceConnectivityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckComputeSourceConnectivityResponse
	GetStatusCode() *int32
	SetBody(v *CheckComputeSourceConnectivityResponseBody) *CheckComputeSourceConnectivityResponse
	GetBody() *CheckComputeSourceConnectivityResponseBody
}

type CheckComputeSourceConnectivityResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckComputeSourceConnectivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckComputeSourceConnectivityResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckComputeSourceConnectivityResponse) GoString() string {
	return s.String()
}

func (s *CheckComputeSourceConnectivityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckComputeSourceConnectivityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckComputeSourceConnectivityResponse) GetBody() *CheckComputeSourceConnectivityResponseBody {
	return s.Body
}

func (s *CheckComputeSourceConnectivityResponse) SetHeaders(v map[string]*string) *CheckComputeSourceConnectivityResponse {
	s.Headers = v
	return s
}

func (s *CheckComputeSourceConnectivityResponse) SetStatusCode(v int32) *CheckComputeSourceConnectivityResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckComputeSourceConnectivityResponse) SetBody(v *CheckComputeSourceConnectivityResponseBody) *CheckComputeSourceConnectivityResponse {
	s.Body = v
	return s
}

func (s *CheckComputeSourceConnectivityResponse) Validate() error {
	return dara.Validate(s)
}
