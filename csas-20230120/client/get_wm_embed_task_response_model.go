// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWmEmbedTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWmEmbedTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWmEmbedTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetWmEmbedTaskResponseBody) *GetWmEmbedTaskResponse
	GetBody() *GetWmEmbedTaskResponseBody
}

type GetWmEmbedTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWmEmbedTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWmEmbedTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWmEmbedTaskResponse) GoString() string {
	return s.String()
}

func (s *GetWmEmbedTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWmEmbedTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWmEmbedTaskResponse) GetBody() *GetWmEmbedTaskResponseBody {
	return s.Body
}

func (s *GetWmEmbedTaskResponse) SetHeaders(v map[string]*string) *GetWmEmbedTaskResponse {
	s.Headers = v
	return s
}

func (s *GetWmEmbedTaskResponse) SetStatusCode(v int32) *GetWmEmbedTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWmEmbedTaskResponse) SetBody(v *GetWmEmbedTaskResponseBody) *GetWmEmbedTaskResponse {
	s.Body = v
	return s
}

func (s *GetWmEmbedTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
