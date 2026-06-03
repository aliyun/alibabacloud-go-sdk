// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTaskListResponse
	GetStatusCode() *int32
	SetBody(v *QueryTaskListResponseBody) *QueryTaskListResponse
	GetBody() *QueryTaskListResponseBody
}

type QueryTaskListResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskListResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTaskListResponse) GetBody() *QueryTaskListResponseBody {
	return s.Body
}

func (s *QueryTaskListResponse) SetHeaders(v map[string]*string) *QueryTaskListResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskListResponse) SetStatusCode(v int32) *QueryTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTaskListResponse) SetBody(v *QueryTaskListResponseBody) *QueryTaskListResponse {
	s.Body = v
	return s
}

func (s *QueryTaskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
