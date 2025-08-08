// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcdpAimResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMcdpAimResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMcdpAimResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMcdpAimResponseBody) *DeleteMcdpAimResponse
	GetBody() *DeleteMcdpAimResponseBody
}

type DeleteMcdpAimResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMcdpAimResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMcdpAimResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcdpAimResponse) GoString() string {
	return s.String()
}

func (s *DeleteMcdpAimResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMcdpAimResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMcdpAimResponse) GetBody() *DeleteMcdpAimResponseBody {
	return s.Body
}

func (s *DeleteMcdpAimResponse) SetHeaders(v map[string]*string) *DeleteMcdpAimResponse {
	s.Headers = v
	return s
}

func (s *DeleteMcdpAimResponse) SetStatusCode(v int32) *DeleteMcdpAimResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMcdpAimResponse) SetBody(v *DeleteMcdpAimResponseBody) *DeleteMcdpAimResponse {
	s.Body = v
	return s
}

func (s *DeleteMcdpAimResponse) Validate() error {
	return dara.Validate(s)
}
