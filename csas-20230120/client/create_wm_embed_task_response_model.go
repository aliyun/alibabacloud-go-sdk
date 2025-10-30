// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWmEmbedTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWmEmbedTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWmEmbedTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateWmEmbedTaskResponseBody) *CreateWmEmbedTaskResponse
	GetBody() *CreateWmEmbedTaskResponseBody
}

type CreateWmEmbedTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWmEmbedTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWmEmbedTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWmEmbedTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWmEmbedTaskResponse) GetBody() *CreateWmEmbedTaskResponseBody {
	return s.Body
}

func (s *CreateWmEmbedTaskResponse) SetHeaders(v map[string]*string) *CreateWmEmbedTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateWmEmbedTaskResponse) SetStatusCode(v int32) *CreateWmEmbedTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWmEmbedTaskResponse) SetBody(v *CreateWmEmbedTaskResponseBody) *CreateWmEmbedTaskResponse {
	s.Body = v
	return s
}

func (s *CreateWmEmbedTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
