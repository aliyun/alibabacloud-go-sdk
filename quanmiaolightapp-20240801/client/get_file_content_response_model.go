// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileContentResponse
	GetStatusCode() *int32
	SetBody(v *GetFileContentResponseBody) *GetFileContentResponse
	GetBody() *GetFileContentResponseBody
}

type GetFileContentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileContentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileContentResponse) GoString() string {
	return s.String()
}

func (s *GetFileContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileContentResponse) GetBody() *GetFileContentResponseBody {
	return s.Body
}

func (s *GetFileContentResponse) SetHeaders(v map[string]*string) *GetFileContentResponse {
	s.Headers = v
	return s
}

func (s *GetFileContentResponse) SetStatusCode(v int32) *GetFileContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileContentResponse) SetBody(v *GetFileContentResponseBody) *GetFileContentResponse {
	s.Body = v
	return s
}

func (s *GetFileContentResponse) Validate() error {
	return dara.Validate(s)
}
