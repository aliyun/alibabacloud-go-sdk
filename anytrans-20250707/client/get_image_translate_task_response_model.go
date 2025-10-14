// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageTranslateTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageTranslateTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageTranslateTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetImageTranslateTaskResponseBody) *GetImageTranslateTaskResponse
	GetBody() *GetImageTranslateTaskResponseBody
}

type GetImageTranslateTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageTranslateTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageTranslateTaskResponse) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageTranslateTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageTranslateTaskResponse) GetBody() *GetImageTranslateTaskResponseBody {
	return s.Body
}

func (s *GetImageTranslateTaskResponse) SetHeaders(v map[string]*string) *GetImageTranslateTaskResponse {
	s.Headers = v
	return s
}

func (s *GetImageTranslateTaskResponse) SetStatusCode(v int32) *GetImageTranslateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageTranslateTaskResponse) SetBody(v *GetImageTranslateTaskResponseBody) *GetImageTranslateTaskResponse {
	s.Body = v
	return s
}

func (s *GetImageTranslateTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
