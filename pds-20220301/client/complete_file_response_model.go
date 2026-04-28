// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompleteFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompleteFileResponse
	GetStatusCode() *int32
	SetBody(v *File) *CompleteFileResponse
	GetBody() *File
}

type CompleteFileResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *File              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompleteFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CompleteFileResponse) GoString() string {
	return s.String()
}

func (s *CompleteFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompleteFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompleteFileResponse) GetBody() *File {
	return s.Body
}

func (s *CompleteFileResponse) SetHeaders(v map[string]*string) *CompleteFileResponse {
	s.Headers = v
	return s
}

func (s *CompleteFileResponse) SetStatusCode(v int32) *CompleteFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CompleteFileResponse) SetBody(v *File) *CompleteFileResponse {
	s.Body = v
	return s
}

func (s *CompleteFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
