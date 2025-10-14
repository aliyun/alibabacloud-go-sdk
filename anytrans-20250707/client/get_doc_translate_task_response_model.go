// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocTranslateTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocTranslateTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocTranslateTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetDocTranslateTaskResponseBody) *GetDocTranslateTaskResponse
	GetBody() *GetDocTranslateTaskResponseBody
}

type GetDocTranslateTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocTranslateTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocTranslateTaskResponse) GoString() string {
	return s.String()
}

func (s *GetDocTranslateTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocTranslateTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocTranslateTaskResponse) GetBody() *GetDocTranslateTaskResponseBody {
	return s.Body
}

func (s *GetDocTranslateTaskResponse) SetHeaders(v map[string]*string) *GetDocTranslateTaskResponse {
	s.Headers = v
	return s
}

func (s *GetDocTranslateTaskResponse) SetStatusCode(v int32) *GetDocTranslateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocTranslateTaskResponse) SetBody(v *GetDocTranslateTaskResponseBody) *GetDocTranslateTaskResponse {
	s.Body = v
	return s
}

func (s *GetDocTranslateTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
