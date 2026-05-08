// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTextTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTextTaskResponse
	GetStatusCode() *int32
	SetBody(v *TextTaskResult) *GetTextTaskResponse
	GetBody() *TextTaskResult
}

type GetTextTaskResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextTaskResult    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTextTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTextTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTextTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTextTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTextTaskResponse) GetBody() *TextTaskResult {
	return s.Body
}

func (s *GetTextTaskResponse) SetHeaders(v map[string]*string) *GetTextTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTextTaskResponse) SetStatusCode(v int32) *GetTextTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTextTaskResponse) SetBody(v *TextTaskResult) *GetTextTaskResponse {
	s.Body = v
	return s
}

func (s *GetTextTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
