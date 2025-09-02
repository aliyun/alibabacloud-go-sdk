// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRemindResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRemindResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRemindResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRemindResponseBody) *DeleteRemindResponse
	GetBody() *DeleteRemindResponseBody
}

type DeleteRemindResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRemindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRemindResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRemindResponse) GoString() string {
	return s.String()
}

func (s *DeleteRemindResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRemindResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRemindResponse) GetBody() *DeleteRemindResponseBody {
	return s.Body
}

func (s *DeleteRemindResponse) SetHeaders(v map[string]*string) *DeleteRemindResponse {
	s.Headers = v
	return s
}

func (s *DeleteRemindResponse) SetStatusCode(v int32) *DeleteRemindResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRemindResponse) SetBody(v *DeleteRemindResponseBody) *DeleteRemindResponse {
	s.Body = v
	return s
}

func (s *DeleteRemindResponse) Validate() error {
	return dara.Validate(s)
}
