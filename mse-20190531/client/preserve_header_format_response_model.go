// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreserveHeaderFormatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreserveHeaderFormatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreserveHeaderFormatResponse
	GetStatusCode() *int32
	SetBody(v *PreserveHeaderFormatResponseBody) *PreserveHeaderFormatResponse
	GetBody() *PreserveHeaderFormatResponseBody
}

type PreserveHeaderFormatResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreserveHeaderFormatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreserveHeaderFormatResponse) String() string {
	return dara.Prettify(s)
}

func (s PreserveHeaderFormatResponse) GoString() string {
	return s.String()
}

func (s *PreserveHeaderFormatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreserveHeaderFormatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreserveHeaderFormatResponse) GetBody() *PreserveHeaderFormatResponseBody {
	return s.Body
}

func (s *PreserveHeaderFormatResponse) SetHeaders(v map[string]*string) *PreserveHeaderFormatResponse {
	s.Headers = v
	return s
}

func (s *PreserveHeaderFormatResponse) SetStatusCode(v int32) *PreserveHeaderFormatResponse {
	s.StatusCode = &v
	return s
}

func (s *PreserveHeaderFormatResponse) SetBody(v *PreserveHeaderFormatResponseBody) *PreserveHeaderFormatResponse {
	s.Body = v
	return s
}

func (s *PreserveHeaderFormatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
