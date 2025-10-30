// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWmExtractTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWmExtractTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWmExtractTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetWmExtractTaskResponseBody) *GetWmExtractTaskResponse
	GetBody() *GetWmExtractTaskResponseBody
}

type GetWmExtractTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWmExtractTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWmExtractTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWmExtractTaskResponse) GoString() string {
	return s.String()
}

func (s *GetWmExtractTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWmExtractTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWmExtractTaskResponse) GetBody() *GetWmExtractTaskResponseBody {
	return s.Body
}

func (s *GetWmExtractTaskResponse) SetHeaders(v map[string]*string) *GetWmExtractTaskResponse {
	s.Headers = v
	return s
}

func (s *GetWmExtractTaskResponse) SetStatusCode(v int32) *GetWmExtractTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWmExtractTaskResponse) SetBody(v *GetWmExtractTaskResponseBody) *GetWmExtractTaskResponse {
	s.Body = v
	return s
}

func (s *GetWmExtractTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
