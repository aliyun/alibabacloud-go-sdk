// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveCdsFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveCdsFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveCdsFileResponse
	GetStatusCode() *int32
	SetBody(v *MoveCdsFileResponseBody) *MoveCdsFileResponse
	GetBody() *MoveCdsFileResponseBody
}

type MoveCdsFileResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveCdsFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveCdsFileResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveCdsFileResponse) GoString() string {
	return s.String()
}

func (s *MoveCdsFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveCdsFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveCdsFileResponse) GetBody() *MoveCdsFileResponseBody {
	return s.Body
}

func (s *MoveCdsFileResponse) SetHeaders(v map[string]*string) *MoveCdsFileResponse {
	s.Headers = v
	return s
}

func (s *MoveCdsFileResponse) SetStatusCode(v int32) *MoveCdsFileResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveCdsFileResponse) SetBody(v *MoveCdsFileResponseBody) *MoveCdsFileResponse {
	s.Body = v
	return s
}

func (s *MoveCdsFileResponse) Validate() error {
	return dara.Validate(s)
}
