// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReservedNodePoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteReservedNodePoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteReservedNodePoolResponse
	GetStatusCode() *int32
	SetBody(v *DeleteReservedNodePoolResponseBody) *DeleteReservedNodePoolResponse
	GetBody() *DeleteReservedNodePoolResponseBody
}

type DeleteReservedNodePoolResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteReservedNodePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteReservedNodePoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteReservedNodePoolResponse) GoString() string {
	return s.String()
}

func (s *DeleteReservedNodePoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteReservedNodePoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteReservedNodePoolResponse) GetBody() *DeleteReservedNodePoolResponseBody {
	return s.Body
}

func (s *DeleteReservedNodePoolResponse) SetHeaders(v map[string]*string) *DeleteReservedNodePoolResponse {
	s.Headers = v
	return s
}

func (s *DeleteReservedNodePoolResponse) SetStatusCode(v int32) *DeleteReservedNodePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteReservedNodePoolResponse) SetBody(v *DeleteReservedNodePoolResponseBody) *DeleteReservedNodePoolResponse {
	s.Body = v
	return s
}

func (s *DeleteReservedNodePoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
