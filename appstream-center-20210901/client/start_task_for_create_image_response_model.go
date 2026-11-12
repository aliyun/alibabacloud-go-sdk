// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTaskForCreateImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartTaskForCreateImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartTaskForCreateImageResponse
	GetStatusCode() *int32
	SetBody(v *StartTaskForCreateImageResponseBody) *StartTaskForCreateImageResponse
	GetBody() *StartTaskForCreateImageResponseBody
}

type StartTaskForCreateImageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartTaskForCreateImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartTaskForCreateImageResponse) String() string {
	return dara.Prettify(s)
}

func (s StartTaskForCreateImageResponse) GoString() string {
	return s.String()
}

func (s *StartTaskForCreateImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartTaskForCreateImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartTaskForCreateImageResponse) GetBody() *StartTaskForCreateImageResponseBody {
	return s.Body
}

func (s *StartTaskForCreateImageResponse) SetHeaders(v map[string]*string) *StartTaskForCreateImageResponse {
	s.Headers = v
	return s
}

func (s *StartTaskForCreateImageResponse) SetStatusCode(v int32) *StartTaskForCreateImageResponse {
	s.StatusCode = &v
	return s
}

func (s *StartTaskForCreateImageResponse) SetBody(v *StartTaskForCreateImageResponseBody) *StartTaskForCreateImageResponse {
	s.Body = v
	return s
}

func (s *StartTaskForCreateImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
