// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalFeishuDocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePersonalFeishuDocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePersonalFeishuDocResponse
	GetStatusCode() *int32
	SetBody(v *CreatePersonalFeishuDocResponseBody) *CreatePersonalFeishuDocResponse
	GetBody() *CreatePersonalFeishuDocResponseBody
}

type CreatePersonalFeishuDocResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePersonalFeishuDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePersonalFeishuDocResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalFeishuDocResponse) GoString() string {
	return s.String()
}

func (s *CreatePersonalFeishuDocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePersonalFeishuDocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePersonalFeishuDocResponse) GetBody() *CreatePersonalFeishuDocResponseBody {
	return s.Body
}

func (s *CreatePersonalFeishuDocResponse) SetHeaders(v map[string]*string) *CreatePersonalFeishuDocResponse {
	s.Headers = v
	return s
}

func (s *CreatePersonalFeishuDocResponse) SetStatusCode(v int32) *CreatePersonalFeishuDocResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePersonalFeishuDocResponse) SetBody(v *CreatePersonalFeishuDocResponseBody) *CreatePersonalFeishuDocResponse {
	s.Body = v
	return s
}

func (s *CreatePersonalFeishuDocResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
