// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUnionTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUnionTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUnionTaskListResponse
	GetStatusCode() *int32
	SetBody(v *QueryUnionTaskListResponseBody) *QueryUnionTaskListResponse
	GetBody() *QueryUnionTaskListResponseBody
}

type QueryUnionTaskListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUnionTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUnionTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUnionTaskListResponse) GoString() string {
	return s.String()
}

func (s *QueryUnionTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUnionTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUnionTaskListResponse) GetBody() *QueryUnionTaskListResponseBody {
	return s.Body
}

func (s *QueryUnionTaskListResponse) SetHeaders(v map[string]*string) *QueryUnionTaskListResponse {
	s.Headers = v
	return s
}

func (s *QueryUnionTaskListResponse) SetStatusCode(v int32) *QueryUnionTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUnionTaskListResponse) SetBody(v *QueryUnionTaskListResponseBody) *QueryUnionTaskListResponse {
	s.Body = v
	return s
}

func (s *QueryUnionTaskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
