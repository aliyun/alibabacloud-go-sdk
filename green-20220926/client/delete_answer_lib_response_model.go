// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAnswerLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAnswerLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAnswerLibResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAnswerLibResponseBody) *DeleteAnswerLibResponse
	GetBody() *DeleteAnswerLibResponseBody
}

type DeleteAnswerLibResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAnswerLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAnswerLibResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAnswerLibResponse) GoString() string {
	return s.String()
}

func (s *DeleteAnswerLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAnswerLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAnswerLibResponse) GetBody() *DeleteAnswerLibResponseBody {
	return s.Body
}

func (s *DeleteAnswerLibResponse) SetHeaders(v map[string]*string) *DeleteAnswerLibResponse {
	s.Headers = v
	return s
}

func (s *DeleteAnswerLibResponse) SetStatusCode(v int32) *DeleteAnswerLibResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAnswerLibResponse) SetBody(v *DeleteAnswerLibResponseBody) *DeleteAnswerLibResponse {
	s.Body = v
	return s
}

func (s *DeleteAnswerLibResponse) Validate() error {
	return dara.Validate(s)
}
