// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseKmsInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseKmsInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseKmsInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseKmsInstanceResponseBody) *ReleaseKmsInstanceResponse
	GetBody() *ReleaseKmsInstanceResponseBody
}

type ReleaseKmsInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseKmsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseKmsInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseKmsInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseKmsInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseKmsInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseKmsInstanceResponse) GetBody() *ReleaseKmsInstanceResponseBody {
	return s.Body
}

func (s *ReleaseKmsInstanceResponse) SetHeaders(v map[string]*string) *ReleaseKmsInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseKmsInstanceResponse) SetStatusCode(v int32) *ReleaseKmsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseKmsInstanceResponse) SetBody(v *ReleaseKmsInstanceResponseBody) *ReleaseKmsInstanceResponse {
	s.Body = v
	return s
}

func (s *ReleaseKmsInstanceResponse) Validate() error {
	return dara.Validate(s)
}
