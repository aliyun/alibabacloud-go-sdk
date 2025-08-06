// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFreeLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFreeLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFreeLicenseResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFreeLicenseResponseBody) *DeleteFreeLicenseResponse
	GetBody() *DeleteFreeLicenseResponseBody
}

type DeleteFreeLicenseResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFreeLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFreeLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFreeLicenseResponse) GoString() string {
	return s.String()
}

func (s *DeleteFreeLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFreeLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFreeLicenseResponse) GetBody() *DeleteFreeLicenseResponseBody {
	return s.Body
}

func (s *DeleteFreeLicenseResponse) SetHeaders(v map[string]*string) *DeleteFreeLicenseResponse {
	s.Headers = v
	return s
}

func (s *DeleteFreeLicenseResponse) SetStatusCode(v int32) *DeleteFreeLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFreeLicenseResponse) SetBody(v *DeleteFreeLicenseResponseBody) *DeleteFreeLicenseResponse {
	s.Body = v
	return s
}

func (s *DeleteFreeLicenseResponse) Validate() error {
	return dara.Validate(s)
}
