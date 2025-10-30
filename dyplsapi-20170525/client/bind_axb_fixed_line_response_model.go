// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxbFixedLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindAxbFixedLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindAxbFixedLineResponse
	GetStatusCode() *int32
	SetBody(v *BindAxbFixedLineResponseBody) *BindAxbFixedLineResponse
	GetBody() *BindAxbFixedLineResponseBody
}

type BindAxbFixedLineResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAxbFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAxbFixedLineResponse) String() string {
	return dara.Prettify(s)
}

func (s BindAxbFixedLineResponse) GoString() string {
	return s.String()
}

func (s *BindAxbFixedLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindAxbFixedLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindAxbFixedLineResponse) GetBody() *BindAxbFixedLineResponseBody {
	return s.Body
}

func (s *BindAxbFixedLineResponse) SetHeaders(v map[string]*string) *BindAxbFixedLineResponse {
	s.Headers = v
	return s
}

func (s *BindAxbFixedLineResponse) SetStatusCode(v int32) *BindAxbFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAxbFixedLineResponse) SetBody(v *BindAxbFixedLineResponseBody) *BindAxbFixedLineResponse {
	s.Body = v
	return s
}

func (s *BindAxbFixedLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
