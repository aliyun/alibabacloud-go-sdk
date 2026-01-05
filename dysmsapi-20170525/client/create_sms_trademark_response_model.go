// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsTrademarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSmsTrademarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSmsTrademarkResponse
	GetStatusCode() *int32
	SetBody(v *CreateSmsTrademarkResponseBody) *CreateSmsTrademarkResponse
	GetBody() *CreateSmsTrademarkResponseBody
}

type CreateSmsTrademarkResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSmsTrademarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSmsTrademarkResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsTrademarkResponse) GoString() string {
	return s.String()
}

func (s *CreateSmsTrademarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSmsTrademarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSmsTrademarkResponse) GetBody() *CreateSmsTrademarkResponseBody {
	return s.Body
}

func (s *CreateSmsTrademarkResponse) SetHeaders(v map[string]*string) *CreateSmsTrademarkResponse {
	s.Headers = v
	return s
}

func (s *CreateSmsTrademarkResponse) SetStatusCode(v int32) *CreateSmsTrademarkResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSmsTrademarkResponse) SetBody(v *CreateSmsTrademarkResponseBody) *CreateSmsTrademarkResponse {
	s.Body = v
	return s
}

func (s *CreateSmsTrademarkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
