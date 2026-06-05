// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComfyUserDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteComfyUserDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteComfyUserDataResponse
	GetStatusCode() *int32
	SetBody(v *DeleteComfyUserDataResponseBody) *DeleteComfyUserDataResponse
	GetBody() *DeleteComfyUserDataResponseBody
}

type DeleteComfyUserDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteComfyUserDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteComfyUserDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteComfyUserDataResponse) GoString() string {
	return s.String()
}

func (s *DeleteComfyUserDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteComfyUserDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteComfyUserDataResponse) GetBody() *DeleteComfyUserDataResponseBody {
	return s.Body
}

func (s *DeleteComfyUserDataResponse) SetHeaders(v map[string]*string) *DeleteComfyUserDataResponse {
	s.Headers = v
	return s
}

func (s *DeleteComfyUserDataResponse) SetStatusCode(v int32) *DeleteComfyUserDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteComfyUserDataResponse) SetBody(v *DeleteComfyUserDataResponseBody) *DeleteComfyUserDataResponse {
	s.Body = v
	return s
}

func (s *DeleteComfyUserDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
