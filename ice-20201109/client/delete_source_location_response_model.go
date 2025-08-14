// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSourceLocationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSourceLocationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSourceLocationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSourceLocationResponseBody) *DeleteSourceLocationResponse
	GetBody() *DeleteSourceLocationResponseBody
}

type DeleteSourceLocationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSourceLocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSourceLocationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSourceLocationResponse) GoString() string {
	return s.String()
}

func (s *DeleteSourceLocationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSourceLocationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSourceLocationResponse) GetBody() *DeleteSourceLocationResponseBody {
	return s.Body
}

func (s *DeleteSourceLocationResponse) SetHeaders(v map[string]*string) *DeleteSourceLocationResponse {
	s.Headers = v
	return s
}

func (s *DeleteSourceLocationResponse) SetStatusCode(v int32) *DeleteSourceLocationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSourceLocationResponse) SetBody(v *DeleteSourceLocationResponseBody) *DeleteSourceLocationResponse {
	s.Body = v
	return s
}

func (s *DeleteSourceLocationResponse) Validate() error {
	return dara.Validate(s)
}
