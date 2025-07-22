// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppStreamingOutTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppStreamingOutTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppStreamingOutTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppStreamingOutTemplateResponseBody) *DeleteAppStreamingOutTemplateResponse
	GetBody() *DeleteAppStreamingOutTemplateResponseBody
}

type DeleteAppStreamingOutTemplateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppStreamingOutTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppStreamingOutTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppStreamingOutTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppStreamingOutTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppStreamingOutTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppStreamingOutTemplateResponse) GetBody() *DeleteAppStreamingOutTemplateResponseBody {
	return s.Body
}

func (s *DeleteAppStreamingOutTemplateResponse) SetHeaders(v map[string]*string) *DeleteAppStreamingOutTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppStreamingOutTemplateResponse) SetStatusCode(v int32) *DeleteAppStreamingOutTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppStreamingOutTemplateResponse) SetBody(v *DeleteAppStreamingOutTemplateResponseBody) *DeleteAppStreamingOutTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteAppStreamingOutTemplateResponse) Validate() error {
	return dara.Validate(s)
}
