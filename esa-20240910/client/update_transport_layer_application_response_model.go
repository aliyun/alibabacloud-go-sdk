// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransportLayerApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTransportLayerApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTransportLayerApplicationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTransportLayerApplicationResponseBody) *UpdateTransportLayerApplicationResponse
	GetBody() *UpdateTransportLayerApplicationResponseBody
}

type UpdateTransportLayerApplicationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTransportLayerApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTransportLayerApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransportLayerApplicationResponse) GoString() string {
	return s.String()
}

func (s *UpdateTransportLayerApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTransportLayerApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTransportLayerApplicationResponse) GetBody() *UpdateTransportLayerApplicationResponseBody {
	return s.Body
}

func (s *UpdateTransportLayerApplicationResponse) SetHeaders(v map[string]*string) *UpdateTransportLayerApplicationResponse {
	s.Headers = v
	return s
}

func (s *UpdateTransportLayerApplicationResponse) SetStatusCode(v int32) *UpdateTransportLayerApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTransportLayerApplicationResponse) SetBody(v *UpdateTransportLayerApplicationResponseBody) *UpdateTransportLayerApplicationResponse {
	s.Body = v
	return s
}

func (s *UpdateTransportLayerApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
