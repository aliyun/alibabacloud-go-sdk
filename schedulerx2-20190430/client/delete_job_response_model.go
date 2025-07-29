// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteJobResponseBody) *DeleteJobResponse
	GetBody() *DeleteJobResponseBody
}

type DeleteJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteJobResponse) GetBody() *DeleteJobResponseBody {
	return s.Body
}

func (s *DeleteJobResponse) SetHeaders(v map[string]*string) *DeleteJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobResponse) SetStatusCode(v int32) *DeleteJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteJobResponse) SetBody(v *DeleteJobResponseBody) *DeleteJobResponse {
	s.Body = v
	return s
}

func (s *DeleteJobResponse) Validate() error {
	return dara.Validate(s)
}
