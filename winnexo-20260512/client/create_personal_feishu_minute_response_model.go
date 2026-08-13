// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalFeishuMinuteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePersonalFeishuMinuteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePersonalFeishuMinuteResponse
	GetStatusCode() *int32
	SetBody(v *CreatePersonalFeishuMinuteResponseBody) *CreatePersonalFeishuMinuteResponse
	GetBody() *CreatePersonalFeishuMinuteResponseBody
}

type CreatePersonalFeishuMinuteResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePersonalFeishuMinuteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePersonalFeishuMinuteResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalFeishuMinuteResponse) GoString() string {
	return s.String()
}

func (s *CreatePersonalFeishuMinuteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePersonalFeishuMinuteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePersonalFeishuMinuteResponse) GetBody() *CreatePersonalFeishuMinuteResponseBody {
	return s.Body
}

func (s *CreatePersonalFeishuMinuteResponse) SetHeaders(v map[string]*string) *CreatePersonalFeishuMinuteResponse {
	s.Headers = v
	return s
}

func (s *CreatePersonalFeishuMinuteResponse) SetStatusCode(v int32) *CreatePersonalFeishuMinuteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePersonalFeishuMinuteResponse) SetBody(v *CreatePersonalFeishuMinuteResponseBody) *CreatePersonalFeishuMinuteResponse {
	s.Body = v
	return s
}

func (s *CreatePersonalFeishuMinuteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
