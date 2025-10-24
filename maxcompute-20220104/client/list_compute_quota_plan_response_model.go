// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeQuotaPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListComputeQuotaPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListComputeQuotaPlanResponse
	GetStatusCode() *int32
	SetBody(v *ListComputeQuotaPlanResponseBody) *ListComputeQuotaPlanResponse
	GetBody() *ListComputeQuotaPlanResponseBody
}

type ListComputeQuotaPlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListComputeQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListComputeQuotaPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s ListComputeQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *ListComputeQuotaPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListComputeQuotaPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListComputeQuotaPlanResponse) GetBody() *ListComputeQuotaPlanResponseBody {
	return s.Body
}

func (s *ListComputeQuotaPlanResponse) SetHeaders(v map[string]*string) *ListComputeQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *ListComputeQuotaPlanResponse) SetStatusCode(v int32) *ListComputeQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *ListComputeQuotaPlanResponse) SetBody(v *ListComputeQuotaPlanResponseBody) *ListComputeQuotaPlanResponse {
	s.Body = v
	return s
}

func (s *ListComputeQuotaPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
