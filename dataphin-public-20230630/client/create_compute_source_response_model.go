// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateComputeSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateComputeSourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateComputeSourceResponseBody) *CreateComputeSourceResponse
	GetBody() *CreateComputeSourceResponseBody
}

type CreateComputeSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateComputeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateComputeSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateComputeSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateComputeSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateComputeSourceResponse) GetBody() *CreateComputeSourceResponseBody {
	return s.Body
}

func (s *CreateComputeSourceResponse) SetHeaders(v map[string]*string) *CreateComputeSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateComputeSourceResponse) SetStatusCode(v int32) *CreateComputeSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateComputeSourceResponse) SetBody(v *CreateComputeSourceResponseBody) *CreateComputeSourceResponse {
	s.Body = v
	return s
}

func (s *CreateComputeSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
