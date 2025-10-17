// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunOcrParseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunOcrParseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunOcrParseResponse
	GetStatusCode() *int32
	SetBody(v *RunOcrParseResponseBody) *RunOcrParseResponse
	GetBody() *RunOcrParseResponseBody
}

type RunOcrParseResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunOcrParseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunOcrParseResponse) String() string {
	return dara.Prettify(s)
}

func (s RunOcrParseResponse) GoString() string {
	return s.String()
}

func (s *RunOcrParseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunOcrParseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunOcrParseResponse) GetBody() *RunOcrParseResponseBody {
	return s.Body
}

func (s *RunOcrParseResponse) SetHeaders(v map[string]*string) *RunOcrParseResponse {
	s.Headers = v
	return s
}

func (s *RunOcrParseResponse) SetStatusCode(v int32) *RunOcrParseResponse {
	s.StatusCode = &v
	return s
}

func (s *RunOcrParseResponse) SetBody(v *RunOcrParseResponseBody) *RunOcrParseResponse {
	s.Body = v
	return s
}

func (s *RunOcrParseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
