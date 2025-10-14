// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTransportLayerApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTransportLayerApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTransportLayerApplicationResponse
	GetStatusCode() *int32
	SetBody(v *GetTransportLayerApplicationResponseBody) *GetTransportLayerApplicationResponse
	GetBody() *GetTransportLayerApplicationResponseBody
}

type GetTransportLayerApplicationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTransportLayerApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTransportLayerApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTransportLayerApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetTransportLayerApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTransportLayerApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTransportLayerApplicationResponse) GetBody() *GetTransportLayerApplicationResponseBody {
	return s.Body
}

func (s *GetTransportLayerApplicationResponse) SetHeaders(v map[string]*string) *GetTransportLayerApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetTransportLayerApplicationResponse) SetStatusCode(v int32) *GetTransportLayerApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTransportLayerApplicationResponse) SetBody(v *GetTransportLayerApplicationResponseBody) *GetTransportLayerApplicationResponse {
	s.Body = v
	return s
}

func (s *GetTransportLayerApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
