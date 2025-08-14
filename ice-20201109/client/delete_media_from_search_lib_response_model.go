// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaFromSearchLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMediaFromSearchLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMediaFromSearchLibResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMediaFromSearchLibResponseBody) *DeleteMediaFromSearchLibResponse
	GetBody() *DeleteMediaFromSearchLibResponseBody
}

type DeleteMediaFromSearchLibResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMediaFromSearchLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMediaFromSearchLibResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaFromSearchLibResponse) GoString() string {
	return s.String()
}

func (s *DeleteMediaFromSearchLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMediaFromSearchLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMediaFromSearchLibResponse) GetBody() *DeleteMediaFromSearchLibResponseBody {
	return s.Body
}

func (s *DeleteMediaFromSearchLibResponse) SetHeaders(v map[string]*string) *DeleteMediaFromSearchLibResponse {
	s.Headers = v
	return s
}

func (s *DeleteMediaFromSearchLibResponse) SetStatusCode(v int32) *DeleteMediaFromSearchLibResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMediaFromSearchLibResponse) SetBody(v *DeleteMediaFromSearchLibResponseBody) *DeleteMediaFromSearchLibResponse {
	s.Body = v
	return s
}

func (s *DeleteMediaFromSearchLibResponse) Validate() error {
	return dara.Validate(s)
}
