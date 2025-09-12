// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenComputePreCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenComputePreCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenComputePreCheckResponse
	GetStatusCode() *int32
	SetBody(v *OpenComputePreCheckResponseBody) *OpenComputePreCheckResponse
	GetBody() *OpenComputePreCheckResponseBody
}

type OpenComputePreCheckResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenComputePreCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenComputePreCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenComputePreCheckResponse) GoString() string {
	return s.String()
}

func (s *OpenComputePreCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenComputePreCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenComputePreCheckResponse) GetBody() *OpenComputePreCheckResponseBody {
	return s.Body
}

func (s *OpenComputePreCheckResponse) SetHeaders(v map[string]*string) *OpenComputePreCheckResponse {
	s.Headers = v
	return s
}

func (s *OpenComputePreCheckResponse) SetStatusCode(v int32) *OpenComputePreCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenComputePreCheckResponse) SetBody(v *OpenComputePreCheckResponseBody) *OpenComputePreCheckResponse {
	s.Body = v
	return s
}

func (s *OpenComputePreCheckResponse) Validate() error {
	return dara.Validate(s)
}
