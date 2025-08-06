// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVodTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVodTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVodTemplateResponseBody) *DeleteVodTemplateResponse
	GetBody() *DeleteVodTemplateResponseBody
}

type DeleteVodTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVodTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVodTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteVodTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVodTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVodTemplateResponse) GetBody() *DeleteVodTemplateResponseBody {
	return s.Body
}

func (s *DeleteVodTemplateResponse) SetHeaders(v map[string]*string) *DeleteVodTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteVodTemplateResponse) SetStatusCode(v int32) *DeleteVodTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVodTemplateResponse) SetBody(v *DeleteVodTemplateResponseBody) *DeleteVodTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteVodTemplateResponse) Validate() error {
	return dara.Validate(s)
}
