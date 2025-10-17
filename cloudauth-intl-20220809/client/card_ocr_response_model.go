// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCardOcrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CardOcrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CardOcrResponse
	GetStatusCode() *int32
	SetBody(v *CardOcrResponseBody) *CardOcrResponse
	GetBody() *CardOcrResponseBody
}

type CardOcrResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CardOcrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CardOcrResponse) String() string {
	return dara.Prettify(s)
}

func (s CardOcrResponse) GoString() string {
	return s.String()
}

func (s *CardOcrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CardOcrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CardOcrResponse) GetBody() *CardOcrResponseBody {
	return s.Body
}

func (s *CardOcrResponse) SetHeaders(v map[string]*string) *CardOcrResponse {
	s.Headers = v
	return s
}

func (s *CardOcrResponse) SetStatusCode(v int32) *CardOcrResponse {
	s.StatusCode = &v
	return s
}

func (s *CardOcrResponse) SetBody(v *CardOcrResponseBody) *CardOcrResponse {
	s.Body = v
	return s
}

func (s *CardOcrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
