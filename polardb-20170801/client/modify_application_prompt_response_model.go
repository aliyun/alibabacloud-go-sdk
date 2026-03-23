// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationPromptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApplicationPromptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApplicationPromptResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApplicationPromptResponseBody) *ModifyApplicationPromptResponse
	GetBody() *ModifyApplicationPromptResponseBody
}

type ModifyApplicationPromptResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApplicationPromptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApplicationPromptResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationPromptResponse) GoString() string {
	return s.String()
}

func (s *ModifyApplicationPromptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApplicationPromptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApplicationPromptResponse) GetBody() *ModifyApplicationPromptResponseBody {
	return s.Body
}

func (s *ModifyApplicationPromptResponse) SetHeaders(v map[string]*string) *ModifyApplicationPromptResponse {
	s.Headers = v
	return s
}

func (s *ModifyApplicationPromptResponse) SetStatusCode(v int32) *ModifyApplicationPromptResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApplicationPromptResponse) SetBody(v *ModifyApplicationPromptResponseBody) *ModifyApplicationPromptResponse {
	s.Body = v
	return s
}

func (s *ModifyApplicationPromptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
