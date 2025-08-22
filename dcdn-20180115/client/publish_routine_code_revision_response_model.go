// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishRoutineCodeRevisionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishRoutineCodeRevisionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishRoutineCodeRevisionResponse
	GetStatusCode() *int32
	SetBody(v *PublishRoutineCodeRevisionResponseBody) *PublishRoutineCodeRevisionResponse
	GetBody() *PublishRoutineCodeRevisionResponseBody
}

type PublishRoutineCodeRevisionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishRoutineCodeRevisionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishRoutineCodeRevisionResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishRoutineCodeRevisionResponse) GoString() string {
	return s.String()
}

func (s *PublishRoutineCodeRevisionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishRoutineCodeRevisionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishRoutineCodeRevisionResponse) GetBody() *PublishRoutineCodeRevisionResponseBody {
	return s.Body
}

func (s *PublishRoutineCodeRevisionResponse) SetHeaders(v map[string]*string) *PublishRoutineCodeRevisionResponse {
	s.Headers = v
	return s
}

func (s *PublishRoutineCodeRevisionResponse) SetStatusCode(v int32) *PublishRoutineCodeRevisionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishRoutineCodeRevisionResponse) SetBody(v *PublishRoutineCodeRevisionResponseBody) *PublishRoutineCodeRevisionResponse {
	s.Body = v
	return s
}

func (s *PublishRoutineCodeRevisionResponse) Validate() error {
	return dara.Validate(s)
}
