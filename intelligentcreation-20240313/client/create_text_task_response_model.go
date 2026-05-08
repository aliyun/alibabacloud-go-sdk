// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTextTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTextTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTextTaskResponse
	GetStatusCode() *int32
	SetBody(v *TextTaskResult) *CreateTextTaskResponse
	GetBody() *TextTaskResult
}

type CreateTextTaskResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextTaskResult    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTextTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTextTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTextTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTextTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTextTaskResponse) GetBody() *TextTaskResult {
	return s.Body
}

func (s *CreateTextTaskResponse) SetHeaders(v map[string]*string) *CreateTextTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTextTaskResponse) SetStatusCode(v int32) *CreateTextTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTextTaskResponse) SetBody(v *TextTaskResult) *CreateTextTaskResponse {
	s.Body = v
	return s
}

func (s *CreateTextTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
