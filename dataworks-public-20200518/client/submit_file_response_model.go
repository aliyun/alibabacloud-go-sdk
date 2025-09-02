// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitFileResponse
	GetStatusCode() *int32
	SetBody(v *SubmitFileResponseBody) *SubmitFileResponse
	GetBody() *SubmitFileResponseBody
}

type SubmitFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitFileResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitFileResponse) GoString() string {
	return s.String()
}

func (s *SubmitFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitFileResponse) GetBody() *SubmitFileResponseBody {
	return s.Body
}

func (s *SubmitFileResponse) SetHeaders(v map[string]*string) *SubmitFileResponse {
	s.Headers = v
	return s
}

func (s *SubmitFileResponse) SetStatusCode(v int32) *SubmitFileResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitFileResponse) SetBody(v *SubmitFileResponseBody) *SubmitFileResponse {
	s.Body = v
	return s
}

func (s *SubmitFileResponse) Validate() error {
	return dara.Validate(s)
}
