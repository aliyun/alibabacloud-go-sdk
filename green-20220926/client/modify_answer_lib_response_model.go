// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAnswerLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAnswerLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAnswerLibResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAnswerLibResponseBody) *ModifyAnswerLibResponse
	GetBody() *ModifyAnswerLibResponseBody
}

type ModifyAnswerLibResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAnswerLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAnswerLibResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAnswerLibResponse) GoString() string {
	return s.String()
}

func (s *ModifyAnswerLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAnswerLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAnswerLibResponse) GetBody() *ModifyAnswerLibResponseBody {
	return s.Body
}

func (s *ModifyAnswerLibResponse) SetHeaders(v map[string]*string) *ModifyAnswerLibResponse {
	s.Headers = v
	return s
}

func (s *ModifyAnswerLibResponse) SetStatusCode(v int32) *ModifyAnswerLibResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAnswerLibResponse) SetBody(v *ModifyAnswerLibResponseBody) *ModifyAnswerLibResponse {
	s.Body = v
	return s
}

func (s *ModifyAnswerLibResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
