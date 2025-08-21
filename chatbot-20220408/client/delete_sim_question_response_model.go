// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSimQuestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSimQuestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSimQuestionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSimQuestionResponseBody) *DeleteSimQuestionResponse
	GetBody() *DeleteSimQuestionResponseBody
}

type DeleteSimQuestionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSimQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSimQuestionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSimQuestionResponse) GoString() string {
	return s.String()
}

func (s *DeleteSimQuestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSimQuestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSimQuestionResponse) GetBody() *DeleteSimQuestionResponseBody {
	return s.Body
}

func (s *DeleteSimQuestionResponse) SetHeaders(v map[string]*string) *DeleteSimQuestionResponse {
	s.Headers = v
	return s
}

func (s *DeleteSimQuestionResponse) SetStatusCode(v int32) *DeleteSimQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSimQuestionResponse) SetBody(v *DeleteSimQuestionResponseBody) *DeleteSimQuestionResponse {
	s.Body = v
	return s
}

func (s *DeleteSimQuestionResponse) Validate() error {
	return dara.Validate(s)
}
