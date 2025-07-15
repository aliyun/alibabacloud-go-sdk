// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFailoverTestJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFailoverTestJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFailoverTestJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFailoverTestJobResponseBody) *DeleteFailoverTestJobResponse
	GetBody() *DeleteFailoverTestJobResponseBody
}

type DeleteFailoverTestJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFailoverTestJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFailoverTestJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFailoverTestJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteFailoverTestJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFailoverTestJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFailoverTestJobResponse) GetBody() *DeleteFailoverTestJobResponseBody {
	return s.Body
}

func (s *DeleteFailoverTestJobResponse) SetHeaders(v map[string]*string) *DeleteFailoverTestJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteFailoverTestJobResponse) SetStatusCode(v int32) *DeleteFailoverTestJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFailoverTestJobResponse) SetBody(v *DeleteFailoverTestJobResponseBody) *DeleteFailoverTestJobResponse {
	s.Body = v
	return s
}

func (s *DeleteFailoverTestJobResponse) Validate() error {
	return dara.Validate(s)
}
