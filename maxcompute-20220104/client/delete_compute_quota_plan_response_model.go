// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComputeQuotaPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteComputeQuotaPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteComputeQuotaPlanResponse
	GetStatusCode() *int32
	SetBody(v *DeleteComputeQuotaPlanResponseBody) *DeleteComputeQuotaPlanResponse
	GetBody() *DeleteComputeQuotaPlanResponseBody
}

type DeleteComputeQuotaPlanResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteComputeQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteComputeQuotaPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteComputeQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteComputeQuotaPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteComputeQuotaPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteComputeQuotaPlanResponse) GetBody() *DeleteComputeQuotaPlanResponseBody {
	return s.Body
}

func (s *DeleteComputeQuotaPlanResponse) SetHeaders(v map[string]*string) *DeleteComputeQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteComputeQuotaPlanResponse) SetStatusCode(v int32) *DeleteComputeQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteComputeQuotaPlanResponse) SetBody(v *DeleteComputeQuotaPlanResponseBody) *DeleteComputeQuotaPlanResponse {
	s.Body = v
	return s
}

func (s *DeleteComputeQuotaPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
