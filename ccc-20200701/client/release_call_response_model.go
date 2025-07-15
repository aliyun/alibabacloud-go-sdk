// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseCallResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseCallResponseBody) *ReleaseCallResponse
	GetBody() *ReleaseCallResponseBody
}

type ReleaseCallResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseCallResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseCallResponse) GoString() string {
	return s.String()
}

func (s *ReleaseCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseCallResponse) GetBody() *ReleaseCallResponseBody {
	return s.Body
}

func (s *ReleaseCallResponse) SetHeaders(v map[string]*string) *ReleaseCallResponse {
	s.Headers = v
	return s
}

func (s *ReleaseCallResponse) SetStatusCode(v int32) *ReleaseCallResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseCallResponse) SetBody(v *ReleaseCallResponseBody) *ReleaseCallResponse {
	s.Body = v
	return s
}

func (s *ReleaseCallResponse) Validate() error {
	return dara.Validate(s)
}
