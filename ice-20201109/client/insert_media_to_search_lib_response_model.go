// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertMediaToSearchLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertMediaToSearchLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertMediaToSearchLibResponse
	GetStatusCode() *int32
	SetBody(v *InsertMediaToSearchLibResponseBody) *InsertMediaToSearchLibResponse
	GetBody() *InsertMediaToSearchLibResponseBody
}

type InsertMediaToSearchLibResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertMediaToSearchLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertMediaToSearchLibResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertMediaToSearchLibResponse) GoString() string {
	return s.String()
}

func (s *InsertMediaToSearchLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertMediaToSearchLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertMediaToSearchLibResponse) GetBody() *InsertMediaToSearchLibResponseBody {
	return s.Body
}

func (s *InsertMediaToSearchLibResponse) SetHeaders(v map[string]*string) *InsertMediaToSearchLibResponse {
	s.Headers = v
	return s
}

func (s *InsertMediaToSearchLibResponse) SetStatusCode(v int32) *InsertMediaToSearchLibResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertMediaToSearchLibResponse) SetBody(v *InsertMediaToSearchLibResponseBody) *InsertMediaToSearchLibResponse {
	s.Body = v
	return s
}

func (s *InsertMediaToSearchLibResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
