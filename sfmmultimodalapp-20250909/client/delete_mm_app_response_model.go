// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMmAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMmAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMmAppResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMmAppResponseBody) *DeleteMmAppResponse
	GetBody() *DeleteMmAppResponseBody
}

type DeleteMmAppResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMmAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMmAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMmAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteMmAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMmAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMmAppResponse) GetBody() *DeleteMmAppResponseBody {
	return s.Body
}

func (s *DeleteMmAppResponse) SetHeaders(v map[string]*string) *DeleteMmAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteMmAppResponse) SetStatusCode(v int32) *DeleteMmAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMmAppResponse) SetBody(v *DeleteMmAppResponseBody) *DeleteMmAppResponse {
	s.Body = v
	return s
}

func (s *DeleteMmAppResponse) Validate() error {
	return dara.Validate(s)
}
