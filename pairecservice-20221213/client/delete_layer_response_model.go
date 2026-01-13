// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLayerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLayerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLayerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLayerResponseBody) *DeleteLayerResponse
	GetBody() *DeleteLayerResponseBody
}

type DeleteLayerResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLayerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLayerResponse) GoString() string {
	return s.String()
}

func (s *DeleteLayerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLayerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLayerResponse) GetBody() *DeleteLayerResponseBody {
	return s.Body
}

func (s *DeleteLayerResponse) SetHeaders(v map[string]*string) *DeleteLayerResponse {
	s.Headers = v
	return s
}

func (s *DeleteLayerResponse) SetStatusCode(v int32) *DeleteLayerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLayerResponse) SetBody(v *DeleteLayerResponseBody) *DeleteLayerResponse {
	s.Body = v
	return s
}

func (s *DeleteLayerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
