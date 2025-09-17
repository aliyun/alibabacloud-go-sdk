// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransportLayerApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTransportLayerApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTransportLayerApplicationResponse
	GetStatusCode() *int32
	SetBody(v *CreateTransportLayerApplicationResponseBody) *CreateTransportLayerApplicationResponse
	GetBody() *CreateTransportLayerApplicationResponseBody
}

type CreateTransportLayerApplicationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTransportLayerApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTransportLayerApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTransportLayerApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateTransportLayerApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTransportLayerApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTransportLayerApplicationResponse) GetBody() *CreateTransportLayerApplicationResponseBody {
	return s.Body
}

func (s *CreateTransportLayerApplicationResponse) SetHeaders(v map[string]*string) *CreateTransportLayerApplicationResponse {
	s.Headers = v
	return s
}

func (s *CreateTransportLayerApplicationResponse) SetStatusCode(v int32) *CreateTransportLayerApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTransportLayerApplicationResponse) SetBody(v *CreateTransportLayerApplicationResponseBody) *CreateTransportLayerApplicationResponse {
	s.Body = v
	return s
}

func (s *CreateTransportLayerApplicationResponse) Validate() error {
	return dara.Validate(s)
}
