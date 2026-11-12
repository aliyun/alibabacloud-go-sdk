// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitImageResponse
	GetStatusCode() *int32
	SetBody(v *SubmitImageResponseBody) *SubmitImageResponse
	GetBody() *SubmitImageResponseBody
}

type SubmitImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitImageResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageResponse) GoString() string {
	return s.String()
}

func (s *SubmitImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitImageResponse) GetBody() *SubmitImageResponseBody {
	return s.Body
}

func (s *SubmitImageResponse) SetHeaders(v map[string]*string) *SubmitImageResponse {
	s.Headers = v
	return s
}

func (s *SubmitImageResponse) SetStatusCode(v int32) *SubmitImageResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitImageResponse) SetBody(v *SubmitImageResponseBody) *SubmitImageResponse {
	s.Body = v
	return s
}

func (s *SubmitImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
