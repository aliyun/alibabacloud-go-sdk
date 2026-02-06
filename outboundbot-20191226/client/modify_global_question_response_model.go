// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalQuestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyGlobalQuestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyGlobalQuestionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyGlobalQuestionResponseBody) *ModifyGlobalQuestionResponse
	GetBody() *ModifyGlobalQuestionResponseBody
}

type ModifyGlobalQuestionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyGlobalQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyGlobalQuestionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalQuestionResponse) GoString() string {
	return s.String()
}

func (s *ModifyGlobalQuestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyGlobalQuestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyGlobalQuestionResponse) GetBody() *ModifyGlobalQuestionResponseBody {
	return s.Body
}

func (s *ModifyGlobalQuestionResponse) SetHeaders(v map[string]*string) *ModifyGlobalQuestionResponse {
	s.Headers = v
	return s
}

func (s *ModifyGlobalQuestionResponse) SetStatusCode(v int32) *ModifyGlobalQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGlobalQuestionResponse) SetBody(v *ModifyGlobalQuestionResponseBody) *ModifyGlobalQuestionResponse {
	s.Body = v
	return s
}

func (s *ModifyGlobalQuestionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
