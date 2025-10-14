// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransportLayerApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTransportLayerApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTransportLayerApplicationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTransportLayerApplicationResponseBody) *DeleteTransportLayerApplicationResponse
	GetBody() *DeleteTransportLayerApplicationResponseBody
}

type DeleteTransportLayerApplicationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTransportLayerApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTransportLayerApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransportLayerApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransportLayerApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTransportLayerApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTransportLayerApplicationResponse) GetBody() *DeleteTransportLayerApplicationResponseBody {
	return s.Body
}

func (s *DeleteTransportLayerApplicationResponse) SetHeaders(v map[string]*string) *DeleteTransportLayerApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransportLayerApplicationResponse) SetStatusCode(v int32) *DeleteTransportLayerApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTransportLayerApplicationResponse) SetBody(v *DeleteTransportLayerApplicationResponseBody) *DeleteTransportLayerApplicationResponse {
	s.Body = v
	return s
}

func (s *DeleteTransportLayerApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
