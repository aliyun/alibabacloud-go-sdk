// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePostPaidInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleasePostPaidInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleasePostPaidInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ReleasePostPaidInstanceResponseBody) *ReleasePostPaidInstanceResponse
	GetBody() *ReleasePostPaidInstanceResponseBody
}

type ReleasePostPaidInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleasePostPaidInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleasePostPaidInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleasePostPaidInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleasePostPaidInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleasePostPaidInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleasePostPaidInstanceResponse) GetBody() *ReleasePostPaidInstanceResponseBody {
	return s.Body
}

func (s *ReleasePostPaidInstanceResponse) SetHeaders(v map[string]*string) *ReleasePostPaidInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleasePostPaidInstanceResponse) SetStatusCode(v int32) *ReleasePostPaidInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleasePostPaidInstanceResponse) SetBody(v *ReleasePostPaidInstanceResponseBody) *ReleasePostPaidInstanceResponse {
	s.Body = v
	return s
}

func (s *ReleasePostPaidInstanceResponse) Validate() error {
	return dara.Validate(s)
}
