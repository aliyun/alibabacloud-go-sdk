// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndexVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIndexVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIndexVersionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIndexVersionResponseBody) *DeleteIndexVersionResponse
	GetBody() *DeleteIndexVersionResponseBody
}

type DeleteIndexVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIndexVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIndexVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndexVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteIndexVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIndexVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIndexVersionResponse) GetBody() *DeleteIndexVersionResponseBody {
	return s.Body
}

func (s *DeleteIndexVersionResponse) SetHeaders(v map[string]*string) *DeleteIndexVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteIndexVersionResponse) SetStatusCode(v int32) *DeleteIndexVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIndexVersionResponse) SetBody(v *DeleteIndexVersionResponseBody) *DeleteIndexVersionResponse {
	s.Body = v
	return s
}

func (s *DeleteIndexVersionResponse) Validate() error {
	return dara.Validate(s)
}
