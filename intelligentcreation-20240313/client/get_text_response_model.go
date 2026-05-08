// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTextResponse
	GetStatusCode() *int32
	SetBody(v *TextResult) *GetTextResponse
	GetBody() *TextResult
}

type GetTextResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextResult        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTextResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTextResponse) GoString() string {
	return s.String()
}

func (s *GetTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTextResponse) GetBody() *TextResult {
	return s.Body
}

func (s *GetTextResponse) SetHeaders(v map[string]*string) *GetTextResponse {
	s.Headers = v
	return s
}

func (s *GetTextResponse) SetStatusCode(v int32) *GetTextResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTextResponse) SetBody(v *TextResult) *GetTextResponse {
	s.Body = v
	return s
}

func (s *GetTextResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
