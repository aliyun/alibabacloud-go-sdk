// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComputeInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteComputeInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteComputeInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteComputeInstanceResponseBody) *DeleteComputeInstanceResponse
	GetBody() *DeleteComputeInstanceResponseBody
}

type DeleteComputeInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteComputeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteComputeInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteComputeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteComputeInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteComputeInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteComputeInstanceResponse) GetBody() *DeleteComputeInstanceResponseBody {
	return s.Body
}

func (s *DeleteComputeInstanceResponse) SetHeaders(v map[string]*string) *DeleteComputeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteComputeInstanceResponse) SetStatusCode(v int32) *DeleteComputeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteComputeInstanceResponse) SetBody(v *DeleteComputeInstanceResponseBody) *DeleteComputeInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteComputeInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
