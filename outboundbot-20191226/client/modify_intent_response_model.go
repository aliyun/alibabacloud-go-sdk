// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIntentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIntentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIntentResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIntentResponseBody) *ModifyIntentResponse
	GetBody() *ModifyIntentResponseBody
}

type ModifyIntentResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIntentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIntentResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIntentResponse) GoString() string {
	return s.String()
}

func (s *ModifyIntentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIntentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIntentResponse) GetBody() *ModifyIntentResponseBody {
	return s.Body
}

func (s *ModifyIntentResponse) SetHeaders(v map[string]*string) *ModifyIntentResponse {
	s.Headers = v
	return s
}

func (s *ModifyIntentResponse) SetStatusCode(v int32) *ModifyIntentResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIntentResponse) SetBody(v *ModifyIntentResponseBody) *ModifyIntentResponse {
	s.Body = v
	return s
}

func (s *ModifyIntentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
