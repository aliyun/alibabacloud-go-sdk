// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUnbindDirectoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchUnbindDirectoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchUnbindDirectoriesResponse
	GetStatusCode() *int32
	SetBody(v *BatchUnbindDirectoriesResponseBody) *BatchUnbindDirectoriesResponse
	GetBody() *BatchUnbindDirectoriesResponseBody
}

type BatchUnbindDirectoriesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUnbindDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUnbindDirectoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchUnbindDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *BatchUnbindDirectoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchUnbindDirectoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchUnbindDirectoriesResponse) GetBody() *BatchUnbindDirectoriesResponseBody {
	return s.Body
}

func (s *BatchUnbindDirectoriesResponse) SetHeaders(v map[string]*string) *BatchUnbindDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *BatchUnbindDirectoriesResponse) SetStatusCode(v int32) *BatchUnbindDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUnbindDirectoriesResponse) SetBody(v *BatchUnbindDirectoriesResponseBody) *BatchUnbindDirectoriesResponse {
	s.Body = v
	return s
}

func (s *BatchUnbindDirectoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
