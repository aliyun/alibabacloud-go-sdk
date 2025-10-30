// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxnFixedLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindAxnFixedLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindAxnFixedLineResponse
	GetStatusCode() *int32
	SetBody(v *BindAxnFixedLineResponseBody) *BindAxnFixedLineResponse
	GetBody() *BindAxnFixedLineResponseBody
}

type BindAxnFixedLineResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAxnFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAxnFixedLineResponse) String() string {
	return dara.Prettify(s)
}

func (s BindAxnFixedLineResponse) GoString() string {
	return s.String()
}

func (s *BindAxnFixedLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindAxnFixedLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindAxnFixedLineResponse) GetBody() *BindAxnFixedLineResponseBody {
	return s.Body
}

func (s *BindAxnFixedLineResponse) SetHeaders(v map[string]*string) *BindAxnFixedLineResponse {
	s.Headers = v
	return s
}

func (s *BindAxnFixedLineResponse) SetStatusCode(v int32) *BindAxnFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAxnFixedLineResponse) SetBody(v *BindAxnFixedLineResponseBody) *BindAxnFixedLineResponse {
	s.Body = v
	return s
}

func (s *BindAxnFixedLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
