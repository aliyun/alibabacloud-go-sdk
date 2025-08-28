// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSeekVideoFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SeekVideoFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SeekVideoFileResponse
	GetStatusCode() *int32
	SetBody(v *SeekVideoFileResponseBody) *SeekVideoFileResponse
	GetBody() *SeekVideoFileResponseBody
}

type SeekVideoFileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SeekVideoFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SeekVideoFileResponse) String() string {
	return dara.Prettify(s)
}

func (s SeekVideoFileResponse) GoString() string {
	return s.String()
}

func (s *SeekVideoFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SeekVideoFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SeekVideoFileResponse) GetBody() *SeekVideoFileResponseBody {
	return s.Body
}

func (s *SeekVideoFileResponse) SetHeaders(v map[string]*string) *SeekVideoFileResponse {
	s.Headers = v
	return s
}

func (s *SeekVideoFileResponse) SetStatusCode(v int32) *SeekVideoFileResponse {
	s.StatusCode = &v
	return s
}

func (s *SeekVideoFileResponse) SetBody(v *SeekVideoFileResponseBody) *SeekVideoFileResponse {
	s.Body = v
	return s
}

func (s *SeekVideoFileResponse) Validate() error {
	return dara.Validate(s)
}
