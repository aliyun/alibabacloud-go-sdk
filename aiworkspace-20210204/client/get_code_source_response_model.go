// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCodeSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCodeSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCodeSourceResponse
	GetStatusCode() *int32
	SetBody(v *GetCodeSourceResponseBody) *GetCodeSourceResponse
	GetBody() *GetCodeSourceResponseBody
}

type GetCodeSourceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCodeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCodeSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCodeSourceResponse) GoString() string {
	return s.String()
}

func (s *GetCodeSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCodeSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCodeSourceResponse) GetBody() *GetCodeSourceResponseBody {
	return s.Body
}

func (s *GetCodeSourceResponse) SetHeaders(v map[string]*string) *GetCodeSourceResponse {
	s.Headers = v
	return s
}

func (s *GetCodeSourceResponse) SetStatusCode(v int32) *GetCodeSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCodeSourceResponse) SetBody(v *GetCodeSourceResponseBody) *GetCodeSourceResponse {
	s.Body = v
	return s
}

func (s *GetCodeSourceResponse) Validate() error {
	return dara.Validate(s)
}
