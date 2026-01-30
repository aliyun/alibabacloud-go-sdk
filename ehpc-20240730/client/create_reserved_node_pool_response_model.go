// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReservedNodePoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateReservedNodePoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateReservedNodePoolResponse
	GetStatusCode() *int32
	SetBody(v *CreateReservedNodePoolResponseBody) *CreateReservedNodePoolResponse
	GetBody() *CreateReservedNodePoolResponseBody
}

type CreateReservedNodePoolResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateReservedNodePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateReservedNodePoolResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateReservedNodePoolResponse) GoString() string {
	return s.String()
}

func (s *CreateReservedNodePoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateReservedNodePoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateReservedNodePoolResponse) GetBody() *CreateReservedNodePoolResponseBody {
	return s.Body
}

func (s *CreateReservedNodePoolResponse) SetHeaders(v map[string]*string) *CreateReservedNodePoolResponse {
	s.Headers = v
	return s
}

func (s *CreateReservedNodePoolResponse) SetStatusCode(v int32) *CreateReservedNodePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateReservedNodePoolResponse) SetBody(v *CreateReservedNodePoolResponseBody) *CreateReservedNodePoolResponse {
	s.Body = v
	return s
}

func (s *CreateReservedNodePoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
