// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishRoutineCodeVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishRoutineCodeVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishRoutineCodeVersionResponse
	GetStatusCode() *int32
	SetBody(v *PublishRoutineCodeVersionResponseBody) *PublishRoutineCodeVersionResponse
	GetBody() *PublishRoutineCodeVersionResponseBody
}

type PublishRoutineCodeVersionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishRoutineCodeVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishRoutineCodeVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishRoutineCodeVersionResponse) GoString() string {
	return s.String()
}

func (s *PublishRoutineCodeVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishRoutineCodeVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishRoutineCodeVersionResponse) GetBody() *PublishRoutineCodeVersionResponseBody {
	return s.Body
}

func (s *PublishRoutineCodeVersionResponse) SetHeaders(v map[string]*string) *PublishRoutineCodeVersionResponse {
	s.Headers = v
	return s
}

func (s *PublishRoutineCodeVersionResponse) SetStatusCode(v int32) *PublishRoutineCodeVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishRoutineCodeVersionResponse) SetBody(v *PublishRoutineCodeVersionResponseBody) *PublishRoutineCodeVersionResponse {
	s.Body = v
	return s
}

func (s *PublishRoutineCodeVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
