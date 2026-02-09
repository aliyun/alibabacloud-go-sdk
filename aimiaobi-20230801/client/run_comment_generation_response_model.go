// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCommentGenerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunCommentGenerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunCommentGenerationResponse
	GetStatusCode() *int32
	SetId(v string) *RunCommentGenerationResponse
	GetId() *string
	SetEvent(v string) *RunCommentGenerationResponse
	GetEvent() *string
	SetBody(v *RunCommentGenerationResponseBody) *RunCommentGenerationResponse
	GetBody() *RunCommentGenerationResponseBody
}

type RunCommentGenerationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                           `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                           `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunCommentGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCommentGenerationResponse) String() string {
	return dara.Prettify(s)
}

func (s RunCommentGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunCommentGenerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunCommentGenerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunCommentGenerationResponse) GetId() *string {
	return s.Id
}

func (s *RunCommentGenerationResponse) GetEvent() *string {
	return s.Event
}

func (s *RunCommentGenerationResponse) GetBody() *RunCommentGenerationResponseBody {
	return s.Body
}

func (s *RunCommentGenerationResponse) SetHeaders(v map[string]*string) *RunCommentGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunCommentGenerationResponse) SetStatusCode(v int32) *RunCommentGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCommentGenerationResponse) SetId(v string) *RunCommentGenerationResponse {
	s.Id = &v
	return s
}

func (s *RunCommentGenerationResponse) SetEvent(v string) *RunCommentGenerationResponse {
	s.Event = &v
	return s
}

func (s *RunCommentGenerationResponse) SetBody(v *RunCommentGenerationResponseBody) *RunCommentGenerationResponse {
	s.Body = v
	return s
}

func (s *RunCommentGenerationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
