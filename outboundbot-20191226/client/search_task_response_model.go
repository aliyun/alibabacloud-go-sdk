// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchTaskResponse
	GetStatusCode() *int32
	SetBody(v *SearchTaskResponseBody) *SearchTaskResponse
	GetBody() *SearchTaskResponseBody
}

type SearchTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchTaskResponse) GoString() string {
	return s.String()
}

func (s *SearchTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchTaskResponse) GetBody() *SearchTaskResponseBody {
	return s.Body
}

func (s *SearchTaskResponse) SetHeaders(v map[string]*string) *SearchTaskResponse {
	s.Headers = v
	return s
}

func (s *SearchTaskResponse) SetStatusCode(v int32) *SearchTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTaskResponse) SetBody(v *SearchTaskResponseBody) *SearchTaskResponse {
	s.Body = v
	return s
}

func (s *SearchTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
