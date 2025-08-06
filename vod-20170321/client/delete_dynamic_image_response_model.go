// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDynamicImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDynamicImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDynamicImageResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDynamicImageResponseBody) *DeleteDynamicImageResponse
	GetBody() *DeleteDynamicImageResponseBody
}

type DeleteDynamicImageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDynamicImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDynamicImageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDynamicImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteDynamicImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDynamicImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDynamicImageResponse) GetBody() *DeleteDynamicImageResponseBody {
	return s.Body
}

func (s *DeleteDynamicImageResponse) SetHeaders(v map[string]*string) *DeleteDynamicImageResponse {
	s.Headers = v
	return s
}

func (s *DeleteDynamicImageResponse) SetStatusCode(v int32) *DeleteDynamicImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDynamicImageResponse) SetBody(v *DeleteDynamicImageResponseBody) *DeleteDynamicImageResponse {
	s.Body = v
	return s
}

func (s *DeleteDynamicImageResponse) Validate() error {
	return dara.Validate(s)
}
