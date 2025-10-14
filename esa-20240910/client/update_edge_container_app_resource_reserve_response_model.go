// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEdgeContainerAppResourceReserveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEdgeContainerAppResourceReserveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEdgeContainerAppResourceReserveResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEdgeContainerAppResourceReserveResponseBody) *UpdateEdgeContainerAppResourceReserveResponse
	GetBody() *UpdateEdgeContainerAppResourceReserveResponseBody
}

type UpdateEdgeContainerAppResourceReserveResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEdgeContainerAppResourceReserveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEdgeContainerAppResourceReserveResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeContainerAppResourceReserveResponse) GoString() string {
	return s.String()
}

func (s *UpdateEdgeContainerAppResourceReserveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEdgeContainerAppResourceReserveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEdgeContainerAppResourceReserveResponse) GetBody() *UpdateEdgeContainerAppResourceReserveResponseBody {
	return s.Body
}

func (s *UpdateEdgeContainerAppResourceReserveResponse) SetHeaders(v map[string]*string) *UpdateEdgeContainerAppResourceReserveResponse {
	s.Headers = v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveResponse) SetStatusCode(v int32) *UpdateEdgeContainerAppResourceReserveResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveResponse) SetBody(v *UpdateEdgeContainerAppResourceReserveResponseBody) *UpdateEdgeContainerAppResourceReserveResponse {
	s.Body = v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
