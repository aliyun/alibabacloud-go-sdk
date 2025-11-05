// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSpecificConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSpecificConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSpecificConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSpecificConfigResponseBody) *DeleteSpecificConfigResponse
	GetBody() *DeleteSpecificConfigResponseBody
}

type DeleteSpecificConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSpecificConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSpecificConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSpecificConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteSpecificConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSpecificConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSpecificConfigResponse) GetBody() *DeleteSpecificConfigResponseBody {
	return s.Body
}

func (s *DeleteSpecificConfigResponse) SetHeaders(v map[string]*string) *DeleteSpecificConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteSpecificConfigResponse) SetStatusCode(v int32) *DeleteSpecificConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSpecificConfigResponse) SetBody(v *DeleteSpecificConfigResponseBody) *DeleteSpecificConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteSpecificConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
