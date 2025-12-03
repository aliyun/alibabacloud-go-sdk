// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTitleGenerateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTitleGenerateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTitleGenerateResponse
	GetStatusCode() *int32
	SetBody(v *GetTitleGenerateResponseBody) *GetTitleGenerateResponse
	GetBody() *GetTitleGenerateResponseBody
}

type GetTitleGenerateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTitleGenerateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTitleGenerateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTitleGenerateResponse) GoString() string {
	return s.String()
}

func (s *GetTitleGenerateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTitleGenerateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTitleGenerateResponse) GetBody() *GetTitleGenerateResponseBody {
	return s.Body
}

func (s *GetTitleGenerateResponse) SetHeaders(v map[string]*string) *GetTitleGenerateResponse {
	s.Headers = v
	return s
}

func (s *GetTitleGenerateResponse) SetStatusCode(v int32) *GetTitleGenerateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTitleGenerateResponse) SetBody(v *GetTitleGenerateResponseBody) *GetTitleGenerateResponse {
	s.Body = v
	return s
}

func (s *GetTitleGenerateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
