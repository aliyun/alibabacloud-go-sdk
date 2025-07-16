// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteResourceLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteResourceLogResponse
	GetStatusCode() *int32
	SetBody(v *DeleteResourceLogResponseBody) *DeleteResourceLogResponse
	GetBody() *DeleteResourceLogResponseBody
}

type DeleteResourceLogResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceLogResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteResourceLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteResourceLogResponse) GetBody() *DeleteResourceLogResponseBody {
	return s.Body
}

func (s *DeleteResourceLogResponse) SetHeaders(v map[string]*string) *DeleteResourceLogResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceLogResponse) SetStatusCode(v int32) *DeleteResourceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceLogResponse) SetBody(v *DeleteResourceLogResponseBody) *DeleteResourceLogResponse {
	s.Body = v
	return s
}

func (s *DeleteResourceLogResponse) Validate() error {
	return dara.Validate(s)
}
