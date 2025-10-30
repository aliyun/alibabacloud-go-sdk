// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWmExtractTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWmExtractTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWmExtractTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateWmExtractTaskResponseBody) *CreateWmExtractTaskResponse
	GetBody() *CreateWmExtractTaskResponseBody
}

type CreateWmExtractTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWmExtractTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWmExtractTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWmExtractTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateWmExtractTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWmExtractTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWmExtractTaskResponse) GetBody() *CreateWmExtractTaskResponseBody {
	return s.Body
}

func (s *CreateWmExtractTaskResponse) SetHeaders(v map[string]*string) *CreateWmExtractTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateWmExtractTaskResponse) SetStatusCode(v int32) *CreateWmExtractTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWmExtractTaskResponse) SetBody(v *CreateWmExtractTaskResponseBody) *CreateWmExtractTaskResponse {
	s.Body = v
	return s
}

func (s *CreateWmExtractTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
