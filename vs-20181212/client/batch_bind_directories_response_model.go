// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchBindDirectoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchBindDirectoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchBindDirectoriesResponse
	GetStatusCode() *int32
	SetBody(v *BatchBindDirectoriesResponseBody) *BatchBindDirectoriesResponse
	GetBody() *BatchBindDirectoriesResponseBody
}

type BatchBindDirectoriesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchBindDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchBindDirectoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchBindDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *BatchBindDirectoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchBindDirectoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchBindDirectoriesResponse) GetBody() *BatchBindDirectoriesResponseBody {
	return s.Body
}

func (s *BatchBindDirectoriesResponse) SetHeaders(v map[string]*string) *BatchBindDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *BatchBindDirectoriesResponse) SetStatusCode(v int32) *BatchBindDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchBindDirectoriesResponse) SetBody(v *BatchBindDirectoriesResponseBody) *BatchBindDirectoriesResponse {
	s.Body = v
	return s
}

func (s *BatchBindDirectoriesResponse) Validate() error {
	return dara.Validate(s)
}
