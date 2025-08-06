// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIImageInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAIImageInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAIImageInfosResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAIImageInfosResponseBody) *DeleteAIImageInfosResponse
	GetBody() *DeleteAIImageInfosResponseBody
}

type DeleteAIImageInfosResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAIImageInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAIImageInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIImageInfosResponse) GoString() string {
	return s.String()
}

func (s *DeleteAIImageInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAIImageInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAIImageInfosResponse) GetBody() *DeleteAIImageInfosResponseBody {
	return s.Body
}

func (s *DeleteAIImageInfosResponse) SetHeaders(v map[string]*string) *DeleteAIImageInfosResponse {
	s.Headers = v
	return s
}

func (s *DeleteAIImageInfosResponse) SetStatusCode(v int32) *DeleteAIImageInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAIImageInfosResponse) SetBody(v *DeleteAIImageInfosResponseBody) *DeleteAIImageInfosResponse {
	s.Body = v
	return s
}

func (s *DeleteAIImageInfosResponse) Validate() error {
	return dara.Validate(s)
}
