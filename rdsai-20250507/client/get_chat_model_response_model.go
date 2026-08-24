// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChatModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChatModelResponse
	GetStatusCode() *int32
	SetBody(v *GetChatModelResponseBody) *GetChatModelResponse
	GetBody() *GetChatModelResponseBody
}

type GetChatModelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatModelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChatModelResponse) GoString() string {
	return s.String()
}

func (s *GetChatModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChatModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChatModelResponse) GetBody() *GetChatModelResponseBody {
	return s.Body
}

func (s *GetChatModelResponse) SetHeaders(v map[string]*string) *GetChatModelResponse {
	s.Headers = v
	return s
}

func (s *GetChatModelResponse) SetStatusCode(v int32) *GetChatModelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatModelResponse) SetBody(v *GetChatModelResponseBody) *GetChatModelResponse {
	s.Body = v
	return s
}

func (s *GetChatModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
