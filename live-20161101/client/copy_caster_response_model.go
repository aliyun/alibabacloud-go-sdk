// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyCasterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyCasterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyCasterResponse
	GetStatusCode() *int32
	SetBody(v *CopyCasterResponseBody) *CopyCasterResponse
	GetBody() *CopyCasterResponseBody
}

type CopyCasterResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyCasterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyCasterResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyCasterResponse) GoString() string {
	return s.String()
}

func (s *CopyCasterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyCasterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyCasterResponse) GetBody() *CopyCasterResponseBody {
	return s.Body
}

func (s *CopyCasterResponse) SetHeaders(v map[string]*string) *CopyCasterResponse {
	s.Headers = v
	return s
}

func (s *CopyCasterResponse) SetStatusCode(v int32) *CopyCasterResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyCasterResponse) SetBody(v *CopyCasterResponseBody) *CopyCasterResponse {
	s.Body = v
	return s
}

func (s *CopyCasterResponse) Validate() error {
	return dara.Validate(s)
}
