// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetComputeSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetComputeSourceResponse
	GetStatusCode() *int32
	SetBody(v *GetComputeSourceResponseBody) *GetComputeSourceResponse
	GetBody() *GetComputeSourceResponseBody
}

type GetComputeSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetComputeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetComputeSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetComputeSourceResponse) GoString() string {
	return s.String()
}

func (s *GetComputeSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetComputeSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetComputeSourceResponse) GetBody() *GetComputeSourceResponseBody {
	return s.Body
}

func (s *GetComputeSourceResponse) SetHeaders(v map[string]*string) *GetComputeSourceResponse {
	s.Headers = v
	return s
}

func (s *GetComputeSourceResponse) SetStatusCode(v int32) *GetComputeSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetComputeSourceResponse) SetBody(v *GetComputeSourceResponseBody) *GetComputeSourceResponse {
	s.Body = v
	return s
}

func (s *GetComputeSourceResponse) Validate() error {
	return dara.Validate(s)
}
