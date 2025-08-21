// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCodeEnhanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCodeEnhanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCodeEnhanceResponse
	GetStatusCode() *int32
	SetBody(v *GetCodeEnhanceResponseBody) *GetCodeEnhanceResponse
	GetBody() *GetCodeEnhanceResponseBody
}

type GetCodeEnhanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCodeEnhanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCodeEnhanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCodeEnhanceResponse) GoString() string {
	return s.String()
}

func (s *GetCodeEnhanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCodeEnhanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCodeEnhanceResponse) GetBody() *GetCodeEnhanceResponseBody {
	return s.Body
}

func (s *GetCodeEnhanceResponse) SetHeaders(v map[string]*string) *GetCodeEnhanceResponse {
	s.Headers = v
	return s
}

func (s *GetCodeEnhanceResponse) SetStatusCode(v int32) *GetCodeEnhanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCodeEnhanceResponse) SetBody(v *GetCodeEnhanceResponseBody) *GetCodeEnhanceResponse {
	s.Body = v
	return s
}

func (s *GetCodeEnhanceResponse) Validate() error {
	return dara.Validate(s)
}
