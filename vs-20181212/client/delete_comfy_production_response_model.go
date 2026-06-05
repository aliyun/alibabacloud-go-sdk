// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComfyProductionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteComfyProductionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteComfyProductionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteComfyProductionResponseBody) *DeleteComfyProductionResponse
	GetBody() *DeleteComfyProductionResponseBody
}

type DeleteComfyProductionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteComfyProductionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteComfyProductionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteComfyProductionResponse) GoString() string {
	return s.String()
}

func (s *DeleteComfyProductionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteComfyProductionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteComfyProductionResponse) GetBody() *DeleteComfyProductionResponseBody {
	return s.Body
}

func (s *DeleteComfyProductionResponse) SetHeaders(v map[string]*string) *DeleteComfyProductionResponse {
	s.Headers = v
	return s
}

func (s *DeleteComfyProductionResponse) SetStatusCode(v int32) *DeleteComfyProductionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteComfyProductionResponse) SetBody(v *DeleteComfyProductionResponseBody) *DeleteComfyProductionResponse {
	s.Body = v
	return s
}

func (s *DeleteComfyProductionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
