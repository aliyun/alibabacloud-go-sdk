// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveMessageAppCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLiveMessageAppCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLiveMessageAppCallbackResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLiveMessageAppCallbackResponseBody) *ModifyLiveMessageAppCallbackResponse
	GetBody() *ModifyLiveMessageAppCallbackResponseBody
}

type ModifyLiveMessageAppCallbackResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLiveMessageAppCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLiveMessageAppCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageAppCallbackResponse) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageAppCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLiveMessageAppCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLiveMessageAppCallbackResponse) GetBody() *ModifyLiveMessageAppCallbackResponseBody {
	return s.Body
}

func (s *ModifyLiveMessageAppCallbackResponse) SetHeaders(v map[string]*string) *ModifyLiveMessageAppCallbackResponse {
	s.Headers = v
	return s
}

func (s *ModifyLiveMessageAppCallbackResponse) SetStatusCode(v int32) *ModifyLiveMessageAppCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLiveMessageAppCallbackResponse) SetBody(v *ModifyLiveMessageAppCallbackResponseBody) *ModifyLiveMessageAppCallbackResponse {
	s.Body = v
	return s
}

func (s *ModifyLiveMessageAppCallbackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
