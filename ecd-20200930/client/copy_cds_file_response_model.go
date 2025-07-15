// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyCdsFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyCdsFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyCdsFileResponse
	GetStatusCode() *int32
	SetBody(v *CopyCdsFileResponseBody) *CopyCdsFileResponse
	GetBody() *CopyCdsFileResponseBody
}

type CopyCdsFileResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyCdsFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyCdsFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyCdsFileResponse) GoString() string {
	return s.String()
}

func (s *CopyCdsFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyCdsFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyCdsFileResponse) GetBody() *CopyCdsFileResponseBody {
	return s.Body
}

func (s *CopyCdsFileResponse) SetHeaders(v map[string]*string) *CopyCdsFileResponse {
	s.Headers = v
	return s
}

func (s *CopyCdsFileResponse) SetStatusCode(v int32) *CopyCdsFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyCdsFileResponse) SetBody(v *CopyCdsFileResponseBody) *CopyCdsFileResponse {
	s.Body = v
	return s
}

func (s *CopyCdsFileResponse) Validate() error {
	return dara.Validate(s)
}
