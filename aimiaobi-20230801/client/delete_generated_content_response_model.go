// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGeneratedContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGeneratedContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGeneratedContentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGeneratedContentResponseBody) *DeleteGeneratedContentResponse
	GetBody() *DeleteGeneratedContentResponseBody
}

type DeleteGeneratedContentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGeneratedContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGeneratedContentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGeneratedContentResponse) GoString() string {
	return s.String()
}

func (s *DeleteGeneratedContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGeneratedContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGeneratedContentResponse) GetBody() *DeleteGeneratedContentResponseBody {
	return s.Body
}

func (s *DeleteGeneratedContentResponse) SetHeaders(v map[string]*string) *DeleteGeneratedContentResponse {
	s.Headers = v
	return s
}

func (s *DeleteGeneratedContentResponse) SetStatusCode(v int32) *DeleteGeneratedContentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGeneratedContentResponse) SetBody(v *DeleteGeneratedContentResponseBody) *DeleteGeneratedContentResponse {
	s.Body = v
	return s
}

func (s *DeleteGeneratedContentResponse) Validate() error {
	return dara.Validate(s)
}
