// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppResourceReserveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEdgeContainerAppResourceReserveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEdgeContainerAppResourceReserveResponse
	GetStatusCode() *int32
	SetBody(v *GetEdgeContainerAppResourceReserveResponseBody) *GetEdgeContainerAppResourceReserveResponse
	GetBody() *GetEdgeContainerAppResourceReserveResponseBody
}

type GetEdgeContainerAppResourceReserveResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEdgeContainerAppResourceReserveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEdgeContainerAppResourceReserveResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppResourceReserveResponse) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppResourceReserveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEdgeContainerAppResourceReserveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEdgeContainerAppResourceReserveResponse) GetBody() *GetEdgeContainerAppResourceReserveResponseBody {
	return s.Body
}

func (s *GetEdgeContainerAppResourceReserveResponse) SetHeaders(v map[string]*string) *GetEdgeContainerAppResourceReserveResponse {
	s.Headers = v
	return s
}

func (s *GetEdgeContainerAppResourceReserveResponse) SetStatusCode(v int32) *GetEdgeContainerAppResourceReserveResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEdgeContainerAppResourceReserveResponse) SetBody(v *GetEdgeContainerAppResourceReserveResponseBody) *GetEdgeContainerAppResourceReserveResponse {
	s.Body = v
	return s
}

func (s *GetEdgeContainerAppResourceReserveResponse) Validate() error {
	return dara.Validate(s)
}
