// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskCopiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTaskCopiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTaskCopiesResponse
	GetStatusCode() *int32
	SetBody(v *GetTaskCopiesResponseBody) *GetTaskCopiesResponse
	GetBody() *GetTaskCopiesResponseBody
}

type GetTaskCopiesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskCopiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskCopiesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTaskCopiesResponse) GoString() string {
	return s.String()
}

func (s *GetTaskCopiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTaskCopiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTaskCopiesResponse) GetBody() *GetTaskCopiesResponseBody {
	return s.Body
}

func (s *GetTaskCopiesResponse) SetHeaders(v map[string]*string) *GetTaskCopiesResponse {
	s.Headers = v
	return s
}

func (s *GetTaskCopiesResponse) SetStatusCode(v int32) *GetTaskCopiesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskCopiesResponse) SetBody(v *GetTaskCopiesResponseBody) *GetTaskCopiesResponse {
	s.Body = v
	return s
}

func (s *GetTaskCopiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
