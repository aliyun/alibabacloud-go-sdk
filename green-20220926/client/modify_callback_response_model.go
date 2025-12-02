// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCallbackResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCallbackResponseBody) *ModifyCallbackResponse
	GetBody() *ModifyCallbackResponseBody
}

type ModifyCallbackResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCallbackResponse) GoString() string {
	return s.String()
}

func (s *ModifyCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCallbackResponse) GetBody() *ModifyCallbackResponseBody {
	return s.Body
}

func (s *ModifyCallbackResponse) SetHeaders(v map[string]*string) *ModifyCallbackResponse {
	s.Headers = v
	return s
}

func (s *ModifyCallbackResponse) SetStatusCode(v int32) *ModifyCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCallbackResponse) SetBody(v *ModifyCallbackResponseBody) *ModifyCallbackResponse {
	s.Body = v
	return s
}

func (s *ModifyCallbackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
