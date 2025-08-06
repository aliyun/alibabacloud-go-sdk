// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAttachedMediaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAttachedMediaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAttachedMediaResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAttachedMediaResponseBody) *DeleteAttachedMediaResponse
	GetBody() *DeleteAttachedMediaResponseBody
}

type DeleteAttachedMediaResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAttachedMediaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAttachedMediaResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAttachedMediaResponse) GoString() string {
	return s.String()
}

func (s *DeleteAttachedMediaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAttachedMediaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAttachedMediaResponse) GetBody() *DeleteAttachedMediaResponseBody {
	return s.Body
}

func (s *DeleteAttachedMediaResponse) SetHeaders(v map[string]*string) *DeleteAttachedMediaResponse {
	s.Headers = v
	return s
}

func (s *DeleteAttachedMediaResponse) SetStatusCode(v int32) *DeleteAttachedMediaResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAttachedMediaResponse) SetBody(v *DeleteAttachedMediaResponseBody) *DeleteAttachedMediaResponse {
	s.Body = v
	return s
}

func (s *DeleteAttachedMediaResponse) Validate() error {
	return dara.Validate(s)
}
