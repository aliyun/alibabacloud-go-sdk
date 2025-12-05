// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKeyMaterialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKeyMaterialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKeyMaterialResponse
	GetStatusCode() *int32
	SetBody(v *DeleteKeyMaterialResponseBody) *DeleteKeyMaterialResponse
	GetBody() *DeleteKeyMaterialResponseBody
}

type DeleteKeyMaterialResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKeyMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKeyMaterialResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKeyMaterialResponse) GoString() string {
	return s.String()
}

func (s *DeleteKeyMaterialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKeyMaterialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKeyMaterialResponse) GetBody() *DeleteKeyMaterialResponseBody {
	return s.Body
}

func (s *DeleteKeyMaterialResponse) SetHeaders(v map[string]*string) *DeleteKeyMaterialResponse {
	s.Headers = v
	return s
}

func (s *DeleteKeyMaterialResponse) SetStatusCode(v int32) *DeleteKeyMaterialResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKeyMaterialResponse) SetBody(v *DeleteKeyMaterialResponseBody) *DeleteKeyMaterialResponse {
	s.Body = v
	return s
}

func (s *DeleteKeyMaterialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
