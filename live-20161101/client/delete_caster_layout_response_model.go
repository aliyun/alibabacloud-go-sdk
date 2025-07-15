// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterLayoutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCasterLayoutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCasterLayoutResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCasterLayoutResponseBody) *DeleteCasterLayoutResponse
	GetBody() *DeleteCasterLayoutResponseBody
}

type DeleteCasterLayoutResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCasterLayoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCasterLayoutResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterLayoutResponse) GoString() string {
	return s.String()
}

func (s *DeleteCasterLayoutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCasterLayoutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCasterLayoutResponse) GetBody() *DeleteCasterLayoutResponseBody {
	return s.Body
}

func (s *DeleteCasterLayoutResponse) SetHeaders(v map[string]*string) *DeleteCasterLayoutResponse {
	s.Headers = v
	return s
}

func (s *DeleteCasterLayoutResponse) SetStatusCode(v int32) *DeleteCasterLayoutResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCasterLayoutResponse) SetBody(v *DeleteCasterLayoutResponseBody) *DeleteCasterLayoutResponse {
	s.Body = v
	return s
}

func (s *DeleteCasterLayoutResponse) Validate() error {
	return dara.Validate(s)
}
