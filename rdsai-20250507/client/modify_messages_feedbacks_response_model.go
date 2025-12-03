// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMessagesFeedbacksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMessagesFeedbacksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMessagesFeedbacksResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMessagesFeedbacksResponseBody) *ModifyMessagesFeedbacksResponse
	GetBody() *ModifyMessagesFeedbacksResponseBody
}

type ModifyMessagesFeedbacksResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMessagesFeedbacksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMessagesFeedbacksResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMessagesFeedbacksResponse) GoString() string {
	return s.String()
}

func (s *ModifyMessagesFeedbacksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMessagesFeedbacksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMessagesFeedbacksResponse) GetBody() *ModifyMessagesFeedbacksResponseBody {
	return s.Body
}

func (s *ModifyMessagesFeedbacksResponse) SetHeaders(v map[string]*string) *ModifyMessagesFeedbacksResponse {
	s.Headers = v
	return s
}

func (s *ModifyMessagesFeedbacksResponse) SetStatusCode(v int32) *ModifyMessagesFeedbacksResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMessagesFeedbacksResponse) SetBody(v *ModifyMessagesFeedbacksResponseBody) *ModifyMessagesFeedbacksResponse {
	s.Body = v
	return s
}

func (s *ModifyMessagesFeedbacksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
