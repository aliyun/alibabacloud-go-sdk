// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeQuotaPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateComputeQuotaPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateComputeQuotaPlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateComputeQuotaPlanResponseBody) *CreateComputeQuotaPlanResponse
	GetBody() *CreateComputeQuotaPlanResponseBody
}

type CreateComputeQuotaPlanResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateComputeQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateComputeQuotaPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateComputeQuotaPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateComputeQuotaPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateComputeQuotaPlanResponse) GetBody() *CreateComputeQuotaPlanResponseBody {
	return s.Body
}

func (s *CreateComputeQuotaPlanResponse) SetHeaders(v map[string]*string) *CreateComputeQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateComputeQuotaPlanResponse) SetStatusCode(v int32) *CreateComputeQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateComputeQuotaPlanResponse) SetBody(v *CreateComputeQuotaPlanResponseBody) *CreateComputeQuotaPlanResponse {
	s.Body = v
	return s
}

func (s *CreateComputeQuotaPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
