// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseValueAddedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseValueAddedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseValueAddedResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseValueAddedResponseBody) *ReleaseValueAddedResponse
	GetBody() *ReleaseValueAddedResponseBody
}

type ReleaseValueAddedResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseValueAddedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseValueAddedResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseValueAddedResponse) GoString() string {
	return s.String()
}

func (s *ReleaseValueAddedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseValueAddedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseValueAddedResponse) GetBody() *ReleaseValueAddedResponseBody {
	return s.Body
}

func (s *ReleaseValueAddedResponse) SetHeaders(v map[string]*string) *ReleaseValueAddedResponse {
	s.Headers = v
	return s
}

func (s *ReleaseValueAddedResponse) SetStatusCode(v int32) *ReleaseValueAddedResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseValueAddedResponse) SetBody(v *ReleaseValueAddedResponseBody) *ReleaseValueAddedResponse {
	s.Body = v
	return s
}

func (s *ReleaseValueAddedResponse) Validate() error {
	return dara.Validate(s)
}
